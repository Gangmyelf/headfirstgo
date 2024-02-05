package geo

import (
	"errors"
	"unicode/utf8"
)

type Landmark struct {
	name string
	Coordinates
}

func (l *Landmark) Name() string {
	return l.name
}

func (l *Landmark) SetName(name string) error {
	if utf8.RuneCountInString(name) > 30 || name == "" {
		return errors.New("Invalid name")
	}
	l.name = name
	return nil
}
