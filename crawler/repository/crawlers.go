package repository

import (
	"Lescatit/crawler/models"
	"Lescatit/db"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const CrawlersCollection = "categories"

// CrawlersRepository is the interface of the crawler backend.
type CrawlersRepository interface {
	Save(url *models.Crawler) error
	GetById(id string) (url *models.Crawler, err error)
	GetDataByURL(url string) (data *models.Crawler, err error)
}

// CRepository provides a mongo collection for database job.
type CRepository struct {
	c *mgo.Collection
}

// NewCrawlersRepository creates a new CrawlersRepository instance.
func NewCrawlersRepository(conn db.Connection) CrawlersRepository {
	return &CRepository{c: conn.DB().C(CrawlersCollection)}
}

// Save adds url to database.
func (r *CRepository) Save(url *models.Crawler) error {
	return r.c.Insert(url)
}

// GetById returns the url based on id.
func (r *CRepository) GetById(id string) (url *models.Crawler, err error) {
	err = r.c.FindId(bson.ObjectIdHex(id)).One(&url)
	return url, err
}

// GetDataByURL returns the data based on url.
func (r *CRepository) GetDataByURL(url string) (data *models.Crawler, err error) {
	err = r.c.Find(bson.M{"url": url}).One(&data)
	return data, err
}

// DeleteAll drops crawlers collection.
func (r *CRepository) DeleteAll() error {
	return r.c.DropCollection()
}
