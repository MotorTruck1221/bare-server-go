package cli

import (
    "github.com/spf13/cobra"
    //tea "github.com/charmbracelet/bubbletea"
    //"github.com/charmbracelet/lipgloss"
    //"fmt"
    "os"
)

var rootCmd = &cobra.Command{
    Use:   "go-bare",
    Short: "The CLI for the bare go project",
    Long: `Go Bare is a TOMPHTTP Spec compliant server written in Go.`,
}

func Init() {
    err := rootCmd.Execute()
    if err != nil { os.Exit(1) }
}
