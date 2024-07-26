package vault

import "context"

type NilClient struct {
}

func NewNilClient() *NilClient {
	return &NilClient{}
}

func (c *NilClient) Extract(ctx context.Context, path, mountPath string) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	m[path] = mountPath
	return m, nil
}
