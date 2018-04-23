package color

type Color interface {
	WhatColor() string
}

func WhatColor(c Color) string {
	return c.WhatColor()
}
