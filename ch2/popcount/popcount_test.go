package popcount

import "testing"

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(1454156789)
	}
}

func BenchmarkPopCount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount2(1454156789)
	}
}

func BenchmarkPopCount3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount3(1454156789)
	}
}

func BenchmarkPopCount4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount4(1454156789)
	}
}

func TestPopCount(t *testing.T) {
	if PopCount(0) != 0 {
		t.Error(`PopCount(0) not equal 0`)
	}
	if PopCount(1) != 1 {
		t.Error(`PopCount(1) not equal 1`)
	}
	if PopCount(2) != 1 {
		t.Error(`PopCount(2) not equal 2`)
	}
	if PopCount(7) != 3 {
		t.Error(`PopCount(7) not equal 3`)
	}
}

func TestPopCount2(t *testing.T) {
	if PopCount2(0) != 0 {
		t.Error(`PopCount2(0) not equal 0`)
	}
	if PopCount2(1) != 1 {
		t.Error(`PopCount2(1) not equal 1`)
	}
	if PopCount2(2) != 1 {
		t.Error(`PopCount2(2) not equal 2`)
	}
	if PopCount2(7) != 3 {
		t.Error(`PopCount2(7) not equal 3`)
	}
}

func TestPopCount3(t *testing.T) {
	if PopCount3(0) != 0 {
		t.Error(`PopCount3(0) not equal 0`)
	}
	if PopCount3(1) != 1 {
		t.Error(`PopCount3(1) not equal 1`)
	}
	if PopCount3(2) != 1 {
		t.Error(`PopCount3(2) not equal 2`)
	}
	if PopCount3(7) != 3 {
		t.Error(`PopCount3(7) not equal 3`)
	}
}

func TestPopCount4(t *testing.T) {
	if PopCount4(0) != 0 {
		t.Error(`PopCount4(0) not equal 0`)
	}
	if PopCount4(1) != 1 {
		t.Error(`PopCount4(1) not equal 1`)
	}
	if PopCount4(2) != 1 {
		t.Error(`PopCount4(2) not equal 2`)
	}
	if PopCount4(7) != 3 {
		t.Error(`PopCount4(7) not equal 3`)
	}
}


