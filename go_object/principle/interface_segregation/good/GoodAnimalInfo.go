package good

type GoodAnimalInfo struct {
	iID   int
	sName string
}

func (me *GoodAnimalInfo) ID() int {
	return me.iID
}
func (me *GoodAnimalInfo) Name() string {
	return me.sName
}
