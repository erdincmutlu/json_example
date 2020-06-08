package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type MyStruct struct {
	AccountNumber string      `json:"accountNumber"`
	CreatedAt     time.Time   `json:"createdAt"`
	UpdatedAt     CustomTime  `json:"updatedAt"`
	ClearedAt     CustomTime  `json:"clearedAt"`
	Type          PaymentType `json:"type"`
}

func TryUnmarshal(b []byte) error {
	var out MyStruct
	err := json.Unmarshal(b, &out)
	if err != nil {
		return err
	}

	fmt.Printf("Out:%+v\n", out)
	return nil
}

func TryMarshal(in MyStruct) ([]byte, error) {
	b, err := json.Marshal(in)
	if err != nil {
		return []byte{}, err
	}

	fmt.Printf("b:%+v\n", b)
	return b, nil
}

type CustomTime struct {
	time.Time
}

func (c *CustomTime) UnmarshalJSON(b []byte) error {
	var t string
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}

	if t == "" {
		*c = CustomTime{time.Time{}}
		return nil
	}

	const format = "2006-01-02"
	parsed, err := time.Parse(format, t)
	if err != nil {
		return err
	}

	*c = CustomTime{Time: parsed}
	return nil
}

type PaymentType string

const (
	PaymentTypeCredit PaymentType = "Credit"
	PaymentTypeDebit  PaymentType = "Debit"
)

var paymentTypeNames = map[PaymentType]string{
	PaymentTypeCredit: "Credit",
	PaymentTypeDebit:  "Debit",
}

func (n PaymentType) String() string {
	if name, ok := paymentTypeNames[n]; ok {
		return name
	}
	return "Unknown"
}

func (v PaymentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.String()[:1])
}
