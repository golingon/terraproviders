// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package networkmanagementconnectivitytest

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Destination struct {
	// Instance: string, optional
	Instance terra.StringValue `hcl:"instance,attr"`
	// IpAddress: string, optional
	IpAddress terra.StringValue `hcl:"ip_address,attr"`
	// Network: string, optional
	Network terra.StringValue `hcl:"network,attr"`
	// Port: number, optional
	Port terra.NumberValue `hcl:"port,attr"`
	// ProjectId: string, optional
	ProjectId terra.StringValue `hcl:"project_id,attr"`
}

type Source struct {
	// Instance: string, optional
	Instance terra.StringValue `hcl:"instance,attr"`
	// IpAddress: string, optional
	IpAddress terra.StringValue `hcl:"ip_address,attr"`
	// Network: string, optional
	Network terra.StringValue `hcl:"network,attr"`
	// NetworkType: string, optional
	NetworkType terra.StringValue `hcl:"network_type,attr"`
	// Port: number, optional
	Port terra.NumberValue `hcl:"port,attr"`
	// ProjectId: string, optional
	ProjectId terra.StringValue `hcl:"project_id,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
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

func (d DestinationAttributes) Instance() terra.StringValue {
	return terra.ReferenceString(d.ref.Append("instance"))
}

func (d DestinationAttributes) IpAddress() terra.StringValue {
	return terra.ReferenceString(d.ref.Append("ip_address"))
}

func (d DestinationAttributes) Network() terra.StringValue {
	return terra.ReferenceString(d.ref.Append("network"))
}

func (d DestinationAttributes) Port() terra.NumberValue {
	return terra.ReferenceNumber(d.ref.Append("port"))
}

func (d DestinationAttributes) ProjectId() terra.StringValue {
	return terra.ReferenceString(d.ref.Append("project_id"))
}

type SourceAttributes struct {
	ref terra.Reference
}

func (s SourceAttributes) InternalRef() terra.Reference {
	return s.ref
}

func (s SourceAttributes) InternalWithRef(ref terra.Reference) SourceAttributes {
	return SourceAttributes{ref: ref}
}

func (s SourceAttributes) InternalTokens() hclwrite.Tokens {
	return s.ref.InternalTokens()
}

func (s SourceAttributes) Instance() terra.StringValue {
	return terra.ReferenceString(s.ref.Append("instance"))
}

func (s SourceAttributes) IpAddress() terra.StringValue {
	return terra.ReferenceString(s.ref.Append("ip_address"))
}

func (s SourceAttributes) Network() terra.StringValue {
	return terra.ReferenceString(s.ref.Append("network"))
}

func (s SourceAttributes) NetworkType() terra.StringValue {
	return terra.ReferenceString(s.ref.Append("network_type"))
}

func (s SourceAttributes) Port() terra.NumberValue {
	return terra.ReferenceNumber(s.ref.Append("port"))
}

func (s SourceAttributes) ProjectId() terra.StringValue {
	return terra.ReferenceString(s.ref.Append("project_id"))
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
	return terra.ReferenceString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("update"))
}

type DestinationState struct {
	Instance  string  `json:"instance"`
	IpAddress string  `json:"ip_address"`
	Network   string  `json:"network"`
	Port      float64 `json:"port"`
	ProjectId string  `json:"project_id"`
}

type SourceState struct {
	Instance    string  `json:"instance"`
	IpAddress   string  `json:"ip_address"`
	Network     string  `json:"network"`
	NetworkType string  `json:"network_type"`
	Port        float64 `json:"port"`
	ProjectId   string  `json:"project_id"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
