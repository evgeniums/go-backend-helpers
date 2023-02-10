package user_service

import (
	"github.com/evgeniums/go-backend-helpers/pkg/api"
	"github.com/evgeniums/go-backend-helpers/pkg/api/api_server"
	"github.com/evgeniums/go-backend-helpers/pkg/user"
	"github.com/evgeniums/go-backend-helpers/pkg/user/user_api"
)

type UserEndpoint[U user.User] struct {
	service *UserService[U]
}

type UserService[U user.User] struct {
	api_server.ServiceBase
	Users        user.Users[U]
	UserTypeName string

	collectionResource api.Resource
	userResource       api.Resource
}

func NewUserService[U user.User](userController user.Users[U],
	setterBuilder func() user.UserFieldsSetter[U],
	userTypeName ...string) *UserService[U] {

	s := &UserService[U]{}

	userType, serviceName, collectionResource, userResource := user_api.PrepareResources(userTypeName...)
	s.collectionResource = collectionResource
	s.userResource = userResource

	s.Init(serviceName)
	s.UserTypeName = userType

	s.Users = userController
	s.AddChild(s.collectionResource)

	s.collectionResource.AddOperation(List(s))
	s.collectionResource.AddOperation(Add(s, setterBuilder))

	s.userResource.AddOperation(Find(s))
	s.userResource.AddChild(SetPhone(s.UserTypeName, s.Users))
	s.userResource.AddChild(SetEmail(s.UserTypeName, s.Users))
	s.userResource.AddChild(SetBlocked(s.UserTypeName, s.Users))
	s.userResource.AddChild(SetPassword(s.UserTypeName, s.Users))

	return s
}

func (s *UserService[U]) CollectionResource() api.Resource {
	return s.collectionResource
}

func (s *UserService[U]) UserResource() api.Resource {
	return s.userResource
}

type SetUserFieldEndpoint struct {
	userTypeName string
	api_server.ResourceEndpoint
	users user.MainFieldSetters
}

func (e *SetUserFieldEndpoint) Init(ep api_server.ResourceEndpointI, userTypeName string, fieldName string, users user.MainFieldSetters, op api.Operation) api_server.ResourceEndpointI {
	api_server.ConstructResourceEndpoint(ep, fieldName, op)
	e.users = users
	e.userTypeName = userTypeName
	return ep
}