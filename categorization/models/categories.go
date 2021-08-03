package models

import (
	"CWS/pb"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Category provides the category instance for categorization job.
type Category struct {
	Id       bson.ObjectId `bson:"_id"`
	Url      string        `bson:"url"`
	Category string        `bson:"category"`
	Created  time.Time     `bson:"created"`
	Updated  time.Time     `bson:"updated"`
	Revision string        `bson:"revision"`
}

// ToProtoBuffer converts the category structure into a protocol buffer category structure.
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

// FromProtoBuffer gets category from protocol buffer and converts to the category structure.
func (c *Category) FromProtoBuffer(category *pb.Category) {
	c.Id = bson.ObjectIdHex(category.GetId())
	c.Url = category.GetUrl()
	c.Category = category.GetCategory()
	c.Created = time.Unix(category.GetCreated(), 0)
	c.Updated = time.Unix(category.GetUpdated(), 0)
	c.Revision = category.GetRevision()
}
