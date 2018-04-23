package main

import (
    "testing"
)

func TestWho(t *testing.T) {
	if "sample!" != Who() {
        t.Errorf("Return value of function `Who()` is not `sample!`.")
    }
}
