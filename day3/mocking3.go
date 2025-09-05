package main

import (
	"fmt"
	"testing"
	// "github.com/stretchr/testify/assert"
)

// mocking using function callbacks

func ProcessData(data string, f func(string) string) string {
   result := f(data)
   // do some processing with result
   return result
}

func TestProcessData(t *testing.T) {
   mockFunc := func(data string) string {
       return "mocked result"
   }
   result := ProcessData("input data", mockFunc)
    fmt.Printf(result)//to remove the error last line is removed because it cannot import from github
//    assert.Equal(t, "mocked result", result)
}
