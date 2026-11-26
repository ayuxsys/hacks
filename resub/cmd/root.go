package cmd

import (
	"errors"
	"log"

	"github.com/ayuxsec/resub/internal/utils"
	"github.com/ayuxsec/resub/pkg/resub"
	"github.com/ayuxsec/resub/pkg/version"
	"github.com/spf13/cobra"
)

var Root *cobra.Command

func init() {
	cobra.EnableTraverseRunHooks = true
	cmdi := &cmdi{}
	Root = cmdi.newRootCmd()
	Root.SilenceUsage = true
	Root.SilenceErrors = true
}

func (cmdi *cmdi) newRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:     "resub <domain|file>",
		Short:   "create permutations of subdomains for fuzzing",
		RunE:    cmdi.runRoot,
		Version: version.String(),
	}
	rootCmd.Flags().StringVarP(&cmdi.identifier, "identifier", "i", "FUZZ", "identifier to replace")
	rootCmd.Flags().StringVarP(&cmdi.wordListMode, "mode", "m", "tiny", "wordlist mode to use (tiny, medium, large, huge)")
	rootCmd.Flags().StringVarP(&cmdi.wordlistPath, "wordlist", "w", "", "path to the wordlist file to use")

	return rootCmd
}

func (cmdi *cmdi) runRoot(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		cmd.Help()
		return errors.New("domain name or path is required")
	}

	domain := args[0]

	var domains []string

	if utils.IsFile(domain) {
		var err error
		domains, err = utils.FilePathToSlice(domain)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		domains = []string{domain}
	}

	resub := resub.New(domains, cmdi.identifier, cmdi.wordListMode) // todo
	resub.WordListPath = cmdi.wordlistPath
	return resub.PermutateDomains()
}
