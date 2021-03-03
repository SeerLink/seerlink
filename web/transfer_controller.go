/*
 * Copyright (c) 2020-2021 The SeerLink developers
 */

package web

import (
	"fmt"
	"net/http"

	"github.com/SeerLink/seerlink/core/services/bulletprooftxmanager"
	"github.com/SeerLink/seerlink/core/services/seerlink"
	"github.com/SeerLink/seerlink/core/store/models"

	"github.com/gin-gonic/gin"
)

// TransfersController can send SEER tokens to another address
type TransfersController struct {
	App seerlink.Application
}

// Create sends ETH from the Seerlink's account to a specified address.
//
// Example: "<application>/withdrawals"
func (tc *TransfersController) Create(c *gin.Context) {
	var tr models.SendEtherRequest
	if err := c.ShouldBindJSON(&tr); err != nil {
		jsonAPIError(c, http.StatusBadRequest, err)
		return
	}

	store := tc.App.GetStore()

	etx, err := bulletprooftxmanager.SendEther(store, tr.FromAddress, tr.DestinationAddress, tr.Amount)
	if err != nil {
		jsonAPIError(c, http.StatusBadRequest, fmt.Errorf("transaction failed: %v", err))
		return
	}

	jsonAPIResponse(c, etx, "eth_tx")
}
