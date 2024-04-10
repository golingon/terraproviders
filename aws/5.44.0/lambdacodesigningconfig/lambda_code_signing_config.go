// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package lambdacodesigningconfig

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type AllowedPublishers struct {
	// SigningProfileVersionArns: set of string, required
	SigningProfileVersionArns terra.SetValue[terra.StringValue] `hcl:"signing_profile_version_arns,attr" validate:"required"`
}

type Policies struct {
	// UntrustedArtifactOnDeployment: string, required
	UntrustedArtifactOnDeployment terra.StringValue `hcl:"untrusted_artifact_on_deployment,attr" validate:"required"`
}

type AllowedPublishersAttributes struct {
	ref terra.Reference
}

func (ap AllowedPublishersAttributes) InternalRef() (terra.Reference, error) {
	return ap.ref, nil
}

func (ap AllowedPublishersAttributes) InternalWithRef(ref terra.Reference) AllowedPublishersAttributes {
	return AllowedPublishersAttributes{ref: ref}
}

func (ap AllowedPublishersAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ap.ref.InternalTokens()
}

func (ap AllowedPublishersAttributes) SigningProfileVersionArns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ap.ref.Append("signing_profile_version_arns"))
}

type PoliciesAttributes struct {
	ref terra.Reference
}

func (p PoliciesAttributes) InternalRef() (terra.Reference, error) {
	return p.ref, nil
}

func (p PoliciesAttributes) InternalWithRef(ref terra.Reference) PoliciesAttributes {
	return PoliciesAttributes{ref: ref}
}

func (p PoliciesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return p.ref.InternalTokens()
}

func (p PoliciesAttributes) UntrustedArtifactOnDeployment() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("untrusted_artifact_on_deployment"))
}

type AllowedPublishersState struct {
	SigningProfileVersionArns []string `json:"signing_profile_version_arns"`
}

type PoliciesState struct {
	UntrustedArtifactOnDeployment string `json:"untrusted_artifact_on_deployment"`
}
