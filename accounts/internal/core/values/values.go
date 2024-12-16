package values

import "errors"

type OrganizationId string
type OrganizationName string
type UserId string

func NewOrganizationName(value string) (OrganizationName, error) {
	if len(value) < 3 {
		return "", errors.New("min len is 3")
	}

	return OrganizationName(value), nil
}

func NewUserId(value string) (UserId, error) {
	return UserId(value), nil
}
