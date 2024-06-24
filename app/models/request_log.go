package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RequestLog struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	RequestID string             `bson:"request_id"`
	Method    string             `bson:"method"`
	Url       string             `bson:"url"`
	Request   any                `bson:"request"`
	Response  any                `bson:"response"`
}

func (u *RequestLog) MarshalBSON() ([]byte, error) {
	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now()
	}
	u.UpdatedAt = time.Now()

	type my RequestLog
	return bson.Marshal((*my)(u))
}
