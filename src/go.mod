module github.com/SimplyVC/oasis_api_server/src

go 1.13

replace (
	github.com/tendermint/tendermint => github.com/oasislabs/tendermint v0.33.4-oasis1
	golang.org/x/crypto/curve25519 => github.com/oasislabs/ed25519/extra/x25519 v0.0.0-20191022155220-a426dcc8ad5f
	golang.org/x/crypto/ed25519 => github.com/oasislabs/ed25519 v0.0.0-20191109133925-b197a691e30d
)

require (
	github.com/claudetech/ini v0.0.0-20140910072410-73e6100d9d51
	github.com/gorilla/mux v1.7.4
	github.com/mackerelio/go-osstat v0.1.0
	github.com/oasislabs/oasis-core/go v0.0.0-20200514075234-edb8515cb538
	github.com/oasisprotocol/oasis-core/go v0.0.0-20200618144736-02a945839e9b
	github.com/prometheus/common v0.9.1
	github.com/tendermint/tendermint v0.33.4
	github.com/zenazn/goji v0.9.0
	google.golang.org/genproto v0.0.0-20200313141609-30c55424f95d // indirect
	google.golang.org/grpc v1.29.1
)
