// Code generated by go-bindata. DO NOT EDIT.
// sources:
// assets/aws-node.yaml (4.472kB)
// assets/coredns-1.14.json (9.919kB)
// assets/coredns-1.15.json (10.18kB)
// assets/coredns-1.16.json (10.18kB)

package defaultaddons

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

var _awsNodeYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x58\x5d\x6f\xdb\x36\x17\xbe\xf7\xaf\x20\x04\xbc\x37\x2f\x26\xd9\x6e\x93\x2c\x13\xb0\x8b\x34\x71\xbb\x60\x89\x6b\xc4\x4d\xb6\x21\x28\x8c\x63\xea\x44\xe6\x4c\x91\x2c\x3f\x14\xbb\xbf\x7e\xa0\xe4\x0f\x49\x96\x95\xa6\xc0\x80\x5e\x4c\x57\xd6\xe1\xe1\xc3\xc3\xe7\x39\x3c\xe2\x71\x18\x86\x3d\x50\xec\x01\xb5\x61\x52\xc4\x44\xcf\x81\x46\xe0\xec\x42\x6a\xf6\x15\x2c\x93\x22\x5a\x9e\x9b\x88\xc9\x7e\x3e\xec\x2d\x99\x48\x62\x72\xc9\x9d\xb1\xa8\xef\x24\xc7\x5e\x86\x16\x12\xb0\x10\xf7\x08\x11\x90\x61\x4c\xe0\xd9\x84\x42\x26\xd8\xd3\x8e\xa3\xf1\xf6\x90\x80\x62\x1f\xb4\x74\xaa\x78\xf5\x4f\x48\xa8\x4e\x0a\x60\xc8\xe0\xab\x14\xf0\x6c\x22\x2a\xb3\x62\x54\xa3\x91\x4e\x53\xac\x38\x07\xff\x0f\x8a\xdf\x39\xea\xf9\x81\xb9\x0a\x4f\x1e\x83\xe0\xf3\x31\x14\x25\x13\xb3\x7b\xf1\x21\x56\xde\x20\x43\xa3\x80\x6e\x4c\xe5\x3a\xe4\x31\xe0\xcc\xd8\xe0\x27\x12\x3c\x83\xa5\x0b\xff\x23\x45\x5b\x2c\xd0\x58\x14\x57\x16\x85\x27\xd0\x1c\x5f\x3e\x01\xcc\xa4\x30\x68\x3b\xd7\xf8\xdc\xeb\x35\x25\xd9\x11\x3f\x45\x9d\x33\x8a\x17\x94\x4a\x27\x6c\x17\xf7\x64\xbf\xa3\x98\x2c\xdd\x1c\x43\xb3\x36\x16\xb3\x43\xf0\xd7\xe9\xfd\x8e\x89\x84\x89\xb4\x53\x76\xc9\xf1\x0e\x9f\xfc\xc8\x96\xa2\x8e\x55\x7a\x84\x1c\x26\xd5\x01\xa6\x71\xf3\xbf\x91\xda\x4d\x36\xb5\x92\xe1\x29\x3d\x20\xa1\x9b\x86\x12\xe7\xaa\x90\x65\x8a\xb6\x46\x0b\x28\x65\x3c\x03\xdf\xc3\x31\x21\x1c\xe6\xc8\x37\xd2\x2f\xcf\x4d\x08\x4a\x55\x77\xa3\x90\xfa\x31\xa7\x12\xb0\x38\xb5\x1a\x2c\xa6\xeb\xd2\xdb\xae\x15\xc6\xe4\x4e\x72\xce\x44\x7a\x5f\x38\x94\xf9\x54\xb5\x6c\x73\x2a\x83\xd5\xbd\x80\x1c\x18\x87\x39\xc7\x98\x04\xc3\xc1\xff\xfc\x89\x30\xc8\x91\x5a\xa9\x4b\xbf\xcc\xe7\xd5\x4d\x25\xa2\xb6\x98\x08\xb1\x98\x29\xbe\x03\xaf\xee\xdb\x3f\xbc\x36\xbf\x1d\x81\x90\xed\xce\xfc\xa3\x34\x93\x9a\xd9\xf5\x25\x07\x63\xc6\x05\x75\x25\x3f\x85\x7b\x48\x35\xb3\x8c\x02\xdf\x78\xc3\xd3\x13\x13\xcc\xae\xf7\x2b\x78\xaf\x8b\x03\xab\x3f\x59\x5f\x1c\xd3\x98\x5c\x39\xcd\x44\x3a\xa5\x0b\x4c\x9c\xa7\xe6\x3a\x15\x72\x67\x1e\xad\x90\x3a\x9f\x66\xd5\x99\x25\xe6\x74\x43\xce\x27\xd4\x99\xa9\x0f\xfb\xe4\x2a\xd8\x1a\xad\x94\x46\x53\x9c\xe7\xa6\x47\xe9\xb5\xc4\x75\x4c\x82\x39\x5a\x88\xbc\xf2\x5a\xa0\xc5\xe2\xd0\x48\x13\xb4\x4c\x20\x44\x2a\xd4\xe0\x25\x21\xd7\xa2\xd5\x21\x07\xee\xb0\x75\xb5\x72\x45\xce\x84\x5b\xbd\x2e\x16\xd0\x74\xf1\x2f\x45\x03\x59\x72\x76\xd2\x11\x0d\x2e\x1b\x95\xbd\x4f\x65\xa6\x9c\xc5\xd0\x27\xf8\x4b\x41\x8d\xa5\xfd\xce\xb8\x9e\x40\xa7\xdb\x23\x53\xb5\xbf\x46\xd5\x1f\x44\xd0\xff\xb4\xac\x69\x69\x6a\xe5\x7e\x7c\x58\xe9\x09\x59\x48\x63\xc7\x68\x9f\xa5\x5e\xc6\xc4\x6a\xb7\xb5\x5b\xc9\x7d\x34\x75\xdd\xc3\x4a\x8c\xa3\x15\x33\x76\x7b\x19\xa0\x52\x58\x60\x02\x75\xcd\x99\x65\x90\x62\x4c\xce\x06\x6f\x4e\x06\xc3\xe1\xc9\xdb\x93\xd3\x37\x51\xb2\xd4\x11\x52\x1d\x39\x13\x3e\xa3\xb1\xe1\x9b\x06\x4d\xe5\x5b\xe8\x4b\x25\x15\x2c\xce\x87\xd1\x59\xf4\xb6\xb2\xdd\x02\x73\xe2\x38\x9f\x48\xce\xe8\x3a\x26\x17\xfc\x19\xd6\xa6\xe2\xa1\xa4\xb6\x0d\x86\xc2\x7d\x80\x13\xa9\x6d\x4c\xce\x86\x67\x3f\x9f\x37\x48\x2c\xbf\x54\x19\x5a\xcd\x68\x15\xaf\xe5\x03\x59\x3e\x1a\x21\x61\x02\x8d\x99\x68\x39\xc7\xfa\x8a\xb8\xda\x17\xf5\xed\x43\x65\x96\x81\xff\x7a\x3e\x06\x7d\x50\xaa\x9f\x6a\x45\xc3\x05\x02\xb7\x8b\x50\x79\x08\x7f\xa1\x09\x21\x49\xf4\xaf\xf1\xe9\x60\x70\x3a\xdc\x5c\x8b\x76\x5b\x17\xcc\x32\xe0\x57\xc8\x61\x3d\x45\x2a\x45\x62\x62\xf2\xf6\xb4\xe2\xc3\x59\x8e\x3f\x54\x40\x28\xf2\xa6\x12\x25\x9d\x17\x7f\x4c\x67\x0f\x93\xcb\xd9\xef\xe7\xd3\xd9\xe5\xf8\x7a\x76\xf3\xf1\xc3\xcd\xe8\x61\x74\xd3\x88\xb0\x48\xf7\x98\x5c\x8d\xde\xdd\x7f\xf8\x26\x9c\x87\xd1\xa7\xdf\x26\x77\xa3\xf7\xd7\x7f\xb6\x23\xa1\x60\x9d\x38\xa3\xf1\xf5\xec\xf6\xd3\x7d\xfb\xe4\xe0\x97\xc1\x60\x18\xb4\xce\xbf\xfd\x6b\x36\xfe\x78\x35\x9a\x8d\x2f\x6e\x47\x6d\x93\xdf\x6b\x99\x1d\x9e\xda\x27\x86\x3c\xd9\xdc\xfb\x5a\xc7\x26\x60\x17\x71\x71\x43\x88\x7c\xf2\xf9\x13\x5c\x4b\xc0\xc6\x7d\x79\x6b\xfe\xe2\xd0\x34\xcf\x00\x21\x54\xb9\x98\x0c\x07\x59\xc5\x6c\x90\xba\xe2\xbe\x21\x85\xc5\x95\xad\xcf\x50\x9a\xe5\x8c\x63\x8a\x49\xad\x32\x14\x7b\x92\xdc\x65\x78\xeb\xab\xca\xc1\x51\xcb\xbc\xb5\x0c\xbc\xef\xab\x4b\x5f\x2a\xdb\xa7\x82\xf5\xe7\xac\x59\xd3\x4a\xee\xa8\x60\xe1\x9c\x89\x30\x61\xfa\x25\x28\xb4\xb4\x80\x12\x68\xa3\xe4\x28\x98\x40\xfb\x2d\x60\x39\xe8\x3e\x97\x69\x2b\x0c\x97\xe9\x0b\x10\x7e\xb6\x76\xa2\x9f\x48\xba\x44\x1d\x19\x49\x97\xad\x48\xe5\xf8\xc1\x70\x17\x98\x59\xb0\xec\x45\xc0\x05\xdb\x0a\x59\x8a\x51\xab\xbc\x5d\xc4\xfa\xbd\x17\xcb\xd6\xc5\x2e\x03\x69\xd3\x2a\xec\xa4\xb6\x13\xae\x5d\xaf\xf0\x28\xc9\x9d\x60\x4d\xbd\xc2\x2e\x86\x5f\x44\x3a\xa6\x5d\x78\x8c\xe6\x57\xa2\xee\x45\x3c\x68\x1f\x41\xb1\x7d\xef\xbb\xef\x1c\xfd\x6d\x74\xd7\x3e\x3a\x63\x65\x76\xb7\x39\xe0\x57\x58\x5c\xe9\x99\x14\x2d\xbd\x15\x0a\x46\xa5\x78\x62\xa9\x89\xda\xff\x1f\xd8\x36\x18\x86\x4a\xdf\x27\x6d\x9a\xc6\x1e\x21\x69\xd9\x67\x1e\xfb\x57\x21\x2f\xe3\xdd\x24\xd6\x96\x97\x7c\x08\x5c\x2d\x60\x58\xb9\x62\x34\x0a\x84\xb1\x52\x17\x1f\xfe\x8d\xad\xe8\xf8\x4a\x10\xc5\x9d\x06\x5e\x8d\xb9\x6c\x81\x98\x48\x1d\x07\x5d\x19\x28\xdb\xc0\x82\x8b\xd1\xf8\xfa\xb2\xb4\xfd\x13\x00\x00\xff\xff\x1b\x39\xca\x65\x78\x11\x00\x00")

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
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa3, 0xa6, 0x37, 0xb2, 0xef, 0xeb, 0xc0, 0x66, 0xbf, 0xa6, 0x9d, 0x6c, 0x95, 0x85, 0x29, 0x70, 0xd5, 0xd7, 0x4d, 0x7d, 0xb2, 0x8a, 0x9a, 0xfb, 0x52, 0x99, 0xfe, 0x3, 0x40, 0x0, 0x78, 0x77}}
	return a, nil
}

var _coredns114Json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\x5f\x53\x1b\x39\x12\x7f\xe7\x53\x4c\x4d\x55\xde\xd6\xc6\x24\x0b\x97\xe5\x8d\x05\x2e\x4b\x5d\xe0\x5c\x40\xf6\xe5\xb2\xb5\xd5\xd6\xb4\x3d\x3a\x34\x6a\x5d\x4b\x63\xf0\xa6\xf8\xee\x57\xf3\xcf\x9e\x3f\x92\xb1\xd9\xe4\x2a\x75\xb5\x7e\x48\x8c\xd4\x6a\xfd\xd4\xfd\xeb\x56\x4b\xf2\x97\x83\x28\x8a\xc1\xc8\x5f\x91\xad\x24\x1d\x9f\x46\xf1\xf2\x28\xfe\xa1\x68\x95\x0e\x33\x1b\x9f\x46\xff\x3a\x88\xa2\x28\xfa\x52\xfe\x1b\x12\x2e\x7b\x1e\xa4\x4e\x8a\xb6\x3b\xe4\xa5\x14\xb8\xe9\xc8\xd0\x41\x02\x0e\xe2\xd3\xb5\x9a\x28\x8a\x35\x64\x58\x88\x3f\xe4\x33\x1c\x25\xda\xae\xe5\xeb\x3e\x6b\x40\x6c\x04\xec\xca\x3a\xcc\xda\x32\xa0\x35\x39\x70\x92\xb4\xed\x28\x8e\xa2\xd8\x30\x65\xe8\x52\xcc\xed\x58\xd2\xa1\x21\x76\x85\x9e\x9f\x8e\x8e\xdf\xb5\x14\x0c\xe4\xac\x60\x30\xe5\x8c\x8e\x73\x8c\xd7\x82\xcf\xad\x49\x15\xcc\x50\x0d\xe6\xc3\x07\x3b\x86\x0c\xfe\x20\x0d\x8f\x76\x2c\x28\x3b\x14\x94\x19\xd2\xa8\x5d\x60\x85\x85\xb9\xde\xdb\x11\x18\xb3\x45\x20\x9f\x21\x6b\x74\x58\xa2\x13\x2a\xb7\x0e\x79\x64\x6b\xeb\x36\x30\xb7\x0c\x69\x2c\x7c\x4e\x8c\x17\x37\x77\xad\x15\x1d\xf4\x56\x16\x5b\x83\xa2\xeb\x9e\xc2\x6a\x1b\xf7\x57\x9f\xf6\xa2\x5b\x1e\xec\x23\xaf\x47\xc7\xa7\xd1\xf1\xbb\x7e\x3b\x93\x23\x41\xaa\x18\xf6\xe9\x62\xda\x1f\xe6\x80\x17\xe8\xa6\xcd\xe0\x56\xe7\xf3\x0f\x3b\xe1\x18\x39\x61\x5e\x83\xe5\xfe\x7c\x1f\x2c\xeb\xef\xbf\xb5\xa8\x61\x51\xa1\x70\xc4\x7d\x72\xf8\xfc\xec\xe5\x96\x45\x5b\x84\xd5\xd9\x7c\x2e\xb5\x74\xab\x62\xc0\x0d\xe9\xb6\x87\x63\xb7\xaa\x08\x7a\x5e\x91\xe1\x6a\xda\x68\xaa\x20\xd5\xda\xf6\x8f\xd5\x33\x21\x28\xd7\x6e\xd7\x90\x15\xc4\xf8\xe7\x23\xf6\x4f\xc7\xd5\x10\xc6\x8b\xe6\x7e\xa5\xb9\x86\xa6\x28\x62\x6a\x2e\x55\xb9\xdc\xf1\xe9\xf1\xbb\xe8\xcb\x67\x5d\x74\x20\x33\xb1\xad\xbe\xa7\x08\xca\xa5\xd5\xf7\x4d\x64\x46\x75\x24\x8f\x15\x09\x50\x91\xd4\x23\x48\x12\x1e\x03\x1b\x88\xa4\x39\xa9\xbe\xd4\xda\xa2\xc8\x50\x62\x23\xa9\x2d\x8a\x9c\xb1\x69\xcc\x8d\x75\x8c\x90\x35\x7f\xcf\x41\x29\x97\x32\xe5\x8b\xd4\xaf\xaf\x12\x7c\xae\xfe\xdb\x64\xbd\xe8\xb4\xc8\x89\x55\xeb\x9c\xf8\x11\x38\x89\xc6\xd1\x21\x3a\x71\xc8\x68\x49\x2d\xc7\x82\xf4\xbc\xea\x17\x20\x52\x8c\xde\x4d\xaa\xbf\x14\x91\xa9\xbe\x31\x2a\x82\xa4\x69\x85\x64\x06\x0a\xb4\xc0\xcf\xfa\xf9\xb3\x8e\x07\x79\xa6\xe1\xdd\x39\xe9\xb9\x5c\x5c\x83\xf9\x8b\x72\x2d\xca\x81\x31\xf6\xd0\x17\xa6\x17\x68\x14\xad\x32\xfc\xff\x0d\xd1\x6f\xb6\x8d\x31\x2d\x18\xad\xbd\x40\x48\x94\xd4\x78\x87\x82\x74\x52\x2c\xe3\x64\x32\x69\xad\x8e\xd1\x28\x29\xa0\xe8\x78\xdb\x69\x5e\xca\xc2\x3b\xbf\x48\xeb\x88\x57\x1f\x65\x26\x8b\x75\x1d\x4d\x76\x48\xfc\x19\x38\x91\x7e\xf4\x59\xed\x75\x76\x7b\x89\x69\xed\x8d\xa9\xb3\xaf\x38\x06\x87\x8b\x55\x1f\x1f\x93\x52\x52\x2f\x3e\x99\x04\x1c\x0e\x11\x66\xf0\x74\x97\xf3\xa2\x34\xfd\xdb\xe3\x37\x7d\x28\x19\x3c\x7d\xd2\xb0\x04\xa9\x60\x56\xa6\xc1\xa3\xd0\x6e\xbd\xde\xb5\x6e\x3b\x13\x7a\xb1\x3a\xcc\x8c\x1a\xc2\xf1\x93\xbd\xec\x11\x8c\x25\x5f\xef\x65\x86\xd6\x41\x56\x18\x47\xe7\x4a\xf5\xe0\x7a\xd9\xfb\x7a\x4f\xbc\xec\x8b\xb6\x37\xfa\x06\x19\xd0\xb4\x6c\x85\xcd\xc6\x3f\x00\xa9\x29\xc1\xb3\x70\x7f\xc9\xd4\xff\xe4\x92\x31\xb9\xc8\x59\xea\xc5\x9d\x48\x31\xc9\x0b\x6b\x5f\x2d\x34\xad\x9b\x2f\x9f\x50\xe4\xae\xca\x36\x43\x1d\xf5\x3c\x77\x35\x9f\xef\x91\xb3\x7e\x05\xb8\xf9\xf8\x86\x47\x6b\xda\x5f\x3e\x19\xae\xca\x99\xb0\x86\x6d\x5a\x4a\x4d\x0f\x58\x16\x41\x33\x74\x30\xee\xa6\x04\x1a\xfa\xa3\x33\x92\x0c\x32\x54\x21\x19\x5f\xe9\xed\xb2\x4b\x50\x39\x6e\x47\x59\xca\x29\xa9\xf3\xa7\x78\x8b\xcc\x6f\xc1\xbe\xe7\x30\x80\x57\x1a\x00\x58\xa4\xff\x7b\x13\x40\x96\x9c\xfc\xf8\x4a\x13\x04\x7a\xfc\x23\x7c\xd2\x43\xc9\xbe\xd4\xc0\xcc\xb1\xa1\xe4\x4c\x3b\xb9\x3d\x72\x0c\xe3\x1c\x79\x9f\xd0\xf1\x99\xc9\xef\xc8\x12\x41\x3d\x7b\x11\x4f\x81\xb8\x8b\xd6\x29\xea\xce\xbf\x99\xf4\x44\xf7\x8a\xb1\xed\x24\xdb\xd0\xac\x49\x68\xdb\xa8\xb2\x1f\xb1\x76\xa7\x56\x14\x4a\xa2\xfd\x4f\x98\x61\x61\x8e\x85\x47\x05\x03\x33\x76\x64\x48\xd1\x62\xf5\x8f\xda\x36\x9d\xe8\x4b\xc9\xba\xb2\x2a\xf1\x93\xd7\xaf\x34\x7e\x44\xb9\x48\xab\xea\x61\xe2\x91\x18\xa2\xef\xa3\xee\x4a\xf4\xa6\x89\x05\x69\x07\x52\x23\xfb\xac\xed\xa1\x3d\xf0\x22\xe4\x97\x78\x54\xd4\xfd\x5e\xd7\xc6\xe5\xe9\xa0\xde\x14\x0f\xd7\xa7\xa0\x21\xf6\xe1\xe0\x58\x66\x50\x95\x13\x6f\xec\x38\x79\xe0\x31\x0a\x1e\xbf\xb1\xe3\x37\xf6\x10\x1f\x6c\xa3\xf3\x74\x79\x34\x3e\x19\x9f\x78\x26\xaf\xc6\x4f\x73\xa5\xa6\xa4\xa4\x28\xfd\x72\x35\xbf\x21\x37\x65\xb4\xed\xd2\xb8\x35\x44\xc9\x25\x6a\xb4\x76\xca\x34\x1b\xd6\x38\xb5\xd0\x1c\xa4\xca\x19\xef\x53\x46\x9b\x92\x2a\x2a\xee\x63\xef\xda\x53\xe7\xcc\x07\x74\xc1\xd8\x8c\x0d\xb8\xb4\x80\x75\x58\x1d\xfd\x02\xc1\xd1\xdc\x48\xbc\x9f\xbc\x9f\x04\x24\xac\x48\xb1\xaa\x7a\x7f\xb9\xbf\x9f\xfa\x78\xe6\x65\x59\x5c\x24\x1a\x09\xea\x02\x15\xac\xda\xa5\xae\x57\xd8\x20\x4b\x4a\x36\x62\x47\x7e\x31\x9b\x0b\x81\xd6\xb6\xcd\x73\xe4\x15\x74\x32\x43\xca\xdd\x46\xe1\xf1\x30\x65\x7b\x7c\xc4\x08\x89\x7c\xd1\x49\xdf\x95\xed\x77\x5a\x57\xf8\x28\xd6\x01\x13\x0a\xc2\xc0\x32\xd7\x61\x3e\xf5\x5f\x6b\x0d\xa6\xf7\x4f\x1d\xf9\x6e\xe3\x76\xa5\xd9\x57\x84\xe6\xb9\xb0\xf3\xc2\xbb\x3f\xff\xea\xf0\x7e\x3a\x7a\x11\x60\x86\x8e\xa5\xd8\xc9\x7e\x21\x80\x3b\x65\x46\x46\x4b\x39\x0b\xf4\x1d\x50\xa2\x2a\x8f\x65\xd2\x85\x7a\xcb\xf3\x51\x46\x5c\x66\xc4\xa3\xbf\x4d\xae\xe5\xee\x09\xa3\x38\x36\xa0\xdd\xa6\x5a\x98\xbc\xd4\x3b\x99\x64\x21\x3b\x6c\x66\x0f\x4e\xbe\x53\xc4\x94\xb7\x5c\xd2\xad\xce\x49\x3b\x7c\x0a\xc5\x7a\x0c\x4a\xd1\xe3\x94\xe5\x52\x2a\x5c\xe0\xa5\x15\xa0\xa0\xae\xcc\xe6\xa0\x2c\x7a\x17\x2a\xc0\xc0\x4c\x2a\xe9\x64\xd0\xca\x85\xea\x24\xd9\x52\xab\xc4\x37\x97\xf7\xbf\xff\x7c\x75\x73\xf1\xfb\xdd\xe5\xed\xaf\x57\xe7\x97\xfe\x02\xc0\xe3\xe1\x72\x74\xc2\x64\xb6\x69\x07\xa5\x02\x0a\xf7\x70\x27\x24\xff\xd4\x6a\x75\x4b\xe4\xfe\x2e\x15\xd6\x37\x3c\xa7\x91\xe3\x1c\x77\xf2\x81\x43\xce\xa4\x2e\xed\x79\x8d\xd6\x16\x3b\x6e\x93\x58\x13\x5c\x1e\xb6\xba\x47\x8a\x16\xbe\xa4\xe6\xd1\xb0\xde\xaf\x0b\x4c\xbe\x31\x4b\x52\x79\x86\xd7\x94\xeb\x7d\xf3\x61\x56\x8c\x59\x43\x6c\x15\x27\x21\xb2\x6e\xd2\xb2\x9e\xcb\xc5\xa8\x9a\x39\x24\xdc\xd8\x33\x64\xc1\xfd\x32\x50\x17\xab\xcb\x82\x99\xaf\xc1\x58\x88\x7c\x85\x52\xb1\x47\xc8\x22\xef\x6e\x5c\x72\x81\x73\xc8\x55\xbf\x7a\x8a\x0d\x4b\x2a\x63\x51\x81\xb5\x37\x35\x9e\x8a\x4e\x23\x4d\x09\x8e\x04\x4b\x27\x05\xa8\xfe\x40\x46\xeb\x80\xdd\x66\x82\x33\xf5\x08\xab\xc1\x55\x96\xad\x8e\x59\xc8\x8d\xee\xa4\xc2\x31\x5a\x77\x0c\x46\x0c\xb3\x43\xbf\x06\xb6\xdd\xe7\x8c\xf0\x45\x5a\x57\xf0\x66\xeb\x4e\xdd\x26\xf4\x07\x06\x81\xd3\x5e\xdd\xf4\x6e\xd2\x1f\x40\xaa\x38\x23\x05\x4e\x66\x9e\x5a\x1c\xe7\x73\x14\xae\x7a\xf1\xa9\xcf\x9f\xde\x30\xa9\x4f\x6a\xa5\xfd\x99\x14\xf6\x6e\x05\x32\xb0\x0e\xb9\xcf\x98\x01\x43\x3d\x00\x6a\xc5\xe7\xb5\x4f\xcf\x92\x84\xb4\x2d\x89\xef\x41\xd1\x3e\x01\x5e\x3e\xc9\x62\xf7\xd8\x8b\x7e\x55\xc8\xed\x6a\x9a\xcc\xb8\xd5\x85\x64\x8f\xbb\xa3\x17\x42\x65\x97\x85\x8b\xf5\x43\x44\x60\xb3\xa9\x69\x79\x4d\x49\x31\xcd\x8f\x6f\xfd\x35\x72\xf7\xb1\xbc\xff\x09\x1e\xf9\x1b\xb3\x37\xc7\xa8\xe0\x71\xb4\x29\x6e\xc3\x07\xae\x28\x74\x6d\xe2\xc5\xdb\xaf\x4d\xf7\x2c\x68\xdb\x99\x73\xbb\xeb\x0f\x7c\x3d\xfb\xbc\x89\xf0\x0c\xc4\x18\x72\x97\x12\xcb\x3f\xca\x98\x1a\x3f\xbc\x2f\xd9\xee\x7b\x28\xa9\x1f\x46\x6f\x49\xbd\xf8\xfb\x83\xef\xf1\xa9\x63\x46\xe4\xac\x63\x30\x46\xea\x45\xb3\xfa\x51\x4d\xc1\xc0\xb3\xb1\xee\xe4\xe6\xd3\x9e\x47\x37\x6f\x23\x9c\xab\x5e\xcc\x75\x56\x07\x46\x7e\x60\xca\xcd\x90\xc5\x71\xdb\xc5\x1d\x3a\x75\xea\xd6\xde\x20\xd4\x89\x21\x59\x6c\xe7\xfe\xd4\xeb\xf9\xc9\x40\x32\x68\x5b\x3f\x59\xd9\x30\x86\x25\xf2\xcc\x33\xbf\x92\x76\xb0\xa7\x3d\x82\x13\x69\x47\x93\xcf\xa2\xdf\xd4\x2c\x45\xea\xde\x7f\x31\x0b\x74\x01\xd8\x07\xed\x96\x6f\x14\x49\x3f\x4b\x9d\x14\x7c\xdc\x2f\xa0\x3a\x06\xf1\x4c\xdd\xbd\xcf\xce\x1d\xe5\xcd\x43\xd4\xd7\xfe\xfd\xcd\x77\x1f\x99\xa4\xf0\x16\xe7\x3d\x6b\xd6\xc4\xdb\xea\xb9\xf6\xd3\xed\xb6\x14\xb8\x07\x18\x9b\xcf\xfe\x8d\xc2\x6d\xc9\x14\x2f\xfc\x76\x24\x7a\xe9\xee\x23\xfc\x10\x1d\xa2\xf5\x41\x1d\x26\xeb\xa9\x3f\x16\xc1\x7d\xf0\x7c\xf0\xdf\x00\x00\x00\xff\xff\x15\x5b\xde\x84\xbf\x26\x00\x00")

func coredns114JsonBytes() ([]byte, error) {
	return bindataRead(
		_coredns114Json,
		"coredns-1.14.json",
	)
}

func coredns114Json() (*asset, error) {
	bytes, err := coredns114JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "coredns-1.14.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc, 0xad, 0x9b, 0x3c, 0xec, 0x20, 0x42, 0x62, 0xcd, 0x55, 0x1d, 0x3d, 0xf9, 0xa0, 0xf2, 0x4d, 0x2f, 0xd0, 0x6c, 0x2c, 0xf9, 0x5b, 0xd0, 0x7b, 0xdd, 0x29, 0x80, 0x86, 0x38, 0x73, 0xeb, 0xcd}}
	return a, nil
}

var _coredns115Json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\x5f\x73\x1b\xb9\x0d\x7f\xf7\xa7\xd8\xd9\x99\xbc\x9d\x64\x29\x3e\xbb\x39\xbf\xf9\x2c\x37\xe7\x69\xec\x6a\x6c\xe7\x5e\x9a\x9b\x1b\x8a\x0b\x49\xac\xb9\x04\x0b\x72\x65\x2b\x19\x7f\xf7\x0e\xf7\x8f\xb4\x7f\xc8\x95\xe4\x3a\x9d\x74\xa6\x7a\x48\xd6\x4b\x00\x04\x01\xfc\x40\x80\xdc\x6f\x47\x51\x14\x33\x2d\x7e\x07\x32\x02\x55\x7c\x1e\xc5\xab\x71\xfc\x93\x7b\x2b\x2c\xa4\x26\x3e\x8f\xfe\x71\x14\x45\x51\xf4\x2d\xff\x37\x44\x9c\x8f\x3c\x0a\x95\xb8\x77\xf7\x40\x2b\xc1\x61\x3b\x90\x82\x65\x09\xb3\x2c\x3e\xdf\x88\x71\x82\x94\x42\xcb\xac\x40\x65\x1a\x03\x51\x14\x6b\xc2\x14\xec\x12\x32\x33\x14\x78\xac\x91\xac\x93\xfb\xcb\xf8\xf4\x64\x23\xd4\x43\x67\x38\x31\x0d\x8e\xd2\x52\x06\xf1\x86\xf0\x65\xcb\x13\x4b\x36\x03\xd9\x99\x0f\x1e\xcd\x90\xa5\xec\x2b\x2a\xf6\x64\x86\x1c\xd3\x63\x8e\xa9\x46\x05\x2a\x9f\xf8\x31\x9b\xc1\x20\x51\xa6\x39\xf9\xe3\x07\x33\x60\x5a\xf7\x10\x64\x33\x20\x05\x16\x72\xed\xb8\xcc\x8c\x05\x1a\x98\xd2\x3a\x95\x9a\x3d\x2c\x8a\xa5\x39\xdd\x25\x12\x4c\x6e\xef\xfd\x2b\xaa\x88\x3c\x3a\xe4\x63\x46\x33\xbe\x25\x30\x6b\x63\x21\xad\x24\x6d\xe4\xc4\x46\x03\x6f\xba\xc7\x59\x7d\xeb\xfe\xe2\x57\x37\x5a\x6d\xea\xf6\xca\x4b\xee\xf8\x3c\x3a\x3d\x69\xbf\x27\xb4\xc8\x51\x3a\xb6\xcf\x93\x69\x9b\xcd\x32\x5a\x80\x9d\x56\xcc\xb5\xc1\x97\x9f\xf6\xd2\x63\x60\xb9\x7e\x8d\x2e\x0f\x97\x87\xe8\xb2\x79\xfe\xa3\x66\x6c\x03\x12\xb8\x45\x6a\x07\x97\x2f\x4e\xbc\x9e\x34\x60\x1c\xac\x2e\xe6\x73\xa1\x84\x5d\x3b\x86\x5b\x54\xf5\x08\x89\xed\xba\x08\xf0\xcb\x22\x98\xae\xa7\x1b\x4f\x1e\xd5\xa4\x1d\x8e\xd5\x0b\xce\x31\x53\xf6\x40\xc8\xfe\xc7\xc0\xe2\x48\xb0\x17\xae\x7a\x23\xbf\x2b\x65\x77\xe0\x1f\x68\xae\xae\x29\x1c\x26\xe7\x42\xe6\x13\x0c\xcf\x4f\x4f\xa2\x6f\x5f\x94\x1b\x00\x22\x24\x53\x3c\x2f\x81\x49\xbb\x2c\x9e\xb7\xc8\x8e\xca\x4c\x30\x94\xc8\x99\x8c\x84\x1a\xb0\x24\xa1\x21\x23\xcd\x22\xa1\xcf\x8a\x87\x52\x5a\x14\x69\x4c\x4c\x24\x94\x01\x9e\x11\x54\x2f\x33\x6d\x2c\x01\x4b\xab\xbf\xe7\x4c\x4a\xbb\x24\xcc\x16\x4b\xbf\xbc\x82\xf0\xa5\xf8\x6f\x9b\x35\xa3\x73\x97\x53\x8b\xb7\x73\xa4\x27\x46\x49\x34\x8c\x8e\xc1\xf2\x63\x02\x83\x72\x35\xe4\xa8\xe6\xc5\x38\x67\x7c\x09\xd1\xc9\xa8\xf8\x4b\x22\xea\xe2\x89\x40\x22\x4b\xaa\xb7\x2c\x99\x31\xc9\x14\x87\x2f\xea\xe5\x8b\xea\xe6\x99\x2a\xee\x2e\x51\xcd\xc5\xe2\x86\xe9\xff\x87\x5c\x2d\xe4\x98\xd6\xe6\xd8\x07\xd3\x09\x68\x89\xeb\x14\x7e\x60\x88\xbe\xd5\x3e\xf6\x1a\xd3\xf6\x6d\x63\x84\x0b\x02\x63\x26\xc0\x12\x29\x14\xdc\x03\x47\x95\x38\x2b\x9c\x8d\x46\xb5\x39\x08\xb4\x14\x9c\xb9\x81\xf7\x8d\xd7\x2b\xe1\xbc\xf3\x9b\x30\x16\x69\xfd\x49\xa4\xc2\x99\x65\x3c\xda\x23\xf1\xa7\xcc\xf2\xe5\x27\x9f\xd1\x5f\x67\xf6\x5d\x81\x5a\xdf\x98\x1a\xfb\x8a\x25\x66\x61\xb1\x6e\xeb\x47\x28\xa5\x50\x8b\xcf\x3a\x61\x16\xba\x1a\xa6\xec\xf9\x3e\xa3\x45\x6e\xf4\xf7\xa7\xef\xda\xaa\xa4\xec\xf9\xb3\x62\x2b\x26\x24\x9b\xe5\x69\x70\x1c\xda\xad\x37\xbb\xd6\x5d\x63\x42\xaf\xae\x16\x52\x2d\xbb\xea\xf8\x83\x3d\xea\x2f\x23\x7b\xec\x9c\x59\x18\x54\x5a\x01\x7f\x1f\x37\xb8\x5e\x5a\x2b\xe5\x04\xb9\xfc\x07\x91\x82\xb1\x2c\x75\xe6\x57\x99\x94\x2d\x32\x2f\xbc\x7a\x74\xd8\xe1\xeb\xdd\xde\xae\xfb\xbb\x6d\xf2\x0e\x10\x0a\x5b\x6d\x4b\x8b\x8e\x92\x0a\x13\xb8\x08\x8f\xe7\x58\xf8\x57\x26\x08\x92\x49\x46\x42\x2d\xee\xf9\x12\x92\xcc\xf9\xf3\x7a\xa1\x70\xf3\xfa\xea\x19\x78\x66\x8b\x7c\xd6\x95\x51\xce\x73\x5f\x22\xe6\x01\x28\x6d\xd7\x98\xdb\x9f\x8f\x3d\xda\x00\xeb\xea\x59\x53\x51\x30\x85\x25\xf4\x49\xc9\x25\x3d\x42\x5e\x66\xcd\xc0\xb2\x61\x33\x67\x61\xd7\x1f\x0d\x4e\xd4\x40\xac\x00\x7d\x7c\xad\xfa\x69\x57\x4c\x66\xd0\xaf\x65\x4e\x27\x85\xca\x9e\xe3\x1e\x9a\x3f\x82\x63\x2f\x61\x05\x5e\x69\x00\x46\x7c\xf9\xdf\x37\x01\x4b\x93\xb3\x9f\x5f\x69\x82\xc0\x88\x9f\xc3\x47\xdd\xa5\x6c\x53\x75\xcc\x1c\x6b\x4c\x2e\x94\x15\xfd\xc8\xd1\x04\x73\xa0\x43\xa0\xe3\x33\x93\xdf\x91\xb9\x06\xe5\xec\x0e\x4f\x01\xdc\x45\x9b\x14\x75\xef\xdf\xae\x5a\xa4\x07\x61\xac\x3f\xc8\xb6\x61\x56\x25\xb4\xbe\x50\x39\x2c\xb0\xf6\x0f\xad\x28\x94\x44\xdb\xbf\x70\x84\x85\x63\x2c\xcc\x15\x04\x66\x6c\x51\xa3\xc4\xc5\xfa\x6f\xa5\x6d\x1a\xe8\x5b\xa2\xb1\x79\x45\xe4\x0f\x5e\xbf\xd0\xf8\x09\xc4\x62\x59\xd4\x27\x23\x0f\x45\x57\xfb\xb6\xd6\x2f\xfd\x1b\x21\x2a\xcb\x84\x02\xf2\x59\xdb\x13\xf6\x8c\x16\x21\xbf\xc4\x03\xd7\x59\x78\x5d\x1b\xe7\xfd\x47\xb9\x29\x1e\x6f\xfa\xac\xae\xee\x5d\xe6\x58\xa4\xac\x28\x58\xde\x99\x61\xf2\x48\x43\xe0\x34\x7c\x67\x86\xef\xcc\x31\x3c\x9a\x4a\xe6\xf9\x6a\x3c\x3c\x1b\x9e\x79\x26\x2f\xf8\xa7\x99\x94\x53\x94\x82\xe7\x7e\xb9\x9e\xdf\xa2\x9d\x12\x98\x7a\xf1\x5d\x63\x91\x62\x05\x0a\x8c\x99\x12\xce\xba\x55\x54\x49\x34\x67\x42\x66\x04\x0f\x4b\x02\xb3\x44\xe9\x6a\xfa\x53\xef\xda\x97\xd6\xea\x8f\x60\x83\xd8\x8c\x35\xb3\x4b\xa7\xd6\x71\xd1\x5c\x06\xc0\x51\x9d\x79\x7c\x18\x7d\x18\x05\x28\x0c\x5f\x42\x51\x71\xff\xf6\xf0\x30\xf5\xc5\x99\x37\xca\x62\x97\x68\x04\x93\x13\x90\x6c\x5d\x2f\xa6\xbd\xc4\x1a\x48\x60\xb2\x25\x1b\xfb\xc9\x4c\xc6\x39\x18\x53\x37\xcf\xd8\x4b\x68\x45\x0a\x98\xd9\xad\xc0\xd3\x6e\xca\xf6\xf8\x28\xdc\x5b\x6c\x35\xf5\x1c\x77\x55\xbf\x80\x2b\x36\x70\x98\xfa\x0f\x98\x3a\xd3\xfb\xa7\x8e\x7c\xe7\x62\xfb\xba\xe3\x0d\x55\xf3\x1c\x9d\x79\xd5\x7b\xb8\x7c\x73\xf5\x7e\x19\xef\x54\x30\x05\x4b\x82\xef\x65\xbf\x90\x82\x7b\x65\x10\x02\x96\x88\xd7\xe0\xd9\xbb\x80\xff\x01\x3c\x7f\x77\x88\x8e\xf7\x82\x28\x81\xc1\x8c\x38\xf8\x1a\xa8\xa8\xc8\xb3\xa9\xb0\xa1\xd1\xbc\x43\x4c\x91\xf2\x8c\x3d\xfe\xcb\xe8\x46\xec\x6f\x00\xd7\xd6\x80\xe9\x13\xcd\x75\x96\xcb\x1d\x8d\xd2\x90\x7f\xb6\xb3\x07\x27\xdf\xcb\x0c\xf9\x39\x9f\xb0\xeb\x4b\x54\x16\x9e\x43\x71\x13\x33\x29\xf1\x69\x4a\x62\x25\x24\x2c\xe0\xca\x70\x26\x59\x59\x39\xce\x99\x34\xe0\x5d\x28\x67\x9a\xcd\x84\x14\x56\x04\xad\xec\x44\x27\x49\x4f\x2d\x15\xdf\x5e\x3d\xfc\xf9\xeb\xf5\xed\xe4\xcf\xfb\xab\xbb\xdf\xaf\x2f\xaf\xfc\x05\x8a\x07\x59\x39\x77\x42\xa8\xfb\xa4\x33\x29\x03\x02\x0f\x70\x27\x4b\xfe\xae\xe4\xfa\x0e\xd1\xfe\x55\x48\x28\x8f\x8a\xce\x23\x4b\x19\xec\xe5\x03\x0b\x94\x0a\x95\xdb\xf3\x06\x8c\x71\x15\x41\x05\xd2\x04\x56\xc7\xb5\xe1\x81\xc4\x85\x6f\x33\xf1\x48\xd8\xd4\x13\x4e\x27\x1f\xcf\x0a\x65\x96\xc2\x0d\x66\xea\xd0\x7d\x28\x75\x3c\x1b\x15\x6b\xc5\x53\x28\x58\xb7\xdb\xa1\x9a\x8b\xc5\xa0\x98\x39\x44\x5c\xd9\x33\x64\xc1\xc3\x32\x7f\x53\x57\x9b\x06\x77\x9c\x4a\x47\x47\xf2\x06\xa5\x6c\x2b\x20\xdd\x7e\xb7\x75\xc9\x04\xe6\x2c\x93\xed\xea\x2e\xd6\x24\x30\xc7\xa2\x64\xc6\xdc\x96\xfa\x14\xe1\x34\xa8\x6e\xf3\x38\x09\x2b\x38\x93\x6d\x5e\x02\x63\x19\xd9\xed\x1c\x17\xf2\x89\xad\x3b\xe7\x79\xa6\xe8\x04\x81\x2a\xf1\x49\xa1\xca\x60\x33\xd0\xe1\xe8\x26\x88\x76\x99\x6e\x9a\x77\x3a\xe1\xd3\xc4\x26\xe1\x6d\x6f\x91\x54\x8f\xe9\x8f\xc4\x38\x4c\x5b\xfb\xc6\xc9\xa8\xcd\x80\xd2\xb5\x71\x81\xe6\xd1\xd3\x2e\xc0\x7c\x0e\xdc\x16\xd7\x5e\x65\x8b\xec\x45\x4a\xd9\x4c\x2a\x4c\x60\x40\x28\xa1\x75\x70\x91\x32\xe7\x99\x76\xd0\x74\x82\xd4\xa3\x40\x29\xf8\xb2\xf4\xe9\x45\x92\xa0\x32\x79\xec\x7b\xb4\xa8\x37\xa9\x57\xcf\xc2\x6d\x20\x07\x45\x60\x81\xba\x7d\x4d\x93\x6a\xbb\x9e\x08\xf2\xb8\x3b\xda\x81\x96\x7d\x16\xce\x37\xb7\x31\x81\xfd\xa6\x0c\xcb\x1b\x4c\xdc\x34\x3f\xbf\xf7\xd7\x08\xcd\x2f\x06\xda\xbf\xe0\xa9\x44\x65\xf6\xaa\xd3\x0b\x76\xcc\x55\xad\x14\xee\x09\xa3\xd0\xc9\x8e\x57\xdf\x76\x5b\x70\x60\x2f\x51\x4f\x9e\xfd\xae\x3f\xf2\x8d\x54\x4f\xfb\x5c\x0c\xd1\x8c\xf1\x21\xcb\xec\x12\x49\x7c\xcd\x31\x35\x7c\xfc\x90\x47\xbb\xef\xb6\xa8\xbc\x1d\xbe\x43\x79\xe8\x47\x18\x3f\xc4\x75\xd1\x0c\xd1\x1a\x4b\x4c\x6b\xa1\x16\xd5\xea\x07\x65\x08\xee\xb8\x98\x2b\xd2\xf3\x79\xcb\xa3\xdb\x0b\x22\xca\x64\x0b\x73\x8d\xd5\x31\x2d\x3e\x12\x66\xba\x1b\xc5\x71\xdd\xc5\x8d\x70\x6a\x94\xae\x2d\x26\x50\x89\x46\xe1\x76\x74\x7f\xea\xf5\x7c\x37\x91\x74\xde\x6d\xae\xbf\x4c\x58\x87\x15\xd0\xcc\x33\xbf\x14\xa6\xb3\xad\x3d\x31\xcb\x97\x0d\x49\x3e\x8b\x7e\x57\xb3\xb8\xd4\x7d\xf8\x62\x16\x60\x03\x6a\x1f\xd5\xdf\x7c\x27\x24\xfd\x2a\x54\xe2\xe2\xf1\x30\x40\x35\x0c\xe2\x99\xba\x79\xe4\x9e\x59\xcc\xaa\xdb\xb8\xb7\xfe\x88\xe9\x87\x47\x26\x4a\xb8\x83\x79\xcb\x9a\x65\xe0\xf5\x7a\xae\x7e\x55\xdc\x97\x02\x0f\x50\xc6\x64\xb3\x7f\x02\xb7\x3d\x99\x62\xc7\x07\x34\xd1\xae\x63\xa7\x9d\x97\xda\x9d\xb0\x3e\x2a\x61\xb2\x99\xfa\x93\x03\xf7\xd1\xcb\xd1\xbf\x03\x00\x00\xff\xff\xcc\xb2\x2a\xb9\xc4\x27\x00\x00")

func coredns115JsonBytes() ([]byte, error) {
	return bindataRead(
		_coredns115Json,
		"coredns-1.15.json",
	)
}

func coredns115Json() (*asset, error) {
	bytes, err := coredns115JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "coredns-1.15.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xb2, 0x4a, 0x2d, 0xc0, 0xc0, 0x44, 0x91, 0xb8, 0x4a, 0x68, 0xe9, 0x44, 0x50, 0x1, 0x27, 0x85, 0x8b, 0xd7, 0x19, 0x2a, 0x90, 0xbe, 0x39, 0x46, 0xc, 0xc8, 0xb8, 0x3a, 0x93, 0xd7, 0x45, 0xf0}}
	return a, nil
}

var _coredns116Json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\x5f\x73\x1b\xb9\x0d\x7f\xf7\xa7\xd8\xd9\x99\xbc\x9d\x64\x29\x3e\xbb\x39\xbf\xf9\x2c\x37\xe7\x69\xec\x6a\x6c\xe7\x5e\x9a\x9b\x1b\x8a\x0b\x49\xac\xb9\x04\x0b\x72\x65\x2b\x19\x7f\xf7\x0e\xf7\x8f\xb4\x7f\xc8\x95\xe4\x3a\x9d\x74\xa6\x7a\x48\xd6\x4b\x00\x04\x01\xfc\x40\x80\xdc\x6f\x47\x51\x14\x33\x2d\x7e\x07\x32\x02\x55\x7c\x1e\xc5\xab\x71\xfc\x93\x7b\x2b\x2c\xa4\x26\x3e\x8f\xfe\x71\x14\x45\x51\xf4\x2d\xff\x37\x44\x9c\x8f\x3c\x0a\x95\xb8\x77\xf7\x40\x2b\xc1\x61\x3b\x90\x82\x65\x09\xb3\x2c\x3e\xdf\x88\x71\x82\x94\x42\xcb\xac\x40\x65\x1a\x03\x51\x14\x6b\xc2\x14\xec\x12\x32\x33\x14\x78\xac\x91\xac\x93\xfb\xcb\xf8\xf4\x64\x23\xd4\x43\x67\x38\x31\x0d\x8e\xd2\x52\x06\xf1\x86\xf0\x65\xcb\x13\x4b\x36\x03\xd9\x99\x0f\x1e\xcd\x90\xa5\xec\x2b\x2a\xf6\x64\x86\x1c\xd3\x63\x8e\xa9\x46\x05\x2a\x9f\xf8\x31\x9b\xc1\x20\x51\xa6\x39\xf9\xe3\x07\x33\x60\x5a\xf7\x10\x64\x33\x20\x05\x16\x72\xed\xb8\xcc\x8c\x05\x1a\x98\xd2\x3a\x95\x9a\x3d\x2c\x8a\xa5\x39\xdd\x25\x12\x4c\x6e\xef\xfd\x2b\xaa\x88\x3c\x3a\xe4\x63\x46\x33\xbe\x25\x30\x6b\x63\x21\xad\x24\x6d\xe4\xc4\x46\x03\x6f\xba\xc7\x59\x7d\xeb\xfe\xe2\x57\x37\x5a\x6d\xea\xf6\xca\x4b\xee\xf8\x3c\x3a\x3d\x69\xbf\x27\xb4\xc8\x51\x3a\xb6\xcf\x93\x69\x9b\xcd\x32\x5a\x80\x9d\x56\xcc\xb5\xc1\x97\x9f\xf6\xd2\x63\x60\xb9\x7e\x8d\x2e\x0f\x97\x87\xe8\xb2\x79\xfe\xa3\x66\x6c\x03\x12\xb8\x45\x6a\x07\x97\x2f\x4e\xbc\x9e\x34\x60\x1c\xac\x2e\xe6\x73\xa1\x84\x5d\x3b\x86\x5b\x54\xf5\x08\x89\xed\xba\x08\xf0\xcb\x22\x98\xae\xa7\x1b\x4f\x1e\xd5\xa4\x1d\x8e\xd5\x0b\xce\x31\x53\xf6\x40\xc8\xfe\xc7\xc0\xe2\x48\xb0\x17\xae\x7a\x23\xbf\x2b\x65\x77\xe0\x1f\x68\xae\xae\x29\x1c\x26\xe7\x42\xe6\x13\x0c\xcf\x4f\x4f\xa2\x6f\x5f\x94\x1b\x00\x22\x24\x53\x3c\x2f\x81\x49\xbb\x2c\x9e\xb7\xc8\x8e\xca\x4c\x30\x94\xc8\x99\x8c\x84\x1a\xb0\x24\xa1\x21\x23\xcd\x22\xa1\xcf\x8a\x87\x52\x5a\x14\x69\x4c\x4c\x24\x94\x01\x9e\x11\x54\x2f\x33\x6d\x2c\x01\x4b\xab\xbf\xe7\x4c\x4a\xbb\x24\xcc\x16\x4b\xbf\xbc\x82\xf0\xa5\xf8\x6f\x9b\x35\xa3\x73\x97\x53\x8b\xb7\x73\xa4\x27\x46\x49\x34\x8c\x8e\xc1\xf2\x63\x02\x83\x72\x35\xe4\xa8\xe6\xc5\x38\x67\x7c\x09\xd1\xc9\xa8\xf8\x4b\x22\xea\xe2\x89\x40\x22\x4b\xaa\xb7\x2c\x99\x31\xc9\x14\x87\x2f\xea\xe5\x8b\xea\xe6\x99\x2a\xee\x2e\x51\xcd\xc5\xe2\x86\xe9\xff\x87\x5c\x2d\xe4\x98\xd6\xe6\xd8\x07\xd3\x09\x68\x89\xeb\x14\x7e\x60\x88\xbe\xd5\x3e\xf6\x1a\xd3\xf6\x6d\x63\x84\x0b\x02\x63\x26\xc0\x12\x29\x14\xdc\x03\x47\x95\x38\x2b\x9c\x8d\x46\xb5\x39\x08\xb4\x14\x9c\xb9\x81\xf7\x8d\xd7\x2b\xe1\xbc\xf3\x9b\x30\x16\x69\xfd\x49\xa4\xc2\x99\x65\x3c\xda\x23\xf1\xa7\xcc\xf2\xe5\x27\x9f\xd1\x5f\x67\xf6\x5d\x81\x5a\xdf\x98\x1a\xfb\x8a\x25\x66\x61\xb1\x6e\xeb\x47\x28\xa5\x50\x8b\xcf\x3a\x61\x16\xba\x1a\xa6\xec\xf9\x3e\xa3\x45\x6e\xf4\xf7\xa7\xef\xda\xaa\xa4\xec\xf9\xb3\x62\x2b\x26\x24\x9b\xe5\x69\x70\x1c\xda\xad\x37\xbb\xd6\x5d\x63\x42\xaf\xae\x16\x52\x2d\xbb\xea\xf8\x83\x3d\xea\x2f\x23\x7b\xec\x9c\x59\x18\x54\x5a\x01\x7f\x1f\x37\xb8\x5e\x5a\x2b\xe5\x04\xb9\xfc\x07\x91\x82\xb1\x2c\x75\xe6\x57\x99\x94\x2d\x32\x2f\xbc\x7a\x74\xd8\xe1\xeb\xdd\xde\xae\xfb\xbb\x6d\xf2\x0e\x10\x0a\x5b\x6d\x4b\x8b\x8e\x92\x0a\x13\xb8\x08\x8f\xe7\x58\xf8\x57\x26\x08\x92\x49\x46\x42\x2d\xee\xf9\x12\x92\xcc\xf9\xf3\x7a\xa1\x70\xf3\xfa\xea\x19\x78\x66\x8b\x7c\xd6\x95\x51\xce\x73\x5f\x22\xe6\x01\x28\x6d\xd7\x98\xdb\x9f\x8f\x3d\xda\x00\xeb\xea\x59\x53\x51\x30\x85\x25\xf4\x49\xc9\x25\x3d\x42\x5e\x66\xcd\xc0\xb2\x61\x33\x67\x61\xd7\x1f\x0d\x4e\xd4\x40\xac\x00\x7d\x7c\xad\xfa\x69\x57\x4c\x66\xd0\xaf\x65\x4e\x27\x85\xca\x9e\xe3\x1e\x9a\x3f\x82\x63\x2f\x61\x05\x5e\x69\x00\x46\x7c\xf9\xdf\x37\x01\x4b\x93\xb3\x9f\x5f\x69\x82\xc0\x88\x9f\xc3\x47\xdd\xa5\x6c\x53\x75\xcc\x1c\x6b\x4c\x2e\x94\x15\xfd\xc8\xd1\x04\x73\xa0\x43\xa0\xe3\x33\x93\xdf\x91\xb9\x06\xe5\xec\x0e\x4f\x01\xdc\x45\x9b\x14\x75\xef\xdf\xae\x5a\xa4\x07\x61\xac\x3f\xc8\xb6\x61\x56\x25\xb4\xbe\x50\x39\x2c\xb0\xf6\x0f\xad\x28\x94\x44\xdb\xbf\x70\x84\x85\x63\x2c\xcc\x15\x04\x66\x6c\x51\xa3\xc4\xc5\xfa\x6f\xa5\x6d\x1a\xe8\x5b\xa2\xb1\x79\x45\xe4\x0f\x5e\xbf\xd0\xf8\x09\xc4\x62\x59\xd4\x27\x23\x0f\x45\x57\xfb\xb6\xd6\x2f\xfd\x1b\x21\x2a\xcb\x84\x02\xf2\x59\xdb\x13\xf6\x8c\x16\x21\xbf\xc4\x03\xd7\x59\x78\x5d\x1b\xe7\xfd\x47\xb9\x29\x1e\x6f\xfa\xac\xae\xee\x5d\xe6\x58\xa4\xac\x28\x58\xde\x99\x61\xf2\x48\x43\xe0\x34\x7c\x67\x86\xef\xcc\x31\x3c\x9a\x4a\xe6\xf9\x6a\x3c\x3c\x1b\x9e\x79\x26\x2f\xf8\xa7\x99\x94\x53\x94\x82\xe7\x7e\xb9\x9e\xdf\xa2\x9d\x12\x98\x7a\xf1\x5d\x63\x91\x62\x05\x0a\x8c\x99\x12\xce\xba\x55\x54\x49\x34\x67\x42\x66\x04\x0f\x4b\x02\xb3\x44\xe9\x6a\xfa\x53\xef\xda\x97\xd6\xea\x8f\x60\x83\xd8\x8c\x35\xb3\x4b\xa7\xd6\x71\xd1\x5c\x06\xc0\x51\x9d\x79\x7c\x18\x7d\x18\x05\x28\x0c\x5f\x42\x51\x71\xff\xf6\xf0\x30\xf5\xc5\x99\x37\xca\x62\x97\x68\x04\x93\x13\x90\x6c\x5d\x2f\xa6\xbd\xc4\x1a\x48\x60\xb2\x25\x1b\xfb\xc9\x4c\xc6\x39\x18\x53\x37\xcf\xd8\x4b\x68\x45\x0a\x98\xd9\xad\xc0\xd3\x6e\xca\xf6\xf8\x28\xdc\x5b\x6c\x35\xf5\x1c\x77\x55\xbf\x80\x2b\x36\x70\x98\xfa\x0f\x98\x3a\xd3\xfb\xa7\x8e\x7c\xe7\x62\xfb\xba\xe3\x0d\x55\xf3\x1c\x9d\x79\xd5\x7b\xb8\x7c\x73\xf5\x7e\x19\xef\x54\x30\x05\x4b\x82\xef\x65\xbf\x90\x82\x7b\x65\x10\x02\x96\x88\xd7\xe0\xd9\xbb\x80\xff\x01\x3c\x7f\x77\x88\x8e\xf7\x82\x28\x81\xc1\x8c\x38\xf8\x1a\xa8\xa8\xc8\xb3\xa9\xb0\xa1\xd1\xbc\x43\x4c\x91\xf2\x8c\x3d\xfe\xcb\xe8\x46\xec\x6f\x00\xd7\xd6\x80\xe9\x13\xcd\x75\x96\xcb\x1d\x8d\xd2\x90\x7f\xb6\xb3\x07\x27\xdf\xcb\x0c\xf9\x39\x9f\xb0\xeb\x4b\x54\x16\x9e\x43\x71\x13\x33\x29\xf1\x69\x4a\x62\x25\x24\x2c\xe0\xca\x70\x26\x59\x59\x39\xce\x99\x34\xe0\x5d\x28\x67\x9a\xcd\x84\x14\x56\x04\xad\xec\x44\x27\x49\x4f\x2d\x15\xdf\x5e\x3d\xfc\xf9\xeb\xf5\xed\xe4\xcf\xfb\xab\xbb\xdf\xaf\x2f\xaf\xfc\x05\x8a\x07\x59\x39\x77\x42\xa8\xfb\xa4\x33\x29\x03\x02\x0f\x70\x27\x4b\xfe\xae\xe4\xfa\x0e\xd1\xfe\x55\x48\x28\x8f\x8a\xce\x23\x4b\x19\xec\xe5\x03\x0b\x94\x0a\x95\xdb\xf3\x06\x8c\x71\x15\x41\x05\xd2\x04\x56\xc7\xb5\xe1\x81\xc4\x85\x6f\x33\xf1\x48\xd8\xd4\x13\x4e\x27\x1f\xcf\x0a\x65\x96\xc2\x0d\x66\xea\xd0\x7d\x28\x75\x3c\x1b\x15\x6b\xc5\x53\x28\x58\xb7\xdb\xa1\x9a\x8b\xc5\xa0\x98\x39\x44\x5c\xd9\x33\x64\xc1\xc3\x32\x7f\x53\x57\x9b\x06\x77\x9c\x4a\x47\x47\xf2\x06\xa5\x6c\x2b\x20\xdd\x7e\xb7\x75\xc9\x04\xe6\x2c\x93\xed\xea\x2e\xd6\x24\x30\xc7\xa2\x64\xc6\xdc\x96\xfa\x14\xe1\x34\xa8\x6e\xf3\x38\x09\x2b\x38\x93\x6d\x5e\x02\x63\x19\xd9\xed\x1c\x17\xf2\x89\xad\x3b\xe7\x79\xa6\xe8\x04\x81\x2a\xf1\x49\xa1\xca\x60\x33\xd0\xe1\xe8\x26\x88\x76\x99\x6e\x9a\x77\x3a\xe1\xd3\xc4\x26\xe1\x6d\x6f\x91\x54\x8f\xe9\x8f\xc4\x38\x4c\x5b\xfb\xc6\xc9\xa8\xcd\x80\xd2\xb5\x71\x81\xe6\xd1\xd3\x2e\xc0\x7c\x0e\xdc\x16\xd7\x5e\x65\x8b\xec\x45\x4a\xd9\x4c\x2a\x4c\x60\x40\x28\xa1\x75\x70\x91\x32\xe7\x99\x76\xd0\x74\x82\xd4\xa3\x40\x29\xf8\xb2\xf4\xe9\x45\x92\xa0\x32\x79\xec\x7b\xb4\xa8\x37\xa9\x57\xcf\xc2\x6d\x20\x07\x45\x60\x81\xba\x7d\x4d\x93\x6a\xbb\x9e\x08\xf2\xb8\x3b\xda\x81\x96\x7d\x16\xce\x37\xb7\x31\x81\xfd\xa6\x0c\xcb\x1b\x4c\xdc\x34\x3f\xbf\xf7\xd7\x08\xcd\x2f\x06\xda\xbf\xe0\xa9\x44\x65\xf6\xaa\xd3\x0b\x76\xcc\x55\xad\x14\xee\x09\xa3\xd0\xc9\x8e\x57\xdf\x76\x5b\x70\x60\x2f\x51\x4f\x9e\xfd\xae\x3f\xf2\x8d\x54\x4f\xfb\x5c\x0c\xd1\x8c\xf1\x21\xcb\xec\x12\x49\x7c\xcd\x31\x35\x7c\xfc\x90\x47\xbb\xef\xb6\xa8\xbc\x1d\xbe\x43\x79\xe8\x47\x18\x3f\xc4\x75\xd1\x0c\xd1\x1a\x4b\x4c\x6b\xa1\x16\xd5\xea\x07\x65\x08\xee\xb8\x98\x2b\xd2\xf3\x79\xcb\xa3\xdb\x0b\x22\xca\x64\x0b\x73\x8d\xd5\x31\x2d\x3e\x12\x66\xba\x1b\xc5\x71\xdd\xc5\x8d\x70\x6a\x94\xae\x2d\x26\x50\x89\x46\xe1\x76\x74\x7f\xea\xf5\x7c\x37\x91\x74\xde\x6d\xae\xbf\x4c\x58\x87\x15\xd0\xcc\x33\xbf\x14\xa6\xb3\xad\x3d\x31\xcb\x97\x0d\x49\x3e\x8b\x7e\x57\xb3\xb8\xd4\x7d\xf8\x62\x16\x60\x03\x6a\x1f\xd5\xdf\x7c\x27\x24\xfd\x2a\x54\xe2\xe2\xf1\x30\x40\x35\x0c\xe2\x99\xba\x79\xe4\x9e\x59\xcc\xaa\xdb\xb8\xb7\xfe\x88\xe9\x87\x47\x26\x4a\xb8\x83\x79\xcb\x9a\x65\xe0\xf5\x7a\xae\x7e\x55\xdc\x97\x02\x0f\x50\xc6\x64\xb3\x7f\x02\xb7\x3d\x99\x62\xc7\x07\x34\xd1\xae\x63\xa7\x9d\x97\xda\x9d\xb0\x3e\x2a\x61\xb2\x99\xfa\x93\x03\xf7\xd1\xcb\xd1\xbf\x03\x00\x00\xff\xff\xcc\xb2\x2a\xb9\xc4\x27\x00\x00")

func coredns116JsonBytes() ([]byte, error) {
	return bindataRead(
		_coredns116Json,
		"coredns-1.16.json",
	)
}

func coredns116Json() (*asset, error) {
	bytes, err := coredns116JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "coredns-1.16.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xb2, 0x4a, 0x2d, 0xc0, 0xc0, 0x44, 0x91, 0xb8, 0x4a, 0x68, 0xe9, 0x44, 0x50, 0x1, 0x27, 0x85, 0x8b, 0xd7, 0x19, 0x2a, 0x90, 0xbe, 0x39, 0x46, 0xc, 0xc8, 0xb8, 0x3a, 0x93, 0xd7, 0x45, 0xf0}}
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
	"aws-node.yaml":     awsNodeYaml,
	"coredns-1.14.json": coredns114Json,
	"coredns-1.15.json": coredns115Json,
	"coredns-1.16.json": coredns116Json,
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
	"aws-node.yaml":     &bintree{awsNodeYaml, map[string]*bintree{}},
	"coredns-1.14.json": &bintree{coredns114Json, map[string]*bintree{}},
	"coredns-1.15.json": &bintree{coredns115Json, map[string]*bintree{}},
	"coredns-1.16.json": &bintree{coredns116Json, map[string]*bintree{}},
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
