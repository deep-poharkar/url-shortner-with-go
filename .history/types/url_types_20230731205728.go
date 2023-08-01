package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type ShortUrlBody struct {
	longUrl string `json = "long_url"`
}

type urlDb struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UrlCode   string             `bson:"url_code"`
	LongUrl   string             `bson:"long_url"`
	ShortUrl  string             `bson:"short_url"`
	CreatedAt int64              `bson:"created_at"`
	ExpiredAt int64              `bson:"expired_at"`
}
