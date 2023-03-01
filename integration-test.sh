#!/bin/bash
set -e

docker run --name sql2csv-cmd-postgres --rm -e POSTGRES_PASSWORD=test -p "53406:5432" -d postgres

# stupid way of making sure the postgres docker image has launched
sleep 15

go build -o sql2csv-cmd-test

./sql2csv-cmd-test -c postgres://postgres:test@localhost:53406/postgres -q "SELECT 1 as number" -o test1.csv

docker exec sql2csv-cmd-postgres gosu postgres psql --dbname=postgres --command="
CREATE TABLE users (
	id INT,
	name VARCHAR
); INSERT INTO users (id, name) VALUES (1, 'test'), (2, 'always');
"

./sql2csv-cmd-test -c postgres://postgres:test@localhost:53406/postgres -q "SELECT * from users" -o test2.csv

docker rm -f sql2csv-cmd-postgres

if [[ "number
1" != "$(cat test1.csv)" ]]; then
  echo "test1 does not match"
  exit 1
else
  echo "test1 pass"
fi

if [[ "id,name
1,test
2,always" != "$(cat test2.csv)" ]]; then
  echo "test2 does not match"
  exit 1
else
  echo "test2 pass"
fi
