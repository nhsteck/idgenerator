package idgenerator

import (
	"math/rand"
	"strings"
	"time"

	"github.com/chilts/sid"
	"github.com/google/uuid"
	"github.com/lithammer/shortuuid"
	"github.com/oklog/ulid"
	"github.com/rs/xid"
)

func GenUUID() string {
	id := uuid.New()
	return id.String()
}

func GenXid() string {
	guid := xid.New()
	return strings.ToUpper(guid.String())
}

func GenShortUuid() string {
	guid := shortuuid.New()
	return guid
}

func GenUlid() string {
	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	return id.String()
}

func GenSid() string {
	id := sid.Id()
	return id
}
