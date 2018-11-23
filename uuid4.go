package uuid4

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"strings"
)

type UUID []byte

func NewUUID4() UUID {
	u := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, u)
	if n != len(u) || err != nil {
		log.Printf("error generating new UUID: %s", err)
		return nil
	}
	u[8] = u[8]&^0xc0 | 0x80
	u[6] = u[6]&^0xf0 | 0x40
	return UUID(u)
}

func NewUUID4Str() string {
	return NewUUID4().URLSafe()
}

func (u UUID) String() string {
	return fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
}

func (u UUID) URLSafe() string {
	s := base64.StdEncoding.EncodeToString(u)
	s = strings.Replace(s, "/", "_", -1)
	s = strings.Replace(s, "+", "_", -1)
	s = strings.Trim(s, "+=")
	return s
}
