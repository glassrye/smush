# SMUSH

Smush is just a simple utility that uses the Gzip library to do the following:

o Look into a directory and match file names based on CLI flag
o It can/will also watch for files with a specific suffix (.log, .txt, etc.)
o There is a default file match that is yesterday's date that follows this format:
o 2023-05-28 (YYYY-MM-DD)
o The match string is configurable but you are only allowed one match string

- You can run this as a daemon or as a one off via cron or what have you.

- The main reason to use this is for the database. Often you'll have files that are bespoke, not ingested properly into a log system like ELK, or Loki, or Graylog. It happens. It just does.

- In any case, these files will be tracke from origin with hash, through compression with a hash. Smush will also upload files that are a certain
  age and above to a GCP bucket (S3 support will be coming later. I just don't like AWS so I started with GCP)

---

```
O_O[james@tethys:~/devel/github.com/glassrye/smush]$ go run cmd/cli/main.go  --help
Old Match  2022-10-16
Usage of /var/folders/4x/6f286ndx6b71brrg5s8j_wx40000gn/T/go-build1093124682/b001/exe/main:
  -db string
        The db name for the database connection. AKA: DB_NAME env variable
  -dir string
        The directory to watch for files.
  -e string
        The environment variable file.
  -host string
        The hostname for the database connection. AKA: DB_HOST env variable
  -m string
        The string to match for files. (default "2023-01-13")
  -pass string
        The password for the database connection. AKA: DB_PASS env variable
  -s string
        The filename suffix to use. (default "log")
  -user string
        The user name for the database connection. AKA: DB_USER env variable
^_^[james@tethys:~/devel/github.com/glassrye/smush]$
```
