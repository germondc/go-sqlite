package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./foo.db")
	checkErr(err, "open database")

	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS userinfo (`uid` INTEGER PRIMARY KEY AUTOINCREMENT,`username` VARCHAR(64) NULL,`departname` VARCHAR(64) NULL,`created` DATE NULL)")
	checkErr(err, "create table statement")

	_, err = statement.Exec()
	checkErr(err, "run create table statement")

	// insert
	statement, err = db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
	checkErr(err, "create insert statement")

	res, err := statement.Exec("astaxie", "研发部门", "2012-12-09")
	checkErr(err, "execute prepare statement")

	id, err := res.LastInsertId()
	checkErr(err, "last insert id")

	fmt.Printf("last inserted id: %d\n", id)
	// update
	statement, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err, "create update statement")

	res, err = statement.Exec("astaxieupdate", id)
	checkErr(err, "execute update statement")

	affect, err := res.RowsAffected()
	checkErr(err, "get affected rows")

	fmt.Printf("rows affected: %d\n", affect)

	// query
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err, "run query")
	var uid int
	var username string
	var department string
	var created time.Time

	for rows.Next() {
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err, "get values")
		fmt.Printf("uid: %d\n", uid)
		fmt.Printf("username: %s\n", username)
		fmt.Printf("department: %s\n", department)
		fmt.Printf("created: %v\n", created)
	}

	rows.Close() //good habit to close

	// delete
	statement, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err, "create delete statement")

	res, err = statement.Exec(id)
	checkErr(err, "run delete statment")

	affect, err = res.RowsAffected()
	checkErr(err, "get affected rows")

	fmt.Printf("rows affected: %d\n", affect)

	db.Close()

}

func checkErr(err error, message string) {
	if err != nil {
		fmt.Printf("error: %s\n", message)
		panic(err)
	}
}
