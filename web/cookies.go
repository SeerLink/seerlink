/*
 * Copyright (c) 2020-2021 The SeerLink developers
 */

package web

import (
	"net/http"
)

// FindSessionCookie returns the cookie with the "clsession" name
func FindSessionCookie(cookies []*http.Cookie) *http.Cookie {
	for _, c := range cookies {
		if c.Name == "clsession" {
			return c
		}
	}

	return nil
}
