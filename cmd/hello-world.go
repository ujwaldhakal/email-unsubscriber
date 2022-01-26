package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"time"
)

var rootCmd = &cobra.Command{
	Use: "hello world",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello world")

		const layout = "2006-01-02"
		snapshot := "2017-07-25"
		t, err := time.Parse(layout, snapshot)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(t.Unix())
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
