package categorizersrps

import (
	"Lescatit/categorizer/models/categorizersmdl"
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
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Panicln(err)
	}
	cfg := db.NewConfig()
	conn, _ := db.NewConnection(cfg)
	defer conn.Close()
	r := NewCategorizersRepository(conn)
	err = r.(*CRepository).DeleteAll()
	if err != nil && err.Error() != "ns not found" {
		log.Panicln(err)
	}
}

// TestCategorizersRepositorySave tests the operation to url add.
func TestCategorizersRepositorySave(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	id := bson.NewObjectId()

	url := &categorizersmdl.Categorizer{
		Id:       id,
		Url:      "https://www.examplecz.com/",
		Category: "NotCategorized",
		Created:  time.Now(),
		Updated:  time.Now(),
		Revision: "0",
		Data:     "Data",
	}
	url.Url = security.Base64Encode(url.Url)
	url.Data = security.Base64Encode(url.Data)

	r := NewCategorizersRepository(conn)
	err = r.Save(url)
	assert.NoError(t, err)

	found, err := r.GetById(url.Id.Hex())
	assert.NoError(t, err)
	assert.NotNil(t, found)
}

// TestCategorizersRepositoryGetById tests the operation to return url based on id.
func TestCategorizersRepositoryGetById(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	id := bson.NewObjectId()

	url := &categorizersmdl.Categorizer{
		Id:       id,
		Url:      "https://www.examplecz2.com/",
		Category: "News",
		Created:  time.Now(),
		Updated:  time.Now(),
		Revision: "0",
		Data:     "Data",
	}
	url.Url = security.Base64Encode(url.Url)
	url.Data = security.Base64Encode(url.Data)

	r := NewCategorizersRepository(conn)
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

// TestCategorizersRepositoryUpdate tests the operation a category update.
func TestCategorizersRepositoryUpdate(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	id := bson.NewObjectId()

	category := &categorizersmdl.Categorizer{
		Id:       id,
		Url:      "https://www.examplecz3.com/",
		Category: "Gambling",
		Created:  time.Now(),
		Updated:  time.Now(),
		Revision: "0",
		Data:     "Data",
	}
	category.Url = security.Base64Encode(category.Url)
	category.Data = security.Base64Encode(category.Data)

	r := NewCategorizersRepository(conn)
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
