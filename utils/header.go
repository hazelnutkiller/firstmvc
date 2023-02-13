package utils

import "net/textproto"

func (h Header) Get(key string) string {
	return textproto.MIMEHeader(h).Get(key)
}
