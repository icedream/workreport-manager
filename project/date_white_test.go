package project

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestDate_dateFormat(t *testing.T) {
	date, err := time.Parse(dateFormat, "2013-04-05")
	require.Nil(t, err)
	require.Equal(t, 2013, date.Year())
	require.Equal(t, time.Month(4), date.Month())
	require.Equal(t, 5, date.Day())
}
