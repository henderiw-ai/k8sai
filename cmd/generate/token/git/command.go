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

package git

import (
	"context"
	"fmt"
	"os/exec"
	"runtime"

	"github.com/fatih/color"
	docs "github.com/henderiw-ai/k8sai/internal/docs/generated/generatedocs"
	"github.com/spf13/cobra"
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
		Use:     "git PROVIDER",
		Args:    cobra.MaximumNArgs(1),
		Short:   docs.GitShort,
		Long:    docs.GitShort + "\n" + docs.GitLong,
		Example: docs.GitExamples,
		RunE:    r.runE,
	}

	r.Command = c
	return r
}

type Runner struct {
	Command  *cobra.Command
	provider string
	Ctx      context.Context
}

func (r *Runner) runE(c *cobra.Command, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("PROVIDER is required, positional arguments; %d provided", len(args))
	}
	r.provider = args[0]
	switch r.provider {
	case "github":
		r.openbrowser("https://github.com/settings/tokens")
	case "gitlab":
		r.openbrowser("https://gitlab.com/-/profile/personal_access_tokens")
	default:

	}
	return nil
}

func (r *Runner) openbrowser(url string) {
	var err error
	isGui := true
	switch runtime.GOOS {
	case "linux":
		_, err = exec.LookPath("xdg-open")
		if err != nil {
			isGui = false
		} else {
			err = exec.Command("xdg-open", url).Start()
		}
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	r.printInstructions(isGui, url)
	if err != nil {
		fmt.Println(err)
	}
}

func (r *Runner) printInstructions(isGui bool, url string) {
	fmt.Println("")
	if isGui {
		color.Green("Opening: %s to generate a key for %s", url, r.provider)
		fmt.Println("")
	} else {
		color.Green("Please open: %s to generate a key for %s", url, r.provider)
		fmt.Println("")
	}
	color.Green("Please copy the generated key and run the following command %s", fmt.Sprintf("`k8sai auth add %s <TOKEN>`", r.provider))
	fmt.Println("")
}

