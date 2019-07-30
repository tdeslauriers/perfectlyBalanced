package methods

import (
	"bufio"
	"fmt"
	"os"

)

func Balanced(){
    intake()
}

func intake() string {
  fmt.Println("Please enter a series of x's and y's.")
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Scan()
  text := scanner.Text()
  return text
}

func checkIntakeForText(s string) bool{

  if s != "" {
    return true
  }
  return false
}
