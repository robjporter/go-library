package xpath

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	SLASH              = string(os.PathSeparator)
	DEFAULT_DIR_ACCESS = 0755
)

const (
	All = iota
	Files
	Folders
)

type Path struct {
	path string
}

func New(path string) (Path, error) {
	p := Path{}
	if path == "" {
		fullPath, err := os.Getwd()
		if err != nil {
			return Path{}, err
		}
		p.UpdatePath(fullPath)
	} else {
		p.UpdatePath(path)
	}
	p.ensurePathEndSlash()
	return Path{path: path}, nil
}

func (p *Path) ensurePathEndSlash() *Path{
	if p.path[len(p.path) - 1:len(p.path)] != SLASH {
		p.path += SLASH
	}
	return p
}

func (p *Path) UpdatePath(path string) *Path {
	p.path = path
	return p
}

func (p *Path) SplitPath() []string {
	return strings.Split(p.path, SLASH)
}

func (p *Path) ParentPath() string {
	list := p.SplitPath()
	var isAbs bool = false
	if strings.HasPrefix(p.path, SLASH) {
		list = list[1:]
		isAbs = true
	}

	if strings.HasSuffix(p.path, SLASH) {
		list = list[:len(list)-1]
	}

	if len(list) <= 0 {
		return SLASH
	} else {
		list = list[:len(list)-1]
		if len(list) <= 0 {
			if isAbs {
				return SLASH
			}
			return ""
		} else {
			parent := strings.Join(list, SLASH)
			if isAbs {
				parent = SLASH + parent
			}
			if parent != "" {
				parent += SLASH
			}
			return parent
		}
	}
}

func (p *Path) GetFileMd5(filename string) string {
	f, err := os.Open(p.path + filename)
	if err != nil {
		panic(err.Error())
	}
	h := md5.New()
	const BUFSIZE = 10240
	buf := make([]byte, BUFSIZE)
	for {
		rlen, err := f.Read(buf)
		if err != nil {
			break
		}
		h.Write(buf[0:rlen])
	}
	return hex.EncodeToString(h.Sum(nil))
}

func (p *Path) GetPath() string {
	file := p.FormatPath()
	pos := strings.LastIndex(file, SLASH)
	return file[0:pos]
}

func (p *Path) BaseName() string {
	list := p.SplitPath()
	if strings.HasSuffix(p.path, SLASH) {
		list = list[:len(list)-1]
	}
	if list != nil && len(list) > 0 {
		return list[len(list)-1]
	}
	return ""
}

func (p *Path) MkDirSpecificMode(dirname string, mode os.FileMode) error {
	exist, err := p.IsExist(p.path + dirname)
	if err == nil {
		if !exist {
			return os.MkdirAll(p.path + dirname, mode)
		} else if exist {
			return nil
		}
	}
	return err
}

func (p *Path) MkDir(dirname string) error {
	exist, err := p.IsExist(p.path + dirname)
	if err == nil {
		if !exist {
			return os.MkdirAll(p.path + dirname, DEFAULT_DIR_ACCESS)
		} else if exist {
			return nil
		}
	}
	return err
}

func (p *Path) IsDir(dirname string) (bool, error) {
	fi, err := os.Stat(p.path + dirname)
	if err != nil {
		return false, err
	}
	if fi.IsDir() {
		return true, nil
	}
	return false, nil
}

func (p *Path) IsExist(name string) (bool, error) {
	_, err := os.Stat(p.path + name)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (p *Path) ListFilesRecursive(prefix string, pa string, b bool) []string {
	if pa == "" {pa = p.path}
	fileInfos, err := ioutil.ReadDir(pa)
	if err != nil {
		return nil
	}

	list := make([]string, 0, 10)
	var dirName string
	if !b {
		dirName = ""
	} else {
		dirName = p.BaseName() + SLASH
	}
	for _, info := range fileInfos {
		if info.IsDir() {
			tmpList := p.ListFilesRecursive(prefix+dirName, p.path+info.Name()+SLASH, true)
			list = append(list, tmpList...)
		} else if info.Mode().IsRegular() {
			list = append(list, prefix+dirName+info.Name())
		}
	}
	return list
}

func (p *Path) FileMode(filename string) (os.FileMode, error) {
	fi, err := os.Stat(p.path + filename)
	if err != nil {
		return 0, err
	}
	return fi.Mode(), nil
}

func (p *Path) FileSize(filename string) (int64, error) {
	fi, err := os.Stat(p.path + filename)
	if err != nil {
		return 0, err
	}
	return fi.Size(), nil
}

func (p *Path) RelativePath(filename string) string {
	rfile := filename[len(p.path):]
	rfile = strings.TrimLeft(rfile, SLASH)
	return rfile
}

func (p *Path) FormatPath() string {
	p.path = strings.Replace(p.path, "\\", SLASH, -1)
	p.path = strings.TrimRight(p.path, SLASH)
	return p.path
}

// SearchFile Search a file in paths.
// this is often used in search config file in /etc ~/
func (p *Path) SearchFile(filename string) (fullpath string, err error) {
	if fullpath = filepath.Join(p.path, filename); p.FileExists(fullpath) {
			return
		}
	err = errors.New(fullpath + " not found in paths")
	return
}

func (p *Path) FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func (p *Path) RemoveFileIfExist(filename string) error {
	fullpath := filepath.Join(p.path, filename)
	if p.FileExists(fullpath) {
		err := os.Remove(filename)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Path) CreateDirIfNotExist(dirname string) error {
	if !p.FileExists(p.path + dirname) {
		err := os.Mkdir(p.path + dirname, 0744)
		if err != nil {
			return fmt.Errorf("Failed to create dir %v %v", dirname, err)
		}
	}
	return nil
}

func (p *Path) GetFilenameNoExtension(filename string) string {
	n := strings.LastIndexByte(filename, '.')
	if n >= 0 {
		return filename[:n]
	}
	return filename
}

func (p *Path) CopyFile(src, dst string) (err error) {
	sfi, err := os.Stat(src)
	if err != nil {
		return
	}
	if !sfi.Mode().IsRegular() {
		return fmt.Errorf("CopyFile: non-regular source file %s (%q)", sfi.Name(), sfi.Mode().String())
	}
	dfi, err := os.Stat(dst)
	if err != nil {
		if !os.IsNotExist(err) {
			return
		}
	} else {
		if !(dfi.Mode().IsRegular()) {
			return fmt.Errorf("CopyFile: non-regular destination file %s (%q)", dfi.Name(), dfi.Mode().String())
		}
		if os.SameFile(sfi, dfi) {
			return
		}
	}
	if err = os.Link(src, dst); err == nil {
		return
	}
	err = p.copyFileContents(src, dst)
	return
}

func (p *Path) copyFileContents(srcFilename, dstFilename string) (err error) {
	in, err := os.Open(p.path + srcFilename)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(p.path + dstFilename)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}

func (p *Path) FileExtension(filename string) string {
	for i := len(filename) - 1; i > -1; i-- {
		if filename[i] == '.' {
			return filename[i+1 : len(filename)]
		}
	}
	return ""
}
