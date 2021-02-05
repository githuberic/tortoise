package bad

import "fmt"

type BadDBConnection struct {
	sURL string
	sUID string
	sPWD string
}

func NewBadDBConnection(url string, uid string, pwd string) *BadDBConnection {
	return &BadDBConnection{
		url, uid, pwd,
	}
}
func (p *BadDBConnection) Execute(sql string, args... interface{}) (error, int) {
	fmt.Printf("BadDBConnection.Execute, sql=%v, args=%v\n", sql, args)
	return nil,0
}