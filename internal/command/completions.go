package command

import (
	"github.com/spf13/cobra"
)

// TODO - Dynamic version for gpm

var completionsCmd = &cobra.Command{
	Use:   "completions",
	Short: "Shell Autocompletions",
	Long:  "Generate shell autocompletions. Takes one arg, one of [zsh, bash, fish, powershell]",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		switch args[0] {
		case "zsh":
			{
				err = cmd.Root().GenZshCompletion(cmd.OutOrStdout())
			}
		case "bash":
			{
				err = cmd.Root().GenBashCompletion(cmd.OutOrStdout())
			}
		case "fish":
			{
				err = cmd.Root().GenFishCompletion(cmd.OutOrStdout(), true)
			}
		case "powershell":
			{
				err = cmd.Root().GenPowerShellCompletion(cmd.OutOrStdout())
			}
		}
		if err != nil {
			cmd.PrintErr(err)
		}
	},
}
