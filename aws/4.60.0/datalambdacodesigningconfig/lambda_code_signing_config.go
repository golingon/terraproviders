// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datalambdacodesigningconfig

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type AllowedPublishers struct{}

type Policies struct{}

type AllowedPublishersAttributes struct {
	ref terra.Reference
}

func (ap AllowedPublishersAttributes) InternalRef() terra.Reference {
	return ap.ref
}

func (ap AllowedPublishersAttributes) InternalWithRef(ref terra.Reference) AllowedPublishersAttributes {
	return AllowedPublishersAttributes{ref: ref}
}

func (ap AllowedPublishersAttributes) InternalTokens() hclwrite.Tokens {
	return ap.ref.InternalTokens()
}

func (ap AllowedPublishersAttributes) SigningProfileVersionArns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](ap.ref.Append("signing_profile_version_arns"))
}

type PoliciesAttributes struct {
	ref terra.Reference
}

func (p PoliciesAttributes) InternalRef() terra.Reference {
	return p.ref
}

func (p PoliciesAttributes) InternalWithRef(ref terra.Reference) PoliciesAttributes {
	return PoliciesAttributes{ref: ref}
}

func (p PoliciesAttributes) InternalTokens() hclwrite.Tokens {
	return p.ref.InternalTokens()
}

func (p PoliciesAttributes) UntrustedArtifactOnDeployment() terra.StringValue {
	return terra.ReferenceString(p.ref.Append("untrusted_artifact_on_deployment"))
}

type AllowedPublishersState struct {
	SigningProfileVersionArns []string `json:"signing_profile_version_arns"`
}

type PoliciesState struct {
	UntrustedArtifactOnDeployment string `json:"untrusted_artifact_on_deployment"`
}
