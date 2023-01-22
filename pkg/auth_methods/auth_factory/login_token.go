package auth_factory

import (
	"fmt"

	"github.com/evgeniums/go-backend-helpers/pkg/auth"
	"github.com/evgeniums/go-backend-helpers/pkg/auth_methods/auth_login_phash"
	"github.com/evgeniums/go-backend-helpers/pkg/auth_methods/auth_token"
	"github.com/evgeniums/go-backend-helpers/pkg/config"
	"github.com/evgeniums/go-backend-helpers/pkg/config/object_config"
	"github.com/evgeniums/go-backend-helpers/pkg/logger"
	"github.com/evgeniums/go-backend-helpers/pkg/utils"
	"github.com/evgeniums/go-backend-helpers/pkg/validator"
)

const LoginphashTokenProtocol = "login_phash_token"

type LoginphashToken struct {
	auth.AuthSchema

	Login auth.AuthHandler
	Token auth.AuthHandler
}

func (l *LoginphashToken) InitLoginToken(cfg config.Config, log logger.Logger, vld validator.Validator, configPath ...string) error {
	l.AuthSchema.SetAggregation(auth.And)

	path := utils.OptionalArg("auth_manager.methods", configPath...)
	loginCfgPath := object_config.Key(path, auth_login_phash.LoginProtocol)
	tokenCfgPath := object_config.Key(path, auth_token.TokenProtocol)

	login := &auth_login_phash.LoginHandler{}
	err := login.Init(cfg, log, vld, loginCfgPath)
	if err != nil {
		return fmt.Errorf("failed to init login handler: %s", err)
	}
	l.Login = login

	token := &auth_token.AuthTokenHandler{}
	err = token.Init(cfg, log, vld, tokenCfgPath)
	if err != nil {
		return fmt.Errorf("failed to init token handler: %s", err)
	}
	l.Token = token

	return nil
}

func (l *LoginphashToken) Init(cfg config.Config, log logger.Logger, vld validator.Validator, configPath ...string) error {

	l.AuthHandlerBase.Init(LoginphashTokenProtocol)

	err := l.InitLoginToken(cfg, log, vld, configPath...)
	if err != nil {
		return err
	}

	l.AuthSchema.AppendHandlers(l.Login, l.Token)
	return nil
}