package classifiersmdl

import (
	"Lescatit/pb"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Classifier provides the cmodel instance for classifer job.
type Classifier struct {
	Id       bson.ObjectId `bson:"_id"`
	Name     string        `bson:"name"`
	Category string        `bson:"category"`
	Created  time.Time     `bson:"created"`
	Updated  time.Time     `bson:"updated"`
	Revision string        `bson:"revision"`
	Data     string        `bson:"data"`
}

// ToProtoBuffer converts the classifier structure into a protocol buffer classifer structure.
func (c *Classifier) ToProtoBuffer() *pb.Classifier {
	return &pb.Classifier{
		Id:       c.Id.Hex(),
		Name:     c.Name,
		Category: c.Category,
		Created:  c.Created.Unix(),
		Updated:  c.Updated.Unix(),
		Revision: c.Revision,
		Data:     c.Data,
	}
}

// FromProtoBuffer gets data from protocol buffer and converts to the classifer structure.
func (c *Classifier) FromProtoBuffer(category *pb.Classifier) {
	c.Id = bson.ObjectIdHex(category.GetId())
	c.Name = category.GetName()
	c.Category = category.GetCategory()
	c.Created = time.Unix(category.GetCreated(), 0)
	c.Updated = time.Unix(category.GetUpdated(), 0)
	c.Revision = category.GetRevision()
	c.Data = category.GetData()
}
