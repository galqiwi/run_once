package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetLockPathSame(t *testing.T) {
	a, err := getLockPath("echo", nil)
	require.NoError(t, err)
	b, err := getLockPath("echo", nil)
	require.NoError(t, err)

	require.Equal(t, a, b)

	a, err = getLockPath("echo", []string{"arg"})
	require.NoError(t, err)
	b, err = getLockPath("echo", []string{"arg"})
	require.NoError(t, err)

	require.Equal(t, a, b)
}

func TestGetLockPathDiff(t *testing.T) {
	a, err := getLockPath("echo1", nil)
	require.NoError(t, err)
	b, err := getLockPath("echo2", nil)
	require.NoError(t, err)

	require.NotEqual(t, a, b)

	a, err = getLockPath("echo", []string{"arg1"})
	require.NoError(t, err)
	b, err = getLockPath("echo", []string{"arg2"})
	require.NoError(t, err)

	require.NotEqual(t, a, b)

	a, err = getLockPath("echo", []string{"arg"})
	require.NoError(t, err)
	b, err = getLockPath("echo", []string{"arg", "arg"})
	require.NoError(t, err)

	require.NotEqual(t, a, b)
}
