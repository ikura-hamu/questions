package repository

import "errors"

var (
	ErrNoRecord = errors.New("no record")
	ErrNoToken  = errors.New("no access token")
)
