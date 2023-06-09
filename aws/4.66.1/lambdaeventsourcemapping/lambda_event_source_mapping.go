// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package lambdaeventsourcemapping

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type AmazonManagedKafkaEventSourceConfig struct {
	// ConsumerGroupId: string, optional
	ConsumerGroupId terra.StringValue `hcl:"consumer_group_id,attr"`
}

type DestinationConfig struct {
	// OnFailure: optional
	OnFailure *OnFailure `hcl:"on_failure,block"`
}

type OnFailure struct {
	// DestinationArn: string, required
	DestinationArn terra.StringValue `hcl:"destination_arn,attr" validate:"required"`
}

type DocumentDbEventSourceConfig struct {
	// CollectionName: string, optional
	CollectionName terra.StringValue `hcl:"collection_name,attr"`
	// DatabaseName: string, required
	DatabaseName terra.StringValue `hcl:"database_name,attr" validate:"required"`
	// FullDocument: string, optional
	FullDocument terra.StringValue `hcl:"full_document,attr"`
}

type FilterCriteria struct {
	// Filter: min=0,max=5
	Filter []Filter `hcl:"filter,block" validate:"min=0,max=5"`
}

type Filter struct {
	// Pattern: string, optional
	Pattern terra.StringValue `hcl:"pattern,attr"`
}

type ScalingConfig struct {
	// MaximumConcurrency: number, optional
	MaximumConcurrency terra.NumberValue `hcl:"maximum_concurrency,attr"`
}

type SelfManagedEventSource struct {
	// Endpoints: map of string, required
	Endpoints terra.MapValue[terra.StringValue] `hcl:"endpoints,attr" validate:"required"`
}

type SelfManagedKafkaEventSourceConfig struct {
	// ConsumerGroupId: string, optional
	ConsumerGroupId terra.StringValue `hcl:"consumer_group_id,attr"`
}

type SourceAccessConfiguration struct {
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// Uri: string, required
	Uri terra.StringValue `hcl:"uri,attr" validate:"required"`
}

type AmazonManagedKafkaEventSourceConfigAttributes struct {
	ref terra.Reference
}

func (amkesc AmazonManagedKafkaEventSourceConfigAttributes) InternalRef() (terra.Reference, error) {
	return amkesc.ref, nil
}

func (amkesc AmazonManagedKafkaEventSourceConfigAttributes) InternalWithRef(ref terra.Reference) AmazonManagedKafkaEventSourceConfigAttributes {
	return AmazonManagedKafkaEventSourceConfigAttributes{ref: ref}
}

func (amkesc AmazonManagedKafkaEventSourceConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return amkesc.ref.InternalTokens()
}

func (amkesc AmazonManagedKafkaEventSourceConfigAttributes) ConsumerGroupId() terra.StringValue {
	return terra.ReferenceAsString(amkesc.ref.Append("consumer_group_id"))
}

type DestinationConfigAttributes struct {
	ref terra.Reference
}

func (dc DestinationConfigAttributes) InternalRef() (terra.Reference, error) {
	return dc.ref, nil
}

func (dc DestinationConfigAttributes) InternalWithRef(ref terra.Reference) DestinationConfigAttributes {
	return DestinationConfigAttributes{ref: ref}
}

func (dc DestinationConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dc.ref.InternalTokens()
}

func (dc DestinationConfigAttributes) OnFailure() terra.ListValue[OnFailureAttributes] {
	return terra.ReferenceAsList[OnFailureAttributes](dc.ref.Append("on_failure"))
}

type OnFailureAttributes struct {
	ref terra.Reference
}

func (of OnFailureAttributes) InternalRef() (terra.Reference, error) {
	return of.ref, nil
}

func (of OnFailureAttributes) InternalWithRef(ref terra.Reference) OnFailureAttributes {
	return OnFailureAttributes{ref: ref}
}

func (of OnFailureAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return of.ref.InternalTokens()
}

func (of OnFailureAttributes) DestinationArn() terra.StringValue {
	return terra.ReferenceAsString(of.ref.Append("destination_arn"))
}

type DocumentDbEventSourceConfigAttributes struct {
	ref terra.Reference
}

func (ddesc DocumentDbEventSourceConfigAttributes) InternalRef() (terra.Reference, error) {
	return ddesc.ref, nil
}

func (ddesc DocumentDbEventSourceConfigAttributes) InternalWithRef(ref terra.Reference) DocumentDbEventSourceConfigAttributes {
	return DocumentDbEventSourceConfigAttributes{ref: ref}
}

func (ddesc DocumentDbEventSourceConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ddesc.ref.InternalTokens()
}

func (ddesc DocumentDbEventSourceConfigAttributes) CollectionName() terra.StringValue {
	return terra.ReferenceAsString(ddesc.ref.Append("collection_name"))
}

func (ddesc DocumentDbEventSourceConfigAttributes) DatabaseName() terra.StringValue {
	return terra.ReferenceAsString(ddesc.ref.Append("database_name"))
}

func (ddesc DocumentDbEventSourceConfigAttributes) FullDocument() terra.StringValue {
	return terra.ReferenceAsString(ddesc.ref.Append("full_document"))
}

type FilterCriteriaAttributes struct {
	ref terra.Reference
}

func (fc FilterCriteriaAttributes) InternalRef() (terra.Reference, error) {
	return fc.ref, nil
}

func (fc FilterCriteriaAttributes) InternalWithRef(ref terra.Reference) FilterCriteriaAttributes {
	return FilterCriteriaAttributes{ref: ref}
}

func (fc FilterCriteriaAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return fc.ref.InternalTokens()
}

func (fc FilterCriteriaAttributes) Filter() terra.SetValue[FilterAttributes] {
	return terra.ReferenceAsSet[FilterAttributes](fc.ref.Append("filter"))
}

type FilterAttributes struct {
	ref terra.Reference
}

func (f FilterAttributes) InternalRef() (terra.Reference, error) {
	return f.ref, nil
}

func (f FilterAttributes) InternalWithRef(ref terra.Reference) FilterAttributes {
	return FilterAttributes{ref: ref}
}

func (f FilterAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return f.ref.InternalTokens()
}

func (f FilterAttributes) Pattern() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("pattern"))
}

type ScalingConfigAttributes struct {
	ref terra.Reference
}

func (sc ScalingConfigAttributes) InternalRef() (terra.Reference, error) {
	return sc.ref, nil
}

func (sc ScalingConfigAttributes) InternalWithRef(ref terra.Reference) ScalingConfigAttributes {
	return ScalingConfigAttributes{ref: ref}
}

func (sc ScalingConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sc.ref.InternalTokens()
}

func (sc ScalingConfigAttributes) MaximumConcurrency() terra.NumberValue {
	return terra.ReferenceAsNumber(sc.ref.Append("maximum_concurrency"))
}

type SelfManagedEventSourceAttributes struct {
	ref terra.Reference
}

func (smes SelfManagedEventSourceAttributes) InternalRef() (terra.Reference, error) {
	return smes.ref, nil
}

func (smes SelfManagedEventSourceAttributes) InternalWithRef(ref terra.Reference) SelfManagedEventSourceAttributes {
	return SelfManagedEventSourceAttributes{ref: ref}
}

func (smes SelfManagedEventSourceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return smes.ref.InternalTokens()
}

func (smes SelfManagedEventSourceAttributes) Endpoints() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](smes.ref.Append("endpoints"))
}

type SelfManagedKafkaEventSourceConfigAttributes struct {
	ref terra.Reference
}

func (smkesc SelfManagedKafkaEventSourceConfigAttributes) InternalRef() (terra.Reference, error) {
	return smkesc.ref, nil
}

func (smkesc SelfManagedKafkaEventSourceConfigAttributes) InternalWithRef(ref terra.Reference) SelfManagedKafkaEventSourceConfigAttributes {
	return SelfManagedKafkaEventSourceConfigAttributes{ref: ref}
}

func (smkesc SelfManagedKafkaEventSourceConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return smkesc.ref.InternalTokens()
}

func (smkesc SelfManagedKafkaEventSourceConfigAttributes) ConsumerGroupId() terra.StringValue {
	return terra.ReferenceAsString(smkesc.ref.Append("consumer_group_id"))
}

type SourceAccessConfigurationAttributes struct {
	ref terra.Reference
}

func (sac SourceAccessConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return sac.ref, nil
}

func (sac SourceAccessConfigurationAttributes) InternalWithRef(ref terra.Reference) SourceAccessConfigurationAttributes {
	return SourceAccessConfigurationAttributes{ref: ref}
}

func (sac SourceAccessConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sac.ref.InternalTokens()
}

func (sac SourceAccessConfigurationAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(sac.ref.Append("type"))
}

func (sac SourceAccessConfigurationAttributes) Uri() terra.StringValue {
	return terra.ReferenceAsString(sac.ref.Append("uri"))
}

type AmazonManagedKafkaEventSourceConfigState struct {
	ConsumerGroupId string `json:"consumer_group_id"`
}

type DestinationConfigState struct {
	OnFailure []OnFailureState `json:"on_failure"`
}

type OnFailureState struct {
	DestinationArn string `json:"destination_arn"`
}

type DocumentDbEventSourceConfigState struct {
	CollectionName string `json:"collection_name"`
	DatabaseName   string `json:"database_name"`
	FullDocument   string `json:"full_document"`
}

type FilterCriteriaState struct {
	Filter []FilterState `json:"filter"`
}

type FilterState struct {
	Pattern string `json:"pattern"`
}

type ScalingConfigState struct {
	MaximumConcurrency float64 `json:"maximum_concurrency"`
}

type SelfManagedEventSourceState struct {
	Endpoints map[string]string `json:"endpoints"`
}

type SelfManagedKafkaEventSourceConfigState struct {
	ConsumerGroupId string `json:"consumer_group_id"`
}

type SourceAccessConfigurationState struct {
	Type string `json:"type"`
	Uri  string `json:"uri"`
}
