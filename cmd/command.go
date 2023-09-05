/*
Copyright 2023 Nokia.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"context"
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/henderiw-ai/k8sai/cmd/auth"
	"github.com/henderiw-ai/k8sai/cmd/download"
	"github.com/henderiw-ai/k8sai/cmd/generate"
	"github.com/henderiw-ai/k8sai/cmd/repo"
	"github.com/henderiw-ai/k8sai/internal/docs/generated/overview"
	"github.com/henderiw-ai/k8sai/pkg/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	defaultConfigFileSubDir = "k8sai"
	defaultConfigFileName   = "k8sai.yaml"
)

var (
	configFile string
)

func GetMain(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:          "k8sai",
		Short:        overview.CliShort,
		Long:         overview.CliLong,
		SilenceUsage: true,
		// We handle all errors in main after return from cobra so we can
		// adjust the error message coming from libraries
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			h, err := cmd.Flags().GetBool("help")
			if err != nil {
				return err
			}
			if h {
				return cmd.Help()
			}
			return cmd.Usage()
		},
	}
	// ensure the viper config directory exists
	cobra.CheckErr(util.EnsureDir(xdg.ConfigHome, defaultConfigFileSubDir))
	// initialize viper settings
	initConfig()

	cmd.AddCommand(generate.GetCommand(ctx))
	cmd.AddCommand(auth.GetCommand(ctx))
	cmd.AddCommand(download.NewCommand(ctx))
	cmd.AddCommand(repo.GetCommand(ctx))
	cmd.AddCommand(versionCmd)

	return cmd
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if configFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(configFile)
	} else {
		viper.AddConfigPath(filepath.Join(xdg.ConfigHome, defaultConfigFileSubDir))
		viper.SetConfigType("yaml")
		viper.SetConfigName("k8sai")

		_ = viper.SafeWriteConfig()
	}

	//viper.Set("kubecontext", kubecontext)
	//viper.Set("kubeconfig", kubeconfig)

	viper.SetEnvPrefix("K8SAI")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		_ = 1
	}
}
