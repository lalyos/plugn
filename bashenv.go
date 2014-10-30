package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

func bashenv() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xa4, 0x58,
		0x6d, 0x53, 0xdc, 0x38, 0x12, 0xfe, 0x3c, 0xf3, 0x2b, 0x1a, 0xc5, 0x09,
		0xcc, 0xdd, 0x39, 0xce, 0xcc, 0xe5, 0xaa, 0xae, 0x60, 0x27, 0x05, 0x49,
		0x48, 0x42, 0x55, 0x16, 0x28, 0x20, 0xa9, 0xdd, 0x0a, 0x24, 0x65, 0x6c,
		0x79, 0xc6, 0x85, 0x2d, 0x4f, 0x59, 0x32, 0x2f, 0x4b, 0xe6, 0xbf, 0x6f,
		0xb7, 0x24, 0xdb, 0xf2, 0x60, 0x20, 0xc9, 0x7e, 0x81, 0x71, 0x4b, 0xea,
		0x7e, 0xfa, 0xe9, 0x17, 0xb5, 0x3d, 0x4c, 0x13, 0xf8, 0xf2, 0x05, 0x98,
		0xf7, 0x7a, 0xe7, 0xf8, 0xc3, 0xb7, 0xcf, 0xbb, 0x47, 0xc7, 0x7b, 0xfb,
		0xef, 0x0e, 0x18, 0xf8, 0x99, 0x02, 0xf6, 0x92, 0xc1, 0xd9, 0xd9, 0x16,
		0xa8, 0x39, 0x17, 0xc3, 0x01, 0x8f, 0xe6, 0x05, 0xb0, 0xb5, 0x35, 0xf8,
		0xb3, 0xa8, 0x4a, 0x90, 0x37, 0x52, 0xf1, 0x1c, 0x5e, 0x87, 0x72, 0x0e,
		0xa9, 0x84, 0xa2, 0x52, 0x50, 0x24, 0x10, 0x87, 0x8a, 0x6f, 0x82, 0xa3,
		0xeb, 0x60, 0x9f, 0x39, 0x27, 0x0f, 0x33, 0x1e, 0x4a, 0x0e, 0xd5, 0x62,
		0x56, 0x86, 0x31, 0x07, 0x55, 0x98, 0xf3, 0x2f, 0xa1, 0x28, 0x61, 0x56,
		0x72, 0x3c, 0x5c, 0x3e, 0xc7, 0xfd, 0x35, 0xa4, 0x8d, 0x4a, 0x84, 0x39,
		0x1f, 0x31, 0x98, 0x4e, 0x81, 0xbd, 0x0d, 0xcb, 0xab, 0x54, 0xb8, 0x80,
		0xb4, 0x5e, 0xfb, 0x0f, 0xd8, 0x81, 0x80, 0x83, 0x63, 0xf8, 0xe3, 0x3f,
		0x50, 0xa1, 0x85, 0x0f, 0x45, 0xce, 0xcf, 0x4b, 0x7e, 0x45, 0x26, 0x52,
		0x21, 0x55, 0x98, 0x65, 0x90, 0xa1, 0x7a, 0xa9, 0xb4, 0xc5, 0x4d, 0x56,
		0x1f, 0x5b, 0x07, 0x00, 0x0f, 0xf4, 0xd6, 0x7a, 0xdf, 0x39, 0x6e, 0x58,
		0x5f, 0xd5, 0x7e, 0x82, 0x16, 0x21, 0x8c, 0x63, 0x48, 0x15, 0xe9, 0x0c,
		0xb8, 0x8a, 0x02, 0x39, 0xe7, 0x59, 0x26, 0x21, 0x14, 0x31, 0x44, 0xf3,
		0x50, 0xcc, 0x38, 0xdc, 0x10, 0x33, 0x68, 0xbf, 0x5c, 0x97, 0xa0, 0x57,
		0x57, 0x0d, 0xc9, 0x2a, 0x2e, 0x70, 0x05, 0xfc, 0x08, 0x98, 0x16, 0x07,
		0x95, 0x2c, 0x83, 0xac, 0x88, 0xc2, 0x2c, 0x38, 0x4f, 0x45, 0x40, 0xb6,
		0xe1, 0xd5, 0x2b, 0x57, 0x3f, 0x5b, 0xef, 0xaa, 0x88, 0xe6, 0x74, 0x5e,
		0xf6, 0x9d, 0xbc, 0x83, 0x7a, 0x07, 0x04, 0x71, 0xc0, 0xcb, 0x3c, 0x15,
		0x61, 0x06, 0x92, 0x4b, 0x99, 0x16, 0x82, 0xe2, 0x45, 0x0c, 0xe2, 0x62,
		0x84, 0x92, 0xb0, 0xbc, 0x21, 0x97, 0x54, 0x78, 0xc1, 0x81, 0x27, 0x09,
		0x8f, 0xd4, 0x73, 0xd6, 0x28, 0x4a, 0x52, 0x0c, 0xdf, 0x35, 0x3a, 0x3d,
		0x19, 0xe2, 0xcf, 0x98, 0x47, 0x59, 0x58, 0x72, 0xf0, 0x77, 0xe0, 0xcd,
		0xef, 0x6f, 0x8f, 0x87, 0xc3, 0x28, 0x8f, 0xfd, 0x2c, 0x95, 0x6a, 0x63,
		0x04, 0xb7, 0xc3, 0x41, 0xbd, 0x1c, 0x73, 0x19, 0x4d, 0xd9, 0x47, 0x94,
		0x23, 0x3b, 0x97, 0x61, 0x9a, 0x85, 0xe7, 0x19, 0x87, 0xa8, 0xc8, 0x73,
		0xe4, 0x4a, 0xb2, 0x76, 0xa3, 0x90, 0x53, 0xe6, 0x8d, 0x51, 0x50, 0xeb,
		0xf1, 0x2f, 0xf8, 0x8d, 0xc4, 0xc8, 0x0b, 0xc9, 0xe0, 0x3b, 0xe2, 0x8d,
		0x81, 0xc9, 0x00, 0x9f, 0x36, 0x83, 0x80, 0x0d, 0x97, 0xad, 0x3d, 0xbd,
		0xaf, 0x6b, 0xb4, 0xd1, 0x95, 0x60, 0x32, 0x5d, 0x60, 0x34, 0x51, 0xcd,
		0xed, 0x1a, 0xc1, 0xfc, 0xb2, 0x7d, 0xb6, 0x64, 0x5b, 0x10, 0xb7, 0xbc,
		0x78, 0x17, 0x04, 0xa2, 0x10, 0x1c, 0x8d, 0x60, 0xde, 0x2d, 0x80, 0x7d,
		0x25, 0x23, 0xda, 0x66, 0x51, 0xaa, 0x8e, 0x25, 0x61, 0xed, 0xfc, 0x8a,
		0x5a, 0xff, 0x12, 0x36, 0x57, 0x74, 0xf2, 0xeb, 0x05, 0x3e, 0xf5, 0xf1,
		0xb5, 0xab, 0x57, 0x90, 0x31, 0x48, 0x2a, 0x11, 0x29, 0x0a, 0x54, 0x48,
		0x4f, 0x96, 0x37, 0x87, 0xb6, 0x44, 0x68, 0x57, 0x71, 0x19, 0xff, 0xdf,
		0x4e, 0x36, 0x7d, 0x6f, 0xbc, 0xc4, 0x65, 0x9d, 0x0c, 0x9a, 0x07, 0xcb,
		0x82, 0x20, 0xb8, 0xde, 0x86, 0xe3, 0xca, 0xa8, 0x8b, 0x37, 0x11, 0xcc,
		0x65, 0xc0, 0x67, 0xf0, 0x0c, 0x73, 0x2f, 0xe6, 0x97, 0x81, 0xa8, 0xb0,
		0x12, 0x9e, 0x3d, 0x33, 0xac, 0x0a, 0xeb, 0xd6, 0x70, 0xa0, 0xfd, 0xa6,
		0xf0, 0x6c, 0x7a, 0xb7, 0xa1, 0x0c, 0x9e, 0xe0, 0x2f, 0x3f, 0x58, 0xb2,
		0xb3, 0xa9, 0xd6, 0xd5, 0xf5, 0xb1, 0x61, 0x6e, 0x25, 0x42, 0xd6, 0x5d,
		0x6f, 0x42, 0x9d, 0xe1, 0x12, 0x11, 0xa3, 0x90, 0x36, 0x02, 0xe2, 0xea,
		0x30, 0x72, 0xca, 0x3c, 0xfa, 0x7f, 0x4a, 0xd9, 0x48, 0x6a, 0x85, 0x04,
		0x6f, 0x0c, 0xa7, 0xec, 0xd4, 0xdb, 0x3e, 0x45, 0xe2, 0x87, 0x83, 0xa5,
		0xcd, 0x1c, 0x63, 0x0f, 0x4c, 0xf8, 0x2d, 0xc4, 0xb1, 0x06, 0x35, 0x6e,
		0x30, 0xd5, 0x60, 0x5a, 0x92, 0x70, 0x71, 0x0b, 0x6b, 0x31, 0x4d, 0x14,
		0xd4, 0x62, 0xdc, 0xd8, 0x91, 0x7f, 0xff, 0x0e, 0xaa, 0xac, 0x78, 0xbd,
		0x8c, 0xfd, 0x41, 0x55, 0x72, 0xfa, 0x42, 0x77, 0xa8, 0x9a, 0xd5, 0x26,
		0x5b, 0x35, 0x8d, 0x5f, 0x3d, 0x94, 0x9f, 0x7a, 0x1d, 0x1e, 0x9b, 0x7e,
		0xe5, 0xdd, 0x3a, 0xfc, 0xe1, 0x3e, 0x76, 0xb6, 0xc4, 0xd3, 0xdb, 0xc4,
		0x43, 0x26, 0xd1, 0x4a, 0xd3, 0xf8, 0x68, 0xcd, 0x6d, 0x74, 0x36, 0x5e,
		0xfb, 0xd8, 0x3b, 0xaa, 0x68, 0x5e, 0x67, 0x04, 0x36, 0x5a, 0xda, 0x48,
		0xeb, 0x16, 0xd9, 0x84, 0x42, 0x9b, 0xd5, 0x5a, 0x08, 0xd6, 0x5d, 0x25,
		0xde, 0x46, 0x22, 0x7c, 0xe2, 0xd5, 0xec, 0x18, 0xd1, 0x79, 0xaa, 0xf2,
		0x95, 0xce, 0x71, 0xa7, 0x6c, 0x75, 0x2b, 0xa3, 0xac, 0x42, 0x9b, 0xdd,
		0xbc, 0xb2, 0x8a, 0x6c, 0x6a, 0x0d, 0x16, 0x65, 0x2a, 0x54, 0x02, 0x0c,
		0xe0, 0xa9, 0x3f, 0xf9, 0xbf, 0x84, 0xa7, 0xf2, 0x14, 0xd3, 0xcc, 0x3a,
		0xd5, 0x31, 0x7f, 0x97, 0x0e, 0x03, 0xc7, 0xe4, 0x5a, 0x03, 0x88, 0xda,
		0x8f, 0x67, 0x3c, 0xd4, 0x0d, 0xc9, 0x46, 0x14, 0xbb, 0xe3, 0xa2, 0xaf,
		0x8e, 0x8e, 0xe7, 0xc5, 0x95, 0x04, 0x5a, 0x45, 0x98, 0x08, 0x38, 0x0f,
		0x75, 0x2d, 0x11, 0xf4, 0xbe, 0x62, 0x0a, 0xcb, 0x19, 0x25, 0xc3, 0xb6,
		0x73, 0xef, 0x90, 0xc8, 0xa1, 0x0e, 0xdb, 0x2e, 0xd4, 0x9e, 0x93, 0x93,
		0x3a, 0x27, 0xe8, 0x07, 0x16, 0xfa, 0x13, 0xbc, 0x56, 0x90, 0x01, 0x3c,
		0x61, 0xb6, 0x39, 0xe9, 0x75, 0x4b, 0x6a, 0x82, 0xa7, 0xe4, 0x59, 0x40,
		0x8d, 0xc2, 0x08, 0x05, 0x8a, 0xa0, 0xfb, 0x0c, 0x81, 0x8f, 0x99, 0xec,
		0x1c, 0xd7, 0xf5, 0xdd, 0x43, 0x8e, 0x45, 0x22, 0x7c, 0xf2, 0xcb, 0x96,
		0xef, 0xb8, 0x49, 0x1e, 0x5b, 0x22, 0xba, 0xf8, 0x89, 0x24, 0xa7, 0x32,
		0x6a, 0xba, 0x34, 0x2b, 0xc3, 0x21, 0x6a, 0x20, 0x6c, 0x7d, 0xe4, 0xed,
		0x09, 0xb9, 0xc0, 0x1b, 0xc0, 0x69, 0x42, 0x78, 0x95, 0xe1, 0xe6, 0x2a,
		0xe7, 0x42, 0xc9, 0xa6, 0xc9, 0xa0, 0x24, 0x4b, 0x05, 0x9f, 0x7a, 0x1b,
		0xea, 0x66, 0xc1, 0xa9, 0x2c, 0x6d, 0xfe, 0xd7, 0xda, 0xda, 0x06, 0xc8,
		0x5c, 0x03, 0x54, 0x27, 0x73, 0x1e, 0xc6, 0xe0, 0x8f, 0x47, 0x76, 0x2e,
		0xf0, 0x39, 0x18, 0xae, 0x48, 0x61, 0x80, 0x64, 0x30, 0x4c, 0x97, 0x25,
		0x6d, 0x0c, 0xaf, 0x2e, 0xc0, 0x7f, 0x37, 0x85, 0xf5, 0x60, 0x1a, 0xdc,
		0xea, 0xac, 0x02, 0xf6, 0x1b, 0x95, 0xe7, 0x2b, 0xb6, 0x5c, 0x07, 0x2a,
		0x4d, 0x60, 0x3a, 0xb5, 0x40, 0x57, 0xb9, 0xcd, 0xab, 0x1f, 0xf6, 0x8a,
		0x96, 0xca, 0x74, 0x41, 0x4f, 0x3a, 0x1d, 0x68, 0x63, 0xdb, 0x92, 0xee,
		0x3a, 0x86, 0x29, 0xdb, 0x82, 0xc7, 0xf8, 0x69, 0xf4, 0xba, 0x43, 0x59,
		0xeb, 0x14, 0x95, 0x07, 0xac, 0xbb, 0x9d, 0xbd, 0xa7, 0x99, 0x4b, 0x4a,
		0x5b, 0x9c, 0x1f, 0x22, 0x5e, 0xf7, 0xc6, 0xba, 0x47, 0x83, 0xae, 0x19,
		0x8a, 0x18, 0xe0, 0xd3, 0xa8, 0x59, 0x02, 0x68, 0x8b, 0xc9, 0x59, 0x68,
		0xf3, 0xb8, 0x55, 0xd9, 0x69, 0x04, 0xc6, 0x33, 0xd4, 0x8b, 0x14, 0x62,
		0x95, 0x83, 0x2f, 0xe0, 0xdf, 0x13, 0xf7, 0xc6, 0x47, 0x77, 0xec, 0x30,
		0x74, 0x8f, 0x3b, 0x7a, 0x4e, 0x0a, 0xf5, 0x70, 0xb1, 0xc8, 0xaa, 0x19,
		0x76, 0x83, 0xa4, 0x2c, 0x72, 0x94, 0xbc, 0xc7, 0x52, 0xfd, 0x74, 0xf4,
		0xd1, 0x71, 0xaf, 0x2a, 0x33, 0xe3, 0x1f, 0x8d, 0x72, 0xd6, 0xb3, 0x08,
		0xaf, 0x75, 0xef, 0xf0, 0xe3, 0xa7, 0xf7, 0x7b, 0xfb, 0xdf, 0x0e, 0x77,
		0x4e, 0x3e, 0x04, 0xcd, 0x90, 0x80, 0x8b, 0x33, 0x54, 0x11, 0x65, 0x54,
		0x5b, 0xcc, 0xc3, 0xc3, 0x0c, 0xbc, 0x89, 0x3e, 0xe1, 0x83, 0xd3, 0x52,
		0x09, 0x62, 0x25, 0x1e, 0x00, 0x79, 0xc4, 0xf3, 0xe2, 0x92, 0x77, 0xd1,
		0x35, 0x2d, 0xcd, 0x48, 0xdd, 0x41, 0xc4, 0x48, 0xec, 0x00, 0x51, 0xe6,
		0xe0, 0x97, 0xc9, 0x0a, 0x44, 0x2e, 0xe8, 0x68, 0x1c, 0x78, 0x66, 0xeb,
		0x7d, 0xdb, 0x1a, 0x23, 0xed, 0x46, 0x84, 0xfa, 0xd0, 0x88, 0x04, 0x7a,
		0x36, 0xd5, 0x55, 0xd5, 0xe2, 0xc2, 0xc8, 0x2d, 0x14, 0x4d, 0x79, 0xe4,
		0xec, 0x2c, 0x2b, 0xce, 0xcd, 0x85, 0xbe, 0x08, 0xd5, 0x5c, 0xf7, 0xde,
		0x7e, 0x93, 0xff, 0xb2, 0x2d, 0xd8, 0xd5, 0x86, 0x3e, 0x6d, 0xe0, 0x6c,
		0xc8, 0x89, 0x7e, 0xf0, 0x48, 0x81, 0x6e, 0xb1, 0x9d, 0x0b, 0x0d, 0xb7,
		0x60, 0xbe, 0x60, 0x21, 0x3e, 0xe4, 0x31, 0x66, 0x10, 0x0d, 0x04, 0x3a,
		0xf5, 0xec, 0x12, 0xdd, 0x8f, 0xfa, 0x39, 0x4e, 0xa5, 0x16, 0x38, 0xaa,
		0x2f, 0x79, 0x49, 0xb3, 0x26, 0xe9, 0x8e, 0x42, 0x65, 0x2c, 0x07, 0x46,
		0xd3, 0x73, 0x55, 0xe4, 0x19, 0x25, 0x1f, 0xfe, 0xf3, 0x67, 0x1c, 0xab,
		0xda, 0x52, 0x05, 0xcc, 0x9e, 0x62, 0x8e, 0x22, 0x3b, 0x2d, 0xfc, 0x8c,
		0x16, 0xb7, 0xb2, 0xb5, 0xa6, 0xce, 0x95, 0xf4, 0x02, 0xaf, 0x24, 0xff,
		0x7f, 0xf4, 0x67, 0xfc, 0xa2, 0xbd, 0x9d, 0x9a, 0xc3, 0x5e, 0x8d, 0x81,
		0x0a, 0x48, 0xf3, 0x43, 0xbf, 0x74, 0xf3, 0xaa, 0xa7, 0x20, 0x1b, 0x9d,
		0xaa, 0x8d, 0x0e, 0xc6, 0x58, 0x95, 0xe9, 0x6c, 0xc6, 0xcb, 0xbe, 0x30,
		0x9f, 0x98, 0x25, 0xbc, 0x94, 0x8a, 0x42, 0x8f, 0x90, 0x35, 0x7f, 0x77,
		0x13, 0x91, 0x76, 0xb8, 0x23, 0xc8, 0xfd, 0xa9, 0x60, 0x82, 0xb2, 0x9a,
		0x0c, 0x75, 0xd0, 0xea, 0x54, 0xa8, 0xbb, 0x59, 0x54, 0x88, 0x24, 0x9d,
		0xd5, 0xf7, 0x41, 0x27, 0x25, 0xb4, 0xa2, 0xd1, 0x88, 0x6e, 0x97, 0xc1,
		0x80, 0xf2, 0xe0, 0xba, 0xa1, 0x23, 0xf0, 0x08, 0x0f, 0xb3, 0xa1, 0xef,
		0x08, 0xed, 0xe0, 0xf2, 0x00, 0x21, 0x06, 0x4a, 0xef, 0xa4, 0xab, 0x57,
		0xb0, 0x5d, 0x34, 0xb5, 0xd4, 0x5b, 0x87, 0xf9, 0x45, 0x9c, 0x96, 0xe0,
		0x2f, 0xfa, 0x2b, 0x91, 0xee, 0x22, 0x01, 0x7e, 0x22, 0x1f, 0xad, 0xc0,
		0xc7, 0x2a, 0x19, 0xb1, 0xda, 0x04, 0xee, 0x03, 0xfb, 0xd6, 0x2c, 0xfd,
		0x73, 0xb4, 0xd8, 0x2f, 0x1e, 0x07, 0x62, 0xe3, 0x84, 0xf9, 0xdc, 0x87,
		0xe5, 0x3d, 0xa6, 0xb9, 0x0d, 0xbc, 0xd9, 0x58, 0x95, 0xe1, 0xca, 0x75,
		0xe2, 0x60, 0xea, 0xb4, 0xdc, 0x50, 0xad, 0x58, 0x37, 0x0a, 0x74, 0x15,
		0xb1, 0x4e, 0x19, 0x39, 0xb4, 0x91, 0x02, 0x17, 0xd6, 0x63, 0x2f, 0x2f,
		0x3f, 0x01, 0xee, 0x47, 0x21, 0x35, 0xb3, 0xfd, 0x5d, 0x92, 0x64, 0x3f,
		0x49, 0xc7, 0xbf, 0x48, 0x12, 0x60, 0xa9, 0x54, 0xf4, 0xf3, 0xbf, 0xf5,
		0x0d, 0xeb, 0xd3, 0xcb, 0x9e, 0x96, 0x36, 0x70, 0x24, 0x7f, 0x10, 0x74,
		0x2f, 0x79, 0xa9, 0x48, 0x7b, 0x81, 0xee, 0xa1, 0x3c, 0x0d, 0xb3, 0xf4,
		0x2f, 0x4c, 0x2e, 0xec, 0x09, 0xf9, 0x42, 0xdd, 0xd4, 0xc0, 0xa9, 0xcf,
		0xdd, 0x9b, 0x53, 0xb8, 0xa0, 0x0a, 0x7a, 0x0f, 0xb8, 0x1f, 0x48, 0x7d,
		0xf4, 0xbe, 0x5c, 0xec, 0x5b, 0x75, 0x2f, 0x61, 0x33, 0x60, 0xb4, 0x00,
		0x63, 0x0b, 0xef, 0x10, 0xe1, 0x09, 0x17, 0xe4, 0x6a, 0x03, 0xd2, 0x0e,
		0x0f, 0xf3, 0x30, 0x15, 0xc6, 0x63, 0xe2, 0xcb, 0xe7, 0x05, 0x2c, 0xd2,
		0x05, 0x4f, 0x50, 0xff, 0x96, 0x99, 0x4c, 0x4e, 0x8e, 0x76, 0xde, 0xec,
		0xd6, 0x7d, 0x45, 0xef, 0xb9, 0xa6, 0xef, 0x0c, 0x3a, 0xd4, 0xfa, 0xe3,
		0xd1, 0xee, 0xfe, 0xe7, 0x69, 0x3d, 0xc8, 0xac, 0xad, 0xb8, 0xbf, 0xfa,
		0x09, 0xc8, 0x7c, 0x5a, 0x6a, 0x37, 0xd0, 0xf7, 0x0d, 0x51, 0x28, 0xad,
		0x57, 0x37, 0xdb, 0xcb, 0xb4, 0x2c, 0x04, 0x0d, 0xb0, 0xac, 0x7e, 0x9d,
		0x98, 0xe8, 0x39, 0xa7, 0xf3, 0xf6, 0x68, 0x87, 0x89, 0x8e, 0xac, 0x19,
		0x31, 0x3a, 0x52, 0xba, 0xcd, 0x3b, 0x02, 0xdb, 0xfa, 0x3b, 0x32, 0xc3,
		0x75, 0x47, 0x64, 0x9b, 0x8c, 0x2b, 0xa3, 0x81, 0xdd, 0xc4, 0x0d, 0xd8,
		0x61, 0x7f, 0xce, 0xba, 0x63, 0x7c, 0xd3, 0x1c, 0xfa, 0xc4, 0xe6, 0xa9,
		0x6f, 0x45, 0xae, 0x1c, 0xa0, 0x84, 0x1c, 0x0e, 0x86, 0xed, 0x1b, 0x83,
		0x69, 0xe6, 0xcb, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x46, 0x76, 0x41,
		0xd9, 0x0c, 0x14, 0x00, 0x00,
	},
		"bashenv",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"bashenv": bashenv,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"bashenv": &_bintree_t{bashenv, map[string]*_bintree_t{
	}},
}}