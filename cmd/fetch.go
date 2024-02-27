package cmd

import (
	"fmt"
	"github.com/Artonus/hermes/internal/cmdutil"
	"github.com/Artonus/hermes/internal/config"
	"github.com/spf13/cobra"
	"os"
)

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetches the data from S3.",
	Long:  `Fetches the data from S3 to the specified directory.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("fetch called")
		accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
		secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
		region := os.Getenv("AWS_REGION")
		bucket := os.Getenv("AWS_BUCKET")
		fetchDir := os.Getenv("FETCH_DIR")
		key := os.Getenv("NETAPP_KEY")
		cfg := config.NewConfig(accessKey, secretKey, config.WithFetchDir(fetchDir), config.WithBucket(bucket), config.WithRegion(region), config.WithKey(key))

		fetchClient, err := cmdutil.CreateFetchClient(cfg)
		if err != nil {
			panic(err)
		}

		err = fetchClient.Fetch(cfg.Key, fetchDir)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fetchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fetchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
