package pkg

import (
	"fmt"
	"github.com/AnakinChoi/testGoModule/internal/pkg"
)
import "testing"

func TestCallText(t *testing.T) {
	text := "test"
	got := pkg.CallText(text)
	want := fmt.Sprintf("Text is %s", text)
	if got != want {
		t.Errorf("CallText got = %s; want %s", got, want)
	}
}
