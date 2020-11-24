// Code generated by go-bindata. DO NOT EDIT.
// sources:
// assets/emr-containers-rbac.yaml (1.46kB)

package authconfigmap

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _emrContainersRbacYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x94\x4f\x8f\xd3\x40\x0c\xc5\xef\xf3\x29\x46\x3d\x6f\x8a\xb8\xa1\x1c\xb9\x70\x5f\x09\x2e\x88\x83\x33\x79\xed\x7a\x77\x32\x1e\xd9\x93\xf2\xe7\xd3\xa3\x99\x34\x2c\x82\x14\x75\x11\xa2\xa7\xba\x76\xf2\xfc\xf3\x53\xec\xae\xeb\x1c\x65\xfe\x00\x35\x96\xd4\x7b\x1d\x28\xec\x69\x2e\x0f\xa2\xfc\x8d\x0a\x4b\xda\x3f\xbd\xb1\x3d\xcb\xab\xd3\x6b\xf7\xc4\x69\xec\xfd\xbd\x44\xb8\x09\x85\x46\x2a\xd4\x3b\xef\x13\x4d\xe8\x3d\x26\xed\x82\xa4\x42\x9c\xa0\xe6\x74\x8e\xb0\x5a\xed\x3c\x65\x7e\xa7\x32\x67\xeb\xfd\xc7\xdd\xee\x93\xf3\xde\x7b\x85\xc9\xac\x01\x2d\x57\x05\x2c\x53\x80\x9d\xab\x27\xe8\xd0\x2a\x47\x94\x96\xba\x42\xc4\xa0\x27\x0e\xa0\x10\x64\x4e\xc5\x76\x77\x7e\x4d\xb5\x38\x48\x3a\xf0\x71\xa2\xdc\xfe\xe1\x84\xf3\x33\x59\xc6\xcd\xae\x77\x7e\x17\xd9\xda\xef\x67\x2a\xe1\xa1\x06\x23\x2c\x28\x0f\x68\x7a\x0a\x2a\x2d\xc2\xc8\x65\xa9\x46\x2c\x99\x25\x0a\x12\x23\x42\xb5\xb0\xe6\x28\x25\x29\xe7\x37\xf2\x2a\x18\x69\x40\xbc\x7e\xc0\xa0\x28\xbf\xc1\x3e\x83\xe4\x67\xce\x95\x64\x41\xdf\x68\x40\x39\xdb\x66\x93\xca\x78\x98\xa3\x61\xb1\x67\x44\x8e\xf2\x75\x6a\x6e\xfd\x63\x97\x5e\xe6\xc8\xf0\x63\x92\x5f\x88\x1f\x65\xb8\x31\x1a\xbe\x14\xa4\xba\x3e\x9b\x8e\x72\x3a\x2a\xcc\xb6\xbf\xed\xff\x07\x79\x71\xb1\xb7\x98\x55\xe2\xb2\x35\x35\x18\x38\x8d\x9c\x8e\xb7\xdf\x12\xf7\x97\xc7\xea\xed\x32\xc0\x15\x37\x4b\x22\xee\x71\xa8\xf5\xd5\xbb\x3f\x34\x71\xde\xff\x74\x10\x2f\x48\xda\x3c\x3c\x22\x14\xeb\x5d\xf7\x22\xcd\xf7\x06\xbd\xa4\xf9\x3d\x00\x00\xff\xff\x37\x5e\x01\xcb\xb4\x05\x00\x00")

func emrContainersRbacYamlBytes() ([]byte, error) {
	return bindataRead(
		_emrContainersRbacYaml,
		"emr-containers-rbac.yaml",
	)
}

func emrContainersRbacYaml() (*asset, error) {
	bytes, err := emrContainersRbacYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "emr-containers-rbac.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xef, 0xb5, 0x76, 0xfa, 0xd2, 0xf2, 0x5a, 0xeb, 0xd, 0xc9, 0xc0, 0xf1, 0xad, 0xfd, 0x89, 0x1c, 0x33, 0x5a, 0x53, 0x82, 0xd8, 0x5a, 0xc0, 0x62, 0x87, 0x6f, 0x91, 0xdd, 0xc8, 0x4e, 0x5c, 0xd6}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"emr-containers-rbac.yaml": emrContainersRbacYaml,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"emr-containers-rbac.yaml": {emrContainersRbacYaml, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
