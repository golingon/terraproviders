// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	serverlessapplicationrepositorycloudformationstack "github.com/golingon/terraproviders/aws/5.10.0/serverlessapplicationrepositorycloudformationstack"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewServerlessapplicationrepositoryCloudformationStack creates a new instance of [ServerlessapplicationrepositoryCloudformationStack].
func NewServerlessapplicationrepositoryCloudformationStack(name string, args ServerlessapplicationrepositoryCloudformationStackArgs) *ServerlessapplicationrepositoryCloudformationStack {
	return &ServerlessapplicationrepositoryCloudformationStack{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ServerlessapplicationrepositoryCloudformationStack)(nil)

// ServerlessapplicationrepositoryCloudformationStack represents the Terraform resource aws_serverlessapplicationrepository_cloudformation_stack.
type ServerlessapplicationrepositoryCloudformationStack struct {
	Name      string
	Args      ServerlessapplicationrepositoryCloudformationStackArgs
	state     *serverlessapplicationrepositoryCloudformationStackState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ServerlessapplicationrepositoryCloudformationStack].
func (scs *ServerlessapplicationrepositoryCloudformationStack) Type() string {
	return "aws_serverlessapplicationrepository_cloudformation_stack"
}

// LocalName returns the local name for [ServerlessapplicationrepositoryCloudformationStack].
func (scs *ServerlessapplicationrepositoryCloudformationStack) LocalName() string {
	return scs.Name
}

// Configuration returns the configuration (args) for [ServerlessapplicationrepositoryCloudformationStack].
func (scs *ServerlessapplicationrepositoryCloudformationStack) Configuration() interface{} {
	return scs.Args
}

// DependOn is used for other resources to depend on [ServerlessapplicationrepositoryCloudformationStack].
func (scs *ServerlessapplicationrepositoryCloudformationStack) DependOn() terra.Reference {
	return terra.ReferenceResource(scs)
}

// Dependencies returns the list of resources [ServerlessapplicationrepositoryCloudformationStack] depends_on.
func (scs *ServerlessapplicationrepositoryCloudformationStack) Dependencies() terra.Dependencies {
	return scs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ServerlessapplicationrepositoryCloudformationStack].
func (scs *ServerlessapplicationrepositoryCloudformationStack) LifecycleManagement() *terra.Lifecycle {
	return scs.Lifecycle
}

// Attributes returns the attributes for [ServerlessapplicationrepositoryCloudformationStack].
func (scs *ServerlessapplicationrepositoryCloudformationStack) Attributes() serverlessapplicationrepositoryCloudformationStackAttributes {
	return serverlessapplicationrepositoryCloudformationStackAttributes{ref: terra.ReferenceResource(scs)}
}

// ImportState imports the given attribute values into [ServerlessapplicationrepositoryCloudformationStack]'s state.
func (scs *ServerlessapplicationrepositoryCloudformationStack) ImportState(av io.Reader) error {
	scs.state = &serverlessapplicationrepositoryCloudformationStackState{}
	if err := json.NewDecoder(av).Decode(scs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", scs.Type(), scs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ServerlessapplicationrepositoryCloudformationStack] has state.
func (scs *ServerlessapplicationrepositoryCloudformationStack) State() (*serverlessapplicationrepositoryCloudformationStackState, bool) {
	return scs.state, scs.state != nil
}

// StateMust returns the state for [ServerlessapplicationrepositoryCloudformationStack]. Panics if the state is nil.
func (scs *ServerlessapplicationrepositoryCloudformationStack) StateMust() *serverlessapplicationrepositoryCloudformationStackState {
	if scs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", scs.Type(), scs.LocalName()))
	}
	return scs.state
}

// ServerlessapplicationrepositoryCloudformationStackArgs contains the configurations for aws_serverlessapplicationrepository_cloudformation_stack.
type ServerlessapplicationrepositoryCloudformationStackArgs struct {
	// ApplicationId: string, required
	ApplicationId terra.StringValue `hcl:"application_id,attr" validate:"required"`
	// Capabilities: set of string, required
	Capabilities terra.SetValue[terra.StringValue] `hcl:"capabilities,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Parameters: map of string, optional
	Parameters terra.MapValue[terra.StringValue] `hcl:"parameters,attr"`
	// SemanticVersion: string, optional
	SemanticVersion terra.StringValue `hcl:"semantic_version,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Timeouts: optional
	Timeouts *serverlessapplicationrepositorycloudformationstack.Timeouts `hcl:"timeouts,block"`
}
type serverlessapplicationrepositoryCloudformationStackAttributes struct {
	ref terra.Reference
}

// ApplicationId returns a reference to field application_id of aws_serverlessapplicationrepository_cloudformation_stack.
func (scs serverlessapplicationrepositoryCloudformationStackAttributes) ApplicationId() terra.StringValue {
	return terra.ReferenceAsString(scs.ref.Append("application_id"))
}

// Capabilities returns a reference to field capabilities of aws_serverlessapplicationrepository_cloudformation_stack.
func (scs serverlessapplicationrepositoryCloudformationStackAttributes) Capabilities() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](scs.ref.Append("capabilities"))
}

// Id returns a reference to field id of aws_serverlessapplicationrepository_cloudformation_stack.
func (scs serverlessapplicationrepositoryCloudformationStackAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(scs.ref.Append("id"))
}

// Name returns a reference to field name of aws_serverlessapplicationrepository_cloudformation_stack.
func (scs serverlessapplicationrepositoryCloudformationStackAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(scs.ref.Append("name"))
}

// Outputs returns a reference to field outputs of aws_serverlessapplicationrepository_cloudformation_stack.
func (scs serverlessapplicationrepositoryCloudformationStackAttributes) Outputs() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](scs.ref.Append("outputs"))
}

// Parameters returns a reference to field parameters of aws_serverlessapplicationrepository_cloudformation_stack.
func (scs serverlessapplicationrepositoryCloudformationStackAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](scs.ref.Append("parameters"))
}

// SemanticVersion returns a reference to field semantic_version of aws_serverlessapplicationrepository_cloudformation_stack.
func (scs serverlessapplicationrepositoryCloudformationStackAttributes) SemanticVersion() terra.StringValue {
	return terra.ReferenceAsString(scs.ref.Append("semantic_version"))
}

// Tags returns a reference to field tags of aws_serverlessapplicationrepository_cloudformation_stack.
func (scs serverlessapplicationrepositoryCloudformationStackAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](scs.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_serverlessapplicationrepository_cloudformation_stack.
func (scs serverlessapplicationrepositoryCloudformationStackAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](scs.ref.Append("tags_all"))
}

func (scs serverlessapplicationrepositoryCloudformationStackAttributes) Timeouts() serverlessapplicationrepositorycloudformationstack.TimeoutsAttributes {
	return terra.ReferenceAsSingle[serverlessapplicationrepositorycloudformationstack.TimeoutsAttributes](scs.ref.Append("timeouts"))
}

type serverlessapplicationrepositoryCloudformationStackState struct {
	ApplicationId   string                                                            `json:"application_id"`
	Capabilities    []string                                                          `json:"capabilities"`
	Id              string                                                            `json:"id"`
	Name            string                                                            `json:"name"`
	Outputs         map[string]string                                                 `json:"outputs"`
	Parameters      map[string]string                                                 `json:"parameters"`
	SemanticVersion string                                                            `json:"semantic_version"`
	Tags            map[string]string                                                 `json:"tags"`
	TagsAll         map[string]string                                                 `json:"tags_all"`
	Timeouts        *serverlessapplicationrepositorycloudformationstack.TimeoutsState `json:"timeouts"`
}
