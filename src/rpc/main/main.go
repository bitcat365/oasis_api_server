package main

import (
	"encoding/base64"
	"fmt"
	"github.com/oasisprotocol/oasis-core/go/common/crypto/address"
	common_signature "github.com/oasisprotocol/oasis-core/go/common/crypto/signature"
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
	var pubKey common_signature.PublicKey
	err := pubKey.UnmarshalText([]byte("CVzqFIADD2Ed0khGBNf4Rvh7vSNtrL1ULTkWYQszDpc="))
	if err != nil {

	}
	fmt.Println(NewAddress(pubKey))

	b, err := base64.StdEncoding.DecodeString(string("AKFreVYqiVbzFffBiBri7pPb/Av4"))
	if err != nil {
	}
	fmt.Println(b)

	var addr address.Address
	err1 := addr.UnmarshalBech32(AddressBech32HRP, []byte("oasis1qzskk72k92y4duc47lqcsxhza6fahlqtlqyesw0p"))
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(addr)
}

func NewAddress(pk common_signature.PublicKey) (a Address) {
	pkData, _ := pk.MarshalBinary()
	return (Address)(address.NewAddress(AddressV0Context, pkData))
}
