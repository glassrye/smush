package main

import (
	"fmt"
	"log"

	"github.com/deepdyve/compress-logs/internal/compress"
	"github.com/spf13/cobra"
)

var compressCmd = &cobra.Command{
	Use:   "compress",
	Short: "Compress files",
	Run: func(cmd *cobra.Command, args []string) {
		directory, _ := cmd.Flags().GetString("directory")
		suffix, _ := cmd.Flags().GetString("suffix")
		level, _ := cmd.Flags().GetInt("level")
		match, _ := cmd.Flags().GetString("match")

		// Implement your compression logic here
		fmt.Printf("Compressing: Directory=%s, Suffix=%s, Level=%d, Match=%s\n", directory, suffix, level, match)

	},
}

var trackCmd = &cobra.Command{
	Use:   "track",
	Short: "Track files",
	Run: func(cmd *cobra.Command, args []string) {
		provider, _ := cmd.Flags().GetString("provider")
		bucket, _ := cmd.Flags().GetString("bucket")
		prefix, _ := cmd.Flags().GetString("prefix")
		db, _ := cmd.Flags().GetString("db")
		user, _ := cmd.Flags().GetString("user")
		pass, _ := cmd.Flags().GetString("pass")

		// Implement your tracking logic here
		fmt.Printf("Tracking: Provider=%s, Bucket=%s, Prefix=%s, DB=%s, User=%s, Pass=%s\n", provider, bucket, prefix, db, user, pass)
	},
}

func cli() {
	compressCmd.Flags().StringP("directory", "d", "", "Directory to compress")
	compressCmd.Flags().BoolP("recurse", "r", false, "Recurse the original directory")
	compressCmd.Flags().StringP("suffix", "s", "", "File suffix to filter")
	compressCmd.Flags().IntP("level", "l", 0, "Compression level")
	compressCmd.Flags().StringP("match", "m", "", "File matching pattern")
	compressCmd.Flags().String("provider", "", "Cloud provider")
	compressCmd.Flags().String("bucket", "", "Bucket name")
	compressCmd.Flags().String("prefix", "", "File prefix")
	compressCmd.Flags().String("db", "", "Database name")
	compressCmd.Flags().String("user", "", "Database user")
	compressCmd.Flags().String("pass", "", "Database password")
	compressCmd.MarkFlagRequired("directory") // Make directory flag required for compress command

	var rootCmd = &cobra.Command{
		Use: "smush",
	}

	var compressCmd = &cobra.Command{
		Use:   "compress",
		Short: "Compress and Archive Files with Tracking",
		Run: func(cmd *cobra.Command, args []string) {
			dir, _ := cmd.Flags().GetString("directory")
			suff, _ := cmd.Flags().GetString("suffix")
			level, _ := cmd.Flags().GetInt("level")
			match, _ := cmd.Flags().GetString("match")
			prov, _ := cmd.Flags().GetStringSlice("provider")
			bucket, _ := cmd.Flags().GetString("bucket")
			prefix, _ := cmd.Flags().GetString("prefix")
			db, _ := cmd.Flags().GetString("db")
			user, _ := cmd.Flags().GetString("user")
			pass, _ := cmd.Flags().GetString("pass")
			recurse, _ := cmd.Flags().GetBool("recurse")

			a := &compress.Archive{}
			a.ArchiveBucket = bucket
			a.ArchivePrefix = prefix
			a.Recurse = recurse
			a.OriginLoc = dir
			a.DBUser = user
			a.DBPass = pass
			a.DBHost = db
			a.Match = match
			a.Suffix = suff
			a.Providers = prov
			a.Level = level

			// err := runCompress(a, a.OriginLoc, suff, match)
			err := a.Compress()
			if err != nil {
				log.Fatalf("there was a fatal error: %v", err)
				return
			}
		},
	}
	rootCmd.AddCommand(compressCmd)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func runCompress(a *compress.Archive, d, s, m string) error {
	fmt.Printf("")
	return nil
}

/*
func runCompress(d,s,m,p,b,pre,db,us,pa string, l int, rec bool) error {
	dir, err := os.ReadDir(d)
	if err != nil {
		return err
	}
	for _, v := range dir {
		a := &compress.Archive{}
		if !v.IsDir() {
			if strings.Contains(v.Name(), s) {
				a.OriginLoc = fmt.Sprintf("%s/%s", d, v.Name())
				a.CompressLoc = fmt.Sprintf("%s/%s.gz", d, v.Name())
				err := a.Compress()
				if err != nil {
					return err
				}
			}
		}
	}
}
*/
