package main

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
)

type StringSlice []string

// https://golang.org/pkg/database/sql/driver/#Valuer implementation for Postrgres String Array
func (stringSlice StringSlice) Value() (driver.Value, error) {
	var quotedStrings []string
	for _, str := range stringSlice {
		quotedStrings = append(quotedStrings, strconv.Quote(str))
	}
	value := fmt.Sprintf("{ %s }", strings.Join(quotedStrings, ","))
	return value, nil
}

// https://golang.org/pkg/database/sql/#Scanner implementation for Postrgres String Array
func (stringSlice *StringSlice) Scan(src interface{}) error {
	val, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("unable to scan")
	}
	value := strings.TrimPrefix(string(val), "{")
	value = strings.TrimSuffix(value, "}")
	*stringSlice = strings.Split(value, ",")
	return nil
}
