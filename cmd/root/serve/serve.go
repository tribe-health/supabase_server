/*
Copyright © 2022 Travis James travis@tribehealthsolutions.com
*/
package serve

import (
	"fmt"
	"github.com/go-faster/errors"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"net/http"
	oas "supabase_server/api"
	"supabase_server/config"
	"supabase_server/ent"
	"supabase_server/pkg/controller"
	"supabase_server/pkg/infrastructure/datastore"
	"supabase_server/pkg/infrastructure/graphql"
	"supabase_server/pkg/infrastructure/router"
	"supabase_server/pkg/registry"
	"supabase_server/pkg/util/environment"
	"sync"
	"time"
)

// serveCmd represents the serve command
var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serve called")

		var loggerConfig zap.Config
		if environment.IsDev() {
			loggerConfig = zap.NewDevelopmentConfig()
		} else {
			loggerConfig = zap.NewProductionConfig()
		}
		loggerConfig.EncoderConfig.TimeKey = "timestamp"
		loggerConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)

		logger, err := loggerConfig.Build()
		if err != nil {
			log.Fatal(err)
		}

		sugar := logger.Sugar()

		wg := new(sync.WaitGroup)

		wg.Add(2)

		client := newDBClient()
		ctrl := newController(client)

		srv := graphql.NewServer(client, ctrl)

		oasServer, _ := oas.NewServer(ctrl.Api, ctrl.Api.SecurityHandler())
		httpServer := http.Server{
			Addr:    ":" + config.C.Server.ApiAddress,
			Handler: oasServer,
		}

		e := router.New(srv)

		go func() {
			defer e.Logger.Info("Server stopped")
			fmt.Println("api server:" + config.C.Server.ApiAddress)
			if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				sugar.Error(errors.Wrap(err, "http"))
			}
			wg.Done()
		}()

		go func() {
			e.Logger.Fatal(e.Start(":" + config.C.Server.Address))
			wg.Done()
		}()

		wg.Wait()
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func newDBClient() *ent.Client {
	client, err := datastore.NewClient()
	if err != nil {
		log.Fatalf("failed opening postgres client: %v", err)
	}

	return client
}

func newController(client *ent.Client) controller.Controller {
	r := registry.New(client)
	return r.NewController()
}
