package e2

/**
ICreateServer 创建云主机
*/
type ICreateServer interface {
	CreateServer(cpu, mem float64) error
}
