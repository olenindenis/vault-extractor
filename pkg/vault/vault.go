package vault

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

type Client struct {
	client *vault.Client
}

func NewClient(ctx context.Context) (*Client, error) {
	client, err := vault.New(
		vault.WithAddress(os.Getenv("VAULT_HOST")),
		vault.WithRequestTimeout(10*time.Second),
	)
	if err != nil {
		return nil, fmt.Errorf("extract, vault create New instance, %w", err)
	}

	token := os.Getenv("VAULT_TOKEN")
	if token == "" {
		resp, err := client.Auth.AppRoleLogin(
			ctx,
			schema.AppRoleLoginRequest{
				RoleId:   os.Getenv("VAULT_APPROLE_ROLE_ID"),
				SecretId: os.Getenv("VAULT_APPROLE_SECRET_ID"),
			},
		)
		if err != nil {
			return nil, fmt.Errorf("NewClient, vault client set AppRoleLogin, %w", err)
		}
		token = resp.Auth.ClientToken
	}

	if err = client.SetToken(token); err != nil {
		return nil, fmt.Errorf("NewClient, vault client set token, %w", err)
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
