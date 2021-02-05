package good

import "fmt"

type GoodMysqlConnection struct {
	sURL string
	sUID string
	sPWD string
}

func NewGoodMysqlConnection(uri string, uid string, pwd string) *GoodMysqlConnection {
	return &GoodMysqlConnection{
		sURL: uri,
		sUID: uid,
		sPWD: pwd,
	}
}

func (p *GoodMysqlConnection) Execute(sql string, args ...interface{}) (error, int)  {
	fmt.Printf("GoodMysqlConnection.Execute, sql=%v, args=%v\n", sql, args)
	return nil, 0
}
