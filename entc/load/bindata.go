// Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// template/main.tmpl
// schema.go
package load

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templateMainTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x91\x41\x6b\xdc\x30\x10\x85\xcf\x9a\x5f\x31\x15\x5b\x22\x81\xa3\xd0\x6b\x61\x4f\xcd\x1e\x52\x68\x52\xd8\x40\x0f\xdb\x25\xc8\xf6\x78\x23\x6a\xcb\xae\xa4\x94\x06\xa1\xff\x5e\x24\xd9\x0b\x3d\xd9\xf2\x7b\xfa\xe6\x3d\x4f\x8c\xd8\xd3\x60\x2c\x21\x9f\xb4\xb1\x1c\x53\x82\xbb\x3b\xfc\x32\xf7\x84\x17\xb2\xe4\x74\xa0\x1e\xdb\x77\xbc\x21\x1b\xba\xeb\xa7\x1b\x85\xf7\x4f\xf8\xf8\xf4\x8c\x87\xfb\x87\x67\x05\x8b\xee\x7e\xe9\x0b\x61\x66\x00\x98\x69\x99\x5d\x40\x01\x8c\xcf\x9e\x03\xe3\xed\x7b\xa0\xfc\x12\x23\x06\x9a\x96\x51\x07\x42\x5e\x5d\xbe\x8c\x2c\xd2\xe2\x8c\x0d\x03\xf2\x8f\xbf\x39\xaa\xef\x2b\x31\x25\x90\x00\x31\xe2\xae\xd5\x9e\xf0\xf3\x1e\xcb\x73\xd3\xf3\xdd\x3f\xda\xa1\xef\x5e\x69\xd2\x1e\xf7\x78\x3a\x93\x0d\xea\xc1\x06\x72\x83\xee\x28\x16\xb4\xd3\xf6\x42\xb8\x7b\x69\x70\x67\xf5\x54\x30\xea\x51\x4f\xe4\x33\x9f\xb1\x18\x6f\x57\x7e\x4a\x2a\x1f\xae\x51\x7c\x4c\x7c\xbd\x93\x52\x53\x58\x64\x7b\xbc\x4d\x09\x12\xc0\xf0\x66\xbb\xd2\x59\x48\x8c\xc0\x72\x90\xd1\x58\xf2\x78\x3a\x9f\xce\xb9\x34\xb0\x61\x76\xf8\xd2\xac\xf9\xf2\xdc\x1a\x65\xcb\x1b\x81\xb1\xb6\x41\x72\x2e\x6b\xdf\xb4\xf3\xaf\x7a\x3c\x16\x51\x54\x8f\x04\xc6\xcc\x50\x1c\x1f\xf6\x68\xcd\x58\xee\xb0\x41\x9b\x51\x90\x73\x59\xce\x15\xea\xdc\x3d\xea\x65\x21\xdb\x8b\x72\x6c\xb0\x95\x90\xd5\xd9\xab\x63\xe8\xe7\xb7\xa0\x7e\x38\x13\x48\x94\x7d\xa8\xaf\xb3\xb1\x9b\xb1\xc6\x15\xfc\xa7\xe5\x52\xca\x6b\xb7\x6d\x4a\x1e\x3f\xbb\x52\xb2\xb2\xc8\xb9\xca\x3a\x06\x67\xec\x25\x7b\xd4\x21\x7b\x84\x94\xc5\x73\xf8\x6b\x82\xf8\x54\x48\xff\x6d\xbd\x96\xaa\x4b\x5f\x7f\x66\x4a\xf0\x2f\x00\x00\xff\xff\xb5\xb1\x2f\xf6\x87\x02\x00\x00")

func templateMainTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateMainTmpl,
		"template/main.tmpl",
	)
}

func templateMainTmpl() (*asset, error) {
	bytes, err := templateMainTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/main.tmpl", size: 647, mode: os.FileMode(420), modTime: time.Unix(1566131522, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x57\xdf\x6f\xdb\xb6\x13\x7f\xb6\xfe\x8a\xfb\x1a\x68\x20\x15\x86\xf2\xee\xc2\x0f\x5f\x6c\x19\x96\x0d\x4b\x87\x76\xdb\x4b\x51\xb4\x8c\x74\x72\xd8\xca\x94\x4b\xd2\x4e\xdb\xc0\xff\xfb\x70\x3f\x68\x49\x96\x9b\x01\x2d\x52\xa0\xb0\x78\xbf\x78\xf7\xb9\x0f\x8f\xcc\xd6\x54\x1f\xcd\x1a\xa1\xed\x4c\x9d\x65\x76\xb3\xed\x7c\x84\x3c\x9b\xcd\xd1\x55\x5d\x6d\xdd\xfa\xf2\x43\xe8\xdc\x3c\x9b\xcd\x9b\x4d\xa4\x1f\x8f\x4d\x8b\x55\x9c\x67\x24\xba\xad\x2e\xd1\xb1\x58\x3f\x2f\x43\x75\x87\x1b\x73\xd9\x58\x6c\xeb\x79\x56\x64\xd9\xe5\x25\xbc\x66\x19\x78\xdc\x7a\x0c\xe8\x62\x00\xe3\x00\x5d\x2c\x55\x11\xef\x4c\x84\x7b\x13\x38\x09\xac\xa1\xf1\xdd\x06\x0c\x54\xdd\x66\xdb\x5a\xac\x61\x17\xd0\x83\x26\x5a\x66\xf1\xcb\x16\x53\xc8\x10\xfd\xae\x8a\xf0\x90\xcd\x6e\xcc\x06\x01\x80\x24\xd6\xad\x01\xe0\x3d\xe5\xbd\x9c\x3b\xb3\xc1\x45\xb7\xb1\x11\x37\xdb\xf8\x65\xfe\x3e\x9b\x5d\xd5\x6b\x0c\x00\xf0\xe6\xed\x73\xfa\x3c\x5a\x22\xc9\xc7\xa6\xbf\x50\x15\x81\x4d\xf9\x33\x99\x72\x75\x27\xb6\xd7\xae\xc6\xcf\x18\xc8\x96\x3f\x93\xad\x15\xf9\xc8\xf8\xc0\xb0\x48\xc8\x29\x2a\x22\xff\x0e\x50\xc4\x71\x8a\xc9\x10\x96\x47\x80\xf9\x8b\x62\xc8\x3f\x2e\xb0\x64\x81\x9a\xd3\x06\x27\xe6\x66\x0d\xdf\x8c\x1e\xcd\x7a\x6c\xfd\xda\x7e\x4d\xc1\x9f\x5b\x17\xf5\x53\xad\x83\xfd\x7a\x12\xfc\xa7\x3b\xe3\x03\xb2\xd9\xf3\x3e\xba\x9a\x57\xa2\x1c\x7b\xfc\xed\xec\xa7\x9d\x6c\x71\xdb\x75\xed\x78\x83\x1d\x2b\xc7\x0e\x37\xb6\x6d\xcd\x6d\x8b\x67\x1d\x9c\x2a\xc7\x2e\x2f\xb7\xd1\x76\xce\xb4\x67\x5d\x3a\x55\x8e\x5d\x7e\xc6\xc6\xec\xda\x78\x3e\xad\x5a\x94\x63\x8f\x7f\x4c\x6b\x6b\x13\x3b\x1f\xa0\x07\x2a\x79\xec\x8f\xca\x33\x84\x62\x3a\x4f\xf9\xc4\xe2\xef\xa0\x13\xfb\x9d\x61\x93\xf6\xe3\xbf\x79\x34\x36\x7c\x84\x41\x27\x86\xa7\xdc\x79\x85\x8d\x6c\x3e\xb6\xf3\xd8\xbc\x9b\xee\xfe\x0a\x1b\xa5\xd9\xe8\x74\x7b\x6c\xbe\xc1\x17\x6d\xcb\x23\x4c\xb9\x76\x7b\xf4\x01\x4f\x4d\xad\x88\x4f\xb7\xff\xb4\xb3\x1e\xeb\x13\x5b\xaf\xe2\x33\x5d\x93\x69\x31\x6d\x9b\xc8\xbf\xa3\x6f\xe2\xd8\x37\x4e\x2b\x3d\xd2\xef\x91\x4a\x05\xb2\xc9\x10\xa5\xd1\x78\x76\x32\xbe\x79\x3b\x6e\xc9\x99\xc1\x28\x45\xde\xe0\x3d\xc7\xae\x3c\x9a\x88\x5c\xa3\x16\x44\xb1\xa5\x2a\xb9\x3a\x88\xf4\xe8\x1b\x53\x61\x99\x35\x3b\x57\x25\xd7\x1c\x8f\x64\x2e\xb4\xb7\x0f\xd9\xcc\x21\x2c\x57\x70\x41\xcb\x87\x6c\xc6\x1c\x5d\x72\x91\x58\xd2\x77\x5e\x2c\xb2\x19\x13\x32\x49\xe9\x5b\xa5\x66\x2d\x42\x92\x9a\xb5\x08\x95\x6a\x4b\x12\xea\xb7\x28\x04\xc4\x25\x5b\x5f\x07\x59\x89\x46\xd9\xb1\x14\x8d\xae\x52\x34\xe9\xfa\x92\x55\x69\xc5\xba\x43\x36\xb3\x0d\x70\xf6\x58\xfe\x3f\x84\xae\xca\x8b\x17\x80\xf0\xbf\x15\x38\xdb\x52\x69\x33\xc7\x29\xc0\xaa\x47\xa0\x60\x3f\x8f\x71\xe7\x1d\x38\x54\x6c\xff\x30\x3e\xdc\x99\x56\xaf\x44\xbe\xb7\xe9\xbc\xe3\xf0\x8a\x3d\x82\x4a\x5f\x1d\x18\xf8\xed\xf5\xcb\x1b\x72\x66\x7e\x55\xc6\xc1\x2d\x42\x8d\xe4\x5a\x8b\x09\x05\x50\xe7\xee\xf6\x03\x56\x51\x7f\xb4\x2b\xa3\x4d\xf3\x90\xf6\x26\xda\xea\x4e\x05\xe4\xb7\xf0\xe6\xed\xed\x97\x88\x0b\x40\xef\xe9\x7f\xe7\x0b\x2a\x2d\x70\xd3\xc4\xf7\x41\xf0\xb6\xae\xb6\x1e\xab\x98\xeb\xfb\x82\x1b\xf5\xb2\xd1\xc8\x45\xa1\xed\x3c\x64\x33\xe5\x18\x87\x5c\xae\x20\x98\x06\x85\x8d\xc9\x56\x90\xf5\x7e\x88\x65\xc2\xcc\xb6\x0b\x68\x36\xb1\xbc\xa2\x5c\x9a\x7c\xae\x89\x3f\xfb\xb4\x84\x67\xfb\xf9\x02\x02\xef\xc3\xc1\x05\xec\xa6\xf3\xf0\x6e\x01\x0d\x6d\xe5\x8d\x23\xae\x0a\xf5\x29\x6a\x60\xf1\x05\x6f\x4f\xeb\x01\xff\x00\x9a\x01\x03\x07\x14\x24\x45\x4f\xc2\x01\x0b\x59\x91\x78\x38\xe0\x1b\xc9\xc7\x8c\x4b\x97\xca\x52\x94\xbf\x9a\xa0\x02\x55\xa7\x9b\x6d\xa9\xbe\x69\xad\xea\x74\x8b\x25\x75\x5a\xab\xba\xbf\x7f\x96\xd0\xa2\xcb\x9b\xb2\x97\xe4\x05\xdb\x1c\xb2\x19\x61\x1c\x16\xd0\x7d\x24\x04\x9a\x32\x97\x17\x03\xdd\xf2\xbe\x78\x41\x62\xc6\x83\x2e\x76\x6e\x12\x6b\xf2\x82\x65\x0d\x2f\x60\x05\x17\xa4\xee\xc3\x55\x93\x70\xfa\x0a\x18\x85\xd4\xcb\x9f\xec\xaa\x64\x70\x0c\x9c\x9e\x0d\x2b\xb8\x50\x3b\x0d\x1f\x4a\x9d\x58\x2b\x30\xdb\x2d\xba\x3a\x4f\x92\x05\x84\x46\x5a\x2d\x0f\xc0\x21\xaf\xf8\xa9\xf8\x94\xb4\xc2\x9e\x56\xbc\xbb\xb0\xaa\x94\x27\xea\x20\xd5\x2b\x49\xad\x1f\x04\x12\x25\xbd\x2d\x87\x39\xeb\x3b\xf4\x29\xb3\xb6\xf5\xe7\x3e\x6f\xcd\x41\x33\x4f\xaf\xe0\x41\xee\xd7\x29\xc9\x0b\xfe\xe2\x26\x52\x15\x44\x3f\x5b\x7f\xe6\xda\x94\x7a\xd2\x91\x25\x8b\xf5\x4c\x8f\x4f\x03\x29\xc6\x67\xe1\x30\x9a\x89\x74\x0b\x95\x3a\x9a\xf2\x50\xe8\x80\xec\x47\x04\xdc\x7b\xb3\x0d\x3c\xdb\xa4\xd8\x44\x8b\x0d\xc6\xbb\xae\x86\x7b\x1b\xef\xc0\x63\xd5\xed\xd1\x43\xec\x00\x5d\xd8\x79\x04\xd7\xc1\xd6\x38\x5b\xd1\x8b\x0c\x36\x12\xde\xba\xb5\x8e\xc2\xc9\x04\x9a\xcc\xc1\x26\xdd\x96\xc7\xe7\xfd\xe9\x44\xac\xb1\x41\x0f\x14\x2e\xe7\x35\x75\x6d\xcf\x20\x4b\x32\x74\x39\xec\x87\x3d\x9c\x91\xff\xea\x4c\xfb\x52\x45\x92\xb0\x76\x72\xcf\x07\xa4\x49\x47\xc0\xd9\x56\xce\xc5\x81\x4e\x8e\x62\x37\x72\xcf\x8b\x05\x5b\xf5\x00\x0a\x27\x27\xf8\x89\xf8\x47\xe1\x1b\x1e\xb4\x09\x7a\x72\x32\x04\x3c\x32\x7c\x42\xec\xa4\x9a\x33\xd0\xa1\x9e\xc8\xc7\x90\x93\x22\x26\xc0\xa5\x23\x31\x81\x2e\x29\x7e\x14\xbc\xf1\x89\x9f\xc0\x67\x8f\x7f\x98\x1e\x5f\x95\x4f\x88\x60\x2a\xea\x0c\x86\xf6\x38\x1b\x1e\x43\x31\x55\xd3\xe3\xc8\x85\x1e\xdf\x07\x11\x86\x2f\x84\x62\xb4\xa2\xdc\x68\x46\xc5\xf2\x77\xeb\xea\xbc\x80\xd5\xea\xa8\xff\x33\x7a\x4e\x9d\x2e\x87\x58\x5e\xb5\xb8\xc9\x47\xa3\x23\x66\x87\xec\xdf\x00\x00\x00\xff\xff\x8a\x04\xaf\xc2\x0d\x11\x00\x00")

func schemaGoBytes() ([]byte, error) {
	return bindataRead(
		_schemaGo,
		"schema.go",
	)
}

func schemaGo() (*asset, error) {
	bytes, err := schemaGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema.go", size: 4365, mode: os.FileMode(420), modTime: time.Unix(1566214133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"template/main.tmpl": templateMainTmpl,
	"schema.go":          schemaGo,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"schema.go": &bintree{schemaGo, map[string]*bintree{}},
	"template": &bintree{nil, map[string]*bintree{
		"main.tmpl": &bintree{templateMainTmpl, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}