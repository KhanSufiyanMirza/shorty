package utils

import (
	"strings"

	"github.com/google/uuid"
)

func GenerateMessageId(id string) string {
	uuid := strings.Replace(uuid.New().String(), "-", "", -1)
	return strings.ToUpper(id) + uuid
}
