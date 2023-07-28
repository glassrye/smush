package main

import (
	"flag"
	"log"
	"os"
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

// getCli parses command line flags and returns a config object and a slive of flag.FlagSet
func getCli() *config {
	var c config
	c.backup = false
	c.track = false

	fs := flag.NewFlagSet("backup", flag.ExitOnError)
	fs.StringVar(&c.backupBucket, "bucket", "", "Specify the bucket name")
	fs.StringVar(&c.backupPrefix, "prefix", "", "Specify the prefix, aka folder, as a location")

	ft := flag.NewFlagSet("track", flag.ExitOnError)
	ft.StringVar(&c.host, "host", "", "The hostname or IP addr of the database")
	ft.StringVar(&c.user, "user", "", "The DB user string")
	ft.StringVar(&c.pass, "pass", "", "The DB user password string")
	ft.StringVar(&c.database, "db", "", "The DB name (where the tables be, yarr....")

	fm := flag.NewFlagSet("files", flag.ExitOnError)
	fm.StringVar(&c.match, "match", "", "File name match. Like a regex.")
	fm.StringVar(&c.suffix, "suff", "", "The suffix of the file, e.g., .log or .txt")
	fm.StringVar(&c.watchDir, "dir", "", "The directory to look in for the files in")

	fe := flag.NewFlagSet("env", flag.ExitOnError)
	fe.StringVar(&c.envFile, "env", "", "An optional environment file")

	switch os.Args[1] {
	case "backup":
		err := fs.Parse(os.Args[2:])
		if err != nil {
			log.Fatalf("there was an error parsing arguments: %v", err)
		}
		c.backupProviders = flag.Args()
		c.backup = true
		return &c
	case "track":
		err := ft.Parse(os.Args[2:])
		if err != nil {
			log.Fatalf("there was an error parsing arguments: %v", err)
		}
		c.track = true
		return &c
	case "files":
		err := fm.Parse(os.Args[2:])
		if err != nil {
			log.Fatalf("there was an error parsing arguments: %v", err)
		}
		return &c
	default:
		return &c
	}
	// return &c
}
