// Code generated by go-bindata.
// sources:
// templates/directory-list.go.html
// templates/error.go.html
// templates/layout.go.html
// templates/root-list.go.html
// DO NOT EDIT!

package main

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

var _directoryListGoHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x55\xdd\x6f\xdb\x36\x10\x7f\xf7\x5f\x71\xe0\xc3\xde\x28\x29\x5d\xb3\x0d\x8a\xec\x62\x6d\x37\xac\x40\x52\x14\x6b\xde\x07\xda\x3c\x59\x44\x69\x52\xa0\x68\x27\xae\xa7\xff\x7d\x47\x59\x52\xf5\x61\x07\x59\x1e\xc2\xd3\x7d\xdf\xef\x3e\x7c\x3a\x49\xcc\x95\x41\x60\x8f\x8f\xca\x6b\x64\x75\xbd\x38\x9d\xa2\x07\xb1\x29\x88\x5b\xd7\xf0\xd3\xd6\xdf\x01\x71\xbe\x08\x5f\x34\x32\x34\x92\xde\xc5\xc0\xf0\x83\x35\x1e\x8d\x0f\xa6\x59\xe5\x8f\x1a\xc1\x1f\x4b\x5c\x32\x8f\xcf\x3e\xde\x54\x15\x5b\x2d\x60\xaf\xa3\x8d\xdb\xef\xd6\x70\x82\x52\x48\xa9\xcc\x96\x6b\xcc\x7d\x0a\xc9\x1d\xd4\x0b\xd0\xaa\x17\x4b\x55\x95\x5a\x1c\x53\x50\x46\x93\xff\x3b\x92\x55\x9e\x37\x7e\x79\xf0\x9b\x82\xb1\x81\x3d\xb0\x4a\x45\xee\xd1\x91\xed\xe6\x9c\x49\x0a\x0c\x62\x60\x63\x1d\x2d\xc8\x0d\x55\xa5\x65\xba\xc6\xdc\x3a\x1c\xe9\x5f\x57\x9e\xfb\x6e\x74\xb3\xb8\x49\x69\xb5\x58\x64\x52\x1d\x60\x43\x16\xd5\x92\x39\xfb\x04\x4f\xa8\x35\x83\x46\xba\x64\x3b\xe1\xb6\xca\x70\x6f\xcb\xf4\xe6\xb6\x7c\x0e\x50\xc0\xd0\x60\x63\x35\x7f\xae\xf8\x6f\x8d\x80\x44\xc5\xcf\x67\x82\xc8\xbd\xee\x95\x42\x4e\xac\x13\x00\xb5\xc3\x09\xb3\x45\x88\xde\x3b\x14\xf2\x43\x90\x56\x04\x3e\xf4\x7f\x99\x56\xd7\x6c\x1b\xb1\x80\xc2\x61\xbe\x64\xd4\xd7\x7b\x65\xbe\xd5\x35\x5b\x11\xf9\x5e\x54\x68\xc4\x8e\xba\x9e\xc5\x62\x68\x91\xc5\x5a\x0d\xa3\x9f\x47\xa0\x93\xed\x75\x9b\x7b\xdc\x25\xdf\xfa\xfa\x64\x72\x1b\x3d\x58\xf9\xa8\x76\x18\xfd\x69\xdd\x4e\x78\x60\x6f\x92\xe4\x17\x9e\xdc\xf0\xe4\x0d\xdc\xdc\xa6\xc9\xdb\x34\xb9\x05\x9e\xfc\x9a\x24\xf0\xf0\xf5\x91\xb5\x6e\x27\x0e\xb0\x65\x07\xe0\x56\xfd\x2c\x66\x71\xf8\x0c\x78\xf6\xc4\x1c\xd8\xb7\x10\xa6\x90\x3b\xb5\x2d\x7c\x87\xf1\x40\x6b\xed\x0d\xdf\x3a\xbb\x2f\xf9\x01\x9d\x57\x1b\x41\x9d\x73\x36\x34\xae\xe1\xf6\xb8\x11\x62\x3f\x2c\x20\x58\x29\xca\x8d\xb5\x30\xbe\x73\xcb\x9b\x01\xc4\x59\x0f\x7e\x2e\x20\x17\x3c\x0c\x30\x5b\x65\xf1\x00\xc3\xbf\x71\xb3\x77\x95\x3a\x20\xdc\x93\xb0\x87\x52\x5c\x8f\x27\x43\xc7\x5d\x1f\x31\xdf\xf9\xe5\x77\x55\xbe\x10\x35\x57\xb4\x30\xc2\xd1\x10\x1f\x90\xdb\x49\xfc\xdf\xcf\xfc\xd7\x04\x2e\x9d\xa2\x21\x3e\x76\x91\xe3\x17\x42\x16\x76\x87\x93\x40\x7f\x11\xeb\x35\x51\x9e\x84\x33\x74\x14\xba\x28\x51\xf4\x52\x1c\x8d\x07\xd4\x3c\xb4\x67\x14\xeb\x8b\x70\xb4\xa0\xd3\x68\xd3\x31\xe9\x9e\xf1\xd6\x9e\x17\xd3\x8b\x35\x5d\xaf\x96\x7b\xfe\x68\xfe\xd3\xf9\x71\xaa\x44\xd9\x7e\x15\x96\xa6\xa5\xa5\xd7\xd6\x49\x74\x28\xbb\xe1\xf2\x05\xad\x64\x5f\xa7\x77\x83\x2a\x7c\xb1\xfa\x4c\xeb\x95\xc5\x44\x8c\xb8\x34\xe2\x2a\x57\x28\x67\x92\x3e\x95\xc1\x0c\x7f\x55\xdf\x2f\xbb\xb8\xc0\xfd\x78\x3f\xe4\x11\xed\x3a\x4c\x06\x69\x66\x7e\x6d\xe5\xb1\x53\xea\x6f\xcb\x1f\x86\x6a\xc6\xe1\x5d\xa1\x6a\xba\x84\x4e\x27\x95\x43\xf4\xa9\xfa\xa8\x5c\x5d\x87\x55\x68\x8f\x02\x03\x29\xbc\xe0\x25\x2d\x68\x50\x0a\x2f\x4d\x2a\x44\xe1\xc2\x0c\x0f\x8a\x97\x57\x2f\xd2\xc8\x86\x96\xfd\xf3\xe5\x8b\x34\xf6\x10\x1c\x86\x9f\xad\xff\x79\x6c\xa6\x6e\x46\x65\x8d\x13\xf4\xf2\x52\x33\xe6\x0e\x50\x57\xf8\x3a\x5b\xca\x37\xf4\x12\xfe\x85\x8a\x1e\x2a\xf9\x52\x3e\xc3\x53\x3b\x2e\x14\xe7\xea\x33\x58\x9b\x7a\x8c\xf5\x97\x6b\xba\x06\xfb\x3b\xa9\x47\xf7\xac\xd7\x9e\xac\xa1\xb4\x4f\x46\x5b\x21\x27\x6b\xd8\x75\x48\x4c\x73\x99\xd6\x72\xb5\xed\xb4\x5a\x6e\x8b\x7e\xc9\xfe\x59\x6b\x61\xbe\xcd\x52\x99\x26\x42\xa0\xa2\x33\x42\xd3\xa5\x0d\xda\xb3\x6c\x5e\x1c\x9e\x1f\x5b\x31\xce\x91\xf8\xdd\x5e\x10\x19\x76\xbd\xbf\x1d\x9d\xda\x7f\x01\x00\x00\xff\xff\x9f\x75\x0f\xb4\x3d\x09\x00\x00")

func directoryListGoHtmlBytes() ([]byte, error) {
	return bindataRead(
		_directoryListGoHtml,
		"directory-list.go.html",
	)
}

func directoryListGoHtml() (*asset, error) {
	bytes, err := directoryListGoHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "directory-list.go.html", size: 2365, mode: os.FileMode(436), modTime: time.Unix(1449454454, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _errorGoHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x8f\x41\xcb\x82\x40\x10\x86\xef\xfe\x8a\x61\x3f\xbe\xab\x12\x78\x0a\xf1\x12\x1d\xbb\x84\x7f\x40\xf0\xb5\x84\x5a\x75\x9c\x8a\x98\xf6\xbf\x37\x6e\x04\x79\xf0\xb2\xec\xfb\xf2\x3c\x33\x8c\x6a\x83\xb6\xf3\x20\x57\x55\x9d\x5c\xe0\x42\x48\x88\xf6\xcc\x3d\x27\xaa\xf0\x8d\xe5\xe4\x07\xda\xf5\x5e\xe0\x65\xc6\x8a\xf3\xa6\x8c\x60\x91\xd9\xcf\x62\x5e\x1e\x30\x4d\xf5\x09\x5b\x6b\x72\x6b\x06\x46\xa9\x9a\x82\x39\x8d\x60\x08\x45\x36\x77\x91\x3d\xd6\x8f\xcf\x9e\x25\x6d\xaf\xc8\xb3\xbd\x0a\xcd\x1e\xbd\x68\xe0\xce\x4b\x4b\xee\xff\x8f\xee\x6e\x39\x01\xe3\x0d\x93\xac\xfa\x8c\x71\xd5\xff\xde\xf6\x0e\x00\x00\xff\xff\xd6\x44\xe1\x7f\x00\x01\x00\x00")

func errorGoHtmlBytes() ([]byte, error) {
	return bindataRead(
		_errorGoHtml,
		"error.go.html",
	)
}

func errorGoHtml() (*asset, error) {
	bytes, err := errorGoHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "error.go.html", size: 256, mode: os.FileMode(436), modTime: time.Unix(1448692517, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _layoutGoHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x54\xdd\x4e\xe3\x3c\x10\xbd\xff\x9e\xc2\x9f\x6f\x51\xe3\x86\x50\x5a\x56\x29\x52\xa1\xd9\x45\x15\xa5\x85\x56\xe9\xc2\x9d\x9b\x3a\x89\x5d\xff\x24\xb6\x43\x13\x10\xef\xbe\x49\x5b\xa0\x8b\xf6\x47\x2b\xed\x5e\x65\x7c\xc6\x9e\x33\x47\x33\x27\xfe\xff\xc3\xc9\xe5\xfc\x7e\x1a\x80\xd4\x0a\x7e\xfe\x9f\xbf\xfb\x00\xe0\xa7\x04\xaf\x9a\xa0\x0e\x2d\xb5\x9c\xec\x62\x00\x9e\x9f\x2d\x11\x19\xc7\x96\x00\x38\x9f\x37\x19\x08\x9c\x97\x97\xdd\x4d\x74\x70\xd5\x37\x91\xa6\x99\x05\x46\x47\x7d\x88\x10\x66\xb8\x74\x12\xa5\x12\x4e\x70\x46\x8d\x13\x29\xb1\xc5\x10\xa7\x4b\x83\x58\x5e\x10\x5d\x21\xd7\x71\x5d\xa7\xbd\x3f\x39\x82\x4a\x87\x19\x78\xee\xa3\x5d\xa9\x7d\x5d\x4e\xe5\x1a\x68\xc2\xfb\xd0\xd8\x8a\x13\x93\x12\x62\x21\x48\x35\x89\xfb\x30\xb5\x36\x33\x9f\x10\x12\xb8\x8c\x56\xd2\x59\x2a\x65\x8d\xd5\x38\x6b\x0e\x0d\xe3\x1b\x80\x3c\xc7\x73\x4e\x51\x64\xcc\x3b\xb6\x25\xac\x11\x08\xa8\xb4\x24\xd1\xd4\x56\x35\x47\x8a\xbd\xde\x49\xcb\xcd\x7b\x62\x3e\x9a\x0c\x66\x65\x8f\xb9\x83\xe2\x08\x77\x16\xc3\x50\x4e\xe9\x31\x5f\x7f\x8e\x37\x9b\x60\x80\x7b\xe9\x70\xb8\x62\x0f\x3c\xbb\x26\x49\x99\xb2\x70\x1c\xb8\x71\xc2\x16\xd3\x2f\x62\xfd\x64\xba\x10\x44\x5a\x19\xa3\x34\x4d\xa8\xec\x43\x2c\x95\xac\x84\x2a\x6a\x75\xff\x58\x54\xcb\xa6\x44\x90\x5f\x49\x8b\xaf\x17\xc7\x37\x6d\x97\x8f\x73\x86\xd7\x17\xeb\xd2\xe3\x68\x7c\x16\xe0\xb4\xd8\x64\xb3\x98\xdc\x3c\x86\xa7\xde\xa8\x43\x9e\xa4\x57\x3c\x3c\xe1\x6c\xde\x2e\xba\xc1\xbd\xf9\x3a\x66\xb7\xe1\x51\x3b\x90\x1d\xfd\x3b\x69\x87\x7b\xf0\xa7\x52\xd8\xc7\xf1\xb0\x1f\x4a\x68\x8b\xd9\x72\x34\x0c\xae\x28\xe6\xb1\x28\x2e\x2e\x6e\xa7\xa7\x83\x93\x5b\x9d\xe9\xbc\x33\x09\xe3\x85\xd7\x9d\xde\xdd\x79\xac\x13\x5c\xe7\xa5\x31\x6e\x15\xe6\x13\x2b\x49\x26\xaf\xc2\xe9\x19\x1e\x75\xcb\xd9\xcf\x25\xfc\xc5\xdd\x8b\x95\xb4\x2d\xbc\x21\x46\x09\x82\x4e\x9c\x4e\xbd\xe8\xcd\xa4\x0e\xe1\xb7\x31\x6d\x2d\x88\x5e\x3d\xe8\x2f\xd5\xaa\xda\x77\xb0\xa2\x8f\x20\xe2\xd8\x98\x3e\x8c\xea\x97\x98\x4a\xa2\xe1\xab\x39\xbf\xcf\xf2\x56\x69\x5a\xae\xfb\x96\xfd\x60\xde\xcb\xfa\x39\x91\xf6\xdd\xbe\x0d\x65\x5d\x60\x4f\xf4\x1a\xfa\x68\xc7\x5e\xb7\xb3\xfd\x37\x7c\x0b\x00\x00\xff\xff\x9f\xbd\x22\xd4\x33\x04\x00\x00")

func layoutGoHtmlBytes() ([]byte, error) {
	return bindataRead(
		_layoutGoHtml,
		"layout.go.html",
	)
}

func layoutGoHtml() (*asset, error) {
	bytes, err := layoutGoHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "layout.go.html", size: 1075, mode: os.FileMode(436), modTime: time.Unix(1449453577, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _rootListGoHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\x90\xcd\x6a\xc3\x30\x10\x84\xcf\xf5\x53\x2c\x2a\xed\xcd\xb6\x1c\x9a\x16\x52\xd5\x97\x42\xa0\x50\xf7\xd0\xf8\x05\x94\x6a\x1d\x0b\x6c\x29\xd8\x4b\xc0\x08\xbf\x7b\x25\xff\x34\x25\xb7\xd1\xc7\xac\x76\x66\x9d\x53\x58\x69\x83\xc0\xca\x52\x53\x83\x6c\x1c\x23\x00\xe7\x92\x42\xfe\xd4\x9e\x8f\x23\x3c\x9e\xe8\x15\xbe\xad\x25\xf8\xd4\x3d\x45\xce\xa1\x51\xde\x15\xfd\x1b\x7d\xb7\x86\xd0\x50\x18\x16\x75\x96\xef\x75\x83\xfd\xd0\x13\xb6\xd3\x5c\x2f\x52\x0f\x23\xa1\xf4\x25\xf7\x9f\x0b\x92\xc7\x06\x83\x0a\xba\x46\xa9\x66\x1d\x5e\x5d\x1e\xdd\x79\x96\x7f\xc9\x16\x45\xea\xc5\xfc\x2c\xac\xd2\x95\x46\x35\xa3\xc5\x9c\x06\xf7\xa2\xae\xbf\x08\x3a\x5a\x35\xac\x26\xe7\x3a\x69\x4e\x08\x49\xc8\xf1\x61\x2a\xdb\x4f\xfd\x96\x65\x40\xa1\xf2\x1b\x73\xee\xdc\x69\x43\x15\xb0\x87\xfb\x0b\x83\xe4\x30\x78\x1b\x03\x25\x49\xc6\x67\x49\x75\x70\x24\x21\x92\xa7\x53\x22\x95\xff\x01\xbf\x5c\x5d\x99\x0f\x5a\xea\x16\x93\xbd\xed\x5a\x49\xc0\x36\x9c\x3f\xc7\x3c\x8b\xf9\x06\xb2\xed\x8e\x3f\xed\xf8\x16\x62\xfe\xc2\x39\x14\x87\x92\xad\xd3\xb7\x85\x42\xee\xf9\xc8\x0b\x5f\x2b\x79\x39\xdf\x4e\xa4\xd3\x31\x17\xdb\x6f\x00\x00\x00\xff\xff\xa9\xd6\xa7\x5b\xc5\x01\x00\x00")

func rootListGoHtmlBytes() ([]byte, error) {
	return bindataRead(
		_rootListGoHtml,
		"root-list.go.html",
	)
}

func rootListGoHtml() (*asset, error) {
	bytes, err := rootListGoHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "root-list.go.html", size: 453, mode: os.FileMode(436), modTime: time.Unix(1448692517, 0)}
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
	"directory-list.go.html": directoryListGoHtml,
	"error.go.html": errorGoHtml,
	"layout.go.html": layoutGoHtml,
	"root-list.go.html": rootListGoHtml,
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
	"directory-list.go.html": &bintree{directoryListGoHtml, map[string]*bintree{}},
	"error.go.html": &bintree{errorGoHtml, map[string]*bintree{}},
	"layout.go.html": &bintree{layoutGoHtml, map[string]*bintree{}},
	"root-list.go.html": &bintree{rootListGoHtml, map[string]*bintree{}},
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

