package dataBase

import (
	"database/sql"
	"fmt"
	"os"
)

// Start and stop Pg data server
// pg_ctl -D "C:\Program Files\PostgreSQL\9.6\data" start
// pg_ctl -D "C:\Program Files\PostgreSQL\9.6\data" stop
// pg_ctl -D "C:\Program Files\PostgreSQL\9.6\data" restart
// Or Winkey+R => services.msc
//   pg_ctl -D "D:\Development\pgsql\data" start
var (
	Store             *sql.DB
	database_data     = os.Getenv("PGDATA") // D:\Development\pgsql\data
	database_host     = "localhost"
	database_user     = os.Getenv("PGUSER")
	database_password = os.Getenv("PGPSWD")
	database_ip       = os.Getenv("PGIP") // "5432"
	database_name     = os.Getenv("PGDB") // "postgres"
	//	connStr := "postgres://username:password@localhost:5432/database_name?sslmode=disable"
	conninfo string = fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		database_user, database_password, database_host, database_ip, database_name)
)
