package storage

import "errors"

var ErrInvalidName = errors.New("storage: user anme can't be empty")

//capital for public and lower for private
type User struct {
	ID   int
	Name string
}

type DataStore struct {
	db map[int]User
}

//constructor for the data store
func NewDataStore() *DataStore {
	return &DataStore{
		db: make(map[int]User),
	}
}

//saves the data to the store
func (ds *DataStore) Save(id int, name string) error {
	if name == "" {
		return ErrInvalidName
	}
	ds.db[id] = User{
		ID:   id,
		Name: name,
	}
	return nil
}

//retrieves the data User
func (ds *DataStore) Get(id int) (User, bool) {
	user, ok := ds.db[id]
	return user, ok
}
