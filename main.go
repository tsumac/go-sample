package main

import (
    "github.com/tsumac/go-sample/color"
    "github.com/tsumac/go-sample/math"
    "fmt"
)

func main() {
	fmt.Println(color.WhatColor(new(color.Red)))
	fmt.Println(math.Increment(1))
}
