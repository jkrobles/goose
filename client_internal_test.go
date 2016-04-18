package goose_test

import (
    "testing"
    "goose"
)


func TestSplit(t *testing.T) {
    t.Log("Starting Split Test")
    g := new(goose.Goose)
    strs := g.Split("we are good", " ")
    test := g.Split("we don't like you"," ")
    if  test[0] == strs[0] {
        t.Log("Correct answer")
    } else {
        t.Log("Arrays don't match ", strs, test)
        t.Fail()
    }
}
