package cmd_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/SeerLink/seerlink/core/cmd"
	"github.com/SeerLink/seerlink/core/internal/cltest"
	"github.com/SeerLink/seerlink/core/internal/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTerminalKeyStoreAuthenticator_WithNoAcctNoPwdCreatesAccount(t *testing.T) {
	t.Parallel()

	store, cleanup := cltest.NewStore(t)
	defer cleanup()
	kst := new(mocks.KeyStoreInterface)
	kst.On("HasAccounts").Return(false)
	kst.On("Unlock", cltest.Password).Return(nil)
	kst.On("NewAccount").Return(accounts.Account{}, nil)
	store.KeyStore = kst

	prompt := &cltest.MockCountingPrompter{
		T: t,
		EnteredStrings: []string{
			cltest.Password,
			"wrongconfirmation",
			cltest.Password,
			cltest.Password,
		},
	}

	auth := cmd.TerminalKeyStoreAuthenticator{Prompter: prompt}
	assert.False(t, store.KeyStore.HasAccounts())
	_, err := auth.Authenticate(store, "")
	assert.NoError(t, err)
	assert.Equal(t, 4, prompt.Count)

	kst.AssertExpectations(t)
}

func TestTerminalKeyStoreAuthenticator_WithNoAcctWithInitialPwdCreatesAcct(t *testing.T) {
	t.Parallel()

	store, cleanup := cltest.NewStore(t)
	kst := new(mocks.KeyStoreInterface)
	kst.On("HasAccounts").Return(false)
	kst.On("Unlock", cltest.Password).Return(nil)
	kst.On("NewAccount").Return(accounts.Account{}, nil)
	kst.On("Accounts").Return([]accounts.Account{})
	store.KeyStore = kst
	defer cleanup()

	auth := cmd.TerminalKeyStoreAuthenticator{Prompter: &cltest.MockCountingPrompter{T: t}}

	assert.Len(t, store.KeyStore.Accounts(), 0)
	_, err := auth.Authenticate(store, cltest.Password)
	assert.NoError(t, err)

	kst.AssertExpectations(t)
}

func TestTerminalKeyStoreAuthenticator_WithAcctNoInitialPwdPromptLoop(t *testing.T) {
	t.Parallel()

	store, cleanup := cltest.NewStore(t)
	defer cleanup()

	cltest.MustAddRandomKeyToKeystore(t, store)

	// prompt loop tries all in array
	prompt := &cltest.MockCountingPrompter{
		T:              t,
		EnteredStrings: []string{"wrongpassword", "wrongagain", cltest.Password},
	}

	auth := cmd.TerminalKeyStoreAuthenticator{Prompter: prompt}
	_, err := auth.Authenticate(store, "")
	assert.NoError(t, err)
	assert.Equal(t, 3, prompt.Count)
}

func TestTerminalKeyStoreAuthenticator_WithAcctAndPwd(t *testing.T) {
	t.Parallel()

	store, cleanup := cltest.NewStore(t)
	defer cleanup()

	cltest.MustAddRandomKeyToKeystore(t, store)

	tests := []struct {
		password  string
		wantError bool
	}{
		{"wrongpassword", true},
	}

	for _, test := range tests {
		t.Run(test.password, func(t *testing.T) {
			auth := cmd.TerminalKeyStoreAuthenticator{Prompter: &cltest.MockCountingPrompter{T: t}}
			_, err := auth.Authenticate(store, test.password)
			assert.Equal(t, test.wantError, err != nil)
		})
	}
}

func TestTerminalKeyStoreAuthenticator_ValidatePasswordStrength(t *testing.T) {
	tests := []struct {
		name               string
		failingPassword    string
		succeedingPassword string
		errorString        string
	}{
		{
			"not long enough",
			"password",
			"passwordpassword",
			"must be longer than 12 characters",
		},
		{
			"not enough lowercase",
			"paSSWORD",
			"password",
			"must contain at least 3 lowercase characters",
		},
		{
			"not enough uppercase",
			"PAssword",
			"PASsword",
			"must contain at least 3 uppercase characters",
		},
		{
			"not enough numbers",
			"password",
			"password123",
			"must contain at least 3 numbers",
		},
		{
			"not enough symbols",
			"password",
			"password!@#",
			"must contain at least 3 symbols",
		},
		{
			"identical consecutive characters",
			"paaaasword",
			"password",
			"must not contain more than 3 identical consecutive characters",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			store, cleanup := cltest.NewStore(t)
			defer cleanup()

			auth := cmd.TerminalKeyStoreAuthenticator{}
			err := auth.ExportedValidatePasswordStrength(store, test.failingPassword)
			require.Error(t, err)
			require.Contains(t, err.Error(), test.errorString)
			err = auth.ExportedValidatePasswordStrength(store, test.succeedingPassword)
			if err != nil {
				require.NotContains(t, err.Error(), test.errorString)
			}
		})
	}
}
