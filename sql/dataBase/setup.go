package dataBase

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

type sqltable struct {
	schemaname, tablename, tableowner, tablespace  sql.NullString
	hasindexes, hasrules, hastriggers, rowsecurity sql.NullBool
}

type Table struct {
	schemaname, tablename, tableowner, tablespace  string
	hasindexes, hasrules, hastriggers, rowsecurity bool
}

func Setup() {
	go func() {
		out, err := exec.Command("pg_ctl", "-D", database_data, "start").Output()
		if err != nil {
			log.Fatal("error:", err)
		}
		fmt.Printf("The date is %s\n", out)
	}()
	var err error

	// open database
	Store, err = sql.Open("postgres", conninfo)
	if err != nil {
		log.Fatal(err)
	}
	// close db
	//defer Store.Close()

	// check db
	err = Store.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected!")
	// Read SQL statement from sql file
	//	f, err := os.Open("dataBase/createTable.sql") // path reference to the root directory, i.e. to the main.go
	f, err := os.Open("dataBase/listTables.sql")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	b := new(strings.Builder)
	io.Copy(b, f)
	sql := b.String()
	//println(sql)

	stmt, err := Store.Prepare(sql)
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = stmt.Exec()
	if err != nil {
		fmt.Println(err.Error())
	}

	rows, err := stmt.Query() // .Exec()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer rows.Close()
	var sqlresult = sqltable{}
	tables := []Table{}
	for rows.Next() {
		err = rows.Scan(&sqlresult.schemaname, &sqlresult.tablename, &sqlresult.tableowner, &sqlresult.tablespace,
			&sqlresult.hasindexes, &sqlresult.hasrules, &sqlresult.hastriggers, &sqlresult.rowsecurity)
		if err != nil {
			fmt.Println("failed to scan", err)
		}
		// fmt.Println(sqlresult)
		result := Table{
			schemaname:  sqlresult.schemaname.String,
			tablename:   sqlresult.tablename.String,
			tableowner:  sqlresult.tableowner.String,
			tablespace:  sqlresult.tablespace.String,
			hasindexes:  sqlresult.hasindexes.Bool,
			hasrules:    sqlresult.hasrules.Bool,
			hastriggers: sqlresult.hastriggers.Bool,
			rowsecurity: sqlresult.rowsecurity.Bool,
		}
		tables = append(tables, result)
	}
	fmt.Println(tables)
}

/*
import (
    "database/sql"
    "fmt"
)

// LookupName returns the username from database ID.
func LookupName(id int) (string, error) {
    db := Connect()
    defer db.Close()
    var name sql.NullString
    err := db.QueryRow("SELECT username FROM accounts WHERE id=?", id).Scan(&name)
    if err != nil {
        return "", fmt.Errorf("lookup name by id %q: %w", id, err)
    }
    return name.String, nil
}
*/
