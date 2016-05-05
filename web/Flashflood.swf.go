package web

import (
	"bytes"
	"compress/gzip"
	"io"
)

// Flashflood_swf returns raw, uncompressed file data.
func Flashflood_swf() []byte {
	gz, err := gzip.NewReader(bytes.NewBuffer([]byte{
0x1f,0x8b,0x08,0x00,0x00,0x09,0x6e,0x88,0x00,0xff,0x00,0x70,
0x0f,0x8f,0xf0,0x43,0x57,0x53,0x0a,0x0e,0x18,0x00,0x00,0x78,
0xda,0x7d,0x38,0x6b,0x73,0x1b,0xd7,0x75,0xf7,0x5e,0xec,0xee,
0x5d,0x2c,0x40,0x12,0x00,0x41,0xca,0x7a,0x50,0x86,0x2c,0x48,
0x94,0xc4,0x25,0x09,0x3d,0x2d,0xd1,0xb2,0x4d,0x8a,0x24,0x08,
0x32,0x22,0x21,0x89,0xd4,0xc3,0x71,0x48,0x60,0x01,0xec,0x72,
0xd7,0x06,0x01,0x08,0xbb,0xa0,0x25,0x3b,0xb6,0x10,0x26,0x91,
0x1f,0x4a,0x9b,0xa4,0xb5,0xdd,0xf4,0x95,0xca,0x51,0xfa,0xb4,
0xd3,0xa4,0x49,0x9b,0xb6,0xd3,0x0f,0xed,0x78,0x26,0x9f,0x3a,
0xd3,0x10,0xe0,0x88,0xfa,0xd0,0x0f,0x6d,0xa7,0x33,0x9d,0x7c,
0xea,0x5f,0x60,0xcf,0xd9,0x05,0x48,0xd1,0x6d,0x3a,0xc2,0xd9,
0x7b,0xee,0xb9,0xe7,0x9c,0x7b,0x5e,0xf7,0xdc,0x4b,0x55,0x48,
0xd7,0x43,0x42,0xfe,0xa1,0x4e,0x0e,0x52,0x32,0x11,0xf6,0x13,
0x42,0xc6,0xd9,0xd6,0xd6,0xd6,0xcb,0x91,0xc7,0x7b,0x08,0x09,
0x91,0x21,0x22,0xfc,0x8a,0xd7,0xe5,0x87,0x12,0x21,0x7f,0x4f,
0x49,0xa7,0x51,0xd4,0x6c,0x73,0x48,0xbf,0xe3,0xe8,0xd5,0x92,
0x56,0x0c,0x4f,0xb6,0x90,0xe9,0x12,0x8c,0x86,0x96,0xd7,0x45,
0xdb,0xd1,0xaa,0x8e,0x92,0x44,0x36,0xa3,0x58,0x2e,0x17,0xe4,
0x64,0xad,0x94,0x77,0xac,0x72,0x29,0xa0,0x15,0x0a,0xe3,0x5a,
0xb1,0x98,0xd3,0xf2,0xaf,0x0b,0xa6,0x76,0x47,0xf7,0x5d,0x2e,
0x2f,0x8b,0xd7,0x74,0xad,0x70,0x57,0x36,0xac,0xa2,0x3e,0xa7,
0xad,0xe8,0x1d,0x3b,0x72,0x43,0xe6,0x1d,0xa5,0x68,0x95,0xf4,
0xb9,0xda,0x4a,0x4e,0xaf,0xfa,0xf3,0xb0,0x60,0x23,0x8b,0xb2,
0xa2,0x3b,0x66,0xb9,0x80,0xa8,0xb0,0xa2,0x59,0x25,0xd1,0xa9,
0xc2,0xb6,0xc2,0x6a,0xd9,0x2a,0x04,0xe6,0x2b,0xba,0x5e,0x58,
0xd0,0x6d,0xb0,0x45,0x70,0x60,0x90,0xe6,0x9d,0xaa,0x55,0x5a,
0x96,0xd2,0xb9,0xd7,0xf4,0xbc,0xe3,0xf7,0x6c,0x2f,0xe9,0x8e,
0x72,0xfd,0xda,0xe5,0x6b,0xfa,0xed,0x1a,0xb0,0xf8,0x87,0xdd,
0xdd,0x5e,0xb6,0x5f,0x94,0x6c,0xa7,0xaa,0x6b,0x2b,0x42,0xb1,
0xac,0x15,0xfc,0xc0,0x30,0xef,0x4d,0x27,0x34,0x47,0xf7,0xdb,
0x4e,0xb9,0x92,0x71,0x2c,0xd8,0xdd,0x75,0xcf,0x45,0xfd,0xb9,
0xbb,0xb0,0x07,0xf0,0x14,0x3a,0x5d,0x6c,0x6c,0x55,0xb3,0x8a,
0x5a,0xae,0xa8,0xb3,0x9c,0xe6,0x47,0xf2,0x25,0xa4,0x06,0x5b,
0x01,0x5b,0xd5,0x4b,0x8e,0x2d,0x4e,0xe2,0xd0,0x55,0xa9,0x96,
0x97,0xab,0xba,0x6d,0xa7,0xb4,0x52,0xa1,0xa8,0x57,0xc3,0xf3,
0xa8,0x13,0xec,0x8c,0x15,0xca,0x6f,0x94,0x70,0xfb,0x40,0xb9,
0xa2,0x97,0x5a,0xab,0xbe,0x52,0xf9,0x0d,0xbe,0xac,0x3b,0x0b,
0xb0,0xa5,0x30,0xab,0x39,0xa6,0x58,0x2d,0xd7,0x4a,0x05,0x37,
0x70,0x31,0xdf,0xbc,0x53,0x40,0xbb,0x41,0x58,0x89,0xb9,0x56,
0xc4,0xac,0x52,0x8c,0x0d,0xc5,0xc4,0xd8,0xec,0xa5,0x61,0xbb,
0x2b,0x5f,0x5e,0xa9,0x14,0x75,0x47,0x6f,0xe9,0x92,0x60,0xd7,
0x5a,0xd1,0x11,0xf2,0x90,0x06,0xa1,0x66,0x95,0x9c,0x80,0x67,
0x5d,0xcd,0xb1,0x8a,0xb6,0x1f,0xed,0x1d,0xab,0x56,0xb5,0xbb,
0x42,0xfa,0xca,0xe4,0x5c,0x08,0xd2,0xe5,0x9a,0x7b,0xd9,0x82,
0x78,0x96,0xf4,0x6a,0xc7,0x95,0x96,0xd9,0x2e,0x55,0xbe,0x72,
0x2d,0x3d,0x75,0x6d,0x72,0x7e,0x5e,0x1e,0x4f,0xcf,0x5e,0xb9,
0x3c,0xb9,0x30,0x29,0xba,0xba,0x84,0x4b,0xe5,0xb2,0x13,0xcc,
0x64,0x3c,0xa3,0x32,0x55,0x3d,0xcf,0x27,0xee,0x96,0xb4,0x15,
0x2b,0xaf,0x78,0x7b,0x39,0x50,0x3b,0xfe,0x05,0xf8,0x24,0x2d,
0xbd,0x58,0x60,0x8e,0x21,0xba,0x7b,0x8a,0x98,0x6b,0x08,0x50,
0xb5,0x5a,0xae,0xfa,0x81,0xd1,0x71,0x31,0x0e,0xda,0x8a,0xba,
0x56,0x0a,0xd9,0xaf,0x5b,0x95,0x4c,0xbe,0x5c,0x02,0xb5,0xb5,
0xbc,0x53,0xae,0x4a,0x15,0xad,0xaa,0xad,0xd8,0x3e,0x47,0x5b,
0x3e,0x64,0x3a,0x4e,0x65,0x64,0x78,0x58,0x2b,0x94,0x73,0xfa,
0x10,0x78,0x3c,0x3c,0x36,0x7f,0x7a,0xf8,0x54,0x22,0x71,0x6e,
0x38,0x57,0xb3,0x8a,0x10,0x58,0xa1,0x52,0xb3,0x4d,0xa9,0xa8,
0x97,0x96,0x1d,0x93,0x1e,0xa3,0xaa,0xf0,0x5a,0xd9,0x2a,0xd1,
0xe3,0x9d,0x7a,0xa9,0xb6,0x92,0x71,0xca,0x2d,0x5b,0xc1,0x68,
0x0b,0xf4,0x6b,0xa5,0xbc,0x5e,0x36,0x50,0x3f,0x78,0x19,0xa8,
0xea,0x2b,0xe5,0x55,0x7d,0xdc,0xb4,0x8a,0x85,0x8e,0x4c,0x26,
0x0f,0xb6,0x54,0x33,0x6e,0xc1,0xc9,0x98,0x12,0x44,0xfc,0xe8,
0xd0,0x78,0xb9,0x08,0xc6,0x42,0x15,0x43,0xe5,0xd9,0x61,0xf0,
0x5e,0x77,0x3c,0x36,0xb0,0x19,0x56,0x14,0xd7,0xdf,0x72,0x75,
0x45,0x73,0x7c,0x97,0xad,0x1c,0xcf,0xd7,0xaa,0xa8,0xbd,0xc3,
0x8b,0x48,0xc1,0xb2,0x2b,0x45,0xed,0xae,0x7f,0xb6,0xbc,0x6a,
0xe9,0xe3,0x45,0xab,0xd2,0x81,0xba,0xb7,0x25,0xc4,0x8c,0xad,
0x95,0x6c,0xc1,0x28,0x97,0x9c,0x70,0x41,0x37,0x34,0x48,0xe0,
0xce,0xa2,0x62,0xeb,0x45,0x28,0x6d,0x2c,0x3b,0x3c,0x79,0xcb,
0xba,0x38,0x8f,0x5f,0xc5,0xc5,0x6f,0x5a,0x05,0x28,0x97,0x37,
0xf0,0x1b,0xde,0x8e,0xf8,0x58,0xcd,0x29,0xcf,0x5b,0x6f,0xea,
0xc2,0xe5,0xc9,0xe4,0x82,0xac,0xb5,0x66,0xc1,0x95,0x72,0xcd,
0xd6,0x27,0x4b,0xa8,0xa8,0x20,0xe3,0x49,0x45,0x97,0xa5,0x63,
0xa5,0x5a,0xb1,0x78,0x9c,0x8e,0xb0,0x91,0x58,0x30,0x5f,0x83,
0x53,0xb0,0x72,0xc5,0x8d,0x3b,0x55,0x44,0xb0,0xd9,0x72,0x24,
0x48,0x49,0x5e,0x73,0x04,0x8c,0x81,0x68,0x9b,0x96,0xe1,0x48,
0xa6,0x6e,0x2d,0x9b,0x4e,0xc0,0x35,0x20,0xe5,0xe2,0x3c,0xe3,
0xc5,0xa2,0x07,0xdc,0xba,0x5a,0xd3,0x8a,0x96,0x61,0xe9,0x85,
0xf1,0xf6,0x79,0xa6,0x6f,0xc9,0x60,0x82,0x9b,0x01,0x1a,0xf3,
0xc5,0x46,0x62,0xf4,0x6d,0xe9,0xb0,0x5b,0x12,0xf4,0x55,0xba,
0x28,0x1b,0xad,0xee,0xa1,0x5c,0x6c,0x63,0x2f,0xf9,0xa1,0xfc,
0x75,0x03,0xca,0xa5,0x20,0xa0,0x79,0x90,0x97,0x5a,0x09,0x8e,
0x94,0x03,0x71,0xc8,0x64,0x44,0xd0,0xbf,0x5c,0x92,0x17,0xd2,
0x57,0x32,0xe8,0x60,0xe7,0xd8,0xc4,0xc4,0xe4,0x44,0x66,0x21,
0x9d,0x99,0x5f,0x18,0x9b,0x9a,0x0c,0x14,0xca,0xd3,0x25,0xcb,
0x99,0xd0,0x21,0xde,0x10,0x3a,0xf7,0x50,0x95,0x6b,0x8e,0x60,
0x01,0x51,0x02,0x83,0xf4,0xaa,0xd3,0xed,0xa5,0x7c,0x57,0xe9,
0xfb,0x51,0x7d,0xd9,0xb9,0x5b,0x81,0x20,0x17,0x2d,0xe8,0x36,
0xf9,0x72,0xe5,0xae,0x84,0x11,0xc8,0xeb,0x12,0x14,0x0d,0x88,
0x71,0x0b,0x8c,0xba,0x93,0x36,0x24,0x4f,0x9c,0x9b,0x9a,0x3d,
0x07,0x41,0xf1,0x41,0xaa,0x7d,0x5a,0xb5,0x2a,0x94,0x70,0x02,
0x27,0x4e,0xb6,0xa0,0x44,0x34,0xa8,0xe0,0x3d,0xb0,0x3d,0x9c,
0xa8,0x0a,0xc8,0xde,0x9d,0xb6,0x27,0xa1,0x14,0x81,0x0e,0xd1,
0x97,0xbc,0x2a,0xf2,0xad,0x68,0x15,0x25,0x6f,0x6a,0xd5,0xf1,
0x72,0x41,0x1f,0x73,0x44,0x0b,0x82,0x35,0xc7,0x5b,0x26,0xcb,
0x46,0xb5,0xbc,0x82,0x48,0xef,0x74,0x69,0x15,0x1c,0x2e,0xc4,
0x0a,0xd0,0xa5,0x62,0x86,0x5b,0x12,0x10,0xc2,0x00,0xf0,0x5d,
0x5f,0x18,0x4f,0x95,0x6b,0x55,0xbb,0xc3,0xc3,0x67,0xad,0x52,
0x0d,0xda,0x43,0x6b,0x36,0xaf,0x43,0xe2,0x0a,0x36,0x1d,0x54,
0x50,0x93,0x17,0x7d,0xac,0xe8,0x59,0x28,0x32,0x13,0xbb,0x0d,
0x76,0x3d,0x24,0xb8,0x2a,0x14,0x5c,0xf1,0xe4,0x11,0x6d,0x09,
0x07,0x00,0x4d,0x42,0xfc,0x5f,0x81,0x33,0x41,0x13,0x92,0xd7,
0xa7,0x7d,0x60,0x65,0x78,0x6e,0x72,0x6a,0x6c,0x61,0xfa,0xc6,
0x64,0x66,0x7a,0x2e,0x39,0x3d,0x37,0xbd,0xf0,0x4a,0xf8,0x4a,
0x7a,0x7e,0x7a,0x17,0x45,0xb6,0xec,0x24,0x06,0x1d,0xba,0x28,
0xf4,0x8b,0x8c,0xa1,0x17,0x0a,0xc1,0x5e,0xd6,0x4b,0x7b,0xe5,
0xde,0x3d,0xbd,0xf1,0xde,0xe7,0x7b,0x5f,0xea,0xbd,0xd4,0x4b,
0xe5,0xd9,0xde,0x2c,0xa3,0xfe,0x6f,0x51,0x4e,0x7d,0x9c,0x89,
0x9c,0x09,0x9c,0x49,0x9c,0x71,0xee,0xf3,0x73,0x16,0xe6,0x2c,
0xc2,0x59,0x88,0xb3,0x6e,0xce,0xa2,0x9c,0xf5,0x70,0xd6,0xcb,
0x85,0x67,0x38,0xdb,0xc7,0xd9,0x7e,0x2e,0x1c,0xe0,0xac,0x8f,
0xb3,0x83,0x9c,0x3d,0xcb,0x59,0x8c,0xb3,0x43,0x9c,0x3d,0xc7,
0xd9,0x61,0x2e,0x1e,0xe1,0xec,0x28,0x67,0xc7,0x39,0x3b,0xc6,
0xd9,0x09,0xce,0x06,0x38,0x53,0x39,0x1b,0xe2,0x6c,0x98,0xb3,
0xb3,0x9c,0x9d,0xe6,0xec,0x1c,0x97,0xce,0x73,0x76,0x81,0xb3,
0x11,0x2e,0xbe,0xc0,0xd9,0x45,0xce,0x5e,0xe4,0xfc,0x65,0xce,
0x46,0x39,0x1b,0xe3,0xf2,0x38,0x67,0x13,0x9c,0x4d,0x72,0x96,
0xe4,0x6c,0x8a,0xb3,0x14,0x67,0xd3,0x9c,0xcd,0x70,0xf6,0x25,
0xce,0x2e,0xef,0xa7,0x5c,0x99,0xe3,0x2c,0xcd,0x95,0x6b,0x9c,
0x2d,0x70,0x76,0x9d,0xb3,0x1b,0x9c,0xdd,0xe4,0xec,0x16,0x67,
0xaf,0x70,0xf6,0x65,0xce,0x5e,0xe5,0xec,0x2b,0x5c,0x5e,0xe4,
0x7c,0x89,0xb3,0x0c,0x0f,0x68,0x9c,0xe5,0x38,0x2b,0x70,0xa6,
0x73,0x66,0x70,0xb6,0xcc,0x03,0x26,0x67,0x16,0x67,0xaf,0x71,
0xf9,0x75,0xce,0x8a,0x9c,0xad,0x70,0x56,0xe2,0xac,0xcc,0x59,
0x80,0xb3,0x0e,0xce,0x6c,0xae,0xd4,0xb8,0xb2,0xca,0xd9,0x1b,
0x5c,0xb9,0xc3,0xd9,0x5d,0xce,0xde,0xe4,0xec,0x2d,0x2e,0x7d,
0x95,0xb3,0x77,0x38,0x7b,0x9f,0x72,0xf6,0x01,0x00,0x44,0x8b,
0xfd,0x06,0xe5,0xd2,0x6f,0xc2,0xf8,0x6d,0x80,0xef,0x02,0xfc,
0x16,0xc0,0x6f,0x03,0x7c,0x08,0xf0,0x11,0xc0,0xc7,0x00,0xbf,
0x03,0xf0,0x3d,0x80,0xdf,0x03,0xf8,0x7d,0x80,0x3f,0x04,0xf8,
0x3e,0xc0,0x1f,0x01,0x3c,0x04,0xf8,0x04,0xdc,0xf9,0x01,0x8c,
0x8f,0x00,0x70,0xfc,0x21,0xc0,0x1f,0x03,0xfc,0x29,0xc0,0x9f,
0x01,0xfc,0x39,0xc0,0xa7,0x00,0x9f,0x01,0xfc,0x08,0xe0,0x2f,
0x01,0x7e,0x0c,0xf0,0x13,0x80,0xbf,0x02,0xf8,0x19,0xc0,0x5f,
0x03,0xfc,0x0d,0xc0,0xcf,0x01,0xfe,0x16,0xe0,0xef,0xe8,0x45,
0x22,0xc3,0x0b,0xa3,0x05,0x54,0x0e,0xee,0x4c,0xb6,0x09,0x54,
0xde,0xfb,0xf4,0x67,0xd7,0x3a,0xf2,0x3f,0x4d,0xf8,0x02,0x30,
0x4f,0x09,0x41,0xce,0x0e,0x42,0xd8,0x19,0xc0,0x5a,0xb4,0x38,
0xa0,0x43,0x30,0x93,0x09,0x2e,0x04,0x49,0x90,0x30,0x1a,0x6c,
0x6f,0x4f,0xdc,0x05,0x14,0x43,0x09,0xe0,0x60,0xe4,0x8c,0x47,
0xa6,0x09,0xa1,0xfd,0xe9,0xf8,0xe2,0xde,0xb8,0x1c,0x71,0x3f,
0xa8,0x07,0x2d,0x3f,0x73,0x9f,0xb6,0xbf,0xff,0x8f,0x95,0x9e,
0x11,0x32,0x7d,0x0a,0x21,0x32,0xeb,0xc0,0x41,0x81,0xaf,0xe8,
0x0f,0x50,0xc2,0xbb,0x08,0x09,0x93,0x6e,0x42,0x22,0x24,0x8a,
0x9f,0x7d,0x94,0xc8,0x07,0x28,0xf1,0x1f,0xa6,0x44,0xe9,0x21,
0x24,0x4e,0xe0,0x71,0x76,0x84,0xc4,0x80,0x3d,0x48,0x06,0xe1,
0x1b,0x22,0x03,0x29,0x42,0xba,0x7d,0x3e,0x4a,0x7a,0x96,0x29,
0xe9,0xd5,0x29,0xd9,0xf3,0x80,0x0e,0x10,0x72,0x82,0x2e,0x1f,
0x22,0x6a,0x12,0x58,0xce,0x10,0x09,0xbe,0xcf,0x13,0xca,0xfc,
0x61,0xca,0x7c,0x61,0xe6,0x93,0x48,0x07,0x3d,0x14,0xa6,0x9d,
0x61,0x12,0x0d,0x0e,0x13,0x3a,0x44,0x4e,0x12,0x96,0x20,0xa7,
0x89,0xef,0x14,0x39,0x4b,0x84,0x33,0xe4,0x62,0x98,0x3e,0xf3,
0x62,0x98,0xed,0x1d,0x0d,0xfb,0xf6,0x4d,0x84,0x85,0xfd,0x63,
0x61,0xf1,0x40,0x26,0x2c,0xf5,0xa9,0x61,0x7e,0x30,0x17,0x96,
0x9f,0x1d,0x20,0x67,0xe9,0x14,0xa1,0x29,0x72,0x9e,0x72,0x42,
0x05,0x22,0x0b,0x94,0x41,0xa4,0x02,0x54,0x11,0x08,0xed,0xa2,
0x31,0x81,0xb0,0x08,0x1d,0x14,0x88,0x2f,0x46,0x07,0x04,0x22,
0x9c,0xa6,0x0f,0x80,0x49,0x3c,0x47,0x93,0x02,0x91,0x2e,0x50,
0x49,0x20,0xfc,0x22,0x53,0x28,0xa5,0xfd,0x59,0xb8,0x94,0xb3,
0xcc,0xf0,0xd5,0x85,0xb4,0xc8,0xb2,0x92,0xaa,0xa8,0x01,0x35,
0xa8,0x76,0xc4,0xbb,0xd4,0x4e,0x55,0x54,0xbb,0xd4,0xd0,0x75,
0x21,0x0d,0x07,0x8e,0x10,0x1f,0xf3,0x51,0xda,0xb9,0xa8,0xcc,
0x28,0xa4,0xae,0x6c,0x36,0x1b,0xe9,0x00,0x18,0x40,0x28,0x01,
0x25,0x88,0xe0,0xc8,0x7c,0xeb,0x09,0x40,0x05,0x46,0xa9,0xaf,
0x73,0x3d,0x51,0xc8,0x76,0x24,0xb2,0x1d,0xb7,0x48,0x9f,0x89,
0xf2,0x52,0x9b,0x93,0x43,0xa4,0xd8,0xde,0xf5,0xc4,0x62,0xa7,
0xba,0xb7,0xf1,0xb0,0x1e,0x9c,0xe9,0xa4,0xf5,0xce,0xcd,0xa5,
0x2e,0xa3,0xab,0x99,0x0e,0x21,0x83,0xec,0x63,0x94,0x9d,0x59,
0x4f,0x84,0x0e,0x11,0xe2,0x5f,0xea,0x59,0xea,0x31,0x7a,0x70,
0xd5,0xe8,0xfd,0x85,0xe9,0x22,0x4b,0x7b,0x8c,0x3d,0x1e,0x21,
0xfd,0x0c,0xf3,0x90,0x38,0x89,0x6e,0x6c,0x6d,0x81,0xac,0x5f,
0x01,0xd9,0xfe,0xf5,0x04,0x78,0xd2,0xdf,0xf2,0xe4,0x34,0x78,
0xd2,0x0d,0x9e,0x1c,0x73,0x3d,0x59,0x8a,0x66,0x23,0xc9,0xfd,
0xa4,0x1e,0x31,0xa3,0xc0,0xae,0x28,0x22,0x85,0xa3,0xbb,0x9e,
0x58,0xea,0x6e,0x93,0xbb,0x97,0xba,0x8d,0xee,0x64,0x1f,0xa9,
0x2d,0x45,0x8d,0x28,0x8e,0x9f,0x0c,0xd2,0x47,0xb5,0x4d,0xb4,
0xa2,0xf9,0xa8,0xf6,0x24,0x7b,0x30,0x5e,0xd8,0xf8,0xc1,0x20,
0x7b,0x04,0xbf,0xe4,0xb3,0xd4,0x8e,0x17,0x1e,0xd5,0xf2,0x02,
0x6c,0x37,0x98,0x8d,0x21,0x4f,0xf2,0x10,0xad,0x07,0xc1,0x2b,
0x35,0x01,0x9f,0x26,0x22,0x27,0xe1,0x93,0x13,0x10,0x3b,0x85,
0x1f,0xcf,0xa8,0x89,0x96,0x51,0xa7,0x5d,0xa3,0x20,0x09,0x67,
0x72,0x42,0xfa,0x39,0x8c,0x93,0x08,0x61,0x66,0x73,0x60,0x51,
0x4f,0x9c,0xa0,0xbf,0x8b,0xe1,0x99,0x30,0x31,0xc1,0xe9,0xc5,
0x23,0x33,0x47,0x88,0xe9,0x7a,0x9e,0xdd,0x6b,0x1c,0x5d,0x3a,
0x60,0x1c,0x80,0xa4,0xf5,0xbb,0x11,0xc8,0x1e,0x33,0x8e,0x2f,
0xed,0x33,0xf6,0xed,0x10,0xf6,0x1a,0x27,0x96,0x0e,0x1b,0x87,
0x5d,0x02,0x68,0x0d,0xec,0xce,0x09,0xed,0x33,0x15,0xa0,0x76,
0x42,0xa4,0x69,0x20,0x3b,0xd0,0x50,0x69,0x52,0x65,0xf5,0x60,
0x0a,0x0f,0x60,0x2b,0x4d,0xc1,0x9d,0x84,0x76,0xed,0x16,0x66,
0x7d,0x66,0x6c,0x0a,0xcf,0x48,0x8b,0x33,0xb4,0xc3,0x19,0xd9,
0xcd,0xe9,0xeb,0x33,0x07,0x81,0xfa,0x8c,0x20,0x51,0x5a,0x8c,
0xad,0xe5,0xc5,0x86,0x71,0xae,0x9e,0x88,0xd5,0x13,0x51,0xf8,
0x9b,0xa6,0x61,0x3c,0x0f,0x3b,0xde,0x20,0xf5,0xc4,0x66,0x9c,
0xd8,0x4f,0xdc,0xa5,0x7a,0x22,0x2f,0x84,0xf6,0x43,0x1a,0x73,
0xc2,0x86,0x71,0x7e,0x0d,0x24,0x3e,0xf7,0x35,0xb3,0x03,0x39,
0xb1,0x65,0x60,0xfa,0x02,0xdd,0xc8,0x09,0xc6,0x88,0xdd,0xf3,
0xaf,0x5b,0x5b,0xae,0x02,0xf5,0x0a,0xc6,0x59,0xbd,0x9a,0x7c,
0xa1,0x15,0xf7,0xf9,0x87,0xae,0x1f,0x7b,0x99,0x48,0xc5,0xf8,
0x7a,0xe2,0xe6,0x89,0x27,0x09,0xd8,0x59,0x68,0x66,0x07,0xa3,
0x8c,0x90,0xa3,0xa9,0x46,0xf3,0x67,0xa9,0x10,0x27,0x64,0x3d,
0xb1,0x91,0x00,0x7a,0x7f,0x8a,0x2a,0x7b,0x60,0x47,0xb2,0x0f,
0x4c,0xa7,0xc7,0xb2,0x03,0xc6,0x70,0xac,0x3e,0x14,0x85,0x8e,
0x30,0x85,0xb8,0xf1,0x12,0x7e,0xd3,0x2f,0xd3,0xec,0x00,0x90,
0xcd,0x61,0x1c,0x12,0xe6,0x49,0x70,0x6a,0x3f,0x9e,0x87,0x68,
0x76,0x20,0x39,0x46,0xea,0x43,0xf5,0x21,0x38,0x11,0xe6,0xa5,
0xe6,0x0d,0x62,0x8e,0xc3,0xda,0x01,0x28,0x71,0xfa,0x5d,0x1a,
0xbb,0x3f,0xb9,0x99,0x4d,0x1a,0x53,0xf5,0xd4,0xe3,0xb6,0xde,
0x1b,0x84,0x64,0x07,0x16,0x87,0x66,0x86,0x08,0xea,0x32,0x86,
0x93,0xd3,0xa4,0x3e,0x59,0x9f,0xdc,0x6c,0xaa,0x79,0x73,0x06,
0x09,0x4d,0xf3,0x4b,0x38,0xf4,0x9b,0x97,0x71,0x68,0x18,0xb3,
0xb1,0xfa,0x5c,0x14,0xcc,0x1d,0xf4,0xd5,0x42,0x6e,0xd0,0x66,
0x8d,0x74,0xcd,0xbc,0x82,0x8b,0xd9,0xab,0xc6,0x35,0x73,0xde,
0xe3,0x5e,0x68,0x71,0x42,0x2b,0x6c,0xb8,0x16,0x5f,0xa7,0x21,
0xc5,0x65,0x6f,0xcd,0x70,0x80,0xb0,0xf4,0x89,0x70,0xea,0x3f,
0x63,0x71,0x92,0x97,0x62,0xf7,0x13,0x79,0x0e,0xa1,0x91,0xb7,
0x7d,0x78,0xd2,0x8c,0xad,0x61,0x66,0xd4,0x4a,0x28,0x4c,0x48,
0xd3,0xb8,0x01,0xe1,0xbc,0x8d,0xe1,0x35,0x6e,0xda,0x30,0xd4,
0x83,0x50,0xe8,0x03,0xc6,0xc9,0x56,0xfe,0xb2,0x03,0xe0,0xed,
0xc9,0x9c,0xa0,0x56,0x61,0x6d,0xa7,0x86,0x3c,0x46,0xb1,0x1f,
0x74,0x75,0x83,0x35,0xc7,0x9b,0xc6,0x2d,0x37,0xe5,0x9f,0x7e,
0xbc,0x1a,0x19,0xc7,0xf6,0x69,0xe7,0x25,0x97,0x06,0xb9,0xe6,
0xa1,0xa3,0x98,0x6b,0x9e,0x93,0xdc,0x64,0xcb,0x9f,0x4b,0x90,
0xe9,0xab,0x90,0x6f,0x79,0x47,0x59,0x4b,0x5f,0x4e,0xca,0x71,
0xcc,0xfb,0x2f,0xb7,0xb6,0xb2,0x03,0x68,0x05,0x70,0x3a,0xc9,
0x57,0x68,0x3d,0x91,0xfc,0x32,0xc5,0x94,0x6c,0x20,0x11,0x48,
0x58,0x07,0xe6,0xab,0x18,0x76,0x63,0xb6,0x3e,0x97,0xf7,0xe7,
0xfc,0x18,0x17,0xcc,0x67,0xa8,0x17,0x36,0x43,0xae,0xf4,0x57,
0xc8,0x6e,0xee,0x7e,0x9c,0x81,0xf6,0x38,0xed,0x84,0xa0,0x1d,
0xdf,0x30,0x16,0x73,0x7e,0x63,0xe9,0xc7,0xab,0xe1,0x7f,0x71,
0x9b,0xca,0x41,0xb1,0x8b,0x8a,0xff,0xe1,0x6b,0x15,0xd3,0x7d,
0xb0,0x06,0x63,0xe7,0x05,0x10,0x22,0x29,0x43,0x10,0xfd,0x48,
0x56,0xfa,0xf3,0x01,0x98,0x07,0x01,0x3a,0x70,0xde,0xb9,0x98,
0x6d,0x24,0xb3,0x14,0xe3,0x96,0x13,0x5c,0x27,0x44,0xb5,0x37,
0xfa,0xef,0x5e,0x10,0x64,0xac,0xfc,0x3c,0x6f,0x80,0x6c,0x28,
0xe0,0x45,0xc1,0x9f,0x93,0x0f,0x42,0x85,0x9f,0xf2,0xcb,0xe1,
0xff,0xde,0xda,0xca,0x71,0x58,0x97,0xd4,0xb7,0x41,0x50,0x39,
0xba,0x8a,0x9a,0x41,0x08,0x22,0x35,0x62,0xdb,0xf9,0x60,0x68,
0x05,0x45,0xe4,0xcf,0x65,0x3b,0xdf,0x01,0xa1,0xe9,0x30,0xce,
0xbb,0x61,0xea,0xcc,0x75,0xaa,0xef,0x44,0x21,0x79,0x8d,0xa4,
0x46,0xd6,0x9e,0xaa,0x73,0x7f,0x2e,0x10,0x81,0xeb,0xaf,0x1f,
0xf4,0xe0,0x6e,0x39,0x45,0xbd,0xea,0x05,0x56,0x01,0xf4,0x5e,
0xae,0x13,0x4f,0x4e,0x9d,0x7a,0x89,0xcc,0x75,0x42,0x36,0x9a,
0xbb,0xe3,0xaf,0xe4,0xe4,0x5c,0xb0,0xe7,0x01,0x58,0x15,0xf8,
0x38,0xe2,0x29,0xb8,0xb7,0xa3,0xe0,0x6b,0x74,0x1b,0x4f,0x85,
0x7e,0x02,0xab,0xa2,0x3a,0x19,0x7d,0x13,0x4b,0x31,0x11,0x85,
0x5b,0x49,0x5d,0xa3,0x29,0xf5,0xeb,0xb4,0xed,0x48,0x03,0xfd,
0xda,0xed,0xcd,0x85,0xa7,0xbc,0xf9,0xf5,0x86,0x42,0x61,0xb8,
0xae,0xee,0x18,0xb7,0x6d,0xd9,0x3f,0x82,0x65,0x8a,0xfa,0x0d,
0xb4,0x23,0x15,0x82,0xc7,0x46,0xe3,0x23,0x2f,0xde,0x2d,0x03,
0xbf,0x49,0xa3,0xa0,0x54,0xbd,0x4f,0x53,0x58,0xd9,0x40,0x78,
0xd7,0x23,0xbc,0x07,0x04,0x38,0xfb,0x8b,0xc1,0xc6,0x4c,0x90,
0xa6,0xe8,0x67,0xf4,0xa7,0xf4,0xe7,0xf8,0x18,0x78,0x96,0x42,
0x37,0x60,0x0d,0x38,0x30,0xd1,0x76,0x97,0xeb,0x11,0x04,0x2a,
0x7e,0x48,0xb1,0x08,0x1e,0x43,0x47,0x79,0xe2,0x1d,0xec,0xcd,
0xfe,0xfe,0xe6,0x7a,0x54,0x71,0xeb,0x1c,0x0f,0x20,0x56,0x79,
0x6b,0x66,0xe4,0x55,0xfa,0xe9,0x2a,0x3a,0x03,0x13,0xf5,0x01,
0x35,0xf3,0x78,0x80,0xd7,0x13,0x8d,0xc4,0xda,0x13,0x8f,0x39,
0x0a,0x4f,0x98,0x26,0x74,0xea,0xc2,0x92,0x6e,0xe8,0x6e,0xa7,
0x0e,0x9d,0x84,0x16,0x85,0xb2,0xe9,0x38,0xe9,0x6e,0xeb,0x59,
0x8a,0x13,0xd0,0x03,0xbc,0x8b,0xc6,0x92,0x0f,0x6f,0xea,0x38,
0x4d,0x1a,0xec,0x78,0x48,0x24,0x64,0x69,0x39,0xbd,0x4c,0xa6,
0x68,0xe8,0xe2,0x28,0x18,0xdd,0x8b,0x1d,0xac,0x13,0xda,0xcb,
0xa9,0x99,0x53,0xc4,0x3c,0xad,0x7e,0x87,0xfa,0xc0,0xee,0x3d,
0x78,0x9b,0xc2,0x75,0x8b,0xf6,0x3e,0xb5,0x97,0xc9,0x96,0x7c,
0x69,0x1f,0x01,0x86,0x6e,0x6c,0xdf,0xd2,0x7a,0x62,0x7d,0x1a,
0x67,0x87,0x70,0x26,0xc3,0x2c,0xf9,0x1a,0x14,0x10,0x21,0xcf,
0x09,0x70,0x5b,0x07,0x60,0xde,0x88,0x93,0x66,0xba,0x88,0x1a,
0x0f,0xe3,0x05,0x7e,0x00,0x49,0xc9,0x12,0x5d,0xb3,0x37,0x9b,
0xf1,0x2d,0x6c,0xad,0xfd,0xa9,0xf5,0x66,0x9c,0xa6,0x8b,0xec,
0x28,0x88,0x1d,0x41,0x5b,0xfc,0xd9,0x4a,0xf6,0x36,0x64,0xf8,
0x47,0x40,0x38,0x8a,0xaf,0x82,0x70,0xf6,0xf6,0x52,0x25,0x5b,
0x39,0x61,0x3f,0xfe,0x27,0xad,0xd2,0x80,0x44,0xc2,0x42,0x5c,
0x80,0x2b,0x03,0xb4,0xdd,0x4c,0xe8,0x34,0x4e,0x56,0xa8,0x4e,
0xd7,0x57,0x98,0xfa,0xbb,0x74,0xf4,0x88,0xfa,0x07,0x74,0xf4,
0xe8,0x75,0x96,0x22,0xec,0x36,0x3e,0x8e,0x2a,0x84,0x56,0x49,
0xbf,0x00,0x17,0xf6,0xf9,0xf5,0xc4,0xf6,0xb5,0xb1,0x0e,0xfa,
0xe1,0xce,0xd8,0x07,0x05,0xd4,0x28,0xac,0xe3,0xa5,0x31,0x46,
0x57,0x23,0x7e,0x08,0xac,0x3b,0x81,0xe3,0xf4,0xb9,0x0f,0xae,
0x8c,0x9e,0x7f,0xdb,0xda,0x6a,0xc2,0x6e,0xc7,0x44,0x90,0x1f,
0xfe,0xdf,0xf2,0x50,0x15,0xfe,0xe6,0xb6,0x82,0x1d,0xb1,0x5f,
0x79,0x62,0xc7,0x19,0xbc,0xaf,0x0e,0x82,0xcb,0xf5,0x20,0x34,
0xfa,0xe4,0x1b,0xb4,0x56,0x7b,0xb2,0x78,0x67,0x23,0x79,0x87,
0x46,0x7c,0x84,0xc4,0xd6,0x52,0x1b,0x36,0x70,0xc5,0x20,0xd4,
0xbe,0x32,0xde,0x7e,0xa9,0x44,0x36,0x75,0x4b,0xe8,0x33,0x07,
0xb2,0x09,0xc3,0x5a,0x7b,0xdc,0x18,0x3d,0x64,0xbe,0xde,0x18,
0x7d,0xce,0x5c,0x69,0x8c,0x1e,0x36,0xcb,0x8d,0xd1,0xb8,0x69,
0x37,0xd4,0x0f,0x69,0x7f,0xda,0x61,0x0d,0xf5,0xe3,0xd6,0xf8,
0xbd,0xd6,0xf8,0x7d,0x6f,0x1c,0xed,0x37,0x6b,0x8d,0xd1,0x63,
0xe6,0x6a,0x43,0xfd,0xa4,0xb5,0xf2,0xd0,0x1d,0xb3,0x41,0xc3,
0x1a,0x3d,0x6e,0xde,0xcd,0x0e,0xf4,0x9b,0x67,0x21,0x1f,0x03,
0xed,0x12,0x1d,0xc4,0x51,0xcc,0x46,0x2e,0x11,0x30,0x66,0xc8,
0x7d,0xaf,0xe1,0xa4,0x1e,0xc1,0xe7,0xda,0x9b,0x14,0x1d,0x19,
0x96,0xe1,0x2a,0xfe,0x25,0x83,0xd6,0xb5,0x19,0xbb,0x1f,0x79,
0x82,0x2d,0x4c,0xc0,0x0f,0xdc,0xcc,0x23,0xfb,0xc7,0xa0,0x18,
0x00,0xfe,0xcf,0xdf,0x0d,0xf7,0xfb,0xd1,0xaf,0x59,0x85,0xdf,
0x3f,0xbb,0x37,0x5b,0xab,0xf4,0xfe,0x84,0xe2,0xcb,0xce,0x17,
0xfa,0x4f,0x38,0x89,0xea,0x6d,0xb7,0x4d,0x43,0xbc,0x3d,0x63,
0x9e,0x6c,0xc4,0x09,0x58,0xb3,0xd1,0x8c,0x13,0xec,0x5a,0xe9,
0xaf,0x22,0x4a,0x5d,0xf4,0x6d,0x44,0x99,0x8b,0xbe,0x43,0x37,
0x52,0xa1,0x1f,0xa2,0xf8,0x5f,0xd0,0xb6,0xfc,0x62,0xa4,0x25,
0x63,0xb7,0x04,0xe2,0xf4,0x93,0x16,0xbf,0x1d,0x27,0xf8,0x6f,
0x26,0x22,0xa5,0x42,0x25,0x94,0xba,0xd7,0x16,0x6a,0x89,0xec,
0xa8,0x81,0xb7,0x80,0x27,0xbd,0x6d,0x58,0x5e,0x5c,0x8c,0xe4,
0x84,0x96,0x6a,0x40,0xda,0xba,0x01,0xf5,0x94,0xe7,0xc4,0xf6,
0xa2,0xe8,0x2d,0x22,0xe2,0x2d,0xe1,0x8e,0x10,0xfa,0x84,0x8f,
0x53,0xf6,0x5f,0x14,0x2b,0x24,0xf2,0xb8,0x91,0xac,0x53,0x02,
0x97,0xc9,0x2f,0xec,0xcd,0x46,0xf2,0x6b,0x80,0xc2,0xeb,0x26,
0xb9,0x86,0x63,0x5e,0x68,0x24,0xbf,0xee,0x22,0x62,0x23,0xf9,
0x0d,0x17,0x91,0x1a,0xc9,0x6f,0x02,0x02,0xf6,0xe1,0x4d,0x1b,
0x57,0x82,0xd0,0xf5,0xd4,0x9f,0x52,0x7c,0x3a,0x62,0xbf,0x50,
0x5d,0xec,0xa1,0xeb,0x00,0x7c,0x37,0xb6,0x19,0x36,0xb6,0x19,
0x36,0x5a,0x0c,0xf7,0xdc,0x87,0x26,0x30,0x04,0x5d,0x06,0xf7,
0xcd,0x89,0x37,0x80,0xea,0xa1,0x0f,0x5b,0xb7,0x39,0x98,0xbe,
0xcd,0x22,0xee,0xb0,0x88,0x4f,0xb3,0x48,0x3b,0x2c,0xd2,0x0e,
0x8b,0xe4,0xb1,0x40,0x19,0x9d,0x64,0x70,0x9a,0xfd,0x8b,0x1f,
0xd0,0x46,0xf2,0x03,0x0a,0x7f,0x85,0x90,0x53,0x48,0xe0,0x8b,
0x77,0x1a,0x70,0x24,0x60,0xaa,0xe2,0xc3,0xf2,0x66,0x36,0x82,
0xa5,0x3f,0x68,0xee,0x6f,0x8c,0x0e,0x99,0x6f,0x35,0x46,0x87,
0xcd,0x7b,0x0d,0x38,0x0e,0xea,0x3b,0xa3,0x09,0xed,0x7c,0xf6,
0x60,0xf6,0x3e,0x5d,0x33,0xde,0xa5,0x35,0xf3,0x5d,0xda,0x9a,
0xbc,0x07,0x93,0xf7,0xda,0x93,0xf7,0x61,0xf2,0x3e,0x4c,0x46,
0x4f,0x9a,0x1f,0xe0,0x70,0xca,0xbc,0x93,0x65,0x69,0x3f,0x36,
0xa9,0x13,0xd8,0x9d,0xf7,0xbb,0x2d,0x0b,0xdb,0x5a,0xac,0x9e,
0x8a,0x42,0x33,0xcc,0x26,0xd7,0xcd,0xa9,0x76,0x57,0x3b,0x8d,
0x2f,0xd1,0x2e,0x3c,0x8b,0x03,0x89,0xec,0xc0,0x2d,0xb1,0xcf,
0x7c,0x80,0x47,0xe4,0x6c,0xfb,0xac,0x9c,0xd9,0x79,0xb4,0x9e,
0xdb,0xfd,0x68,0x95,0xfa,0xcc,0x24,0x50,0xcf,0xb7,0x39,0x9f,
0xdf,0xe1,0x1c,0xc1,0x77,0x1d,0xbc,0x97,0x1a,0xcd,0x74,0x06,
0x5f,0xd5,0x2f,0x7c,0x61,0x7e,0x01,0x7b,0xd9,0xde,0x6d,0x4d,
0xfc,0xc4,0xe8,0x88,0xc9,0xfb,0x4c,0x29,0x2b,0x8d,0xbe,0x60,
0x72,0x7c,0x70,0x77,0xe3,0xed,0xb2,0xfd,0x1f,0x2f,0x04,0xdb,
0xf6,0xff,0x00,0xb9,0xd1,0xf7,0x3e,0x01,0x00,0x00,0xff,0xff,
0xbd,0xf5,0xc2,0x06,0x70,0x0f,0x00,0x00,
	}))

	if err != nil {
		panic("Decompression failed: " + err.Error())
	}

	var b bytes.Buffer
	io.Copy(&b, gz)
	gz.Close()

	return b.Bytes()
}