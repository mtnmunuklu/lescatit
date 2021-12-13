package repository

import (
	"Lescatit/authentication/models"
	"Lescatit/db"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Init initializes the database connection and drops users collection.
func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Panicln(err)
	}

	cfg := db.NewConfig()
	conn, _ := db.NewConnection(cfg)
	defer conn.Close()

	r := NewUsersRepository(conn)
	err = r.(*URepository).DeleteAll()
	if err != nil && err.Error() != "ns not found" {
		log.Panicln(err)
	}

}

// TestUsersRepositorySave tests the user create operation.
func TestUsersRepositorySave(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	id := bson.NewObjectId()

	user := &models.User{
		Id:       id,
		Name:     "TEST",
		Email:    fmt.Sprintf("%s@email.test", id.Hex()),
		Password: "123456789",
		Role:     "user",
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	r := NewUsersRepository(conn)
	err = r.Save(user)
	assert.NoError(t, err)

	found, err := r.GetById(user.Id.Hex())
	assert.NoError(t, err)
	assert.NotNil(t, found)
}

// TestUsersRepositoryGetById tests the operation to return user based on id.
func TestUsersRepositoryGetById(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	id := bson.NewObjectId()

	user := &models.User{
		Id:       id,
		Name:     "TEST",
		Email:    fmt.Sprintf("%s@email.test", id.Hex()),
		Password: "123456789",
		Role:     "user",
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	r := NewUsersRepository(conn)
	err = r.Save(user)
	assert.NoError(t, err)

	found, err := r.GetById(user.Id.Hex())
	assert.NoError(t, err)
	assert.Equal(t, user.Id, found.Id)
	assert.Equal(t, user.Name, found.Name)
	assert.Equal(t, user.Email, found.Email)
	assert.Equal(t, user.Password, found.Password)

	found, err = r.GetById(bson.NewObjectId().Hex())
	assert.Error(t, err)
	assert.EqualError(t, mgo.ErrNotFound, err.Error())
	assert.Nil(t, found)
}

// TestUsersRepositoryGetByEmail tests the operation to return user based on email.
func TestUsersRepositoryGetByEmail(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	id := bson.NewObjectId()

	user := &models.User{
		Id:       id,
		Name:     "TEST",
		Email:    fmt.Sprintf("%s@email.test", id.Hex()),
		Password: "123456789",
		Role:     "user",
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	r := NewUsersRepository(conn)
	err = r.Save(user)
	assert.NoError(t, err)

	found, err := r.GetByEmail(user.Email)
	assert.NoError(t, err)
	assert.Equal(t, user.Id, found.Id)
	assert.Equal(t, user.Name, found.Name)
	assert.Equal(t, user.Email, found.Email)
	assert.Equal(t, user.Password, found.Password)

	found, err = r.GetByEmail("")
	assert.Error(t, err)
	assert.EqualError(t, mgo.ErrNotFound, err.Error())
	assert.Nil(t, found)
}

// TestUsersRepositoryUpdate tests the user update operation.
func TestUsersRepositoryUpdate(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	id := bson.NewObjectId()

	user := &models.User{
		Id:       id,
		Name:     "TEST",
		Email:    fmt.Sprintf("%s@email.test", id.Hex()),
		Password: "123456789",
		Role:     "user",
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	r := NewUsersRepository(conn)
	err = r.Save(user)
	assert.NoError(t, err)

	found, err := r.GetById(user.Id.Hex())
	assert.NoError(t, err)
	assert.NotNil(t, found)

	user.Name = "UpdateTest"
	err = r.Update(user)
	assert.NoError(t, err)

	found, err = r.GetById(user.Id.Hex())
	assert.NoError(t, err)
	assert.Equal(t, "UpdateTest", found.Name)
}

// TestUsersRepositoryDelete tests the user delete operation.
func TestUsersRepositoryDelete(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	id := bson.NewObjectId()

	user := &models.User{
		Id:       id,
		Name:     "TEST",
		Email:    fmt.Sprintf("%s@email.test", id.Hex()),
		Password: "123456789",
		Role:     "user",
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	r := NewUsersRepository(conn)
	err = r.Save(user)
	assert.NoError(t, err)

	found, err := r.GetById(user.Id.Hex())
	assert.NoError(t, err)
	assert.NotNil(t, found)

	err = r.DeleteById(user.Id.Hex())
	assert.NoError(t, err)

	found, err = r.GetById(user.Id.Hex())
	assert.Error(t, err)
	assert.EqualError(t, mgo.ErrNotFound, err.Error())
	assert.Nil(t, found)
}

// TestUsersRepositoryGetAll tests the operation the return all users.
func TestUsersRepositoryGetAll(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	id := bson.NewObjectId()

	user := &models.User{
		Id:       id,
		Name:     "TEST",
		Email:    fmt.Sprintf("%s@email.test", id.Hex()),
		Password: "123456789",
		Role:     "user",
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	r := NewUsersRepository(conn)
	err = r.Save(user)
	assert.NoError(t, err)

	users, err := r.GetAll()
	assert.NoError(t, err)
	assert.NotEmpty(t, users)
}
