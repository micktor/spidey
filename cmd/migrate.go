package cmd

import "github.com/spf13/cobra"
import "github.com/micktor/spidey/internal/repository"

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate tables",
	Long:  `migrate up/down db tables`,
	Run:   runMigration,
}

func init() {
	rootCmd.AddCommand(migrateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runMigration(cmd *cobra.Command, args []string) {
	dbClient := repository.Connect(&conf)
	repository.Migrate(dbClient)

	return
}
