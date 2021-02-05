package good

type IGoodDBConnection interface {
	Execute(sql string, args ...interface{}) (error, int)
}