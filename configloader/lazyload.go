package configloader

//AssemblerLazyLoader assembler lazy loader struct
type AssemblerLazyLoader struct {
	Assembler *Assembler
}

//LazyLoadConfig lazeload data into interface.
//Return any error if raised
func (l *AssemblerLazyLoader) LazyLoadConfig(v interface{}) error {
	_, err := l.Assembler.Assemble(v)
	return err
}

//NewLazyLoader create new assembler lazy loader
func NewLazyLoader() *AssemblerLazyLoader {
	return &AssemblerLazyLoader{}
}

//LazyLoaderFunc lazy loader func interface
type LazyLoaderFunc func(v interface{}) error

//LazyLoader lazy loader interface
type LazyLoader interface {
	//LazyLoad lazeload data into interface.
	//Return any error if raised
	LazyLoadConfig(v interface{}) error
}
