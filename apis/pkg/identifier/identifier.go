package identifier

import "github.com/google/uuid"

type ID = uuid.UUID

func New() ID {
	return uuid.New()
}

func FromString(s string) (ID, error) {
	return uuid.Parse(s)
}
