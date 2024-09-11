package signer_engines

import (
	"errors"

	"github.com/mavryk-network/mavpay/constants"
	"github.com/mavryk-network/mvgo/codec"
	"github.com/mavryk-network/mvgo/mavryk"
	"github.com/mavryk-network/mvgo/signer"
)

type InMemorySigner struct {
	Key mavryk.PrivateKey
}

func InitInMemorySigner(key string) (*InMemorySigner, error) {
	tkey, err := mavryk.ParsePrivateKey(key)
	if err != nil {
		return nil, errors.Join(constants.ErrSignerLoadFailed, err)
	}
	return &InMemorySigner{
		Key: tkey,
	}, nil
}

func (inMemSigner *InMemorySigner) GetId() string {
	return "InMemorySigner"
}

func (inMemSigner *InMemorySigner) GetPKH() mavryk.Address {
	return inMemSigner.Key.Address()
}

func (inMemSigner *InMemorySigner) GetKey() mavryk.Key {
	return inMemSigner.Key.Public()
}

func (inMemSigner *InMemorySigner) Sign(op *codec.Op) error {
	if err := op.Sign(inMemSigner.Key); err != nil {
		return err
	}
	return nil
}

func (inMemSigner *InMemorySigner) GetSigner() signer.Signer {
	return signer.NewFromKey(inMemSigner.Key)
}
