package confirmation_control_server

import (
	"github.com/evgeniums/go-backend-helpers/pkg/background_worker"
	"github.com/evgeniums/go-backend-helpers/pkg/config/object_config"
	"github.com/evgeniums/go-backend-helpers/pkg/multitenancy/app_with_multitenancy"
	"github.com/evgeniums/go-backend-helpers/pkg/op_context"
	"github.com/evgeniums/go-backend-helpers/pkg/utils"
)

type ConfirmationControlServer struct {
	*ExternalServer
	*InternalServer
}

func New(externalServerCfg ...ExternalServerCfg) *ConfirmationControlServer {
	s := &ConfirmationControlServer{}
	s.ExternalServer = NewExternalServer(externalServerCfg...)
	s.InternalServer = NewInternalServer()
	return s
}

func (s *ConfirmationControlServer) Init(app app_with_multitenancy.AppWithMultitenancy, ctx op_context.Context, configPath ...string) error {

	// setup
	c := ctx.TraceInMethod("ConfirmationControlServer.Init")
	var err error
	onExit := func() {
		if err != nil {
			c.SetError(err)
		}
		ctx.TraceOutMethod()
	}
	defer onExit()

	path := utils.OptionalArg("confirmation_control_server", configPath...)
	externalServerConfigPath := object_config.Key(path, "external_server")
	internalServerConfigPath := object_config.Key(path, "internal_server")

	// init external server
	err = s.ExternalServer.Init(app, ctx, externalServerConfigPath)
	if err != nil {
		c.SetMessage("failed to init external server")
		return err
	}

	// init internal server
	err = s.InternalServer.Init(app, ctx, s.ExternalServer.BaseUrl(), internalServerConfigPath)
	if err != nil {
		c.SetMessage("failed to init internal server")
		return err
	}

	// done
	return nil
}

func (s *ConfirmationControlServer) Run(fin background_worker.Finisher) {
	s.ExternalServer.ApiServer().Run(fin)
	s.InternalServer.ApiServer().Run(fin)
}
