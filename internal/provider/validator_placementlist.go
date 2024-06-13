package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/juju/collections/set"
)

// Ensure the implementation satisfies the expected interfaces
var _ basetypes.StringTypable = CustomStringType{}

type CustomStringType struct {
	basetypes.StringType
	// ... potentially other fields ...
}

func (t CustomStringType) Equal(o attr.Type) bool {
	other, ok := o.(CustomStringType)

	if !ok {
		return false
	}

	return t.StringType.Equal(other.StringType)
}

func (t CustomStringType) String() string {
	return "CustomStringType"
}

func (t CustomStringType) ValueFromString(_ context.Context, in basetypes.StringValue) (basetypes.StringValuable, diag.Diagnostics) {
	// CustomStringValue defined in the value type section
	value := CustomStringValue{
		StringValue: in,
	}

	return value, nil
}

func (t CustomStringType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.StringType.ValueFromTerraform(ctx, in)

	if err != nil {
		return nil, err
	}

	stringValue, ok := attrValue.(basetypes.StringValue)

	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	stringValuable, diags := t.ValueFromString(ctx, stringValue)

	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting StringValue to StringValuable: %v", diags)
	}

	return stringValuable, nil
}

func (t CustomStringType) ValueType(_ context.Context) attr.Value {
	// CustomStringValue defined in the value type section
	return CustomStringValue{}
}

// Ensure the implementation satisfies the expected interfaces
var _ basetypes.StringValuable = CustomStringValue{}

type CustomStringValue struct {
	basetypes.StringValue
	// ... potentially other fields ...
}

func (v CustomStringValue) Equal(o attr.Value) bool {
	other, ok := o.(CustomStringValue)

	if !ok {
		return false
	}

	planSet := set.NewStrings()
	for _, v := range strings.Split(v.ValueString(), ",") {
		planSet.Add(strings.TrimSpace(v))
	}

	stateSet := set.NewStrings()
	for _, v := range strings.Split(other.ValueString(), ",") {
		stateSet.Add(strings.TrimSpace(v))
	}

	return planSet.Difference(stateSet).IsEmpty() && stateSet.Difference(planSet).IsEmpty()
}

func (v CustomStringValue) Type(_ context.Context) attr.Type {
	// CustomStringType defined in the schema type section
	return CustomStringType{}
}
