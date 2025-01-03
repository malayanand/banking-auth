package domain

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/malayanand/banking/errs"
	"github.com/malayanand/banking/logger"
)

type AuthToken struct {
	token *jwt.Token
}

func (t AuthToken) NewAccessToken() (string, *errs.AppError) {
	signedString, err := t.token.SignedString([]byte(HMAC_SECRET))
	if err != nil {
		logger.Error("Failed while signing the access token: " + err.Error())
		return "", errs.NewAuthenticationError("Cannot generate access token")
	}
	return signedString, nil
}

func (t AuthToken) newRefreshToken() (string, *errs.AppError) {
	c := t.token.Claims.(AccessTokenClaims)
	refreshClaims := c.RefreshTokenClaims()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	signedString, err := token.SignedString([]byte(HMAC_SECRET))
	if err != nil {
		logger.Error("Failed while signing the access token: " + err.Error())
		return "", errs.NewAuthenticationError("Cannot generate refresh token")
	}

	return signedString, nil
}

func NewAuthToken(claims AccessTokenClaims) AuthToken {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return AuthToken{token: token}
}

func NewAccessTokenFromRefreshToken(refreshToken string) (string, *errs.AppError) {
	token, err := jwt.ParseWithClaims(refreshToken, &RefreshTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(HMAC_SECRET), nil
	})

	if err != nil {
		return "", errs.NewAuthenticationError("Invalid or expired refresh token")
	}

	r := token.Claims.(*RefreshTokenClaims)
	accessTokenClaims := r.AccessTokenClaims()
	authToken := NewAuthToken(accessTokenClaims)

	return authToken.NewAccessToken()
}
