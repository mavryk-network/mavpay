package mock

import "github.com/mavryk-network/mvgo/mavryk"

func GetRandomAddress() mavryk.Address {
	k, _ := mavryk.GenerateKey(mavryk.KeyTypeEd25519)
	return k.Address()
}
