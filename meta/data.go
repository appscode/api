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

var _metaConfigYaml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x93\xc1\x6e\xdb\x30\x0c\x86\xef\x7a\x0a\x22\x39\xcf\x46\x76\xcc\x2d\xcb\x8a\xcd\x40\x07\x14\x4b\xb7\xbb\x2c\xb1\x31\x11\x99\x12\x44\xca\x43\xde\x7e\x70\x82\x64\x2e\x8c\xb6\xde\xc9\xb0\xf9\x7d\xbf\x28\x59\x5c\xc3\x3e\xf2\x0b\x1d\x4b\xb6\x4a\x91\x05\x5e\x62\x86\xdd\x53\xf3\xc0\x3e\x45\x62\x95\x6a\x7c\x3b\x60\x1e\x30\xc3\x1f\x0a\xc1\xac\xc1\x3a\x85\xc8\x40\xec\x69\x20\x5f\x6c\x80\x1e\xb5\x8b\x5e\xa0\xa4\xc8\xa0\x1d\x82\x7b\x15\x6a\xd6\x90\x72\x1c\xc8\xa3\xaf\x8c\xd1\x73\xc2\x2d\xac\x1a\x56\xcc\x0e\x93\xc6\x7c\x6d\x61\x65\x06\x1b\xc8\x5f\x94\x49\x71\x6b\x00\x62\xf6\x98\xb7\xb0\x31\x00\x1c\x7f\xdf\xa9\xb1\x04\xf0\x09\x56\x75\x87\x36\x68\x57\x7d\xbf\x3c\xea\x83\x5a\x2d\xb2\x32\xb6\x68\x87\xac\xe4\xde\xc9\xfc\x6c\x00\x0a\x4f\x48\xf4\xb7\xbd\x7f\x14\x7f\xab\x8e\x72\x35\x6c\x5a\x54\xbb\xa9\x76\xaf\xd6\xac\x1f\xe3\x91\x78\x31\x1a\x8b\x2e\x63\x9f\xe3\x09\x27\xb1\xa7\xd2\x62\x66\x54\x94\xbb\xb0\x0f\x45\x14\xb3\x8c\xdd\x66\x2d\xe9\xe0\x32\x25\x5d\xa6\x34\x2c\x6a\xd9\xe1\x97\x73\xf3\xf4\xcf\xe8\x2d\x05\xe2\x63\x20\xd1\xbb\xf2\xe3\xfa\xed\x91\x44\xeb\x03\xb2\x7f\x18\xa1\xff\x50\x4a\x2b\x2e\x53\x8b\xcb\x95\x5f\x2c\x73\x89\x6d\x8f\x92\xac\xc3\xbb\xf2\x8c\xb6\x97\x7a\x9f\xd1\xea\x02\xee\x1b\xea\xc7\x50\x23\xbb\xc1\x52\xb0\x6d\xb8\x24\xca\x59\x14\xfb\xdd\xfb\x17\xc7\xd1\x64\x13\xe3\xe1\xce\x5a\x9a\x13\x5f\x31\xe0\x94\x68\xad\x3b\x95\x34\xf9\x4b\x84\xac\x52\xff\xc4\xdb\x90\xbd\xcd\x5e\xe7\x76\x7e\x0e\x6f\x60\xb7\x95\xff\x06\x00\x00\xff\xff\x13\xe0\xc6\x5b\x15\x04\x00\x00")

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

	info := bindataFileInfo{name: "meta/config.yaml", size: 1045, mode: os.FileMode(420), modTime: time.Unix(1476262750, 0)}
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

