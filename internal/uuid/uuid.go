package uuid

import (
	"github.com/google/uuid"
)

func Generate() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

func MustGenerate() string {
	id, err := Generate()
	if err != nil {
		panic(err)
	}
	return id
}
