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

package token

import (
	"context"

	"github.com/henderiw-ai/k8sai/cmd/generate/token/git"
	docs "github.com/henderiw-ai/k8sai/internal/docs/generated/generatedocs"
	"github.com/spf13/cobra"
)

func GetCommand(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "token",
		Short: docs.TokenShort,
		Long:  docs.TokenLong,
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

	cmd.AddCommand(
		git.NewCommand(ctx),
	)

	return cmd
}
