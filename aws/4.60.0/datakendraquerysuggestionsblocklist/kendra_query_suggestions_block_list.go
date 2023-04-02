// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datakendraquerysuggestionsblocklist

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type SourceS3Path struct{}

type SourceS3PathAttributes struct {
	ref terra.Reference
}

func (ssp SourceS3PathAttributes) InternalRef() (terra.Reference, error) {
	return ssp.ref, nil
}

func (ssp SourceS3PathAttributes) InternalWithRef(ref terra.Reference) SourceS3PathAttributes {
	return SourceS3PathAttributes{ref: ref}
}

func (ssp SourceS3PathAttributes) InternalTokens() hclwrite.Tokens {
	return ssp.ref.InternalTokens()
}

func (ssp SourceS3PathAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(ssp.ref.Append("bucket"))
}

func (ssp SourceS3PathAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(ssp.ref.Append("key"))
}

type SourceS3PathState struct {
	Bucket string `json:"bucket"`
	Key    string `json:"key"`
}
