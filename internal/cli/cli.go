package cli

import (
    "github.com/spf13/cobra"
    "os"
)

var rootCmd = &cobra.Command{
    Use:   "bare-go",
    Short: "The CLI for the bare go project",
    Long: `Bare Go is a TOMPHTTP Spec compliant server written in Go.`,
}

func Init() {
    err := rootCmd.Execute()
    if err != nil { os.Exit(1) }
}
