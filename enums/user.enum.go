package enums

type _UserRole struct {
	ADMIN string
	USER  string
}

var UserRole = _UserRole{
	ADMIN: "ADMIN",
	USER:  "USER",
}

type _UserStatus struct {
	NOT_ACTIVATED string
	ACTIVE        string
	IS_DISABLED   string
}

var UserStatus = _UserStatus{
	NOT_ACTIVATED: "NOT_ACTIVATED",
	ACTIVE:        "ACTIVE",
	IS_DISABLED:   "IS_DISABLED",
}
