module github.com/SimplyVC/oasis_api_server/src

go 1.15

replace (
	// Fixes vulnerabilities in etcd v3.3.{10,13} (dependencies via viper).
	// Can be removed once there is a spf13/viper release with updated etcd.
	// https://github.com/spf13/viper/issues/956
	github.com/coreos/etcd => github.com/coreos/etcd v3.3.25+incompatible
	// Updates the version used in spf13/cobra (dependency via tendermint) as
	// there is no release yet with the fix. Remove once an updated release of
	// spf13/cobra exists and tendermint is updated to include it.
	// https://github.com/spf13/cobra/issues/1091
	github.com/gorilla/websocket => github.com/gorilla/websocket v1.4.2
	github.com/tendermint/tendermint => github.com/oasisprotocol/tendermint v0.34.15-oasis1
	golang.org/x/crypto/curve25519 => github.com/oasisprotocol/ed25519/extra/x25519 v0.0.0-20210127160119-f7017427c1ea
	golang.org/x/crypto/ed25519 => github.com/oasisprotocol/ed25519 v0.0.0-20210127160119-f7017427c1ea
)

require (
	github.com/claudetech/ini v0.0.0-20140910072410-73e6100d9d51
	github.com/gorilla/mux v1.7.4
	github.com/mackerelio/go-osstat v0.1.0
	github.com/oasisprotocol/oasis-core/go v0.2202.7
	github.com/prometheus/common v0.37.0
	github.com/tendermint/tendermint v0.34.21
	github.com/zenazn/goji v0.9.0
	google.golang.org/grpc v1.49.0
)
