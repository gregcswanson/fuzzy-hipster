package interfaces

import (
  "net/http"
)

type BaseRepository struct {
	request  *http.Request
}