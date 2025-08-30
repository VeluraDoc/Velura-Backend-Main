package model

type Role int

const (
	Admin Role = iota
	Owner
	Member
)

func (r Role) String() string {
	switch r {
	case Admin:
		return "Admin"
	case Owner:
		return "Owner"
	case Member:
		return "Member"
	default:
		return "Unknown"
	}
}
