package jsonstructure

import (
	"fmt"
	"strings"
)

type SchemaError struct {
	Scope string
	Err   error
}

type EnumError struct {
	SchemaError
}

func (e *SchemaError) Error() string {
	return fmt.Sprintf("At %s: %s", e.Scope, e.Err.Error())
}

func (e *EnumError) Error() string {
	return fmt.Sprintf("At %s: %s", e.Scope, e.Err.Error())
}

func errorAt(err error, scope []string) error {
	if err == nil {
		return nil
	}
	scopeStr := "/" + strings.Join(scope, "/")
	return &SchemaError{
		Scope: scopeStr,
		Err:   err,
	}
}

func enumError(err error, scope []string) error {
	if err == nil {
		return nil
	}
	scopeStr := "/" + strings.Join(scope, "/")
	return &EnumError{
		SchemaError: SchemaError{
			Scope: scopeStr,
			Err:   err,
		},
	}
}
