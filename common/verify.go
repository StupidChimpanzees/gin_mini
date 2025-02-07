package common

import (
	"errors"
	"gin_work/extend/jwt"
	"gin_work/wrap/log"
)

func CheckToken(accessToken, domain, ip string) (*jwt.TokenClaims, error) {
	if accessToken == "" {
		return nil, errors.New(jwt.TokenErr)
	}
	var claims *jwt.TokenClaims
	var err error
	claims, err = jwt.ParseToken(accessToken, domain)
	// access过期
	if err != nil {
		if err.Error() == jwt.ExpiresErr { // 超过有效期
			return claims, err
		}
		return nil, err
	} else if claims.Ip != ip { // IP不对应
		return nil, errors.New(jwt.IpErr)
	}
	return claims, nil
}

func RefreshToken(uuid, domain, ip string) (accessToken string, err error) {
	accessToken, err = jwt.GenerateToken(uuid, domain, ip)
	if err != nil {
		log.Warning(err.Error())
		return "", err
	}
	return accessToken, err
}
