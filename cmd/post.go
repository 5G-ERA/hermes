package cmd

import (
	"fmt"
	"github.com/Artonus/hermes/internal/cmdutil"
	"github.com/Artonus/hermes/internal/config"
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
		accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
		secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
		region := os.Getenv("AWS_REGION")
		bucket := os.Getenv("AWS_BUCKET")
		postDir := os.Getenv("POST_DIR")
		key := os.Getenv("NETAPP_KEY")
		cfg := config.NewConfig(accessKey, secretKey, config.WithPostDir(postDir), config.WithBucket(bucket), config.WithRegion(region), config.WithKey(key))

		postClient, err := cmdutil.CreatePostClient(cfg)
		if err != nil {
			panic(err)
		}

		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

		fmt.Println("Waiting for SIGTERM signal...")

		// Block until a signal is received
		sigReceived := <-sigCh
		// Notify the channel when a SIGTERM signal is received
		fmt.Printf("Received signal: %v\n", sigReceived)
		// Upload the data to S3
		err = postClient.Post(cfg.Key, postDir)
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
