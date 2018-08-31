package types

import (
	"github.com/DSiSc/craft/signature/keypair"
)

type Account struct {
	PrivateKey keypair.PrivateKey
	PublicKey  keypair.PublicKey
	Address    Address
	SigScheme  keypair.SignatureScheme
}

func (*Account) PrivKey() keypair.PrivateKey {
	return nil
}

//get signer's public key

func (*Account) PubKey() keypair.PublicKey {
	return nil
}

func (*Account) Scheme() keypair.SignatureScheme {
	//var temp keypair.SignatureScheme
	var byt byte = 'a'
	return keypair.SignatureScheme(byt)
}
