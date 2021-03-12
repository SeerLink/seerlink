// +build !windows

package main

import (
	"io/ioutil"
	"testing"

	"github.com/SeerLink/seerlink/core/cmd"
	"github.com/SeerLink/seerlink/core/internal/cltest"
)

func ExampleRun() {
	t := &testing.T{}
	tc, cleanup := cltest.NewConfig(t)
	defer cleanup()
	tc.Config.Set("SEERLINK_DEV", false)
	testClient := &cmd.Client{
		Renderer:               cmd.RendererTable{Writer: ioutil.Discard},
		Config:                 tc.Config,
		AppFactory:             cmd.SeerlinkAppFactory{},
		KeyStoreAuthenticator:  cmd.TerminalKeyStoreAuthenticator{Prompter: &cltest.MockCountingPrompter{}},
		FallbackAPIInitializer: &cltest.MockAPIInitializer{},
		Runner:                 cmd.SeerlinkRunner{},
		HTTP:                   cltest.NewMockAuthenticatedHTTPClient(tc.Config, "session"),
		ChangePasswordPrompter: cltest.MockChangePasswordPrompter{},
	}

	Run(testClient, "core.test", "--help")
	Run(testClient, "core.test", "--version")
	// Output:
	// NAME:
	//    core.test - CLI for Seerlink

	// USAGE:
	//    core.test [global options] command [command options] [arguments...]

	// VERSION:
	//    unset@unset

	// COMMANDS:
	//    admin           Commands for remotely taking admin related actions
	//    attempts, txas  Commands for managing Ethereum Transaction Attempts
	//    bridges         Commands for Bridges communicating with External Adapters
	//    config          Commands for the node's configuration
	//    jobs            Commands for managing Jobs
	//    keys            Commands for managing various types of keys used by the Seerlink node
	//    node, local     Commands for admin actions that must be run locally
	//    runs            Commands for managing Runs
	//    txs             Commands for handling Ethereum transactions
	//    help, h         Shows a list of commands or help for one command

	// GLOBAL OPTIONS:
	//    --json, -j     json output as opposed to table
	//    --help, -h     show help
	//    --version, -v  print the version
	// core.test version unset@unset
}
