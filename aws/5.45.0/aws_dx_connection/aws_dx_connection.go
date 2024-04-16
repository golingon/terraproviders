// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_dx_connection

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource aws_dx_connection.
type Resource struct {
	Name      string
	Args      Args
	state     *awsDxConnectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (adc *Resource) Type() string {
	return "aws_dx_connection"
}

// LocalName returns the local name for [Resource].
func (adc *Resource) LocalName() string {
	return adc.Name
}

// Configuration returns the configuration (args) for [Resource].
func (adc *Resource) Configuration() interface{} {
	return adc.Args
}

// DependOn is used for other resources to depend on [Resource].
func (adc *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(adc)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (adc *Resource) Dependencies() terra.Dependencies {
	return adc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (adc *Resource) LifecycleManagement() *terra.Lifecycle {
	return adc.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (adc *Resource) Attributes() awsDxConnectionAttributes {
	return awsDxConnectionAttributes{ref: terra.ReferenceResource(adc)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (adc *Resource) ImportState(state io.Reader) error {
	adc.state = &awsDxConnectionState{}
	if err := json.NewDecoder(state).Decode(adc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", adc.Type(), adc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (adc *Resource) State() (*awsDxConnectionState, bool) {
	return adc.state, adc.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (adc *Resource) StateMust() *awsDxConnectionState {
	if adc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", adc.Type(), adc.LocalName()))
	}
	return adc.state
}

// Args contains the configurations for aws_dx_connection.
type Args struct {
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
}

type awsDxConnectionAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_dx_connection.
func (adc awsDxConnectionAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(adc.ref.Append("arn"))
}

// AwsDevice returns a reference to field aws_device of aws_dx_connection.
func (adc awsDxConnectionAttributes) AwsDevice() terra.StringValue {
	return terra.ReferenceAsString(adc.ref.Append("aws_device"))
}

// Bandwidth returns a reference to field bandwidth of aws_dx_connection.
func (adc awsDxConnectionAttributes) Bandwidth() terra.StringValue {
	return terra.ReferenceAsString(adc.ref.Append("bandwidth"))
}

// EncryptionMode returns a reference to field encryption_mode of aws_dx_connection.
func (adc awsDxConnectionAttributes) EncryptionMode() terra.StringValue {
	return terra.ReferenceAsString(adc.ref.Append("encryption_mode"))
}

// HasLogicalRedundancy returns a reference to field has_logical_redundancy of aws_dx_connection.
func (adc awsDxConnectionAttributes) HasLogicalRedundancy() terra.StringValue {
	return terra.ReferenceAsString(adc.ref.Append("has_logical_redundancy"))
}

// Id returns a reference to field id of aws_dx_connection.
func (adc awsDxConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(adc.ref.Append("id"))
}

// JumboFrameCapable returns a reference to field jumbo_frame_capable of aws_dx_connection.
func (adc awsDxConnectionAttributes) JumboFrameCapable() terra.BoolValue {
	return terra.ReferenceAsBool(adc.ref.Append("jumbo_frame_capable"))
}

// Location returns a reference to field location of aws_dx_connection.
func (adc awsDxConnectionAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(adc.ref.Append("location"))
}

// MacsecCapable returns a reference to field macsec_capable of aws_dx_connection.
func (adc awsDxConnectionAttributes) MacsecCapable() terra.BoolValue {
	return terra.ReferenceAsBool(adc.ref.Append("macsec_capable"))
}

// Name returns a reference to field name of aws_dx_connection.
func (adc awsDxConnectionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(adc.ref.Append("name"))
}

// OwnerAccountId returns a reference to field owner_account_id of aws_dx_connection.
func (adc awsDxConnectionAttributes) OwnerAccountId() terra.StringValue {
	return terra.ReferenceAsString(adc.ref.Append("owner_account_id"))
}

// PartnerName returns a reference to field partner_name of aws_dx_connection.
func (adc awsDxConnectionAttributes) PartnerName() terra.StringValue {
	return terra.ReferenceAsString(adc.ref.Append("partner_name"))
}

// PortEncryptionStatus returns a reference to field port_encryption_status of aws_dx_connection.
func (adc awsDxConnectionAttributes) PortEncryptionStatus() terra.StringValue {
	return terra.ReferenceAsString(adc.ref.Append("port_encryption_status"))
}

// ProviderName returns a reference to field provider_name of aws_dx_connection.
func (adc awsDxConnectionAttributes) ProviderName() terra.StringValue {
	return terra.ReferenceAsString(adc.ref.Append("provider_name"))
}

// RequestMacsec returns a reference to field request_macsec of aws_dx_connection.
func (adc awsDxConnectionAttributes) RequestMacsec() terra.BoolValue {
	return terra.ReferenceAsBool(adc.ref.Append("request_macsec"))
}

// SkipDestroy returns a reference to field skip_destroy of aws_dx_connection.
func (adc awsDxConnectionAttributes) SkipDestroy() terra.BoolValue {
	return terra.ReferenceAsBool(adc.ref.Append("skip_destroy"))
}

// Tags returns a reference to field tags of aws_dx_connection.
func (adc awsDxConnectionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](adc.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_dx_connection.
func (adc awsDxConnectionAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](adc.ref.Append("tags_all"))
}

// VlanId returns a reference to field vlan_id of aws_dx_connection.
func (adc awsDxConnectionAttributes) VlanId() terra.NumberValue {
	return terra.ReferenceAsNumber(adc.ref.Append("vlan_id"))
}

type awsDxConnectionState struct {
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
	PartnerName          string            `json:"partner_name"`
	PortEncryptionStatus string            `json:"port_encryption_status"`
	ProviderName         string            `json:"provider_name"`
	RequestMacsec        bool              `json:"request_macsec"`
	SkipDestroy          bool              `json:"skip_destroy"`
	Tags                 map[string]string `json:"tags"`
	TagsAll              map[string]string `json:"tags_all"`
	VlanId               float64           `json:"vlan_id"`
}
