// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package ecrreplicationconfiguration

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ReplicationConfiguration struct {
	// Rule: min=1,max=10
	Rule []Rule `hcl:"rule,block" validate:"min=1,max=10"`
}

type Rule struct {
	// Destination: min=1,max=25
	Destination []Destination `hcl:"destination,block" validate:"min=1,max=25"`
	// RepositoryFilter: min=0,max=100
	RepositoryFilter []RepositoryFilter `hcl:"repository_filter,block" validate:"min=0,max=100"`
}

type Destination struct {
	// Region: string, required
	Region terra.StringValue `hcl:"region,attr" validate:"required"`
	// RegistryId: string, required
	RegistryId terra.StringValue `hcl:"registry_id,attr" validate:"required"`
}

type RepositoryFilter struct {
	// Filter: string, required
	Filter terra.StringValue `hcl:"filter,attr" validate:"required"`
	// FilterType: string, required
	FilterType terra.StringValue `hcl:"filter_type,attr" validate:"required"`
}

type ReplicationConfigurationAttributes struct {
	ref terra.Reference
}

func (rc ReplicationConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return rc.ref, nil
}

func (rc ReplicationConfigurationAttributes) InternalWithRef(ref terra.Reference) ReplicationConfigurationAttributes {
	return ReplicationConfigurationAttributes{ref: ref}
}

func (rc ReplicationConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rc.ref.InternalTokens()
}

func (rc ReplicationConfigurationAttributes) Rule() terra.ListValue[RuleAttributes] {
	return terra.ReferenceAsList[RuleAttributes](rc.ref.Append("rule"))
}

type RuleAttributes struct {
	ref terra.Reference
}

func (r RuleAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r RuleAttributes) InternalWithRef(ref terra.Reference) RuleAttributes {
	return RuleAttributes{ref: ref}
}

func (r RuleAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r RuleAttributes) Destination() terra.ListValue[DestinationAttributes] {
	return terra.ReferenceAsList[DestinationAttributes](r.ref.Append("destination"))
}

func (r RuleAttributes) RepositoryFilter() terra.ListValue[RepositoryFilterAttributes] {
	return terra.ReferenceAsList[RepositoryFilterAttributes](r.ref.Append("repository_filter"))
}

type DestinationAttributes struct {
	ref terra.Reference
}

func (d DestinationAttributes) InternalRef() (terra.Reference, error) {
	return d.ref, nil
}

func (d DestinationAttributes) InternalWithRef(ref terra.Reference) DestinationAttributes {
	return DestinationAttributes{ref: ref}
}

func (d DestinationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return d.ref.InternalTokens()
}

func (d DestinationAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("region"))
}

func (d DestinationAttributes) RegistryId() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("registry_id"))
}

type RepositoryFilterAttributes struct {
	ref terra.Reference
}

func (rf RepositoryFilterAttributes) InternalRef() (terra.Reference, error) {
	return rf.ref, nil
}

func (rf RepositoryFilterAttributes) InternalWithRef(ref terra.Reference) RepositoryFilterAttributes {
	return RepositoryFilterAttributes{ref: ref}
}

func (rf RepositoryFilterAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rf.ref.InternalTokens()
}

func (rf RepositoryFilterAttributes) Filter() terra.StringValue {
	return terra.ReferenceAsString(rf.ref.Append("filter"))
}

func (rf RepositoryFilterAttributes) FilterType() terra.StringValue {
	return terra.ReferenceAsString(rf.ref.Append("filter_type"))
}

type ReplicationConfigurationState struct {
	Rule []RuleState `json:"rule"`
}

type RuleState struct {
	Destination      []DestinationState      `json:"destination"`
	RepositoryFilter []RepositoryFilterState `json:"repository_filter"`
}

type DestinationState struct {
	Region     string `json:"region"`
	RegistryId string `json:"registry_id"`
}

type RepositoryFilterState struct {
	Filter     string `json:"filter"`
	FilterType string `json:"filter_type"`
}
