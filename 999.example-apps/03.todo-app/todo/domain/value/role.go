package value

import "errors"

//go:generate enumer -type=Role -json -transform=upper -trimprefix=Role -output=role_enum.go
type Role int

const (
	ADMIN Role = iota
	REGULAR_USER
)

var ERROR_VALIDATE_ROLE = errors.New("role validation error")

func (r *Role) validate() error {
	if !r.IsARole() {
		return ERROR_VALIDATE_ROLE
	}
}
