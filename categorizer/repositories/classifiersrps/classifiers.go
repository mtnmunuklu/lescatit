package classifiersrps

import (
	"Lescatit/categorizer/models/classifiersmdl"
	"Lescatit/db"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const ClassifiersCollection = "classifiers"

// ClassifierRepository is the interface of the classifier backend.
type ClassifiersRepository interface {
	Save(classifer *classifiersmdl.Classifier) error
	GetById(id string) (classifer *classifiersmdl.Classifier, err error)
	GetByName(name string) (classifer *classifiersmdl.Classifier, err error)
	GetAllClassifiersByCategory(category string, count int) (classifer []*classifiersmdl.Classifier, err error)
	Update(classifer *classifiersmdl.Classifier) error
	DeleteById(id string) error
}

// ClRepository provides a mongo collection for classifier job.
type CRepository struct {
	c *mgo.Collection
}

// NewClassifiersRepository creates a new ClassifiersRepository instance.
func NewClassifiersRepository(conn db.Connection) ClassifiersRepository {
	return &CRepository{c: conn.DB().C(ClassifiersCollection)}
}

// Save adds classifier to database.
func (r *CRepository) Save(classifer *classifiersmdl.Classifier) error {
	return r.c.Insert(classifer)
}

// GetById returns the classifier based on id.
func (r *CRepository) GetById(id string) (classifer *classifiersmdl.Classifier, err error) {
	err = r.c.FindId(bson.ObjectIdHex(id)).One(&classifer)
	return classifer, err
}

// GetByName returns the classifier based on name.
func (r *CRepository) GetByName(name string) (classifer *classifiersmdl.Classifier, err error) {
	err = r.c.Find(bson.M{"name": name}).One(&classifer)
	return classifer, err
}

// GetAllClassifiersByCategory returns all classifiers based on category.
func (r *CRepository) GetAllClassifiersByCategory(category string, count int) (classifer []*classifiersmdl.Classifier, err error) {
	err = r.c.Find(bson.M{"category": category}).Limit(count).All(&classifer)
	return classifer, err
}

// Update updates the classifier.
func (r *CRepository) Update(classifer *classifiersmdl.Classifier) error {
	return r.c.UpdateId(classifer.Id, classifer)
}

// Delete deletes the classifier based on id.
func (r *CRepository) DeleteById(id string) error {
	return r.c.RemoveId(bson.ObjectIdHex(id))
}

// DeleteAll drops classifier collection.
func (r *CRepository) DeleteAll() error {
	return r.c.DropCollection()
}
