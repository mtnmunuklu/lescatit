package repository

import (
	"Lescatit/authentication/models"
	"Lescatit/db"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const UsersCollection = "users"

// UsersRepository is the interface of the authentication backend.
type UsersRepository interface {
	Save(user *models.User) error
	GetById(id string) (user *models.User, err error)
	GetByEmail(email string) (user *models.User, err error)
	GetAll() (user []*models.User, err error)
	Update(user *models.User) error
	DeleteById(id string) error
}

// URepository provides a mongo collection for database job.
type URepository struct {
	c *mgo.Collection
}

// NewUsersRepository creates a new UsersRepository instance.
func NewUsersRepository(conn db.Connection) UsersRepository {
	return &URepository{c: conn.DB().C(UsersCollection)}
}

// Save creates a user.
func (r *URepository) Save(user *models.User) error {
	return r.c.Insert(user)
}

// GetById returns the user based on id.
func (r *URepository) GetById(id string) (user *models.User, err error) {
	err = r.c.FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

// GetByEmail returns the user based on email.
func (r *URepository) GetByEmail(email string) (user *models.User, err error) {
	err = r.c.Find(bson.M{"email": email}).One(&user)
	return user, err
}

// GetAll returns all users.
func (r *URepository) GetAll() (user []*models.User, err error) {
	err = r.c.Find(bson.M{}).All(&user)
	return user, err
}

// Update updates the user.
func (r *URepository) Update(user *models.User) error {
	return r.c.UpdateId(user.Id, user)
}

// Delete deletes the user based on id.
func (r *URepository) DeleteById(id string) error {
	return r.c.RemoveId(bson.ObjectIdHex(id))
}

// DeleteAll drops users collection.
func (r *URepository) DeleteAll() error {
	return r.c.DropCollection()
}
