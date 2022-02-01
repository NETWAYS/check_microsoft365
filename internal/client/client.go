package client

import (
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	a "github.com/microsoft/kiota/authentication/go/azure"
	msgraphsdkgo "github.com/microsoftgraph/msgraph-sdk-go"
)

type Client struct {
	TenantId           string
	ClientIid          string
	ClientSecret       string
	GraphServiceClient *msgraphsdkgo.GraphServiceClient
	Scope              string
}

func NewClient(tenantId, clientId, clientSecret, scope string) *Client {
	return &Client{
		TenantId:     tenantId,
		ClientIid:    clientId,
		ClientSecret: clientSecret,
		Scope:        scope,
	}
}

func (c *Client) NewGraphServiceClient() error {
	cred, err := azidentity.NewClientSecretCredential(
		c.TenantId,
		c.ClientIid,
		c.ClientSecret,
		&azidentity.ClientSecretCredentialOptions{})

	if err != nil {
		return fmt.Errorf("error creating credentials: %v", err)
	}

	auth, err := a.NewAzureIdentityAuthenticationProviderWithScopes(cred, []string{c.Scope})
	if err != nil {
		return fmt.Errorf("error creating authentication provider: %v", err)
	}

	adapter, err := msgraphsdkgo.NewGraphRequestAdapter(auth)
	if err != nil {
		return fmt.Errorf("error creating graph adapter: %v", err)
	}

	client := msgraphsdkgo.NewGraphServiceClient(adapter)
	c.GraphServiceClient = client

	return nil
}
