package models

import (
	"CWS/pb"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Category struct {
	Id       bson.ObjectId `bson:"_id"`
	Url      string        `bson:"url"`
	Category string        `bson:"category"`
	Created  time.Time     `bson:"created"`
	Updated  time.Time     `bson:"updated"`
	Revision int64         `bson:"revision"`
}

func (c *Category) ToProtoBuffer() *pb.Category {
	return &pb.Category{
		Id:       c.Id.Hex(),
		Url:      c.Url,
		Category: c.Category,
		Created:  c.Created.Unix(),
		Updated:  c.Updated.Unix(),
		Revision: c.Revision,
	}
}

func (c *Category) FromProtoBuffer(category *pb.Category) {
	c.Id = bson.ObjectIdHex(category.GetId())
	c.Url = category.GetUrl()
	c.Category = category.GetCategory()
	c.Created = time.Unix(category.GetCreated(), 0)
	c.Updated = time.Unix(category.GetUpdated(), 0)
	c.Revision = category.GetRevision()
}
