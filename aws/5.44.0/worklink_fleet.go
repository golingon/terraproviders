// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	worklinkfleet "github.com/golingon/terraproviders/aws/5.44.0/worklinkfleet"
	"io"
)

// NewWorklinkFleet creates a new instance of [WorklinkFleet].
func NewWorklinkFleet(name string, args WorklinkFleetArgs) *WorklinkFleet {
	return &WorklinkFleet{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*WorklinkFleet)(nil)

// WorklinkFleet represents the Terraform resource aws_worklink_fleet.
type WorklinkFleet struct {
	Name      string
	Args      WorklinkFleetArgs
	state     *worklinkFleetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [WorklinkFleet].
func (wf *WorklinkFleet) Type() string {
	return "aws_worklink_fleet"
}

// LocalName returns the local name for [WorklinkFleet].
func (wf *WorklinkFleet) LocalName() string {
	return wf.Name
}

// Configuration returns the configuration (args) for [WorklinkFleet].
func (wf *WorklinkFleet) Configuration() interface{} {
	return wf.Args
}

// DependOn is used for other resources to depend on [WorklinkFleet].
func (wf *WorklinkFleet) DependOn() terra.Reference {
	return terra.ReferenceResource(wf)
}

// Dependencies returns the list of resources [WorklinkFleet] depends_on.
func (wf *WorklinkFleet) Dependencies() terra.Dependencies {
	return wf.DependsOn
}

// LifecycleManagement returns the lifecycle block for [WorklinkFleet].
func (wf *WorklinkFleet) LifecycleManagement() *terra.Lifecycle {
	return wf.Lifecycle
}

// Attributes returns the attributes for [WorklinkFleet].
func (wf *WorklinkFleet) Attributes() worklinkFleetAttributes {
	return worklinkFleetAttributes{ref: terra.ReferenceResource(wf)}
}

// ImportState imports the given attribute values into [WorklinkFleet]'s state.
func (wf *WorklinkFleet) ImportState(av io.Reader) error {
	wf.state = &worklinkFleetState{}
	if err := json.NewDecoder(av).Decode(wf.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", wf.Type(), wf.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [WorklinkFleet] has state.
func (wf *WorklinkFleet) State() (*worklinkFleetState, bool) {
	return wf.state, wf.state != nil
}

// StateMust returns the state for [WorklinkFleet]. Panics if the state is nil.
func (wf *WorklinkFleet) StateMust() *worklinkFleetState {
	if wf.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", wf.Type(), wf.LocalName()))
	}
	return wf.state
}

// WorklinkFleetArgs contains the configurations for aws_worklink_fleet.
type WorklinkFleetArgs struct {
	// AuditStreamArn: string, optional
	AuditStreamArn terra.StringValue `hcl:"audit_stream_arn,attr"`
	// DeviceCaCertificate: string, optional
	DeviceCaCertificate terra.StringValue `hcl:"device_ca_certificate,attr"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// OptimizeForEndUserLocation: bool, optional
	OptimizeForEndUserLocation terra.BoolValue `hcl:"optimize_for_end_user_location,attr"`
	// IdentityProvider: optional
	IdentityProvider *worklinkfleet.IdentityProvider `hcl:"identity_provider,block"`
	// Network: optional
	Network *worklinkfleet.Network `hcl:"network,block"`
}
type worklinkFleetAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_worklink_fleet.
func (wf worklinkFleetAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(wf.ref.Append("arn"))
}

// AuditStreamArn returns a reference to field audit_stream_arn of aws_worklink_fleet.
func (wf worklinkFleetAttributes) AuditStreamArn() terra.StringValue {
	return terra.ReferenceAsString(wf.ref.Append("audit_stream_arn"))
}

// CompanyCode returns a reference to field company_code of aws_worklink_fleet.
func (wf worklinkFleetAttributes) CompanyCode() terra.StringValue {
	return terra.ReferenceAsString(wf.ref.Append("company_code"))
}

// CreatedTime returns a reference to field created_time of aws_worklink_fleet.
func (wf worklinkFleetAttributes) CreatedTime() terra.StringValue {
	return terra.ReferenceAsString(wf.ref.Append("created_time"))
}

// DeviceCaCertificate returns a reference to field device_ca_certificate of aws_worklink_fleet.
func (wf worklinkFleetAttributes) DeviceCaCertificate() terra.StringValue {
	return terra.ReferenceAsString(wf.ref.Append("device_ca_certificate"))
}

// DisplayName returns a reference to field display_name of aws_worklink_fleet.
func (wf worklinkFleetAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(wf.ref.Append("display_name"))
}

// Id returns a reference to field id of aws_worklink_fleet.
func (wf worklinkFleetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(wf.ref.Append("id"))
}

// LastUpdatedTime returns a reference to field last_updated_time of aws_worklink_fleet.
func (wf worklinkFleetAttributes) LastUpdatedTime() terra.StringValue {
	return terra.ReferenceAsString(wf.ref.Append("last_updated_time"))
}

// Name returns a reference to field name of aws_worklink_fleet.
func (wf worklinkFleetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(wf.ref.Append("name"))
}

// OptimizeForEndUserLocation returns a reference to field optimize_for_end_user_location of aws_worklink_fleet.
func (wf worklinkFleetAttributes) OptimizeForEndUserLocation() terra.BoolValue {
	return terra.ReferenceAsBool(wf.ref.Append("optimize_for_end_user_location"))
}

func (wf worklinkFleetAttributes) IdentityProvider() terra.ListValue[worklinkfleet.IdentityProviderAttributes] {
	return terra.ReferenceAsList[worklinkfleet.IdentityProviderAttributes](wf.ref.Append("identity_provider"))
}

func (wf worklinkFleetAttributes) Network() terra.ListValue[worklinkfleet.NetworkAttributes] {
	return terra.ReferenceAsList[worklinkfleet.NetworkAttributes](wf.ref.Append("network"))
}

type worklinkFleetState struct {
	Arn                        string                                `json:"arn"`
	AuditStreamArn             string                                `json:"audit_stream_arn"`
	CompanyCode                string                                `json:"company_code"`
	CreatedTime                string                                `json:"created_time"`
	DeviceCaCertificate        string                                `json:"device_ca_certificate"`
	DisplayName                string                                `json:"display_name"`
	Id                         string                                `json:"id"`
	LastUpdatedTime            string                                `json:"last_updated_time"`
	Name                       string                                `json:"name"`
	OptimizeForEndUserLocation bool                                  `json:"optimize_for_end_user_location"`
	IdentityProvider           []worklinkfleet.IdentityProviderState `json:"identity_provider"`
	Network                    []worklinkfleet.NetworkState          `json:"network"`
}
