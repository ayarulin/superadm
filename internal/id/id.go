package id

import "github.com/google/uuid"

func Random() string {
	uuid, err := uuid.NewRandom()
	if err != nil {
		panic("Unable to generate uniq Id")
	}
	return uuid.String()
}
