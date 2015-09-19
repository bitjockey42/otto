// Code generated by go-bindata.
// sources:
// data/aws-vpc-public-private/deploy/cloud-init.sh.tpl
// data/aws-vpc-public-private/deploy/main.tf.tpl
// data/common/dev/Vagrantfile.tpl
// data/common/dev-dep/Vagrantfile.fragment.tpl
// DO NOT EDIT!

package dockerext

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path/filepath"
)

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
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
	name string
	size int64
	mode os.FileMode
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

var _dataAwsVpcPublicPrivateDeployCloudInitShTpl = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x91\x41\x6b\xe3\x30\x10\x85\xef\xfa\x15\x2f\xc9\x59\x11\xec\x6e\x2e\x21\xbb\x97\xcd\x2e\x14\x02\x85\x86\xd2\x43\x29\x41\x91\x27\xb1\xa8\x2d\x0b\x69\x9c\x34\xa4\xf9\xef\x95\x6c\xb7\x09\xf4\x66\x9e\xbe\xf7\xe6\xcd\x78\x32\x52\x5b\xeb\xd4\x56\xc7\x52\x44\x62\x48\x12\x62\x82\x3b\x17\x59\x57\x15\xcc\xe3\xc3\x0a\x76\x87\x23\xa1\xd4\x07\x02\x37\x42\x7b\x96\xfb\x04\xb6\xbe\xd0\x4c\x90\xa7\x2f\xc5\x0e\x26\x79\x82\x69\x43\x75\x9b\xb3\x6c\xcc\x2b\x05\x91\x65\xc8\xb8\x5e\xa1\x64\xf6\x71\xae\x54\xf2\x4d\x8b\xee\x71\x6a\x9a\x5a\xe1\x1d\xa9\x47\x32\xfe\x0d\x94\xd3\xb9\x24\x98\xc6\xb1\xb6\x2e\xd9\x7b\x10\xa6\x7f\x3b\x9f\x11\x5a\xb7\xd1\x61\x1f\x71\xb9\x40\x4a\xa7\x6b\xfa\x3d\x4e\x72\xfe\x48\xd2\x38\x23\xbd\x67\x63\x6b\xbd\xcf\x5a\xce\x7e\x0a\x76\x88\x8e\x14\x0e\xd6\x90\x30\x9a\xf1\x47\x11\x1b\x65\x9d\x65\x75\x8d\x48\xa5\xdc\x0e\x8b\xc5\xbf\xfb\xff\xa2\xa0\x68\x82\xf5\x6c\x1b\x87\xf1\x72\xa8\xf2\xd9\x6d\x8e\x9b\xb9\x42\xa4\xad\x03\x23\x81\x3b\x5b\x51\x3c\x45\xa6\x1a\xda\x15\xe8\x74\x2a\x86\x56\x09\x6b\x7c\xa6\xd2\x1e\x15\x1d\xa8\xc2\xf3\xe8\xc7\xcf\x5f\xb3\x17\x21\x02\x45\xaf\x8f\x4e\x08\xdf\x44\x96\x1d\x47\x6f\x64\x10\x2b\x22\x8f\x59\x9a\xd0\x75\x11\x80\x6a\x63\xe8\x7e\xe1\x70\x9d\x7e\xf4\xb5\x8d\xa0\x3c\xb7\xa7\xf3\x1a\xe9\x00\xeb\x0e\xb9\x3d\xc0\x37\xd3\x47\x00\x00\x00\xff\xff\x19\x5f\x9d\x4a\x17\x02\x00\x00"

func dataAwsVpcPublicPrivateDeployCloudInitShTplBytes() ([]byte, error) {
	return bindataRead(
		_dataAwsVpcPublicPrivateDeployCloudInitShTpl,
		"data/aws-vpc-public-private/deploy/cloud-init.sh.tpl",
	)
}

func dataAwsVpcPublicPrivateDeployCloudInitShTpl() (*asset, error) {
	bytes, err := dataAwsVpcPublicPrivateDeployCloudInitShTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/aws-vpc-public-private/deploy/cloud-init.sh.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dataAwsVpcPublicPrivateDeployMainTfTpl = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x56\x4d\x6f\xe3\x36\x10\xbd\xfb\x57\x0c\x94\x66\x91\x05\x2a\xc9\x9b\x2e\x7a\x29\xda\x4b\x8b\xf6\xd6\x02\xbd\xf4\x50\x2c\x04\x4a\x1c\xd9\x44\x28\x92\xe0\x87\x53\xc3\xd0\x7f\xef\x90\xb4\x22\xc9\x49\x16\x31\x90\xa2\x0e\xe2\x48\x6f\x86\x7c\x33\xf3\x86\xc3\xdc\xc0\x6f\xa8\xd0\x32\x8f\x1c\xda\x23\xfc\xe1\xbd\xfe\x16\xb8\x06\xa5\x3d\x20\x17\x1e\x06\xa6\x02\x93\xf2\xb8\xd9\x1c\x98\x15\xac\x95\x08\x85\x50\xbd\x65\x8d\xe0\x05\x9c\xc6\x05\xcc\x1e\x5d\xc3\xba\x0e\x9d\x6b\x1e\xf0\xf8\x82\xd1\x61\x67\xd1\xbf\x62\xb4\xb8\x13\x5a\x5d\x18\xc8\xb5\x51\x6c\xc0\x04\x2f\x17\x0c\x82\x20\xe0\xd8\xb3\x20\x3d\xfc\x98\x90\xf2\xfe\xd3\xf7\xdf\x6d\xf9\xe7\xcf\x05\x8c\xab\x68\x9d\x67\xaa\xc3\xc6\x1f\x0d\x5e\xac\xf2\xf7\x95\x1b\x28\xbd\xf5\x0a\x63\xc5\x81\x2a\xd2\xb8\xd0\x2a\x8a\xf7\x59\xa2\x26\xb4\x52\x74\xaf\x9a\x0f\xa6\x6b\x3a\xc1\xed\x0b\xf0\xd9\x77\x81\xb6\xcc\x79\xca\xbb\xd9\x6b\xe7\x2f\x16\x4c\xa6\xe0\x30\xef\xb5\x31\x56\x1f\x04\x47\x9b\x2a\x46\xd0\x06\x60\x2e\x78\xcc\xe7\x9b\x13\x2d\xaf\xd6\x42\x8c\x05\xb9\xcd\xa5\x5f\xbb\xcd\x78\x72\xcb\x22\x40\xfc\xac\xdc\x32\x4e\x2e\x14\x84\x45\xa7\x83\xed\x66\x4d\x83\x15\xfe\xd8\xec\xac\x0e\xa6\x20\xd0\x98\x1c\x59\xd4\x2d\xef\x73\x3a\xe5\x97\x71\x2c\xf3\x96\x53\x03\x25\xce\x5c\x97\x99\x2f\xbf\x93\x89\x6c\x42\xed\x88\xce\xa5\xfd\x00\x28\x7d\xaf\x3b\x2d\x73\x78\xe5\xa7\x04\xf6\x56\x0f\x8d\xd1\xd6\x27\x70\x9b\x30\xaf\x27\x64\xc6\xa2\x22\x4d\x2b\x75\xf7\xe0\x08\xfb\x7b\x41\x16\x2d\x63\xf1\x85\xbc\x46\xfa\xc5\xff\x8c\x71\x5b\xa5\x9f\x7a\x7b\xe6\xa2\x5a\xde\xc0\x2f\x68\xa4\x3e\x02\x23\x85\x3c\xe8\x1e\xa6\x6e\x75\x17\x75\x9e\xf0\x65\x85\xa9\xe7\x61\xfe\xcc\x82\x0d\x22\x15\x76\xd5\xf8\xb3\x79\x05\xe7\xe6\x98\x1a\x79\xbd\xcf\xb3\x63\x90\x9c\xa7\x23\x79\x41\x3a\xc1\xc9\x27\x76\x6c\xc3\x99\x67\xb3\x4f\x2f\x24\xde\xd1\x5f\xc3\xfc\xbe\x1a\x34\x0f\x12\xc7\xba\x93\x3a\xf0\x52\x28\xe1\x2b\xb7\x2f\x3e\x66\xcd\xa3\x24\xeb\xae\x22\xea\x49\xb3\xe7\x2d\x57\x51\x39\xaa\x18\xdb\x97\xb8\xd8\xb3\xdd\xa4\xdd\xef\x31\xc8\x55\xf7\x15\xa9\xec\xf4\xd5\x69\xa5\xb0\x8b\x87\xeb\xec\x1b\x03\x5e\x16\x32\xb4\x41\xf9\x50\x24\x5b\x3c\x99\xeb\x22\x3b\x94\xfd\x53\x75\x84\x19\xb3\xdf\xf2\x24\xcf\x75\x59\xa2\x17\x8e\x89\xf4\x99\x63\x44\xe7\x48\x6f\xe0\x2f\x46\x53\xb8\xd7\x16\xe6\x62\xc1\x1d\x2a\x17\xa8\x3f\x9e\x24\x06\xe1\xa0\x0f\x34\xa6\xa1\xd5\x3a\xcd\x72\xa4\x25\x08\x03\xcd\x0b\xb5\x03\xad\x3e\x6e\x52\x33\x1f\x84\x23\x8e\x38\x40\x2c\x0e\xe4\x58\xe2\x3f\xd8\x15\xe7\x22\x08\x25\x85\xc2\x54\xe8\xc7\x3d\xc9\x45\x7d\x41\x37\x81\x79\x10\x52\x42\xb9\x5d\xf2\xdf\xff\x54\x73\x3c\xd4\x8a\x18\x7f\x88\xb7\x85\x93\x88\x06\xee\xe3\xb3\xc2\x73\x77\xd3\xd7\xe9\x36\x05\xce\x85\xa5\xbd\xe9\x31\x28\xea\x89\x98\x23\x21\xae\x6a\x83\x90\x1c\x6e\xc7\x94\xe5\xaf\x4f\x46\x20\xbd\x68\x95\xd4\xda\x54\x3f\x13\xea\x29\xda\x71\x84\xbb\xe4\x7e\x65\x1a\xc3\x43\xe4\x2e\x0d\xd4\x7e\x30\xb5\xa6\xcb\xad\x9e\xa3\x28\x5f\x24\x9a\xa3\x5f\xf1\xc4\xf6\x9d\x08\xce\xc7\x32\xb7\x56\x24\x18\xc7\x3a\x2b\xcb\x91\x34\x54\x39\x0d\x32\x5f\xc1\xfa\x22\xe9\xd7\x93\xeb\xf8\x5b\xd3\xa2\xf2\x7d\xf8\x10\xdb\x6e\x0f\x55\x3d\x30\xa1\xe2\x69\x9b\xe6\x1d\x89\x84\x8a\x47\x9d\x6e\xdf\x20\x1a\xcf\xe3\xea\xcd\xaa\x65\xff\x77\x95\x2d\x6f\xf9\x3f\xa9\xf7\x55\xf2\x77\x14\xf1\x55\x9e\xab\xb4\xbc\x81\x3f\x89\xfd\x80\xc0\xd4\x11\x3c\x0e\x74\x45\x31\x7b\x8c\x59\xd3\xf4\xd3\x56\xd0\x04\x79\xa4\x21\xc1\x38\xa6\x4b\x6d\xa1\xb6\x83\x3b\xd1\xc7\x65\x57\x4a\x67\x07\x28\x6d\x3f\xe7\x34\x5f\x75\x3a\x78\x13\x3c\xfd\x3b\x76\xbe\xbc\x0e\x4c\x86\xf3\xad\xb4\xbc\xdf\xd2\x38\x5f\xcd\xd6\x71\xf3\x6f\x00\x00\x00\xff\xff\xbc\xac\xe0\x57\xa5\x0a\x00\x00"

func dataAwsVpcPublicPrivateDeployMainTfTplBytes() ([]byte, error) {
	return bindataRead(
		_dataAwsVpcPublicPrivateDeployMainTfTpl,
		"data/aws-vpc-public-private/deploy/main.tf.tpl",
	)
}

func dataAwsVpcPublicPrivateDeployMainTfTpl() (*asset, error) {
	bytes, err := dataAwsVpcPublicPrivateDeployMainTfTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/aws-vpc-public-private/deploy/main.tf.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dataCommonDevVagrantfileTpl = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x54\x4d\x6f\xeb\x36\x10\xbc\xeb\x57\x6c\x99\xf7\x02\x1b\x88\x24\x34\x28\x7a\x48\x93\xa0\x69\x3e\x90\x5c\xdc\x20\x76\x7b\x09\x02\x81\x16\x57\x12\x11\x8a\x24\x48\x4a\xb6\x61\xfb\xbf\x77\x29\xd9\x6d\x82\x1c\xfa\x0e\x12\xc8\xe5\xec\x70\x76\x77\xa4\xbf\x79\xed\xb8\x0e\x59\x69\x74\x25\xeb\xce\xe1\x84\x9d\xb3\x29\x08\x03\xbb\x31\xb4\x4b\x00\xc6\x55\xd6\xb7\xd9\xd2\xac\xe1\x0a\x58\xc3\x7d\x23\x4b\xe3\x6c\x6e\x1d\x96\xd2\xe3\xaf\xbf\xb0\x84\x80\x27\x30\xc7\xd0\x59\xf0\xa6\x45\xf0\xa1\xab\xaa\x4f\xd9\xd6\x99\x5e\x7a\x69\x34\x30\xdf\xa0\x52\xec\x0c\xa4\x56\x52\xe3\x05\x7c\xf3\xa5\x93\x36\x8c\x24\x0f\xa6\xd3\x82\x87\x08\x3c\xea\x1a\x77\x13\x59\x01\xd7\x9b\x29\xa1\xb6\xdf\xa1\x32\x0e\x84\x74\xc4\x41\xcb\x63\x46\x41\x11\x9f\x09\xec\xe1\xfb\x9e\x60\xf1\x9c\x04\xe7\x26\x04\x93\xff\x87\x4a\xb7\xdb\x98\xae\x8c\xb1\xd9\x2d\x45\x03\x3a\xd8\xef\xd9\x27\xb5\x7e\xa3\x4b\x14\x45\x65\x94\xa0\x53\x46\x19\x91\x8c\x50\x67\x71\xf1\xa3\x85\xb1\x52\xc0\xc9\x96\x12\xf6\x70\x7a\x0a\x4b\x6a\xdc\x61\x9b\xb7\x5c\xea\xcc\x37\x6c\x2c\x06\xb5\x88\xf5\x90\xe8\xa1\x05\x2f\xc8\x45\x2c\x2c\x34\x08\x95\xe3\x75\x8b\x3a\xd0\x86\x07\x58\x21\x74\x1e\x81\x7b\xe0\x20\xd0\x12\x1a\x7b\xae\x26\x0f\x52\x61\xe6\x28\x6b\x12\x95\x1e\x53\x0a\xcb\x43\x13\x35\x4f\xcf\x60\x29\xb5\x90\xba\x9e\xc6\x0b\x5e\x59\xdf\xae\xb8\xc3\xa2\xea\xa2\x6c\xd2\x7b\x0c\xac\x8c\x7b\xf7\x61\x68\x12\x7b\xcb\x90\x97\xcd\x60\x06\xcd\x5b\x8c\x56\xf8\x52\x35\xf5\x66\x12\x0f\x47\xcf\xd8\x11\x03\x60\x33\xd4\x7c\xa9\xb0\xe8\x5b\xd7\xe9\x42\xda\x82\x7a\xfd\x4e\xd6\xb8\x82\x8a\x2b\x8f\x03\x8c\x6a\x4e\xc6\x77\x7c\x92\x83\x07\x08\x72\x79\x39\xbf\x7d\x79\x7a\x5e\x24\x1e\x03\xa4\x98\x24\x27\x10\x07\x98\xe2\x1a\xcb\x0b\x88\xef\x2e\x20\x29\x69\x5b\xae\x05\xac\x24\xd5\x68\xba\x60\xbb\x00\xca\xd4\x35\x0a\x58\xd2\x52\x9b\x40\x83\xf2\x56\xf1\x0d\x8a\xc4\xe0\x64\x0a\x5b\xf8\xf6\x3b\x9c\x5f\x9f\xfe\x0c\xbb\x11\xe9\x20\x0d\x03\x35\x5c\x43\x4e\xa6\xc9\x75\xa7\xd4\x6f\xb0\xff\xf7\x46\x42\x5d\x1c\xb9\x39\x90\xdb\x2b\xb9\x26\xfe\x16\xbd\xe7\x35\x26\x46\x0d\xac\x58\x36\x06\xd8\x6b\xcc\x78\xa3\x2b\xd8\x81\xe1\xf6\x60\x5e\xea\x3a\xcc\xe7\x8f\x83\x61\x2b\xee\xa3\xd9\x88\x57\xea\x84\xec\xfc\x13\xd4\x0e\x2d\xb0\xbf\x3c\xde\xcd\xe6\x24\x9a\x41\x8e\xa1\xcc\xbd\x6f\xe2\x23\x8a\xb1\xdf\x70\xfd\x41\x1e\x99\x42\xc7\xc6\x0d\xb7\x7e\x48\xdc\x81\xef\x68\x0a\x01\x11\x52\xfe\x7f\x34\x44\x60\x70\x4c\xf0\xe8\x7a\x59\xd2\x86\xcc\xe9\x90\xa6\xef\x42\x52\xc9\x24\xc1\xb5\x35\x2e\xc0\xdd\xfd\x1f\x4f\x37\xb3\xe2\xe1\xe5\xcf\xd9\xe2\x7e\x76\x77\xa5\x8d\x96\xf1\x8b\xe1\x65\x90\x3d\x0d\xc7\x28\x60\x4f\x9a\xd2\x94\x8a\x95\x3e\x2e\x16\xcf\x73\x10\x8e\xce\xdc\x50\xf2\x8d\x0d\x59\x96\xb1\xe4\x78\x1d\xb7\x21\xad\x69\xb0\x9d\xa5\x8f\x11\xbf\x84\xe5\x48\x05\xe9\x66\x08\x05\xfa\x3f\xf9\xa8\x23\x6d\x42\xb0\x3e\x39\x38\xe3\x9f\x00\x00\x00\xff\xff\xe4\xa7\xd3\xe3\xb9\x04\x00\x00"

func dataCommonDevVagrantfileTplBytes() ([]byte, error) {
	return bindataRead(
		_dataCommonDevVagrantfileTpl,
		"data/common/dev/Vagrantfile.tpl",
	)
}

func dataCommonDevVagrantfileTpl() (*asset, error) {
	bytes, err := dataCommonDevVagrantfileTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/common/dev/Vagrantfile.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dataCommonDevDepVagrantfileFragmentTpl = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x91\x4f\x6f\x9c\x30\x10\xc5\xef\xfe\x14\x23\xd2\xec\xa9\x80\xda\x63\x9a\x8d\xd4\xe6\x8f\x94\xcb\xb6\x4a\x72\x47\x2e\x1e\xc0\x0a\xcc\x58\xb6\x21\x8d\x36\xfb\xdd\x3b\x86\x15\x5d\x55\x55\x15\x4e\x9e\x99\xc7\x9b\xdf\xb3\x3f\x38\x8f\xa1\xc3\xbe\x87\x2d\x5c\x5e\x3e\x5e\x3f\xdc\xff\x78\x52\xb6\x81\x9a\x87\x41\x93\x81\x7c\x02\xed\x62\xde\x62\x84\xab\xd2\xe0\x54\xd2\x28\xda\xcf\x57\x9b\x4f\x5f\x20\x76\x48\x0a\xe4\xc3\x5f\x8e\x7d\x84\x9b\xdb\x6f\xf7\x5f\x77\xd5\xdd\xc3\xf7\xdd\xd3\xed\xee\x66\x4b\x4c\x96\x22\x7a\x5d\x47\x3b\xe1\xac\x0c\xa3\xe1\xd5\xd0\x52\x88\x5a\xdc\xf2\xd7\xb9\x15\xbd\xa6\x90\x8c\xf2\x2e\x46\x17\xfe\xde\xa7\x1a\xab\x8e\x7c\xaa\x66\x6a\x6c\x5b\x4c\x43\xe1\x3c\x4f\x36\x58\x26\xc8\xe6\x18\xd9\x47\xb1\xed\x2d\xe1\x05\xac\xd1\xfe\x2d\x37\x5c\x3f\xa3\xcf\x40\x80\xde\xcc\x9b\xd0\x99\xc2\x8f\xd2\xdf\xef\x61\x19\x55\x76\xd0\x2d\xc2\xe1\x20\x9e\xda\xb7\xe1\x62\x9e\x89\xa6\x4a\x55\xea\x2b\x24\xa3\xd4\x19\x3c\xbe\x52\x0d\x3c\x7a\xe0\x17\x02\x83\x0e\x1a\xee\x0d\x7a\x21\x49\x77\xe4\xf1\x04\x20\x88\x14\x4d\x75\x14\x24\x43\xa7\x63\x57\xbc\xb0\x7f\xb6\xd4\x2e\xcb\xd6\x6e\x3b\x62\x88\xd5\xc9\x2c\x2d\xbb\xe3\x91\x8c\x8e\x29\xc3\x62\x3b\xfa\xa5\x6a\xd8\xcb\xf2\x29\x01\xa8\xfd\xf9\x52\xda\x19\xa2\x59\x7f\xa9\xa4\x13\x0a\x51\x55\x09\xf3\xfc\xa0\x92\x62\x0b\x59\xc9\x31\x72\xf9\x47\x97\x0b\x02\xe9\x21\xa5\x4f\x47\xf1\xea\x99\x5d\x71\x2d\x02\x79\xd0\x19\xe5\x7f\x99\x92\xeb\x1c\x45\x0e\xef\x7b\xac\xac\x36\x70\xb6\x17\xf9\x01\x36\x1b\xf8\xa9\x43\x77\x2c\xcb\x41\x5b\x2a\x42\x97\xa5\x4c\x72\xdf\x29\x96\x70\xab\xdf\x01\x00\x00\xff\xff\xdd\x96\xb5\xa6\xb9\x02\x00\x00"

func dataCommonDevDepVagrantfileFragmentTplBytes() ([]byte, error) {
	return bindataRead(
		_dataCommonDevDepVagrantfileFragmentTpl,
		"data/common/dev-dep/Vagrantfile.fragment.tpl",
	)
}

func dataCommonDevDepVagrantfileFragmentTpl() (*asset, error) {
	bytes, err := dataCommonDevDepVagrantfileFragmentTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/common/dev-dep/Vagrantfile.fragment.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"data/aws-vpc-public-private/deploy/cloud-init.sh.tpl": dataAwsVpcPublicPrivateDeployCloudInitShTpl,
	"data/aws-vpc-public-private/deploy/main.tf.tpl": dataAwsVpcPublicPrivateDeployMainTfTpl,
	"data/common/dev/Vagrantfile.tpl": dataCommonDevVagrantfileTpl,
	"data/common/dev-dep/Vagrantfile.fragment.tpl": dataCommonDevDepVagrantfileFragmentTpl,
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
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"data": &bintree{nil, map[string]*bintree{
		"aws-vpc-public-private": &bintree{nil, map[string]*bintree{
			"deploy": &bintree{nil, map[string]*bintree{
				"cloud-init.sh.tpl": &bintree{dataAwsVpcPublicPrivateDeployCloudInitShTpl, map[string]*bintree{
				}},
				"main.tf.tpl": &bintree{dataAwsVpcPublicPrivateDeployMainTfTpl, map[string]*bintree{
				}},
			}},
		}},
		"common": &bintree{nil, map[string]*bintree{
			"dev": &bintree{nil, map[string]*bintree{
				"Vagrantfile.tpl": &bintree{dataCommonDevVagrantfileTpl, map[string]*bintree{
				}},
			}},
			"dev-dep": &bintree{nil, map[string]*bintree{
				"Vagrantfile.fragment.tpl": &bintree{dataCommonDevDepVagrantfileFragmentTpl, map[string]*bintree{
				}},
			}},
		}},
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
