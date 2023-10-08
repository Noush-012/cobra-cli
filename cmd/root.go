package cmd

import (
	"fmt"
	"os"

	"github.com/Noush-012/cobra-cli/utils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Test",
	Short: "This app will greet you !",
	Long: `A Fast and Flexible Static Site Generator built with
	love by spf13 and friends in Go.
	Complete documentation is available at http://hugo.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {
		//Greeting
		fmt.Printf("%sHey buddy ! you are powerful%s\n", utils.BrightCyan, utils.Reset)
		utils.TextMorphing()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
