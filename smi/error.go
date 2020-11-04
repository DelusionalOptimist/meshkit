package smi

import (
	"fmt"

	"github.com/layer5io/meshkit/errors"
)

// ErrSmiInit is the error for smi init method
func ErrSmiInit(des string) error {
	return errors.NewDefault(errors.ErrSmiInit, des)
}

// ErrInstallSmi is the error for installing smi tool
func ErrInstallSmi(err error) error {
	return errors.NewDefault(errors.ErrInstallSmi, fmt.Sprintf("Error installing smi tool: %s", err.Error()))
}

// ErrConnectSmi is the error for connecting to smi tool
func ErrConnectSmi(err error) error {
	return errors.NewDefault(errors.ErrConnectSmi, fmt.Sprintf("Error connecting to smi tool: %s", err.Error()))
}

// ErrRunSmi is the error for running conformance test
func ErrRunSmi(err error) error {
	return errors.NewDefault(errors.ErrRunSmi, fmt.Sprintf("Error running smi tool: %s", err.Error()))
}

// ErrDeleteSmi is the error for deleteing smi tool
func ErrDeleteSmi(err error) error {
	return errors.NewDefault(errors.ErrDeleteSmi, fmt.Sprintf("Error deleting smi tool: %s", err.Error()))
}
