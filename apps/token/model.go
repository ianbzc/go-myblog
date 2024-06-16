package token

import (
	"time"

	"github.com/rs/xid"
)

type Token struct {
	UserId                int64  `json:"user_id" gorm:"user_id"`
	UserName              string `json:"username" gorm:"username"`
	AccessToken           string `json:"access_token" gorm:"access_token"`
	AccessTokenExpiresAt  int    `json:"access_token_expires_at" gorm:"access_token_expires_at"`
	RefreshToken          string `json:"refresh_token" gorm:"refresh_token"`
	RefreshTokenExpiresAt int    `json:"refresh_token_expires_at" gorm:"refresh_token_expires_at"`
	CreatedAt             int64  `json:"created_at" gorm:"created_at"`
	UpdatedAt             int64  `json:"updated_at" gorm:"updated_at`
}

func NewToken() *Token {
	return &Token{
		AccessToken:           xid.New().String(),
		AccessTokenExpiresAt:  7200,
		RefreshToken:          xid.New().String(),
		RefreshTokenExpiresAt: 3600 * 24 * 7,
		CreatedAt:             time.Now().Unix(),
	}
}

func (t *Token) TableName() string {
	return "tokens"
}
