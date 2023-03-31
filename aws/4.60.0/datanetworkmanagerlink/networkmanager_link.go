// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datanetworkmanagerlink

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Bandwidth struct{}

type BandwidthAttributes struct {
	ref terra.Reference
}

func (b BandwidthAttributes) InternalRef() terra.Reference {
	return b.ref
}

func (b BandwidthAttributes) InternalWithRef(ref terra.Reference) BandwidthAttributes {
	return BandwidthAttributes{ref: ref}
}

func (b BandwidthAttributes) InternalTokens() hclwrite.Tokens {
	return b.ref.InternalTokens()
}

func (b BandwidthAttributes) DownloadSpeed() terra.NumberValue {
	return terra.ReferenceAsNumber(b.ref.Append("download_speed"))
}

func (b BandwidthAttributes) UploadSpeed() terra.NumberValue {
	return terra.ReferenceAsNumber(b.ref.Append("upload_speed"))
}

type BandwidthState struct {
	DownloadSpeed float64 `json:"download_speed"`
	UploadSpeed   float64 `json:"upload_speed"`
}
