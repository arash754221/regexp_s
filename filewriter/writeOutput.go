package filewriter

import (
	"s/dto"
)

func Making(w []dto.IOut) []byte {
	// |tokens|status|type|
	// |-|-|-|
	// |a|✅|variable|
	// |c|🚫|e|
	// |f|g|e|

	File := []byte{124, 116, 111, 107, 101, 110, 115, 124, 115, 116, 97, 116, 117, 115, 124, 116, 121, 112, 101, 124, 10, 124, 45, 124, 45, 124, 45, 124, 10}

	for _, b := range w {
		f := []byte{124}
		f = append(f, []byte(b.Token)...)
		f = append(f, 124)

		if b.Status {
			f = append(f, []byte("✅")...)
		} else {
			f = append(f, []byte("🚫")...)

		}
		f = append(f, 124)
		f = append(f, []byte(b.Typee)...)
		f = append(f, 124)
		f = append(f, 10)

		File = append(File, f...)
	}
	return File
}
