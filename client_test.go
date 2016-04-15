package goose

import (
    "testing"
)

func TestSplit(t *testing.T) {
    t.Log("Starting Split Test")
    g := new(Goose)
    strs := g.Split("we are good", " ")
    test := g.Split("we don't like you"," ")
    if  test[0] == strs[0] {
        t.Log("Correct answer")
    } else {
        t.Log("Arrays don't match ", strs, test)
        t.Fail()
    }
}
