// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package cloudassetprojectfeed

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Condition struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Expression: string, required
	Expression terra.StringValue `hcl:"expression,attr" validate:"required"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Title: string, optional
	Title terra.StringValue `hcl:"title,attr"`
}

type FeedOutputConfig struct {
	// PubsubDestination: required
	PubsubDestination *PubsubDestination `hcl:"pubsub_destination,block" validate:"required"`
}

type PubsubDestination struct {
	// Topic: string, required
	Topic terra.StringValue `hcl:"topic,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type ConditionAttributes struct {
	ref terra.Reference
}

func (c ConditionAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c ConditionAttributes) InternalWithRef(ref terra.Reference) ConditionAttributes {
	return ConditionAttributes{ref: ref}
}

func (c ConditionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c ConditionAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("description"))
}

func (c ConditionAttributes) Expression() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("expression"))
}

func (c ConditionAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("location"))
}

func (c ConditionAttributes) Title() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("title"))
}

type FeedOutputConfigAttributes struct {
	ref terra.Reference
}

func (foc FeedOutputConfigAttributes) InternalRef() (terra.Reference, error) {
	return foc.ref, nil
}

func (foc FeedOutputConfigAttributes) InternalWithRef(ref terra.Reference) FeedOutputConfigAttributes {
	return FeedOutputConfigAttributes{ref: ref}
}

func (foc FeedOutputConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return foc.ref.InternalTokens()
}

func (foc FeedOutputConfigAttributes) PubsubDestination() terra.ListValue[PubsubDestinationAttributes] {
	return terra.ReferenceAsList[PubsubDestinationAttributes](foc.ref.Append("pubsub_destination"))
}

type PubsubDestinationAttributes struct {
	ref terra.Reference
}

func (pd PubsubDestinationAttributes) InternalRef() (terra.Reference, error) {
	return pd.ref, nil
}

func (pd PubsubDestinationAttributes) InternalWithRef(ref terra.Reference) PubsubDestinationAttributes {
	return PubsubDestinationAttributes{ref: ref}
}

func (pd PubsubDestinationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pd.ref.InternalTokens()
}

func (pd PubsubDestinationAttributes) Topic() terra.StringValue {
	return terra.ReferenceAsString(pd.ref.Append("topic"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type ConditionState struct {
	Description string `json:"description"`
	Expression  string `json:"expression"`
	Location    string `json:"location"`
	Title       string `json:"title"`
}

type FeedOutputConfigState struct {
	PubsubDestination []PubsubDestinationState `json:"pubsub_destination"`
}

type PubsubDestinationState struct {
	Topic string `json:"topic"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
