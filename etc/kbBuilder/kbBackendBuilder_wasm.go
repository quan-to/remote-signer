package kbBuilder

import (
	"github.com/quan-to/remote-signer"
	"github.com/quan-to/remote-signer/keyBackend"
)

func BuildKeyBackend() keyBackend.Backend {
	return keyBackend.MakeSaveToDiskBackend(remote_signer.PrivateKeyFolder, remote_signer.KeyPrefix)
}