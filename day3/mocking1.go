package main

import (
	"errors"
)
//technique1: manually creating mock objects

// interface defining the behavior of a real dependencies
type Database interface{
	Get(Key string) (string, error)
	Set(Key, Value string) error
}
// created new struct for mock
type MockDatabase struct{
	Data map[string]string
}

func (m *MockDatabase) Get(Key string) (string, error){
   data, error := m.Data[Key]
   if !error {
	   return "", errors.New("key not found")
   }
   return data, nil
}

func (m *MockDatabase) Set(Key, Value string) error{
	m.Data[Key] = Value
	return nil
}


