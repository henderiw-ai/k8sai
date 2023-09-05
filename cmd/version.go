package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version = "0.0.0"
	commit  = "none"
	date    = "unknown"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show k8sai version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("k8sai: %s (%s), built at: %s\n", version, commit, date)
	},
}
