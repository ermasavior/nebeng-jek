package internal_consul

import (
	capi "github.com/hashicorp/consul/api"
)

func (client *KVClient) GetKeyValue(key string, options *capi.QueryOptions) ([]byte, error) {
	pair, _, err := client.kv.Get(key, options)
	if err != nil {
		return nil, err
	}

	return pair.Value, nil
}
