/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package migrate

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"supabase_server/ent"
	"supabase_server/ent/migrate"
	"supabase_server/pkg/infrastructure/datastore"
)

// migrateCmd represents the migrate command
var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("migrate called")
		client, err := datastore.NewClient()
		if err != nil {
			log.Fatalf("failed opening postgres client: %v", err)
		}
		defer client.Close()
		createDBSchema(client)
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func createDBSchema(client *ent.Client) {
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithForeignKeys(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
