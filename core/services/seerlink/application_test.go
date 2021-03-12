// +build !windows

package seerlink_test

import (
	"syscall"
	"testing"

	"github.com/SeerLink/seerlink/core/services/eth"

	"github.com/SeerLink/seerlink/core/internal/cltest"
	"github.com/SeerLink/seerlink/core/store/models"

	"github.com/onsi/gomega"
	"github.com/stretchr/testify/require"
	"github.com/tevino/abool"
)

func TestSeerlinkApplication_SignalShutdown(t *testing.T) {

	rpcClient, gethClient, _, assertMocksCalled := cltest.NewEthMocksWithStartupAssertions(t)
	defer assertMocksCalled()
	app, cleanup := cltest.NewApplication(t,
		eth.NewClientWith(rpcClient, gethClient),
	)
	defer cleanup()
	completed := abool.New()
	app.Exiter = func(code int) {
		completed.Set()
	}

	require.NoError(t, app.Start())
	syscall.Kill(syscall.Getpid(), syscall.SIGTERM)

	gomega.NewGomegaWithT(t).Eventually(func() bool {
		return completed.IsSet()
	}).Should(gomega.BeTrue())
}

func TestSeerlinkApplication_resumesPendingConnection_Happy(t *testing.T) {
	rpcClient, gethClient, _, assertMocksCalled := cltest.NewEthMocksWithStartupAssertions(t)
	defer assertMocksCalled()
	app, cleanup := cltest.NewApplication(t,
		eth.NewClientWith(rpcClient, gethClient),
	)
	defer cleanup()
	store := app.Store

	j := cltest.NewJobWithWebInitiator()
	require.NoError(t, store.CreateJob(&j))

	jr := cltest.CreateJobRunWithStatus(t, store, j, models.RunStatusPendingConnection)

	require.NoError(t, app.StartAndConnect())
	_ = cltest.WaitForJobRunToComplete(t, store, jr)
}

func TestSeerlinkApplication_resumesPendingConnection_Archived(t *testing.T) {
	rpcClient, gethClient, _, assertMocksCalled := cltest.NewEthMocksWithStartupAssertions(t)
	defer assertMocksCalled()
	app, cleanup := cltest.NewApplication(t,
		eth.NewClientWith(rpcClient, gethClient),
	)
	defer cleanup()
	store := app.Store

	j := cltest.NewJobWithWebInitiator()
	require.NoError(t, store.CreateJob(&j))

	jr := cltest.CreateJobRunWithStatus(t, store, j, models.RunStatusPendingConnection)

	require.NoError(t, store.ArchiveJob(j.ID))

	require.NoError(t, app.StartAndConnect())
	_ = cltest.WaitForJobRunToComplete(t, store, jr)
}
