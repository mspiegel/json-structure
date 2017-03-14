// Code generated by go-bindata.
// sources:
// json-structure.json
// DO NOT EDIT!

package bootstrap

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

var _jsonStructureJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x58\x4d\x6f\x1b\x21\x10\xbd\xfb\x57\x20\xce\x5b\xa9\x6a\x7b\xa8\x7c\x4b\xd5\x4b\xa2\x26\x39\x24\x39\xa5\x3d\x10\x67\xec\x12\x01\xbb\xe2\x23\x4a\x14\xf9\xbf\x57\xcb\xc2\x2e\xb3\x86\xb5\x9b\x66\x53\x9f\x2c\xb3\x8f\x37\xf3\x86\x19\x06\x78\x59\x10\x42\x2d\xb7\x02\xe8\x92\xd0\xb3\xab\xcb\x0b\x72\x65\xb5\x5b\x59\xa7\x81\x48\xb0\xec\x83\x89\x7f\x69\xd5\x62\xd7\x9a\x6d\x24\x28\x6b\xe8\x92\xb4\x93\x09\xa1\xaa\x56\x0a\x36\xcc\xf2\x47\xe8\x07\x09\xa1\x92\x2b\x2e\x9d\xa4\x4b\xf2\xd1\x0f\x6d\xab\x0e\x5e\x37\x96\xd7\x8a\x89\x14\xab\x9c\x10\xec\xce\x3b\x61\xb5\x83\x2a\x8e\xdf\xc3\x9a\x39\x61\xe9\x92\xb4\x08\x44\xb3\xaa\xa5\xac\x55\x4a\x32\x80\xe3\x50\xab\xed\xb9\xf1\xd2\x1e\x4c\xad\x68\x35\x8c\xff\x74\x9f\x3e\x7f\xfd\x42\x97\xe4\x36\xf1\x88\xfc\x0a\x80\x6d\xef\x01\x28\x2f\x21\xc3\x68\xc0\x1e\x40\x98\x20\xb8\x05\x69\x10\xd7\xd8\xbf\xfe\xc3\x76\x91\xfe\x46\xc9\x8d\xe6\x92\xb7\x61\x4e\x59\x90\xe1\x10\x94\xc1\x2c\x5d\xd7\x5a\xb2\x51\x4c\xf6\xbb\xda\x6b\xb4\x9a\xab\x0d\xdd\x89\x4a\xb2\x5e\x99\xc8\xdc\xd5\xb5\x00\x86\xc2\x3d\xac\xcd\x9a\x09\x03\x58\xde\x22\x50\x7b\x82\x24\xb1\xbe\x05\x9e\x44\x6c\xe2\x98\x5b\x0d\xf1\xa7\x6b\x0e\xe2\xde\x94\x65\x26\xa1\xcb\x09\xcd\x2e\x49\x10\x5f\xa5\xdf\x42\x3a\xdc\x0e\x22\xfb\x9c\x29\xad\xda\x55\x47\x73\x94\x22\xc2\xd7\x54\x43\xc2\xdb\x30\x6b\x41\xab\xbf\xa1\x9e\x2e\x2b\x4c\x2f\xb9\xfa\x01\x6a\x63\x7f\x97\x0c\x70\x65\x61\x03\xfa\x00\x0b\x15\xde\x85\x4a\x06\xd9\xd3\x3b\x18\xcc\x26\xc1\x99\xa9\x8f\x34\x8f\xfd\xc6\xb3\xd7\xff\x0b\x27\xef\x40\x1f\xa7\x02\xd5\xf9\x56\x5a\x74\x27\x2c\x6f\x04\x5c\xae\x4b\xec\x61\xfe\x9b\x65\x59\xdf\xf7\xfe\xcd\x5a\x39\x89\xe7\xa4\x87\xa7\x95\x70\x86\x3f\xc2\xf9\xbc\x32\x06\x3b\x33\xe8\xc9\xa6\xf0\x69\x28\xee\x42\xeb\xbc\x88\x59\x34\x99\xc9\xaf\xcf\xd2\xb8\xb7\x1c\xd2\x2f\xda\x82\x3a\xca\x52\x0b\x1e\x14\xd6\x34\xe3\x4f\xc2\x2c\x59\x83\x69\x73\xe7\xa1\x04\x7f\xfd\xdc\xc0\x77\x58\x09\x9a\x7c\xde\xee\x89\xdd\x89\xd6\xec\xf9\x38\x43\xc7\xbc\x6b\x85\xc8\x4d\x1e\x0d\x77\x03\x31\xda\x6e\x4e\xa7\x66\xcf\xd0\x44\xe7\xb7\x97\x5d\xdc\x73\xd6\x94\x8a\xf7\x24\x04\x77\xa6\xda\x6d\x53\x77\x7f\xdd\x82\xfd\x4f\xee\xb5\x37\x91\xbd\xee\xdd\x28\x7e\xac\x47\x10\xe7\x5d\x2b\xe4\x1b\xbe\x13\x8c\x78\xdf\x63\x4f\xb9\x31\xa0\xaf\xb1\xb6\x57\xc7\x6e\xe7\x7e\x76\x58\xdc\xf6\x78\xd8\xab\xca\x78\xd8\xc5\xb6\x5a\x14\xa3\xd9\x5f\x64\x96\x84\xbc\x0c\xf3\xe2\xdd\x0b\x6f\x53\x7d\x0b\x4d\xa1\xb1\xb1\x22\xa8\x8a\x07\x46\x04\x8d\x4d\x36\x45\x9a\x78\x3f\x42\xc8\x70\x6b\x42\xc8\x87\xee\x10\x4d\x30\xd2\x1f\xad\xc7\x8c\x5d\x07\x1d\x33\xfa\xde\x95\x22\x59\xe8\x17\x18\x19\xca\x15\x51\x82\x0d\x96\x11\x25\x8c\xf8\xa4\xdf\xa2\xc6\xb0\x76\xe3\x42\x30\x17\x6a\x11\xc3\xba\x0a\xc5\x40\x13\x62\x88\x81\x31\x23\xc7\xf9\xd0\x5f\xa4\x25\xe3\x43\xad\xe7\xb3\x75\x37\x57\xfb\x77\xa0\x37\x7d\x2a\xe0\xb2\xa9\xb5\x9d\xa8\x89\x69\x5a\x5c\xe1\xd3\x8d\xb2\x58\x2d\xc9\x7b\x92\x59\x69\xee\x0d\xbe\xb5\xce\xdd\x87\xb1\x19\x95\x16\x5e\x8b\xa6\xca\x7c\x1e\x47\x32\x67\x93\x1d\x67\x50\x32\x4e\xce\x4e\xd2\x78\xb1\xfd\x13\x00\x00\xff\xff\xfc\xa0\x17\xa3\x9b\x14\x00\x00")

func jsonStructureJsonBytes() ([]byte, error) {
	return bindataRead(
		_jsonStructureJson,
		"json-structure.json",
	)
}

func jsonStructureJson() (*asset, error) {
	bytes, err := jsonStructureJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "json-structure.json", size: 5275, mode: os.FileMode(420), modTime: time.Unix(1489461195, 0)}
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
	"json-structure.json": jsonStructureJson,
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
	"json-structure.json": &bintree{jsonStructureJson, map[string]*bintree{}},
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