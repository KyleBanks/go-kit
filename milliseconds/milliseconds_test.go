package milliseconds

import (
	"testing"
	"time"
)

func TestToMilliseconds(t *testing.T) {
	res := ToMilliseconds(time.Now())
	if res <= 0 {
		t.Error("Unexpected time returned: ", res)
	}
}

func TestNowInMilliseconds(t *testing.T) {
	now := NowInMilliseconds()
	if now <= 0 {
		t.Error("Unexpected time returned: ", now)
	}
}
