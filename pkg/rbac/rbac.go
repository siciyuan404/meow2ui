package rbac

type Role string

const (
	RoleOwner  Role = "owner"
	RoleAdmin  Role = "admin"
	RoleMember Role = "member"
	RoleViewer Role = "viewer"
)

func Authorize(role Role, action string) bool {
	switch role {
	case RoleOwner, RoleAdmin:
		return true
	case RoleMember:
		return action != "admin" && action != "export"
	case RoleViewer:
		return action == "read"
	default:
		return false
	}
}
