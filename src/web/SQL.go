package web

import (
	"GoMango/transfer"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	c = transfer.X()
)

func (D *Data) Connect() (*sql.DB, error) {
	//d, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/gomango")
	//if err != nil {
	//	return &sql.DB{}, err
	//}
	//defer d.Close()
	cfg := c.Config("dbconfig.json")
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", cfg.DbAuth.User, cfg.DbAuth.Pass, "localhost:3306", cfg.Database.Db))
	if err != nil {
		log.Fatal(err)
	}

	//db, err := sql.Open(
	//	conf.Database.DbType,
	//	conf.DbAuth.User+":"+conf.DbAuth.Pass+
	//		"@tcp(localhost:"+
	//		strconv.Itoa(conf.Database.Port)+")/"+
	//		conf.Database.Db,
	//)
	return db, nil
}

func (D *Data) Write(db *sql.DB, query string) error {
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	return nil
}

func X() Data {
	x := Data{}
	return x
}
