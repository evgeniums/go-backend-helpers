package test_utils

import (
	"testing"

	"github.com/evgeniums/go-backend-helpers/pkg/app_context"
	"github.com/evgeniums/go-backend-helpers/pkg/common"
	"github.com/evgeniums/go-backend-helpers/pkg/generic_error"
	"github.com/evgeniums/go-backend-helpers/pkg/op_context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func CheckGenericError(t *testing.T, err error, expectedCode string, expectedMessage ...string) {
	assert.Error(t, err)
	gErr, ok := err.(generic_error.Error)
	require.True(t, ok)
	assert.Equal(t, expectedCode, gErr.Code())
	if len(expectedMessage) != 0 {
		assert.Equal(t, expectedMessage[0], gErr.Message())
	}
}

func ObjectEqual(t *testing.T, left common.Object, right common.Object) {
	if left.GetCreatedAt().Equal(right.GetCreatedAt()) {
		right.SetCreatedAt(left.GetCreatedAt())
	}
	if left.GetUpdatedAt().Equal(right.GetUpdatedAt()) {
		right.SetUpDatedAt(left.GetCreatedAt())
	}
	assert.Equal(t, left, right)
}

func NoError(t *testing.T, ctx op_context.Context, err error) {
	if err != nil && ctx != nil {
		ctx.Close()
	}
	require.NoError(t, err)
}

func NoErrorApp(t *testing.T, app app_context.Context, err error) {
	if err != nil && app != nil {
		app.Logger().CheckFatalStack(app.Logger())
	}
	require.NoError(t, err)
}
