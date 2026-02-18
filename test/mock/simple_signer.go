package mock

import (
	signer_engines "github.com/mavryk-network/mavpay/engines/signer"
	"github.com/mavryk-network/gomavryk/mavryk"
)

func InitSimpleSigner() *signer_engines.InMemorySigner {
	key, _ := mavryk.GenerateKey(mavryk.KeyTypeEd25519)
	encoded := key.String()

	result, _ := signer_engines.InitInMemorySigner(encoded)
	return result
}
