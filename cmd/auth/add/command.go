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

package add

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/henderiw-ai/k8sai/apis"
	docs "github.com/henderiw-ai/k8sai/internal/docs/generated/authdocs"
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
		Use:     "add PROVIDER TOKEN",
		Args:    cobra.MaximumNArgs(2),
		Short:   docs.AddShort,
		Long:    docs.AddShort + "\n" + docs.AddLong,
		Example: docs.AddExamples,
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
		return fmt.Errorf("PROVIDER and TOKEN are required, positional arguments; %d provided", len(args))
	}
	provider := args[0]
	switch provider {
	case "github":
		viper.Set(apis.ViperGithubTokenKey, args[1])
	case "gitlab":
		viper.Set(apis.ViperGitlabTokenKey, args[1])
	default:

	}
	if err := viper.WriteConfig(); err != nil {
		color.Red("Error writing config file: %s", err.Error())
		return err
	}
	return nil
}
