package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "webgames",
	Short: "WebGames is a browser-based gaming platform that brings the classic console experience directly to your web engine",
	Long: `WebGames

A high-performance, browser-based gaming platform that brings the classic console 
experience directly to your web engine. No downloads, no plugins—just select your 
console, pick your game, and play.`,
}

func Execute() {
	if len(os.Args) == 1 {
		cmd, _, err := rootCmd.Find(os.Args[1:])
		if err != nil || cmd.Args == nil {
			args := append([]string{"server"}, os.Args[1:]...)
			rootCmd.SetArgs(args)
		}
	}
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
