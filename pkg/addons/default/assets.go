// Code generated by go-bindata.
// sources:
// assets/aws-node-1.10.yaml
// assets/aws-node.yaml
// assets/coredns-1.11.json
// assets/coredns-1.12.json
// assets/coredns-1.13.json
// DO NOT EDIT!

package defaultaddons

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

var _awsNode110Yaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x56\xdf\x6f\xdb\x36\x10\x7e\xd7\x5f\x71\x70\x1f\x0a\x0c\x95\x1c\x75\x59\x97\xe9\xcd\x75\xbc\xac\x58\xe2\x1a\xf1\x9a\x62\x28\x02\x83\xa6\x2e\x32\x67\x8a\xe4\xc8\xa3\x1d\xf7\xaf\x1f\x28\xf9\x87\x14\xcb\xee\x36\x0c\x18\x5f\x2c\xf3\xee\xbe\xe3\x7d\xfc\xc8\x63\x1c\xc7\x11\x33\xe2\x01\xad\x13\x5a\x65\x60\xe7\x8c\x27\xcc\xd3\x42\x5b\xf1\x95\x91\xd0\x2a\x59\x5e\xb9\x44\xe8\xfe\x2a\x8d\x5e\xc1\xd2\xcf\xd1\x2a\x24\x74\xb0\xaa\x43\x1c\xcc\xf1\x49\x5b\x84\x34\xb9\x4a\x2e\xc0\x2d\xb4\x97\x39\x78\x87\x67\xa1\xe6\x48\x2c\x8d\x96\x42\xe5\x19\x0c\xa5\x77\x84\xf6\x5e\x4b\x8c\x4a\x24\x96\x33\x62\x59\x04\xa0\x58\x89\x19\xb0\xb5\x8b\x95\xce\x31\xb2\x5e\xa2\xcb\xa2\x18\x98\x11\x37\x56\x7b\xe3\x82\x53\x0c\xdc\xe6\x15\x2e\x2b\xd9\x57\xad\xd8\xda\x25\x5c\x97\x11\x80\x45\xa7\xbd\xe5\xb8\x75\xeb\x7d\xd7\xab\x7e\x03\xaa\x33\x8c\xa3\x8b\x20\xd4\x30\x6f\xd8\x9b\xd8\xf0\xa5\xd7\x7b\x3c\x86\x31\x3a\x77\x35\x8e\xce\xd1\x9d\x42\x84\x2f\x3d\x29\x1c\xf5\xde\x40\x6f\xcd\x88\x2f\xc2\x47\x81\xd4\x7b\x7c\x99\x02\x9f\x09\x55\x45\x63\x57\xb2\x9c\x61\xa9\x95\x43\x3a\x83\xfc\x18\xbd\xdc\xc2\xd5\x8e\xd8\x29\xda\x95\xe0\x38\xe0\x5c\x7b\x45\xe7\xb8\x85\x43\x11\x59\xb5\xc7\xb1\xdb\x38\xc2\xf2\x08\xfb\xff\x95\xc7\x7b\xa1\x72\xa1\x8a\xb3\x2a\xd1\x12\xef\xf1\x29\x58\x76\x44\x9f\x59\x75\x04\x70\xac\xc1\x23\x4c\xe7\xe7\x7f\x20\xa7\x4a\x7c\x9d\xcc\xfe\x33\x3e\x6b\x88\xeb\x6a\x6f\xa7\x48\x2d\x7e\x99\x31\xee\x6f\x50\xf9\x53\x9b\xca\x83\x8a\xf6\xdc\xfd\x9b\xcd\x06\x90\x6c\x8e\xb2\x12\x1f\xc0\xf2\xca\xc5\xcc\x98\x26\x0f\x06\x79\xb0\x79\x93\x33\xc2\x29\x59\x46\x58\x6c\x6a\x6f\xda\x18\xcc\xe0\x5e\x4b\x29\x54\xf1\xa9\x72\x88\x00\x1c\x4a\xe4\xa4\x6d\xed\x53\x06\xc1\xde\x36\x52\x74\x25\x01\x20\x2c\x8d\x64\x84\xdb\xa0\x46\x21\x61\xc8\x56\x7c\x37\x42\x18\x4c\x29\x4d\xd5\x5e\x37\x9c\x1d\x5f\x60\xee\x25\xda\x84\x49\xb3\x60\xc9\x81\xe4\xa0\x3b\x6e\x05\x09\xce\x64\x6c\x74\x9e\xc1\xeb\xd7\x55\xd8\xae\xe8\xea\xbb\xb5\xed\xe3\x97\xb4\x86\xb1\xd0\x8e\xc6\x48\x6b\x6d\x97\x19\x90\xf5\xbb\x79\xd2\x12\x6d\x7b\x39\x31\x68\x13\xe6\xb4\xcd\x60\xf4\x2c\x5c\x75\xca\xc3\xe0\x5a\x11\x13\x0a\x6d\xc3\x55\x94\xac\xc0\x0c\xde\x5d\xbc\xbd\xbc\x48\xd3\xcb\xef\x2f\x7f\x78\x9b\xe4\x4b\x9b\x20\xb7\x89\x77\xf1\x1a\x1d\xc5\x6f\xdb\x77\x60\xbf\xfe\x17\x07\x86\xb8\x12\xd9\x2a\x4d\x2e\x93\x74\xcf\x45\x85\x38\xf1\x52\x4e\xb4\x14\x7c\x93\xc1\x40\xae\xd9\xc6\xed\xed\x46\x5b\x6a\x50\x17\x1f\x96\x35\xd1\x96\x32\x78\x97\xbe\xfb\xf1\x6a\x6f\xde\xa9\xac\x44\xb2\x82\x1f\x50\x8e\xb4\x57\x0f\x54\xab\xac\x11\x1b\x6f\xfd\x06\x9f\xa7\xb3\x87\xc9\x70\xf6\xeb\xd5\x74\x36\x1c\x7f\x98\xdd\x7e\xbc\xb9\x1d\x3d\x8c\x6e\x1b\xae\x00\x2b\x26\x3d\x66\x70\x3d\x7a\xff\xe9\xa6\x03\xe3\xee\xf7\xd9\xf8\xe3\xf5\x68\x36\x1e\xdc\x8d\x8e\xe3\x7e\xb6\xba\xcc\x5a\xd3\x00\x4f\x02\x65\xbe\xbd\x34\x3a\x2c\x13\x46\x8b\xac\xd2\x41\x12\x6a\x08\xdb\xde\x91\xf6\xf3\xe0\xb7\xe1\x2f\x55\xd2\xe9\x64\x30\xfc\x2f\x33\xef\x4e\x40\xb2\x3f\xb6\x7b\xef\x56\xbf\x38\x4c\xfe\xe9\xd1\x91\x6b\x83\x72\xe3\x33\x48\x2f\xca\xc3\x59\x40\xee\xad\xa0\xcd\x50\x2b\xc2\x67\x6a\x7a\x1b\x2b\x56\x42\x62\x81\x79\x4b\xc3\x00\x2b\x2d\x7d\x89\x77\x41\xfd\x2d\x69\x94\x61\xa6\x5e\x6d\x3f\x9c\x80\xbe\x36\xd4\xe7\x4a\xf4\xe7\x42\x1d\x49\x84\x2b\x11\xcf\x85\x8a\x73\x61\xcf\x41\x20\xf1\x0a\x42\x21\x25\x79\x27\x88\x42\xfa\x16\xc8\x8a\xd9\xbe\xd4\xc5\x51\xb8\xd4\xc5\x99\xd0\x10\x65\xbd\xea\xe7\x9a\x2f\xd1\x26\x4e\xf3\xe5\x11\x42\x6d\x6b\x98\x6a\x6e\x1a\x47\xf6\x74\xb5\x61\x69\x55\xaa\x26\xe7\x75\xea\x63\xe2\xe2\x33\x15\x9f\x01\xea\xa2\x2f\x3e\x51\xfd\x19\x98\x36\x81\xf1\xa9\xe2\xbf\x89\xf1\x92\xce\x97\x0f\x0b\x66\xc4\xa1\x8b\x9d\x78\x08\x78\x47\xba\xbc\xdf\x4a\xfe\x1a\x9f\x84\x12\xe1\x42\xed\xe8\x75\xa8\x04\xd7\xea\x49\x14\x2e\xe9\x7e\x1e\xee\x6e\x75\xc7\x75\xe8\x5b\xdb\xf6\x1f\x01\x14\xf5\x8b\xe1\xd4\xa3\xf2\x15\xa4\x49\x7a\x11\x9a\xae\x83\xde\xb6\x2f\xf7\xde\x00\x93\x12\x14\xae\xd1\x56\xed\x78\x67\x70\xbd\xfa\xd9\xb6\x7b\x96\x55\x3d\x27\xdd\xf5\xdf\x9a\x28\x23\xbd\x65\xb2\xb9\xe2\xba\xeb\x08\x55\x78\xc9\x6c\xc3\x50\x37\xe5\x8a\x89\xd1\xf8\xc3\xb0\x9e\xfb\x2b\x00\x00\xff\xff\x9c\xa8\x6b\xec\xbf\x0b\x00\x00")

func awsNode110YamlBytes() ([]byte, error) {
	return bindataRead(
		_awsNode110Yaml,
		"aws-node-1.10.yaml",
	)
}

func awsNode110Yaml() (*asset, error) {
	bytes, err := awsNode110YamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "aws-node-1.10.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _awsNodeYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x56\xdf\x8f\x1a\x37\x10\x7e\xdf\xbf\x62\x44\x1f\x22\x55\xdd\xe5\x48\xaf\xe9\x75\xdf\x08\x47\xd3\xa8\x77\x04\x1d\xcd\x45\x55\x74\x42\xc6\x3b\x07\x2e\x5e\xdb\xb5\xc7\x70\xe4\xaf\xaf\xec\xe5\xc7\x2e\x2c\xa4\xad\x2a\xd5\x2f\xb0\x9e\x99\x6f\x3c\x9f\x3f\x7b\x9c\xa6\x69\xc2\x8c\x78\x44\xeb\x84\x56\x39\xd8\x19\xe3\x19\xf3\xb4\xd0\x56\x7c\x61\x24\xb4\xca\x96\x37\x2e\x13\xba\xbb\xea\x25\xdf\xc0\xd2\xcf\xd0\x2a\x24\x74\xb0\xaa\x42\x1c\xcc\xf0\x59\x5b\x84\x5e\x76\x93\x5d\x81\x5b\x68\x2f\x0b\xf0\x0e\x2f\x42\xcd\x90\x58\x2f\x59\x0a\x55\xe4\x30\x90\xde\x11\xda\x07\x2d\x31\x29\x91\x58\xc1\x88\xe5\x09\x80\x62\x25\xe6\xc0\xd6\x2e\x55\xba\xc0\xc4\x7a\x89\x2e\x4f\x52\x60\x46\xbc\xb3\xda\x1b\x17\x9c\x52\xe0\xb6\x88\xb8\xac\x64\x5f\xb4\x62\x6b\x97\x71\x5d\x26\x00\x16\x9d\xf6\x96\xe3\xd6\xad\xf3\x6d\x27\xfe\x06\x54\x67\x18\x47\x97\x40\xa8\x61\x56\xb3\xd7\xb1\xe1\x73\xa7\xf3\x74\x0a\x63\x74\xe1\x2a\x1c\x5d\xa0\x3b\x87\x08\x9f\x3b\x52\x38\xea\x7c\x07\x9d\x35\x23\xbe\x08\x7f\xe6\x48\x9d\xa7\xe3\x14\xf8\x42\xa8\x22\x8d\x6d\xc9\x0a\x86\xa5\x56\x0e\xe9\x02\xf2\x53\x72\xbc\x85\xab\x1d\xb1\x13\xb4\x2b\xc1\xb1\xcf\xb9\xf6\x8a\x2e\x71\x0b\x87\x22\xf2\xb8\xc7\xa9\xdb\x38\xc2\xf2\x04\xfb\xff\x95\xc7\x5b\xa1\x0a\xa1\xe6\x17\x55\xa2\x25\x3e\xe0\x73\xb0\xec\x88\xbe\xb0\xea\x04\xe0\x54\x83\x27\x98\xce\xcf\xfe\x40\x4e\x51\x7c\xad\xcc\xfe\x33\x3e\x2b\x88\xdb\xb8\xb7\x13\xa4\x06\xbf\xcc\x18\xf7\x37\xa8\xfc\xa9\x49\xe5\x41\x45\x7b\xee\xfe\xcd\x66\x03\x48\x36\x43\x19\xc5\x07\xb0\xbc\x71\x29\x33\xa6\xce\x83\x41\x1e\x6c\xde\x14\x8c\x70\x42\x96\x11\xce\x37\x95\x37\x6d\x0c\xe6\xf0\xa0\xa5\x14\x6a\xfe\x31\x3a\x24\x00\x0e\x25\x72\xd2\xb6\xf2\x29\x83\x60\xef\x6a\x29\xda\x92\x00\x10\x96\x46\x32\xc2\x6d\x50\xad\x90\x30\x64\x23\xbe\x1d\x21\x0c\xa6\x94\xa6\xb8\xd7\x35\x67\xc7\x17\x58\x78\x89\x36\x63\xd2\x2c\x58\x76\x20\x39\xe8\x8e\x5b\x41\x82\x33\x99\x1a\x5d\xe4\xf0\xea\x55\x0c\xdb\x15\x1d\x86\xb1\x42\x5b\x41\x9b\x81\x64\xce\x8d\x22\xab\x15\x75\x31\x71\xba\x8b\xdf\x7a\xbb\x86\x48\x46\xc7\x9b\x10\xc6\x42\x3b\x1a\x21\xad\xb5\x5d\xe6\x40\xd6\xef\xe6\x49\x4b\xb4\xcd\xc5\xa7\xa0\x4d\x98\xd3\x36\x87\xe1\x8b\x70\xf1\x4e\x08\x83\x6b\x45\x4c\x28\xb4\x35\x57\x51\xb2\x39\xe6\xf0\xe6\xea\xf5\xf5\x55\xaf\x77\xfd\xfd\xf5\x0f\xaf\xb3\x62\x69\x33\xe4\x36\xf3\x2e\x5d\xa3\xa3\xf4\x75\xf3\xc6\xec\x56\x5f\x69\xe0\x93\x2b\x91\xaf\x7a\xd9\x75\xd6\xdb\x33\x17\x11\xc7\x5e\xca\xb1\x96\x82\x6f\x72\xe8\xcb\x35\xdb\xb8\xbd\xdd\x68\x4b\x35\xa2\xd3\xc3\xb2\xc6\xda\x52\x0e\x6f\x7a\x6f\x7e\xbc\xd9\x9b\x77\x9a\x2c\x91\xac\xe0\x07\x94\x13\xa5\x56\x03\xd5\x2a\xaf\xc5\xa6\x5b\xbf\xfe\xa7\xc9\xf4\x71\x3c\x98\xfe\x7a\x33\x99\x0e\x46\xef\xa7\x77\x1f\xde\xdd\x0d\x1f\x87\x77\x35\x57\x80\x15\x93\x1e\x73\xb8\x1d\xbe\xfd\xf8\xae\x05\xe3\xfe\xf7\xe9\xe8\xc3\xed\x70\x3a\xea\xdf\x0f\x4f\xe3\x7e\xb6\xba\xcc\x1b\xd3\x00\xcf\x02\x65\xb1\xbd\x62\x5a\x2c\x63\x46\x8b\x3c\xaa\x26\x0b\x35\x84\x6d\x6f\x49\xfb\xa9\xff\xdb\xe0\x97\x98\x74\x32\xee\x0f\xfe\xcb\xcc\xbb\xf3\x92\xed\x0f\xf9\xde\xbb\xd1\x5d\x0e\x93\x7f\x7a\x74\xe4\x9a\xa0\xdc\xf8\x1c\x7a\x57\xe5\xe1\xe4\x20\xf7\x51\xfa\x5a\x11\xbe\x50\xdd\xdb\x58\xb1\x12\x12\xe7\x58\x34\x34\x0c\xb0\xd2\xd2\x97\x78\x1f\xd4\xdf\x90\x46\x19\x66\xaa\xd5\x76\xc3\x09\xe8\x6a\x43\x5d\xae\x44\x77\x26\xd4\x89\x44\xb8\x12\xe9\x4c\xa8\xb4\x10\xf6\x12\x04\x12\x8f\x10\x0a\x29\x2b\x5a\x41\x14\xd2\xd7\x40\x56\xcc\x76\xa5\x9e\x9f\x84\x4b\x3d\xbf\x10\x1a\xa2\xac\x57\xdd\x42\xf3\x25\xda\xcc\x69\xbe\x3c\x41\xa8\x6c\x35\x53\xc5\x4d\xed\xc8\x9e\xaf\x36\x2c\x2d\xa6\xaa\x73\x5e\xa5\x3e\x25\x2e\xbd\x50\xf1\x05\xa0\x36\xfa\xd2\x33\xd5\x5f\x80\x69\x12\x98\x9e\x2b\xfe\xab\x18\xc7\x74\x1e\x3f\x43\x98\x11\x87\x9e\x77\xe6\xd9\xe0\x1d\xe9\xf2\x61\x2b\xf9\x5b\x7c\x16\x4a\x84\x0b\xb5\xa5\x33\xa2\x12\x5c\xab\x67\x31\x77\x59\xfb\x63\x72\xd7\x03\x1c\xd7\xa1\xcb\x6d\x1f\x0b\x09\xc0\xbc\x7a\x5f\x9c\x7b\x82\xee\x1a\x77\x55\xe5\x8e\x8e\x55\x2f\x36\x9f\x5e\xad\x4d\x1c\x1d\x1d\x47\xda\xc6\x0b\x7c\x3b\x17\x8f\x72\x05\x62\xa4\xb7\x4c\xd6\xd7\x5c\x75\x29\xa1\xe6\x5e\x32\x5b\x33\x54\x4d\x3c\x72\x31\x1c\xbd\x1f\x54\x73\x7f\x05\x00\x00\xff\xff\xd1\x54\xb9\x74\xef\x0b\x00\x00")

func awsNodeYamlBytes() ([]byte, error) {
	return bindataRead(
		_awsNodeYaml,
		"aws-node.yaml",
	)
}

func awsNodeYaml() (*asset, error) {
	bytes, err := awsNodeYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "aws-node.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _coredns111Json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x19\xc9\x72\x1b\xb9\xf5\xae\xaf\xe8\xea\xb3\x49\x91\xd6\x62\x85\x37\x8d\xa4\x78\x54\x65\x29\x2c\x49\x9e\x4b\xec\x9a\x7a\x44\x3f\x92\x88\xd0\x00\x82\x85\x12\xc7\xa5\x7f\x4f\xa1\xf7\x05\x68\x92\x8a\x27\x49\x55\x78\xb0\x5b\x78\x6b\xbf\xfd\xa1\x7f\x1c\x45\x51\x0c\x92\xfe\x86\x4a\x53\xc1\xe3\x59\x14\x6f\xa6\xf1\x07\x77\x4a\x0d\xa6\x3a\x9e\x45\x7f\x3f\x8a\xa2\x28\xfa\x91\xfd\x1b\x42\xce\x20\xcf\x94\x27\xee\xec\x11\xd5\x86\x12\xac\x01\x29\x1a\x48\xc0\x40\x3c\xab\xd8\x38\x46\x9c\x0b\x03\x86\x0a\xae\x5b\x80\x28\x8a\xa5\x12\x29\x9a\x35\x5a\x3d\xa6\xe2\x58\x0a\x65\x1c\xdf\xbf\x4c\xcf\x4e\x2a\xa6\x1e\x3c\x4d\x14\x48\x74\x98\x46\x59\x8c\x2b\xc4\xb7\x9a\x26\x66\xb0\x40\xd6\x93\x87\xcf\x7a\x0c\x29\xfc\x21\x38\xbc\xe8\x31\x11\xe9\x31\x11\xa9\x14\x1c\x79\x26\xf8\xd9\x2e\x70\x94\x70\xdd\x16\xfe\x7c\xa1\x47\x20\xe5\x00\x82\x5d\xa0\xe2\x68\x30\xd3\x8e\x30\xab\x0d\xaa\x91\x2e\xac\x53\xaa\x39\x40\xc2\x21\xcd\xf0\xae\x84\xc2\xeb\xfb\x47\xff\x1b\x95\x48\x1e\x1d\x32\x98\x96\x40\x6a\x04\xbd\xd5\x06\xd3\x92\x53\xc5\x27\xd6\x12\x49\xdb\x3d\xce\xea\xb5\xfb\xf3\x5f\xd3\x68\x0d\xd1\xdd\x37\x2f\xa8\xe3\x59\x74\x76\xd2\x3d\x57\xc2\x08\x22\x98\x23\xfb\x7a\x3d\xef\x92\x19\x50\x2b\x34\xf3\x92\xb8\x01\x7c\xfb\xb0\x97\x1e\x23\x43\xe4\x7b\x74\x79\xba\x3a\x44\x97\xea\xf9\x7b\xc3\xd8\x1a\x19\x12\x23\x54\x37\xb8\x7c\x71\xe2\xf5\xa4\x46\xed\xd2\xea\x72\xb9\xa4\x9c\x9a\xad\x23\xb8\x17\xbc\x19\x21\xb1\xd9\xe6\x01\x7e\x95\x07\xd3\xed\xbc\xf2\xe4\x51\x83\xdb\xe1\xb9\x7a\x49\x88\xb0\xdc\xec\x4a\xd9\x77\x65\x0f\x11\x0a\xf7\x4a\x9e\xc1\xf0\xee\x73\xd9\x1d\xdd\x07\xda\xa4\xff\xbe\x2e\xf1\x96\x94\x65\x02\xc6\xb3\xb3\x93\xe8\xc7\x37\xee\x00\xa8\x94\x50\x3a\x7f\x5e\x23\x30\xb3\xce\x9f\xeb\xf4\x8d\x8a\x74\x1f\x33\x41\x80\x45\x94\x8f\x20\x49\xd4\x18\x94\x84\x88\xca\xf3\xfc\xa1\xe0\x16\x45\x52\x24\x3a\xa2\x5c\x23\xb1\x0a\xcb\x43\x2b\xb5\x51\x08\x69\xf9\xf7\x12\x18\x33\x6b\x25\xec\x6a\xed\xe7\x97\x23\xbe\xe5\xff\xd5\xa5\x31\x9a\xb9\xc2\x59\x9d\xbe\x6e\xa3\x71\x74\x8c\x86\x1c\x2b\xd4\x82\x6d\xc6\x44\xf0\x65\x0e\x25\x40\xd6\x18\x9d\x4c\xbe\xf1\xb7\x6f\xbc\x5f\x23\xca\x98\xb9\x12\x7c\x49\x57\x77\x20\xff\xff\xc2\x05\xa4\xd4\xc7\xbe\x3c\xba\x46\xc9\xc4\x36\xc5\xdd\x39\xd4\x69\x7b\xff\x76\x73\xda\xdb\x5e\x3f\xab\xd1\xbc\xc7\xb4\x43\x7d\x46\x89\x95\x42\xad\xaf\x11\x12\x46\x39\x3e\x22\x11\x3c\x71\x56\xf8\x38\x3d\xfd\x74\x7a\x71\x72\x7e\xfa\xa9\x21\x4a\xa1\x64\x94\x40\x06\x6f\x1d\x6f\xa8\x73\xd2\xaf\x54\x1b\xa1\xb6\x5f\x68\x4a\x9d\x75\xa6\x93\x3d\x0a\x74\x0a\x86\xac\xbf\xf8\x6c\xff\x3e\xeb\xef\x8a\xd7\x66\x03\x69\xd5\x7f\xa3\xc0\xe0\x6a\xdb\xd5\x4f\x09\xc6\x28\x5f\x7d\x95\x09\x18\xec\x6b\x98\xc2\xeb\xa3\x55\x2b\x07\x99\x7e\xe8\x81\xbe\x72\xd8\x00\x65\xb0\xc8\xaa\xd8\x34\xd4\x51\xab\xce\xf2\xd0\x12\xe6\xd5\xd3\x60\x2a\x59\x5f\x15\x7f\xbc\x67\x10\xa2\x30\x8b\xf8\x27\x9a\xa2\x36\x90\x3a\xc3\x70\xcb\x58\x47\x5d\x6f\xfc\xbf\xdf\x0b\xbb\xfd\xd0\xf4\x44\xd7\x20\xbd\x48\xcd\x4e\xa1\x6e\xce\x3d\x25\xb9\x48\xf0\x32\x0c\xcf\xa2\xf4\x9f\x96\x2a\x4c\xae\xad\xa2\x7c\xf5\x48\xd6\x98\x58\x67\xed\xdb\x15\x17\xd5\xf1\xcd\x2b\x12\x6b\xf2\x82\xd3\xe7\x51\xc8\x79\x2c\x62\xf9\x09\x55\xda\x9d\xd2\xea\x9f\x8f\x3c\xaa\x42\xfe\xe6\x55\xaa\x7c\xe4\x08\x73\x18\xe2\x92\x71\x7a\xc6\x6c\x50\x59\xa0\x81\x71\xbb\xa8\x88\xbe\x3f\x5a\x94\x42\xa2\x82\x3c\x1d\xe3\x5b\x3e\x8c\xbb\x01\x66\x71\x58\xcb\x0c\x8f\x51\x6e\x5f\xe3\x01\x9c\xef\x41\xd8\x5b\x58\x81\x77\x1a\x00\x14\x59\xff\xe7\x4d\x00\x69\x72\x7e\xfa\x4e\x13\x04\x20\x7e\x0a\x1f\x76\x1f\xb3\x8b\xd5\xfe\xbb\x63\xf4\x98\x08\x6e\x80\x72\x54\xbe\x17\xf5\x24\x14\xa8\x55\xc8\x24\xf1\xc8\x0d\x36\x5e\x9b\xc6\xd9\xf8\x53\x14\x8d\xe3\x6a\xc8\xeb\x61\x7e\xef\x13\xc7\x34\x85\xac\xd4\xc6\xe7\x93\x8f\xa7\x93\xe9\xf4\xf4\xe4\xf4\xec\xe3\x38\x79\x56\x63\x24\x6a\x8c\x76\xf4\x82\xda\x8c\xa6\x9d\x6a\x85\xcf\xba\x94\x37\xdb\x4c\xc7\xd3\xf1\x89\x47\xb1\x9c\xf7\xdc\x32\x36\x17\x8c\x92\x2c\xac\x6e\x97\xf7\xc2\xcc\x15\xea\xe6\x64\xd1\x20\x61\x74\x83\x1c\xb5\x9e\x2b\xb1\xe8\xf7\x86\x02\x69\x09\x94\x59\x85\x4f\x6b\x85\x7a\x2d\x98\x1b\x58\xce\xbc\x76\x59\x1b\x23\x3f\xa3\x09\xf0\x71\xcd\x1a\xcc\xda\xa9\x75\x9c\x4f\xbd\x81\x88\x2d\x37\xae\x8b\xc9\xc5\x24\x80\xa1\xc9\x1a\xf3\x71\xe2\xd7\xa7\xa7\xb9\x2f\x5c\xbd\xf9\x18\xbb\xe2\x4a\x81\x5d\x23\x83\x6d\x3d\x29\x9c\x7b\xc5\xc4\x12\x15\x15\x49\x8d\x36\xf5\xa3\x69\x4b\x08\x6a\xdd\x34\xcf\xd4\x8b\x68\x68\x8a\xc2\x9a\x9a\xe1\x59\x3f\xdc\x3d\x3e\x0a\x0f\x4e\xb5\xa6\x9e\x65\xbb\xfc\x05\x5c\x51\xa5\xca\xdc\xbf\xde\xf6\xc4\xfb\x45\x47\xbe\xad\x7c\x5f\x77\xfc\x44\xd5\x3c\x8b\xbb\x57\x3d\xb7\xa8\xff\x64\xf5\xdc\x7a\xb4\x43\xc1\x14\x8d\xa2\x64\x2f\xfb\x85\x14\xdc\xab\xba\xb8\x95\xcc\x2a\x82\xbe\x21\x28\xca\xf3\x3d\xa5\x26\x04\xcd\x66\xb0\x54\xa8\xac\x72\x4c\x3f\x4d\xee\xe8\xfe\x89\xe5\x46\x13\xd4\x43\xac\x89\xb4\x19\xdf\xc9\x24\x0d\xd9\xa1\x96\x1e\x14\xbe\x57\xc6\x64\x8b\x30\x35\xdb\x2b\xc1\x0d\xbe\x86\xea\x51\x0c\x8c\x89\x97\xb9\xa2\x1b\xca\x70\x85\x37\x9a\x00\x83\x62\x70\x5a\x02\xd3\xe8\x7d\x51\x02\x12\x16\x94\x51\x43\x83\x56\x76\xac\x93\x64\xa0\xd5\xc6\xf7\x37\x4f\xbf\xff\x72\x7b\x7f\xfd\xfb\xe3\xcd\xc3\x6f\xb7\x57\x37\xfe\x7e\xeb\xf1\x70\x46\x9d\x28\x21\x87\xb8\x03\x63\x01\x86\x07\xb8\x13\x92\xbf\x71\xb6\x7d\x10\xc2\xfc\x95\x32\x2c\xf6\xb1\x59\x64\x94\xc5\xbd\x7c\x60\x50\xa5\x94\x67\xf6\xbc\x43\xad\x5d\x67\x2a\x8b\x7f\x82\x9b\xe3\x06\x78\xc4\xc4\xca\x57\xd4\x3c\x1c\xaa\xbe\xe6\x74\xf2\xd1\x6c\x04\xb3\x29\xde\x09\xcb\x0f\xad\x87\xa9\xa3\xa9\x54\x6c\x34\xf8\x50\xb0\xd6\x65\x99\x2f\xe9\x6a\x94\x4b\x0e\x21\x97\xf6\x0c\x59\xd0\x9b\xe0\x83\xa3\x4f\x27\x38\x5c\x0d\xac\xcd\x73\x8d\x4b\xb0\xac\xdb\xf1\x63\xa9\xa8\xc8\xf2\x82\x81\xd6\xf7\x85\xfe\xb9\x6b\x47\x6e\x2b\x18\x11\x45\x0d\x25\xc0\xba\x84\xca\xad\x5d\xca\xd4\x02\x2e\xd9\x0b\x6c\x7b\x6b\xab\xce\x37\x12\x54\x25\xef\x24\xd7\x63\x54\x01\x7a\x14\xfd\x4c\xed\xce\x74\xba\x7d\xc5\x18\x5e\x9a\xdb\x88\xf7\x83\x5d\xb3\x19\x5c\x9f\x15\x10\x9c\x77\x7a\xfd\xc9\xa4\x4b\x20\x98\x1b\xb6\x03\xdb\x8e\x67\xb6\x2c\x06\xfb\xab\xc2\xa4\x97\x49\x22\xb8\xce\x62\xc0\x13\xb7\xcd\x49\xfe\xe6\x95\xba\x42\x7a\x90\xf7\xf3\xe8\xdb\x53\x33\x52\x5d\xc0\x05\x0a\x63\xe1\xb6\x3b\x91\x38\x13\x9e\x7e\xf4\xcf\x3d\xed\x0f\x3c\xdd\x5f\x70\x77\x2c\xed\x52\x8e\xcd\xa1\xad\xa5\x1a\x16\xc3\x03\x76\x14\x5a\x23\xbc\xfa\x76\xe7\xa8\x03\x87\xaf\x66\x96\x0f\xfb\xe6\xc8\x07\x29\x9f\xf6\xb9\x26\x54\x0b\x20\x63\xb0\x66\x2d\x14\xfd\x23\x8b\xb9\xf1\xf3\x45\xb6\x13\xfa\xee\x0e\x8b\xcb\xfc\x07\xc1\x76\x7e\x33\xfb\xaf\xdd\x10\x2e\x84\x30\xda\x28\x90\x92\xf2\x55\xf9\x8a\xa3\x22\xce\x76\xdc\xc5\xe6\x05\x6a\xd6\x71\x5b\x7d\x27\xa8\x2c\xeb\x44\x7e\xeb\xed\x40\xd2\xcf\x4a\x58\xd9\x0f\xd5\xb8\xe9\xc7\x56\xcc\xb4\x06\xa9\x0e\x11\xf2\x44\x0a\xea\xfa\x8b\xbf\xfe\x78\xbe\x65\x25\xbd\xb3\xea\xc6\x53\x87\x75\xd8\xa0\x5a\x78\xe4\x33\xaa\x7b\x85\xfd\x05\x0c\x59\xb7\x38\xf9\x2c\xfa\xa7\x9a\xc5\xf5\x8f\xc3\x5f\x66\x85\x26\xa0\xf6\x51\xf3\xe4\x4f\x4a\x97\x5f\x28\x4f\x5c\x3c\xbe\xff\x4b\xb3\x4f\x74\xfb\x12\xc7\x1a\x61\xcb\x9b\xd7\x9f\xfd\x61\xf9\x7f\x3e\x33\x05\xc3\x07\x5c\x76\xac\x59\x04\xde\xa0\xe7\x9a\x5f\x07\x86\xea\xdc\x01\xca\x68\xbb\xf8\x07\x12\x33\x50\x29\x76\x7c\xd4\x8c\x76\x2d\xe3\x3b\xbf\x63\xf4\xc2\xfa\xa8\x48\x93\x4a\xf4\x17\x97\xdc\x47\x6f\x47\xff\x0a\x00\x00\xff\xff\x48\xb8\xdf\xe8\x58\x21\x00\x00")

func coredns111JsonBytes() ([]byte, error) {
	return bindataRead(
		_coredns111Json,
		"coredns-1.11.json",
	)
}

func coredns111Json() (*asset, error) {
	bytes, err := coredns111JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "coredns-1.11.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _coredns112Json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x59\x5b\x53\x23\xbb\x11\x7e\xe7\x57\x4c\xcd\xf3\xda\xd8\xdc\x96\xf8\x8d\x03\x64\x0f\x55\x0b\x71\x01\x7b\x5e\xb2\x5b\xa7\xda\x9a\xb6\xad\xa0\x91\x94\x96\xc6\xe0\xb3\xc5\x7f\x4f\x69\xee\x17\xcd\xd8\x10\x4e\x92\xaa\xf8\x61\x77\x90\x5a\xdd\x3d\xdd\x5f\xdf\x34\x3f\x0f\x82\x20\x04\xcd\x7f\x43\x32\x5c\xc9\x70\x16\x84\x9b\x69\xf8\xc9\xad\x72\x8b\xb1\x09\x67\xc1\xdf\x0f\x82\x20\x08\x7e\xa6\xff\xf6\x11\xa7\x3b\x4f\x5c\x46\x6e\xed\x01\x69\xc3\x19\x56\x1b\x31\x5a\x88\xc0\x42\x38\x2b\xd9\x38\x46\x52\x2a\x0b\x96\x2b\x69\x1a\x1b\x41\x10\x6a\x52\x31\xda\x35\x26\x66\xcc\xd5\xa1\x56\x64\x1d\xdf\xbf\x4c\x4f\x8f\x4b\xa6\x1e\x3a\xc3\x08\x34\x3a\x4a\x4b\x09\x86\x25\xe1\x6b\x75\x26\x14\xb0\x40\xd1\x91\x87\x4f\x66\x0c\x31\xfc\xa1\x24\x3c\x9b\x31\x53\xf1\x21\x53\xb1\x56\x12\x65\x2a\xf8\x29\x59\xe0\x28\x92\xa6\x29\xfc\xe9\xdc\x8c\x40\xeb\x01\x82\x64\x81\x24\xd1\x62\xaa\x1d\x13\x89\xb1\x48\x23\x93\x5b\xa7\x50\x73\xe0\x88\x84\x38\xa5\xbb\x54\x84\x57\x77\x0f\xfe\x37\x2a\x88\x3c\x3a\xa4\x7b\x46\x03\xab\x08\xcc\xd6\x58\x8c\x0b\x4e\x25\x9f\xd0\x68\x64\x4d\xf7\x38\xab\x57\xee\xcf\x7e\x75\xa3\xd5\x44\xb7\xdf\x3c\x3f\x1d\xce\x82\xd3\xe3\xf6\x3a\x29\xab\x98\x12\xee\xd8\xb7\xab\x79\xfb\x98\x05\x5a\xa1\x9d\x17\x87\x6b\x9b\xaf\x9f\xf6\xd2\x63\x64\x99\x7e\x8f\x2e\x8f\x97\x6f\xd1\xa5\x7c\xfe\x51\x33\xb6\x41\x81\xcc\x2a\x6a\x83\xcb\x87\x13\xaf\x27\x0d\x1a\x17\x56\x17\xcb\x25\x97\xdc\x6e\xdd\x81\x3b\x25\xeb\x08\x09\xed\x36\x03\xf8\x65\x06\xa6\x9b\x79\xe9\xc9\x83\x1a\xb7\xb7\xc7\xea\x05\x63\x2a\x91\x76\x57\xc8\xbe\x2b\x7a\x98\x22\xdc\x2b\x78\x06\xe1\xdd\xe5\xb2\x1b\xdd\x6f\xb4\x49\xf7\x7d\x5d\xe0\x2d\xb9\x48\x05\x8c\x67\xa7\xc7\xc1\xcf\xef\xd2\x6d\x20\x91\x22\x93\x3d\xaf\x11\x84\x5d\x67\xcf\x55\xf8\x06\x79\xb8\x8f\x85\x62\x20\x02\x2e\x47\x10\x45\x34\x06\xd2\x10\x70\x7d\x96\x3d\xe4\xdc\x82\x40\xab\xc8\x04\x5c\x1a\x64\x09\x61\xb1\x98\x68\x63\x09\x21\x2e\xfe\x5e\x82\x10\x76\x4d\x2a\x59\xad\xfd\xfc\x32\xc2\xd7\xec\xbf\x2a\x35\x06\x33\x97\x38\xcb\xd5\x97\x6d\x30\x0e\x0e\xd1\xb2\x43\x42\xa3\xc4\x66\xcc\x94\x5c\x66\xbb\x0c\xd8\x1a\x83\xe3\x49\xf6\x97\x50\x4a\x67\x4f\x84\x42\x41\x54\xac\x42\xb4\x00\x01\x92\xe1\x77\xf9\xfa\x5d\x76\x53\x49\x01\xad\x4b\x25\x97\x7c\x75\x0b\xfa\xff\x0f\x55\xa0\xb5\x39\xf4\x85\xdb\x15\x6a\xa1\xb6\x31\xee\x0e\xb5\x56\x75\xfc\xb7\x6b\xd8\xde\xf6\xfa\xa8\x7a\xf4\x1e\xd3\x0e\x95\x23\x52\x2b\x42\x63\xae\x10\x22\xc1\x25\x3e\x20\x53\x32\x72\x56\x38\x9a\x9e\x7c\x3e\x39\x3f\x3e\x3b\xf9\x5c\x13\x45\xa8\x05\x67\x90\xee\x37\x96\x37\xdc\x39\xe9\x57\x6e\xac\xa2\xed\x57\x1e\x73\x67\x9d\xe9\x64\x8f\x3c\x1e\x83\x65\xeb\xaf\x3e\xdb\xbf\xcf\xfa\xbb\xf0\x5a\xaf\x33\x8d\x32\x61\x09\x2c\xae\xb6\x6d\xfd\x48\x09\xc1\xe5\xea\x9b\x8e\xc0\x62\x57\xc3\x18\x5e\x1e\x12\x5a\xb9\x9d\xe9\xa7\xce\xd6\x37\x09\x1b\xe0\x02\x16\x69\xb2\x9b\xf6\x15\xde\xb2\x00\xdd\x37\x84\x79\xf5\xb4\x18\x6b\xd1\x55\xc5\x8f\xf7\x74\x87\x11\xa6\x88\x7f\xe4\x31\x1a\x0b\xb1\x33\x8c\x4c\x84\x68\xa9\xeb\xc5\xff\xfb\xbd\xb0\xdb\x0f\x75\x4f\xb4\x0d\xd2\x41\x6a\xba\x0a\x55\x0d\xef\x28\x29\x55\x84\x17\xfd\xfb\x29\x4a\xff\x99\x70\xc2\xe8\x2a\x21\x2e\x57\x0f\x6c\x8d\x51\xe2\xac\x7d\xb3\x92\xaa\x5c\xbe\x7e\x41\x96\xd8\x2c\xe1\x74\x79\xe4\x72\x1e\x72\x2c\x3f\x22\xc5\xed\x66\xae\xfa\xf9\x8e\x07\x25\xe4\xaf\x5f\x34\x65\x9d\x49\x3f\x87\x21\x2e\x29\xa7\x27\x4c\xfb\x99\x05\x5a\x18\x37\x93\x8a\xea\xfa\xa3\x71\x52\x69\x24\xc8\xc2\x31\xbc\x91\xc3\xb4\x1b\x10\x09\x0e\x6b\x99\xd2\x09\x2e\x93\x97\x70\x80\xe6\x47\xef\xde\x6b\xbf\x02\xef\x34\x00\x10\x5b\xff\xe7\x4d\x00\x71\x74\x76\xf2\x4e\x13\xf4\xec\xf8\x4f\xf8\xa8\xbb\x94\x6d\xaa\xe6\xdf\x2d\xa3\x87\x4c\x49\x0b\x5c\x22\xf9\x5e\xd4\x13\x50\x40\xab\x3e\x93\x84\x23\xd7\xff\x78\x6d\x1a\xa6\x5d\x52\x9e\x34\x0e\xcb\x5e\xb0\x43\xf9\xa3\x7b\x38\xe4\x31\xa4\xa9\x36\x3c\x9b\x1c\x9d\x4c\xa6\xd3\x93\xe3\x93\xd3\xa3\x71\xf4\x44\x63\x64\x34\xc6\x64\xf4\x8c\xc6\x8e\xa6\xad\x6c\x85\x4f\xa6\x90\x37\xdb\x4c\xc7\x47\xe3\x23\x8f\x62\x19\xef\x79\x22\xc4\x5c\x09\xce\x52\x58\xdd\x2c\xef\x94\x9d\x13\x9a\x7a\x67\x51\x3b\x22\xf8\x06\x25\x1a\x33\x27\xb5\xe8\xd6\x86\x9c\x68\x09\x5c\x24\x84\x8f\x6b\x42\xb3\x56\xc2\x35\x2c\xa7\x5e\xbb\xac\xad\xd5\x5f\xd0\xf6\xf0\x71\xc5\x1a\xec\xda\xa9\x75\x98\x35\xc7\x3d\x88\x2d\x06\xb3\xf3\xc9\xf9\xa4\x87\xc2\xb0\x35\x66\xed\xc4\xaf\x8f\x8f\x73\x1f\x5c\xbd\xf1\x18\xba\xe4\xca\x41\x5c\xa1\x80\x6d\xd5\x29\x9c\x79\xc5\x84\x1a\x89\xab\xa8\x22\x9b\xfa\xc9\x4c\xc2\x18\x1a\x53\x37\xcf\xd4\x4b\x68\x79\x8c\x2a\xb1\x15\xc3\xd3\x2e\xdc\x3d\x3e\xea\x6f\x9c\x2a\x4d\x3d\x33\x79\xf1\xeb\x71\x45\x19\x2a\x73\xff\x14\xdc\x11\xef\x17\x1d\xf8\x86\xf7\x7d\xdd\xf1\x81\xaa\x79\xe6\x7b\xaf\x7a\x6e\x9e\xff\x60\xf5\xdc\x14\xb5\x43\xc1\x18\x2d\x71\xb6\x97\xfd\xfa\x14\xdc\x2b\xbb\xb8\xc9\x2d\x21\x86\xbe\x26\x28\xc8\xe2\x3d\xe6\xb6\x6f\x37\xed\xc1\x62\x45\x69\xe6\x98\x7e\x9e\xdc\xf2\xfd\x03\xcb\xb5\x26\x68\x86\x58\x33\x9d\xa4\x7c\x27\x93\xb8\xcf\x0e\x95\xf4\x5e\xe1\x7b\x45\x4c\x3a\x2f\x73\xbb\xbd\x54\xd2\xe2\x4b\x5f\x3e\x0a\x41\x08\xf5\x3c\x27\xbe\xe1\x02\x57\x78\x6d\x18\x08\xc8\x1b\xa7\x25\x08\x83\xde\x17\x65\xa0\x61\xc1\x05\xb7\xbc\xd7\xca\x8e\x75\x14\x0d\x94\xda\xf0\xee\xfa\xf1\xf7\x5f\x6e\xee\xae\x7e\x7f\xb8\xbe\xff\xed\xe6\xf2\xda\x5f\x6f\x3d\x1e\x4e\x4f\x47\xa4\xf4\x10\x77\x10\xa2\x87\xe1\xde\xee\xd4\xa4\xd8\x6d\x7a\xf3\x93\x0e\xa7\x4b\x48\x84\xaf\x7e\xa4\x8e\x87\xe8\x6f\x52\x6c\xef\x95\xb2\x7f\xe5\x02\xf3\xc9\x6d\x16\x58\x4a\x70\x2f\x6f\x59\xa4\x98\xcb\xd4\xf2\xb7\x68\x8c\xab\x61\x45\x99\x88\x70\x73\x58\xdb\x1e\x09\xb5\xf2\xa5\x3f\x0f\x87\xb2\x02\x3a\x9d\x7c\x67\x36\x4a\x24\x31\xa6\xef\xf8\xc6\xcc\x19\xbb\x33\xa5\x8a\xb5\x56\xa0\x0f\xd6\x55\x02\x97\x4b\xbe\x1a\x65\x92\xfb\x88\x0b\x7b\xf6\x59\xd0\x9b\x0a\x06\x9b\xa4\x16\x8c\x5c\xb6\xac\xcc\xe3\xf7\x6d\xa8\x89\xab\x34\x82\x04\x18\x73\x97\xeb\x9f\xb9\x76\x54\x5c\x58\x33\xe2\x96\x33\x10\xed\xb3\xe4\x66\x34\xb2\x95\x8c\x0b\xf1\x0c\xdb\xce\x8c\x6b\xb2\xf1\x05\xa9\x60\x1f\x65\xaa\x8c\xca\x8d\xce\x89\x6e\x58\xb7\x1b\x40\xd3\xbc\xb6\xec\x9f\xb0\x9b\x84\x77\x83\x25\xb6\x8e\xaf\x2f\x04\x0c\xe7\xad\xc6\xe0\x78\xd2\x3e\xa0\x84\xeb\xcc\x7b\x46\x23\x4f\x23\x8a\xcb\x25\x32\x9b\xdd\xec\xe6\x73\x9d\x17\xb5\xf9\xb4\xe0\x46\xb8\x11\x29\x81\xad\x91\x21\x06\xe7\x99\x76\xf0\x77\x82\xce\xa3\x40\xce\xf8\x32\xf7\xe9\x45\x14\x29\x69\x52\x1c\x7a\xb4\xa8\xcf\x1d\xd7\x2f\xdc\xa5\xfd\x37\x21\x30\x8b\x80\x3d\x4d\xc3\xca\xeb\xc2\x9e\x34\x9e\xe3\xe6\x56\x45\xce\x87\x27\x47\xfe\x2e\xad\xf9\xd5\xaa\xfd\xeb\x9d\x74\x0b\xbb\x14\x4d\x7e\xdf\x8c\x55\xb6\xb6\xfd\xe3\x40\xd0\x37\xf4\x78\xf5\x6d\x77\x7d\x6f\x6c\x15\xeb\x99\x66\xd8\x37\x07\xbe\x9d\xe2\x69\x9f\x4b\x4d\x5a\x00\x1b\x43\x62\xd7\x8a\xf8\x1f\x29\xe8\xc7\x4f\xe7\x29\x1c\x7d\x37\x9d\xf9\x17\x8a\x7b\x25\x76\x7e\x08\xfc\xaf\xdd\x67\x2e\x94\xb2\xc6\x12\x68\xcd\xe5\xaa\x78\xc5\x51\x8e\xb3\x1d\x37\xc7\x59\x92\x9c\xb5\xdc\x56\xdd\x60\x52\x22\x5a\xc8\x6f\xbc\x1d\x68\xfe\x85\x54\xa2\xbb\x50\x0d\xeb\x7e\x6c\x60\xa6\xd1\xf6\xb5\x0e\xa1\x8c\xb4\xe2\xae\xc6\xf9\x13\xa0\xe7\x03\x5d\xd4\x59\x2b\xef\x67\x4d\xbf\x0e\x1b\xa4\x85\x47\xbe\xe0\xa6\x53\x5c\x9e\xc1\xb2\x75\x83\x93\xcf\xa2\x7f\xaa\x59\x5c\x02\x7d\xfb\xcb\xac\xd0\xf6\xa8\x7d\x50\x5f\xf9\x93\xc2\xe5\x17\x2e\x23\x87\xc7\xf7\x7f\x3e\xf7\x89\x6e\x5e\x39\x25\x56\x25\xc5\x3d\xf1\x47\x7f\x2d\xff\x9f\x8f\x4c\x25\xf0\x1e\x97\x2d\x6b\xe6\xc0\x1b\xf4\x5c\xfd\x5b\xc6\x50\x9e\x7b\x83\x32\x26\x59\xfc\x03\x99\x1d\xc8\x14\x3b\xbe\xd4\x06\xbb\xae\x0e\x76\x7e\x75\xe9\xc0\xfa\x20\x0f\x93\x52\xf4\x57\x17\xdc\x07\xaf\x07\xff\x0a\x00\x00\xff\xff\x0b\xba\x03\x53\x2d\x22\x00\x00")

func coredns112JsonBytes() ([]byte, error) {
	return bindataRead(
		_coredns112Json,
		"coredns-1.12.json",
	)
}

func coredns112Json() (*asset, error) {
	bytes, err := coredns112JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "coredns-1.12.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _coredns113Json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\x4b\x73\xdb\x38\x12\xbe\xfb\x57\xb0\x78\x8e\x64\xc9\xaf\x78\x74\xf3\xd8\xde\x8c\x6b\x63\xaf\xca\x76\xe6\xb2\x99\x9a\x82\xc0\x96\x84\x35\x88\xc6\x02\xa0\x6c\x4d\xca\xff\x7d\x0b\xe0\x9b\x04\x68\xc9\x9b\x6c\xe5\xb0\x3c\x24\x34\xd0\xe8\x6e\x76\x7f\xfd\x00\xa0\x6f\x07\x51\x14\x13\xc9\x7e\x07\xa5\x19\x8a\x78\x16\xc5\x9b\x69\xfc\xc1\x8e\x32\x03\xa9\x8e\x67\xd1\x3f\x0f\xa2\x28\x8a\xbe\xb9\x7f\x43\xc4\x6e\xe6\x89\x89\xc4\x8e\x3d\x80\xda\x30\x0a\xf5\x44\x0a\x86\x24\xc4\x90\x78\x56\xb1\xb1\x8c\x84\x40\x43\x0c\x43\xa1\x5b\x13\x51\x14\x4b\x85\x29\x98\x35\x64\x7a\xcc\xf0\x50\xa2\x32\x96\xef\x2f\xd3\xd3\xe3\x8a\xa9\x87\x4e\x53\x45\x24\x58\x4a\xa3\x32\x88\x2b\xc2\xd7\x7a\x4d\xcc\xc9\x02\x78\x4f\x1e\x3c\xe9\x31\x49\xc9\x5f\x28\xc8\xb3\x1e\x53\x4c\x0f\x29\xa6\x12\x05\x08\x27\xf8\x29\x5b\xc0\x28\x11\xba\x2d\xfc\xe9\x5c\x8f\x88\x94\x03\x04\xd9\x02\x94\x00\x03\x4e\x3b\xca\x33\x6d\x40\x8d\x74\x61\x9d\x52\xcd\x81\x25\x82\xa4\x8e\xee\x12\x15\x5c\xdd\x3d\xf8\xbf\xa8\x24\xf2\xe8\xe0\xe6\xb4\x24\xb4\x26\xd0\x5b\x6d\x20\x2d\x39\x55\x7c\x62\x2d\x81\xb6\xdd\x63\xad\x5e\xbb\x3f\x7f\x9a\x46\x6b\x88\xee\x7e\x79\xb1\x3a\x9e\x45\xa7\xc7\xdd\x71\x85\x06\x29\x72\xbb\xec\xcb\xd5\xbc\xbb\xcc\x10\xb5\x02\x33\x2f\x17\x37\x26\x5f\x3f\xec\xa4\xc7\xc8\x50\xf9\x1e\x5d\x1e\x2f\xf7\xd1\xa5\x7a\xff\xa3\x61\x6c\x0d\x1c\xa8\x41\xd5\x05\x97\x0f\x27\x5e\x4f\x6a\xd0\x36\xac\x2e\x96\x4b\x26\x98\xd9\xda\x05\x77\x28\x9a\x08\x89\xcd\x36\x07\xf8\x65\x0e\xa6\x9b\x79\xe5\xc9\x83\x06\xb7\xfd\x63\xf5\x82\x52\xcc\x84\xd9\x33\x64\xff\xeb\xc0\xa2\xa8\x60\xa7\xb8\x1a\x44\x7e\x9f\xcb\xdb\xc0\xdf\xd3\x5c\x7d\x53\xd8\x98\x5c\x32\xee\x04\x8c\x67\xa7\xc7\xd1\xb7\xaf\xc2\x4e\x80\x52\xa8\x74\xfe\xbe\x06\xc2\xcd\x3a\x7f\xaf\x23\x3b\x2a\x32\xc1\x98\x23\x25\x3c\x62\x62\x44\x92\x44\x8d\x89\x92\x24\x62\xf2\x2c\x7f\x29\xb8\x45\x91\xc4\x44\x47\x4c\x68\xa0\x99\x82\x72\x30\x93\xda\x28\x20\x69\xf9\xf7\x92\x70\x6e\xd6\x0a\xb3\xd5\xda\xcf\x2f\x27\x7c\xcd\xff\xab\xb3\x66\x34\xb3\x39\xb5\x1a\x7d\xd9\x46\xe3\xe8\x10\x0c\x3d\x54\xa0\x91\x6f\xc6\x14\xc5\x32\x9f\xa5\x84\xae\x21\x3a\x9e\xe4\x7f\x71\x44\x99\xbf\x29\xe0\x48\x92\x72\x94\x24\x0b\xc2\x89\xa0\xf0\x55\xbc\x7e\x15\xfd\x2c\x53\xa2\xee\x12\xc5\x92\xad\x6e\x89\xfc\x3f\xe0\x1a\x80\x23\x52\xea\x43\x5f\x90\x5e\x81\xe4\xb8\x4d\xe1\x27\x0e\xd0\xef\x55\xc5\xde\x63\xda\xa1\x22\xa6\x70\xa5\x40\xeb\x2b\x20\x09\x67\x02\x1e\x80\xa2\x48\xac\x15\x8e\xa6\x27\x1f\x4f\xce\x8f\xcf\x4e\x3e\x36\x44\x29\x90\x9c\x51\xe2\xe6\x5b\xc3\x1b\x66\x9d\xf4\x1b\xd3\x06\xd5\xf6\x33\x4b\x99\x09\xb1\x08\x55\x81\x94\x18\xba\xfe\xec\xf3\xc1\xfb\xbc\xf0\x16\x6e\x9b\x55\xaa\x55\x64\x8c\x22\x06\x56\xdb\xae\x7e\x0a\x39\x67\x62\xf5\x45\x26\xc4\x40\x5f\xc3\x94\xbc\x3c\x64\x6a\x65\x67\xa6\x1f\x7a\x53\x5f\x04\xd9\x10\xc6\xc9\xc2\xe5\xc3\x69\xa8\x6c\x57\xe5\xeb\xbe\x25\xcc\xab\xa7\x81\x54\xf2\xbe\x2a\x7e\xdc\xbb\x19\xaa\xc0\x21\xff\x91\xa5\xa0\x0d\x49\xad\x61\x44\xc6\x79\x47\x5d\x6f\x1c\xbc\xdf\x0b\x6f\xfb\xa1\xe9\x89\xae\x41\x7a\x88\x75\xa3\xa4\xee\x00\x7a\x4a\x0a\x4c\xe0\x22\x3c\xef\xd0\xfa\xef\x8c\x29\x48\xae\x32\xc5\xc4\xea\x81\xae\x21\xc9\xac\xb5\x6f\x56\x02\xab\xe1\xeb\x17\xa0\x99\xc9\x13\x4f\x9f\x47\x21\xe7\xa1\xc0\xf2\x23\xa8\xb4\xdb\x0a\xd6\x8f\x6f\x79\x54\x41\xfe\xfa\x45\xaa\xbc\xaf\x09\x73\x18\xe2\xe2\x38\x3d\x81\xeb\x86\x16\x60\xc8\xb8\x9d\x5c\xb0\xef\x8f\xd6\x4a\x94\xa0\x48\x1e\x8e\xf1\x8d\x18\xa6\xdd\x10\x9e\xc1\xb0\x96\x8e\x8e\x33\x91\xbd\xc4\x03\x34\x7f\x04\xe7\x5e\xc3\x0a\xbc\xd3\x00\x44\xd1\xf5\xff\xde\x04\x24\x4d\xce\x4e\xde\x69\x82\xc0\x8c\x7f\x85\x8f\xba\x4f\xd9\xa5\xea\x99\x39\x96\x98\x5c\x08\xc3\x86\x23\x47\x2a\x58\x82\xda\x27\x74\x7c\x66\xf2\x3b\xd2\x69\x50\x48\xb7\xf1\x14\x88\xbb\xa8\x4a\x51\x0f\xfe\x42\xd2\x21\xdd\x2b\xc6\x86\x41\x56\xc3\xac\x4c\x68\x43\x50\xd9\x0f\x58\xbb\x43\x2b\x0a\x25\xd1\xee\x13\x46\x58\x18\x63\xe1\x55\xc1\xc0\x8c\x0d\x4a\xe4\xb8\xda\xfe\xbd\xb0\x4d\x2b\xfa\xd6\xa8\x8d\x6b\x5d\xfc\xe0\xf5\x33\x8d\x9f\x81\xad\xd6\xb6\xa6\x4c\x27\x13\x0f\x45\x5f\xfb\xae\xd6\x6d\x8a\x8e\x98\x98\xa2\x30\x84\x09\x50\x3e\x6b\x7b\x60\x4f\xd4\x2a\xe4\x97\x78\x64\xb7\x00\x5e\xd7\xc6\x6e\xa3\x50\x14\xc5\xc3\x6a\x3b\xd4\xd7\xbd\xbf\x38\x66\x29\x71\xad\x44\x7c\x36\x39\x3a\x99\x4c\xa7\x27\xc7\x27\xa7\x47\xe3\xe4\x49\x8d\x81\xaa\x31\x64\xa3\x67\xd0\x66\x34\xed\x54\x63\x78\xd2\xa5\xbc\xd9\x66\x3a\x3e\x1a\x9f\x79\x14\xcb\x79\xcf\x33\xce\xe7\xc8\x19\x75\x3e\xbb\x59\xde\xa1\x99\x2b\xd0\xcd\x0e\xba\xb1\x84\xb3\x0d\x08\xd0\x7a\xae\x70\xd1\xef\x7d\x0a\xa2\x25\x61\x3c\x53\xf0\xb8\x56\xa0\xd7\xc8\x6d\x63\x7e\xea\xb5\xcb\xda\x18\xf9\x09\x4c\x30\x6e\x63\x49\xcc\xda\xaa\x75\x98\xef\x0f\x03\x81\x53\x1e\x5b\x9c\x4f\xce\x27\x01\x0a\x4d\xd7\x90\xb7\xcd\xbf\x3d\x3e\xce\x7d\x18\xf4\x22\x30\xb6\x49\x88\x11\x7e\x05\x9c\x6c\xeb\x8e\xf8\xcc\x2b\x26\x96\xa0\x18\x26\x35\xd9\xd4\x4f\xa6\x33\x4a\x41\xeb\xa6\x79\xa6\x5e\x42\xc3\x52\xc0\xcc\xd4\x0c\x4f\xfb\xe9\xdc\xe3\xa3\xf0\x06\xa1\xd6\xd4\x73\x62\x55\x3e\x01\x57\x54\xa1\x32\xf7\x9f\x11\xf5\xc4\xfb\x45\x47\xbe\xa3\xad\x5d\xdd\xf1\x1d\x55\xf3\x9c\x7e\x79\xd5\x7b\xbc\xfc\xee\xea\xfd\x32\x7d\x53\xc1\x14\x8c\x62\x74\x27\xfb\x85\x14\xdc\x29\xbb\x28\xd0\x98\x29\x0a\xbe\x26\x3f\xca\xe3\x3d\x65\x26\x34\xeb\xf6\x18\x29\x2a\x97\x39\xa6\x1f\x27\xb7\x6c\xf7\xc0\xb2\xad\x37\xe8\x21\xd6\x54\x66\x8e\xef\x64\x92\x86\xec\x50\x4b\x0f\x0a\xdf\x29\x62\xdc\x91\x11\x33\xdb\x4b\x14\x06\x5e\x42\xf9\x28\x26\x9c\xe3\xf3\x5c\xb1\x0d\xe3\xb0\x82\x6b\x4d\x09\x27\x45\x77\xb3\x24\x5c\x83\xf7\x43\x29\x91\x64\xc1\x38\x33\x2c\x68\x65\xcb\x3a\x49\x06\xea\x7d\x7c\x77\xfd\xf8\xe7\xaf\x37\x77\x57\x7f\x3e\x5c\xdf\xff\x7e\x73\x79\xed\x2f\xa2\x1e\x0f\xbb\xd5\x89\x42\x39\xc4\x9d\x70\x1e\x60\xb8\xb3\x3b\xa5\x42\x7a\xeb\xce\x45\xdd\x21\xcc\x92\x64\xdc\x57\x3f\x9c\xe3\x49\xf2\x0f\xc1\xb7\xf7\x88\xe6\x6f\x8c\x43\x71\x42\x31\x8b\x8c\xca\x60\x27\x6f\x19\x50\x29\x13\xce\xf2\xb7\xa0\xb5\xad\x61\x65\x99\x48\x60\x73\xd8\x98\x1e\x71\x5c\xf9\xd2\x9f\x87\x43\x55\x01\xad\x4e\xbe\x35\x1b\xe4\x59\x0a\xee\x1b\xf7\xcc\x9c\xa9\x5d\x53\xa9\xd8\x68\x05\x42\xb0\xae\x13\xb8\x58\xb2\xd5\x28\x97\x1c\x22\x2e\xed\x19\xb2\xe0\xfe\x4d\x52\x07\x46\x36\x5b\xd6\xe6\xf1\xfb\x36\x96\x8a\xa1\x8b\x20\x4e\xb4\xbe\x2b\xf4\xcf\x5d\x3b\xb2\xfb\xe3\x11\x55\xcc\x30\x4a\x78\x77\xa1\x02\x6d\x88\x32\xb5\x80\x0b\xfe\x4c\xb6\xbd\x03\x1c\x9d\x6f\x30\x40\x95\xbc\x93\x5c\x8f\x51\x35\xd1\x5b\xd1\x8f\xe9\x6e\xf7\xa7\xdb\x27\xfa\xe1\xe3\xa3\x36\xe1\xdd\x60\x7d\x6d\x82\xeb\x93\x22\x14\xe6\x9d\xae\xe0\x78\xd2\x5d\x80\xdc\xee\x0e\x02\x7b\x12\x4f\x17\x0a\xcb\x25\x50\x93\x5f\x7a\x14\x3b\x2f\x2f\x64\x8b\x3d\x8a\xb3\xbf\x42\x0e\x9d\xfd\x70\x4a\xb4\x01\xd5\x8d\xfc\x5e\xc4\x79\x14\x28\x18\x5f\x16\x3e\xbd\x48\x12\x14\xda\x81\xd0\xa3\x45\x73\xef\x73\xfd\xc2\x6c\xce\xdf\x0b\x7e\x39\xfc\x77\x34\x0d\xad\x8e\xcb\x03\x39\xbc\xc0\xcd\x2d\x26\xd6\x87\x27\x47\xfe\x16\xad\x7d\xa1\xdb\x7d\x82\xbb\xd1\xd2\x2e\x65\x87\x1f\xdc\x29\x95\x7d\x6d\x78\x2f\x10\x85\x76\xf4\x5e\x7d\xbb\x2d\xdf\x9e\x7d\x62\x33\xcd\x0c\xfb\xe6\xc0\x37\x53\xbe\xed\x72\x72\xaf\x16\x84\x8e\x49\x66\xd6\xa8\xd8\x5f\x0e\xf4\xe3\xa7\x73\x07\x47\xdf\x71\x7e\x71\x79\x77\x8f\x7c\xdf\x3b\xf2\x9f\xe2\x3c\x7f\x81\x68\xb4\x51\x44\x4a\x26\x56\xe5\xd7\x8f\x0a\x08\xbe\x71\x73\x92\x27\xcf\x59\xc7\xa3\xf5\x09\xbe\xca\x78\x27\x28\x5a\x5f\x47\x24\xfb\xa4\x30\x93\x7d\x14\xc7\x4d\x17\xb7\xe0\xd4\x6a\x07\x3b\x8b\x40\x24\x12\x99\xad\x7d\xfe\xdc\xe8\xb9\xd6\x4e\x7a\x63\xd5\xfd\x84\x0e\xeb\xb0\x01\xb5\xf0\xc8\xe7\x4c\xf7\x8a\xce\x33\x31\x74\xdd\xe2\xe4\xb3\xe8\x0f\x35\x8b\xcd\xad\xfb\x7f\xcc\x0a\x4c\x40\xed\x83\xe6\xc8\x0f\x8a\xa4\x5f\x99\x48\x2c\x1e\xf7\x0b\xa8\x96\x41\x3c\xa2\xdb\x47\xad\x99\xc1\xac\xbc\x1f\xf9\xde\xbf\x31\xf9\xe9\x23\x13\x39\xdc\xc3\xb2\x63\xcd\x02\x78\x83\x9e\x6b\xde\xe5\x0d\xa5\xc0\x3d\x94\xd1\xd9\xe2\x5f\x40\xcd\x40\xa6\x78\xe3\xf7\x0d\xd1\x5b\x47\x0a\x6f\xde\x3a\xf6\x60\x7d\x50\x84\x49\x25\xfa\xb3\x0d\xee\x83\xd7\x83\xff\x04\x00\x00\xff\xff\xce\xda\xa9\x5f\x63\x25\x00\x00")

func coredns113JsonBytes() ([]byte, error) {
	return bindataRead(
		_coredns113Json,
		"coredns-1.13.json",
	)
}

func coredns113Json() (*asset, error) {
	bytes, err := coredns113JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "coredns-1.13.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"aws-node-1.10.yaml": awsNode110Yaml,
	"aws-node.yaml": awsNodeYaml,
	"coredns-1.11.json": coredns111Json,
	"coredns-1.12.json": coredns112Json,
	"coredns-1.13.json": coredns113Json,
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
	"aws-node-1.10.yaml": &bintree{awsNode110Yaml, map[string]*bintree{}},
	"aws-node.yaml": &bintree{awsNodeYaml, map[string]*bintree{}},
	"coredns-1.11.json": &bintree{coredns111Json, map[string]*bintree{}},
	"coredns-1.12.json": &bintree{coredns112Json, map[string]*bintree{}},
	"coredns-1.13.json": &bintree{coredns113Json, map[string]*bintree{}},
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

