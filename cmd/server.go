package cmd

import (
	"os"

	"github.com/devShahriar/xm/internal/adapters/db"
	"github.com/devShahriar/xm/internal/adapters/http"
	"github.com/devShahriar/xm/internal/config"
	"github.com/devShahriar/xm/internal/usecase"
	"github.com/spf13/cobra"
)

var server = &cobra.Command{
	Use:   "server",
	Short: "server",

	Run: func(cmd *cobra.Command, args []string) {

		conf := config.GetCmdConfig()
		conf.ReadConfig()
		//Create Database if doesn't exist
		db.CreateDBIfNotExists()

		dbInstance := db.NewDBInstance()
		dbInstance.RunMigrations()

		companyUsecase := usecase.NewCompanyUsecase(dbInstance)

		server := http.NewServer(companyUsecase)

		server.RegisterRoutes()
		server.Start()

	},
}

func init() {
	rootCmd.AddCommand(server)
	registerFlags(server)
}

func registerFlags(c *cobra.Command) {

	conf := config.GetCmdConfig()
	c.Flags().StringVarP(&conf.ConfigPath, "config", "c", os.Getenv("CONFIG_PATH"), "Config path for the general config in yaml. Ex: DB conf etc")
}
