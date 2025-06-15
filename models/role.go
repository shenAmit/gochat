package models

import (
	"github.com/kamva/mgm/v3"
)

type Role struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	Status           bool   `json:"status" bson:"status"`
}
