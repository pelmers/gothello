package gothello

import "testing"

func TestMask2RC(t *testing.T) {
	r, c := Mask2RC(D4)
	b := RC2Mask(r, c)
	b.String()
	if b != D4 {
		t.Errorf("Mask-RC conversion not working.")
	}
}
