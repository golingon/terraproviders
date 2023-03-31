// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datamobilenetworkslice

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type SingleNetworkSliceSelectionAssistanceInformation struct{}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type SingleNetworkSliceSelectionAssistanceInformationAttributes struct {
	ref terra.Reference
}

func (snssai SingleNetworkSliceSelectionAssistanceInformationAttributes) InternalRef() terra.Reference {
	return snssai.ref
}

func (snssai SingleNetworkSliceSelectionAssistanceInformationAttributes) InternalWithRef(ref terra.Reference) SingleNetworkSliceSelectionAssistanceInformationAttributes {
	return SingleNetworkSliceSelectionAssistanceInformationAttributes{ref: ref}
}

func (snssai SingleNetworkSliceSelectionAssistanceInformationAttributes) InternalTokens() hclwrite.Tokens {
	return snssai.ref.InternalTokens()
}

func (snssai SingleNetworkSliceSelectionAssistanceInformationAttributes) SliceDifferentiator() terra.StringValue {
	return terra.ReferenceString(snssai.ref.Append("slice_differentiator"))
}

func (snssai SingleNetworkSliceSelectionAssistanceInformationAttributes) SliceServiceType() terra.NumberValue {
	return terra.ReferenceNumber(snssai.ref.Append("slice_service_type"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() terra.Reference {
	return t.ref
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("read"))
}

type SingleNetworkSliceSelectionAssistanceInformationState struct {
	SliceDifferentiator string  `json:"slice_differentiator"`
	SliceServiceType    float64 `json:"slice_service_type"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}
