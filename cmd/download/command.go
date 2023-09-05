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

package download

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/henderiw-ai/k8sai/apis"
	docs "github.com/henderiw-ai/k8sai/internal/docs/generated/downloaddocs"
	"github.com/henderiw-ai/k8sai/pkg/git"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewCommand(ctx context.Context) *cobra.Command {
	return NewRunner(ctx).Command
}

// NewRunner returns a command runner.
func NewRunner(ctx context.Context) *Runner {
	r := &Runner{
		ctx: ctx,
	}
	c := &cobra.Command{
		Use:     "download",
		Args:    cobra.MaximumNArgs(2),
		Short:   docs.DownloadShort,
		Long:    docs.DownloadShort + "\n" + docs.DownloadLong,
		Example: docs.DownloadExamples,
		RunE:    r.runE,
	}

	r.Command = c
	return r
}

type Runner struct {
	Command *cobra.Command
	ctx     context.Context
}

func (r *Runner) runE(c *cobra.Command, args []string) error {
	gitProviders := r.ctx.Value("gitProviders")
	gitRegistry, ok := gitProviders.(git.Registry)
	if !ok {
		return fmt.Errorf("cannot get git registry")
	}

	for _, repo := range viper.GetStringSlice(apis.ViperReposKey) {
		apiRepo := apis.Repo{}
		if err := json.Unmarshal([]byte(repo), &apiRepo); err != nil {
			return err
		}
		parsedRepo, err := git.ParseRepoURL(apiRepo.RepoURL)
		if err != nil {
			return err
		}
		g, err := gitRegistry.NewGitOfProvider(parsedRepo.Host, apis.BaseFilePath)
		if err != nil {
			return err
		}
		if err := g.GetContent(r.ctx, parsedRepo.Owner, parsedRepo.Repo, apiRepo.RepoPath); err != nil {
			return err
		}
	}

	return nil
}
