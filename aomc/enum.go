package aomc

// AuthImplStatus is a string indicating the implementation status of an auth method
// in a hook class belonging to a class.
// It's restricted to a few specific values (see the constants).
type AuthImplStatus string

// AuthImplStatus values
const (
	Unknown AuthImplStatus = "UNKNOWN"
	Yes     AuthImplStatus = "YES"
	No      AuthImplStatus = "NO"
)

func (s *AuthImplStatus) String() string {
	return string(*s)
}

// UserRole is a string containing the role that's required to C/R/U/D an object of a class.
// It's used for example in a class' "RequiredRoleCreate".
// When a UserRole is set to "AppAdmin", the "AllowedRolesCreate" kick in.
// This doesn't make much sense, but it's defined that way by ApiOmat.
// The string is restricted to a few specific values (see the constants).
type UserRole string

// UserRole values
const (
	Everyone   UserRole = "Guest"
	User       UserRole = "User"
	Owner      UserRole = "Owner"
	AppAdmin   UserRole = "AppAdmin"
	OrgAdmin   UserRole = "OrgAdmin"
	SuperAdmin UserRole = "SuperAdmin"
)

func (s *UserRole) String() string {
	return string(*s)
}
