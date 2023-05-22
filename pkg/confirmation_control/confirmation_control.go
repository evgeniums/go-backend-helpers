package confirmation_control

import "github.com/evgeniums/go-backend-helpers/pkg/multitenancy"

const StatusSuccess string = "success"

type ConfirmationSender interface {
	SendConfirmation(ctx multitenancy.TenancyContext, operationId string, recipient string, failedUrl string, parameters ...map[string]interface{}) (redirectUrl string, err error)
}

type ConfirmationCallbackHandler interface {
	ConfirmationCallback(ctx multitenancy.TenancyContext, operationId string, codeOrStatus string) (redirectUrl string, err error)
}
