package model

import (
	"time"
)

type Profile struct{
	ID string `bson:"id"`
	FirstName string `bson:"first_name"`
	LastName string `bson:"last_name"`
	Email string `bson:"email"`
	Password string `bson:"password"`
	CreatedAt time.Time `bson:"created_at"`
	UpdateAt time.Time `bson:"update_at"`
	
}

type Profiles []Profile