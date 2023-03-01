package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/frederikhs/sql2csv"
	"log"
	"os"
	"time"
)

func main() {
	q := flag.String("q", "", "query to run")
	f := flag.String("f", "", "file containing query to run eg. query.sql")
	o := flag.String("o", "", "output filename eg. result.csv")
	t := flag.Int("t", 0, "query timeout in seconds")
	v := flag.Bool("v", false, "verbose mode")
	d := flag.String("d", "", "hostname for database as defined in .pgpass")
	c := flag.String("c", "", "connection string for the database: postgres://user:pass@host:port/dbname")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Lmicroseconds|log.Ltime)
	loggerFn := func(ln string) {
		verboseLog(logger, *v, ln)
	}

	// must always give an output name
	if *o == "" {
		flag.Usage()
		os.Exit(1)
	}

	// must always either get a file or a query, not both
	if (*f != "" && *q != "") || (*f == "" && *q == "") {
		flag.Usage()
		os.Exit(1)
	}

	if *t < 0 {
		flag.Usage()
		os.Exit(1)
	}

	var conn *sql2csv.Connection
	var err error

	if (*d != "" && *c != "") || (*d == "" && *c == "") {
		flag.Usage()
		os.Exit(1)
	}

	if *d != "" {
		loggerFn("connecting to database using pgpass")
		conn, err = sql2csv.NewConnectionFromPgPass(context.Background(), *d)
	} else {
		loggerFn("connecting to database using connection string")
		conn, err = sql2csv.NewConnection(context.Background(), *c)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	loggerFn("connected to database")

	var query *sql2csv.Query
	if *f != "" {
		query, err = readQueryFromFile(*f)
	} else {
		query, err = sql2csv.NewQuery(*q)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ctx := createContext(*t)
	err = conn.WriteQuery(ctx, query, *o, loggerFn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func readQueryFromFile(path string) (*sql2csv.Query, error) {
	f, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return sql2csv.NewQuery(string(f))
}

func createContext(timeout int) context.Context {
	if timeout == 0 {
		return context.Background()
	}

	c, _ := context.WithTimeout(context.Background(), time.Second*time.Duration(timeout))

	return c
}

func verboseLog(logger *log.Logger, verbose bool, ln string) {
	if verbose {
		logger.Println(ln)
	}
}
