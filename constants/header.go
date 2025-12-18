package constants

import "net/textproto"

var (
	XApiKey       = textproto.CanonicalMIMEHeaderKey("x-api-key")
	XRequestAt    = textproto.CanonicalMIMEHeaderKey("x-request-at")
	Authorization = textproto.CanonicalMIMEHeaderKey("authorization")
)
