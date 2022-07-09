package main

import (
	"sql/dataBase"

	_ "github.com/lib/pq"
)

func main() {

	//var outb, errb bytes.Buffer
	// run pg server
	/*	go func() {
			cmd := exec.Command("pg_ctl", "-D", database_data, "start")

			cmd.Stdout = &outb
			cmd.Stderr = &errb
			err := cmd.Run()
			if err != nil {
				log.Fatal(err)
			}
		}()
		fmt.Println("out:", outb.String(), "err:", errb.String()) */

	dataBase.Setup()
	dataBase.RegisterTrigger()
	dataBase.Listener()

	/*

		age := 21
		rows, err := db.Query("select name from users where age = $1", age)
		if err != nil {
			log.Fatal(err)
		}
		_ = rows
		//for index, row := range rows.Columns() {
		//		fmt.Println(index, row)
		//	}

		var name string
		stmt, err := db.Prepare("select name from users where id = $1")
		if err != nil {
			log.Fatal(err)
		}

		idc := 4
		err = stmt.QueryRow(idc).Scan(&name)
		if err != nil {
			log.Fatal(err)
		}

		var userid int
		err = db.QueryRow("insert into album (title, artist) values ($1, $2)", "title", "artist").Scan(&userid)
		if err != nil {
			log.Fatal(err)
		} */
}
