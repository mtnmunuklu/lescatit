package categorizersrps

import (
	"Lescatit/categorizer/models/categorizersmdl"
	"Lescatit/db"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const CategorizersCollection = "categories"

// CategorizersRepository is the interface of the categorizer backend.
type CategorizersRepository interface {
	Save(url *categorizersmdl.Categorizer) error
	GetById(id string) (url *categorizersmdl.Categorizer, err error)
	Update(category *categorizersmdl.Categorizer) error
}

// CRepository provides a mongo collection for categorizer job.
type CRepository struct {
	c *mgo.Collection
}

// NewCategorizersRepository creates a new CategorizersRepository instance.
func NewCategorizersRepository(conn db.Connection) CategorizersRepository {
	return &CRepository{c: conn.DB().C(CategorizersCollection)}
}

// Save adds url to database.
func (r *CRepository) Save(url *categorizersmdl.Categorizer) error {
	return r.c.Insert(url)
}

// GetById returns the url based on id.
func (r *CRepository) GetById(id string) (url *categorizersmdl.Categorizer, err error) {
	err = r.c.FindId(bson.ObjectIdHex(id)).One(&url)
	return url, err
}

// Update updates the category.
func (r *CRepository) Update(category *categorizersmdl.Categorizer) error {
	return r.c.UpdateId(category.Id, category)
}

// DeleteAll drops categorizers collection.
func (r *CRepository) DeleteAll() error {
	return r.c.DropCollection()
}
