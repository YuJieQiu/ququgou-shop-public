package utils

import (
	"fmt"
	"testing"
)

func TestCreateUUID(t *testing.T) {
	uuid := CreateUUID()

	fmt.Println("test create uuid:" + uuid)
}
