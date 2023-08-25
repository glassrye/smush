package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/deepdyve/compress-logs/internal/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	// Set some variables for application user
	c := getCli()
	var err error
	// year, month, day := time.Now().Date()
	// defMatch := fmt.Sprintf("%v-%02d-%v", year, int(month), day-1)

	old := time.Now()
	oldTime := old.AddDate(0, 0, -90)
	oldMatch := fmt.Sprintf("%v-%02d-%v", oldTime.Year(), int(oldTime.Month()), oldTime.Day())
	fmt.Println("Old Match ", oldMatch)

	if c.envFile != "" {
		err = godotenv.Load(c.envFile)
		if err != nil {
			log.Printf("error loading envFile: %v\n", err)
		}
	}

	if c.track {
		if c.user == "" && os.Getenv("DB_USER") == "" {
			fmt.Printf("You must specify a user name or set the DB_USER variable")
			return
		}
		if c.pass == "" && os.Getenv("DB_PASS") == "" {
			fmt.Printf("You must specify a user password or set the DB_PASS variable")
			return
		}
		if c.host == "" && os.Getenv("DB_HOST") == "" {
			fmt.Printf("You must specify a database host or set the DB_HOST variable")
			return
		}
		if c.db == "" && os.Getenv("DB_NAME") == "" {
			fmt.Printf("You must specify a database name or set the DB_NAME variable")
			return
		}
		/*dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", c.user, c.pass, c.host, c.db)
		models.DB.DB, err = sql.Open("mysql", dsn)
		if err != nil {
			fmt.Printf("Error with DB: %v", err)
		}
		defer models.DB.DB.Close()()
		*/
	}

	if c.watchDir == "" {
		fmt.Printf("You must specify a directory to watch.\n")
		return
	}

	// Walk the directory
	dirList, err := os.ReadDir(c.watchDir)
	if err != nil {
		fmt.Printf("Error reading directory: %v", err)
	}
	var fl []string
	for _, e := range dirList {
		if !e.IsDir() {
			if strings.Contains(e.Name(), c.match) && strings.HasSuffix(e.Name(), c.suffix) {
				fp := c.watchDir + e.Name()
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
		compFile := f + ".gz"
		arch := &models.Archive{
			Archived:      false,
			Online:        true,
			OriginLoc:     f,
			CompressLoc:   compFile,
			ArchiveBucket: "smush-test", // This will become dynamic
			ArchivePrefix: "jbk-test",   // this will also become dynamic
		}

		// arch.CompressLoc, arch.OriginHash, arch.CompressHash, err = process.Compress(arch.OriginLoc)
		err = arch.Compress()
		if err != nil {
			fmt.Println(err)
			return
		}

		if err != nil {
			fmt.Printf("Error processing file: %v\n", err)
			return
		}
		if c.track {
			d, err := models.NewDatabase(c.user, c.pass, c.db, c.host)
			if err != nil {
				fmt.Printf("Error getting DB connection: %v", err)
				return 
			}
			t := models.Tracker{
				DB: d,
			}
			t.Archive = *arch
			res, err := t.AddRecord()
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
}
