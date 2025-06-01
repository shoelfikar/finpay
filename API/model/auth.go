package model

import "time"

type AuthData struct {
	AccessToken         string
	Expired             time.Duration
	RefreshToken        string
	RefreshTokenExpired time.Duration
}
