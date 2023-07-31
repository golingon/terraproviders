// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	ssmmaintenancewindowtarget "github.com/golingon/terraproviders/aws/5.10.0/ssmmaintenancewindowtarget"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSsmMaintenanceWindowTarget creates a new instance of [SsmMaintenanceWindowTarget].
func NewSsmMaintenanceWindowTarget(name string, args SsmMaintenanceWindowTargetArgs) *SsmMaintenanceWindowTarget {
	return &SsmMaintenanceWindowTarget{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SsmMaintenanceWindowTarget)(nil)

// SsmMaintenanceWindowTarget represents the Terraform resource aws_ssm_maintenance_window_target.
type SsmMaintenanceWindowTarget struct {
	Name      string
	Args      SsmMaintenanceWindowTargetArgs
	state     *ssmMaintenanceWindowTargetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SsmMaintenanceWindowTarget].
func (smwt *SsmMaintenanceWindowTarget) Type() string {
	return "aws_ssm_maintenance_window_target"
}

// LocalName returns the local name for [SsmMaintenanceWindowTarget].
func (smwt *SsmMaintenanceWindowTarget) LocalName() string {
	return smwt.Name
}

// Configuration returns the configuration (args) for [SsmMaintenanceWindowTarget].
func (smwt *SsmMaintenanceWindowTarget) Configuration() interface{} {
	return smwt.Args
}

// DependOn is used for other resources to depend on [SsmMaintenanceWindowTarget].
func (smwt *SsmMaintenanceWindowTarget) DependOn() terra.Reference {
	return terra.ReferenceResource(smwt)
}

// Dependencies returns the list of resources [SsmMaintenanceWindowTarget] depends_on.
func (smwt *SsmMaintenanceWindowTarget) Dependencies() terra.Dependencies {
	return smwt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SsmMaintenanceWindowTarget].
func (smwt *SsmMaintenanceWindowTarget) LifecycleManagement() *terra.Lifecycle {
	return smwt.Lifecycle
}

// Attributes returns the attributes for [SsmMaintenanceWindowTarget].
func (smwt *SsmMaintenanceWindowTarget) Attributes() ssmMaintenanceWindowTargetAttributes {
	return ssmMaintenanceWindowTargetAttributes{ref: terra.ReferenceResource(smwt)}
}

// ImportState imports the given attribute values into [SsmMaintenanceWindowTarget]'s state.
func (smwt *SsmMaintenanceWindowTarget) ImportState(av io.Reader) error {
	smwt.state = &ssmMaintenanceWindowTargetState{}
	if err := json.NewDecoder(av).Decode(smwt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", smwt.Type(), smwt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SsmMaintenanceWindowTarget] has state.
func (smwt *SsmMaintenanceWindowTarget) State() (*ssmMaintenanceWindowTargetState, bool) {
	return smwt.state, smwt.state != nil
}

// StateMust returns the state for [SsmMaintenanceWindowTarget]. Panics if the state is nil.
func (smwt *SsmMaintenanceWindowTarget) StateMust() *ssmMaintenanceWindowTargetState {
	if smwt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", smwt.Type(), smwt.LocalName()))
	}
	return smwt.state
}

// SsmMaintenanceWindowTargetArgs contains the configurations for aws_ssm_maintenance_window_target.
type SsmMaintenanceWindowTargetArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// OwnerInformation: string, optional
	OwnerInformation terra.StringValue `hcl:"owner_information,attr"`
	// ResourceType: string, required
	ResourceType terra.StringValue `hcl:"resource_type,attr" validate:"required"`
	// WindowId: string, required
	WindowId terra.StringValue `hcl:"window_id,attr" validate:"required"`
	// Targets: min=1,max=5
	Targets []ssmmaintenancewindowtarget.Targets `hcl:"targets,block" validate:"min=1,max=5"`
}
type ssmMaintenanceWindowTargetAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of aws_ssm_maintenance_window_target.
func (smwt ssmMaintenanceWindowTargetAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(smwt.ref.Append("description"))
}

// Id returns a reference to field id of aws_ssm_maintenance_window_target.
func (smwt ssmMaintenanceWindowTargetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(smwt.ref.Append("id"))
}

// Name returns a reference to field name of aws_ssm_maintenance_window_target.
func (smwt ssmMaintenanceWindowTargetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(smwt.ref.Append("name"))
}

// OwnerInformation returns a reference to field owner_information of aws_ssm_maintenance_window_target.
func (smwt ssmMaintenanceWindowTargetAttributes) OwnerInformation() terra.StringValue {
	return terra.ReferenceAsString(smwt.ref.Append("owner_information"))
}

// ResourceType returns a reference to field resource_type of aws_ssm_maintenance_window_target.
func (smwt ssmMaintenanceWindowTargetAttributes) ResourceType() terra.StringValue {
	return terra.ReferenceAsString(smwt.ref.Append("resource_type"))
}

// WindowId returns a reference to field window_id of aws_ssm_maintenance_window_target.
func (smwt ssmMaintenanceWindowTargetAttributes) WindowId() terra.StringValue {
	return terra.ReferenceAsString(smwt.ref.Append("window_id"))
}

func (smwt ssmMaintenanceWindowTargetAttributes) Targets() terra.ListValue[ssmmaintenancewindowtarget.TargetsAttributes] {
	return terra.ReferenceAsList[ssmmaintenancewindowtarget.TargetsAttributes](smwt.ref.Append("targets"))
}

type ssmMaintenanceWindowTargetState struct {
	Description      string                                    `json:"description"`
	Id               string                                    `json:"id"`
	Name             string                                    `json:"name"`
	OwnerInformation string                                    `json:"owner_information"`
	ResourceType     string                                    `json:"resource_type"`
	WindowId         string                                    `json:"window_id"`
	Targets          []ssmmaintenancewindowtarget.TargetsState `json:"targets"`
}
