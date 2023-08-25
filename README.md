# SMUSH

Please consider this currently like... pre-pre-alpha. I doodle when I feel like it.

Smush is just a simple utility that uses the Gzip library to do the following:

- Look into a directory and match file names based on CLI flag and then compress them in that same directory.
  - You can also specifiy a suffix to match on
  - E.g., File name match: `access` File suffix match: `log`
- There is a default file match that is yesterday's date that follows this format:
  - 2023-05-28 (YYYY-MM-DD)
- The match string is configurable but you are only allowed one match string

- You can run this as a daemon or as a one off via cron or what have you.

- The main reason to use this is for the database. Often you'll have files that are bespoke, not ingested properly into a log system like ELK, or Loki, or Graylog. It happens. It just does.

- In any case, these files will be tracke from origin with hash, through compression with a hash. Smush will also upload files that are a certain
  age and above to a GCP bucket (S3 support will be coming later. I just don't like AWS so I started with GCP)

---

## Build

```
make build
```

This produces an amd64 and darwin build. Feel free to create a PR for others. I just don't wanna..

```
^_^[james@tethys ~/devel/github.com/glassrye/smush] (bucket-db) $ tree build/
build/
├── smush-amd64
└── smush-darwin

1 directory, 2 files
^_^[james@tethys ~/devel/github.com/glassrye/smush] (bucket-db) $

```

It is up to you to use the `db/init.sql` file to create a database for your self. You DO NOT have to actually track files in a database, but you'll need if you want to! :)

After you have created the database and run `make build` you can check out the options:

```
^_^[james@tethys ~/devel/github.com/glassrye/smush] (bucket-db) $ ./build/smush-darwin --help
Usage:
  smush [flags]

Flags:
  -b, --backup          Enable backup
      --bucket string   Bucket name
  -c, --compress        Enable compression (required)
      --db string       Database name
  -d, --dir string      Directory for compression
      --folder string   Folder in bucket
  -h, --help            help for smush
      --host string     Host address
      --match string    Matching name for files
      --pass string     Database user pass
      --suff string     Suffix for files
  -T, --track           Enable tracking
      --user string     Database user name
Old Match  2023-05-26
You must specify a directory to watch.
^_^[james@tethys ~/devel/github.com/glassrye/smush] (bucket-db) $
```

Here is an example of how I use `smush` to compress some files and create an entry in a database for later interrogation.

```
./build/smush-darwin -c -d ~/tmp/smush/ --match test_comp --suff log -T --db smushTracker --host localhost --user postgres --pass testme
```

### NOTE: GCP / S3 buckets are under construction
