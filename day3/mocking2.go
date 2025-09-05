package main



// mocking via implementing interface
type Database1 interface{
	Get1(Key string) (string, error)
	Set1(Key, Value string) error 
}

type MockDatabase1 struct{
	GetFunc func(Key string) (string, error)
	SetFunc func(Key, Value string) error
}

func (m *MockDatabase1) Get1(Key string) (string, error){
	return m.GetFunc(Key)
}

func (m *MockDatabase1) Set1(Key, Value string) error{
    return m.SetFunc(Key, Value)
}