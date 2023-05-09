// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package kendraexperience

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Endpoints struct{}

type Configuration struct {
	// ContentSourceConfiguration: optional
	ContentSourceConfiguration *ContentSourceConfiguration `hcl:"content_source_configuration,block"`
	// UserIdentityConfiguration: optional
	UserIdentityConfiguration *UserIdentityConfiguration `hcl:"user_identity_configuration,block"`
}

type ContentSourceConfiguration struct {
	// DataSourceIds: set of string, optional
	DataSourceIds terra.SetValue[terra.StringValue] `hcl:"data_source_ids,attr"`
	// DirectPutContent: bool, optional
	DirectPutContent terra.BoolValue `hcl:"direct_put_content,attr"`
	// FaqIds: set of string, optional
	FaqIds terra.SetValue[terra.StringValue] `hcl:"faq_ids,attr"`
}

type UserIdentityConfiguration struct {
	// IdentityAttributeName: string, required
	IdentityAttributeName terra.StringValue `hcl:"identity_attribute_name,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type EndpointsAttributes struct {
	ref terra.Reference
}

func (e EndpointsAttributes) InternalRef() (terra.Reference, error) {
	return e.ref, nil
}

func (e EndpointsAttributes) InternalWithRef(ref terra.Reference) EndpointsAttributes {
	return EndpointsAttributes{ref: ref}
}

func (e EndpointsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return e.ref.InternalTokens()
}

func (e EndpointsAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("endpoint"))
}

func (e EndpointsAttributes) EndpointType() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("endpoint_type"))
}

type ConfigurationAttributes struct {
	ref terra.Reference
}

func (c ConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c ConfigurationAttributes) InternalWithRef(ref terra.Reference) ConfigurationAttributes {
	return ConfigurationAttributes{ref: ref}
}

func (c ConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c ConfigurationAttributes) ContentSourceConfiguration() terra.ListValue[ContentSourceConfigurationAttributes] {
	return terra.ReferenceAsList[ContentSourceConfigurationAttributes](c.ref.Append("content_source_configuration"))
}

func (c ConfigurationAttributes) UserIdentityConfiguration() terra.ListValue[UserIdentityConfigurationAttributes] {
	return terra.ReferenceAsList[UserIdentityConfigurationAttributes](c.ref.Append("user_identity_configuration"))
}

type ContentSourceConfigurationAttributes struct {
	ref terra.Reference
}

func (csc ContentSourceConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return csc.ref, nil
}

func (csc ContentSourceConfigurationAttributes) InternalWithRef(ref terra.Reference) ContentSourceConfigurationAttributes {
	return ContentSourceConfigurationAttributes{ref: ref}
}

func (csc ContentSourceConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return csc.ref.InternalTokens()
}

func (csc ContentSourceConfigurationAttributes) DataSourceIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](csc.ref.Append("data_source_ids"))
}

func (csc ContentSourceConfigurationAttributes) DirectPutContent() terra.BoolValue {
	return terra.ReferenceAsBool(csc.ref.Append("direct_put_content"))
}

func (csc ContentSourceConfigurationAttributes) FaqIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](csc.ref.Append("faq_ids"))
}

type UserIdentityConfigurationAttributes struct {
	ref terra.Reference
}

func (uic UserIdentityConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return uic.ref, nil
}

func (uic UserIdentityConfigurationAttributes) InternalWithRef(ref terra.Reference) UserIdentityConfigurationAttributes {
	return UserIdentityConfigurationAttributes{ref: ref}
}

func (uic UserIdentityConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return uic.ref.InternalTokens()
}

func (uic UserIdentityConfigurationAttributes) IdentityAttributeName() terra.StringValue {
	return terra.ReferenceAsString(uic.ref.Append("identity_attribute_name"))
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

type EndpointsState struct {
	Endpoint     string `json:"endpoint"`
	EndpointType string `json:"endpoint_type"`
}

type ConfigurationState struct {
	ContentSourceConfiguration []ContentSourceConfigurationState `json:"content_source_configuration"`
	UserIdentityConfiguration  []UserIdentityConfigurationState  `json:"user_identity_configuration"`
}

type ContentSourceConfigurationState struct {
	DataSourceIds    []string `json:"data_source_ids"`
	DirectPutContent bool     `json:"direct_put_content"`
	FaqIds           []string `json:"faq_ids"`
}

type UserIdentityConfigurationState struct {
	IdentityAttributeName string `json:"identity_attribute_name"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
