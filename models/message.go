package models

import "time"

type Message struct {
	ID         string    `bson:"_id,omitempty" json:"id"`
	SenderID   string    `bson:"sender_id" json:"sender_id"`
	ReceiverID string    `bson:"receiver_id" json:"receiver_id"`
	Text       string    `bson:"text" json:"text"`
	CreatedAt  time.Time `bson:"created_at" json:"created_at"`
}
