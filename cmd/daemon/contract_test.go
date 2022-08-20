package main

import (
	"testing"

	"github.com/athanorlabs/atomic-swap/common"
	"github.com/athanorlabs/atomic-swap/tests"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestGetOrDeploySwapFactory(t *testing.T) {
	pk := tests.GetTakerTestKey(t)
	ec, chainID := tests.NewEthClient(t)
	tmpDir := t.TempDir()

	_, addr, err := getOrDeploySwapFactory(ethcommon.Address{},
		common.Development,
		tmpDir,
		chainID,
		pk,
		ec,
	)
	require.NoError(t, err)
	t.Log(addr)

	_, addr2, err := getOrDeploySwapFactory(addr,
		common.Development,
		tmpDir,
		chainID,
		pk,
		ec,
	)
	require.NoError(t, err)
	require.Equal(t, addr, addr2)
}
