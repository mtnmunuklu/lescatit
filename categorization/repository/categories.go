package repository

import (
	"github.com/mtnmunuklu/lescatit/categorization/models"
	"github.com/mtnmunuklu/lescatit/db"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const CategoriesCollection = "categories"

// CategoriesRepository is the interface of the categorization backend.
type CategoriesRepository interface {
	Save(url *models.Category) error
	GetById(id string) (url *models.Category, err error)
	GetCategoryByURL(url string) (category *models.Category, err error)
	GetAllURLsByCategory(category string, count int) (url []*models.Category, err error)
	Update(category *models.Category) error
	DeleteById(id string) error
}

// CRepository provides a mongo collection for database job.
type CRepository struct {
	c *mgo.Collection
}

// NewCategoriesRepository creates a new CategoriesRepository instance.
func NewCategoriesRepository(conn db.Connection) CategoriesRepository {
	return &CRepository{c: conn.DB().C(CategoriesCollection)}
}

// Save adds url to database.
func (r *CRepository) Save(url *models.Category) error {
	return r.c.Insert(url)
}

// GetById returns the url based on id.
func (r *CRepository) GetById(id string) (url *models.Category, err error) {
	err = r.c.FindId(bson.ObjectIdHex(id)).One(&url)
	return url, err
}

// GetCategoryByURL returns the category based on url.
func (r *CRepository) GetCategoryByURL(url string) (category *models.Category, err error) {
	err = r.c.Find(bson.M{"url": url}).One(&category)
	return category, err
}

// GetAllURLsByCategory returns all urls based on category.
func (r *CRepository) GetAllURLsByCategory(category string, count int) (url []*models.Category, err error) {
	err = r.c.Find(bson.M{"category": category}).Limit(count).All(&url)
	return url, err
}

// Update updates the category.
func (r *CRepository) Update(category *models.Category) error {
	return r.c.UpdateId(category.Id, category)
}

// Delete deletes the url based on id.
func (r *CRepository) DeleteById(id string) error {
	return r.c.RemoveId(bson.ObjectIdHex(id))
}

// DeleteAll drops categories collection.
func (r *CRepository) DeleteAll() error {
	return r.c.DropCollection()
}
