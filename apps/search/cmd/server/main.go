package main

import (
	"os"

	"github.com/phanhotboy/nien-su-viet/apps/search/internal/shared/app"
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:              "search-microservice",
	Short:            "search-microservice based on DDD architecture",
	Long:             `This is a command runner or cli for api architecture in golang.`,
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {
		app.NewApp().Run()
	},
}

// @contact.name Phan Nguyen
// @contact.url https://github.com/phanhotboy
// @title Search Service Api
// @version 1.0
// @description Search Service Api
func main() {
	pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithStyle("Search", pterm.FgLightGreen.ToStyle()),
		putils.LettersFromStringWithStyle(" Service", pterm.FgLightMagenta.ToStyle())).
		Render()

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
