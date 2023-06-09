// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	cloudformationstacksetinstance "github.com/golingon/terraproviders/aws/5.6.2/cloudformationstacksetinstance"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudformationStackSetInstance creates a new instance of [CloudformationStackSetInstance].
func NewCloudformationStackSetInstance(name string, args CloudformationStackSetInstanceArgs) *CloudformationStackSetInstance {
	return &CloudformationStackSetInstance{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudformationStackSetInstance)(nil)

// CloudformationStackSetInstance represents the Terraform resource aws_cloudformation_stack_set_instance.
type CloudformationStackSetInstance struct {
	Name      string
	Args      CloudformationStackSetInstanceArgs
	state     *cloudformationStackSetInstanceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CloudformationStackSetInstance].
func (cssi *CloudformationStackSetInstance) Type() string {
	return "aws_cloudformation_stack_set_instance"
}

// LocalName returns the local name for [CloudformationStackSetInstance].
func (cssi *CloudformationStackSetInstance) LocalName() string {
	return cssi.Name
}

// Configuration returns the configuration (args) for [CloudformationStackSetInstance].
func (cssi *CloudformationStackSetInstance) Configuration() interface{} {
	return cssi.Args
}

// DependOn is used for other resources to depend on [CloudformationStackSetInstance].
func (cssi *CloudformationStackSetInstance) DependOn() terra.Reference {
	return terra.ReferenceResource(cssi)
}

// Dependencies returns the list of resources [CloudformationStackSetInstance] depends_on.
func (cssi *CloudformationStackSetInstance) Dependencies() terra.Dependencies {
	return cssi.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CloudformationStackSetInstance].
func (cssi *CloudformationStackSetInstance) LifecycleManagement() *terra.Lifecycle {
	return cssi.Lifecycle
}

// Attributes returns the attributes for [CloudformationStackSetInstance].
func (cssi *CloudformationStackSetInstance) Attributes() cloudformationStackSetInstanceAttributes {
	return cloudformationStackSetInstanceAttributes{ref: terra.ReferenceResource(cssi)}
}

// ImportState imports the given attribute values into [CloudformationStackSetInstance]'s state.
func (cssi *CloudformationStackSetInstance) ImportState(av io.Reader) error {
	cssi.state = &cloudformationStackSetInstanceState{}
	if err := json.NewDecoder(av).Decode(cssi.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cssi.Type(), cssi.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CloudformationStackSetInstance] has state.
func (cssi *CloudformationStackSetInstance) State() (*cloudformationStackSetInstanceState, bool) {
	return cssi.state, cssi.state != nil
}

// StateMust returns the state for [CloudformationStackSetInstance]. Panics if the state is nil.
func (cssi *CloudformationStackSetInstance) StateMust() *cloudformationStackSetInstanceState {
	if cssi.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cssi.Type(), cssi.LocalName()))
	}
	return cssi.state
}

// CloudformationStackSetInstanceArgs contains the configurations for aws_cloudformation_stack_set_instance.
type CloudformationStackSetInstanceArgs struct {
	// AccountId: string, optional
	AccountId terra.StringValue `hcl:"account_id,attr"`
	// CallAs: string, optional
	CallAs terra.StringValue `hcl:"call_as,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ParameterOverrides: map of string, optional
	ParameterOverrides terra.MapValue[terra.StringValue] `hcl:"parameter_overrides,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// RetainStack: bool, optional
	RetainStack terra.BoolValue `hcl:"retain_stack,attr"`
	// StackSetName: string, required
	StackSetName terra.StringValue `hcl:"stack_set_name,attr" validate:"required"`
	// DeploymentTargets: optional
	DeploymentTargets *cloudformationstacksetinstance.DeploymentTargets `hcl:"deployment_targets,block"`
	// OperationPreferences: optional
	OperationPreferences *cloudformationstacksetinstance.OperationPreferences `hcl:"operation_preferences,block"`
	// Timeouts: optional
	Timeouts *cloudformationstacksetinstance.Timeouts `hcl:"timeouts,block"`
}
type cloudformationStackSetInstanceAttributes struct {
	ref terra.Reference
}

// AccountId returns a reference to field account_id of aws_cloudformation_stack_set_instance.
func (cssi cloudformationStackSetInstanceAttributes) AccountId() terra.StringValue {
	return terra.ReferenceAsString(cssi.ref.Append("account_id"))
}

// CallAs returns a reference to field call_as of aws_cloudformation_stack_set_instance.
func (cssi cloudformationStackSetInstanceAttributes) CallAs() terra.StringValue {
	return terra.ReferenceAsString(cssi.ref.Append("call_as"))
}

// Id returns a reference to field id of aws_cloudformation_stack_set_instance.
func (cssi cloudformationStackSetInstanceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cssi.ref.Append("id"))
}

// OrganizationalUnitId returns a reference to field organizational_unit_id of aws_cloudformation_stack_set_instance.
func (cssi cloudformationStackSetInstanceAttributes) OrganizationalUnitId() terra.StringValue {
	return terra.ReferenceAsString(cssi.ref.Append("organizational_unit_id"))
}

// ParameterOverrides returns a reference to field parameter_overrides of aws_cloudformation_stack_set_instance.
func (cssi cloudformationStackSetInstanceAttributes) ParameterOverrides() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cssi.ref.Append("parameter_overrides"))
}

// Region returns a reference to field region of aws_cloudformation_stack_set_instance.
func (cssi cloudformationStackSetInstanceAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(cssi.ref.Append("region"))
}

// RetainStack returns a reference to field retain_stack of aws_cloudformation_stack_set_instance.
func (cssi cloudformationStackSetInstanceAttributes) RetainStack() terra.BoolValue {
	return terra.ReferenceAsBool(cssi.ref.Append("retain_stack"))
}

// StackId returns a reference to field stack_id of aws_cloudformation_stack_set_instance.
func (cssi cloudformationStackSetInstanceAttributes) StackId() terra.StringValue {
	return terra.ReferenceAsString(cssi.ref.Append("stack_id"))
}

// StackSetName returns a reference to field stack_set_name of aws_cloudformation_stack_set_instance.
func (cssi cloudformationStackSetInstanceAttributes) StackSetName() terra.StringValue {
	return terra.ReferenceAsString(cssi.ref.Append("stack_set_name"))
}

func (cssi cloudformationStackSetInstanceAttributes) DeploymentTargets() terra.ListValue[cloudformationstacksetinstance.DeploymentTargetsAttributes] {
	return terra.ReferenceAsList[cloudformationstacksetinstance.DeploymentTargetsAttributes](cssi.ref.Append("deployment_targets"))
}

func (cssi cloudformationStackSetInstanceAttributes) OperationPreferences() terra.ListValue[cloudformationstacksetinstance.OperationPreferencesAttributes] {
	return terra.ReferenceAsList[cloudformationstacksetinstance.OperationPreferencesAttributes](cssi.ref.Append("operation_preferences"))
}

func (cssi cloudformationStackSetInstanceAttributes) Timeouts() cloudformationstacksetinstance.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cloudformationstacksetinstance.TimeoutsAttributes](cssi.ref.Append("timeouts"))
}

type cloudformationStackSetInstanceState struct {
	AccountId            string                                                     `json:"account_id"`
	CallAs               string                                                     `json:"call_as"`
	Id                   string                                                     `json:"id"`
	OrganizationalUnitId string                                                     `json:"organizational_unit_id"`
	ParameterOverrides   map[string]string                                          `json:"parameter_overrides"`
	Region               string                                                     `json:"region"`
	RetainStack          bool                                                       `json:"retain_stack"`
	StackId              string                                                     `json:"stack_id"`
	StackSetName         string                                                     `json:"stack_set_name"`
	DeploymentTargets    []cloudformationstacksetinstance.DeploymentTargetsState    `json:"deployment_targets"`
	OperationPreferences []cloudformationstacksetinstance.OperationPreferencesState `json:"operation_preferences"`
	Timeouts             *cloudformationstacksetinstance.TimeoutsState              `json:"timeouts"`
}
