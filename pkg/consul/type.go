package internal_consul

import (
	capi "github.com/hashicorp/consul/api"
)

type KVClient struct {
	kv *capi.KV
}
