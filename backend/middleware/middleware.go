package middleware

import "backend/external/sso"

type Client struct {
	SSOExternal *sso.Client
}
