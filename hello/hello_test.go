package hello

import(
	"testing"
)

func TestHello(t *testing.T)  {
	v := Say("f")
	t.Error("Hello Test" + v)
}