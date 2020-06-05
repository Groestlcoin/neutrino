module github.com/lightninglabs/neutrino

require (
	github.com/btcsuite/btcd v0.20.1-beta
	github.com/btcsuite/btclog v0.0.0-20170628155309-84c8d2346e9f
	github.com/btcsuite/btcutil v0.0.0-20190425235716-9e5f4b9a998d
	github.com/btcsuite/btcwallet/wallet/txauthor v1.0.0
	github.com/btcsuite/btcwallet/walletdb v1.2.0
	github.com/btcsuite/btcwallet/wtxmgr v1.0.0
	github.com/davecgh/go-spew v1.1.1
	github.com/lightningnetwork/lnd/queue v1.0.1
)

go 1.13

replace (
	github.com/btcsuite/btcd => github.com/Groestlcoin/grsd v0.20.1-grs
	github.com/btcsuite/btcutil => github.com/Groestlcoin/grsutil v0.5.0-grs3
	github.com/btcsuite/btcwallet/wallet/txauthor => github.com/Groestlcoin/grswallet/wallet/txauthor v1.0.0-grs
	github.com/btcsuite/btcwallet/walletdb => github.com/Groestlcoin/grswallet/walletdb v1.2.0-grs
	github.com/btcsuite/btcwallet/wtxmgr => github.com/Groestlcoin/grswallet/wtxmgr v1.0.0-grs
)
