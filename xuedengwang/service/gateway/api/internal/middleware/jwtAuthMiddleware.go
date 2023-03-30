package middleware

import (
	"context"
	"xuedengwang/core/errx"
	"xuedengwang/core/jwtx"
	"xuedengwang/core/result"

	"net/http"
)

const (
	jwtAudience  = "aud"
	jwtExpire    = "exp"
	jwtId        = "jti"
	jwtIssueAt   = "iat"
	jwtIssuer    = "iss"
	jwtNotBefore = "nbf"
	jwtSubject   = "sub"
)

type JwtAuthMiddleware struct {
	secret string
}

func NewJwtAuthMiddleware(secret string) *JwtAuthMiddleware {
	return &JwtAuthMiddleware{
		secret: secret,
	}
}

func (m *JwtAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		claims, err := jwtx.ParseRequest(r, m.secret)
		if err != nil {
			result.AuthHttpResult(r, w, nil, errx.NewErrCode(errx.TOKEN_EXPIRE_ERROR))
			return
		}

		ctx := r.Context()
		for k, v := range claims {
			switch k {
			case jwtAudience, jwtExpire, jwtId, jwtIssueAt, jwtIssuer, jwtNotBefore, jwtSubject:
			default:
				ctx = context.WithValue(ctx, k, v)
			}
		}

		next.ServeHTTP(w, r.WithContext(ctx))

	}
}
