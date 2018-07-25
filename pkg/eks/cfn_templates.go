// Code generated by go-bindata. DO NOT EDIT.
// sources:
// assets/1.10.3/2018-07-18/amazon-eks-cluster.yaml
// assets/1.10.3/2018-07-18/amazon-eks-nodegroup.yaml
// assets/1.10.3/2018-07-18/amazon-eks-service-role.yaml
// assets/1.10.3/2018-07-18/amazon-eks-vpc-sample.yaml

package eks

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

var _amazonEksClusterYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x93\x41\x4f\xe3\x30\x10\x85\xef\xf9\x15\x03\x17\x4e\x41\x85\x1b\xd6\x6a\xa5\xa8\x74\x51\xd5\x55\x17\x35\x08\xce\xae\x3b\x21\xd6\x3a\x76\x34\x1e\x53\xb1\xbf\x7e\x65\x27\x29\xc9\x92\x72\xd8\x5b\x9b\x19\xbf\xf7\x79\xde\x38\xcf\xf3\xac\x78\x29\x9f\xb0\x69\x8d\x64\xfc\xe1\xa8\x91\xfc\x8c\xe4\xb5\xb3\x02\xae\x6e\x17\x37\x8b\x7c\x71\x97\x2f\xee\xae\xb2\x7b\xf4\x8a\x74\xcb\x5d\xa5\x68\xe4\x1f\x67\x61\xb5\x29\x21\x87\xa5\xb3\x4c\xce\xc0\xa3\x91\x16\xaf\xb2\xec\x51\x92\x6c\x90\x91\xbc\xc8\x00\x96\x26\x78\x46\xda\xca\x06\xe3\x5f\x80\x89\x52\x01\x56\x36\x08\x95\x23\xe0\x1a\x93\xa0\xea\x0e\x5c\xa7\xe6\xa7\xf7\x16\x05\x94\x4c\xda\xbe\x66\x51\xad\xf3\x4a\x56\x25\xaa\x40\x9a\xdf\x1f\xc8\x85\xd6\xcf\x8a\x1b\xed\x19\x5c\x05\xbe\x6f\x85\xd7\xd8\x0b\xeb\x7b\x0f\xec\x40\x32\x4b\x55\xc7\x5f\xc9\x7b\xbb\xf6\xc0\xb5\x64\x90\xc6\xb8\x23\x30\xc9\xaa\xd2\xea\x54\x8e\x68\xfd\x4d\xdb\x68\x3f\x06\xfc\xa9\x3d\x7f\x2b\x5e\x4a\x21\x56\xcb\x5b\x21\x26\x64\x42\xac\x0f\xdf\x23\x7b\x19\xf6\x16\xf9\x0c\xa8\xd7\x1d\x68\xea\x49\x80\xda\xc2\xb1\xd6\xaa\x86\x77\x17\xe0\xa8\x8d\x81\x3d\x82\xdc\x1b\x8c\x44\x46\x06\xab\x6a\x38\x3a\xfa\x8d\x04\xd6\x1d\xd0\x7f\xc5\x93\x54\x3f\x40\x90\xde\xb4\xc2\x9d\x33\x58\xec\xb6\x73\x3c\xbb\x6d\x84\x89\xd7\xf6\x5d\x2f\x90\x33\xd8\xf3\xc4\x49\x04\x8f\xdd\x08\x15\x83\xb3\x11\x91\x60\x8f\xb5\x34\xd5\x7c\x6e\x9b\xb0\x47\xb2\xc8\xe8\x87\xf5\xfa\xec\xfa\x54\xe3\xa8\x0f\xde\xba\xc6\x01\xe4\xeb\xd5\xe8\xc4\x2a\x19\x0c\x0b\xb8\xbc\xb9\xbe\x59\x5c\xa6\x6f\x45\x8c\x12\x0f\xcf\xd2\x04\xec\x27\x9f\x0f\xf5\x6c\x87\xde\x05\x52\x5d\x61\xb5\x29\xfb\x55\x15\x23\xf9\xcb\x6e\x88\x9b\x52\x88\xbe\xda\xe9\x3e\x92\x6b\x91\x58\x0f\xa2\x00\x69\xc1\xe1\x62\x87\xd5\x78\xe5\xfb\xe2\xc9\xe9\xb9\x55\x4b\x67\x2b\xfd\x3a\x1c\x8b\x69\x8c\x96\x65\x7d\xf0\x83\xc8\xd9\x4d\xff\x38\x99\x62\xfd\x38\xd2\x2f\xd8\xe0\x19\xe3\x25\x3b\xd4\x26\x99\xf7\x2d\xa7\xa7\x9e\x5a\x3e\x65\x94\x65\xbf\x02\xb7\x81\xff\x99\xcf\xb2\x98\x09\x6f\x19\xa7\x51\x69\x25\x19\xa1\x08\x5c\xbb\xf4\xdc\xee\x25\xcb\xf3\x8f\x3b\xa5\x22\xe0\xe2\x01\xb9\x60\x1e\x39\x5c\x8f\xd4\x4e\x62\x51\x2b\x9b\x80\xac\xec\xa1\x75\xda\xf2\x0c\xce\x50\x9a\x9a\x7f\x7e\xbe\xe7\x11\x06\x85\xa9\xe5\xfc\x7b\x59\x1f\xd0\x46\x5e\xa4\xff\xb9\x6c\x41\x36\xfb\x1b\x00\x00\xff\xff\xa5\xb4\x23\x1c\x8a\x05\x00\x00")

func amazonEksClusterYamlBytes() ([]byte, error) {
	return bindataRead(
		_amazonEksClusterYaml,
		"amazon-eks-cluster.yaml",
	)
}

func amazonEksClusterYaml() (*asset, error) {
	bytes, err := amazonEksClusterYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "amazon-eks-cluster.yaml", size: 1418, mode: os.FileMode(420), modTime: time.Unix(1533210802, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xb6, 0x63, 0xfa, 0xca, 0x5b, 0xc1, 0xe3, 0x56, 0x74, 0x8b, 0xe1, 0xc0, 0x73, 0x44, 0x50, 0xba, 0x4d, 0xe3, 0x28, 0xf1, 0xda, 0xa3, 0x8d, 0x53, 0x29, 0xe5, 0x98, 0x92, 0x18, 0xa5, 0x67, 0xec}}
	return a, nil
}

var _amazonEksNodegroupYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x3a\x6b\x73\xdb\xb6\xb2\xdf\xf5\x2b\xb6\xba\x9e\xd1\xbd\x77\x44\x52\x94\xa5\x3c\xd8\x9b\xde\x51\x65\x25\xd5\x49\x6c\x6b\x2c\x25\x3d\x3d\x69\xc6\x03\x91\x2b\x19\x35\x09\xb0\x00\x18\xdb\x4d\x73\x7e\xfb\x19\x80\xa4\xc4\x87\x28\x3a\x69\x66\x9a\x0f\x91\x00\xec\x7b\x17\xbb\x8b\x95\x2d\xcb\xea\x4c\x7e\x5e\xae\x30\x8a\x43\xa2\xf0\x25\x17\x11\x51\xef\x50\x48\xca\x99\x07\xbd\xe1\xc0\x1d\x58\x83\xe7\xd6\xe0\x79\xaf\x73\x86\xd2\x17\x34\x56\xe9\xc9\x24\x22\x7f\x70\x06\xb3\xd7\x4b\xb0\xe0\x82\x07\x08\xaf\x04\x4f\xe2\x5e\xa7\xb3\x20\x82\x44\xa8\x50\x48\xaf\xd3\x01\x78\x8d\x0f\x17\x24\x42\xaf\x03\x00\x50\xa2\xb1\xba\x41\x98\x4d\x87\x1a\x02\x16\x84\x0a\x50\x1c\x48\x18\xf2\x3b\x58\x2e\x7f\x02\xe2\xfb\x28\xa5\xde\x53\x37\x08\x94\x49\x45\x98\x8f\xd2\x90\x59\x3d\xc4\xe8\xc1\xe4\xe7\xa5\xe7\xcd\xa6\x43\xcf\x7b\x8d\x0f\x9a\x80\xf9\xa2\x99\x69\xbe\x5a\xa6\x79\x44\xb6\x38\x0f\xbc\x83\x48\xe6\xd0\xf3\xe6\x41\x5d\xb2\xc9\xf9\x1c\x68\x00\x1b\x2e\x0c\x73\xa6\xd5\xdb\x49\x60\xef\xa8\x67\x3b\x86\x70\x9d\x88\x56\x2d\x47\x02\xf5\x10\x63\x03\xbd\x82\x70\x4b\x25\x28\xdb\x66\xa4\x36\x24\x09\x95\x07\x6a\x68\x47\x18\xd0\x24\x32\xdb\x13\x6d\x1f\x0c\xde\x91\x30\x41\x99\x32\xb5\x34\x88\x8c\x48\x18\xee\x97\x05\x0c\xb3\x0e\x89\xd8\xe2\x7e\x79\x5f\x59\x0f\x4b\x1b\xd1\x69\x19\x3f\x3a\xb5\x2b\xc7\x55\xf0\x0a\xfe\xc8\xae\x2c\xab\xc7\x35\xf8\x51\x75\xc3\x1d\x94\x77\xc6\x76\x65\x59\x3d\x1e\x56\x37\x46\xd5\x0d\xb7\x06\x32\x2c\xc3\xf8\x65\xb9\xfd\x8a\xdc\x7e\x55\x6e\xbf\x2a\xb7\x3f\xb2\x9f\x95\x37\xca\x62\xfb\x15\xb1\xfd\xaa\xd8\x7e\x55\x6c\x7f\x6c\x3f\xaf\x6e\xb8\x65\x1e\xb4\xec\x1d\x5a\xf1\x0e\xad\x7a\x87\x9e\x56\x78\xd0\x53\xbb\x46\xd1\x7d\x52\xda\x11\x15\xa2\xa2\x4a\x54\x54\x89\x8a\x2a\x51\x51\x36\xae\xa8\x18\x57\x54\x8d\x2b\xaa\xc6\x15\x55\xe3\x8a\x51\x55\xcc\x7b\xf7\xc0\xce\x69\x99\x6e\x5c\x89\xfe\x78\x58\x21\x1b\x0f\xab\x44\xe2\xaa\xb6\x71\x55\xb9\xb8\x62\xb1\x29\x67\x52\x09\x42\x99\x2a\x65\x84\x28\x91\x0a\xd6\x08\x04\x3e\x92\x90\x06\xf5\x0c\x91\x67\x96\x49\xa2\xf8\xd2\x27\x21\x65\x5b\x93\x55\xcf\x29\x5b\xd2\x3f\xb0\x98\xc7\x2e\x92\x68\x8d\xa2\x9e\x75\xce\x29\xa3\x51\x12\x81\xa4\x7f\x20\xf0\x4d\x21\x35\xc3\x64\xf9\xca\x2e\xe7\x16\xb7\x91\x21\xb9\x7f\x34\x43\x72\xff\x78\x86\xa7\x9a\xe1\x34\x4c\xa4\x42\x71\xa4\x2a\xf8\x29\x04\x30\x12\x21\xc4\x82\x7f\xa4\x01\x06\x70\x77\x83\xcc\xe4\xcf\xfc\xf8\x8e\x48\xf0\x05\x12\x85\x81\x0d\x30\xdf\x00\x55\x40\x25\x50\xe6\x73\x21\xd0\x57\x7d\x93\x69\x25\xdc\xd1\x30\x04\xc6\x53\xdb\xaf\x43\xd4\x45\xe5\x37\x4e\x4b\xc4\xec\x7a\x1a\xce\x6c\x63\x74\x69\x10\xf6\x2d\xa3\xbf\x27\x08\x34\x40\xa6\xe8\x86\xa2\xd8\xe5\xf8\xbd\x19\x0e\x53\xce\x8c\x30\xe5\x4c\x09\x1e\x2e\x42\xc2\x70\x89\x7e\x22\xa8\x7a\x30\x58\x0d\x96\x91\x19\x0c\x6c\x8d\x85\xf9\xa6\x64\x10\x3f\xa5\x06\xb1\x26\x67\x1f\xac\x7a\x65\x26\xba\xfa\x75\x00\xde\xc5\x7e\x5e\x25\x6b\x1c\xdf\x2d\xa6\x39\x9b\x3b\x2e\x6e\x51\xb4\xd4\xe2\x77\x8b\x69\x4e\x76\x99\xac\x19\x2a\xd9\xa4\x4a\x7a\xaa\xdd\x2a\x72\xda\x12\x7c\xc2\xb4\x9b\x72\xb7\x16\x58\xbc\xa1\x52\xfd\x5f\x41\x11\x83\xae\x59\xfd\xa0\x79\x9d\x13\x46\xb6\x18\x2c\x78\x48\xfd\x87\x89\x60\xb2\x18\xbb\x53\x1e\x45\xe4\x0c\x43\x1a\x51\x85\x81\x26\xd4\xe9\x9c\x93\x38\xa6\x6c\x6b\xe0\xce\xc9\xfd\x82\x07\x72\x81\x42\xbb\x2d\xc5\xcc\xab\x41\xba\xda\xc1\x78\x30\x7c\x9e\x9f\xdf\x1f\x06\x18\x3f\xcb\x01\x86\xad\x10\xa3\x06\x88\xe1\xe9\x28\x07\x79\xd6\x06\x32\x6e\x11\x74\xdc\x26\xe8\xb8\x55\xd0\x71\xbb\xa0\xbb\x4a\x75\x0c\xc4\x6d\x52\xe6\xe9\xe9\x53\xb3\x95\x57\xb3\x26\x65\x76\xe5\xad\x49\xd4\x7d\xbd\x3b\x02\xd1\xa6\xcc\xbe\x24\x1e\x03\xc9\x53\x7e\xa3\x32\xbb\x46\xaa\x06\xe1\x0e\x73\x80\xa3\xda\x46\x6d\xda\x46\xcd\xda\xba\x6e\x06\xd2\x18\xc7\x83\xfc\xbc\x85\x47\x6b\x1c\x47\xed\x71\x5c\x68\xe9\x9a\x61\x5a\x02\x39\x6a\x0b\xe4\xa8\x35\x90\xa3\xf6\x40\x2e\x74\x8a\xc7\x60\x86\x4d\x74\x72\xe7\xef\x1a\x8d\x26\x59\xf6\x9d\x47\x23\x9f\x42\x2f\xd2\x0c\xd3\x1a\xf0\x71\x7b\x34\xc7\x47\xa2\x39\x87\x11\x6d\xc1\x28\x5a\x25\x11\xed\x57\x4f\xb4\x0b\x2b\x5a\x32\xb3\x68\x8b\x68\xd1\x1a\xd1\xa2\x3d\xa2\x45\x7b\x66\x2e\x74\xa8\x8d\x61\x92\xbf\xdd\x6a\x00\xcf\xf2\xe3\xa6\x14\xb2\xc3\x3f\xcc\xe0\x74\x9c\x9f\x37\x48\x30\x1a\xe5\x00\x4d\xc6\xc8\x20\x0a\x6d\x75\xa3\xaa\x85\x46\xfb\x20\x4c\xe7\x1c\x15\x09\x88\x22\xfa\xd4\x14\xf1\x69\xc8\x93\x20\x9d\x37\xe8\x96\xc0\x9b\x33\x85\x62\x43\xfc\x0c\x7f\x37\x45\x30\x8d\x8a\xcc\x89\x5a\xd9\x27\xc0\x1b\xb2\xc6\x9d\xd5\xf4\xbf\x20\x6f\x32\xbb\xb3\xd7\xcb\xbc\xbf\xea\xee\x00\x8a\x63\x89\x3d\x92\x55\xec\x46\x0f\xed\x37\x36\x68\x5f\x22\xd0\xcf\x69\xe3\x64\x7a\xc2\x29\x67\x1b\xba\x4d\x84\x51\xbb\x5d\xbc\x52\x07\x5a\x3b\x69\x78\x28\xb4\xc3\xa5\xfd\x7d\x0d\xae\x38\xd2\xa8\x1f\xa6\xd3\x94\xd2\x7e\x3e\x70\xf9\x0a\x63\xa0\xd2\x3d\xdf\x97\xda\xc3\xb4\xaa\xa5\x9d\xac\xcb\xec\x74\xae\x50\xf2\x44\xf8\x98\xce\x9d\x8a\xea\x2c\x04\xdf\xd0\x10\x6b\x73\xa0\xf9\xe4\x5c\x07\x5e\x09\x28\x8d\x3e\xc1\x63\x14\x8a\xe2\x8e\xfb\x82\xa8\x1b\x0f\xba\x4e\x2e\xe1\x15\x0f\xf7\x87\x16\x7c\x77\x85\x9b\x12\x4b\x7d\x5e\x15\x43\xef\x1d\x96\xc1\x40\x1f\x66\x3c\x91\x32\x89\x0c\x6a\xda\xdb\x9e\x71\x3f\x89\x90\xa9\xbd\x59\x4a\xe3\xba\xa1\xe5\x0e\x2c\xf7\x69\x6f\x77\xba\x54\x44\x61\x19\xc1\x82\xd9\x66\x83\xbe\xf2\xd2\x71\x52\xc1\x9a\x0b\x41\x99\x4f\x63\x52\x72\x1d\xc0\x12\xc5\x47\xea\x63\x79\xd3\x02\xf4\x87\x36\x31\x73\x40\x72\x27\x6d\x9f\x47\x85\xf3\x89\x6f\xae\x75\xc9\x53\x52\x49\x6f\xaf\x4e\x83\x61\xeb\x8d\x7c\x6a\xdc\xda\x7e\x6e\xdd\x03\xcf\xa6\xa3\x8f\x9e\x06\x3b\x9b\xb3\xd2\x23\x65\x59\x7e\x6b\xe9\x97\x1d\x31\x0f\x49\xfd\xa6\x2c\xbf\x1e\x33\x1a\x85\x77\x94\xfe\x67\xe4\x2e\x06\xec\x8a\x6c\x0b\x31\xf3\x1a\x1f\x3c\xf8\x6e\x99\xac\xa1\x7b\x9b\xac\x51\x30\x54\x28\x6d\xca\x9d\x8c\xa8\x73\xf2\xa9\x90\x9c\x3e\xef\xef\x86\x99\xfe\x79\xd0\xe3\x77\x0c\x83\xde\x41\x3b\xcc\xd9\x56\xa0\x94\x8f\x30\x47\x06\x69\x00\xcf\x30\x46\x16\xc8\x4b\xe6\xd5\x29\x36\x98\xad\x3c\x37\x35\xd3\x5b\x33\xde\x54\x1c\x7c\x1e\x45\x09\xa3\x3e\x51\x08\x77\x54\xdd\x00\x12\xff\x06\xb8\xba\xd9\xd9\x2b\x15\x20\xf0\xf6\xd7\xe7\x50\x8a\x5d\x9a\x5b\x5d\x16\xba\x0d\x67\x1e\x2f\x04\x57\xdc\xe7\xa1\x07\x3d\xcb\xcd\xaf\xc3\x4b\xc1\xa3\x05\x17\xca\x83\x41\xee\x12\x9e\xae\x9f\x8c\xc7\xa7\xe3\x83\xa6\xd4\x38\xc5\x2a\xf0\x77\x9a\x36\x7b\x81\xbf\x4e\xd6\x18\xea\xb7\x33\x61\x01\xc4\x3c\x30\x63\x72\x81\x3e\xd2\x8f\x58\x30\x3b\xe5\x0c\x36\x82\x47\xcd\x33\x82\x6f\xe6\x87\x47\x56\xcb\xa2\x5b\x94\x1f\xd7\xbc\xe2\x0e\x86\xe3\x26\xc7\x14\x69\xcf\x8c\x61\x57\xfc\x6b\xee\xff\xec\x5b\x3b\xa5\xd1\xba\x07\x2f\x41\xd9\x85\x3b\x0f\x1e\xf4\xc4\x23\xad\x7a\x86\x52\x51\x66\xfc\xfd\xf5\xb7\xe4\x8b\xdd\xd1\x22\xdb\xdf\x79\x4d\xf2\x2b\x51\x33\x7e\xd1\x55\x93\xc5\xdc\x14\xb6\x86\x6c\xf4\x48\xdb\xff\xd5\xe4\xb4\x37\x7b\x6e\xe1\xd1\xe8\xb4\xe6\x08\xbd\x57\x1c\x47\xd6\xac\x5a\x68\xf0\x4a\x8b\x16\xc3\x51\x81\xc1\x94\xc4\xc4\xa7\xea\xa1\x20\xf4\xf1\x76\xf1\x0d\x49\x98\x7f\x53\xea\xdb\xcc\x78\x74\x8f\x5f\x84\xc8\xcb\x7a\x36\xc2\x3e\xc2\xa4\xd4\xbb\xe6\x13\xe8\xc7\x0a\xf5\x6e\x31\xfd\x17\x67\x38\xdf\x4d\x62\x2b\x65\x38\xef\x12\x1b\x0b\x71\xa9\xbf\xce\x2a\x6c\x5a\x9c\xcb\x65\xd8\x3a\xf9\x54\xea\xc9\x3f\x5b\x7a\x59\x68\x5c\x05\x8f\xc9\x96\x28\x9c\xa8\xd4\x0c\x1e\xf4\x94\x48\xb0\x77\xa0\xee\xf7\x1e\x53\xf7\x7b\x4d\x75\xff\x71\x1c\xdf\xc6\x01\x51\x59\xef\xb8\xeb\x29\xf7\xa6\xbc\xe2\xa1\xfe\x48\xa1\xf6\x36\x3b\xa7\x2c\xef\x5a\xe5\x9c\xe5\xfd\x1f\xf4\xdc\x3d\xdb\x73\x72\xff\x23\x51\xfe\x4d\xea\x26\x7d\x92\x45\x68\xd1\xfb\xc7\x03\xf5\x40\x24\x35\x77\xc1\xdc\xa7\x5a\x91\x64\x1d\x52\x7f\x1e\x4f\x82\xc0\xa4\x97\xb2\x71\xe7\x24\xaa\xf6\xfc\xf5\xd6\xbc\xd8\xe8\x03\xe4\xbf\x11\x17\xe0\x4a\xef\x9c\xd2\xaf\xbc\x75\x62\x85\x97\x52\xfe\x4b\x77\x0a\x54\x7e\x19\x95\xd2\xc0\xa1\x77\xc3\xa1\x3c\xf1\x56\xa2\x38\xcb\xde\xcc\x59\x46\x60\x9e\xf7\x23\x91\xf8\x64\x54\xec\xac\xf5\xee\x3f\x38\x65\x1e\xbc\x2f\x75\xe8\xdd\x6e\xbf\xb4\x2e\x9f\x02\x74\xff\xeb\x3b\x67\x4d\x99\xb3\x26\xf2\x06\xac\x7b\xfc\x95\x55\x10\x00\xba\xd3\xc9\xf5\x74\x76\xb5\x9a\xbf\x9c\x4f\x27\xab\xd9\xf5\xd9\xfc\x6a\x36\x5d\x5d\x5e\xfd\xf2\xc2\x41\xe5\x3b\xfb\xf8\x75\xe2\x5b\xda\xed\x43\xf7\x11\x34\x5e\xce\xdf\xcc\xae\x17\x93\xd5\x4f\x2f\x4e\x9a\xa8\x3b\x3e\xb1\x7d\xa1\x9a\x08\x9e\x5f\x9e\xcd\xde\xec\xa1\x53\x62\xff\x76\x6c\x72\x27\x1d\xbc\x95\xc7\xd1\x0a\xec\x0f\xd1\xd1\x04\xac\xe1\xc0\x7d\x6a\xb9\xae\x35\x70\x6d\xc6\x45\x44\x42\xfb\x37\xc9\x59\x13\xdd\xe8\x36\xa0\x02\xac\x18\x1a\xf5\x69\xc7\x3c\x24\x4a\x13\x96\x9f\x88\x10\x2c\x9e\x23\xed\x14\x82\x1b\xa5\x62\xe9\x39\x8e\x3c\xb5\x12\x69\xdd\xa1\x54\x56\xe5\x85\xe6\xa4\x2b\x0b\x6f\xa5\xe3\xda\xee\xc0\x3e\x75\x86\x03\xf7\x99\x35\x78\x62\x0d\xc6\x5f\xa1\x3a\xb9\x93\xba\xe7\x31\xf7\x17\x81\x04\x81\x15\xf1\x00\x43\xb0\x2c\x99\xe6\x8c\x6c\x6d\xae\xa2\xe3\xd4\x24\xde\xc3\x99\xdf\xfa\x8e\x38\x4f\x73\xc2\x5b\x09\x81\x29\xf7\x6b\xb4\xf2\x3a\x6e\x59\x02\xb7\x94\xb3\x17\xdd\x3e\x7c\x82\x2b\xdc\x78\xd0\x35\xa9\xe6\xca\x6c\x77\xe1\x73\xbf\x0b\x96\xa1\x5f\x00\x29\x64\xd8\x0c\xe0\xf7\x04\xc5\x03\xf4\xf2\xdf\x03\x3f\xf9\x3a\x03\x6d\x4c\xf3\x30\x49\xd4\x0d\xd7\xb7\xd3\xdc\x46\x38\x74\x62\x07\x44\x91\x3e\x20\x0b\x62\x4e\x99\xf2\x76\xdf\x3e\xf7\xe0\x07\x70\x54\x14\x3b\xb9\xe4\xd7\x19\x8b\x6b\x81\x32\x09\xd5\x51\xfb\xfa\x44\xb5\x23\xc3\x9f\xb0\x15\x18\x1f\x94\x4b\x4b\x0c\x7f\x02\xb9\xbb\x85\xde\xa7\x58\x50\xa6\xe0\x64\xf8\xb9\x07\x7f\x82\xc4\x00\x7a\xd2\x79\xdf\xff\xb5\xfb\xc1\x71\xb6\x7a\x6b\x6d\x52\x0b\x58\x01\xfc\x00\xb5\x60\xde\x39\xad\xf1\x7a\x4d\x96\xab\xd9\xd5\xf5\xec\xe2\x6c\x71\x39\xbf\x58\xbd\x38\xf9\xef\x2f\x92\x3e\x37\xd8\xe3\xa4\xfd\x9f\x26\x29\xe6\x17\xab\xd9\xd5\xc5\xe4\xcd\xf5\x7c\xa1\x25\x30\x57\x45\x9a\x9b\xe1\x39\x8e\xfb\xe4\xb9\x3d\x1c\x8f\xec\xec\xd3\x09\x89\x42\xa9\x9c\x08\x15\xb1\xb4\x03\x9d\x90\xfb\x24\xb4\x68\xfc\x71\xd4\xc8\x40\x4b\x62\x51\x90\xfd\x8a\xbe\xfd\x93\xea\xc6\x16\x9c\x8f\x44\x38\x21\x5d\x9b\x44\x19\xa2\x32\x9f\xe9\x7d\x69\x25\x3f\x7d\xf3\xd6\x90\xbb\x98\x9c\xcf\xfa\x4d\x81\x0b\xdd\xbf\xc8\xe5\x6a\xf6\x6a\x7e\x79\xd1\x6f\xbe\x3b\x29\x07\x9d\xec\xe5\x83\x54\x18\x05\xd9\x67\xce\xcc\xce\x2e\xf0\x23\xcc\xf5\xcf\xeb\xc5\xe5\xd9\x32\xe5\xd5\xd5\x45\xeb\x25\x65\xc1\x9c\x9d\x93\xb8\xeb\xc1\xfb\xca\xaf\xb0\x3b\x81\xaa\xd5\x56\xcb\x94\x81\xc2\x87\x6f\x2b\x60\xbb\x3f\xbf\x01\x9b\x42\x80\xf6\x4f\x8a\x8b\xbf\x44\xfe\xec\x62\x79\x9d\x47\xcc\x7c\xf1\xc2\x1d\xd8\xee\x60\x60\xeb\xff\x9b\x30\xe8\x06\xde\xbf\x87\xa2\x04\xf0\xe2\x05\xb8\x03\xfb\x7f\xe1\xc3\x07\xf8\x5e\xbf\x95\x18\x54\xc9\x3e\x1d\xda\xc3\x94\xec\xf7\xb0\x69\x2c\xf8\x3b\x5d\xcb\xe8\xfd\x93\xca\x7a\x0b\xdf\xc4\xa2\xc5\x2c\x35\x79\xbb\xfa\xe9\xf2\x6a\xbe\xfa\xc5\xe4\xab\x7e\x63\x12\x6b\xb9\x36\xd0\x7e\x3b\xe7\xb3\x8b\xd5\xf5\x74\xd2\xce\xe7\x51\x3a\x36\x32\x34\xe0\xbe\x0a\x21\x20\x18\x71\x66\x09\x0c\x39\x09\x1a\x6d\xb2\x03\x17\x28\x15\x11\x0a\x32\x46\x4d\x08\x0e\x8f\x95\xa3\xfb\x26\xdd\x08\xfa\x1b\x66\x49\xba\x65\x24\x04\x0b\xe1\xe4\xff\xa1\x0e\xbf\xfb\x66\x59\x52\x11\xff\x16\xaa\x99\x63\xa9\x77\x75\x76\xd2\xc9\xe3\x08\xb6\xc8\xa6\xf5\xfb\x47\xed\x71\x6e\x69\x91\xaf\xb1\x2b\x26\xaa\x5f\x0b\xbf\x1f\xe8\x7f\x1f\x3a\xc5\xef\x9d\xcb\x44\xc5\x49\xfa\xe7\x28\x87\xa7\xf2\xb5\x3f\x51\x29\xfd\xa1\x26\x88\x7c\x68\x9d\xbf\x0e\x5f\xa1\x9a\x28\x55\x23\x66\x4f\x04\xeb\x74\xfe\x13\x00\x00\xff\xff\xf1\x7e\x73\x3e\x5a\x2b\x00\x00")

func amazonEksNodegroupYamlBytes() ([]byte, error) {
	return bindataRead(
		_amazonEksNodegroupYaml,
		"amazon-eks-nodegroup.yaml",
	)
}

func amazonEksNodegroupYaml() (*asset, error) {
	bytes, err := amazonEksNodegroupYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "amazon-eks-nodegroup.yaml", size: 11098, mode: os.FileMode(420), modTime: time.Unix(1533210802, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x98, 0xc8, 0xca, 0x11, 0x86, 0x93, 0xfd, 0x9c, 0x58, 0xb8, 0x3e, 0x36, 0x35, 0x66, 0xbd, 0x6d, 0x51, 0xf9, 0x25, 0x61, 0x51, 0xe8, 0xc0, 0x38, 0xea, 0x60, 0x52, 0xce, 0xdd, 0xb4, 0x22, 0x7d}}
	return a, nil
}

var _amazonEksServiceRoleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x53\x41\x6f\xdb\x3c\x0c\xbd\xeb\x57\xb0\xc5\x07\x18\x08\xe0\x7c\x49\x2e\x43\x75\x73\xdb\xb4\x18\xd2\x66\x45\x1c\xb4\x87\x61\x07\x46\x61\x5a\x21\xb2\x65\x48\xd4\xb2\x6c\xd8\x7f\x1f\x24\x3b\x59\x8c\x02\x19\x76\x98\x4f\x06\xdf\x23\x45\xbe\x47\xe6\x79\x2e\x8a\x97\x72\x49\x55\x63\x90\xe9\xce\xba\x0a\xf9\x99\x9c\xd7\xb6\x96\x90\x4d\x46\xe3\x51\x3e\xba\xca\x47\x57\x99\xb8\x25\xaf\x9c\x6e\xb8\x45\x8a\x0a\xbf\xdb\x1a\xa6\xb3\x12\x4a\x72\x5f\xb5\x22\x58\x58\x43\x99\x10\x62\x41\xde\x06\xa7\xc8\x4b\x21\x00\x8a\x97\xb2\x23\x44\xfc\xce\xba\x36\x73\x3a\x2b\xa5\x00\x00\x58\xee\x1b\x92\x91\x25\xe5\xc7\xe2\x51\xca\xc8\x4a\xc0\x93\xb3\x0d\x39\xd6\xb1\x0e\xa4\xaf\xf0\x3e\x54\xa9\xcc\x93\x35\x5a\xed\x6f\xad\x0a\x15\xd5\x7c\xc0\x01\x7a\x9d\x4f\xf2\xf1\x28\x1f\x7f\xc8\x8e\x68\xc9\xc8\xd4\x4f\xc8\x61\xba\xd9\x90\x62\x09\x85\x31\x76\x77\x8c\xc7\xe7\x75\xad\x74\x83\x46\x9e\x04\xe1\x30\x6c\x3f\x98\x03\x6d\xfd\x10\xd3\x60\xb8\xf3\x43\x65\xab\x13\xbc\x50\x49\x33\x71\xca\xf7\xec\xe5\xef\x71\x3a\xe8\x11\x6b\x7c\xa5\x75\x3b\x5c\xe1\x6a\x7f\xda\x27\xba\x5a\xe2\xce\x4b\x8d\x95\x4c\x3f\x4d\xa2\xfd\x7f\x94\xb3\xeb\xac\xcd\xfe\x9b\xc4\x1b\x13\x3c\x93\x3b\x26\x0a\x80\x39\xf1\xce\xba\xed\x83\xc5\xf5\x35\x1a\xac\xd5\x01\x8e\xfe\xe1\x79\x03\x4f\x1a\x78\x6f\x61\x0b\xce\xb1\x22\x09\x17\x65\x58\xc1\xe5\x7f\x3f\x52\x6a\xc9\xa8\xb6\x31\xfe\x33\x9f\x3f\x5c\x5f\x76\xf4\x28\x8e\x97\xf0\xf9\x62\x41\x9b\x33\x9b\xf4\xa5\x57\xfd\x9f\xae\xc5\x61\xb9\x25\x64\x83\xec\x0f\x26\x93\x41\xcf\x5a\x19\x8b\xeb\x55\x52\x51\xd7\xaf\x72\xd0\xa7\xa8\x89\xbc\x71\x84\x4c\x25\xa9\xe0\x34\xef\xef\x9d\x0d\xcd\x3b\x4e\x7b\x7c\x2b\x1a\x08\xf1\x29\x70\x13\xb8\x3d\xae\x28\x44\xe1\xba\x67\x7b\x07\xba\x7c\x23\x70\xd6\x10\xf0\x1b\x72\x3a\xd3\x9d\x36\x06\x82\x27\x60\x0b\x2a\x3d\x19\x15\x05\x77\xb8\x56\xd8\x58\x07\xb3\xb0\x22\x57\x13\x93\x07\xd5\xae\x85\x4f\xb5\x9f\xd1\x84\x68\xd9\x3d\x71\xc1\x7c\xc6\x8a\x61\xe1\xea\x94\x31\xfd\xd6\x58\x77\x94\xf4\xbc\xe3\xdd\x18\x97\x42\xfc\x0a\x00\x00\xff\xff\x46\xf1\xdd\x9c\x90\x04\x00\x00")

func amazonEksServiceRoleYamlBytes() ([]byte, error) {
	return bindataRead(
		_amazonEksServiceRoleYaml,
		"amazon-eks-service-role.yaml",
	)
}

func amazonEksServiceRoleYaml() (*asset, error) {
	bytes, err := amazonEksServiceRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "amazon-eks-service-role.yaml", size: 1168, mode: os.FileMode(420), modTime: time.Unix(1533210993, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xd3, 0x8, 0xc0, 0x76, 0x12, 0x30, 0x16, 0x74, 0x4, 0x59, 0x88, 0x45, 0x46, 0x1a, 0x3f, 0x6f, 0x5d, 0xd3, 0x7f, 0x50, 0x1, 0xe3, 0xfd, 0x9a, 0xe, 0x1f, 0xd1, 0x23, 0xed, 0x8, 0xd3, 0x82}}
	return a, nil
}

var _amazonEksVpcSampleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x57\x51\x6f\xdb\x36\x10\x7e\xd7\xaf\xb8\x1a\x03\xbc\x01\xb1\x2b\xc9\x43\x10\x13\xc3\x00\x4f\x49\xba\xac\x5d\x6b\x58\x86\x03\xac\xe8\x03\x2d\x9d\x6d\x22\x14\x29\x90\x54\x02\x6f\xd8\x7f\x1f\x44\xc9\x32\x6d\x4b\x5d\x8c\x06\x0d\xfc\x60\x88\x77\xe4\xf7\xdd\xc7\xbb\xd3\x69\x30\x18\x78\x93\xfb\x78\x8e\x59\xce\xa9\xc1\x5b\xa9\x32\x6a\x16\xa8\x34\x93\x82\x40\x3f\xf4\x03\x7f\xe0\x8f\x07\xfe\xb8\xef\x5d\xa3\x4e\x14\xcb\x4d\x65\x99\x64\xf4\x6f\x29\xe0\xe6\x7d\x0c\x31\xcd\x72\x8e\xb0\x98\x46\x7d\xcf\x9b\x52\x45\x33\x34\xa8\x34\xf1\x3c\x80\x45\x9e\xfc\xc6\x65\xf2\x40\x3c\x00\x80\xf9\x36\x47\x02\xb1\x51\x4c\xac\xed\xc2\x35\xae\x68\xc1\x0d\x81\x60\x1c\x0e\x83\xcb\xab\xa1\x3f\xf4\xdf\x06\x97\xb5\xcd\xc1\x9b\x6f\x10\xa2\xbb\xeb\x19\x28\x2a\xd6\x08\x2b\xa9\xc0\x6c\x2c\xe6\x10\xe6\x1b\xa6\x41\x6f\x64\xc1\x53\x58\x22\x50\x78\xa4\x9c\xa5\x90\x2b\xf6\x48\x0d\xc2\x8f\xb3\xdb\x08\x82\x71\x70\xf5\x93\x73\xc2\xb0\x24\x37\x79\xa4\x8c\xd3\x25\xe3\xcc\x6c\xff\x92\x02\xb5\xcb\xf2\x03\xd3\xe6\x97\xc9\x7d\x4c\xc8\x4d\x14\x12\x72\xec\x4b\xc8\x47\x9a\xe1\xaf\xe5\x31\x71\xb1\x14\x68\xfc\xe0\xac\x40\x2f\x7f\x2e\x23\xbd\x3a\x8d\x34\x62\xa9\xb2\x27\xd9\x20\xb5\x3d\x1b\xfc\x00\x9e\x98\xd9\x30\xb1\x8b\xda\xc1\x0d\xcf\xc2\x0d\xc2\xab\xb3\x80\xc3\x6e\xe0\xd1\x79\xc0\xe3\xf0\x2c\xe0\xd1\x09\xf0\x9f\x68\x68\x4a\x0d\x2d\x11\xed\xc5\x44\x5c\x16\x69\x95\xb2\xe5\x41\xe4\x4e\x18\x54\x2b\x9a\x60\xc5\xa9\x49\xc5\x77\x4a\x16\x79\x7d\xb9\x00\x83\xfa\x1f\xe0\x03\x5d\x22\x27\xcd\x23\x40\xba\x23\xdd\xbb\x97\xea\x01\x15\x7c\x44\xf3\x24\xd5\x03\x44\x52\xac\xd8\xba\x50\x16\xa8\xd7\xec\x70\x93\x7d\x7f\xca\xa0\xc9\xfa\x83\xc5\x83\x34\x69\xb3\x84\x9d\x96\x4a\x69\xcf\x9b\xa1\x96\x85\x4a\xaa\x44\x5d\x4c\x23\x57\xfb\x7d\xaa\x96\x62\x59\x76\x4a\xe6\xa8\x0c\xc3\x86\x5d\xa3\x34\x01\x78\x33\xc3\xd5\x31\xd1\x1b\x41\x97\x1c\xaf\x85\x8e\x8b\x3c\x97\xca\x10\x30\xaa\xc0\x63\xe3\xef\x52\x1b\x41\x33\xd4\x07\xe6\x39\x5d\xef\x25\x86\xf7\xb8\x25\x50\x56\x48\x13\xce\x82\xf2\x02\x09\xbc\x89\x8b\x25\xf4\x7f\xf8\xc7\xd2\x8d\x0d\x4d\x1e\x4a\xaf\x7f\x07\x55\xf7\x00\xb0\x77\x28\xd0\xbc\xa3\x06\x9f\xe8\xd6\x8d\xb0\xb7\x0f\xf1\xc8\xab\xe7\x55\x7a\xd4\x8f\x13\x63\x68\xb2\xc9\x50\x98\x8e\xed\x6d\xae\xbd\x0e\xcd\x8e\xa0\xee\x52\x52\x49\x77\xb4\x5e\x7b\x2f\xf2\xa4\xf1\xa8\xab\x65\x26\x0b\x83\xf3\x52\xbb\xf6\xeb\xda\xdb\x3b\x18\x1c\x9f\x79\x96\xde\xd3\x62\xc9\x59\x52\x27\x92\x3e\xf2\xaf\xb2\xbb\x7d\x4b\x43\x9d\xd4\x15\x9b\xa3\x48\xf5\x27\x41\x5a\x85\xee\x8e\xac\x23\xa8\x7d\xd4\x4d\x6c\x47\x42\xd8\x2e\x61\x98\xb0\x35\xe7\x64\xae\x3f\xb4\xbf\xb7\x7e\xed\xf5\x7f\x17\xe3\x74\xe8\xf6\x0b\xa8\xac\xd6\xe4\xb6\x18\x5b\x30\x32\xb3\x69\x54\x9f\x00\x7e\xd0\x11\xce\xc9\xeb\xa1\x11\xf5\x56\x10\x12\x23\xc7\xc4\xec\xd7\x06\xd0\xf7\xfb\xce\x93\xe5\x7d\xf2\x36\x3a\x29\xda\x66\xc7\x0c\x57\xa4\xb5\x9f\x54\xa9\x72\xe8\x77\x76\xca\xd8\x12\xed\x9d\x96\xe8\x0e\xb0\xe7\xbe\x7c\xbe\x5d\xd2\xf0\xa5\x24\x0d\x5e\x4a\xd2\xf0\x7b\x4b\x1a\xba\x92\x8e\xbe\x5d\xd2\xd1\x4b\x49\x1a\xbe\x94\xa4\xa3\xef\x2d\xe9\xc8\x95\x34\xd8\xb7\x96\x89\xd6\x32\x61\xd5\xc0\xf0\x15\x9d\x5b\x37\x74\xa8\x5a\xed\x68\xfa\xcf\x0e\xf3\x59\x8d\xce\xa9\xa4\x57\xe0\x18\x9e\xc9\x71\xf4\x0a\x1c\x47\xcf\xe5\x18\x49\x61\x94\xe4\x53\x4e\x05\xc6\x98\x14\x8a\x99\xad\x9d\xfc\x3a\xf8\xb9\x2e\x1d\x8c\xac\xed\x70\x56\xe5\x85\x36\xa8\x20\x91\x59\x56\x08\x96\xd8\x78\xec\xa4\x0a\x4f\xd5\xc8\x28\x64\xda\x14\xc5\xc9\x34\xf0\xa9\x30\x79\x61\xaa\x4f\xa3\x5d\xb8\x9a\x9c\x8e\xc4\x13\xce\xeb\x41\x58\x83\x33\x02\xbb\xe9\xff\x87\x64\x02\x3e\x43\xef\xa2\x77\x01\x9f\x0f\x33\xef\xe2\xf0\x92\x0f\x1f\x47\xf0\x05\xbe\x58\x7c\x57\x80\x36\x12\x3b\x07\x58\x97\x1e\xcd\x37\x57\xd2\x48\x60\xf5\x86\xbc\x14\xfc\x59\x82\x7c\x85\x7a\xe7\xe5\xed\xd8\x3a\x3d\xe3\xe4\xc3\x70\x31\x8d\xe0\x2e\x3d\x80\xd8\x29\xfe\x5f\x00\x00\x00\xff\xff\xc0\x32\x41\x5b\xe0\x0e\x00\x00")

func amazonEksVpcSampleYamlBytes() ([]byte, error) {
	return bindataRead(
		_amazonEksVpcSampleYaml,
		"amazon-eks-vpc-sample.yaml",
	)
}

func amazonEksVpcSampleYaml() (*asset, error) {
	bytes, err := amazonEksVpcSampleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "amazon-eks-vpc-sample.yaml", size: 3808, mode: os.FileMode(420), modTime: time.Unix(1533210922, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa9, 0x61, 0xc3, 0xc2, 0x8c, 0x17, 0x5a, 0x19, 0xf0, 0xe1, 0xc0, 0xc3, 0x6, 0xfb, 0x2d, 0x1e, 0xce, 0x11, 0xf3, 0xea, 0xdf, 0x89, 0x99, 0x37, 0x5d, 0xc1, 0x5d, 0x93, 0xd5, 0x9c, 0x4d, 0x14}}
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
	"amazon-eks-cluster.yaml": amazonEksClusterYaml,

	"amazon-eks-nodegroup.yaml": amazonEksNodegroupYaml,

	"amazon-eks-service-role.yaml": amazonEksServiceRoleYaml,

	"amazon-eks-vpc-sample.yaml": amazonEksVpcSampleYaml,
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
	"amazon-eks-cluster.yaml":      &bintree{amazonEksClusterYaml, map[string]*bintree{}},
	"amazon-eks-nodegroup.yaml":    &bintree{amazonEksNodegroupYaml, map[string]*bintree{}},
	"amazon-eks-service-role.yaml": &bintree{amazonEksServiceRoleYaml, map[string]*bintree{}},
	"amazon-eks-vpc-sample.yaml":   &bintree{amazonEksVpcSampleYaml, map[string]*bintree{}},
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
