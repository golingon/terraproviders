// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package cognitouserpoolclient

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type AnalyticsConfiguration struct {
	// ApplicationArn: string, optional
	ApplicationArn terra.StringValue `hcl:"application_arn,attr"`
	// ApplicationId: string, optional
	ApplicationId terra.StringValue `hcl:"application_id,attr"`
	// ExternalId: string, optional
	ExternalId terra.StringValue `hcl:"external_id,attr"`
	// RoleArn: string, optional
	RoleArn terra.StringValue `hcl:"role_arn,attr"`
	// UserDataShared: bool, optional
	UserDataShared terra.BoolValue `hcl:"user_data_shared,attr"`
}

type TokenValidityUnits struct {
	// AccessToken: string, optional
	AccessToken terra.StringValue `hcl:"access_token,attr"`
	// IdToken: string, optional
	IdToken terra.StringValue `hcl:"id_token,attr"`
	// RefreshToken: string, optional
	RefreshToken terra.StringValue `hcl:"refresh_token,attr"`
}

type AnalyticsConfigurationAttributes struct {
	ref terra.Reference
}

func (ac AnalyticsConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return ac.ref, nil
}

func (ac AnalyticsConfigurationAttributes) InternalWithRef(ref terra.Reference) AnalyticsConfigurationAttributes {
	return AnalyticsConfigurationAttributes{ref: ref}
}

func (ac AnalyticsConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ac.ref.InternalTokens()
}

func (ac AnalyticsConfigurationAttributes) ApplicationArn() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("application_arn"))
}

func (ac AnalyticsConfigurationAttributes) ApplicationId() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("application_id"))
}

func (ac AnalyticsConfigurationAttributes) ExternalId() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("external_id"))
}

func (ac AnalyticsConfigurationAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("role_arn"))
}

func (ac AnalyticsConfigurationAttributes) UserDataShared() terra.BoolValue {
	return terra.ReferenceAsBool(ac.ref.Append("user_data_shared"))
}

type TokenValidityUnitsAttributes struct {
	ref terra.Reference
}

func (tvu TokenValidityUnitsAttributes) InternalRef() (terra.Reference, error) {
	return tvu.ref, nil
}

func (tvu TokenValidityUnitsAttributes) InternalWithRef(ref terra.Reference) TokenValidityUnitsAttributes {
	return TokenValidityUnitsAttributes{ref: ref}
}

func (tvu TokenValidityUnitsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tvu.ref.InternalTokens()
}

func (tvu TokenValidityUnitsAttributes) AccessToken() terra.StringValue {
	return terra.ReferenceAsString(tvu.ref.Append("access_token"))
}

func (tvu TokenValidityUnitsAttributes) IdToken() terra.StringValue {
	return terra.ReferenceAsString(tvu.ref.Append("id_token"))
}

func (tvu TokenValidityUnitsAttributes) RefreshToken() terra.StringValue {
	return terra.ReferenceAsString(tvu.ref.Append("refresh_token"))
}

type AnalyticsConfigurationState struct {
	ApplicationArn string `json:"application_arn"`
	ApplicationId  string `json:"application_id"`
	ExternalId     string `json:"external_id"`
	RoleArn        string `json:"role_arn"`
	UserDataShared bool   `json:"user_data_shared"`
}

type TokenValidityUnitsState struct {
	AccessToken  string `json:"access_token"`
	IdToken      string `json:"id_token"`
	RefreshToken string `json:"refresh_token"`
}
