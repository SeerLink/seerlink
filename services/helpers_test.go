/*
 * Copyright (c) 2020-2021 The SeerLink developers
 */

package services

func (ht *HeadTracker) ExportedDone() chan struct{} {
	return ht.done
}
