// Package backend provides the portion of top-level swapd instance
// management that is shared by both the maker and the taker.
package backend

import (
	"context"
	"sync"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/libp2p/go-libp2p/core/peer"

	rcommon "github.com/athanorlabs/go-relayer/common"
	rnet "github.com/athanorlabs/go-relayer/net"

	"github.com/athanorlabs/atomic-swap/common"
	"github.com/athanorlabs/atomic-swap/common/types"
	mcrypto "github.com/athanorlabs/atomic-swap/crypto/monero"
	"github.com/athanorlabs/atomic-swap/db"
	contracts "github.com/athanorlabs/atomic-swap/ethereum"
	"github.com/athanorlabs/atomic-swap/ethereum/extethclient"
	"github.com/athanorlabs/atomic-swap/monero"
	"github.com/athanorlabs/atomic-swap/protocol/swap"
	"github.com/athanorlabs/atomic-swap/protocol/txsender"
)

// MessageSender is implemented by a Host
type MessageSender interface {
	SendSwapMessage(common.Message, types.Hash) error
	CloseProtocolStream(id types.Hash)
}

// RelayerHost contains required network functionality for discovering
// and messaging relayers.
type RelayerHost interface {
	Discover(time.Duration) ([]peer.ID, error)
	SubmitTransaction(who peer.ID, msg *rnet.TransactionRequest) (*rcommon.SubmitTransactionResponse, error)
}

// RecoveryDB is implemented by *db.RecoveryDB
type RecoveryDB interface {
	PutContractSwapInfo(id types.Hash, info *db.EthereumSwapInfo) error
	GetContractSwapInfo(id types.Hash) (*db.EthereumSwapInfo, error)
	PutSwapPrivateKey(id types.Hash, keys *mcrypto.PrivateSpendKey) error
	GetSwapPrivateKey(id types.Hash) (*mcrypto.PrivateSpendKey, error)
	PutCounterpartySwapPrivateKey(id types.Hash, keys *mcrypto.PrivateSpendKey) error
	GetCounterpartySwapPrivateKey(id types.Hash) (*mcrypto.PrivateSpendKey, error)
	PutSwapRelayerInfo(id types.Hash, info *types.OfferExtra) error
	GetSwapRelayerInfo(id types.Hash) (*types.OfferExtra, error)
	PutCounterpartySwapKeys(id types.Hash, sk *mcrypto.PublicKey, vk *mcrypto.PrivateViewKey) error
	GetCounterpartySwapKeys(id types.Hash) (*mcrypto.PublicKey, *mcrypto.PrivateViewKey, error)
	DeleteSwap(id types.Hash) error
}

// Backend provides an interface for both the XMRTaker and XMRMaker into the Monero/Ethereum chains.
// It also interfaces with the network layer.
type Backend interface {
	XMRClient() monero.WalletClient
	ETHClient() extethclient.EthClient
	MessageSender

	RecoveryDB() RecoveryDB

	// NewTxSender creates a new transaction sender, called per-swap
	NewTxSender(asset ethcommon.Address, erc20Contract *contracts.IERC20) (txsender.Sender, error)

	// helpers
	NewSwapFactory(addr ethcommon.Address) (*contracts.SwapFactory, error)

	// getters
	Ctx() context.Context
	Env() common.Environment
	SwapManager() swap.Manager
	Contract() *contracts.SwapFactory
	ContractAddr() ethcommon.Address
	SwapTimeout() time.Duration
	XMRDepositAddress(offerID *types.Hash) *mcrypto.Address

	// setters
	SetSwapTimeout(timeout time.Duration)
	SetXMRDepositAddress(*mcrypto.Address, types.Hash)
	ClearXMRDepositAddress(types.Hash)

	// relayer functions
	DiscoverRelayers() ([]peer.ID, error)
	SubmitTransactionToRelayer(peer.ID, *rcommon.SubmitTransactionRequest) (*rcommon.SubmitTransactionResponse, error)
}

type backend struct {
	ctx         context.Context
	env         common.Environment
	swapManager swap.Manager
	recoveryDB  RecoveryDB

	// wallet/node endpoints
	moneroWallet monero.WalletClient
	ethClient    extethclient.EthClient

	// Monero deposit address. When the XMR xmrtaker has transferBack set to
	// true (default), claimed funds are swept back to the primary XMR wallet
	// address used by swapd. This sweep destination address can be overridden
	// on a per-swap basis, by setting an address indexed by the offerID/swapID
	// in the map below.
	perSwapXMRDepositAddrRWMu sync.RWMutex
	perSwapXMRDepositAddr     map[types.Hash]*mcrypto.Address

	// swap contract
	contract     *contracts.SwapFactory
	contractAddr ethcommon.Address
	swapTimeout  time.Duration

	// network interface
	MessageSender

	// relayer network interface
	rnet RelayerHost
}

// Config is the config for the Backend
type Config struct {
	Ctx            context.Context
	MoneroClient   monero.WalletClient
	EthereumClient extethclient.EthClient
	Environment    common.Environment

	SwapContract        *contracts.SwapFactory
	SwapContractAddress ethcommon.Address

	SwapManager swap.Manager

	RecoveryDB RecoveryDB

	Net         MessageSender
	RelayerHost RelayerHost
}

// NewBackend returns a new Backend
func NewBackend(cfg *Config) (Backend, error) {
	if cfg.SwapContract == nil || (cfg.SwapContractAddress == ethcommon.Address{}) {
		return nil, errNilSwapContractOrAddress
	}

	return &backend{
		ctx:                   cfg.Ctx,
		env:                   cfg.Environment,
		moneroWallet:          cfg.MoneroClient,
		ethClient:             cfg.EthereumClient,
		contract:              cfg.SwapContract,
		contractAddr:          cfg.SwapContractAddress,
		swapManager:           cfg.SwapManager,
		swapTimeout:           common.SwapTimeoutFromEnv(cfg.Environment),
		MessageSender:         cfg.Net,
		perSwapXMRDepositAddr: make(map[types.Hash]*mcrypto.Address),
		recoveryDB:            cfg.RecoveryDB,
		rnet:                  cfg.RelayerHost,
	}, nil
}

func (b *backend) XMRClient() monero.WalletClient {
	return b.moneroWallet
}

func (b *backend) ETHClient() extethclient.EthClient {
	return b.ethClient
}

func (b *backend) NewTxSender(asset ethcommon.Address, erc20Contract *contracts.IERC20) (txsender.Sender, error) {
	if !b.ethClient.HasPrivateKey() {
		return txsender.NewExternalSender(b.ctx, b.env, b.ethClient.Raw(), b.contractAddr, asset)
	}

	return txsender.NewSenderWithPrivateKey(b.ctx, b.ETHClient(), b.contract, erc20Contract), nil
}

func (b *backend) RecoveryDB() RecoveryDB {
	return b.recoveryDB
}

func (b *backend) Contract() *contracts.SwapFactory {
	return b.contract
}

func (b *backend) ContractAddr() ethcommon.Address {
	return b.contractAddr
}

func (b *backend) Ctx() context.Context {
	return b.ctx
}

func (b *backend) Env() common.Environment {
	return b.env
}

func (b *backend) SwapManager() swap.Manager {
	return b.swapManager
}

func (b *backend) SwapTimeout() time.Duration {
	return b.swapTimeout
}

// SetSwapTimeout sets the duration between the swap being initiated on-chain and the timeout t0,
// and the duration between t0 and t1.
func (b *backend) SetSwapTimeout(timeout time.Duration) {
	b.swapTimeout = timeout
}

func (b *backend) NewSwapFactory(addr ethcommon.Address) (*contracts.SwapFactory, error) {
	return contracts.NewSwapFactory(addr, b.ethClient.Raw())
}

// XMRDepositAddress returns the per-swap override deposit address, if a
// per-swap address was set. Otherwise the primary swapd Monero wallet address
// is returned.
func (b *backend) XMRDepositAddress(offerID *types.Hash) *mcrypto.Address {
	b.perSwapXMRDepositAddrRWMu.RLock()
	defer b.perSwapXMRDepositAddrRWMu.RUnlock()

	if offerID != nil {
		addr, ok := b.perSwapXMRDepositAddr[*offerID]
		if ok {
			return addr
		}
	}

	return b.XMRClient().PrimaryAddress()
}

// SetXMRDepositAddress sets a per-swap override deposit address to use when
// sweeping funds out of the shared swap wallet. When transferBack is set
// (default), funds will be swept to this override address instead of to swap's
// primary monero wallet.
func (b *backend) SetXMRDepositAddress(addr *mcrypto.Address, offerID types.Hash) {
	b.perSwapXMRDepositAddrRWMu.Lock()
	defer b.perSwapXMRDepositAddrRWMu.Unlock()
	b.perSwapXMRDepositAddr[offerID] = addr
}

// ClearXMRDepositAddress clears the per-swap, override deposit address from the
// map if a value was set.
func (b *backend) ClearXMRDepositAddress(offerID types.Hash) {
	b.perSwapXMRDepositAddrRWMu.Lock()
	defer b.perSwapXMRDepositAddrRWMu.Unlock()
	delete(b.perSwapXMRDepositAddr, offerID)
}

func (b *backend) DiscoverRelayers() ([]peer.ID, error) {
	const defaultDiscoverTime = time.Second * 3
	return b.rnet.Discover(defaultDiscoverTime)
}

func (b *backend) SubmitTransactionToRelayer(
	to peer.ID,
	req *rcommon.SubmitTransactionRequest,
) (*rcommon.SubmitTransactionResponse, error) {
	msg := &rnet.TransactionRequest{
		SubmitTransactionRequest: *req,
	}

	return b.rnet.SubmitTransaction(to, msg)
}
