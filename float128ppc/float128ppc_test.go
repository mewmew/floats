package float128ppc

import (
	"math/big"
	"testing"
)

func TestRoundTrip(t *testing.T) {
	golden := []struct {
		h, l uint64
	}{
		{h: 0x0000000000000000, l: 0x0000000000000000}, // "0xM00000000000000000000000000000000"
		{h: 0x3DF0000000000000, l: 0x0000000000000000}, // "0xM3DF00000000000000000000000000000"
		{h: 0x3FF0000000000000, l: 0x0000000000000000}, // "0xM3FF00000000000000000000000000000"
		{h: 0x4000000000000000, l: 0x0000000000000000}, // "0xM40000000000000000000000000000000"
		{h: 0x400C000000000030, l: 0x0000000010000000}, // "0xM400C0000000000300000000010000000"
		{h: 0x400F000000000000, l: 0xBCB0000000000000}, // "0xM400F000000000000BCB0000000000000"
		{h: 0x403B000000000000, l: 0x0000000000000000}, // "0xM403B0000000000000000000000000000"
		{h: 0x405EDA5E353F7CEE, l: 0x0000000000000000}, // "0xM405EDA5E353F7CEE0000000000000000"
		{h: 0x4093B40000000000, l: 0x0000000000000000}, // "0xM4093B400000000000000000000000000"
		{h: 0x41F0000000000000, l: 0x0000000000000000}, // "0xM41F00000000000000000000000000000"
		{h: 0x4D436562A0416DE0, l: 0x0000000000000000}, // "0xM4D436562A0416DE00000000000000000"
		{h: 0x8000000000000000, l: 0x0000000000000000}, // "0xM80000000000000000000000000000000"
		{h: 0x818F2887B9295809, l: 0x800000000032D000}, // "0xM818F2887B9295809800000000032D000"
		{h: 0xC00547AE147AE148, l: 0x3CA47AE147AE147A}, // "0xMC00547AE147AE1483CA47AE147AE147A"
	}
	for _, g := range golden {
		f1 := NewFromBits(g.h, g.l)
		fbig, nan := f1.Big()
		_ = nan
		f2, acc := NewFromBig(fbig)
		_ = acc
		h, l := f2.Bits()
		if g.h != h {
			t.Errorf("0xM%016X%016X; high mismatch; expected 0x%016X, got 0x%016X", g.h, g.l, g.h, h)
		}
		if g.l != l {
			t.Errorf("0xM%016X%016X; low mismatch; expected 0x%016X, got 0x%016X", g.h, g.l, g.l, l)
		}
		if acc != big.Exact {
			t.Errorf("0xM%016X%016X; round-trip result accuracy inexact; expected %v, got %v", g.h, g.l, big.Exact, acc)
		}
	}
}
