package config

import "time"

var JwtSecret = []byte("your_secret_key")

const (
	AccessTokenExpiry  = time.Minute * 15
	RefreshTokenExpiry = time.Hour * 24 * 7
)
