/*
 * Copyright (c) 2020-2021 The SeerLink developers
 */

package adapters

import (
	"github.com/SeerLink/seerlink/core/store"
	"github.com/SeerLink/seerlink/core/store/models"
)

type ResultCollect struct{}

func (r ResultCollect) TaskType() models.TaskType {
	return TaskTypeResultCollect
}

func (r ResultCollect) Perform(input models.RunInput, store *store.Store) models.RunOutput {
	updatedCollection := make([]interface{}, 0)
	for _, c := range input.ResultCollection().Array() {
		updatedCollection = append(updatedCollection, c.Value())
	}
	updatedCollection = append(updatedCollection, input.Result().Value())
	ro, err := input.Data().Add(models.ResultCollectionKey, updatedCollection)
	if err != nil {
		return models.NewRunOutputError(err)
	}
	return models.NewRunOutputComplete(ro)
}