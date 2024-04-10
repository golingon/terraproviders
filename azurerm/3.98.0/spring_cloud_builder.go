// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	springcloudbuilder "github.com/golingon/terraproviders/azurerm/3.98.0/springcloudbuilder"
	"io"
)

// NewSpringCloudBuilder creates a new instance of [SpringCloudBuilder].
func NewSpringCloudBuilder(name string, args SpringCloudBuilderArgs) *SpringCloudBuilder {
	return &SpringCloudBuilder{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SpringCloudBuilder)(nil)

// SpringCloudBuilder represents the Terraform resource azurerm_spring_cloud_builder.
type SpringCloudBuilder struct {
	Name      string
	Args      SpringCloudBuilderArgs
	state     *springCloudBuilderState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SpringCloudBuilder].
func (scb *SpringCloudBuilder) Type() string {
	return "azurerm_spring_cloud_builder"
}

// LocalName returns the local name for [SpringCloudBuilder].
func (scb *SpringCloudBuilder) LocalName() string {
	return scb.Name
}

// Configuration returns the configuration (args) for [SpringCloudBuilder].
func (scb *SpringCloudBuilder) Configuration() interface{} {
	return scb.Args
}

// DependOn is used for other resources to depend on [SpringCloudBuilder].
func (scb *SpringCloudBuilder) DependOn() terra.Reference {
	return terra.ReferenceResource(scb)
}

// Dependencies returns the list of resources [SpringCloudBuilder] depends_on.
func (scb *SpringCloudBuilder) Dependencies() terra.Dependencies {
	return scb.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SpringCloudBuilder].
func (scb *SpringCloudBuilder) LifecycleManagement() *terra.Lifecycle {
	return scb.Lifecycle
}

// Attributes returns the attributes for [SpringCloudBuilder].
func (scb *SpringCloudBuilder) Attributes() springCloudBuilderAttributes {
	return springCloudBuilderAttributes{ref: terra.ReferenceResource(scb)}
}

// ImportState imports the given attribute values into [SpringCloudBuilder]'s state.
func (scb *SpringCloudBuilder) ImportState(av io.Reader) error {
	scb.state = &springCloudBuilderState{}
	if err := json.NewDecoder(av).Decode(scb.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", scb.Type(), scb.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SpringCloudBuilder] has state.
func (scb *SpringCloudBuilder) State() (*springCloudBuilderState, bool) {
	return scb.state, scb.state != nil
}

// StateMust returns the state for [SpringCloudBuilder]. Panics if the state is nil.
func (scb *SpringCloudBuilder) StateMust() *springCloudBuilderState {
	if scb.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", scb.Type(), scb.LocalName()))
	}
	return scb.state
}

// SpringCloudBuilderArgs contains the configurations for azurerm_spring_cloud_builder.
type SpringCloudBuilderArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SpringCloudServiceId: string, required
	SpringCloudServiceId terra.StringValue `hcl:"spring_cloud_service_id,attr" validate:"required"`
	// BuildPackGroup: min=1
	BuildPackGroup []springcloudbuilder.BuildPackGroup `hcl:"build_pack_group,block" validate:"min=1"`
	// Stack: required
	Stack *springcloudbuilder.Stack `hcl:"stack,block" validate:"required"`
	// Timeouts: optional
	Timeouts *springcloudbuilder.Timeouts `hcl:"timeouts,block"`
}
type springCloudBuilderAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_spring_cloud_builder.
func (scb springCloudBuilderAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(scb.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_spring_cloud_builder.
func (scb springCloudBuilderAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(scb.ref.Append("name"))
}

// SpringCloudServiceId returns a reference to field spring_cloud_service_id of azurerm_spring_cloud_builder.
func (scb springCloudBuilderAttributes) SpringCloudServiceId() terra.StringValue {
	return terra.ReferenceAsString(scb.ref.Append("spring_cloud_service_id"))
}

func (scb springCloudBuilderAttributes) BuildPackGroup() terra.SetValue[springcloudbuilder.BuildPackGroupAttributes] {
	return terra.ReferenceAsSet[springcloudbuilder.BuildPackGroupAttributes](scb.ref.Append("build_pack_group"))
}

func (scb springCloudBuilderAttributes) Stack() terra.ListValue[springcloudbuilder.StackAttributes] {
	return terra.ReferenceAsList[springcloudbuilder.StackAttributes](scb.ref.Append("stack"))
}

func (scb springCloudBuilderAttributes) Timeouts() springcloudbuilder.TimeoutsAttributes {
	return terra.ReferenceAsSingle[springcloudbuilder.TimeoutsAttributes](scb.ref.Append("timeouts"))
}

type springCloudBuilderState struct {
	Id                   string                                   `json:"id"`
	Name                 string                                   `json:"name"`
	SpringCloudServiceId string                                   `json:"spring_cloud_service_id"`
	BuildPackGroup       []springcloudbuilder.BuildPackGroupState `json:"build_pack_group"`
	Stack                []springcloudbuilder.StackState          `json:"stack"`
	Timeouts             *springcloudbuilder.TimeoutsState        `json:"timeouts"`
}
