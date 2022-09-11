package noop

import "github.com/concourse/concourse/v7/atc/creds"

type noopFactory struct{}

func NewNoopFactory() *noopFactory {
	return &noopFactory{}
}

func (*noopFactory) NewSecrets() creds.Secrets {
	return &Noop{}
}
