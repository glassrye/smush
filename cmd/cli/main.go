package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/deepdyve/compress-logs/internal/models"
	"github.com/deepdyve/compress-logs/internal/process"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type config struct {
	User     string
	Pass     string
	Host     string
	Database string
	Backup   bool
}

// I think I should refactor and create a type that defines a file and the attributes it can have

func main() {
	// Set some variables for application user
	var c config
	var watchDir string
	var match string
	var suffix string
	var envFile string
	var err error
	year, month, day := time.Now().Date()
	defMatch := fmt.Sprintf("%v-%02d-%v", year, int(month), day-1)

	old := time.Now()
	oldTime := old.AddDate(0, 0, -90)
	oldMatch := fmt.Sprintf("%v-%02d-%v", oldTime.Year(), int(oldTime.Month()), oldTime.Day())
	fmt.Println("Old Match ", oldMatch)

	flag.BoolVar(&c.Backup, "b", false, "Do you want to backup the files to a bucket and track in a database")
	flag.StringVar(&c.User, "user", "", "The user name for the database connection. AKA: DB_USER env variable")
	flag.StringVar(&c.Pass, "pass", "", "The password for the database connection. AKA: DB_PASS env variable")
	flag.StringVar(&c.Host, "host", "", "The hostname for the database connection. AKA: DB_HOST env variable")
	flag.StringVar(&c.Database, "db", "", "The db name for the database connection. AKA: DB_NAME env variable")
	flag.StringVar(&watchDir, "dir", "", "The directory to watch for files.")
	flag.StringVar(&match, "m", defMatch, "The string to match for files.")
	flag.StringVar(&suffix, "s", "log", "The filename suffix to use.")
	flag.StringVar(&envFile, "e", "", "The environment variable file.")
	flag.Parse()

	// This is available for use
	// The idea is that I can create the container and run it
	// but also gen an environment file via GSM during the CI and/or deployment process
	// So, we have these way to reference the necessary variable:
	// 1.) just set the ENV variables manually (ick)
	// 2.) set the variables via cli flags (pretty good!)
	// 3.) set the variables by loading an environment file (also pretty OK, but just gen this during CI)
	err = godotenv.Load(envFile)
	if err != nil {
		fmt.Println("Skipping local environment loading")
	}

	if c.Backup && os.Getenv("DB_USER") == "" {
		fmt.Printf("You must specify a user name or set the DB_USER variable")
		return
	}
	if c.Backup && os.Getenv("DB_PASS") == "" {
		fmt.Printf("You must specify a user password or set the DB_PASS variable")
		return
	}
	if c.Backup && os.Getenv("DB_HOST") == "" {
		fmt.Printf("You must specify a database host or set the DB_HOST variable")
		return
	}
	if c.Backup && os.Getenv("DB_NAME") == "" {
		fmt.Printf("You must specify a database name or set the DB_NAME variable")
		return
	}
	if watchDir == "" {
		fmt.Printf("You must specify a directory to watch.\n")
		return
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", c.User, c.Pass, c.Host, c.Database)
	models.DB, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Error with DB: %v", err)
	}
	defer models.DB.Close()

	// Walk the directory
	dirList, err := os.ReadDir(watchDir)
	if err != nil {
		fmt.Printf("Error reading directory: %v", err)
	}
	var fl []string
	for _, e := range dirList {
		if !e.IsDir() {
			if strings.Contains(e.Name(), match) && strings.HasSuffix(e.Name(), suffix) {
				fp := watchDir + e.Name()
				// fmt.Println("Found a file: ", fp)
				fl = append(fl, fp)
			}
		}
	}

	// TODO: Create a channel and wait group
	// Let the channel hold the return from Compress()
	// Send a pointer to the waitgroup to Compress() and have Compress() defer the close
	// Ask about this in gophers.
	for _, f := range fl {
		// fmt.Printf("File Path: %s\n", f)
		arch := models.Archive{
			Archived:      false,
			Online:        true,
			OriginLoc:     f,
			ArchiveBucket: "dd-kaia-test", // This will become dynamic
			ArchivePrefix: "jbk-test",     // this will also become dynamic
		}
		arch.CompressLoc, arch.OriginHash, arch.CompressHash, err = process.Compress(arch.OriginLoc)

		if err != nil {
			fmt.Printf("Error processing file: %v\n", err)
			return
		}
		res, err := arch.AddRecord()
		if err != nil {
			fmt.Printf("Error adding record: %v\n", err)
			continue
		}
		id, err := res.LastInsertId()
		if err != nil {
			fmt.Printf("Error getting last insert id: %v\n", err)
		}
		s := strconv.Itoa(int(id))
		fmt.Printf("Insert ID: %s New File Hash: %s\n", s, arch.CompressHash)
	}
}
