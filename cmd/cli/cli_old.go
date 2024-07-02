package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

type config struct {
	user            string
	pass            string
	host            string
	db        string
	watchDir        string
	match           string
	suffix          string
	backupBucket    string
	backupPrefix    string
	backupProviders []string
	envFile         string
	track           bool
	backup          bool
	compress bool
}

// GetCli - returns a pointer to a config struct config
func getCli() *config {
	 c := &config{}
	 var rootCmd = &cobra.Command{
		Use: "smush",
		Run: func(cmd *cobra.Command, args []string) {
			if !c.compress {
				fmt.Println("Error: --compress flag is required!")
				os.Exit(1)
			}
			fmt.Println("Compression enabled.")
			fmt.Printf("Directory: %s\n", c.watchDir)
			fmt.Printf("Matching Name: %s\n", c.match)
			fmt.Printf("Suffix: %s\n", c.suffix)
			
			if c.track {
				fmt.Println("Tracking enabled.")
				fmt.Printf("Database: %s\n", c.db)
				fmt.Printf("Host: %s\n", c.host)
				fmt.Printf("User: %s\n", c.user)
				fmt.Printf("Pass: %s\n", c.pass)
			}
			if c.backup {
				fmt.Println("Backup enabled.")
				fmt.Printf("Bucket: %s\n", c.backupBucket)
				fmt.Printf("Folder: %s\n", c.backupPrefix)
			}
		}, 
	}
	rootCmd.Flags().BoolVarP(&c.compress, "compress", "c", false, "Enable compression (required)")
	rootCmd.Flags().StringVarP(&c.watchDir, "dir", "d", "", "Directory for compression")
	rootCmd.Flags().StringVar(&c.match, "match", "", "Matching name for files")
	rootCmd.Flags().StringVar(&c.suffix, "suff", "", "Suffix for files")

	rootCmd.Flags().BoolVarP(&c.track, "track", "T", false, "Enable tracking")
	rootCmd.Flags().StringVar(&c.db, "db", "", "Database name")
	rootCmd.Flags().StringVar(&c.host, "host", "", "Host address")
	rootCmd.Flags().StringVar(&c.user, "user", "", "Database user name")
	rootCmd.Flags().StringVar(&c.pass, "pass", "", "Database user pass")

	rootCmd.Flags().BoolVarP(&c.backup, "backup", "b", false, "Enable backup")
	rootCmd.Flags().StringVar(&c.backupBucket, "bucket", "", "Bucket name")
	rootCmd.Flags().StringVar(&c.backupPrefix, "folder", "", "Folder in bucket")
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("there was an error parsing arguments: %v", err)
	}

	return c	
}
