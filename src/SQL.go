package src

import (
	"GoMango/transfer"
	"database/sql"
	"strconv"
)

var (
	c = transfer.X()
)

func (Db *Data) Connect() {
	conf := c.Config("dbconfig.json")
	db, err := sql.Open(
		conf.DbType,
		conf.DbAuth.User+":"+conf.DbAuth.Pass+
			"@tcp(localhost:"+
			strconv.Itoa(conf.Port)+")/"+
			conf.Db,
	)
	c.Errs(err)
	defer db.Close()
}

func X() Data {
	x := Data{}
	return x
}
