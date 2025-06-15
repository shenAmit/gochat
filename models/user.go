package models

import (
	"time"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string               `json:"name,omitempty" bson:"name,omitempty"`
	Email            string               `json:"email,omitempty" bson:"email,omitempty"`
	Username         string               `json:"username,omitempty" bson:"username,omitempty"`
	Password         string               `json:"password,omitempty" bson:"password,omitempty"`
	OTP              *string              `json:"otp,omitempty" bson:"otp,omitempty"`
	VerifiedAt       *time.Time           `json:"verified_at,omitempty" bson:"verified_at,omitempty"`
	OTPExpireAt      *time.Time           `json:"otp_expire_at,omitempty" bson:"otp_expire_at,omitempty"`
	Roles            []primitive.ObjectID `json:"roles,omitempty" bson:"roles,omitempty"`
}
