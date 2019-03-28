package model

import (
	uuid "github.com/satori/go.uuid"
	"github.com/tMatSuZ/serverless-go-sample/pkg/datastore"
)

type PersonRepository struct {
	datastore datastore.Datastore
}

func NewPersonRepository(ds datastore.Datastore) *PersonRepository {
	return &PersonRepository{datastore: ds}
}

func (r *PersonRepository) Get(id string) (*Person, error) {
	var person *Person
	if err := r.datastore.Get(id, &person); err != nil {
		return nil, err
	}
	return person, nil
}

func (r *PersonRepository) Store(person *Person) error {
	id, _ := uuid.NewV4()
	person.ID = id.String()
	return r.datastore.Store(person)
}

func (r *PersonRepository) List() (*[]Person, error) {
	var persons *[]Person
	if err := r.datastore.List(&persons); err != nil {
		return nil, err
	}
	return persons, nil
}
