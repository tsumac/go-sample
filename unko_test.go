package main

import (
    "testing"
)

func TestWho(t *testing.T) {
	if "unko" != Who() {
        t.Errorf("Return value of function `Who()` is not `unko`.")
    }
}
