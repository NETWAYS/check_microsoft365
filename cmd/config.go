package cmd

import (
	"github.com/NETWAYS/check_microsoft365/internal/client"
	"os"
	"reflect"
)

type Config struct {
	TenantId       string `env:"AZURE_TENANT_ID"`
	ClientId       string `env:"AZURE_CLIENT_ID"`
	ClientSecret   string `env:"AZURE_CLIENT_SECRET"`
	Scope          string
	ServiceName    string
	ServiceNames   []string
	StateOverrides []string
}

var cliConfig Config

func (c *Config) Client() *client.Client {
	return client.NewClient(c.TenantId, c.ClientId, c.ClientSecret, c.Scope)
}

func (c *Config) LoadEnv() {
	v := reflect.ValueOf(c).Elem()
	el := reflect.TypeOf(c).Elem()

	for i := 0; i < el.NumField(); i++ {
		envKey := el.Field(i).Tag.Get("env")
		if envKey != "" {
			currentValue := v.Field(i).String()

			if value := os.Getenv(envKey); value != "" && currentValue == "" {
				v.Field(i).SetString(value)
			}
		}
	}
}
