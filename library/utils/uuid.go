package utils

import (
	"github.com/gofrs/uuid"
	//"github.com/satori/go.uuid"
	"strings"
)

// CreateUUID is creation uuid
func CreateUUID() string {

	//u1 := uuid.Must(uuid.NewV4())
	//fmt.Printf("UUIDv4: %s\n", u1)

	// or error handling
	u2, err := uuid.NewV4()

	if err != nil {
		//fmt.Printf("Something went wrong: %s", err)
		return ""
	}

	//fmt.Printf("UUIDv4: %s\n", u2)

	v4 := strings.Replace(u2.String(), "-", "", -1)
	return v4
}
