package methods

import (
  "testing"
)

func TestCheckIntakeForText(t *testing.T){

  input0 := checkIntakeForText("")
  if input0 != false {
    t.Log("Test input is empty, should be false.")
    t.Fail()
  }

  input1 := checkIntakeForText("test")
  if input1 != true {
    t.Error("Test input not empty, should be true.")
  }



}
