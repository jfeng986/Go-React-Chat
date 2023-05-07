package models

type User struct {
	ID       string `bson:"_id,omitempty" json:"id"`
	Username string `bson:"username" json:"username" binding:"required"`
	Password string `bson:"password" json:"password" binding:"required"`
}

type RegisterRequest struct {
	Username string `bson:"username" json:"username" binding:"required"`
	Password string `bson:"password" json:"password" binding:"required"`
}

type RegisterResponse struct {
	ID       string `bson:"_id,omitempty" json:"id"`
	Username string `bson:"username" json:"username" binding:"required"`
}

type LoginRequest struct {
	Username string `bson:"username" json:"username" binding:"required"`
	Password string `bson:"password" json:"password" binding:"required"`
}

type LoginResponse struct {
	ID       string `bson:"_id,omitempty" json:"id"`
	Username string `bson:"username" json:"username" binding:"required"`
}
