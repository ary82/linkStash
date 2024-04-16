package utils

import "net/http"

// Set a cookie with this config
func SetCookie(w http.ResponseWriter, name string, cookieVal string) {
	cookie := &http.Cookie{
		Name:     name,
		Value:    cookieVal,
		Path:     "/",
		MaxAge:   24 * 3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, cookie)
}

// Clears the cookie with name parameter
func ClearJwtCookie(w http.ResponseWriter, name string) {
	cookie := &http.Cookie{
		Name:     name,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, cookie)
}
