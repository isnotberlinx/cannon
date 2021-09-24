//go:build mips

package oracle

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
)

func Preimage(hash common.Hash) []byte {
	f, err := os.Open(fmt.Sprintf("/tmp/eth/%s", hash))
	if err != nil {
		panic("missing preimage")
	}

	defer f.Close()
	ret, err := ioutil.ReadAll(f)
	if err != nil {
		panic("preimage read failed")
	}
	return ret
}

// these are stubs in embedded world
func PrefetchStorage(blockNumber *big.Int, addr common.Address, skey common.Hash) {}
func PrefetchAccount(blockNumber *big.Int, addr common.Address)                   {}
func PrefetchCode(blockNumber *big.Int, addrHash common.Hash)                     {}