package signer_engines

import (
	"context"
	"errors"
	"net/url"

	"github.com/mavryk-network/mavpay/constants"
	"github.com/mavryk-network/mvgo/codec"
	"github.com/mavryk-network/mvgo/mavryk"
	"github.com/mavryk-network/mvgo/signer"
	"github.com/mavryk-network/mvgo/signer/remote"
)

type RemoteSignerSpecs struct {
	Pkh string `json:"pkh"`
	Url string `json:"url"`
}

type RemoteSigner struct {
	Address mavryk.Address
	Remote  *remote.RemoteSigner
	Key     mavryk.Key
}

func InitRemoteSignerFromSpecs(specs RemoteSignerSpecs) (*RemoteSigner, error) {
	return InitRemoteSigner(specs.Pkh, specs.Url)
}

func InitRemoteSigner(address string, remoteUrl string) (*RemoteSigner, error) {
	if _, err := url.Parse(remoteUrl); err != nil {
		return nil, errors.Join(constants.ErrSignerLoadFailed, err)
	}
	rs, err := remote.New(remoteUrl, nil)
	if err != nil {
		return nil, errors.Join(constants.ErrSignerLoadFailed, err)
	}
	addr, err := mavryk.ParseAddress(address)
	if err != nil {
		return nil, errors.Join(constants.ErrSignerLoadFailed, err)
	}

	key, err := rs.GetKey(context.Background(), addr)
	if err != nil {
		return nil, errors.Join(constants.ErrSignerLoadFailed, err)
	}

	return &RemoteSigner{
		Address: addr,
		Remote:  rs,
		Key:     key,
	}, nil
}

func (remoteSigner *RemoteSigner) GetId() string {
	return "RemoteSigner"
}

func (remoteSigner *RemoteSigner) GetPKH() mavryk.Address {
	return remoteSigner.Address
}

func (remoteSigner *RemoteSigner) GetKey() mavryk.Key {
	return remoteSigner.Key
}

func (remoteSigner *RemoteSigner) GetSigner() signer.Signer {
	return remoteSigner.Remote
}

func (remoteSigner *RemoteSigner) Sign(op *codec.Op) error {
	sig, err := remoteSigner.Remote.SignOperation(context.Background(), remoteSigner.Address, op)
	if err != nil {
		return err
	}
	op.WithSignature(sig)
	return nil
}
