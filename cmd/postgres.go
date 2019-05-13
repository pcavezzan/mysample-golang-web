package cmd

import (
	"github.com/pcavezzan/sample-apiserver/app/cli"
	"github.com/spf13/cobra"
)

var postgresCmd = &cobra.Command{
	Use:   "postgres",
	Short: "Postgres interaction",
	Long:  `Permit to run specific action on postgres database`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

var initSubCmd = &cobra.Command{
	Use: "init",
	Short: "Initialize postgres database",
	Long: "Create the default schema",
	Run: func(cmd *cobra.Command, args []string) {
		cli.NewInitCommand(&cfg).Execute()
	},
}


func init() {
	postgresCmd.AddCommand(initSubCmd)
	rootCmd.AddCommand(postgresCmd)
}
