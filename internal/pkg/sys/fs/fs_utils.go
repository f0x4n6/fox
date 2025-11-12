package fs

import (
	"io"
)

func Map(file File) ([]byte, error) {
	b, err := io.ReadAll(file)

	if err != nil {
		return nil, err
	}

	_, err = file.Seek(0, io.SeekStart)

	if err != nil {
		return nil, err
	}

	return b, nil
}

func Open(path string) File {
	f, _ := _fs.Open(path)
	return f
}

func Create(path string) File {
	f, _ := _fs.Create(path)
	return f
}

func Exists(path string) bool {
	return _fs.Exists(path)
}
