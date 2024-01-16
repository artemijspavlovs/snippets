package utils

import (
	"time"

	"github.com/go-faker/faker/v4"
)

func GenerateRandomTimestampFromUnix() time.Time {
	return time.Unix(faker.UnixTime(), 0)
}
