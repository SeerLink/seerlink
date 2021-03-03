/*
 * Copyright (c) 2020-2021 The SeerLink developers
 */

package main

import (
	"os"

	"github.com/SeerLink/seerlink/core/cmd"
	"github.com/SeerLink/seerlink/core/logger"
	"github.com/SeerLink/seerlink/core/store/orm"
)

func main() {
	Run(NewProductionClient(), os.Args...)
}

// Run runs the CLI, providing further command instructions by default.
func Run(client *cmd.Client, args ...string) {
	app := cmd.NewApp(client)
	logger.WarnIf(app.Run(args))
}

// NewProductionClient configures an instance of the CLI to be used
// in production.
func NewProductionClient() *cmd.Client {
	config := orm.NewConfig()
	prompter := cmd.NewTerminalPrompter()
	cookieAuth := cmd.NewSessionCookieAuthenticator(config, cmd.DiskCookieStore{Config: config})
	return &cmd.Client{
		Renderer:                       cmd.RendererTable{Writer: os.Stdout},
		Config:                         config,
		AppFactory:                     cmd.SeerlinkAppFactory{},
		KeyStoreAuthenticator:          cmd.TerminalKeyStoreAuthenticator{Prompter: prompter},
		FallbackAPIInitializer:         cmd.NewPromptingAPIInitializer(prompter),
		Runner:                         cmd.SeerlinkRunner{},
		HTTP:                           cmd.NewAuthenticatedHTTPClient(config, cookieAuth),
		CookieAuthenticator:            cookieAuth,
		FileSessionRequestBuilder:      cmd.NewFileSessionRequestBuilder(),
		PromptingSessionRequestBuilder: cmd.NewPromptingSessionRequestBuilder(prompter),
		ChangePasswordPrompter:         cmd.NewChangePasswordPrompter(),
		PasswordPrompter:               cmd.NewPasswordPrompter(),
	}
}
