package dataBase

// Package listen is a self-contained Go program which uses the LISTEN / NOTIFY
// mechanism to avoid polling the database while waiting for more work to arrive.
//
// You can see the program in action by defining a function similar to
// the following:
//
// CREATE OR REPLACE FUNCTION public.get_work()
//   RETURNS bigint
//   LANGUAGE sql
//   AS $$
//     SELECT CASE WHEN random() >= 0.2 THEN int8 '1' END
//   $$
// ;

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/lib/pq"
)

func RegisterTrigger() {
	reportProblem := func(ev pq.ListenerEventType, err error) {
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	minReconn := 10 * time.Second
	maxReconn := time.Minute
	listener := pq.NewListener(conninfo, minReconn, maxReconn, reportProblem)
	err := listener.Listen("getwork")
	if err != nil {
		panic(err)
	}

	// Read SQL statement from sql file
	f, err := os.Open("dataBase/trigger.sql") // path reference to the root directory, i.e. to the main.go
	if err != nil {
		panic(err)
	}
	defer f.Close()
	b := new(strings.Builder)
	io.Copy(b, f)
	//println(b.String())
	// Or read as bytes
	/*	b, err := os.ReadFile("db/trigger.sql")
		if err != nil {
			panic(err)
		}
		os.Stdout.Write(b) */
	stmt, err := Store.Prepare(b.String())
	/*	stmt, err := db.Prepare(`
		CREATE or REPLACE public.messages_notify_trigger() RETURNS trigger AS $$
		DECLARE
			BEGIN
				PERFORM pg_notify(CAST ('new_message' AS text), row_to_json(NEW)::text),
				RETURN new;
			END;
		$$ LANGUAGE plpgsql;

		CREATE TRIGGER messages_new_trigger AFTER INSERT ON public.messages
		FOR EACH ROW EXECUTE PROCEDURE public.messages_notify_trigger();
		`) */
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = stmt.Exec()
	if err != nil {
		fmt.Println(err.Error())
	}
}
