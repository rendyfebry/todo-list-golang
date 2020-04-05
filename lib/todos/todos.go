package todos

import (
	"fmt"

	couchdb "github.com/leesper/couchdb-golang"
	uuid "github.com/satori/go.uuid"
)

// Task ...
type Task struct {
	ID   string
	Text string
	Done bool
}

// TodoSvc ...
type TodoSvc struct {
	db *couchdb.Database
	// dbName     string
	// dbHost     string
	// dbUser     string
	// dbPassword string
}

// TodoService ...
type TodoService interface {
	Add(text string) (map[string]interface{}, error)
	List() ([]map[string]interface{}, error)
	Delete(id string) error
	Complete(id string) (map[string]interface{}, error)
}

// NewTodoService ...
func NewTodoService(dbUser, dbPassword, dbHost, dbName string) (TodoService, error) {
	var err error
	dbString := fmt.Sprintf("http://%s:%s@%s:5984/%s_rendyfebry", dbUser, dbPassword, dbHost, dbName)

	db, err := couchdb.NewDatabase(dbString)
	if err != nil {
		return nil, err
	}

	return &TodoSvc{
		db: db,
	}, nil
}

// Add ...
func (s *TodoSvc) Add(text string) (map[string]interface{}, error) {
	newDoc := map[string]interface{}{
		"_id":  uuid.NewV4().String(),
		"text": text,
		"done": false,
	}

	_, _, err := s.db.Save(newDoc, nil)
	if err != nil {
		return nil, err
	}

	return newDoc, nil
}

// List ...
func (s *TodoSvc) List() ([]map[string]interface{}, error) {
	docs, err := s.db.QueryJSON(`{"selector": {}, "limit": 1000}`)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return docs, nil
}

// Delete ...
func (s *TodoSvc) Delete(id string) error {
	err := s.db.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

// Complete ...
func (s *TodoSvc) Complete(id string) (map[string]interface{}, error) {
	doc, err := s.db.Get(id, nil)
	if err != nil {
		return nil, err
	}

	updatedDoc := map[string]interface{}{
		"_id":  doc["_id"],
		"_rev": doc["_rev"],
		"text": doc["text"],
		"done": true,
	}

	_, _, err = s.db.Save(updatedDoc, nil)
	if err != nil {
		return nil, err
	}

	return updatedDoc, nil
}
