package web

import (
	"GoMango/transfer"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	c = transfer.X()
)

func (D *Data) Connect() (*sql.DB, error) {
	d, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/gomango")
	if err != nil {
		return &sql.DB{}, err
	}
	defer d.Close()

	//conf := c.Config("dbconfig.json")
	//db, err := sql.Open(
	//	conf.Database.DbType,
	//	conf.DbAuth.User+":"+conf.DbAuth.Pass+
	//		"@tcp(localhost:"+
	//		strconv.Itoa(conf.Database.Port)+")/"+
	//		conf.Database.Db,
	//)
	return d, nil
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
