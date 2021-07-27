package repository

import (
	"CWS/authentication/models"
	"CWS/db"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const UsersCollection = "users"

type UsersRepository interface {
	Save(user *models.User) error
	GetById(id string) (user *models.User, err error)
	GetByEmail(email string) (user *models.User, err error)
	GetAll() (user []*models.User, err error)
	Update(user *models.User) error
	Delete(id string) error
}

type URepository struct {
	c *mgo.Collection
}

func NewUsersRepository(conn db.Connection) UsersRepository {
	return &URepository{c: conn.DB().C(UsersCollection)}
}

func (r *URepository) Save(user *models.User) error {
	return r.c.Insert(user)
}

func (r *URepository) GetById(id string) (user *models.User, err error) {
	err = r.c.FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

func (r *URepository) GetByEmail(email string) (user *models.User, err error) {
	err = r.c.Find(bson.M{"email": email}).One(&user)
	return user, err
}

func (r *URepository) GetAll() (user []*models.User, err error) {
	err = r.c.Find(bson.M{}).All(&user)
	return user, err
}

func (r *URepository) Update(user *models.User) error {
	return r.c.UpdateId(user.Id, user)
}

func (r *URepository) Delete(id string) error {
	return r.c.RemoveId(bson.ObjectIdHex(id))
}

func (r *URepository) DeleteAll() error {
	return r.c.DropCollection()
}
