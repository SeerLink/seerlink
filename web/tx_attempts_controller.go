/*
 * Copyright (c) 2020-2021 The SeerLink developers
 */

package web

import (
	"github.com/SeerLink/seerlink/core/services/seerlink"
	"github.com/SeerLink/seerlink/core/store/presenters"

	"github.com/gin-gonic/gin"
)

// TxAttemptsController lists TxAttempts requests.
type TxAttemptsController struct {
	App seerlink.Application
}

// Index returns paginated transaction attempts
func (tac *TxAttemptsController) Index(c *gin.Context, size, page, offset int) {
	attempts, count, err := tac.App.GetStore().EthTxAttempts(offset, size)
	ptxs := make([]presenters.EthTx, len(attempts))
	for i, attempt := range attempts {
		ptxs[i] = presenters.NewEthTxFromAttempt(attempt)
	}
	paginatedResponse(c, "transactions", size, page, ptxs, count, err)
}
