package repository

import (
	"CWS/categorization/models"
	"CWS/db"
	"CWS/security"
	"log"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Panicln(err)
	}
	cfg := db.NewConfig()
	conn, _ := db.NewConnection(cfg)
	defer conn.Close()
	r := NewCategoriesRepository(conn)
	err = r.(*categoriesRepository).DeleteAll()
	if err != nil && err.Error() != "ns not found" {
		log.Panicln(err)
	}
}

func TestCategoriesRepositorySave(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	id := bson.NewObjectId()

	category := &models.Category{
		Id:       id,
		Url:      "https://www.haberler1.com/",
		Category: "News",
		Created:  time.Now(),
		Updated:  time.Now(),
		Revision: 0,
	}
	category.Url = security.Base64Encode(category.Url)

	r := NewCategoriesRepository(conn)
	err = r.Save(category)
	assert.NoError(t, err)

	found, err := r.GetById(category.Id.Hex())
	assert.NoError(t, err)
	assert.NotNil(t, found)
}

func TestCategoriesRepositoryGetById(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	id := bson.NewObjectId()

	category := &models.Category{
		Id:       id,
		Url:      "https://www.haberler2.com/",
		Category: "News",
		Created:  time.Now(),
		Updated:  time.Now(),
		Revision: 0,
	}
	category.Url = security.Base64Encode(category.Url)

	r := NewCategoriesRepository(conn)
	err = r.Save(category)
	assert.NoError(t, err)

	found, err := r.GetById(category.Id.Hex())
	assert.NoError(t, err)
	assert.Equal(t, category.Id, found.Id)
	assert.Equal(t, category.Url, found.Url)
	assert.Equal(t, category.Category, found.Category)
	assert.Equal(t, category.Revision, found.Revision)

	found, err = r.GetById(bson.NewObjectId().Hex())
	assert.Error(t, err)
	assert.EqualError(t, mgo.ErrNotFound, err.Error())
	assert.Nil(t, found)
}

func TestCategoriesRepositoryGetCategoryByUrl(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	id := bson.NewObjectId()

	category := &models.Category{
		Id:       id,
		Url:      "https://www.haberler3.com/",
		Category: "News",
		Created:  time.Now(),
		Updated:  time.Now(),
		Revision: 0,
	}
	category.Url = security.Base64Encode(category.Url)

	r := NewCategoriesRepository(conn)
	err = r.Save(category)
	assert.NoError(t, err)

	found, err := r.GetCategoryByUrl(category.Url)
	assert.NoError(t, err)
	assert.Equal(t, category.Id, found.Id)
	assert.Equal(t, category.Url, found.Url)
	assert.Equal(t, category.Category, found.Category)
	assert.Equal(t, category.Revision, found.Revision)

	found, err = r.GetCategoryByUrl("")
	assert.Error(t, err)
	assert.EqualError(t, mgo.ErrNotFound, err.Error())
	assert.Nil(t, found)
}

func TestCategoriesRepositoryGetAllUrlsByCategory(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	id := bson.NewObjectId()

	category := &models.Category{
		Id:       id,
		Url:      "https://www.haberler4.com/",
		Category: "News",
		Created:  time.Now(),
		Updated:  time.Now(),
		Revision: 0,
	}
	category.Url = security.Base64Encode(category.Url)

	id2 := bson.NewObjectId()

	category2 := &models.Category{
		Id:       id2,
		Url:      "https://www.haberler5.com/",
		Category: "News",
		Created:  time.Now(),
		Updated:  time.Now(),
		Revision: 0,
	}
	category2.Url = security.Base64Encode(category2.Url)

	r := NewCategoriesRepository(conn)
	err = r.Save(category)
	assert.NoError(t, err)

	err = r.Save(category2)
	assert.NoError(t, err)

	founds, err := r.GetAllUrlsByCategory(category.Category, 10)
	assert.NoError(t, err)
	assert.NotEmpty(t, founds)
}

func TestCategoriesRepositoryUpdate(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	id := bson.NewObjectId()

	category := &models.Category{
		Id:       id,
		Url:      "https://www.haberler6.com/",
		Category: "Gambling",
		Created:  time.Now(),
		Updated:  time.Now(),
		Revision: 0,
	}
	category.Url = security.Base64Encode(category.Url)

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

func TestCategoriesRepositoryDelete(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	id := bson.NewObjectId()

	category := &models.Category{
		Id:       id,
		Url:      "https://www.haberler.com/",
		Category: "News",
		Created:  time.Now(),
		Updated:  time.Now(),
		Revision: 0,
	}
	category.Url = security.Base64Encode(category.Url)

	r := NewCategoriesRepository(conn)
	err = r.Save(category)
	assert.NoError(t, err)

	found, err := r.GetById(category.Id.Hex())
	assert.NoError(t, err)
	assert.NotNil(t, found)

	err = r.Delete(category.Id.Hex())
	assert.NoError(t, err)

	found, err = r.GetById(category.Id.Hex())
	assert.Error(t, err)
	assert.EqualError(t, mgo.ErrNotFound, err.Error())
	assert.Nil(t, found)
}
