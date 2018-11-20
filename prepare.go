package qrcode

// mode defines the type for "mode" (method of representing a defined character set as a bit string)
// as defined in 4.1.13 of ISO_IEC-18004-2006.
type mode uint8

const (
	numeric mode = 1 << iota
	alphanumeric
	byte_
)

// versions defines the type of a set of Versions. Version is defined in 4.1.244.1.244.1.244.1.24 of ISO_IEC-18004-2006.
type versions uint8

const (
	v1to9 versions = 1 + iota
	v10to26
	v27to40
)

var characterCount = map[mode][versions]int{
	numeric:      map[versions]int{v1to9: 10, v10to26: 12, v27to40: 14},
	alphanumeric: map[versions]int{v1to9: 9, v10to26: 11, v27to40: 13},
	byte_:        map[versions]int{v1to9: 8, v10to26: 16, v27to40: 16},
}
