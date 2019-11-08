package assets

import (
	"bytes"
	"os"
	"path"
	"time"
)

// File An asset file.
type File struct {
	// The full asset file path
	Path string

	// The asset file mode
	FileMode os.FileMode

	// The asset modification time
	Mtime time.Time

	// The asset data. Note that this data might be in gzip compressed form.
	Data []byte

	fs       *FileSystem
	buf      *bytes.Reader
	dirIndex int
}

// Name read file name
func (f *File) Name() string {
	return path.Base(f.Path)
}

// Mode read file mode
func (f *File) Mode() os.FileMode {
	return f.FileMode
}

// ModTime read file time
func (f *File) ModTime() time.Time {
	return f.Mtime
}

// IsDir determine whether the file is a path
func (f *File) IsDir() bool {
	return f.FileMode.IsDir()
}

// Size read file size
func (f *File) Size() int64 {
	return int64(len(f.Data))
}

// Sys todo
func (f *File) Sys() interface{} {
	return nil
}

// Close close file
func (f *File) Close() error {
	f.buf = nil
	f.dirIndex = 0

	return nil
}

// Stat read file information
func (f *File) Stat() (os.FileInfo, error) {
	return f, nil
}

// Readdir read information for all files in the path
func (f *File) Readdir(count int) ([]os.FileInfo, error) {
	if f.IsDir() {
		ret, err := f.fs.readDir(f.Path, f.dirIndex, count)
		f.dirIndex += len(ret)

		return ret, err
	}
	return nil, os.ErrInvalid
}

// Read read file content
func (f *File) Read(data []byte) (int, error) {
	if f.buf == nil {
		f.buf = bytes.NewReader(f.Data)
	}

	return f.buf.Read(data)
}

// Seek implements the io.Seeker interface.
func (f *File) Seek(offset int64, whence int) (int64, error) {
	if f.buf == nil {
		f.buf = bytes.NewReader(f.Data)
	}

	return f.buf.Seek(offset, whence)
}
