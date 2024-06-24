// Copyright 2024 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	jujustorage "github.com/juju/juju/storage"
	"github.com/juju/utils/v3"
)

// storageSetRequiresReplace is a plan modifier function that determines if the storage set requires a replace.
// It compares the storage set in the plan with the storage set in the state.
// Return false if new items were added and old items were not changed.
// Return true if old items were removed
func storageSetRequiresReplace(ctx context.Context, req planmodifier.SetRequest, resp *setplanmodifier.RequiresReplaceIfFuncResponse) {
	planSet := make(map[string]jujustorage.Constraints)
	if !req.PlanValue.IsNull() {
		var planStorageSlice []nestedStorage
		resp.Diagnostics.Append(req.PlanValue.ElementsAs(ctx, &planStorageSlice, false)...)
		if resp.Diagnostics.HasError() {
			return
		}
		if len(planStorageSlice) > 0 {
			for _, storage := range planStorageSlice {
				storageName := storage.Label.ValueString()
				storageSize := storage.Size.ValueString()
				storagePool := storage.Pool.ValueString()
				storageCount := storage.Count.ValueInt64()

				// Validate storage size
				parsedStorageSize, err := utils.ParseSize(storageSize)
				if err != nil {
					resp.Diagnostics.AddError("1Invalid Storage Size", fmt.Sprintf("1Invalid storage size %q: %s", storageSize, err))
					return
				}

				planSet[storageName] = jujustorage.Constraints{
					Size:  parsedStorageSize,
					Pool:  storagePool,
					Count: uint64(storageCount),
				}
			}
		}
	}

	stateSet := make(map[string]jujustorage.Constraints)

	// print the state stateSet
	if !req.StateValue.IsNull() {
		var stateStorageSlice []nestedStorage
		resp.Diagnostics.Append(req.StateValue.ElementsAs(ctx, &stateStorageSlice, false)...)
		if resp.Diagnostics.HasError() {
			return
		}
		if len(stateStorageSlice) > 0 {
			for _, storage := range stateStorageSlice {
				storageName := storage.Label.ValueString()
				storageSize := storage.Size.ValueString()
				storagePool := storage.Pool.ValueString()
				storageCount := storage.Count.ValueInt64()

				// Validate storage size
				parsedStorageSize, err := utils.ParseSize(storageSize)
				if err != nil {
					resp.Diagnostics.AddError("2Invalid Storage Size", fmt.Sprintf("2Invalid storage size %q [%q]: %s", storageSize, stateStorageSlice, err))
					return
				}

				stateSet[storageName] = jujustorage.Constraints{
					Size:  parsedStorageSize,
					Pool:  storagePool,
					Count: uint64(storageCount),
				}
			}
		}
	}

	// Return false if new items were added and old items were not changed
	for key, value := range planSet {
		stateValue, ok := stateSet[key]
		if !ok {
			resp.RequiresReplace = false
			return
		}
		if (value.Size != stateValue.Size) || (value.Pool != stateValue.Pool) || (value.Count != stateValue.Count) {
			resp.RequiresReplace = true
			return
		}
	}

	// Return true if old items were removed
	for key := range stateSet {
		if _, ok := planSet[key]; !ok {
			resp.RequiresReplace = true
			return
		}
	}

	resp.RequiresReplace = false
}
