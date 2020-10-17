package hw02_unpack_string //nolint:golint,stylecheck
//package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type test struct {
	input    string
	expected string
	err      error
}

func TestUnpack(t *testing.T) {
	for _, tst := range [...]test{
		{
			input:    "a4bc2d5e",
			expected: "aaaabccddddde",
		},
		{
			input:    "abccd",
			expected: "abccd",
		},
		{
			input:    "3abc",
			expected: "",
			err:      ErrInvalidString,
		},
		{
			input:    "45",
			expected: "",
			err:      ErrInvalidString,
		},
		{
			input:    "aaa10b",
			expected: "",
			err:      ErrInvalidString,
		},
		{
			input:    "",
			expected: "",
		},
		{
			input:    "aaa0b",
			expected: "aab",
		},
	} {
		result, err := Unpack(tst.input)
		require.Equal(t, tst.err, err)
		require.Equal(t, tst.expected, result)
	}
}

func TestUnpackWithEscape(t *testing.T) {
	//t.Skip() // Remove if task with asterisk completed

	for _, tst := range [...]test{
		{
			input:    `qwe\4\5`,
			expected: `qwe45`,
		},
		{
			input:    `qwe\45`,
			expected: `qwe44444`,
		},
		{
			input:    `qwe\\5`,
			expected: `qwe\\\\\`,
		},
		{
			input:    `qwe\\\3`,
			expected: `qwe\3`,
		},
	} {
		result, err := Unpack(tst.input)
		require.Equal(t, tst.err, err)
		require.Equal(t, tst.expected, result)
	}
}

func TestTranslateString(t *testing.T) {
	testSource := []struct {
		source   string
		expected string
	}{
		{`a4c`, `aaaac`},
		{`a\\b\5`, `a\b5`},
		{`\\5`, `\\\\\`},
		{`абв`, `абв`},
		{`зздд4`, `ззддддд`},
		{"as\n5", "as\n\n\n\n\n"},
	}
	for _, tc := range testSource {
		tc := tc
		t.Run(tc.source, func(t *testing.T) {
			got, err := translateString(tc.source)
			require.NoError(t, err)
			require.Equal(t, tc.expected, got)
		})
	}

}

func TestCheckString(t *testing.T) {
	testSource := []struct {
		source  string
		matched bool
	}{
		{`1s`, false},
		{`s1`, true},
		{`\\\`, false},
		{`\55d`, true},
		{`  a`, true},
		{`s2\\2\3\\5`, true},
		{`asd{}`, false},
		{`a\\\n5`, false},
		{`\\\\\\\\\\`, true},
		{`    5`, true},
		{"as\n5", true},
	}
	for _, tc := range testSource {
		tc := tc
		t.Run(tc.source, func(t *testing.T) {
			got := CheckString(tc.source)
			require.Equal(t, tc.matched, got)
		})
	}
}
