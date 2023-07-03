// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package kendrafaq

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type S3Path struct {
	// Bucket: string, required
	Bucket terra.StringValue `hcl:"bucket,attr" validate:"required"`
	// Key: string, required
	Key terra.StringValue `hcl:"key,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
}

type S3PathAttributes struct {
	ref terra.Reference
}

func (sp S3PathAttributes) InternalRef() (terra.Reference, error) {
	return sp.ref, nil
}

func (sp S3PathAttributes) InternalWithRef(ref terra.Reference) S3PathAttributes {
	return S3PathAttributes{ref: ref}
}

func (sp S3PathAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sp.ref.InternalTokens()
}

func (sp S3PathAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("bucket"))
}

func (sp S3PathAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("key"))
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

type S3PathState struct {
	Bucket string `json:"bucket"`
	Key    string `json:"key"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
}