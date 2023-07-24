// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	subnetserviceendpointstoragepolicy "github.com/golingon/terraproviders/azurerm/3.66.0/subnetserviceendpointstoragepolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSubnetServiceEndpointStoragePolicy creates a new instance of [SubnetServiceEndpointStoragePolicy].
func NewSubnetServiceEndpointStoragePolicy(name string, args SubnetServiceEndpointStoragePolicyArgs) *SubnetServiceEndpointStoragePolicy {
	return &SubnetServiceEndpointStoragePolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SubnetServiceEndpointStoragePolicy)(nil)

// SubnetServiceEndpointStoragePolicy represents the Terraform resource azurerm_subnet_service_endpoint_storage_policy.
type SubnetServiceEndpointStoragePolicy struct {
	Name      string
	Args      SubnetServiceEndpointStoragePolicyArgs
	state     *subnetServiceEndpointStoragePolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SubnetServiceEndpointStoragePolicy].
func (ssesp *SubnetServiceEndpointStoragePolicy) Type() string {
	return "azurerm_subnet_service_endpoint_storage_policy"
}

// LocalName returns the local name for [SubnetServiceEndpointStoragePolicy].
func (ssesp *SubnetServiceEndpointStoragePolicy) LocalName() string {
	return ssesp.Name
}

// Configuration returns the configuration (args) for [SubnetServiceEndpointStoragePolicy].
func (ssesp *SubnetServiceEndpointStoragePolicy) Configuration() interface{} {
	return ssesp.Args
}

// DependOn is used for other resources to depend on [SubnetServiceEndpointStoragePolicy].
func (ssesp *SubnetServiceEndpointStoragePolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(ssesp)
}

// Dependencies returns the list of resources [SubnetServiceEndpointStoragePolicy] depends_on.
func (ssesp *SubnetServiceEndpointStoragePolicy) Dependencies() terra.Dependencies {
	return ssesp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SubnetServiceEndpointStoragePolicy].
func (ssesp *SubnetServiceEndpointStoragePolicy) LifecycleManagement() *terra.Lifecycle {
	return ssesp.Lifecycle
}

// Attributes returns the attributes for [SubnetServiceEndpointStoragePolicy].
func (ssesp *SubnetServiceEndpointStoragePolicy) Attributes() subnetServiceEndpointStoragePolicyAttributes {
	return subnetServiceEndpointStoragePolicyAttributes{ref: terra.ReferenceResource(ssesp)}
}

// ImportState imports the given attribute values into [SubnetServiceEndpointStoragePolicy]'s state.
func (ssesp *SubnetServiceEndpointStoragePolicy) ImportState(av io.Reader) error {
	ssesp.state = &subnetServiceEndpointStoragePolicyState{}
	if err := json.NewDecoder(av).Decode(ssesp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ssesp.Type(), ssesp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SubnetServiceEndpointStoragePolicy] has state.
func (ssesp *SubnetServiceEndpointStoragePolicy) State() (*subnetServiceEndpointStoragePolicyState, bool) {
	return ssesp.state, ssesp.state != nil
}

// StateMust returns the state for [SubnetServiceEndpointStoragePolicy]. Panics if the state is nil.
func (ssesp *SubnetServiceEndpointStoragePolicy) StateMust() *subnetServiceEndpointStoragePolicyState {
	if ssesp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ssesp.Type(), ssesp.LocalName()))
	}
	return ssesp.state
}

// SubnetServiceEndpointStoragePolicyArgs contains the configurations for azurerm_subnet_service_endpoint_storage_policy.
type SubnetServiceEndpointStoragePolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Definition: min=0,max=2
	Definition []subnetserviceendpointstoragepolicy.Definition `hcl:"definition,block" validate:"min=0,max=2"`
	// Timeouts: optional
	Timeouts *subnetserviceendpointstoragepolicy.Timeouts `hcl:"timeouts,block"`
}
type subnetServiceEndpointStoragePolicyAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_subnet_service_endpoint_storage_policy.
func (ssesp subnetServiceEndpointStoragePolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ssesp.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_subnet_service_endpoint_storage_policy.
func (ssesp subnetServiceEndpointStoragePolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ssesp.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_subnet_service_endpoint_storage_policy.
func (ssesp subnetServiceEndpointStoragePolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ssesp.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_subnet_service_endpoint_storage_policy.
func (ssesp subnetServiceEndpointStoragePolicyAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ssesp.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_subnet_service_endpoint_storage_policy.
func (ssesp subnetServiceEndpointStoragePolicyAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ssesp.ref.Append("tags"))
}

func (ssesp subnetServiceEndpointStoragePolicyAttributes) Definition() terra.ListValue[subnetserviceendpointstoragepolicy.DefinitionAttributes] {
	return terra.ReferenceAsList[subnetserviceendpointstoragepolicy.DefinitionAttributes](ssesp.ref.Append("definition"))
}

func (ssesp subnetServiceEndpointStoragePolicyAttributes) Timeouts() subnetserviceendpointstoragepolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[subnetserviceendpointstoragepolicy.TimeoutsAttributes](ssesp.ref.Append("timeouts"))
}

type subnetServiceEndpointStoragePolicyState struct {
	Id                string                                               `json:"id"`
	Location          string                                               `json:"location"`
	Name              string                                               `json:"name"`
	ResourceGroupName string                                               `json:"resource_group_name"`
	Tags              map[string]string                                    `json:"tags"`
	Definition        []subnetserviceendpointstoragepolicy.DefinitionState `json:"definition"`
	Timeouts          *subnetserviceendpointstoragepolicy.TimeoutsState    `json:"timeouts"`
}
