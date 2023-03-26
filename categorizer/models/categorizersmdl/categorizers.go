package categorizersmdl

import (
	"time"

	"github.com/mtnmunuklu/lescatit/pb"

	"gopkg.in/mgo.v2/bson"
)

// Categorizer provides the categorizer instance for categorize job.
type Categorizer struct {
	Id       bson.ObjectId `bson:"_id"`
	Url      string        `bson:"url"`
	Category string        `bson:"category"`
	Created  time.Time     `bson:"created"`
	Updated  time.Time     `bson:"updated"`
	Revision string        `bson:"revision"`
	Data     string        `bson:"data"`
}

// ToProtoBuffer converts the categorizer structure into a protocol buffer categorizer structure.
func (c *Categorizer) ToProtoBuffer() *pb.Categorizer {
	return &pb.Categorizer{
		Id:       c.Id.Hex(),
		Url:      c.Url,
		Category: c.Category,
		Created:  c.Created.Unix(),
		Updated:  c.Updated.Unix(),
		Revision: c.Revision,
		Data:     c.Data,
	}
}

// FromProtoBuffer gets data from protocol buffer and converts to the categorizer structure.
func (c *Categorizer) FromProtoBuffer(category *pb.Categorizer) {
	c.Id = bson.ObjectIdHex(category.GetId())
	c.Url = category.GetUrl()
	c.Category = category.GetCategory()
	c.Created = time.Unix(category.GetCreated(), 0)
	c.Updated = time.Unix(category.GetUpdated(), 0)
	c.Revision = category.GetRevision()
	c.Data = category.GetData()
}
