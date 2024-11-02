package utils

import "errors"

type FailingType struct{}

func (f FailingType) MarshalJSON() ([]byte, error) {
	return nil, errors.New("forced marshal failure")
}
