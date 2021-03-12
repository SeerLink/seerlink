package cmd

import (
	"github.com/SeerLink/seerlink/core/store"
)

func (auth TerminalKeyStoreAuthenticator) ExportedValidatePasswordStrength(store *store.Store, password string) error {
	return auth.validatePasswordStrength(store, password)
}
