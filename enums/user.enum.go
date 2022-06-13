package enums

type __UserRole struct {
	ADMIN string
	USER  string
}

var _UserRole = __UserRole{
	ADMIN: "ADMIN",
	USER:  "USER",
}

type __UserStatus struct {
	NOT_ACTIVATED string
	ACTIVE        string
	IS_DISABLED   string
}

var _UserStatus = __UserStatus{
	NOT_ACTIVATED: "NOT_ACTIVATED",
	ACTIVE:        "ACTIVE",
	IS_DISABLED:   "IS_DISABLED",
}

type _User struct {
	Role   __UserRole
	Status __UserStatus
}

var User = _User{
	Role:   _UserRole,
	Status: _UserStatus,
}
