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

package delete

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fatih/color"
	"github.com/henderiw-ai/k8sai/apis"
	docs "github.com/henderiw-ai/k8sai/internal/docs/generated/repodocs"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewCommand(ctx context.Context) *cobra.Command {
	return NewRunner(ctx).Command
}

// NewRunner returns a command runner.
func NewRunner(ctx context.Context) *Runner {
	r := &Runner{
		Ctx: ctx,
	}
	c := &cobra.Command{
		Use:     "delete REPO_URL REPO_PATH",
		Args:    cobra.MaximumNArgs(2),
		Short:   docs.DeleteShort,
		Long:    docs.DeleteShort + "\n" + docs.DeleteLong,
		Example: docs.DeleteExamples,
		RunE:    r.runE,
	}

	r.Command = c
	return r
}

type Runner struct {
	Command *cobra.Command
	Ctx     context.Context
}

func (r *Runner) runE(c *cobra.Command, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("REPO_URL and REPO_PATH are required, positional arguments; %d provided", len(args))
	}

	repos := viper.GetStringSlice(apis.ViperReposKey)
	for idx, repo := range repos {
		r := apis.Repo{}
		if err := json.Unmarshal([]byte(repo), &r); err != nil {
			return err
		}
		if r.RepoURL == args[0] && r.RepoPath == args[1] {
			repos = append(repos[:(idx)], repos[idx+1:]...)
			viper.Set(apis.ViperReposKey, repos)
			if err := viper.WriteConfig(); err != nil {
				color.Red("Error writing config file: %s", err.Error())
				return err
			}
			break
		}
	}

	return nil
}
