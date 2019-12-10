package source

func IsSame(src Source, dst Source) bool {
	return src.ID() == dst.ID()
}

type Source interface {
	ReadRaw() ([]byte, error)
	AbsolutePath() string
	ID() string
}

func Read(s Source) ([]byte, error) {
	return s.ReadRaw()
}
