package util

import (
	"fmt"
)

type Color struct {
	R uint8
	G uint8
	B uint8
}

func GetContrastBackground(hex string) string {
	var c Color
	var err error
	if hex == "" {
		return "FFFFFF"
	}

	if c, err = parseHexColor(hex); err != nil {
		panic(fmt.Sprintf("Failed to parse color %s", hex))
	}

	// http://www.w3.org/TR/AERT#color-contrast
	var yiq = ((int32(c.R) * 299) + (int32(c.G) * 587) + (int32(c.B) * 114)) / 1000
	if yiq >= 128 {
		return "000000"
	} else {
		return "FFFFFF"
	}
}

func parseHexColor(hex string) (c Color, err error) {
	switch len(hex) {
	case 7:
		_, err = fmt.Sscanf(hex, "#%02x%02x%02x", &c.R, &c.G, &c.B)
	case 4:
		_, err = fmt.Sscanf(hex, "#%1x%1x%1x", &c.R, &c.G, &c.B)
		// Double the hex digits:
		c.R *= 17
		c.G *= 17
		c.B *= 17
	default:
		err = fmt.Errorf("invalid length, must be 7 or 4")

	}
	return
}
