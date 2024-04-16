// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_securityhub_account

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

// Resource represents the Terraform resource aws_securityhub_account.
type Resource struct {
	Name      string
	Args      Args
	state     *awsSecurityhubAccountState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (asa *Resource) Type() string {
	return "aws_securityhub_account"
}

// LocalName returns the local name for [Resource].
func (asa *Resource) LocalName() string {
	return asa.Name
}

// Configuration returns the configuration (args) for [Resource].
func (asa *Resource) Configuration() interface{} {
	return asa.Args
}

// DependOn is used for other resources to depend on [Resource].
func (asa *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(asa)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (asa *Resource) Dependencies() terra.Dependencies {
	return asa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (asa *Resource) LifecycleManagement() *terra.Lifecycle {
	return asa.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (asa *Resource) Attributes() awsSecurityhubAccountAttributes {
	return awsSecurityhubAccountAttributes{ref: terra.ReferenceResource(asa)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (asa *Resource) ImportState(state io.Reader) error {
	asa.state = &awsSecurityhubAccountState{}
	if err := json.NewDecoder(state).Decode(asa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", asa.Type(), asa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (asa *Resource) State() (*awsSecurityhubAccountState, bool) {
	return asa.state, asa.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (asa *Resource) StateMust() *awsSecurityhubAccountState {
	if asa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", asa.Type(), asa.LocalName()))
	}
	return asa.state
}

// Args contains the configurations for aws_securityhub_account.
type Args struct {
	// AutoEnableControls: bool, optional
	AutoEnableControls terra.BoolValue `hcl:"auto_enable_controls,attr"`
	// ControlFindingGenerator: string, optional
	ControlFindingGenerator terra.StringValue `hcl:"control_finding_generator,attr"`
	// EnableDefaultStandards: bool, optional
	EnableDefaultStandards terra.BoolValue `hcl:"enable_default_standards,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}

type awsSecurityhubAccountAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_securityhub_account.
func (asa awsSecurityhubAccountAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(asa.ref.Append("arn"))
}

// AutoEnableControls returns a reference to field auto_enable_controls of aws_securityhub_account.
func (asa awsSecurityhubAccountAttributes) AutoEnableControls() terra.BoolValue {
	return terra.ReferenceAsBool(asa.ref.Append("auto_enable_controls"))
}

// ControlFindingGenerator returns a reference to field control_finding_generator of aws_securityhub_account.
func (asa awsSecurityhubAccountAttributes) ControlFindingGenerator() terra.StringValue {
	return terra.ReferenceAsString(asa.ref.Append("control_finding_generator"))
}

// EnableDefaultStandards returns a reference to field enable_default_standards of aws_securityhub_account.
func (asa awsSecurityhubAccountAttributes) EnableDefaultStandards() terra.BoolValue {
	return terra.ReferenceAsBool(asa.ref.Append("enable_default_standards"))
}

// Id returns a reference to field id of aws_securityhub_account.
func (asa awsSecurityhubAccountAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(asa.ref.Append("id"))
}

type awsSecurityhubAccountState struct {
	Arn                     string `json:"arn"`
	AutoEnableControls      bool   `json:"auto_enable_controls"`
	ControlFindingGenerator string `json:"control_finding_generator"`
	EnableDefaultStandards  bool   `json:"enable_default_standards"`
	Id                      string `json:"id"`
}
