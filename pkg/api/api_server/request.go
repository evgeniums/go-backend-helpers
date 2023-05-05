package api_server

import (
	"github.com/evgeniums/go-backend-helpers/pkg/api"
	"github.com/evgeniums/go-backend-helpers/pkg/app_context"
	"github.com/evgeniums/go-backend-helpers/pkg/auth"
	"github.com/evgeniums/go-backend-helpers/pkg/common"
	"github.com/evgeniums/go-backend-helpers/pkg/db"
	"github.com/evgeniums/go-backend-helpers/pkg/logger"
	"github.com/evgeniums/go-backend-helpers/pkg/op_context/default_op_context"
	"github.com/evgeniums/go-backend-helpers/pkg/validator"
)

// Interface of request to server API.
type Request interface {
	auth.AuthContext
	common.WithParameters

	Server() Server
	Response() Response
	Endpoint() Endpoint

	ParseValidate(cmd interface{}) error
}

type RequestBase struct {
	auth.UserContextBase
	auth.SessionBase
	endpoint Endpoint
}

func (r *RequestBase) Init(app app_context.Context, log logger.Logger, db db.DB, endpoint Endpoint, fields ...logger.Fields) {
	r.ContextBase = default_op_context.NewContext()
	r.ContextBase.Init(app, log, db, fields...)
	r.endpoint = endpoint
}

func (r *RequestBase) Endpoint() Endpoint {
	return r.endpoint
}

func FullRequestPath(r Request) string {
	return r.Endpoint().Resource().BuildActualPath(r.ResourceIds())
}

func FullRequestServicePath(r Request) string {
	return r.Endpoint().Resource().BuildActualPath(r.ResourceIds(), true)
}

func ParseDbQuery(request Request, model interface{}, queryName string, cmd ...api.Query) (*db.Filter, error) {

	var q api.Query
	if len(cmd) == 0 {
		q = &api.DbQuery{}
	} else {
		q = cmd[0]
	}
	c := request.TraceInMethod("ParseDbQuery", logger.Fields{"query_name": queryName})
	defer request.TraceOutMethod()

	err := request.ParseValidate(q)
	if err != nil {
		c.SetMessage("failed to parse/verify query")
		return nil, c.SetError(err)
	}
	if q.Query() == "" {
		return nil, nil
	}

	filter, err := db.ParseQuery(request.Db(), q.Query(), model, queryName, db.EmptyFilterValidator(request.App().Validator()))
	if err != nil {
		vErr, ok := err.(*validator.ValidationError)
		if ok {
			request.SetGenericError(vErr.GenericError(), true)
		}
		c.SetMessage("failed to parse/validate db query")
		return nil, c.SetError(err)
	}

	return filter, nil
}
