// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewDxConnection(name string, args DxConnectionArgs) *DxConnection {
	return &DxConnection{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DxConnection)(nil)

type DxConnection struct {
	Name  string
	Args  DxConnectionArgs
	state *dxConnectionState
}

func (dc *DxConnection) Type() string {
	return "aws_dx_connection"
}

func (dc *DxConnection) LocalName() string {
	return dc.Name
}

func (dc *DxConnection) Configuration() interface{} {
	return dc.Args
}

func (dc *DxConnection) Attributes() dxConnectionAttributes {
	return dxConnectionAttributes{ref: terra.ReferenceResource(dc)}
}

func (dc *DxConnection) ImportState(av io.Reader) error {
	dc.state = &dxConnectionState{}
	if err := json.NewDecoder(av).Decode(dc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dc.Type(), dc.LocalName(), err)
	}
	return nil
}

func (dc *DxConnection) State() (*dxConnectionState, bool) {
	return dc.state, dc.state != nil
}

func (dc *DxConnection) StateMust() *dxConnectionState {
	if dc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dc.Type(), dc.LocalName()))
	}
	return dc.state
}

func (dc *DxConnection) DependOn() terra.Reference {
	return terra.ReferenceResource(dc)
}

type DxConnectionArgs struct {
	// Bandwidth: string, required
	Bandwidth terra.StringValue `hcl:"bandwidth,attr" validate:"required"`
	// EncryptionMode: string, optional
	EncryptionMode terra.StringValue `hcl:"encryption_mode,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ProviderName: string, optional
	ProviderName terra.StringValue `hcl:"provider_name,attr"`
	// RequestMacsec: bool, optional
	RequestMacsec terra.BoolValue `hcl:"request_macsec,attr"`
	// SkipDestroy: bool, optional
	SkipDestroy terra.BoolValue `hcl:"skip_destroy,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// DependsOn contains resources that DxConnection depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type dxConnectionAttributes struct {
	ref terra.Reference
}

func (dc dxConnectionAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(dc.ref.Append("arn"))
}

func (dc dxConnectionAttributes) AwsDevice() terra.StringValue {
	return terra.ReferenceString(dc.ref.Append("aws_device"))
}

func (dc dxConnectionAttributes) Bandwidth() terra.StringValue {
	return terra.ReferenceString(dc.ref.Append("bandwidth"))
}

func (dc dxConnectionAttributes) EncryptionMode() terra.StringValue {
	return terra.ReferenceString(dc.ref.Append("encryption_mode"))
}

func (dc dxConnectionAttributes) HasLogicalRedundancy() terra.StringValue {
	return terra.ReferenceString(dc.ref.Append("has_logical_redundancy"))
}

func (dc dxConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceString(dc.ref.Append("id"))
}

func (dc dxConnectionAttributes) JumboFrameCapable() terra.BoolValue {
	return terra.ReferenceBool(dc.ref.Append("jumbo_frame_capable"))
}

func (dc dxConnectionAttributes) Location() terra.StringValue {
	return terra.ReferenceString(dc.ref.Append("location"))
}

func (dc dxConnectionAttributes) MacsecCapable() terra.BoolValue {
	return terra.ReferenceBool(dc.ref.Append("macsec_capable"))
}

func (dc dxConnectionAttributes) Name() terra.StringValue {
	return terra.ReferenceString(dc.ref.Append("name"))
}

func (dc dxConnectionAttributes) OwnerAccountId() terra.StringValue {
	return terra.ReferenceString(dc.ref.Append("owner_account_id"))
}

func (dc dxConnectionAttributes) PortEncryptionStatus() terra.StringValue {
	return terra.ReferenceString(dc.ref.Append("port_encryption_status"))
}

func (dc dxConnectionAttributes) ProviderName() terra.StringValue {
	return terra.ReferenceString(dc.ref.Append("provider_name"))
}

func (dc dxConnectionAttributes) RequestMacsec() terra.BoolValue {
	return terra.ReferenceBool(dc.ref.Append("request_macsec"))
}

func (dc dxConnectionAttributes) SkipDestroy() terra.BoolValue {
	return terra.ReferenceBool(dc.ref.Append("skip_destroy"))
}

func (dc dxConnectionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](dc.ref.Append("tags"))
}

func (dc dxConnectionAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](dc.ref.Append("tags_all"))
}

func (dc dxConnectionAttributes) VlanId() terra.StringValue {
	return terra.ReferenceString(dc.ref.Append("vlan_id"))
}

type dxConnectionState struct {
	Arn                  string            `json:"arn"`
	AwsDevice            string            `json:"aws_device"`
	Bandwidth            string            `json:"bandwidth"`
	EncryptionMode       string            `json:"encryption_mode"`
	HasLogicalRedundancy string            `json:"has_logical_redundancy"`
	Id                   string            `json:"id"`
	JumboFrameCapable    bool              `json:"jumbo_frame_capable"`
	Location             string            `json:"location"`
	MacsecCapable        bool              `json:"macsec_capable"`
	Name                 string            `json:"name"`
	OwnerAccountId       string            `json:"owner_account_id"`
	PortEncryptionStatus string            `json:"port_encryption_status"`
	ProviderName         string            `json:"provider_name"`
	RequestMacsec        bool              `json:"request_macsec"`
	SkipDestroy          bool              `json:"skip_destroy"`
	Tags                 map[string]string `json:"tags"`
	TagsAll              map[string]string `json:"tags_all"`
	VlanId               string            `json:"vlan_id"`
}
