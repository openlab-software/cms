package publicid

import "github.com/google/uuid"

type PublicID string

func New() PublicID {
	return PublicID(uuid.New().String())
}
