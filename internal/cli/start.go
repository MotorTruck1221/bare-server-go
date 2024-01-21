package cli

import (
    "github.com/spf13/cobra"
    
    //stuff from project
    "github.com/ruby-network/bare-server-go/internal/routes"
)

func init() {
    rootCmd.AddCommand(startCmd)
    startCmd.Flags().StringP("port", "p", "8080", "Port to listen on")
    startCmd.Flags().StringP("directory", "d", "/", "Directory to serve from")
    startCmd.Flags().StringP("host", "H", "0.0.0.0", "Host to listen on")
}

var startCmd = &cobra.Command{
    Use:   "start",
    Short: "Start the server",
    Long: `Start the server`,
    Run: func(cmd *cobra.Command, args []string) {
        routes.Init(cmd.Flag("host").Value.String(), cmd.Flag("port").Value.String(), cmd.Flag("directory").Value.String())
    },
}
