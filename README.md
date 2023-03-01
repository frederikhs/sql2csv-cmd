# sql2csv-cmd

[![GoDoc](https://godoc.org/github.com/frederikhs/sql2csv-cmd?status.svg)](https://godoc.org/github.com/frederikhs/sql2csv-cmd)
[![Quality](https://goreportcard.com/badge/github.com/frederikhs/sql2csv-cmd)](https://goreportcard.com/report/github.com/frederikhs/sql2csv-cmd)
[![Test](https://github.com/frederikhs/sql2csv-cmd/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/frederikhs/sql2csv-cmd/actions/workflows/test.yml)
[![Release](https://img.shields.io/github/v/release/frederikhs/sql2csv-cmd.svg)](https://github.com/frederikhs/sql2csv-cmd/releases/latest)
[![License](https://img.shields.io/github/license/frederikhs/sql2csv-cmd)](LICENSE)

*cli application for extracting data as csv from a PostgreSQL database using sql*

## Usage

```text
Usage of sql2csv-cmd:
  -c string
    	connection string for the database: postgres://user:pass@host:port/dbname
  -d string
    	hostname for database as defined in .pgpass
  -f string
    	file containing query to run eg. query.sql
  -o string
    	output filename eg. result.csv
  -q string
    	query to run
  -t int
    	query timeout in seconds
  -v	verbose mode
```

## Example

```shell
sql2csv-cmd -d db.example.com -q "SELECT * FROM public.users" -o users.csv
```

```shell
sql2csv-cmd -d db.example.com -f query.sql -o query.csv
```

## Installation

### Linux amd64

```bash
# install
curl -L https://github.com/frederikhs/sql2csv-cmd/releases/latest/download/sql2csv-cmd_Linux_x86_64.tar.gz -o sql2csv-cmd.tar.gz
tar -xvf sql2csv-cmd.tar.gz
sudo mv sql2csv-cmd /usr/local/bin/sql2csv-cmd

# clean up
rm sql2csv-cmd.tar.gz
```

## .pgpass

sql2csv-cmd uses the `.pgpass` file that resides in the `$HOME` directory of the user running the program.

---

Example of a connection configuration in the `.pgpass` file

```
<HOST>:<PORT>:<DATABASE>:<USERNAME>:<PASSWORD>
```
