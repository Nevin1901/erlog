## erlog

erlog is a simple lightweight logging platform. you can log events from your code, and then they will show up in the backend

## running

`go run -tags json1 main.go`

### looking for people to help build the commercial version

will be hosted by us and will be easier for the customer. dm Nevin#7114 on discord if you want to be a part of it.

doesn't matter what you do so long as you're dedicated to building the worlds next logging software

## rationale

if we want fts, sqlite is really fast https://supabase.com/blog/postgres-full-text-search-vs-the-rest.
we can then split up the data into 500 chunks to save into the db. no case of failure if the logging server goes down.

sqlite can also use json. we can look through and see if there is a key

scaling -> clickhouse maybe? only I'm not using a docker image

docker images are an excuse to throw whatever and hope it works. erlog should be able to run fast and simple without having 8gb of ram and 2 cores.

If you just want to run this on your vps then you should be able to do that without pulling 10,000 depdenencies

## todo

set WAL, set pragma syncrhonous=NORMAL or off, batch inserts, use custom ORM

- benchmark sqlite with gorm and also with sqlite3 go, compilation options

- maybe try to set db.Set("gorm:table_options", "mystring").Migrator().CreateTable(&User{}) to append for creating table
