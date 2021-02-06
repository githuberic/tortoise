package good

type IGoodBird interface {
	ID() int
	Name() string

	// 是否可用鸣叫
	Tweet() error
}

type IFlyableBird interface {
	IGoodBird

	Fly() error
}

type IRunnableBird interface {
	IGoodBird

	Run() error
}