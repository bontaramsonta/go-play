package main

import "errors"

func validateStatus(status string) error {
	switch len := len(status); {
	case len == 0:
		return errors.New("status cannot be empty")
	case len > 140:
		return errors.New("status exceeds 140 characters")
	default:
		return errors.New("")
	}
}
