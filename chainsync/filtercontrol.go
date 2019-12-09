package chainsync

import (
	"fmt"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
)

// ErrCheckpointMismatch is returned if given filter headers don't pass our
// control check.
var ErrCheckpointMismatch = fmt.Errorf("checkpoint doesn't match")

// filterHeaderCheckpoints holds a mapping from heights to filter headers for
// various heights. We use them to check whether peers are serving us the
// expected filter headers.
var filterHeaderCheckpoints = map[wire.BitcoinNet]map[uint32]*chainhash.Hash{
	// Mainnet filter header checkpoints.
	chaincfg.MainNetParams.Net: map[uint32]*chainhash.Hash{
		10000:   hashFromStr("9c7fe7042671e992908e1045e8e61e7fc74b8453604c2e841bf7c610501a5d93"),
		100000:  hashFromStr("83ac9e700fc7924da60cd63b51a44772fb40909bf624c9b8bcf62fc93d54355d"),
		200000:  hashFromStr("fe3e81ce1f343310c9560d65782d72c7c5fac3a9ce8583c3caf08f6a3fa938a5"),
		400000:  hashFromStr("4ecc73b1b3569748babd677f826b5ed8a6774815c9cfb657510061786ea8960e"),
		600000:  hashFromStr("e285d551fe98579d99c8a2ead53c98a8d148ce0bb06dc9a0acde83b65145144a"),
		800000:  hashFromStr("51fd5b3639793dfe7c6af8e47bb5601a26d3adb3fe88a2332e35cdc2f937c9b8"),
		1000000: hashFromStr("5eeee7c52b08b711e503055760cc426b09c9121e95318711b0fec46992351cdb"),
		1200000: hashFromStr("f32e13c4b23616484f18e97c55a89ba84cf2227cf7fb4f0fe0e3abbfb65f8e72"),
		1400000: hashFromStr("cc81c82045212f47768e9792865f1e4a8ce65a0e1dde5834b3e2eb1fc04deb1d"),
		1600000: hashFromStr("24bac78950d662be824d27e63b1313a9f2085a1555809fe18cb071bde734dfff"),
		1800000: hashFromStr("025ecdfba84d4462a310cbf30a77142ead8ea200b9627d312df607c600e0d0b5"),
		2000000: hashFromStr("c957a17cac1d2367e7c9042d1289cc7abd75bbac532c0b4a950fe53b7b50440b"),
		2200000: hashFromStr("2c9d448cb73ea3423830e4c7c9cb63da7a3e8741f8faf652e0f20e3fba7443b7"),
		2400000: hashFromStr("960a7908e8834b4fa4cc05f24a3c8d5a30868c7d12f8a52791dce5ca9c9cf228"),
		2600000: hashFromStr("e2d2b42d63e4000f2401840e2c9f651bce5e26bbd901a49b2d5e2d3481d758d0"),
		2800000: hashFromStr("183a2abe554d2b500f3dc8be4e1b4b05d99285c2606bbbdc3f5055d996bb1e67"),
	},

	// Testnet filter header checkpoints.
	chaincfg.TestNet3Params.Net: map[uint32]*chainhash.Hash{
		10000:   hashFromStr("f5605a9b304f3299414892e1335221e111351a5e55076e397bdac347542330f0"),
		100000:  hashFromStr("939bc97ffd0e5264175796c6b2dc0bc9fcad74cf5525fc343932ef13c7b7de5e"),
		200000:  hashFromStr("4949fbe1d2e3ca5eae253a8a3710db730656707b227fa1cc871d923c4c7f22e6"),
		400000:  hashFromStr("242ef3df03fd6328a3651b27ee0b312f3917b276ff903c277174d675db882480"),
		600000:  hashFromStr("78efa1617a33a5cde8417286640aa3ae0d1aabbd41594f6db936b34b080e21ca"),
		800000:  hashFromStr("39da2257576b7a2ad07ecc7c647cf1e54e83c24d5c2bdcad1ec7794b5c053dc7"),
		1000000: hashFromStr("528e01fcc8046fbc6915d0f5aba3c83a9b8ce59446c8bf7db56229624df1725a"),
		1200000: hashFromStr("1638a565f26c6e28f566b4c13acb24e632e4d8eaa634379504a7c919164e4929"),
		1400000: hashFromStr("4abf73b0678dcf500f9cdecf31b760d74c6380da0a8361cff64c39c1fbb17158"),
	},
}

// ControlCFHeader controls the given filter header against our list of
// checkpoints. It returns ErrCheckpointMismatch if we have a checkpoint at the
// given height, and it doesn't match.
func ControlCFHeader(params chaincfg.Params, fType wire.FilterType,
	height uint32, filterHeader *chainhash.Hash) error {

	if fType != wire.GCSFilterRegular {
		return fmt.Errorf("unsupported filter type %v", fType)
	}

	control, ok := filterHeaderCheckpoints[params.Net]
	if !ok {
		return nil
	}

	hash, ok := control[height]
	if !ok {
		return nil
	}

	if *filterHeader != *hash {
		return ErrCheckpointMismatch
	}

	return nil
}

// hashFromStr makes a chainhash.Hash from a valid hex string. If the string is
// invalid, a nil pointer will be returned.
func hashFromStr(hexStr string) *chainhash.Hash {
	hash, _ := chainhash.NewHashFromStr(hexStr)
	return hash
}
