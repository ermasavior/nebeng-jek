package internal_consul

import (
	capi "github.com/hashicorp/consul/api"
)

func NewConsulKVClient(address string) (*KVClient, error) {
	client, err := capi.NewClient(&capi.Config{
		Address: address,
	})
	if err != nil {
		return nil, err
	}

	return &KVClient{
		kv: client.KV(),
	}, nil
}
