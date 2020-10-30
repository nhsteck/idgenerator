package idgenerator

import (
	"fmt"
	"testing"
	"time"
)

const LIMIT = 1000000

func TestUUID(t *testing.T) {
	startTime := time.Now()
	for i := 0; i < LIMIT; i++ {
		_ = GenUUID()
	}
	allTime := time.Now().Sub(startTime)
	fmt.Printf("TestUUID \t >> %d\n", allTime)
}

func TestXid(t *testing.T) {
	startTime := time.Now()
	for i := 0; i < LIMIT; i++ {
		_ = GenXid()
	}
	allTime := time.Now().Sub(startTime)
	fmt.Printf("TestGenXid \t >> %d\n", allTime)
}

func TestShortUUID(t *testing.T) {
	t.SkipNow()
	startTime := time.Now()
	for i := 0; i < LIMIT; i++ {
		_ = GenShortUuid()
	}
	allTime := time.Now().Sub(startTime)
	fmt.Printf("TestShortUUID \t >> %d\n", allTime)
}

func TestGenUlid(t *testing.T) {
	t.SkipNow()
	startTime := time.Now()
	for i := 0; i < LIMIT; i++ {
		_ = GenUlid()
	}
	allTime := time.Now().Sub(startTime)
	fmt.Printf("GenUlid \t >> %d\n", allTime)
}

func TestGenSid(t *testing.T) {
	startTime := time.Now()
	for i := 0; i < LIMIT; i++ {
		_ = GenSid()
	}
	allTime := time.Now().Sub(startTime)
	fmt.Printf("GenSid \t\t >> %d\n", allTime)
}
