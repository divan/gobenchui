// Code generated by go-bindata.
// sources:
// assets/.main.css.swp
// assets/index.html
// assets/main.css
// assets/skeleton.css
// assets/websocket.js
// DO NOT EDIT!

package main

import (
	"github.com/elazarl/go-bindata-assetfs"
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

var _assetsMainCssSwp = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x99\x4d\x6b\x1b\x47\x18\xc7\xc7\xa5\xa5\xa8\xae\xed\x52\x97\x52\x7a\x9a\xda\x35\x6e\xa9\x56\xda\x95\x2c\xea\x4a\xbd\xf8\x52\xdc\x43\xdd\x16\x63\xd3\xd2\xd3\xec\xee\x48\x3b\x74\xb5\x23\x76\x57\xd6\x8b\x71\x0b\xbd\xb7\xd7\x1e\x03\xb9\xe4\x03\x84\x84\x84\x90\x63\xf0\x07\xc8\x25\x97\x84\x5c\xf2\x01\xf2\x72\xc9\x2d\xff\x95\x56\xb2\xf5\x62\x64\x92\x10\x13\x78\x7e\xcb\xcf\xcb\x8c\x9f\x79\xe6\x79\x66\xf6\x62\x6c\x9b\xfb\x3f\xfe\xc4\xbf\xcd\x6d\x30\xf0\x11\x63\x8b\x7b\x5f\xee\x57\xbf\xbf\xf2\xee\xf6\x3f\x73\xcc\x55\x07\x22\x60\xb3\xe9\xc5\x19\x75\xe1\xe4\x7c\xed\x08\xff\xcc\xb8\xbf\x7a\x81\xf9\x28\x74\xf2\x35\x15\x7b\x4d\x3b\xe7\xe8\x7a\xbe\x3f\x59\xd3\xb6\x0c\x1c\xaf\xa9\xf2\x22\x8a\x64\x1c\xe5\xeb\x42\x05\x39\x27\x8a\xce\xb1\x3f\x41\x10\x33\x68\xc6\x55\x63\x73\x81\x15\x0b\x96\x99\x0c\x57\x57\xbe\xe0\xcb\x1f\xef\x5d\x74\x55\x04\x41\x10\x04\x41\x10\x04\x41\xbc\x41\xe2\xc6\x1c\xfb\x1b\xef\x77\xd2\x71\x25\x7d\xcf\x8d\xbd\x09\x82\x20\x08\x82\x20\x08\x82\x20\x08\x82\x78\x7b\x11\x2e\x63\xb7\x3f\x60\xec\xce\x3c\xeb\xfd\xff\x7f\xf0\xf7\xff\xf3\x25\xc6\x6e\xc1\x9b\xf0\x06\xbc\x0a\xbb\xd0\x83\x7b\x70\x0b\x7e\x03\x3f\x85\x9f\xc0\x65\xb8\x08\x19\x3c\x5e\x64\xec\x1a\xbc\x0c\xff\x83\x87\xb0\x03\xdb\x30\x82\x12\x66\xe1\xe7\xf0\x7d\xf8\x74\x81\xb1\x87\xf0\x01\xbc\x0f\xef\xc1\x4b\xf0\x5f\xd8\x85\x55\xb8\x0b\x7f\x85\xbf\xc0\x1f\xe0\x26\x5c\x83\x4b\xf0\xee\x87\x8c\xfd\x0f\xff\x80\xbf\xc3\xdf\xe0\x0e\xac\xc0\x2c\xfc\x0c\xbe\x07\x9f\xa1\xc7\x27\xf0\x31\x7c\x04\x8f\xe7\xfb\x7d\x5f\x9f\xbf\xc0\x0b\x20\x08\x82\x20\x08\x82\x20\x88\x59\x1c\x1d\xb1\x8c\xad\x43\x57\x86\x65\x6e\x35\xda\x7c\x6c\xe8\xab\x9a\x17\xd7\x42\xd9\xe1\x91\xf6\x95\x5b\x61\xb9\xfe\x6f\xa5\xcb\x0f\xb1\x98\x65\x1c\xed\x6b\xc4\xae\x9a\xd5\xe4\xa9\xb0\x4c\x55\x07\xb1\x11\xa9\xae\x44\x82\x42\xa3\x5d\xc9\xb0\x4c\x2c\xdb\xb1\x21\x90\x2a\x28\xf3\x30\x49\x88\xb0\x86\x70\x5d\x15\xd4\xca\xbc\x84\x18\x96\xa9\x8b\xb0\xa6\x02\x23\xd6\x0d\x2c\x33\x93\xa9\x5c\x55\xeb\x58\x86\xe9\x36\xb6\x6e\x1b\x91\x27\x5c\xdd\x2a\xf3\x01\x08\xe3\x45\x88\x5d\x78\x58\xb3\xc5\x57\x96\x69\x66\xf9\xc9\x0f\x33\xb7\xf1\xdd\xd7\xc8\x6d\xd4\x75\xd7\x18\x4b\x30\x58\x5b\x9a\xb5\xb6\x25\xed\x3f\x55\x3c\xb2\xfc\xbc\x6b\xd3\x9e\x6c\x1d\xc7\xba\x8e\xb6\xfa\x9d\x0e\x1b\x4f\xca\xee\x75\x9a\x69\x29\x37\xf6\x92\xbe\xcd\x35\xf4\xed\x49\xe1\x0e\xfb\xee\xe7\xc0\x9e\x78\xac\x8d\x5e\xb4\xaf\x02\x69\x78\x32\x39\xc7\xf2\x60\xae\x77\xe8\xad\x74\xae\x68\x9a\x63\xf7\x70\x2a\xa8\x2a\xea\xca\xef\x94\xf9\xfa\xb6\xf4\x0f\x64\xac\x1c\xc1\x77\x64\x53\xae\x67\xf9\xfa\xcf\x0d\x19\xf0\x5d\x11\x44\x18\x44\x78\x19\x91\x0c\x15\xae\xd4\x2b\x4d\x29\xa6\x60\x4e\x16\x93\xce\x8d\x14\x63\x8d\x17\x73\x3a\xe8\x25\x8a\x19\x7e\x70\x96\x65\xa1\xb4\xe2\xb4\x73\x9a\x56\x5a\x61\xb2\xb4\xc2\x44\x69\x9b\xaf\xb3\xb4\x42\x5a\xda\xe9\xaf\xdf\x91\x01\xbe\xe9\xca\x49\xbd\xc9\x69\x9c\x71\x9a\xc5\x29\x25\x4f\x5c\x6d\xe9\xd5\x4e\xd3\xb3\xd2\x22\x47\x12\xac\x8c\x26\x58\xc9\xf2\xe1\x44\x96\x6f\x85\x4a\xf8\xa3\x59\x6c\xed\x76\x90\xe7\x45\x00\x00\x00\xff\xff\xb2\x37\x81\x17\x00\x30\x00\x00")

func assetsMainCssSwpBytes() ([]byte, error) {
	return bindataRead(
		_assetsMainCssSwp,
		"assets/.main.css.swp",
	)
}

func assetsMainCssSwp() (*asset, error) {
	bytes, err := assetsMainCssSwpBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/.main.css.swp", size: 12288, mode: os.FileMode(420), modTime: time.Unix(1445221669, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _assetsIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x56\xdb\x6e\xe3\x36\x10\x7d\x76\xbe\x82\x50\x0b\xd8\x46\x63\xc9\xf6\x36\x5b\x54\x91\x05\x34\x5b\xf4\xf2\xd0\xc5\x02\x69\x9f\x8a\x22\xa0\xa4\xb1\x45\x5b\x22\x55\x92\xf2\x05\xa9\xff\xbd\x43\x52\x52\x2c\x7b\x83\xa4\x41\xd1\x7a\x81\x8d\x34\x33\x67\xce\x70\x6e\x62\x94\xeb\xb2\x88\xaf\xa2\x1c\x68\x16\x5f\x0d\x22\x95\x4a\x56\x69\xa2\x64\xba\xf0\x72\xad\xab\x30\x08\x52\x91\x81\xbf\xfe\xb3\x06\x79\xf0\x53\x51\x06\xee\x71\x32\xf3\xbf\xf5\x67\xfe\x5a\x79\x71\x14\x38\xd4\x33\xf8\x52\x94\xc0\xf5\x5a\x59\x70\x26\x76\xbc\x10\x34\x53\x8d\xf8\x35\x0e\x6c\x00\x39\x5b\xe5\x69\x4e\xa5\x76\x7e\x4e\x5e\xdf\xe8\xa1\x14\x59\x5d\x80\x0a\x60\x5f\x09\xa9\x19\x5f\x5d\x38\x2a\x18\xdf\x10\x09\xc5\xc2\x53\xfa\x80\xa6\x39\x80\xf6\x88\x3e\x54\xb0\xf0\x34\xec\x75\x90\x2a\xe5\x91\x5c\xc2\x72\xe1\x05\x4a\x53\xcd\xd2\x40\x6d\xa0\x00\x2d\xb8\x6f\x74\x6f\x75\x52\x52\xd6\x3a\x20\xf8\x6b\x4f\x74\x02\x5a\xd3\x2d\x75\xd2\xc6\xc6\xfc\xb6\x54\x12\x25\xd2\x0d\x59\x10\x5e\x17\xc5\x6d\x4f\xb1\x53\xb5\x64\xa8\xf1\x76\x0a\x13\x32\x9b\x7f\xe3\x4f\xf1\xdf\x2c\x7c\x3f\x9f\xcf\x83\x9d\xf2\x6e\xaf\xae\x06\x03\x63\x28\x2a\xcd\x04\x57\x0f\x9a\x95\x80\xf6\x8f\x28\x1e\xd8\xb4\x85\xe4\xd1\x86\x10\x92\xa1\xaa\xf0\x58\x30\xbc\x26\xa9\xe0\x1c\x52\xfd\x11\xe9\x54\x48\xb4\xac\xe1\x9a\x1c\xaf\x0d\x44\x33\x5d\x80\x85\x60\xbc\x21\xf1\x1e\x1f\x89\xff\x69\xb3\xfa\x48\xd1\xeb\xf1\xe8\x35\x56\xaa\x4e\xce\x0d\x2d\x6f\x02\x3c\xcd\x4b\x2a\x37\x64\xc4\x55\x20\xaa\x71\x0b\xd8\x7f\xb7\x67\xca\x58\x4b\xd8\x82\x54\x90\xb5\xac\x29\xd5\xb0\x12\x92\x01\x6a\x7f\xff\xa3\x8d\xe2\xd0\x98\x9b\xe7\x41\x41\x13\x28\xba\xb7\x8b\x08\x7f\x45\xe2\x96\x66\x30\x58\x0a\x59\x52\xad\x41\x86\x64\x59\xf3\xd4\xe4\x84\x8c\xc6\x96\x58\xd7\x92\x13\x9d\x33\xe5\x6f\x69\x51\x03\xf9\x8a\x0c\xb9\x1a\xde\x92\xa3\x45\xda\xff\x9b\x1c\x08\x51\x68\x56\x19\x8e\x54\x0a\xa5\x72\xca\x64\x97\x26\x7c\x93\x6d\xf8\x0d\x2b\xe6\x88\x2d\x89\x6f\x02\xb9\x07\x73\x14\xcc\x14\x31\x89\x3b\x91\xfc\x45\xd6\x4a\xf0\x07\xa5\xb1\xf8\x15\x64\x8d\x05\x70\xf3\x84\x3e\x8e\x17\x75\x2c\xa1\xfc\x3f\xca\x88\xb4\x42\x1e\x4e\x0b\x79\xf7\x1f\xd5\xf1\x17\xcb\xfc\xd6\x4a\xde\xfd\x8b\x85\xc4\x48\xfa\x75\x7c\x12\xbc\xaa\x8c\x3b\xc6\x71\x5f\xfa\xc2\x6e\x4c\x2c\x62\x1b\xbe\x89\xde\xd0\x7c\x39\x1a\x7e\x61\xc6\xe5\xa1\xcb\xf2\x70\x7c\xb2\xe6\x46\xa7\xa3\x3c\xbe\x6d\x11\x58\x99\x97\x00\x68\xe2\xec\x4d\x23\x59\x5d\xbb\x0e\x5e\xa2\x3c\x87\xb9\xee\x7b\x81\xd6\x82\xf0\xc8\x83\xe7\xd6\x78\xbb\x18\x77\x90\x98\x05\x07\xe7\x9f\x8d\x28\x70\xdf\xaf\x28\x11\xd9\xc1\x80\x33\xb6\x25\x69\x41\x95\xc2\x4f\x00\x6a\x40\x9a\x5d\x3c\x88\xf2\x59\xfc\xa3\x20\x77\x5d\x4f\xfe\xf6\x33\x22\x67\x06\x10\x20\xe2\x0c\x28\xc5\xce\xa1\x4e\x64\x82\xc3\x04\xfb\x45\x66\x38\x32\x45\x5d\x72\x8f\xd8\xad\xee\xb6\xf2\x84\x16\x6c\xc5\x43\x52\xc0\x52\x5b\x24\x12\xce\xe3\xfe\xc4\x20\xdf\xbc\x51\xbd\x8b\x3f\x51\x9d\x87\xa4\x31\x30\x2f\xce\xe0\x9d\x65\x6d\x22\xfa\xe7\xf4\x29\x7e\x56\x9b\x03\xdb\x00\xee\x31\x77\x35\x76\x6b\xa4\x2a\xca\x09\xcb\xcc\x97\xc8\x48\x3c\x1b\x9a\xd3\x5a\x62\xa3\x8f\xbb\x00\x9b\x1e\xfe\x50\x4b\x89\x0e\x3f\x88\xb2\x64\xda\x35\x27\x7a\xbd\xb1\x7e\x52\x2b\x7c\x48\x0a\xac\x89\x17\x37\x96\xc4\x49\x4f\xf9\x9c\xc4\xf1\xf5\xfc\xf9\xdf\xe3\xb4\xfb\x3f\xd8\x01\x25\xde\x7c\x3a\x7d\x3f\x99\xce\x26\xd3\x39\x99\xdd\x84\xd3\xaf\xc3\xe9\x8d\x67\x06\x63\x74\x89\xbb\xaf\x93\x35\x2e\x2c\xd4\x8e\xa3\x44\x92\x20\xfe\x89\xaa\x26\x95\x7d\x43\x23\xef\x1d\xee\xa6\x3d\x5c\x37\x6b\x6f\xcf\xb4\xc4\x16\xee\x57\xba\x00\x8e\x11\x58\x6a\x3b\xf4\xee\xe4\xaa\x57\x75\x34\xdb\x31\xac\x35\x12\xe9\xce\xf8\x78\x34\xb1\x3f\x93\x0e\x93\x85\x2e\x62\x32\x21\xad\x07\x5c\x0f\xb0\x7f\xe2\x9b\x91\xd7\x7b\x39\x6f\xb3\xd7\x0e\x40\x22\x24\x4e\x13\x64\x9e\x2d\x6c\x7f\x13\x98\x89\xfc\x4c\x2a\xfb\x90\xde\x1a\x78\x42\x7c\x8e\x7f\x29\x84\xee\x26\xd7\x14\xd9\x3c\x54\xf1\x0a\x38\x48\x3c\x62\xe6\x92\x10\xd1\xe6\xf6\xd4\xdc\xf5\x56\x28\xac\x13\x77\xd9\x64\x5b\xca\x83\x95\xb0\x7c\x35\xc3\x6b\x17\x95\x2b\xd0\x0b\x0f\x5b\x96\x72\x64\xef\x54\x51\x40\x63\x92\x1c\x5e\x74\x76\xe9\xc2\x8a\x0d\x3c\x0a\xaa\xd3\x73\xb8\xbf\x51\xe0\x36\x12\xa6\xdb\xde\xb3\xff\x0e\x00\x00\xff\xff\xde\x13\x06\xad\x6f\x0b\x00\x00")

func assetsIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsIndexHtml,
		"assets/index.html",
	)
}

func assetsIndexHtml() (*asset, error) {
	bytes, err := assetsIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/index.html", size: 2927, mode: os.FileMode(420), modTime: time.Unix(1445221677, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _assetsMainCss = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x92\xcb\xce\x9b\x30\x10\x85\xd7\xe1\x29\x46\xa9\x2a\x5a\x09\x22\xe3\x04\xa9\x25\xab\xee\xba\x6a\x17\x7d\x02\x13\x0f\x60\xd5\x60\x64\xdc\xe6\x52\xf5\xdd\x3b\x06\x72\x21\x44\x55\x16\xf9\xb1\xb8\x68\x38\x67\xfc\xcd\x78\x72\x23\x8f\xf0\x27\x58\x14\xa6\x71\x71\x21\x6a\xa5\x8f\x19\x2c\xbf\xa2\xfe\x8d\x4e\xed\x04\x7c\xc3\x5f\xb8\x8c\xe0\x12\x88\xe0\x8b\x55\x42\x47\xd0\x89\xa6\x8b\x3b\xb4\xaa\xd8\x06\x7f\x83\xa0\x4a\x66\x59\xc2\x69\x96\x30\x82\xf0\x7b\x8b\x0d\xfc\x20\x67\x38\x4d\x30\x18\x3b\x75\xc2\x0c\x52\xd6\x1e\xce\x91\x3d\xaa\xb2\x72\x19\xac\x19\xa3\x90\x56\x0d\xc6\xd5\x39\xc4\x7b\x59\x2d\x6c\xa9\x9a\x0c\x38\xb9\x80\xf5\x2f\x8a\x3a\x3c\xb8\x58\x68\x55\xd2\x9f\x1d\x36\x0e\xed\x00\xc9\x3d\xe4\xce\x68\x63\x33\x78\x97\x24\xc9\xf6\x35\xc8\xfc\xd3\x1c\x99\xcf\x91\xf9\x14\x99\xd1\x4a\x7a\x5e\x8f\xb6\x7e\x23\xb4\x07\xdd\x4c\x1e\xa0\xb1\x19\x1a\xbf\xa0\xa5\xaf\x3a\xda\x64\xf3\xd4\xd1\x8e\xb2\x49\x9f\x36\x23\xcc\xaa\x42\x21\xd1\x7a\xa2\xbd\x92\xae\xea\xab\x79\x4f\xea\x56\x48\xa9\x9a\x92\x4a\xf1\x83\x70\x93\x20\xce\x8d\x73\xa6\x26\x61\xda\x47\x69\xe7\xfc\xa7\x72\x14\x3e\xc4\x5d\x25\xa4\xd9\x67\x5e\x0f\x6b\xba\x49\x01\xb6\xcc\xc5\x07\x4a\x1a\xc1\xf5\xc1\x56\x9b\xcf\x1f\xbd\xb7\x36\xa7\x89\x91\xae\x67\xbd\x77\x36\xb8\xf5\xf2\xff\x7b\x7d\xdd\x85\x31\x6e\xa8\x7b\x2c\xcb\x99\x36\x1b\xe7\xe7\x5a\x7c\x3a\x1b\x7f\xeb\x7b\x7a\x77\x0c\x7e\x0e\x17\xd7\x69\x63\x85\x5f\xc3\x36\xb9\xb1\xd4\x5e\x94\x7e\xa3\xe1\x9b\xf4\x44\xa7\x7d\x9a\xd2\xe2\x11\x3a\xa3\x95\xf4\xe2\x7f\x01\x00\x00\xff\xff\x76\xd1\x9d\x81\x3b\x04\x00\x00")

func assetsMainCssBytes() ([]byte, error) {
	return bindataRead(
		_assetsMainCss,
		"assets/main.css",
	)
}

func assetsMainCss() (*asset, error) {
	bytes, err := assetsMainCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/main.css", size: 1083, mode: os.FileMode(420), modTime: time.Unix(1445221646, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _assetsSkeletonCss = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x3a\xdb\x8e\xdb\x38\xb2\xef\xfe\x0a\xa2\x07\x03\x4c\x02\x5b\xed\x7b\xba\xdd\x98\x83\x93\x4b\x77\x66\x81\x49\x82\x4d\xb2\xbb\x0f\x8b\x79\xa0\x25\xca\x22\x5a\x12\x35\x24\x15\xb7\x27\x18\x60\xff\x61\xff\x70\xbf\x64\xab\x28\x51\x37\x93\x8e\xf7\xc9\x71\xd2\xb0\x55\xc5\x2a\x16\x8b\x75\xa5\x78\xfd\x7c\xf4\x9c\x7c\x7a\x64\x29\xd3\x22\x27\x7f\x9f\x07\xd3\x60\x09\x90\xd7\xa2\x38\x48\xbe\x4b\x34\x99\x4f\x67\xcb\x31\x79\x43\xbf\x30\xf2\x96\x66\x34\x4c\x18\xa0\xf7\xfb\x7d\xb0\x63\x5a\xd5\x74\x41\x28\x32\x80\x3e\x48\xc6\x88\x16\xa4\x54\x8c\x94\x79\xc4\x24\xd1\x09\x23\xef\xfe\xf2\x99\xa4\x3c\x64\xb9\x62\x01\x0c\x4a\xb4\x2e\x36\xd7\xd7\xc8\x41\x14\x00\x14\xa5\x0c\x59\x20\xe4\xee\xba\x1e\xa4\xae\x33\xae\x27\x96\xa2\x48\x0a\x20\x9a\xcd\xaf\xe7\xb7\xd7\x28\xca\xe8\xf9\xf5\x68\x34\xba\x7e\x4e\x3e\xd3\x6d\xca\x88\x88\x49\x28\x72\xcd\x72\xad\x46\xff\xf9\xd7\xbf\xbf\xc3\xff\xa3\x09\x79\x2b\x79\x04\x5f\xaf\x28\xe8\xe5\x93\x3e\xa4\x4c\xc1\xd3\xe7\x43\x21\x76\x92\x16\xc9\x01\x1e\x7e\xe5\xf9\x23\x02\x5f\x95\x1a\xd4\x89\xbf\x1e\x84\xcc\x94\xc1\x28\x8d\xdf\xaf\x45\xc4\x90\x0a\x57\x8d\xcf\x9f\x0a\x1a\xf2\x7c\x07\xbf\xfe\xa6\x79\xca\x35\x37\xd0\xd7\x29\xa3\xb2\x02\xbf\x63\x11\xa7\xe4\xaf\x25\x93\x88\xb2\x5a\x33\xa2\x5c\x5c\x27\xce\xff\x04\x64\x0c\x70\x33\x29\xcf\xc1\x74\xbe\x8e\x08\x29\x84\x82\x95\x89\x7c\x43\x24\x4b\xa9\xe6\x5f\xd8\x1d\x40\xf7\x3c\xd2\xc9\x86\xcc\xa6\xd3\x1f\xf1\x31\xa3\x4f\x93\x1a\x74\xbb\x9e\x16\x4f\x15\x4c\xee\x38\x90\x4d\x09\x2d\xb5\x40\x48\x41\xa3\x08\x14\x83\xa0\x79\x3d\x68\x2b\x9e\x26\x8a\xff\x61\xa0\x5b\x21\xc1\x5e\x27\x00\xba\x23\x7f\xa2\x18\x69\x99\xe5\x63\xfb\x43\x19\x69\x06\xf3\xc6\xa9\xa0\x7a\x43\x52\x16\xeb\xd3\xdc\x50\xef\xb0\x9d\x24\x62\x5f\xc0\xa8\x15\x49\x41\x38\xe3\x1b\x34\x27\xcb\x29\x08\x83\x2b\xff\xff\xcc\xec\xd7\x4f\x19\xcf\xed\x6a\x0c\xee\x99\x99\x7a\xa0\x97\x46\x96\x9b\x95\x11\xa5\xbb\x3c\x9c\xf1\xf4\x9c\xab\x95\x7f\x4e\x83\xfb\xc6\x9c\xb0\x7c\x98\xc3\x0c\xa8\xb4\xd4\xfc\x54\xf5\xc8\x4a\xfd\x13\x54\x0d\x2c\xa3\x3f\x7c\x13\x73\xa9\xf4\x24\x4c\x78\x1a\x75\x49\xbb\x70\x17\x1b\xb3\x30\x1c\x2f\x72\xd6\x9d\xb9\x7d\x54\xc4\xf1\xf9\x6a\xc5\x5e\x06\xeb\xe6\xf3\xc2\x8a\xa4\xf7\xe2\x2c\xe2\xd9\x22\x58\x34\x9f\x86\x38\x81\x78\xe7\x27\x6f\x88\xe7\x73\xa0\xe8\x7c\x0c\x71\x0c\x71\xef\xc4\xd4\x0d\xf1\x62\x1a\x1c\x8b\x1d\x83\x27\x9c\x45\x7c\xeb\x10\x5b\xf1\xa7\xf3\x14\x76\xe3\x10\x5b\xb1\x2f\x2c\x3f\x63\xcd\xab\xb5\x43\x6c\x86\x09\xe5\x0c\xe2\xf5\xca\x21\x76\xce\x4f\x6e\x74\x43\xfc\x62\x19\x4c\xbb\x82\x57\x5b\x75\x4a\xe8\x0e\xf1\xcd\xdc\x25\x76\x7a\x72\xd1\x0d\xf1\xed\xcc\x65\x24\x7b\x96\x9e\xd8\xac\xaf\xbd\xa8\xe2\xb7\xf8\x89\x4e\xb8\x8c\x6a\x36\x3e\x1e\x4e\x5b\x01\x13\xaf\x88\x95\x93\xfa\x84\xd2\xed\xd4\x09\x4d\x63\xf7\xcc\x03\x63\x31\x24\x10\x78\x3e\xc4\xb1\x82\xe2\x00\x43\x0c\xb2\x30\x4f\x93\xed\x61\x32\xf4\x5c\x07\x42\x75\x79\xf7\xb4\x71\xe3\xf2\xdf\x96\x45\xeb\xc9\x63\x1f\xc2\xcf\x7b\xf6\xc2\xb1\x73\x1d\x16\x1d\x47\x1f\x72\x77\xc5\x80\x01\xf7\xf9\xda\xe1\x48\x2d\x8b\x4e\x24\x18\x7b\x31\xca\x27\xfa\x62\xe9\xd8\xf3\x0e\x8b\x36\x52\x0c\x99\xbb\x62\xc8\x80\xf9\xd2\x15\xf6\x5a\x16\x6d\x24\x19\xfb\x10\x7e\x9d\xaf\x5c\x51\xb1\xc3\xa2\xe3\x73\x43\xee\x2e\x77\x1c\x70\x5f\xbb\x5c\xa1\x65\xd1\x8d\x44\x63\x3f\xca\xcb\xdd\x15\x57\x5b\x16\x9d\x50\x35\xf6\x62\xbc\x4a\x7f\x71\xe3\x08\x60\x1d\x8b\xf3\xa9\xc5\x19\xe3\x86\x4e\xe4\x0a\xcb\x9d\xb5\xa7\x7e\xad\x3b\xa3\xe0\x80\xfd\xad\x3b\x86\x74\xfd\xbc\x17\xc8\x8e\xc3\x40\x0f\xad\xfe\x77\x7b\x3f\x8a\x76\xc7\xd1\xa0\x8f\x57\xe7\xec\xee\xb1\x98\x9d\x98\xe8\x58\x44\x07\xab\xbc\xb6\x0f\x6c\xff\xac\x8a\xf2\x6e\x63\x70\xf1\x32\xdc\x5b\x9b\x83\xa0\xef\x3f\x7c\xbe\x1f\x25\x3a\x4b\x09\x57\x04\xd6\x8b\xcd\xde\x7a\x1e\xac\x7e\x24\x4a\x60\x6d\xa9\x09\x4d\x53\xd3\xf4\x7d\xbc\x7f\x47\x32\x46\x55\x29\x59\x86\xcd\x19\x00\xa5\x28\x77\x89\x28\x75\xd3\x6d\x8e\xa8\x64\x64\x0b\x6b\x8f\x08\xb4\x9e\x33\x2c\x4a\xab\x02\x3a\x20\x9f\x04\x22\x78\x08\xec\x0e\x64\x16\xac\x80\x0b\xf9\x99\xcc\x56\x30\x64\xf3\x0c\x85\x31\x42\x60\x95\x18\x43\x91\x8a\x75\x37\xdb\x54\x92\xa0\x5e\xb7\x22\x3a\x0c\x91\xc0\x84\x65\x77\x98\x9d\xc2\x52\x4a\x10\x09\x18\xb3\x4c\x91\x90\x62\xb7\x1a\x82\x74\x19\x08\x53\xee\x48\xc6\x15\x87\x86\x52\x16\x92\x69\x90\x05\x3a\x0f\x18\x05\xf2\x19\xa6\x20\x38\x2e\xa7\xca\x6c\x29\xc7\xad\x36\xf1\x02\xf9\xaf\xef\xec\x8c\xfb\x1a\x06\x05\x7c\x03\x8b\x69\xc6\xd3\xc3\x86\x5c\x7d\xa4\x29\xdb\xd3\xc3\xd5\x98\x5c\xfd\x82\xa5\x81\x86\x55\xbe\x67\x25\xeb\x01\x48\x0d\x69\x00\x63\xf2\x52\x72\x9a\x8e\x89\xa2\xb9\x82\x08\x28\x79\x8c\xac\xc1\xc6\x84\xdc\x90\x1f\xe6\xf3\xb9\x31\x28\xd3\x18\xb7\x8d\xe5\xc5\xed\xc6\x6b\x4c\xc9\x6c\x4c\x92\x39\xfc\x2d\xe0\x6f\x09\x7f\x2b\xf8\x5b\x9b\x4d\xab\x3d\x45\x8b\x02\x4b\xa0\x16\xb0\x15\xd0\x1b\x67\x90\x4e\x61\x47\x8e\x54\xbd\x98\x9a\x6a\x29\x99\x81\xab\x75\x76\x1d\x6a\x41\x1c\x3d\xdc\x2a\x50\x16\xf4\x6d\x1a\x76\x79\xa2\xaa\x3e\x7a\x43\x26\xc1\x0c\x87\x02\x8f\x79\x9f\xc7\x22\x58\x3b\x79\xac\xee\x7c\x3c\x50\x90\xc5\x90\x89\x53\x90\x85\x5f\x10\x64\xb2\xec\x33\x99\x07\x4b\x27\x13\xa7\x24\xd3\x1b\xcb\x65\xd5\xe7\x32\x0b\x6e\x5c\x5c\x56\x4e\x51\xa6\x2b\xcb\x65\x3d\xe4\xb2\x72\x71\x59\x3b\xb8\x4c\x6d\x13\xfc\x6b\xa7\x09\x2d\x12\x3c\xc9\xd0\xdf\x6e\x43\x87\x5b\xba\xaa\x35\x89\x81\x7f\xb8\x55\xcb\x60\xde\xe2\x8e\x76\x60\xdd\xe2\x96\xee\xdd\x31\xb8\x95\x5b\xe9\x06\xe7\x51\x82\xe9\xb8\x0b\x87\xf5\x5a\x9f\xac\xce\x77\x2e\xee\x79\x5e\x77\xa4\x46\x76\x1b\x4d\x66\xf7\x2f\xef\xdf\xbc\x42\xe1\xe9\x26\x11\x5f\xea\x03\x00\x8b\x9d\x3e\xbc\x9c\xbe\xbe\x6f\x96\x66\x8f\xac\x2e\xbe\x0e\xef\xe2\x82\xad\x11\x71\x3c\xb2\xdf\x3c\x2f\x4a\xfd\x4f\x7d\x28\xd8\xcf\x57\xaa\xdc\x66\x5c\x5f\xfd\xd6\x87\x4a\x06\xd9\x6d\x08\xac\xc8\xaf\x7e\x33\xda\x88\xb8\x2a\x52\x0a\x01\x9d\xe7\xc6\x05\xb6\xa9\x08\x1f\x31\x2c\x59\x5f\x58\xdc\x54\xa7\x4c\x9d\x83\xa7\x45\x7d\xf0\x64\x35\xb9\x5a\xad\xf0\x51\xb3\x27\x3d\xa1\x29\xdf\xe5\x1b\x12\x32\xcc\x3d\x77\x83\xdc\x35\xab\xe8\x7a\x11\x6f\x5d\x25\x97\x9e\xff\xd9\x39\x87\x1e\x58\x45\x14\x3b\x95\x96\x90\x43\x62\x21\x21\x96\x96\x45\xc1\x64\x08\x29\xb8\x41\x46\x2c\x14\x92\x56\x87\x6e\x39\xd4\x32\xe6\xc0\x2d\xe1\x9a\x19\x6e\x0c\x81\x7b\x48\x2e\xe6\xc4\x8b\x86\x8f\x3b\xc8\xea\x79\x34\xa9\x57\x64\x38\x17\x14\xd3\x6b\x75\x24\x66\x8e\xc1\x24\x8d\x78\xa9\xc0\x3b\xed\xb1\x1b\x42\x61\x55\x98\xef\x45\xca\x23\xf2\xc3\x76\xbb\x35\x7a\x29\xa5\x42\x36\x85\xe0\x56\x0b\x27\xce\xe8\xaa\xed\xa8\x0c\xd4\xee\xad\x7d\x72\xed\xb0\x0b\x57\xef\xb3\x0b\x65\x77\xdb\xe2\xec\x7c\xb1\x08\x4b\xd5\xcc\x57\x3f\x39\xe7\x73\xe0\xec\x7c\x0e\x54\x33\x9f\xc1\xf5\x3c\x0e\xca\xd0\x8e\x3a\x2d\xf4\xe6\xe6\x06\xa1\x50\x51\xa1\x0d\xd4\xd1\xa6\x96\xb2\xfe\x9a\x14\x92\x43\x40\x3a\x58\x71\x8f\xc0\x2e\xb9\x4f\x0e\xaa\x17\x70\x72\x8c\x5d\xc9\x60\x50\x6f\x49\x0f\x0f\x0f\x6e\x13\x82\xc5\xbe\x5e\x3c\x4c\x1d\xeb\xad\x11\xde\x45\xf6\x2d\xc1\x83\x3c\x63\xc1\x27\xec\xe4\x8c\x91\x9e\xc5\x0f\xac\x68\x88\xed\x19\x95\x07\x79\x8e\xe8\x7e\x93\x3b\x63\xa4\x4f\xf4\x63\x83\xf4\xef\x5e\x9d\x3a\x8e\x77\xaf\xcd\x29\xf6\x6c\x3a\xfb\x9e\x73\x46\x57\x2f\x2c\xa3\x3c\x1d\x26\x83\xbc\xcc\xb6\x4c\x0e\xa1\x8a\x51\x19\x26\x43\x28\x06\xd6\x63\xd8\x11\xcb\x52\x1e\x81\x0a\xaa\xd4\x1e\x14\x89\x70\xe4\x02\xa1\x95\x8e\x47\x0a\xda\x90\x50\x57\xd5\x91\x2f\xdb\xac\x21\xba\x62\x4b\x65\x5a\x9e\xcf\xd0\x91\x21\x00\x4c\x50\xd7\x4d\x55\x95\x68\x94\x09\xfa\xd8\xde\x3c\x3c\x8c\x09\x64\x20\x21\xa1\x1b\xdb\x1e\xc8\x3f\xd8\xf6\x91\xd7\x7d\x8e\x63\x9b\xe3\x38\xf6\x84\xf2\x37\x33\xfc\x77\x2a\xfa\x43\x40\x4f\x68\x24\xf6\x6d\x7e\xf1\x07\x79\x90\xfd\x23\xcb\xc0\x75\x14\xa1\xfb\xc7\x3d\x95\x11\x89\x58\x4c\xcb\x54\x13\x65\xfa\x65\x14\x5d\x61\xeb\x66\xb4\xa6\x20\x45\x4a\xc2\x3f\x7c\xfa\xce\xb7\xd0\x6c\x1d\x64\x72\xd4\xf1\x84\x42\x0e\xa6\x90\x38\xab\xe4\x5a\x69\x04\x3e\x93\x4c\xfc\xe1\xc3\x99\xcf\x11\x0e\xf4\xd5\x9b\x00\xeb\x69\x6b\x1f\xeb\x55\xcf\x3e\xaa\xfa\x74\xdd\x87\xd9\x06\x0b\xc1\xc0\xcb\xa1\x3f\x57\xd4\xb0\x5a\x74\x86\xa9\x5a\x97\x2e\x5c\xa5\x51\x37\xc6\x33\x95\xd1\xae\x0b\xd1\xea\xd8\x62\xad\x22\xec\x73\xe5\x32\x9d\x50\xe6\x30\xdd\x36\xf3\xf4\x73\x6a\x4a\xb7\x2c\x1d\x8f\x52\xb6\x63\x79\xd4\x2f\xfe\x9a\xaa\x6f\xd0\xa3\x56\x7d\x81\xab\x64\x03\x7e\x31\x67\x69\x84\xa7\x27\x5f\x7b\xe5\x61\xc7\x67\xea\x2e\x68\x3a\xdc\x86\x30\x61\xe1\x23\xb8\xc6\x51\xb9\x0a\x3e\x26\xdc\x85\x69\xb3\x00\xf2\x7f\x24\x30\x3f\x26\xcd\xf1\x88\xb7\x86\xed\x1d\x56\xb9\xd7\x02\x91\x22\xa3\x69\xa7\xc1\x51\xdf\xed\x1b\x78\x0c\x06\x65\x75\x5a\x94\x82\x98\x13\x13\x39\xa0\xd4\xe6\x32\x4c\x31\x70\x28\x1e\x19\x35\x89\xe3\x31\x50\x0d\x43\x0e\x4c\x7b\x83\xc6\xa4\x66\x66\xfd\xc6\xbe\xad\x71\x75\x7f\x30\xb4\x04\xdb\x81\x2f\x20\xc4\x19\xea\xaf\x9a\x85\x7d\x4f\x5d\x1f\x71\x4d\xed\x8f\x45\x57\xe5\x55\x13\x70\x5b\xbd\x76\x4d\x79\xb7\xcd\xb4\xf6\x66\x4f\x0c\xcc\x5e\x98\xab\x02\x17\xd7\xba\x77\x2b\x42\x10\xaf\x6f\xfa\xa6\x6b\x6f\xed\xac\x7d\x79\x1f\xcc\xdd\x7a\x38\xa7\x27\xc1\x32\x65\x86\xff\x3c\x89\xea\x7e\x86\xff\x3c\x89\x0a\x54\x59\x48\x06\x2e\xd3\x08\x7b\xec\xf2\x8d\xf8\xa8\x7c\x7b\x14\x30\x94\x0c\xb8\xb4\x07\x73\xd5\xdd\x8d\x8b\x6f\x82\x77\x67\x74\x02\x81\x33\xea\xef\xcd\x6c\x8e\x85\x44\x9d\x3d\xba\x9d\x6a\x7b\xef\xa1\x4e\xd8\xb5\x25\x1e\x69\x18\xd3\x52\xd2\x7f\xeb\xaf\xa3\xa3\xb7\xfd\x43\x67\xaa\xa8\x52\xda\x23\x6a\x1f\x7b\x34\xb2\x0a\x49\xed\x71\x8b\xbd\x1c\x73\x71\xad\x7a\x55\x6d\x8f\x22\xea\x5a\xfb\x94\x4f\x9b\x40\x7f\x5c\xfe\x8d\xfb\x79\x64\x48\xdc\x9c\x4c\x81\x05\x42\x63\x81\x46\xfb\x7b\x29\x34\xfc\x8e\x52\x24\xdd\x95\x08\xd7\x68\x92\xe3\x51\x81\x11\x0a\xe3\x12\x60\x20\xae\xbb\x18\xce\x1b\x86\x46\xc3\xed\xa5\xa3\x8b\xab\xd3\xab\xe3\xa0\x9c\xc4\x65\x9a\x56\xd9\xd4\x75\x83\xe7\xc4\xe1\x42\x39\xc1\x6b\x45\x03\xf2\xce\x4d\xa3\xb3\x58\x14\x48\x5e\xdd\xa2\x33\x6f\x23\xaa\x0b\x43\x06\xd0\x1d\x81\x26\xdf\x1d\x60\x3c\xcb\x2a\xfa\x1d\x57\xe1\x77\xac\xe3\x44\x1e\x9d\x7a\x2e\x7a\x61\xbc\x31\xa0\x45\x13\x20\x87\x45\x4e\x0b\x32\xf4\xce\x08\x52\x25\x36\x7b\xbb\xed\xe2\x4b\xf7\xea\xc3\x44\x1f\x96\xc6\x8d\xac\xe4\xad\x10\x51\xce\x94\xea\x5f\x6f\xdb\xd0\x58\x9b\xf3\x00\x09\x8d\x90\xfd\x5d\x4e\xc2\xb8\xee\xb3\xcd\x95\xc6\x0d\xb9\xba\xba\xeb\xe6\x1e\xe3\xae\xe6\xa4\x0c\xb9\xa3\xc5\xe9\xa4\xb5\x94\xde\x85\xbf\x8b\x6b\xc3\xab\xa2\xeb\xe7\xa3\xf7\x10\x88\x36\xa6\x2b\xdd\x32\xa5\xc9\x9e\x1e\xf0\x4d\xa2\xd2\xb2\x0c\x35\xc4\x25\xf3\x06\x11\x5f\xcb\x89\x98\x54\xef\x06\x7e\xaf\x56\x85\xaf\x1d\x61\x60\x08\x61\x50\x57\xa3\x6a\xc4\x28\x07\x7d\x18\x80\xc4\xf7\xd4\x34\xd7\x26\x73\x07\xe6\x0e\x1c\x7b\xa2\x59\x01\x61\x8e\xf0\x98\x1c\x44\x09\xd3\x81\x76\x23\xc3\x28\xa1\xf9\xae\x62\x54\xb7\x93\xd8\x3f\x6e\xeb\x13\x6d\x6c\x2d\x33\x7c\xa1\x59\x5f\xa2\x1b\x43\xc2\x51\xf5\xbc\x99\xd8\xf2\xb4\x9a\xfe\x50\x15\x09\x65\x01\x55\xa2\xc1\xd5\xf4\x44\x41\x94\xe6\xc0\x84\x42\xd7\x60\xd8\x13\xae\x21\xa9\x31\x89\xf7\x60\xeb\xfb\x98\xdd\x57\x22\x35\xcf\xd3\x97\x01\xfd\x6f\x52\x7e\xa2\xa9\x12\xd5\xd9\x29\xd4\x20\x2c\x27\x3b\x09\x3e\xb4\x65\x21\xb4\xc7\xd0\x3e\x87\x78\x7d\xf2\xd9\x37\xde\xb7\x1c\x73\xd7\xa7\x5e\xd3\xbc\xf0\x92\x45\x4c\x3d\x82\x37\x7b\xe8\x20\x7a\xfa\x08\xdf\xd4\x84\xbf\xbc\xf1\xd1\xce\x1b\xda\xff\x06\x00\x00\xff\xff\xf0\x1b\x32\x3e\xbc\x2c\x00\x00")

func assetsSkeletonCssBytes() ([]byte, error) {
	return bindataRead(
		_assetsSkeletonCss,
		"assets/skeleton.css",
	)
}

func assetsSkeletonCss() (*asset, error) {
	bytes, err := assetsSkeletonCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/skeleton.css", size: 11452, mode: os.FileMode(420), modTime: time.Unix(1445220472, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _assetsWebsocketJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x55\x51\x6b\x1c\x37\x10\x7e\xde\xfb\x15\xd3\x6b\xe0\xf6\xc8\x45\xee\xf3\xb9\x2e\x34\xb8\xc5\x2d\xb5\x1d\x70\xa0\x94\x10\x8c\x6e\x35\xbe\x55\xbd\x92\x8c\xa4\x3d\x7b\x29\xf7\xdf\x3b\x23\x69\x2f\x1b\x1f\x09\x29\xa5\xd0\x87\xf8\xe1\x2c\xcd\x68\xbe\x99\xf9\xe6\x93\x76\x27\x3d\x04\xd7\xdc\xc3\x19\x58\x7c\x84\xdf\x71\x73\x43\x3b\x8c\xf5\x63\xe8\xbd\x5e\x9e\xce\x66\xec\x15\xce\xba\x07\xb4\x74\xe8\xae\xb7\x4d\xd4\xce\xd6\x4b\xf8\x0b\x1a\x67\x83\xeb\x50\x74\x6e\x5b\xcf\x69\x63\xb1\x89\xa8\x20\x3a\x98\xc3\x4b\x28\x08\xb0\x1f\x21\x9a\xce\x05\x9c\x62\xe0\xa7\x40\xc8\x09\xe9\xb4\x82\x9a\xa1\x50\x34\x4e\x21\x2d\xe6\xcb\xf9\x14\xd1\x60\x08\x72\x7b\x84\x39\xab\xb8\xad\xa6\x95\x3e\xde\x46\x6d\xd8\xff\xa2\x5e\x7c\xcb\xcb\xdb\x0d\xda\xa6\x35\xd2\xdf\x2f\x96\xa2\xd5\xdb\x36\x9d\x0a\x35\x75\x3a\x09\x32\x68\x4a\x0c\xad\x3e\x17\x32\xab\x4c\xd8\xd2\xd1\x5f\x6f\xae\xaf\xc4\x83\xf4\x01\x6b\x14\x4a\x46\xc9\xcc\x55\xfa\x0e\x6a\xf2\x8b\x38\x3c\x50\x0d\x67\x67\xb0\x08\x51\xc6\x3e\x2c\x52\x8d\x95\x72\x4d\x6f\xd0\x46\xb1\xc5\xf8\x53\x87\xbc\x7c\x3d\xfc\xa2\xea\xc3\x29\xa1\x89\x0d\x7f\xf1\xf6\xf2\x37\x4a\xc1\x40\xd9\x41\x79\x0f\xd0\xd9\x92\xc0\xe7\x3f\x6b\xab\x43\x8b\x6a\x9e\xe1\x3f\x8d\xdf\x38\x63\x74\xbc\xdd\x74\xc4\x22\x65\x09\x71\x20\xfe\x77\x3a\xe8\x8d\xee\x74\x1c\x28\xd9\xbc\xd5\x4a\xa1\x9d\x73\xaa\x3d\x60\x47\x73\xcb\x90\x32\x32\x9b\xc6\x31\x58\x2a\x20\x63\x71\xcf\xb8\x14\x77\xce\x1b\x19\xeb\xc5\x1f\xf4\xf7\xea\xf2\xf2\xd5\xf9\x39\x5c\x5c\xac\x8d\x59\x07\xea\xe6\xf4\x0b\x6a\x7a\xd6\x73\x4a\xf7\x12\x16\x50\x2f\xe8\xdf\x24\x5d\xe8\x37\x7f\x92\x4e\xd8\xb7\x64\xd7\xe2\xfb\x8d\x87\x93\x1f\xd2\xf2\x42\x86\x76\x0d\xcf\x02\x5a\x32\xa6\x66\xe8\xa7\x34\x74\x3c\x1d\x8f\xa1\xef\x62\x99\x4e\xde\x14\xe2\xf3\x86\x67\x5a\xbd\x10\x28\x9b\xb6\xce\x16\x11\x30\xae\x3e\x88\xef\x1e\x87\x15\xec\x64\xd7\x63\x19\x01\x4b\x2a\xc9\x87\x70\x92\xfd\xdd\x77\xef\x4f\x8f\x89\x2c\x60\xff\x84\x4b\x06\x39\x39\x81\x24\xef\x24\x48\x26\x4b\xb2\x35\x25\xae\x74\x4c\x1a\xce\x9b\xca\x4a\x83\xeb\x44\xe7\x2a\x1b\x86\x75\xae\x4b\x5c\x85\x37\xe8\xaf\x1f\xb2\x79\x9f\x71\xab\x80\x5e\x23\xa9\x6a\x72\x87\x78\x60\x75\x09\x21\xb0\x3c\xcd\x24\xc3\x7c\x98\xaf\x32\xd5\x53\x22\x65\xe7\x51\xaa\x01\xf0\x49\x87\x18\x72\xca\xec\x12\x52\xa9\x37\x4e\x53\xd3\x5c\xe1\x0a\xa2\xef\x91\x18\x94\x34\x91\x82\x39\xea\x8d\xe1\xd2\xa3\x94\x03\x33\xc8\xa4\x20\x02\xba\x49\x9e\xba\x34\xc9\xa4\xca\x35\xbc\x63\xe0\xf7\xa5\xcf\x4a\xab\x43\xa7\x54\xf6\x68\xcd\x7c\x1c\xd9\xf7\x63\x09\xb3\xf4\x53\x38\xa6\x37\xc0\xf9\xe1\x5f\xb1\xfc\x63\x47\x77\x0d\xd5\xeb\x21\xe2\x17\xf0\x4d\x09\xff\x57\x74\x73\x3d\xff\x3d\xdb\xd9\x30\x4b\x94\x5f\xb9\xc7\x15\x10\xb2\xe7\x6b\xe2\x76\xe8\xe1\xde\xba\x47\x5b\x6a\x5b\x81\xb4\x0a\xb4\xa5\x5d\x04\xdb\x77\x5d\xc8\x51\x44\x4f\xba\x64\x44\x88\x47\x30\x3a\x04\x6d\xb7\xe2\xc3\x9d\x9d\x88\x67\x04\x3a\x5c\x5d\x6d\x15\x3e\xad\x72\x82\x72\x79\xef\x5c\x4f\x69\xce\x32\x5b\xa9\xd4\x82\x93\x0e\xa5\x67\x7e\x0a\xa0\x9e\x52\xc9\xa6\x44\xa7\x61\xf1\x5e\x70\xfb\xf4\xc0\x24\x61\x8c\xce\x03\x38\x0f\x24\xd3\x40\x2f\x4e\xec\xbd\x9d\xa4\x2b\xcc\x64\x9a\x18\xee\x9b\x14\x35\x62\xe4\x32\x0e\x03\xfe\x8c\x08\x99\xa3\xa2\xb8\x8f\x25\x30\xcb\x22\x9c\x52\xff\xb6\x45\x08\x5c\x31\x3d\x40\x1f\x49\xbf\xb8\xaf\xcf\xaf\xd7\xf4\x6e\xed\x90\xbf\xf3\x01\xe9\xa3\x27\xf9\xa3\x3f\xf2\xf0\x9c\x6d\xd6\xce\x57\xb2\x8f\xc9\xa6\x74\xfb\xd9\xdf\x01\x00\x00\xff\xff\x3d\xa6\xc4\xfa\x7e\x09\x00\x00")

func assetsWebsocketJsBytes() ([]byte, error) {
	return bindataRead(
		_assetsWebsocketJs,
		"assets/websocket.js",
	)
}

func assetsWebsocketJs() (*asset, error) {
	bytes, err := assetsWebsocketJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/websocket.js", size: 2430, mode: os.FileMode(420), modTime: time.Unix(1445217722, 0)}
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
	"assets/.main.css.swp": assetsMainCssSwp,
	"assets/index.html": assetsIndexHtml,
	"assets/main.css": assetsMainCss,
	"assets/skeleton.css": assetsSkeletonCss,
	"assets/websocket.js": assetsWebsocketJs,
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
	"assets": &bintree{nil, map[string]*bintree{
		".main.css.swp": &bintree{assetsMainCssSwp, map[string]*bintree{
		}},
		"index.html": &bintree{assetsIndexHtml, map[string]*bintree{
		}},
		"main.css": &bintree{assetsMainCss, map[string]*bintree{
		}},
		"skeleton.css": &bintree{assetsSkeletonCss, map[string]*bintree{
		}},
		"websocket.js": &bintree{assetsWebsocketJs, map[string]*bintree{
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


func assetFS() *assetfs.AssetFS {
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, Prefix: k}
	}
	panic("unreachable")
}
