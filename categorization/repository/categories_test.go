package repository

import (
	"log"
	"testing"
	"time"

	"github.com/mtnmunuklu/lescatit/categorization/models"
	"github.com/mtnmunuklu/lescatit/db"
	"github.com/mtnmunuklu/lescatit/security"

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

	r := NewCategoriesRepository(conn)
	err = r.(*CRepository).DeleteAll()
	if err != nil && err.Error() != "ns not found" {
		log.Panicln(err)
	}
}

// TestCategoriesRepositorySave tests the operation to url add.
func TestCategoriesRepositorySave(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	id := bson.NewObjectId()

	category := &models.Category{
		Id:       id,
		Url:      "https://www.examplect.com/",
		Category: "News",
		Created:  time.Now(),
		Updated:  time.Now(),
		Revision: "0",
		Data:     "Data",
	}
	category.Url = security.Base64Encode(category.Url)
	category.Data = security.Base64Encode(category.Data)

	r := NewCategoriesRepository(conn)
	err = r.Save(category)
	assert.NoError(t, err)

	found, err := r.GetById(category.Id.Hex())
	assert.NoError(t, err)
	assert.NotNil(t, found)
}

// TestCategoriesRepositoryGetById tests the operation to return url based on id.
func TestCategoriesRepositoryGetById(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	id := bson.NewObjectId()

	category := &models.Category{
		Id:       id,
		Url:      "https://www.examplect2.com/",
		Category: "News",
		Created:  time.Now(),
		Updated:  time.Now(),
		Revision: "0",
		Data:     "Data",
	}
	category.Url = security.Base64Encode(category.Url)
	category.Data = security.Base64Encode(category.Data)

	r := NewCategoriesRepository(conn)
	err = r.Save(category)
	assert.NoError(t, err)

	found, err := r.GetById(category.Id.Hex())
	assert.NoError(t, err)
	assert.Equal(t, category.Id, found.Id)
	assert.Equal(t, category.Url, found.Url)
	assert.Equal(t, category.Category, found.Category)
	assert.Equal(t, category.Revision, found.Revision)
	assert.Equal(t, category.Data, found.Data)

	found, err = r.GetById(bson.NewObjectId().Hex())
	assert.Error(t, err)
	assert.EqualError(t, mgo.ErrNotFound, err.Error())
	assert.Nil(t, found)
}

// TestCategoriesRepositoryGetCategoryByURL tests the operation to return category based on url.
func TestCategoriesRepositoryGetCategoryByURL(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	id := bson.NewObjectId()

	category := &models.Category{
		Id:       id,
		Url:      "https://www.examplect3.com/",
		Category: "News",
		Created:  time.Now(),
		Updated:  time.Now(),
		Revision: "0",
		Data:     "Data",
	}
	category.Url = security.Base64Encode(category.Url)
	category.Data = security.Base64Encode(category.Data)

	r := NewCategoriesRepository(conn)
	err = r.Save(category)
	assert.NoError(t, err)

	found, err := r.GetCategoryByURL(category.Url)
	assert.NoError(t, err)
	assert.Equal(t, category.Id, found.Id)
	assert.Equal(t, category.Url, found.Url)
	assert.Equal(t, category.Category, found.Category)
	assert.Equal(t, category.Revision, found.Revision)
	assert.Equal(t, category.Data, found.Data)

	found, err = r.GetCategoryByURL("")
	assert.Error(t, err)
	assert.EqualError(t, mgo.ErrNotFound, err.Error())
	assert.Nil(t, found)
}

// TestCategoriesRepositoryGetAllURLsByCategory tests the operation to return all urls based on category.
func TestCategoriesRepositoryGetAllURLsByCategory(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	id := bson.NewObjectId()

	category := &models.Category{
		Id:       id,
		Url:      "https://www.examplect4.com/",
		Category: "News",
		Created:  time.Now(),
		Updated:  time.Now(),
		Revision: "0",
		Data:     "Data",
	}
	category.Url = security.Base64Encode(category.Url)
	category.Data = security.Base64Encode(category.Data)

	id2 := bson.NewObjectId()

	category2 := &models.Category{
		Id:       id2,
		Url:      "https://www.examplect5.com/",
		Category: "News",
		Created:  time.Now(),
		Updated:  time.Now(),
		Revision: "0",
		Data:     "Data",
	}
	category2.Url = security.Base64Encode(category2.Url)
	category2.Data = security.Base64Encode(category2.Data)

	r := NewCategoriesRepository(conn)
	err = r.Save(category)
	assert.NoError(t, err)

	err = r.Save(category2)
	assert.NoError(t, err)

	founds, err := r.GetAllURLsByCategory(category.Category, 10)
	assert.NoError(t, err)
	assert.NotEmpty(t, founds)
}

// TestCategoriesRepositoryUpdate tests the category update operation.
func TestCategoriesRepositoryUpdate(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	id := bson.NewObjectId()

	category := &models.Category{
		Id:       id,
		Url:      "https://www.examplect6.com/",
		Category: "Gambling",
		Created:  time.Now(),
		Updated:  time.Now(),
		Revision: "0",
		Data:     "Data",
	}
	category.Url = security.Base64Encode(category.Url)
	category.Data = security.Base64Encode(category.Data)

	r := NewCategoriesRepository(conn)
	err = r.Save(category)
	assert.NoError(t, err)

	found, err := r.GetById(category.Id.Hex())
	assert.NoError(t, err)
	assert.NotNil(t, found)

	category.Category = "News"
	err = r.Update(category)
	assert.NoError(t, err)

	found, err = r.GetById(category.Id.Hex())
	assert.NoError(t, err)
	assert.Equal(t, "News", found.Category)

}

// TestCategoriesRepositoryDelete tests the url delete operation.
func TestCategoriesRepositoryDelete(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	id := bson.NewObjectId()

	category := &models.Category{
		Id:       id,
		Url:      "https://www.examplect7.com/",
		Category: "News",
		Created:  time.Now(),
		Updated:  time.Now(),
		Revision: "0",
		Data:     "Data",
	}
	category.Url = security.Base64Encode(category.Url)
	category.Data = security.Base64Encode(category.Data)

	r := NewCategoriesRepository(conn)
	err = r.Save(category)
	assert.NoError(t, err)

	found, err := r.GetById(category.Id.Hex())
	assert.NoError(t, err)
	assert.NotNil(t, found)

	err = r.DeleteById(category.Id.Hex())
	assert.NoError(t, err)

	found, err = r.GetById(category.Id.Hex())
	assert.Error(t, err)
	assert.EqualError(t, mgo.ErrNotFound, err.Error())
	assert.Nil(t, found)
}
