// Code generated by go-bindata.
// sources:
// meta/config.yaml
// DO NOT EDIT!

package meta

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

var _metaConfigYaml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x93\xc1\x8e\x9b\x30\x10\x86\xef\x7e\x8a\x51\x72\xae\x51\x7b\xcc\x0d\xd1\x55\x1b\x69\xb7\xaa\x4a\xdb\xbb\xb1\x67\xc3\x28\x66\x6c\xd9\x63\xaa\x7d\xfb\x8a\x44\xb4\x44\x48\xb0\x27\x06\xfd\xdf\x7c\x0c\xb6\xe6\x08\x4d\xe0\x57\xba\x94\x64\x84\x02\x67\x78\x0d\x09\xea\xef\xe7\x27\x76\x31\x10\x4b\xd6\xd3\x5b\x8b\x69\xc4\x04\x7f\xc8\x7b\x75\x04\x63\x05\x02\x03\xb1\xa3\x91\x5c\x31\x1e\x06\x94\x3e\xb8\x0c\x25\x06\x06\xe9\x11\xec\x83\x54\x1d\x21\xa6\x30\x92\x43\xa7\x95\x92\xb7\x88\x27\x38\x9c\x59\x30\x59\x8c\x12\xd2\x7d\x84\x83\x1a\x8d\x27\x77\x6b\x59\x84\x27\x05\x10\x92\xc3\x74\x82\x8f\x0a\x80\xc3\xef\x7f\xd4\x14\x01\x7c\x80\x43\xd5\xa3\xf1\xd2\xeb\xaf\xb7\x47\xd5\x8a\x91\x92\x0f\xca\x14\xe9\x91\x85\xec\x86\xf3\x93\x02\x28\xbc\x20\xd1\xcd\xff\xbe\xa7\x9f\xd3\xa9\x59\xd7\x0f\xdf\xaa\x9e\xc3\x85\x78\x17\x09\x45\xb6\x99\x9f\xe1\x8a\x0b\xcd\xb5\x74\x98\x18\x05\xb3\x6e\x7c\xc9\x82\x29\x4f\xd3\x24\x29\xb1\xb5\x89\xe2\xc2\x36\x18\xf2\xc4\x17\x4f\x59\xf4\xcb\xbd\x7e\xa6\x2c\x55\x8b\xec\x9e\xa6\xf0\x1d\x68\xe9\xb2\x4d\xd4\xe1\x3e\xfa\x8b\xf3\x1a\x66\x33\x60\x8e\xc6\xa2\xfe\x36\x57\x55\x93\xd0\xc8\x0e\xf3\x05\x65\x1b\x38\xe7\x7a\x34\xe4\x4d\xe7\x6f\xa6\xfc\x96\x05\x87\x7a\xfb\x0a\x2d\xe9\x17\x33\x9d\xd8\x6a\x84\xff\xc9\x67\xf4\xb8\x4c\x3a\x63\xaf\x25\xea\xc6\x13\xb2\x54\x3f\x70\x5e\x95\x35\x72\x5f\x90\x95\xfa\x31\x9d\xf5\x7f\x03\x00\x00\xff\xff\xc3\x0d\x15\x2b\x74\x03\x00\x00")

func metaConfigYamlBytes() ([]byte, error) {
	return bindataRead(
		_metaConfigYaml,
		"meta/config.yaml",
	)
}

func metaConfigYaml() (*asset, error) {
	bytes, err := metaConfigYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "meta/config.yaml", size: 884, mode: os.FileMode(436), modTime: time.Unix(1474268295, 0)}
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
	"meta/config.yaml": metaConfigYaml,
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
	"meta": &bintree{nil, map[string]*bintree{
		"config.yaml": &bintree{metaConfigYaml, map[string]*bintree{}},
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

