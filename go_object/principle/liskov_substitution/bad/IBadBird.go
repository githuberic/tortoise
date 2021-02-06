package bad

type IBadBird interface {
	ID() int
	Name() string

	Tweet() error
	Fly() error
}
