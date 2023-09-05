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

package repo

import (
	"context"

	"github.com/henderiw-ai/k8sai/cmd/repo/add"
	"github.com/henderiw-ai/k8sai/cmd/repo/delete"
	docs "github.com/henderiw-ai/k8sai/internal/docs/generated/repodocs"
	"github.com/spf13/cobra"
)

/*
go run *.go repo add "https://github.com/kubernetes/kubernetes" "api/openapi-spec/swagger.json"   
go run *.go repo add "https://github.com/nokia/k8s-ipam" "config/crd/bases"   
go run *.go repo add "https://github.com/GoogleCloudPlatform/k8s-config-connector" "crds"   
*/

func GetCommand(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "repo",
		Short: docs.RepoShort,
		Long:  docs.RepoLong,
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
		add.NewCommand(ctx),
		delete.NewCommand(ctx),
	)

	return cmd
}
