package main

import (
	"fmt"
)

type RGB struct {
	R, G, B int
}

func rgbConvert(hexCode string) (c RGB, err error) {
	switch len(hexCode) {
	case 6:
		_, err = fmt.Sscanf(hexCode, "%02x%02x%02x", &c.R, &c.G, &c.B)
	case 3:
		_, err = fmt.Sscanf(hexCode, "%1x%1x%1x", &c.R, &c.G, &c.B)
		c.R *= 17
		c.G *= 17
		c.B *= 17
	}
	return
}

func toANSI(R, G, B int) int {
	if R == G && R == B && G == B {
		if R < 8.0 {
			return 16
		} else if R > 248 {
			return 231
		} else {
			return ((R-8)/247)*24 + 232
		}
	}

	r := (R * 5) / 255
	g := (G * 5) / 255
	b := (B * 5) / 255
	return 16 + 36*r + 6*g + b
}

func main() {
	var hexCode string

	fmt.Print("Hex code: ")
	fmt.Scan(&hexCode)
	color, err := rgbConvert(hexCode)
	if err != nil {
		fmt.Print("Invalid hex code")
		return
	}
	output := fmt.Sprintf("38;5;%dm", toANSI(color.R, color.G, color.B))
	fmt.Println("Color code:", output)
	fmt.Printf("Example: \x1b[%vcolor\x1b[0m", output)
}
