package api_server

import (
	"github.com/evgeniums/go-backend-helpers/pkg/message"
	"github.com/evgeniums/go-backend-helpers/pkg/op_context"
)

// Interface of request to server API.
type Request interface {
	op_context.WithCtx

	Server() Server
	Service() Service
	Tenancy() Tenancy

	WithAuth
	message.WithMessage

	Response() Response
}
