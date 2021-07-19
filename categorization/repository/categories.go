package repository

import (
	"CWS/categorization/models"
	"CWS/db"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const CategoriesCollection = "categories"

type CategoriesRepository interface {
	Save(url *models.Category) error
	GetById(id string) (url *models.Category, err error)
	GetCategoryByUrl(url string) (category *models.Category, err error)
	GetAllUrlsByCategory(category string, count int) (url []*models.Category, err error)
	Update(category *models.Category) error
	Delete(id string) error
}

type categoriesRepository struct {
	c *mgo.Collection
}

func NewCategoriesRepository(conn db.Connection) CategoriesRepository {
	return &categoriesRepository{c: conn.DB().C(CategoriesCollection)}
}

func (r *categoriesRepository) Save(url *models.Category) error {
	return r.c.Insert(url)
}

func (r *categoriesRepository) GetById(id string) (url *models.Category, err error) {
	err = r.c.FindId(bson.ObjectIdHex(id)).One(&url)
	return url, err
}

func (r *categoriesRepository) GetCategoryByUrl(url string) (category *models.Category, err error) {
	err = r.c.Find(bson.M{"url": url}).One(&category)
	return category, err
}

func (r *categoriesRepository) GetAllUrlsByCategory(category string, count int) (url []*models.Category, err error) {
	err = r.c.Find(bson.M{"category": category}).Limit(count).All(&url)
	return url, err
}

func (r *categoriesRepository) Update(category *models.Category) error {
	return r.c.UpdateId(category.Id, category)
}

func (r *categoriesRepository) Delete(id string) error {
	return r.c.RemoveId(bson.ObjectIdHex(id))
}

func (r *categoriesRepository) DeleteAll() error {
	return r.c.DropCollection()
}
