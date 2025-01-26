package omdb

import (
	"fmt"
)

const (
	host        = "http://www.omdbapi.com"
	apiKeyParam = "apikey"
)

func (a *Adapter) baseURL() string {
	return fmt.Sprintf("%s/?%s=%s", host, apiKeyParam, a.config.Key)
}
