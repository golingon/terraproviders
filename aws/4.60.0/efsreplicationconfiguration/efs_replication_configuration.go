// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package efsreplicationconfiguration

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Destination struct {
	// AvailabilityZoneName: string, optional
	AvailabilityZoneName terra.StringValue `hcl:"availability_zone_name,attr"`
	// KmsKeyId: string, optional
	KmsKeyId terra.StringValue `hcl:"kms_key_id,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
}

type DestinationAttributes struct {
	ref terra.Reference
}

func (d DestinationAttributes) InternalRef() terra.Reference {
	return d.ref
}

func (d DestinationAttributes) InternalWithRef(ref terra.Reference) DestinationAttributes {
	return DestinationAttributes{ref: ref}
}

func (d DestinationAttributes) InternalTokens() hclwrite.Tokens {
	return d.ref.InternalTokens()
}

func (d DestinationAttributes) AvailabilityZoneName() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("availability_zone_name"))
}

func (d DestinationAttributes) FileSystemId() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("file_system_id"))
}

func (d DestinationAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("kms_key_id"))
}

func (d DestinationAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("region"))
}

func (d DestinationAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("status"))
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

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("delete"))
}

type DestinationState struct {
	AvailabilityZoneName string `json:"availability_zone_name"`
	FileSystemId         string `json:"file_system_id"`
	KmsKeyId             string `json:"kms_key_id"`
	Region               string `json:"region"`
	Status               string `json:"status"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
}
