package cmd

import (
	"fmt"
	"github.com/Artonus/hermes/internal/cmdutil"
	"github.com/Artonus/hermes/internal/config"
	"os"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete files from S3.",
	Long:  `This command deletes all uploaded files from S3 bucket.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("delete called")
		accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
		secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
		region := os.Getenv("AWS_REGION")
		bucket := os.Getenv("AWS_BUCKET")
		key := os.Getenv("NETAPP_KEY")
		cfg := config.NewConfig(accessKey, secretKey, config.WithBucket(bucket), config.WithRegion(region), config.WithKey(key))

		deleteClient, err := cmdutil.CreateDeleteClient(cfg)
		if err != nil {
			return err
		}

		res, deleteErr := deleteClient.Delete(cfg.Key)
		if !res && deleteErr != nil {
			return deleteErr
		}
		if res && deleteErr != nil {
			fmt.Println(deleteErr.Error())
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
