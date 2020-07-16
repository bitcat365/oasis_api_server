module github.com/SimplyVC/oasis_api_server/src

go 1.13

replace (
	github.com/tendermint/tendermint => github.com/oasisprotocol/tendermint v0.33.4-oasis2
	golang.org/x/crypto/curve25519 => github.com/oasisprotocol/ed25519/extra/x25519 v0.0.0-20200528083105-55566edd6df0
	golang.org/x/crypto/ed25519 => github.com/oasisprotocol/ed25519 v0.0.0-20200528083105-55566edd6df0
)

require (
	github.com/claudetech/ini v0.0.0-20140910072410-73e6100d9d51
	github.com/gorilla/mux v1.7.4
	github.com/mackerelio/go-osstat v0.1.0
	github.com/oasisprotocol/oasis-core/go v0.0.0-20200709155302-d52975aadf46
	github.com/prometheus/common v0.9.1
	github.com/tendermint/go-amino v0.15.0
	github.com/tendermint/tendermint v0.33.6
	github.com/zenazn/goji v0.9.0
	google.golang.org/genproto v0.0.0-20200313141609-30c55424f95d // indirect
	google.golang.org/grpc v1.29.1
)
