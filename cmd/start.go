/*
Copyright © 2024 Aza M mukhamejanov.aza@gmail.com
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/creative-cranels/smart-search/configs"
	"github.com/creative-cranels/smart-search/internal/server"
	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
	"github.com/spf13/cobra"
)

var targetDB, searchProtocol, managerProtocol string

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the backend server",
	Long:  `how to use flags in order to use different types of db, protocols`, //TODO: add more detailed description
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("starting server...")

		srvConfig := server.ServerConfig{
			TargetDB:        server.TargetDBType(targetDB),
			SearchProtocol:  server.SearchProtocolType(searchProtocol),
			ManagerProtocol: server.ManagerProtocolType(managerProtocol),
		}

		if err := godotenv.Load(".env"); err != nil {
			log.Fatal("Error loading .env file:", err)
		}

		var cfg configs.Config
		if err := envconfig.Process(cmd.Context(), &cfg); err != nil {
			log.Fatal("Error processing envconfig:", err)
		}

		server.Start(srvConfig, &cfg)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	startCmd.Flags().StringVar(&targetDB, "db", "postgresql", "Choose database: postgresql, mongodb, elasticsearch")
	startCmd.Flags().StringVar(&searchProtocol, "cli", "rest", "Choose protocol: rest, grpc")
	startCmd.Flags().StringVar(&managerProtocol, "manager", "rest", "Choose protocol: rest, grpc")
}
