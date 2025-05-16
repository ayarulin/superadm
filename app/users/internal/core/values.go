package core

import (
  "superadmin.ru/pkg/validation"
)

type CompanyName struct {
  value string `validate:"required"`

}

func NewCompanyName(value string) (*CompanyName, error) {
	s := CompanyName{value}

  if err := validation.Struct(s); err != nil {
    return nil, err
  }

	return &s, nil
}

func (s *CompanyName) Value() string {
	return s.value
}
