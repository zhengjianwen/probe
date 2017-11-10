package auth

import "errors"

func Init(c *AuthConfig) {
	initRpc([]string{c.UicAddr}, []string{c.CloudAddr})
	initCookie(c.CookieSecret)
}

type AuthConfig struct {
	CookieSecret string
	UicAddr      string
	CloudAddr    string
}

func (c *AuthConfig) Validate() error {
	if len(c.CookieSecret) == 0 {
		return errors.New("param cookie secret not found")
	}

	if len(c.UicAddr) == 0 {
		return errors.New("param uic service addr not found")
	}

	if len(c.CloudAddr) == 0 {
		return errors.New("param cloud service addr not found")
	}

	return nil
}
