package command

import (
	"github.com/spf13/cobra"
	"os"
)

var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish|powershell]",
	Short: "Generate completion script",
	Long: `To load completions:

Bash:

$ source <(gpm completion bash)

# To load completions for each session, execute once:
Linux:
  $ gpm completion bash > /etc/bash_completion.d/gpm
MacOS:
  $ gpm completion bash > /usr/local/etc/bash_completion.d/gpm

Zsh:

# If shell completion is not already enabled in your environment you will need
# to enable it.  You can execute the following once:

$ echo "autoload -U compinit; compinit" >> ~/.zshrc

# To load completions for each session, execute once:
$ gpm completion zsh > "${fpath[1]}/_gpm"

# You will need to start a new shell for this setup to take effect.

Fish:

$ gpm completion fish | source

# To load completions for each session, execute once:
$ gpm completion fish > ~/.config/fish/completions/gpm.fish

Powershell:

PS> gpm completion powershell | Out-String | Invoke-Expression

# To load completions for every new session, run:
PS> gpm completion powershell > gpm.ps1
# and source this file from your powershell profile.
`,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		switch args[0] {
		case "bash":
			err = cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			err = cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			err = cmd.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			err = cmd.Root().GenPowerShellCompletion(os.Stdout)
		}
		if err != nil {
			cmd.PrintErr(err)
		}
	},
}
