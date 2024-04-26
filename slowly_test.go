package slowly

import (
	"log"
	"testing"
	"time"
)

func TestSlowly(t *testing.T) {
	ticker := NewTicker(1, time.Second)
	defer ticker.Stop()
	for i := 0; i < 10; i++ {
		ticker.Exec()
		log.Println(i, time.Now().UnixMicro())
	}
}
