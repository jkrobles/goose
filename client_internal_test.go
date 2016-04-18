package goose_test

import (
    "testing"
    "github.com/jkrobles/goose"
)


func TestSplit(t *testing.T) {
    t.Log("Starting Split Test")
    g := new(goose.Goose)
    strs := g.Split("we are good", " ")
    var test []string = []string{"we","are","good"}
    if  test[0] == strs[0] && test[1] == strs[1] {
        t.Log("Correct answer")
    } else {
        t.Log("Arrays don't match ", strs, test)
        t.Fail()
    }
}
