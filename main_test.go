package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTryUnmarshal(t *testing.T) {
	data := []byte(`{	
		"accountNumber":"123456789",
		"createAt": "2020-05-27T01:02:03Z01:00",
		"updatedAt": "2020-05-28",
		"clearedAt": "",
		"type": "C"
		}`)
	exp := MyStruct{
		AccountNumber: "123456789",
		CreatedAt:     time.Time{},
		UpdatedAt:     CustomTime{time.Date(2020, 5, 28, 0, 0, 0, 0, time.UTC)},
		ClearedAt:     CustomTime{},
		Type:          "Credit",
	}
	out, err := TryUnmarshal(data)
	require.NoError(t, err)
	require.Equal(t, exp, out)
}

func TestTryMarshal(t *testing.T) {
	data := MyStruct{
		AccountNumber: "12345678",
		CreatedAt:     time.Date(2020, 1, 2, 3, 4, 5, 0, time.UTC),
		Type:          PaymentTypeCredit,
	}

	expected := []byte(
		"{\"accountNumber\":\"12345678\"," +
			"\"createdAt\":\"2020-01-02T03:04:05Z\"," +
			"\"updatedAt\":\"0001-01-01T00:00:00Z\"," +
			"\"clearedAt\":\"0001-01-01T00:00:00Z\"," +
			"\"type\":\"C\"}")
	b, err := TryMarshal(data)
	require.NoError(t, err)
	require.Equal(t, expected, b)
}
