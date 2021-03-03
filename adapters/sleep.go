/*
 * Copyright (c) 2020-2021 The SeerLink developers
 */

package adapters

import (
	"time"

	"github.com/SeerLink/seerlink/core/logger"
	"github.com/SeerLink/seerlink/core/store"
	"github.com/SeerLink/seerlink/core/store/models"
	"github.com/SeerLink/seerlink/core/utils"
)

// Sleep adapter allows a job to do nothing for some amount of wall time.
type Sleep struct {
	Until models.AnyTime `json:"until"`
}

// TaskType returns the type of Adapter.
func (adapter *Sleep) TaskType() models.TaskType {
	return TaskTypeSleep
}

// Perform returns the input RunResult after waiting for the specified Until parameter.
func (adapter *Sleep) Perform(input models.RunInput, str *store.Store) models.RunOutput {
	duration := adapter.Duration()
	if duration > 0 {
		logger.Debugw("Task sleeping...", "duration", duration)
		<-str.Clock.After(duration)
	}

	return models.NewRunOutputComplete(models.JSON{})
}

// Duration returns the amount of sleeping this task should be paused for.
func (adapter *Sleep) Duration() time.Duration {
	return utils.DurationFromNow(adapter.Until.Time)
}
