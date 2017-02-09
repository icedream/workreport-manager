package project_test

import (
	"testing"
	"time"

	"git.dekart811.net/icedream/workreportmgr/project"

	"github.com/stretchr/testify/require"
)

func TestDate_String(t *testing.T) {
	date := project.Date{time.Date(2016, 8, 22, 11, 22, 33, 97, time.Local)}
	require.Equal(t, "22.08.2016", date.String())
}

func TestDate_UnmarshalJSON_Quotes(t *testing.T) {
	date := project.Date{}
	require.Nil(t, date.UnmarshalJSON([]byte(`"2016-08-22"`)))
	require.Equal(t, 2016, date.Time.Year())
	require.Equal(t, time.Month(8), date.Time.Month())
	require.Equal(t, 22, date.Time.Day())
}

func TestDate_UnmarshalJSON_NoQuotes(t *testing.T) {
	date := project.Date{}
	require.Nil(t, date.UnmarshalJSON([]byte(`2016-08-22`)))
	require.Equal(t, 2016, date.Time.Year())
	require.Equal(t, time.Month(8), date.Time.Month())
	require.Equal(t, 22, date.Time.Day())
}
