// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDxHostedConnection creates a new instance of [DxHostedConnection].
func NewDxHostedConnection(name string, args DxHostedConnectionArgs) *DxHostedConnection {
	return &DxHostedConnection{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DxHostedConnection)(nil)

// DxHostedConnection represents the Terraform resource aws_dx_hosted_connection.
type DxHostedConnection struct {
	Name      string
	Args      DxHostedConnectionArgs
	state     *dxHostedConnectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DxHostedConnection].
func (dhc *DxHostedConnection) Type() string {
	return "aws_dx_hosted_connection"
}

// LocalName returns the local name for [DxHostedConnection].
func (dhc *DxHostedConnection) LocalName() string {
	return dhc.Name
}

// Configuration returns the configuration (args) for [DxHostedConnection].
func (dhc *DxHostedConnection) Configuration() interface{} {
	return dhc.Args
}

// DependOn is used for other resources to depend on [DxHostedConnection].
func (dhc *DxHostedConnection) DependOn() terra.Reference {
	return terra.ReferenceResource(dhc)
}

// Dependencies returns the list of resources [DxHostedConnection] depends_on.
func (dhc *DxHostedConnection) Dependencies() terra.Dependencies {
	return dhc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DxHostedConnection].
func (dhc *DxHostedConnection) LifecycleManagement() *terra.Lifecycle {
	return dhc.Lifecycle
}

// Attributes returns the attributes for [DxHostedConnection].
func (dhc *DxHostedConnection) Attributes() dxHostedConnectionAttributes {
	return dxHostedConnectionAttributes{ref: terra.ReferenceResource(dhc)}
}

// ImportState imports the given attribute values into [DxHostedConnection]'s state.
func (dhc *DxHostedConnection) ImportState(av io.Reader) error {
	dhc.state = &dxHostedConnectionState{}
	if err := json.NewDecoder(av).Decode(dhc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dhc.Type(), dhc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DxHostedConnection] has state.
func (dhc *DxHostedConnection) State() (*dxHostedConnectionState, bool) {
	return dhc.state, dhc.state != nil
}

// StateMust returns the state for [DxHostedConnection]. Panics if the state is nil.
func (dhc *DxHostedConnection) StateMust() *dxHostedConnectionState {
	if dhc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dhc.Type(), dhc.LocalName()))
	}
	return dhc.state
}

// DxHostedConnectionArgs contains the configurations for aws_dx_hosted_connection.
type DxHostedConnectionArgs struct {
	// Bandwidth: string, required
	Bandwidth terra.StringValue `hcl:"bandwidth,attr" validate:"required"`
	// ConnectionId: string, required
	ConnectionId terra.StringValue `hcl:"connection_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// OwnerAccountId: string, required
	OwnerAccountId terra.StringValue `hcl:"owner_account_id,attr" validate:"required"`
	// Vlan: number, required
	Vlan terra.NumberValue `hcl:"vlan,attr" validate:"required"`
}
type dxHostedConnectionAttributes struct {
	ref terra.Reference
}

// AwsDevice returns a reference to field aws_device of aws_dx_hosted_connection.
func (dhc dxHostedConnectionAttributes) AwsDevice() terra.StringValue {
	return terra.ReferenceAsString(dhc.ref.Append("aws_device"))
}

// Bandwidth returns a reference to field bandwidth of aws_dx_hosted_connection.
func (dhc dxHostedConnectionAttributes) Bandwidth() terra.StringValue {
	return terra.ReferenceAsString(dhc.ref.Append("bandwidth"))
}

// ConnectionId returns a reference to field connection_id of aws_dx_hosted_connection.
func (dhc dxHostedConnectionAttributes) ConnectionId() terra.StringValue {
	return terra.ReferenceAsString(dhc.ref.Append("connection_id"))
}

// HasLogicalRedundancy returns a reference to field has_logical_redundancy of aws_dx_hosted_connection.
func (dhc dxHostedConnectionAttributes) HasLogicalRedundancy() terra.StringValue {
	return terra.ReferenceAsString(dhc.ref.Append("has_logical_redundancy"))
}

// Id returns a reference to field id of aws_dx_hosted_connection.
func (dhc dxHostedConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dhc.ref.Append("id"))
}

// JumboFrameCapable returns a reference to field jumbo_frame_capable of aws_dx_hosted_connection.
func (dhc dxHostedConnectionAttributes) JumboFrameCapable() terra.BoolValue {
	return terra.ReferenceAsBool(dhc.ref.Append("jumbo_frame_capable"))
}

// LagId returns a reference to field lag_id of aws_dx_hosted_connection.
func (dhc dxHostedConnectionAttributes) LagId() terra.StringValue {
	return terra.ReferenceAsString(dhc.ref.Append("lag_id"))
}

// LoaIssueTime returns a reference to field loa_issue_time of aws_dx_hosted_connection.
func (dhc dxHostedConnectionAttributes) LoaIssueTime() terra.StringValue {
	return terra.ReferenceAsString(dhc.ref.Append("loa_issue_time"))
}

// Location returns a reference to field location of aws_dx_hosted_connection.
func (dhc dxHostedConnectionAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dhc.ref.Append("location"))
}

// Name returns a reference to field name of aws_dx_hosted_connection.
func (dhc dxHostedConnectionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dhc.ref.Append("name"))
}

// OwnerAccountId returns a reference to field owner_account_id of aws_dx_hosted_connection.
func (dhc dxHostedConnectionAttributes) OwnerAccountId() terra.StringValue {
	return terra.ReferenceAsString(dhc.ref.Append("owner_account_id"))
}

// PartnerName returns a reference to field partner_name of aws_dx_hosted_connection.
func (dhc dxHostedConnectionAttributes) PartnerName() terra.StringValue {
	return terra.ReferenceAsString(dhc.ref.Append("partner_name"))
}

// ProviderName returns a reference to field provider_name of aws_dx_hosted_connection.
func (dhc dxHostedConnectionAttributes) ProviderName() terra.StringValue {
	return terra.ReferenceAsString(dhc.ref.Append("provider_name"))
}

// Region returns a reference to field region of aws_dx_hosted_connection.
func (dhc dxHostedConnectionAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(dhc.ref.Append("region"))
}

// State returns a reference to field state of aws_dx_hosted_connection.
func (dhc dxHostedConnectionAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(dhc.ref.Append("state"))
}

// Vlan returns a reference to field vlan of aws_dx_hosted_connection.
func (dhc dxHostedConnectionAttributes) Vlan() terra.NumberValue {
	return terra.ReferenceAsNumber(dhc.ref.Append("vlan"))
}

type dxHostedConnectionState struct {
	AwsDevice            string  `json:"aws_device"`
	Bandwidth            string  `json:"bandwidth"`
	ConnectionId         string  `json:"connection_id"`
	HasLogicalRedundancy string  `json:"has_logical_redundancy"`
	Id                   string  `json:"id"`
	JumboFrameCapable    bool    `json:"jumbo_frame_capable"`
	LagId                string  `json:"lag_id"`
	LoaIssueTime         string  `json:"loa_issue_time"`
	Location             string  `json:"location"`
	Name                 string  `json:"name"`
	OwnerAccountId       string  `json:"owner_account_id"`
	PartnerName          string  `json:"partner_name"`
	ProviderName         string  `json:"provider_name"`
	Region               string  `json:"region"`
	State                string  `json:"state"`
	Vlan                 float64 `json:"vlan"`
}