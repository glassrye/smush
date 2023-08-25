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

// getCli parses command line flags and returns a config object that can be used for archiving files.
/*
func getCli() *config {
	var c config

	var rootCmd = &cobra.Command{Use: "smush"}

	var cmdBackup = &cobra.Command{
		Use:   "backup",
		Short: "Backup command",
		Run: func(cmd *cobra.Command, args []string) {
			c.backupProviders = args
			c.backup = true
		},
	}
	cmdBackup.Flags().StringVar(&c.backupBucket, "bucket", "", "Specify the bucket name")
	cmdBackup.Flags().StringVar(&c.backupPrefix, "prefix", "", "Specify the prefix, aka folder, as a location")

	var cmdTrack = &cobra.Command{
		Use:   "track",
		Short: "Track command",
		Run: func(cmd *cobra.Command, args []string) {
			c.track = true
		},
	}
	cmdTrack.Flags().StringVar(&c.host, "host", "", "The hostname or IP addr of the database")
	cmdTrack.Flags().StringVar(&c.user, "user", "", "The DB user string")
	cmdTrack.Flags().StringVar(&c.pass, "pass", "", "The DB user password string")
	cmdTrack.Flags().StringVar(&c.dsn, "db", "", "The DB name (where the tables be, yarr....")

	var cmdFiles = &cobra.Command{
		Use:   "compress",
		Short: "Compress files command",
		Run:   func(cmd *cobra.Command, args []string) {},
	}
	cmdFiles.Flags().StringVar(&c.match, "match", "", "File name match. Like a regex.")
	cmdFiles.Flags().StringVar(&c.suffix, "suff", "", "The suffix of the file, e.g., .log or .txt")
	cmdFiles.Flags().StringVar(&c.watchDir, "dir", "", "The directory to look in for the files in")

	// I haven't decided if I want to actually do this yet. Seems like bullshit extra stuff to me
	// but I could be wrong so I'm leaving this hear for the time being
	/*var cmdEnv = &cobra.Command{
		Use:   "env",
		Short: "Env command",
		Run:   func(cmd *cobra.Command, args []string) {},
	}
	cmdEnv.Flags().StringVar(&c.envFile, "env", "", "An optional environment file")
	*/

	// rootCmd.AddCommand(cmdBackup, cmdTrack, cmdFiles, cmdEnv)
	/*
	rootCmd.AddCommand(cmdBackup, cmdTrack, cmdFiles)
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("there was an error parsing arguments: %v", err)
	}
	return &c
}
	*/


/*var rootCmd = &cobra.Command{
	Use: "yourcli",
	Run: func(cmd *cobra.Command, args []string) {
		if !c.compress {
			fmt.Println("Error: --compress flag is required!")
			os.Exit(1)
		}
		fmt.Println("Compression enabled.")
		fmt.Printf("Directory: %s\n", dir)
		fmt.Printf("Matching Name: %s\n", matchName)
		fmt.Printf("Suffix: %s\n", suffix)
		
		if c.track {
			fmt.Println("Tracking enabled.")
			fmt.Printf("Database: %s\n", db)
			fmt.Printf("Host: %s\n", host)
			fmt.Printf("User: %s\n", user)
		}
		if c.backup {
			fmt.Println("Backup enabled.")
			fmt.Printf("Bucket: %s\n", bucket)
			fmt.Printf("Folder: %s\n", folder)
		}
	},
}*/
// var compressFlag, track, backup bool
// var dir, matchName, suffix, db, host, user, bucket, folder string

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
