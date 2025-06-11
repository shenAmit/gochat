package models

import (
	"github.com/kamva/mgm/v3"
)

type User struct {
	mgm.DefaultModel `bson:",inline"` // Adds ID, CreatedAt, UpdatedAt
	Username         string           `json:"username" bson:"username"`
	Email            string           `json:"email" bson:"email"`
	Password         string           `json:"password" bson:"password"`
	CreatedAt        int64            `json:"createdAt" bson:"createdAt"`
}
