package main

import (
	"database/sql/driver"
	"fmt"
	"strings"
)

type StringSlice []string

// https://golang.org/pkg/database/sql/driver/#Valuer implementation for Postrgres String Array
func (a StringSlice) Value() (driver.Value, error) {
	var quotedStrings []string
	for _, str := range a {
		quotedStrings = append(quotedStrings, fmt.Sprintf(`"%s"`, str))
	}
	value := fmt.Sprintf("{%s}", strings.Join(quotedStrings, ","))
	return value, nil
}

// https://golang.org/pkg/database/sql/#Scanner implementation for Postrgres String Array
func (a *StringSlice) Scan(src interface{}) error {
	val, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("unable to scan")
	}
	value := strings.TrimPrefix(string(val), "{")
	value = strings.TrimSuffix(value, "}")
	*a = strings.Split(value, ",")
	return nil
}
