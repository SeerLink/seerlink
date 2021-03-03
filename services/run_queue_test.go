/*
 * Copyright (c) 2020-2021 The SeerLink developers
 */

package services_test

import (
	"testing"

	"github.com/SeerLink/seerlink/core/internal/cltest"
	"github.com/SeerLink/seerlink/core/internal/mocks"
	"github.com/SeerLink/seerlink/core/services"

	"github.com/onsi/gomega"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/mock"
)

func TestRunQueue(t *testing.T) {
	t.Parallel()
	g := gomega.NewGomegaWithT(t)

	runExecutor := new(mocks.RunExecutor)
	runQueue := services.NewRunQueue(runExecutor)

	executeJobChannel := make(chan struct{})

	runQueue.Start()
	defer runQueue.Stop()

	runExecutor.On("Execute", mock.Anything).
		Return(nil, nil).
		Run(func(mock.Arguments) {
			executeJobChannel <- struct{}{}
		})

	runQueue.Run(uuid.NewV4())

	g.Eventually(func() int {
		return runQueue.WorkerCount()
	}).Should(gomega.Equal(1))

	cltest.CallbackOrTimeout(t, "Execute", func() {
		<-executeJobChannel
	})

	runExecutor.AssertExpectations(t)

	g.Eventually(func() int {
		return runQueue.WorkerCount()
	}).Should(gomega.Equal(0))
}

func TestRunQueue_OneWorkerPerRun(t *testing.T) {
	t.Parallel()
	g := gomega.NewGomegaWithT(t)

	runExecutor := new(mocks.RunExecutor)
	runQueue := services.NewRunQueue(runExecutor)

	executeJobChannel := make(chan struct{})

	runQueue.Start()
	defer runQueue.Stop()

	runExecutor.On("Execute", mock.Anything).
		Return(nil, nil).
		Run(func(mock.Arguments) {
			executeJobChannel <- struct{}{}
		})

	runQueue.Run(uuid.NewV4())
	runQueue.Run(uuid.NewV4())

	g.Eventually(func() int {
		return runQueue.WorkerCount()
	}).Should(gomega.Equal(2))

	cltest.CallbackOrTimeout(t, "Execute", func() {
		<-executeJobChannel
		<-executeJobChannel
	})

	runExecutor.AssertExpectations(t)

	g.Eventually(func() int {
		return runQueue.WorkerCount()
	}).Should(gomega.Equal(0))
}

func TestRunQueue_OneWorkerForSameRunTriggeredMultipleTimes(t *testing.T) {
	t.Parallel()
	g := gomega.NewGomegaWithT(t)

	runExecutor := new(mocks.RunExecutor)
	runQueue := services.NewRunQueue(runExecutor)

	executeJobChannel := make(chan struct{})

	runQueue.Start()
	defer runQueue.Stop()

	runExecutor.On("Execute", mock.Anything).
		Return(nil, nil).
		Run(func(mock.Arguments) {
			executeJobChannel <- struct{}{}
		})

	id := uuid.NewV4()
	runQueue.Run(id)
	runQueue.Run(id)

	g.Eventually(func() int {
		return runQueue.WorkerCount()
	}).Should(gomega.Equal(1))

	g.Consistently(func() int {
		return runQueue.WorkerCount()
	}).Should(gomega.BeNumerically("<", 2))

	cltest.CallbackOrTimeout(t, "Execute", func() {
		<-executeJobChannel
		<-executeJobChannel
	})

	runExecutor.AssertExpectations(t)

	g.Eventually(func() int {
		return runQueue.WorkerCount()
	}).Should(gomega.Equal(0))
}
