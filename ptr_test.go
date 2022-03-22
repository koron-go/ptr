package ptr_test

import (
	"fmt"
	"testing"

	"github.com/koron-go/ptr"
)

func checkTV[T any](t *testing.T, p *T, want string) {
	t.Helper()
	got := fmt.Sprintf("%T %#v", p, *p)
	if got != want {
		t.Errorf("unmatch want=%q got=%q", want, got)
	}
}

type Item0 struct {}

func TestP(t *testing.T) {
	checkTV(t, ptr.P(1), "*int 1")
	checkTV(t, ptr.P(-123), "*int -123")
	checkTV(t, ptr.P(float64(1)), "*float64 1")
	checkTV(t, ptr.P("abc"), "*string \"abc\"")
	checkTV(t, ptr.P(Item0{}), "*ptr_test.Item0 ptr_test.Item0{}")
}

func checkT[T any](t *testing.T, p *T, want string) {
	t.Helper()
	got := fmt.Sprintf("%T", p)
	if got != want {
		t.Errorf("unmatch want=%q got=%q", want, got)
	}
}

func TestPP(t *testing.T) {
	checkT(t, ptr.P(ptr.P(1)), "**int")
	checkT(t, ptr.P(ptr.P(-123)), "**int")
	checkT(t, ptr.P(ptr.P(float64(1))), "**float64")
	checkT(t, ptr.P(ptr.P("abc")), "**string")
	checkT(t, ptr.P(ptr.P(Item0{})), "**ptr_test.Item0")
}

func checkV[T any](t *testing.T, v T, want string) {
	t.Helper()
	got := fmt.Sprintf("%[1]T %#[1]v", v)
	if got != want {
		t.Errorf("unmatch want=%q got=%q", want, got)
	}
}

func TestV(t *testing.T) {
	checkV(t, ptr.V(ptr.P(1)), "int 1")
	checkV(t, ptr.V(ptr.P(-123)), "int -123")
	checkV(t, ptr.V(ptr.P(float64(1))), "float64 1")
	checkV(t, ptr.V(ptr.P("abc")), "string \"abc\"")
	checkV(t, ptr.V(ptr.P(Item0{})), "ptr_test.Item0 ptr_test.Item0{}")
}

func TestVNil(t *testing.T) {
	checkV(t, ptr.V((*int)(nil)), "int 0")
	checkV(t, ptr.V((*float64)(nil)), "float64 0")
	checkV(t, ptr.V((*string)(nil)), "string \"\"")
	checkV(t, ptr.V((*Item0)(nil)), "ptr_test.Item0 ptr_test.Item0{}")
}
