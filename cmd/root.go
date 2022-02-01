package cmd

import (
	"github.com/NETWAYS/go-check"
	"github.com/spf13/cobra"
	"os"
)

var (
	Timeout = 30
)

var rootCmd = &cobra.Command{
	Use:   "check_microsoft365",
	Short: "Icinga check plugin to check Microsoft365",
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		go check.HandleTimeout(Timeout)
	},
	Run: Help,
}

func Execute(version string) {
	defer check.CatchPanic()

	rootCmd.Version = version
	rootCmd.VersionTemplate()

	if err := rootCmd.Execute(); err != nil {
		check.ExitError(err)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.DisableAutoGenTag = true

	rootCmd.SetHelpCommand(&cobra.Command{
		Use:    "no-help",
		Hidden: true,
	})

	pfs := rootCmd.PersistentFlags()
	pfs.StringVarP(&cliConfig.TenantId, "tenantid", "T", "",
		"The tenant id from your app registration (env:AZURE_TENANT_ID)")
	pfs.StringVarP(&cliConfig.ClientId, "clientid", "c", "",
		"The client id from your app registration (env: AZURE_CLIENT_ID)")
	pfs.StringVarP(&cliConfig.ClientSecret, "clientsecret", "s", "",
		"The client secret from your app registration (env: AZURE_CLIENT_SECRET)")
	pfs.StringVarP(&cliConfig.Scope, "scope", "S", "https://graph.microsoft.com/.default",
		"Represents the definition of a delegated permission")
	pfs.IntVarP(&Timeout, "timeout", "t", Timeout,
		"Timeout for the check")

	pfs.SortFlags = false
	rootCmd.Flags().SortFlags = false
}

func Help(cmd *cobra.Command, strings []string) {
	_ = cmd.Usage()

	os.Exit(3)
}
