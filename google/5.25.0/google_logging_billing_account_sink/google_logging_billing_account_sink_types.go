// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_logging_billing_account_sink

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type BigqueryOptions struct {
	// UsePartitionedTables: bool, required
	UsePartitionedTables terra.BoolValue `hcl:"use_partitioned_tables,attr" validate:"required"`
}

type Exclusions struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Disabled: bool, optional
	Disabled terra.BoolValue `hcl:"disabled,attr"`
	// Filter: string, required
	Filter terra.StringValue `hcl:"filter,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
}

type BigqueryOptionsAttributes struct {
	ref terra.Reference
}

func (bo BigqueryOptionsAttributes) InternalRef() (terra.Reference, error) {
	return bo.ref, nil
}

func (bo BigqueryOptionsAttributes) InternalWithRef(ref terra.Reference) BigqueryOptionsAttributes {
	return BigqueryOptionsAttributes{ref: ref}
}

func (bo BigqueryOptionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return bo.ref.InternalTokens()
}

func (bo BigqueryOptionsAttributes) UsePartitionedTables() terra.BoolValue {
	return terra.ReferenceAsBool(bo.ref.Append("use_partitioned_tables"))
}

type ExclusionsAttributes struct {
	ref terra.Reference
}

func (e ExclusionsAttributes) InternalRef() (terra.Reference, error) {
	return e.ref, nil
}

func (e ExclusionsAttributes) InternalWithRef(ref terra.Reference) ExclusionsAttributes {
	return ExclusionsAttributes{ref: ref}
}

func (e ExclusionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return e.ref.InternalTokens()
}

func (e ExclusionsAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("description"))
}

func (e ExclusionsAttributes) Disabled() terra.BoolValue {
	return terra.ReferenceAsBool(e.ref.Append("disabled"))
}

func (e ExclusionsAttributes) Filter() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("filter"))
}

func (e ExclusionsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("name"))
}

type BigqueryOptionsState struct {
	UsePartitionedTables bool `json:"use_partitioned_tables"`
}

type ExclusionsState struct {
	Description string `json:"description"`
	Disabled    bool   `json:"disabled"`
	Filter      string `json:"filter"`
	Name        string `json:"name"`
}
