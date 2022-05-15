package builduuid

import (
	"log"

	"github.com/google/uuid"
)

func GetUUID() string {
	ud, err := uuid.NewRandom()
	if err != nil {
		log.Fatal(err)
	}

	return ud.String()
}
