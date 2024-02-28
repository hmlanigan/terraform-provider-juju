// Copyright 2024 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package provider

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/juju/collections/set"
)

// A file for defining any needed stringplanmodifier.RequiresReplaceIfFunc implementations.

// orderIndependentRequiresReplace is of type stringplanmodifier.RequiresReplaceIfFunc where
// two string attributes, designated as a comma deliminated list, are compared
// based on contents, not order.
func orderIndependentRequiresReplace(_ context.Context, request planmodifier.StringRequest, response *stringplanmodifier.RequiresReplaceIfFuncResponse) {
	plan := set.NewStrings()
	for _, value := range strings.Split(request.PlanValue.ValueString(), ",") {
		plan.Add(strings.TrimSpace(value))
	}
	state := set.NewStrings()
	for _, value := range strings.Split(request.StateValue.ValueString(), ",") {
		state.Add(strings.TrimSpace(value))
	}
	if !plan.Difference(state).IsEmpty() {
		response.RequiresReplace = true
	}
}

var orderIndependentRequiresReplaceDescription = "If the contents of this attribute change, Terraform will destroy and recreate the resource."
var orderIndependentRequiresReplaceMarkdownDescription = "If the contents of this attribute change, Terraform will destroy and recreate the resource."
