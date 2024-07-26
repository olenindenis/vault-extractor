package vault

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/vault-client-go"
)

type Client struct {
	client *vault.Client
}

func NewClient(host, token string) (*Client, error) {
	client, err := vault.New(
		vault.WithAddress(host),
		vault.WithRequestTimeout(10*time.Second),
	)
	if err != nil {
		return nil, fmt.Errorf("extract, vault create New instance, %w", err)
	}

	if err = client.SetToken(token); err != nil {
		return nil, fmt.Errorf("extract, vault client set token, %w", err)
	}

	return &Client{
		client,
	}, nil
}

func (v *Client) Extract(ctx context.Context, path, mountPath string) (map[string]interface{}, error) {
	s, err := v.client.Secrets.KvV2Read(ctx, path, vault.WithMountPath(mountPath))
	if err != nil {
		return nil, fmt.Errorf("extract, vault client secrets KvV2Read, %w", err)
	}

	return s.Data.Data, nil
}
