package configurator

import "os"

type OsWrapper interface {
	ReadFile(name string) ([]byte, error)
	Stat(name string) (os.FileInfo, error)
	WriteFile(name string, data []byte, perm os.FileMode) error
}

type BuiltInOs struct{}

func (bio BuiltInOs) ReadFile(filePath string) ([]byte, error) {
	return os.ReadFile(filePath)
}

func (bio BuiltInOs) Stat(filePath string) (os.FileInfo, error) {
	return os.Stat(filePath)
}

func (bio BuiltInOs) WriteFile(name string, data []byte, perm os.FileMode) error {
	return os.WriteFile(name, data, perm)
}
