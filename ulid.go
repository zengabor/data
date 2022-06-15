package data

import (
	"math/rand"
	// "strings"
	"time"

	"github.com/oklog/ulid/v2"
)

func NewULID() string {
	entropy := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	u, err := ulid.New(ulid.Timestamp(time.Now().UTC()), entropy)
	if err != nil {
		return NewUUID4Str()
	}
	return string(u.String())
}

// func NewLowerCaseULID(t time.Time) string {
// 	return strings.ToLower(NewULID(t))
// }
