package main

import (
	"encoding/base64"
	"fmt"
	"github.com/oasisprotocol/oasis-core/go/common/crypto/address"
	"github.com/oasisprotocol/oasis-core/go/common/crypto/signature"
	tmcrypto "github.com/oasisprotocol/oasis-core/go/consensus/tendermint/crypto"
)

var (
	// AddressV0Context is the unique context for v0 staking account addresses.
	AddressV0Context = address.NewContext("oasis-core/address: staking", 0)
	// AddressBech32HRP is the unique human readable part of Bech32 encoded
	// staking account addresses.
	AddressBech32HRP = address.NewBech32HRP("oasis")
)

type Address address.Address

func main() {
	var pubKey signature.PublicKey
	err := pubKey.UnmarshalText([]byte("3IbxhcOu3j2o/0Mk5V3qXOfasEEm42pJpzuWAf8SREg="))
	if err != nil {

	}
	fmt.Println(NewAddress(pubKey))

	b, err := base64.StdEncoding.DecodeString(string("ND60v9LrslMFI46zbMsUxyeabzg="))
	if err != nil {
	}
	fmt.Println(b)

	tmAddr := tmcrypto.PublicKeyToTendermint(&pubKey).Address().String()
	fmt.Println(tmAddr)

	var addr address.Address
	err1 := addr.UnmarshalBech32(AddressBech32HRP, []byte("oasis1qpc34h9cm0wrvkrepvzhcz0mpmjmr4dv6g24yayr"))
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(addr)
}

func NewAddress(pk signature.PublicKey) (a Address) {
	pkData, _ := pk.MarshalBinary()
	return (Address)(address.NewAddress(AddressV0Context, pkData))
}
