package classifiersrps

import (
	"Lescatit/categorizer/models/classifiersmdl"
	"Lescatit/categorizer/util"
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

// Init initializes the database connection and drops classifiers collection.
func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Panicln(err)
	}

	cfg := db.NewConfig()
	conn, _ := db.NewConnection(cfg)
	defer conn.Close()

	r := NewClassifiersRepository(conn)
	err = r.(*CRepository).DeleteAll()
	if err != nil && err.Error() != "ns not found" {
		log.Panicln(err)
	}
}

// TestClassifiersRepositorySave tests the classifier add operation.
func TestClassifiersRepositorySave(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	id := bson.NewObjectId()
	name := util.GenerateRandomFileName("", ".nbc")

	classifier := &classifiersmdl.Classifier{
		Id:       id,
		Name:     name,
		Category: "NB",
		Created:  time.Now(),
		Updated:  time.Now(),
		Revision: "0",
		Data:     "Data",
	}
	classifier.Data = security.Base64Encode(classifier.Data)

	r := NewClassifiersRepository(conn)
	err = r.Save(classifier)
	assert.NoError(t, err)

	found, err := r.GetById(classifier.Id.Hex())
	assert.NoError(t, err)
	assert.NotNil(t, found)
}

// TestClassifiersRepositoryGetById tests the operation to return classifer based on id.
func TestClassifiersRepositoryGetById(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	id := bson.NewObjectId()
	name := util.GenerateRandomFileName("", ".nbc")

	classifier := &classifiersmdl.Classifier{
		Id:       id,
		Name:     name,
		Category: "NB",
		Created:  time.Now(),
		Updated:  time.Now(),
		Revision: "0",
		Data:     "Data",
	}
	classifier.Data = security.Base64Encode(classifier.Data)

	r := NewClassifiersRepository(conn)
	err = r.Save(classifier)
	assert.NoError(t, err)

	found, err := r.GetById(classifier.Id.Hex())
	assert.NoError(t, err)
	assert.Equal(t, classifier.Id, found.Id)
	assert.Equal(t, classifier.Name, found.Name)
	assert.Equal(t, classifier.Category, found.Category)
	assert.Equal(t, classifier.Revision, found.Revision)
	assert.Equal(t, classifier.Data, found.Data)

	found, err = r.GetById(bson.NewObjectId().Hex())
	assert.Error(t, err)
	assert.EqualError(t, mgo.ErrNotFound, err.Error())
	assert.Nil(t, found)
}

// TestClassifiersRepositoryGetByName tests the operation to return classifier based on name.
func TestClassifiersRepositoryGetByName(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	id := bson.NewObjectId()
	name := util.GenerateRandomFileName("", ".nbc")

	classifier := &classifiersmdl.Classifier{
		Id:       id,
		Name:     name,
		Category: "NB",
		Created:  time.Now(),
		Updated:  time.Now(),
		Revision: "0",
		Data:     "Data",
	}
	classifier.Data = security.Base64Encode(classifier.Data)

	r := NewClassifiersRepository(conn)
	err = r.Save(classifier)
	assert.NoError(t, err)

	found, err := r.GetByName(classifier.Name)
	assert.NoError(t, err)
	assert.Equal(t, classifier.Id, found.Id)
	assert.Equal(t, classifier.Name, found.Name)
	assert.Equal(t, classifier.Category, found.Category)
	assert.Equal(t, classifier.Revision, found.Revision)
	assert.Equal(t, classifier.Data, found.Data)

	found, err = r.GetByName("")
	assert.Error(t, err)
	assert.EqualError(t, mgo.ErrNotFound, err.Error())
	assert.Nil(t, found)
}

// TestClassifiersRepositoryGetAllClassiffiersByCategory tests the operation to return all classifiers based on category.
func TestClassifiersRepositoryGetAllClassiffiersByCategory(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	id := bson.NewObjectId()
	name := util.GenerateRandomFileName("", ".nbc")

	classifier := &classifiersmdl.Classifier{
		Id:       id,
		Name:     name,
		Category: "NB",
		Created:  time.Now(),
		Updated:  time.Now(),
		Revision: "0",
		Data:     "Data",
	}
	classifier.Data = security.Base64Encode(classifier.Data)

	id2 := bson.NewObjectId()
	name2 := util.GenerateRandomFileName("", ".nbc")

	classifier2 := &classifiersmdl.Classifier{
		Id:       id2,
		Name:     name2,
		Category: "NB",
		Created:  time.Now(),
		Updated:  time.Now(),
		Revision: "0",
		Data:     "Data",
	}
	classifier2.Data = security.Base64Encode(classifier2.Data)

	r := NewClassifiersRepository(conn)
	err = r.Save(classifier)
	assert.NoError(t, err)

	err = r.Save(classifier2)
	assert.NoError(t, err)

	founds, err := r.GetAllClassifiersByCategory(classifier.Category, 10)
	assert.NoError(t, err)
	assert.NotEmpty(t, founds)
}

// TestClassifiersRepositoryUpdate tests the classifier update operation.
func TestClassifiersRepositoryUpdate(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	id := bson.NewObjectId()
	name := util.GenerateRandomFileName("", ".nbc")

	classifier := &classifiersmdl.Classifier{
		Id:       id,
		Name:     name,
		Category: "NB",
		Created:  time.Now(),
		Updated:  time.Now(),
		Revision: "0",
		Data:     "Data",
	}
	classifier.Data = security.Base64Encode(classifier.Data)

	r := NewClassifiersRepository(conn)
	err = r.Save(classifier)
	assert.NoError(t, err)

	found, err := r.GetById(classifier.Id.Hex())
	assert.NoError(t, err)
	assert.NotNil(t, found)

	classifier.Category = "KNN"
	err = r.Update(classifier)
	assert.NoError(t, err)

	found, err = r.GetById(classifier.Id.Hex())
	assert.NoError(t, err)
	assert.Equal(t, "KNN", found.Category)
}

// TestClassifiersRepositoryDelete tests the classifier delete operation.
func TestClassifiersRepositoryDelete(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	id := bson.NewObjectId()
	name := util.GenerateRandomFileName("", ".nbc")

	classifier := &classifiersmdl.Classifier{
		Id:       id,
		Name:     name,
		Category: "NB",
		Created:  time.Now(),
		Updated:  time.Now(),
		Revision: "0",
		Data:     "Data",
	}
	classifier.Data = security.Base64Encode(classifier.Data)

	r := NewClassifiersRepository(conn)
	err = r.Save(classifier)
	assert.NoError(t, err)

	found, err := r.GetById(classifier.Id.Hex())
	assert.NoError(t, err)
	assert.NotNil(t, found)

	err = r.Delete(classifier.Id.Hex())
	assert.NoError(t, err)

	found, err = r.GetById(classifier.Id.Hex())
	assert.Error(t, err)
	assert.EqualError(t, mgo.ErrNotFound, err.Error())
	assert.Nil(t, found)
}
