package s3uploader

import (
	"bytes"
	"errors"
)

type sig struct {
	bytes []byte
	t     string
}

var signatures = [...]sig{
	// image formats
	{bytes: []byte{0x89, 0x50, 0x4e, 0x47}, t: "image/png"},
	{bytes: []byte{0x42, 0x4d, 0x78, 0x78, 0x78, 0x78, 0x0, 0x0}, t: "image/bmp"},
	{bytes: []byte{0x47, 0x49, 0x46, 0x38}, t: "image/gif"},
	{bytes: []byte{0xff, 0xd8, 0xff}, t: "image/jpeg"},
	{bytes: []byte{0x4d, 0x4d, 0x0, 0x2a}, t: "image/tiff"},
	// video formats
	{bytes: []byte{0x66, 0x74, 0x79, 0x70, 0x69, 0x73, 0x6f, 0x6d}, t: "video/mp4"},
	{bytes: []byte{0x47, 0x3f, 0xff, 0x10}, t: "video/mpeg"},
	{bytes: []byte{0x6d, 0x64, 0x61, 0x74}, t: "video/mov"},
	{bytes: []byte{0x66, 0x74, 0x79, 0x70, 0x33, 0x67, 0x65}, t: "video/3gpp"},
	{bytes: []byte{0x4f, 0x67, 0x67, 0x53}, t: "video/ogg"},
	{bytes: []byte{0x6d, 0x64, 0x61, 0x74}, t: "video/quicktime"},
	{bytes: []byte{0x1a, 0x45, 0xdf, 0xa3}, t: "video/webm"},
	{bytes: []byte{0x46, 0x4c, 0x56}, t: "video/flv"},
	{bytes: []byte{0x52, 0x49, 0x46, 0x46}, t: "video/avi"},
	// audio formats
	{bytes: []byte{0xb, 0x77}, t: "audio/ac3"},
	{bytes: []byte{0x66, 0x74, 0x79, 0x70, 0x4d, 0x34, 0x41}, t: "audio/mp4"},
	{bytes: []byte{0x0, 0x0, 0xff, 0xfb}, t: "audio/mpeg"},
	{bytes: []byte{0x49, 0x44, 0x33}, t: "audio/mp3"},
	{bytes: []byte{0x41, 0x49, 0x46, 0x46}, t: "audio/aiff"},
	{bytes: []byte{0x57, 0x41, 0x56, 0x45}, t: "audio/wav"},
	{bytes: []byte{0x2e, 0x73, 0x6e, 0x64}, t: "audio/basic"},
	// anything else we can just store as 'octet-stream'
}

func detectMimeType(data []byte) (string, error) {
	for _, sig := range signatures {
		if bytes.HasPrefix(data, sig.bytes) == true {
			return sig.t, nil
		}
	}

	// Return a default
	return "application/octet-stream", errors.New("Coulnd't detect file type.")
}