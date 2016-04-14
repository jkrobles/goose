package goose

import (
    "testing"
)

func TestConvert(t *testing.T) {
    t.Log("Starting Test")
    str := 4
    if str == 0 {
        t.Log("Error should not be nil", str)
        t.Fail()
    }

    if str == 4 {
        t.Log("Correct answer")
    } else {
        t.Log("Should be 4 but got ", str)
        t.Fail()
    }
}
