package cmd

import (
	"github.com/Matt-Gleich/statuser"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "Quick-Git",
	Short: "Git for busy people",
	Long: `
⚡ An advanced version of the git cli. Speed and amazing UX.
🐙 GitHub Repository: https://github.com/Matt-Gleich/Quick-Git

   ____       _      _           ___ _ _
  /___ \_   _(_) ___| | __      / _ (_) |_
 //  / / | | | |/ __| |/ /____ / /_\/ | __|
/ \_/ /| |_| | | (__|   <_____/ /_\\| | |_
\___,_\ \__,_|_|\___|_|\_\    \____/|_|\__|
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		statuser.Error("Failed to run root cmd", err, 0)
	}
}
