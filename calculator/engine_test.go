package calculator

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// setup statement
	setup()

	// run tests
	e := m.Run()

	// cleaning statement
	teardown()

	// report system exit code
	os.Exit(e)
}

func setup() {
	log.Println("Setting up ....")
}

func teardown() {
	log.Println("tearing down.")
}

func TestAdd(t *testing.T) {
	defer func() {
		log.Println("Deferred tearing down")
	}()

	// Arrange (Initial)
	e := Engine{}

	// ActAssert (Implement and validity)
	actAssert := func(x, y, want float64) {
		// Act
		got := e.Add(x, y)

		// Assert
		if got != want {
			t.Errorf("Add(%.2f,%.2f) incorrect, got: %.2f, want: %.2f", x, y, got, want)
		}
	}

	t.Run("positive input", func(t *testing.T) {
		x, y := 2.5, 2.5
		want := 5
		actAssert(x, y, float64(want))
	})

	t.Run("negative input", func(t *testing.T) {
		x, y := -2.5, -3.5
		want := -6.0
		actAssert(x, y, want)
	})
}

func BenchmarkAdd(b *testing.B) {
	e := Engine{}

	// run the Add function b.N times
	for i := 0; i < b.N; i++ {
		e.Add(2, 3)
	}
}