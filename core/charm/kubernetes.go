// Copyright 2021 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package charm

import (
	"github.com/juju/charm/v9"
	"github.com/juju/collections/set"

	"github.com/juju/juju/core/series"
)

func IsKubernetes(cm charm.CharmMeta) bool {
	switch Format(cm) {
	case FormatV2:
		if len(cm.Meta().Containers) > 0 {
			return true
		}
		// TODO (hml) 2021-04-19
		// Enhance with logic around Assumes once that is finalized.
	case FormatV1:
		if set.NewStrings(cm.Meta().Series...).Contains(series.Kubernetes.String()) {
			return true
		}
	}
	return false
}
