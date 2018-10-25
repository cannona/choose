package choose

import "testing"

func didPanic(f func()) (ret bool) {
	defer func() {
		if r := recover(); r != nil {
			ret = true
		}
	}()
	f()
	return
}

func TestChoose(t *testing.T) {
	// Let's use Pascal's triangle for the values.
	row := append([]int64(nil), 1)
	for n := int64(0); n < 62; n++ {
		for k, v := range row {
			if x := Choose(n, int64(k)); x != v {
				t.Fatalf("%v choose %v returned %v, not %v", n, k, x, v)
			}
		}
		newRow := make([]int64, len(row)+1)
		newRow[0] = 1
		for i := 1; i < len(row); i++ {
			newRow[i] = row[i-1] + row[i]
		}
		newRow[len(row)] = 1
		row = newRow
	}
	// Ensure it panics when it should.
	if !didPanic(func() {
		Choose(3, 4)
	}) {
		t.Fatal("Choose did not panic when k > n.")
	}
	if !didPanic(func() {
		Choose(3, -1)
	}) {
		t.Fatal("Choose did not panic when k < 0")
	}
}

func BenchmarkChoose(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Choose(61, 30)
	}
}
