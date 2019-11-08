package assets

import (
	"bytes"
	"net/http"
	"os"
	"path"
	"time"
)

// FileSystem An in-memory asset file system. The file system implements the
// http.FileSystem interface.
type FileSystem struct {
	// A map of directory paths to the files in those directories.
	Dirs map[string][]string

	// A map of file/directory paths to assets.File types.
	Files map[string]*File

	// Override loading assets from local path. Useful for development.
	LocalPath string
}

// NewFileSystem create a file system instance
func NewFileSystem(dirs map[string][]string, files map[string]*File, localPath string) *FileSystem {
	fs := &FileSystem{
		Dirs:      dirs,
		Files:     files,
		LocalPath: localPath,
	}

	for _, f := range fs.Files {
		f.fs = fs
	}

	return fs
}

// NewFile create a file
func (f *FileSystem) NewFile(path string, fileMode os.FileMode, mTime time.Time, data []byte) *File {
	return &File{
		Path:     path,
		FileMode: fileMode,
		Mtime:    mTime,
		Data:     data,

		fs: f,
	}
}

// Open Implementation of http.FileSystem
func (f *FileSystem) Open(p string) (http.File, error) {
	p = path.Clean(p)

	if len(f.LocalPath) != 0 {
		return http.Dir(f.LocalPath).Open(p)
	}

	if fi, ok := f.Files[p]; ok {
		if !fi.IsDir() {
			// Make a copy for reading
			ret := fi
			ret.buf = bytes.NewReader(ret.Data)

			return ret, nil
		}

		return fi, nil
	}

	return nil, os.ErrNotExist
}

func (f *FileSystem) readDir(p string, index int, count int) ([]os.FileInfo, error) {
	if d, ok := f.Dirs[p]; ok {
		maxl := index + count

		if maxl > len(d) {
			maxl = len(d)
		}

		ret := make([]os.FileInfo, 0, maxl-index)

		for i := index; i < maxl; i++ {
			ret = append(ret, f.Files[path.Join(p, d[i])])
		}

		return ret, nil
	}

	return nil, os.ErrNotExist
}
