package repository

import (
	"Lescatit/crawler/models"
	"Lescatit/db"
	"Lescatit/security"
	"log"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Init initializes the database connection and drops categories collection.
func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Panicln(err)
	}

	cfg := db.NewConfig()
	conn, _ := db.NewConnection(cfg)
	defer conn.Close()

	r := NewCrawlersRepository(conn)
	err = r.(*CRepository).DeleteAll()
	if err != nil && err.Error() != "ns not found" {
		log.Panicln(err)
	}
}

// TestCategoriesRepositorySave tests the url add operation.
func TestCrawlersRepositorySave(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	id := bson.NewObjectId()

	url := &models.Crawler{
		Id:       id,
		Url:      "https://www.examplecw.com/",
		Category: "NotCategorized",
		Created:  time.Now(),
		Updated:  time.Now(),
		Revision: "0",
		Data:     "Data",
	}
	url.Url = security.Base64Encode(url.Url)
	url.Data = security.Base64Encode(url.Data)

	r := NewCrawlersRepository(conn)
	err = r.Save(url)
	assert.NoError(t, err)

	found, err := r.GetById(url.Id.Hex())
	assert.NoError(t, err)
	assert.NotNil(t, found)
}

// TestCrawlersRepositoryGetById tests the operation to return url based on id.
func TestCrawlersRepositoryGetById(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	id := bson.NewObjectId()

	url := &models.Crawler{
		Id:       id,
		Url:      "https://www.examplecw2.com/",
		Category: "News",
		Created:  time.Now(),
		Updated:  time.Now(),
		Revision: "0",
		Data:     "Data",
	}
	url.Url = security.Base64Encode(url.Url)
	url.Data = security.Base64Encode(url.Data)

	r := NewCrawlersRepository(conn)
	err = r.Save(url)
	assert.NoError(t, err)

	found, err := r.GetById(url.Id.Hex())
	assert.NoError(t, err)
	assert.Equal(t, url.Id, found.Id)
	assert.Equal(t, url.Url, found.Url)
	assert.Equal(t, url.Category, found.Category)
	assert.Equal(t, url.Revision, found.Revision)
	assert.Equal(t, url.Data, found.Data)

	found, err = r.GetById(bson.NewObjectId().Hex())
	assert.Error(t, err)
	assert.EqualError(t, mgo.ErrNotFound, err.Error())
	assert.Nil(t, found)
}

// TestCrawlersRepositoryGetDataByURL tests the operation to return data based on url.
func TestCrawlersRepositoryGetDataByURL(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	id := bson.NewObjectId()

	url := &models.Crawler{
		Id:       id,
		Url:      "https://www.examplecw3.com/",
		Category: "News",
		Created:  time.Now(),
		Updated:  time.Now(),
		Revision: "0",
		Data:     "Data",
	}
	url.Url = security.Base64Encode(url.Url)
	url.Data = security.Base64Encode(url.Data)

	r := NewCrawlersRepository(conn)
	err = r.Save(url)
	assert.NoError(t, err)

	found, err := r.GetDataByURL(url.Url)
	assert.NoError(t, err)
	assert.Equal(t, url.Id, found.Id)
	assert.Equal(t, url.Url, found.Url)
	assert.Equal(t, url.Category, found.Category)
	assert.Equal(t, url.Revision, found.Revision)
	assert.Equal(t, url.Data, found.Data)

	found, err = r.GetDataByURL("")
	assert.Error(t, err)
	assert.EqualError(t, mgo.ErrNotFound, err.Error())
	assert.Nil(t, found)
}
