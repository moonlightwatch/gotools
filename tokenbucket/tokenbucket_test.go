package tokenbucket_test

import (
	"testing"
	"time"

	"github.com/moonlightwatch/gotools/tokenbucket"
)

func TestNewTokenBucket(t *testing.T) {
	tb := tokenbucket.NewTokenBucket(10, 1*time.Second/20)
	if tb == nil {
		t.Error("NewTokenBucket should not return nil")
	}
}

func TestTokenBucket_Take(t *testing.T) {
	tb := tokenbucket.NewTokenBucket(10, 1*time.Second/20)
	time.Sleep(1 * time.Second)
	tb.Take()
}

func TestTokenBucket_TakeNum(t *testing.T) {
	tb := tokenbucket.NewTokenBucket(10, 1*time.Second/20)
	tb.TakeNum(5)
}

func TestTokenBucket_TakeNum2(t *testing.T) {
	tb := tokenbucket.NewTokenBucket(10, 1*time.Second/20)
	if tb.TakeNum(11) != 10 {
		t.Error("TakeNum should return 10")
	}
}

func TestTokenBucket_TryTake(t *testing.T) {
	tb := tokenbucket.NewTokenBucket(10, 1*time.Second/20)
	if !tb.TryTake() {
		t.Error("TryTake should return true")
	}
}

func TestTokenBucket_TryTakeNum(t *testing.T) {
	tb := tokenbucket.NewTokenBucket(10, 1*time.Second)
	if !tb.TryTakeNum(5) {
		t.Error("TryTakeNum should return true")
	}
}

func TestTokenBucket_TryTakeNum2(t *testing.T) {
	tb := tokenbucket.NewTokenBucket(10, 1*time.Second)
	if tb.TryTakeNum(11) {
		t.Error("TryTakeNum should return false")
	}
}

func TestTokenBucket_TryTakeNum3(t *testing.T) {
	tb := tokenbucket.NewTokenBucket(10, 1*time.Second)
	if tb.TryTakeNum(5) {
		if !tb.TryTakeNum(5) {
			t.Error("TryTakeNum should return true")
		}
	} else {
		t.Error("TryTakeNum should return true")
	}
}

func TestTokenBucket_TryTakeNum4(t *testing.T) {
	tb := tokenbucket.NewTokenBucket(10, 1*time.Second)
	if tb.TryTakeNum(5) {
		if tb.TryTakeNum(6) {
			t.Error("TryTakeNum should return false")
		}
	} else {
		t.Error("TryTakeNum should return true")
	}
}

func TestTokenBucket_TryTakeNum5(t *testing.T) {
	tb := tokenbucket.NewTokenBucket(10, 1*time.Second)
	if tb.TryTakeNum(5) {
		if !tb.TryTakeNum(5) {
			t.Error("TryTakeNum should return false")
		}
	} else {
		t.Error("TryTakeNum should return true")
	}
}

func TestTokenBucket_AvailableTokens(t *testing.T) {
	tb := tokenbucket.NewTokenBucket(10, 1*time.Second)
	if tb.AvailableTokens() != 10 {
		t.Error("AvailableTokens should return 10")
	}
}

func TestTokenBucket_Capacity(t *testing.T) {
	tb := tokenbucket.NewTokenBucket(10, 1*time.Second)
	if tb.Capacity() != 10 {
		t.Error("Capacity should return 10")
	}
}

func TestTokenBucket_FillInterval(t *testing.T) {
	tb := tokenbucket.NewTokenBucket(10, 1*time.Second)
	if tb.FillInterval() != 1*time.Second {
		t.Error("FillInterval should return 1*time.Second")
	}
}

func TestTokenBucket_LastTokenTime(t *testing.T) {
	tb := tokenbucket.NewTokenBucket(10, 1*time.Second)
	tb.Take()
	if tb.LastTokenTime().UnixNano() == 0 {
		t.Error("LastTokenTime should not return 0")
	}
}

func TestTokenBucket_Reset(t *testing.T) {
	tb := tokenbucket.NewTokenBucket(10, 1*time.Second)
	tb.Take()
	tb.Reset()
	if tb.AvailableTokens() != 10 {
		t.Error("AvailableTokens should return 10")
	}
}
