// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	ecscapacityprovider "github.com/golingon/terraproviders/aws/4.60.0/ecscapacityprovider"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEcsCapacityProvider creates a new instance of [EcsCapacityProvider].
func NewEcsCapacityProvider(name string, args EcsCapacityProviderArgs) *EcsCapacityProvider {
	return &EcsCapacityProvider{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EcsCapacityProvider)(nil)

// EcsCapacityProvider represents the Terraform resource aws_ecs_capacity_provider.
type EcsCapacityProvider struct {
	Name      string
	Args      EcsCapacityProviderArgs
	state     *ecsCapacityProviderState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [EcsCapacityProvider].
func (ecp *EcsCapacityProvider) Type() string {
	return "aws_ecs_capacity_provider"
}

// LocalName returns the local name for [EcsCapacityProvider].
func (ecp *EcsCapacityProvider) LocalName() string {
	return ecp.Name
}

// Configuration returns the configuration (args) for [EcsCapacityProvider].
func (ecp *EcsCapacityProvider) Configuration() interface{} {
	return ecp.Args
}

// DependOn is used for other resources to depend on [EcsCapacityProvider].
func (ecp *EcsCapacityProvider) DependOn() terra.Reference {
	return terra.ReferenceResource(ecp)
}

// Dependencies returns the list of resources [EcsCapacityProvider] depends_on.
func (ecp *EcsCapacityProvider) Dependencies() terra.Dependencies {
	return ecp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [EcsCapacityProvider].
func (ecp *EcsCapacityProvider) LifecycleManagement() *terra.Lifecycle {
	return ecp.Lifecycle
}

// Attributes returns the attributes for [EcsCapacityProvider].
func (ecp *EcsCapacityProvider) Attributes() ecsCapacityProviderAttributes {
	return ecsCapacityProviderAttributes{ref: terra.ReferenceResource(ecp)}
}

// ImportState imports the given attribute values into [EcsCapacityProvider]'s state.
func (ecp *EcsCapacityProvider) ImportState(av io.Reader) error {
	ecp.state = &ecsCapacityProviderState{}
	if err := json.NewDecoder(av).Decode(ecp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ecp.Type(), ecp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [EcsCapacityProvider] has state.
func (ecp *EcsCapacityProvider) State() (*ecsCapacityProviderState, bool) {
	return ecp.state, ecp.state != nil
}

// StateMust returns the state for [EcsCapacityProvider]. Panics if the state is nil.
func (ecp *EcsCapacityProvider) StateMust() *ecsCapacityProviderState {
	if ecp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ecp.Type(), ecp.LocalName()))
	}
	return ecp.state
}

// EcsCapacityProviderArgs contains the configurations for aws_ecs_capacity_provider.
type EcsCapacityProviderArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// AutoScalingGroupProvider: required
	AutoScalingGroupProvider *ecscapacityprovider.AutoScalingGroupProvider `hcl:"auto_scaling_group_provider,block" validate:"required"`
}
type ecsCapacityProviderAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_ecs_capacity_provider.
func (ecp ecsCapacityProviderAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ecp.ref.Append("arn"))
}

// Id returns a reference to field id of aws_ecs_capacity_provider.
func (ecp ecsCapacityProviderAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ecp.ref.Append("id"))
}

// Name returns a reference to field name of aws_ecs_capacity_provider.
func (ecp ecsCapacityProviderAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ecp.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_ecs_capacity_provider.
func (ecp ecsCapacityProviderAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ecp.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_ecs_capacity_provider.
func (ecp ecsCapacityProviderAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ecp.ref.Append("tags_all"))
}

func (ecp ecsCapacityProviderAttributes) AutoScalingGroupProvider() terra.ListValue[ecscapacityprovider.AutoScalingGroupProviderAttributes] {
	return terra.ReferenceAsList[ecscapacityprovider.AutoScalingGroupProviderAttributes](ecp.ref.Append("auto_scaling_group_provider"))
}

type ecsCapacityProviderState struct {
	Arn                      string                                              `json:"arn"`
	Id                       string                                              `json:"id"`
	Name                     string                                              `json:"name"`
	Tags                     map[string]string                                   `json:"tags"`
	TagsAll                  map[string]string                                   `json:"tags_all"`
	AutoScalingGroupProvider []ecscapacityprovider.AutoScalingGroupProviderState `json:"auto_scaling_group_provider"`
}
