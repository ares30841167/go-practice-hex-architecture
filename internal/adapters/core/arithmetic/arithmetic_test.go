package arithmetic_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"guanyu.dev/arithmetic/internal/adapters/core/arithmetic"
)

func TestAddition(t *testing.T) {
	arith := arithmetic.NewAdapter()

	answer, err := arith.Addition(1, 1)
	if err != nil {
		t.Fatalf("expected: %v", err)
	}

	require.Equal(t, answer, int32(2))
}

func TestSubtraction(t *testing.T) {
	arith := arithmetic.NewAdapter()

	answer, err := arith.Subtraction(1, 1)
	if err != nil {
		t.Fatalf("expected: %v", err)
	}

	require.Equal(t, answer, int32(0))
}

func TestMultiplication(t *testing.T) {
	arith := arithmetic.NewAdapter()

	answer, err := arith.Multiplication(1, 1)
	if err != nil {
		t.Fatalf("expected: %v", err)
	}

	require.Equal(t, answer, int32(1))
}

func TestDivision(t *testing.T) {
	arith := arithmetic.NewAdapter()

	answer, err := arith.Division(1, 1)
	if err != nil {
		t.Fatalf("expected: %v", err)
	}

	require.Equal(t, answer, int32(1))
}
