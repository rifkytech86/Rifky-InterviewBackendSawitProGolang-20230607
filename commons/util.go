package commons

import "io/ioutil"

//go:generate mockery --name IFileReader
type IFileReader interface {
	ReadFile(filename string) ([]byte, error)
}

type fileReader struct{}

func NewFileReader() IFileReader {
	return &fileReader{}
}
func (f *fileReader) ReadFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}
