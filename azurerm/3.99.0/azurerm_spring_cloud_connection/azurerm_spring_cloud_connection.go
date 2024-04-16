// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_spring_cloud_connection

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

// Resource represents the Terraform resource azurerm_spring_cloud_connection.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermSpringCloudConnectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (ascc *Resource) Type() string {
	return "azurerm_spring_cloud_connection"
}

// LocalName returns the local name for [Resource].
func (ascc *Resource) LocalName() string {
	return ascc.Name
}

// Configuration returns the configuration (args) for [Resource].
func (ascc *Resource) Configuration() interface{} {
	return ascc.Args
}

// DependOn is used for other resources to depend on [Resource].
func (ascc *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(ascc)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (ascc *Resource) Dependencies() terra.Dependencies {
	return ascc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (ascc *Resource) LifecycleManagement() *terra.Lifecycle {
	return ascc.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (ascc *Resource) Attributes() azurermSpringCloudConnectionAttributes {
	return azurermSpringCloudConnectionAttributes{ref: terra.ReferenceResource(ascc)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (ascc *Resource) ImportState(state io.Reader) error {
	ascc.state = &azurermSpringCloudConnectionState{}
	if err := json.NewDecoder(state).Decode(ascc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ascc.Type(), ascc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (ascc *Resource) State() (*azurermSpringCloudConnectionState, bool) {
	return ascc.state, ascc.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (ascc *Resource) StateMust() *azurermSpringCloudConnectionState {
	if ascc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ascc.Type(), ascc.LocalName()))
	}
	return ascc.state
}

// Args contains the configurations for azurerm_spring_cloud_connection.
type Args struct {
	// ClientType: string, optional
	ClientType terra.StringValue `hcl:"client_type,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SpringCloudId: string, required
	SpringCloudId terra.StringValue `hcl:"spring_cloud_id,attr" validate:"required"`
	// TargetResourceId: string, required
	TargetResourceId terra.StringValue `hcl:"target_resource_id,attr" validate:"required"`
	// VnetSolution: string, optional
	VnetSolution terra.StringValue `hcl:"vnet_solution,attr"`
	// Authentication: required
	Authentication *Authentication `hcl:"authentication,block" validate:"required"`
	// SecretStore: optional
	SecretStore *SecretStore `hcl:"secret_store,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermSpringCloudConnectionAttributes struct {
	ref terra.Reference
}

// ClientType returns a reference to field client_type of azurerm_spring_cloud_connection.
func (ascc azurermSpringCloudConnectionAttributes) ClientType() terra.StringValue {
	return terra.ReferenceAsString(ascc.ref.Append("client_type"))
}

// Id returns a reference to field id of azurerm_spring_cloud_connection.
func (ascc azurermSpringCloudConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ascc.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_spring_cloud_connection.
func (ascc azurermSpringCloudConnectionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ascc.ref.Append("name"))
}

// SpringCloudId returns a reference to field spring_cloud_id of azurerm_spring_cloud_connection.
func (ascc azurermSpringCloudConnectionAttributes) SpringCloudId() terra.StringValue {
	return terra.ReferenceAsString(ascc.ref.Append("spring_cloud_id"))
}

// TargetResourceId returns a reference to field target_resource_id of azurerm_spring_cloud_connection.
func (ascc azurermSpringCloudConnectionAttributes) TargetResourceId() terra.StringValue {
	return terra.ReferenceAsString(ascc.ref.Append("target_resource_id"))
}

// VnetSolution returns a reference to field vnet_solution of azurerm_spring_cloud_connection.
func (ascc azurermSpringCloudConnectionAttributes) VnetSolution() terra.StringValue {
	return terra.ReferenceAsString(ascc.ref.Append("vnet_solution"))
}

func (ascc azurermSpringCloudConnectionAttributes) Authentication() terra.ListValue[AuthenticationAttributes] {
	return terra.ReferenceAsList[AuthenticationAttributes](ascc.ref.Append("authentication"))
}

func (ascc azurermSpringCloudConnectionAttributes) SecretStore() terra.ListValue[SecretStoreAttributes] {
	return terra.ReferenceAsList[SecretStoreAttributes](ascc.ref.Append("secret_store"))
}

func (ascc azurermSpringCloudConnectionAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](ascc.ref.Append("timeouts"))
}

type azurermSpringCloudConnectionState struct {
	ClientType       string                `json:"client_type"`
	Id               string                `json:"id"`
	Name             string                `json:"name"`
	SpringCloudId    string                `json:"spring_cloud_id"`
	TargetResourceId string                `json:"target_resource_id"`
	VnetSolution     string                `json:"vnet_solution"`
	Authentication   []AuthenticationState `json:"authentication"`
	SecretStore      []SecretStoreState    `json:"secret_store"`
	Timeouts         *TimeoutsState        `json:"timeouts"`
}
