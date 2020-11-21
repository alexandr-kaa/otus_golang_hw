package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func delTestFile() error {
	if _, err := os.Open("./test.txt"); err == nil {
		return os.Remove("./test.txt")
	}
	return nil
}

func TestCopy(t *testing.T) {
	// Place your code here
	defer func() {
		err := delTestFile()
		require.NoError(t, err)
	}()
	t.Run("length 100", func(t *testing.T) {
		err := Copy("./testdata/input.txt", "./test.txt", 0, 100)
		require.NoError(t, err)
	})
	t.Run("length 2000", func(t *testing.T) {
		err := Copy("./testdata/input.txt", "./test.txt", 0, 2000)
		require.NoError(t, err)
	})
	t.Run("length 0", func(t *testing.T) {
		err := Copy("./testdata/input.txt", "./test.txt", 0, 0)
		require.NoError(t, err)
	})

}

func TestLimitOffset(t *testing.T) {
	defer func() {
		err := delTestFile()
		require.NoError(t, err)
	}()
	err := Copy("./testdata/input.txt", "./test.txt", 100000000, 100)
	require.EqualError(t, ErrOffsetExceedsFileSize, err.Error())
}

func TestLimitLength(t *testing.T) {
	defer func() {
		err := delTestFile()
		require.NoError(t, err)
	}()
	err := Copy("./testdata/input.txt", "./test.txt", 0, 1000000000)
	require.NoError(t, err)
}
func TestMinusData(t *testing.T) {
	defer func() {
		err := delTestFile()
		require.NoError(t, err)
	}()
	t.Run("Minus offset", func(t *testing.T) {
		err := Copy("./testdata/input.txt", "./test.txt", -1, 100)
		require.EqualError(t, ErrInvalidSize, err.Error())
	})

	t.Run("Minus length", func(t *testing.T) {
		err := Copy("./testdata/input.txt", "./test.txt", 0, -100)
		require.EqualError(t, ErrInvalidSize, err.Error())
	})

}
func TestEmptyPath(t *testing.T) {
	defer func() {
		err := delTestFile()
		require.NoError(t, err)
	}()
	t.Run("Empty from", func(t *testing.T) {
		err := Copy("", "./test.txt", 1, 100)
		require.Error(t, err)
	})

	t.Run("Empty to", func(t *testing.T) {
		err := Copy("./testdata/input.txt", "", 0, 100)
		require.Error(t, err)
	})
}
