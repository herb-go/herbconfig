package configuration

func IsSame(src Configuration, dst Configuration) bool {
	return src.ID() == dst.ID()
}

type Configuration interface {
	ReadRaw() ([]byte, error)
	AbsolutePath() string
	ID() string
}

func Read(cf Configuration) ([]byte, error) {
	return cf.ReadRaw()
}
