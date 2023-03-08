package domain

import "time"

type Question struct {
	Id        int
	Question  string
	Answer    string
	CreatedAt time.Time
}
