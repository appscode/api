package dtypes_test

import (
	"fmt"
	"testing"

	"github.com/appscode/api/dtypes"
	"github.com/appscode/errors"
	"github.com/stretchr/testify/assert"
)

func TestNewStatusOk(t *testing.T) {
	ok := dtypes.NewStatusOK()
	assert.NotNil(t, ok)

	fmt.Println(ok.Code, ok.Message)

	ok = dtypes.NewStatusOK("with-mesage")
	assert.NotNil(t, ok)

	fmt.Println(ok.Code, ok.Message)
}

func TestNewStatusUnauthorized(t *testing.T) {
	ok := dtypes.NewStatusUnauthorized()
	assert.NotNil(t, ok)

	fmt.Println(ok.Code, ok.Message)
}

func TestNewStatusBadRequest(t *testing.T) {
	ok := dtypes.NewStatusBadRequest()
	assert.NotNil(t, ok)

	fmt.Println(ok.Code, ok.Message)
}

func TestNewStatusFromError(t *testing.T) {
	err := errors.New().WithMessage("no error").WithCause(errors.NewGoError("foo-error")).Internal()
	c := dtypes.NewStatusFromError(err)

	assert.NotNil(t, c)
	fmt.Println(c.Code, c.Message)
}
