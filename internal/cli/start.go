package cli

import (
    "github.com/spf13/cobra"
    
    //stuff from project
    "github.com/Ruby-Network/bare-go/internal/routes"
)

func init() {
    rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
    Use:   "start",
    Short: "Start the server",
    Long: `Start the server`,
    Run: func(cmd *cobra.Command, args []string) {
        routes.Init()
    },
}

