package main

import (
	"errors"
	"testing"
)


func TestMockDatabase1(t *testing.T){
 // creating a basic mock instance
	createMock := func() *MockDatabase1{
		return &MockDatabase1{
			GetFunc: func(Key string) (string, error){
			      return "", errors.New("default GetFunc: key not found")
			},
            SetFunc: func(Key, Value string) (error){
				return nil
			},
		}
	}

	// test a successful retrieval

	t.Run("Get a valid key", func(t *testing.T) {
		mockDB := createMock()

		mockDB.GetFunc = func(Key string) (string, error){
			if Key == "test_key"{
				return "test_value",nil
			}
			return "", errors.New("unexpected Key")
		}

		value, err := mockDB.Get1("test_key")

		if err != nil{
			t.Errorf("Expected no error, but got: %v", err)
		}
		if value != "test_value"{
			t.Errorf("Expected value 'test_value', burt got: %s", value)
		}
	})
// test an error condition during retrieval
	t.Run("Get with an error", func(t *testing.T) {
		mockDB := createMock()

		mockDB.GetFunc = func(Key string) (string, error){
			return "", errors.New("mocked get error")
		}

		_, err := mockDB.Get1("any_key")

		if err == nil{
			t.Errorf("expected error, but got nil")
		}
		if err.Error() != "mocked get error"{
			t.Errorf("expected error : 'mocked get error', but got : %s", err)
		}
	})

	// test a successful set operation

	t.Run("set a key-value pair", func(t *testing.T) {
		mockDB := createMock()
		
		var recievedKey, recievedValue string
		mockDB.SetFunc = func(Key, Value string) error{
			recievedKey = Key
			recievedValue = Value
			return nil
		}

		err := mockDB.Set1("new_key", "new_value")

		if err != nil{
			t.Errorf("expected no error, but got: %v", err)
		}
		if recievedKey != "new_key" || recievedValue != "new_value"{
			t.Errorf("Set func was not called with the right arguments. Got key: %s, value: %s", recievedKey,recievedValue)
		}
	})
}