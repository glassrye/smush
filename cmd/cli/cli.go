package main

import (
	"github.com/spf13/cobra"
	"log"
)

type config struct {
	user            string
	pass            string
	host            string
	database        string
	watchDir        string
	match           string
	suffix          string
	backupBucket    string
	backupPrefix    string
	backupProviders []string
	envFile         string
	track           bool
	backup          bool
}

// getCli parses command line flags and returns a config object that can be used for archiving files.
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
	cmdTrack.Flags().StringVar(&c.database, "db", "", "The DB name (where the tables be, yarr....")

	var cmdFiles = &cobra.Command{
		Use:   "files",
		Short: "Files command",
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
	rootCmd.AddCommand(cmdBackup, cmdTrack, cmdFiles)
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("there was an error parsing arguments: %v", err)
	}

	return &c
}
