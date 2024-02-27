package cmd

import (
	"fmt"
	"github.com/Artonus/hermes/internal/cmdutil"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

// postCmd represents the post command
var postCmd = &cobra.Command{
	Use:   "post",
	Short: "Sends the data to S3.",
	Long:  `Sends the data from specified directory directly to S3.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("post called")
		fetchDir := os.Getenv("POST_DIR")
		postClient, err := cmdutil.CreatePostClient()
		if err != nil {
			panic(err)
		}

		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

		fmt.Println("Waiting for SIGTERM signal...")

		// Block until a signal is received
		sigReceived := <-sigCh
		// Notify the channel when a SIGTERM signal is received
		fmt.Println("Received signal: %v", sigReceived)
		// Upload the data to S3
		err = postClient.Post("", fetchDir)
		if err != nil {
			return err
		}

		fmt.Println("fetch completed")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(postCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// postCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// postCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
