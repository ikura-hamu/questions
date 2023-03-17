package domain

import "github.com/google/uuid"

type Member struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	DisplayName string    `json:"string"`
}

func NewMember(id uuid.UUID, name string, displayName string) Member {
	return Member{
		Id:          id,
		Name:        name,
		DisplayName: displayName,
	}
}
