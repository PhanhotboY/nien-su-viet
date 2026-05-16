package main

import (
	"os"

	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/shared/app"
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:              "billing-microservice",
	Short:            "billing-microservice based on vertical slice architecture",
	Long:             `This is a command runner or cli for api architecture in golang.`,
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {
		app.NewApp().Run()
	},
}

// https://github.com/swaggo/swag#how-to-use-it-with-gin

// @contact.name Phan Nguyen
// @contact.url https://github.com/phanhotboy
// @title Billing Service Api
// @version 1.0
// @description Billing Service Api
func main() {
	pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithStyle("Billing", pterm.FgLightGreen.ToStyle()),
		putils.LettersFromStringWithStyle(" Service", pterm.FgLightMagenta.ToStyle())).
		Render()

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
