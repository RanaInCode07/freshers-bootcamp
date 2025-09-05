package main

import "testing"

// manually creating mock objects
func TestMockDatabase(t *testing.T){
	//mock object instantiated
	mockDB := &MockDatabase{
		Data : make(map[string]string),
	}
	// testing the Set Method
    err := mockDB.Set("name", "Go")
	if err != nil{
		t.Fatalf("Set failed with error %v", err)
	}
    // testing the Get Method for succesful retrieval
	data, error := mockDB.Get("name")
	if err != nil{
		t.Fatalf("Get failed with error %v", error)
	}
	expectedValue := "Go"
	if data != expectedValue{
		t.Errorf("Get returned unexpected value. Got: %s, Expected: %s", data, expectedValue)
	}
    
	// testing the Get method for non-existing key
	_, err1 := mockDB.Get("age")
	if err1 == nil{
		t.Errorf("Get for non existent key should have returned an error but it didn't")
	}
	expectedError := "key not found"
	if err1.Error() != expectedError{
        t.Errorf("Get returned unexpected error message. Got: %s, Expected: %s", err1, expectedError)
	}
}

