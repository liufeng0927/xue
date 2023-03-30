// 主要是配合jwt来生成用户登录token

package jwtx

import (
	"github.com/golang-jwt/jwt/v4/request"
	"github.com/pkg/errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	errInvalidToken = errors.New("invalid auth token")
	errNoClaims     = errors.New("no auth params")
	errExpiredToken = errors.New("token is expired")
)

// secretFunc 验证密钥格式.
func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(secret), nil
	}
}

// ParseRequest 从header获取token并解析token。
func ParseRequest(r *http.Request, secret string) (jwt.MapClaims, error) {

	token, err := doParseToken(r, secret)
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errInvalidToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errNoClaims
	}

	if int64(claims["exp"].(float64)) < jwt.TimeFunc().Unix() {
		return nil, errExpiredToken
	}

	return claims, nil
}

func doParseToken(r *http.Request, secret string) (*jwt.Token, error) {
	return request.ParseFromRequest(r, request.AuthorizationHeaderExtractor, secretFunc(secret))
}

// GenerateToken
// iss: （Issuer）签发者
// iat: （Issued At）签发时间，用Unix时间戳表示
// exp: （Expire Time）过期时间，用Unix时间戳表示
// aud: （Audience）接收该JWT的一方
// sub: （Subject）该JWT的主题
// nbf: （Not Before）不要早于这个时间
// jti: （JWT ID）用于标识JWT的唯一ID
func GenerateToken(payload map[string]interface{}, secret string, timeout int64) (tokenString string, err error) {
	now := time.Now().Unix()
	claims := make(jwt.MapClaims)
	claims["nbf"] = now
	claims["iat"] = now
	claims["exp"] = now + timeout

	for k, v := range payload {
		claims[k] = v
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 用指定的密钥对token签名.
	return token.SignedString([]byte(secret))

}
