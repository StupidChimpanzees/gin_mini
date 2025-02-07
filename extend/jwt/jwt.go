package jwt

import (
	"errors"
	"gin_work/wrap/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenClaims struct {
	Ip             string
	CreateTime     time.Time
	RefreshExpires time.Time
	jwt.RegisteredClaims
}

const (
	Token = "token"

	AuthPrefix = "authorization_"
	HeaderName = "Authorization"

	TokenErr      = "token_is_empty"
	ExpiresErr    = "expires"
	RefreshExpErr = "refresh_expires"
	IpErr         = "ip_address_error"
)

func GenerateToken(uuid, domain, ip string, args ...interface{}) (string, error) {
	jwtConf := config.Mapping.JWT
	if len(args) > 0 && args[0].(int) > 0 {
		jwtConf.Expires = args[0].(int)
	}

	iJwtClaims := TokenClaims{
		Ip:             ip,
		CreateTime:     time.Now(),
		RefreshExpires: time.Now().Add(time.Duration(jwtConf.RefreshExpires) * time.Second),
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        uuid,
			Issuer:    config.Mapping.App.Name,
			Subject:   domain,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(jwtConf.Expires) * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now().Add(time.Duration(jwtConf.Issued) * time.Second)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, iJwtClaims)

	return token.SignedString([]byte(jwtConf.SignKey))
}

func ParseToken(tokenStr, domain string) (*TokenClaims, error) {
	jwtConf := config.Mapping.JWT

	now := time.Now()
	token, err := jwt.ParseWithClaims(tokenStr, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtConf.SignKey), nil
	}, jwt.WithIssuer(config.Mapping.App.Name), jwt.WithSubject(domain), jwt.WithExpirationRequired(), jwt.WithIssuedAt())
	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		return claims, nil
	} else if claims.RefreshExpires.Before(now) {
		return nil, errors.New(RefreshExpErr)
	} else if claims.ExpiresAt.Before(now) {
		return claims, errors.New(ExpiresErr)
	}

	return nil, err
}
