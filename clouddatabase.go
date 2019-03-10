package main

import (
	_ "github.com/lib/pq"
)

type Database struct{}

func NewDatabase() *Database {
	return &Database{}
}

/*
var db *sql.DB

type Database struct {
	*sql.DB
}

var (
	connectionName = "mbtio-234017:us-east4:mbtio"
	dbUser         = "postgres"
	dbPassword     = "Testing123"
	dsn            = fmt.Sprintf("%s:%s@unix(/cloudsql/%s)/mbtio", dbUser, dbPassword, connectionName)
)

//manages google cloud postgres db connection
func NewDatabase() (*Database, error) {

	//	db, err := sql.Open("postgres", mustGetenv("POSTGRES_CONNECTION"))
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		return nil, err
	}

	d := &Database{db}
	d.CreateUserTable()

	fmt.Println("X")

	return d, nil
}

func (db *Database) CreateUserTable() error {
	stmt, err := db.Prepare(`CREATE TABLE IF NOT EXISTSCREATE TABLE Users (uid int PRIMARY KEY NOT NULL ,pin  int NOT NULL,
		 name varchar(30) NOT NULL,	phone varchar(30) NOT NULL,mbti varchar(4), bio varchar(512), height int NOT NULL,
		 age int NOT NULL, sex varchar(10) NOT NULL, interest varchar(10) NOT NULL, lat float8 NOT NULL,long float8 NOT NULL);`)

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec("mbtio")

	if err != nil {
		return err
	}

	return nil
}
*/
