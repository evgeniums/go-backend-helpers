package admin

import (
	"github.com/evgeniums/go-backend-helpers/pkg/common"
	"github.com/evgeniums/go-backend-helpers/pkg/user/user_session_default"
	"github.com/evgeniums/go-backend-helpers/pkg/user_manager"
)

type Role struct {
	common.IDBase
	Name string
}

type Admin struct {
	user_session_default.User
	Roles []Role `gorm:"-:all"`
}

func NewAdmin() *Admin {
	// TODO fill roles from database
	a := &Admin{}
	a.Roles = make([]Role, 1)
	a.Roles[0] = Role{Name: "superadmin"}
	return a
}

type AdminSession struct {
	user_manager.SessionBase
}

func NewAdminSession() *AdminSession {
	return &AdminSession{}
}

type AdminSessionClient struct {
	user_session_default.UserSessionClient
}

func NewAdminSessionClient() *AdminSessionClient {
	return &AdminSessionClient{}
}
