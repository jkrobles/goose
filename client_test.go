package goose

import (
    "testing"
)

func TestSplit(t *testing.T) {
    t.Log("Starting Split Test")
    g := new(Goose)
    strs := g.Split("we are good", " ")
    t.Log(strs[0])
    var test []string = new []string{"we"}
   // test[0] = "we"
    //test[1] = "are"
   // test[2] = "good"
    t.Log(test)
   if  test[0] == strs[0] {
        t.Log("Correct answer")
    } else {
        t.Log("Arrays don't match ", strs, test)
        t.Fail()
    }
}
