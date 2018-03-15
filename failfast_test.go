package failfast

import (
	"testing"
	"time"
)

func TestA(t *testing.T) {
	time.Sleep(5 * time.Second)
	t.Fatal("failed test")
}
func TestB(*testing.T) {
	time.Sleep(1 * time.Second)
}
