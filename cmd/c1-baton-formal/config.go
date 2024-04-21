package main

import (
	"context"
	"fmt"

	"github.com/conductorone/baton-sdk/pkg/cli"
	"github.com/spf13/cobra"
)

// config defines the external configuration required for the connector to run.
type config struct {
	cli.BaseConfig `mapstructure:",squash"` // Puts the base config options in the same place as the connector options
	FormalAPIKey   string                   `mapstructure:"formal-api-key"`
}

// validateConfig is run after the configuration is loaded, and should return an error if it isn't valid.
func validateConfig(ctx context.Context, cfg *config) error {
	if cfg.FormalAPIKey == "" {
		return fmt.Errorf("missing Formal API key")
	}
	return nil
}

func parseCmd(cmd *cobra.Command) {
	cmd.PersistentFlags().String("formal-api-key", "", "Your Formal API KEY ($BATON_FORMAL_API_KEY)")
}
