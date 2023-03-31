// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datakendrafaq

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type S3Path struct{}

type S3PathAttributes struct {
	ref terra.Reference
}

func (sp S3PathAttributes) InternalRef() terra.Reference {
	return sp.ref
}

func (sp S3PathAttributes) InternalWithRef(ref terra.Reference) S3PathAttributes {
	return S3PathAttributes{ref: ref}
}

func (sp S3PathAttributes) InternalTokens() hclwrite.Tokens {
	return sp.ref.InternalTokens()
}

func (sp S3PathAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("bucket"))
}

func (sp S3PathAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("key"))
}

type S3PathState struct {
	Bucket string `json:"bucket"`
	Key    string `json:"key"`
}
