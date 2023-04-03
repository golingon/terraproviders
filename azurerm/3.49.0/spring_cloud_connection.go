// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	springcloudconnection "github.com/golingon/terraproviders/azurerm/3.49.0/springcloudconnection"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSpringCloudConnection creates a new instance of [SpringCloudConnection].
func NewSpringCloudConnection(name string, args SpringCloudConnectionArgs) *SpringCloudConnection {
	return &SpringCloudConnection{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SpringCloudConnection)(nil)

// SpringCloudConnection represents the Terraform resource azurerm_spring_cloud_connection.
type SpringCloudConnection struct {
	Name      string
	Args      SpringCloudConnectionArgs
	state     *springCloudConnectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SpringCloudConnection].
func (scc *SpringCloudConnection) Type() string {
	return "azurerm_spring_cloud_connection"
}

// LocalName returns the local name for [SpringCloudConnection].
func (scc *SpringCloudConnection) LocalName() string {
	return scc.Name
}

// Configuration returns the configuration (args) for [SpringCloudConnection].
func (scc *SpringCloudConnection) Configuration() interface{} {
	return scc.Args
}

// DependOn is used for other resources to depend on [SpringCloudConnection].
func (scc *SpringCloudConnection) DependOn() terra.Reference {
	return terra.ReferenceResource(scc)
}

// Dependencies returns the list of resources [SpringCloudConnection] depends_on.
func (scc *SpringCloudConnection) Dependencies() terra.Dependencies {
	return scc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SpringCloudConnection].
func (scc *SpringCloudConnection) LifecycleManagement() *terra.Lifecycle {
	return scc.Lifecycle
}

// Attributes returns the attributes for [SpringCloudConnection].
func (scc *SpringCloudConnection) Attributes() springCloudConnectionAttributes {
	return springCloudConnectionAttributes{ref: terra.ReferenceResource(scc)}
}

// ImportState imports the given attribute values into [SpringCloudConnection]'s state.
func (scc *SpringCloudConnection) ImportState(av io.Reader) error {
	scc.state = &springCloudConnectionState{}
	if err := json.NewDecoder(av).Decode(scc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", scc.Type(), scc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SpringCloudConnection] has state.
func (scc *SpringCloudConnection) State() (*springCloudConnectionState, bool) {
	return scc.state, scc.state != nil
}

// StateMust returns the state for [SpringCloudConnection]. Panics if the state is nil.
func (scc *SpringCloudConnection) StateMust() *springCloudConnectionState {
	if scc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", scc.Type(), scc.LocalName()))
	}
	return scc.state
}

// SpringCloudConnectionArgs contains the configurations for azurerm_spring_cloud_connection.
type SpringCloudConnectionArgs struct {
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
	Authentication *springcloudconnection.Authentication `hcl:"authentication,block" validate:"required"`
	// SecretStore: optional
	SecretStore *springcloudconnection.SecretStore `hcl:"secret_store,block"`
	// Timeouts: optional
	Timeouts *springcloudconnection.Timeouts `hcl:"timeouts,block"`
}
type springCloudConnectionAttributes struct {
	ref terra.Reference
}

// ClientType returns a reference to field client_type of azurerm_spring_cloud_connection.
func (scc springCloudConnectionAttributes) ClientType() terra.StringValue {
	return terra.ReferenceAsString(scc.ref.Append("client_type"))
}

// Id returns a reference to field id of azurerm_spring_cloud_connection.
func (scc springCloudConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(scc.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_spring_cloud_connection.
func (scc springCloudConnectionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(scc.ref.Append("name"))
}

// SpringCloudId returns a reference to field spring_cloud_id of azurerm_spring_cloud_connection.
func (scc springCloudConnectionAttributes) SpringCloudId() terra.StringValue {
	return terra.ReferenceAsString(scc.ref.Append("spring_cloud_id"))
}

// TargetResourceId returns a reference to field target_resource_id of azurerm_spring_cloud_connection.
func (scc springCloudConnectionAttributes) TargetResourceId() terra.StringValue {
	return terra.ReferenceAsString(scc.ref.Append("target_resource_id"))
}

// VnetSolution returns a reference to field vnet_solution of azurerm_spring_cloud_connection.
func (scc springCloudConnectionAttributes) VnetSolution() terra.StringValue {
	return terra.ReferenceAsString(scc.ref.Append("vnet_solution"))
}

func (scc springCloudConnectionAttributes) Authentication() terra.ListValue[springcloudconnection.AuthenticationAttributes] {
	return terra.ReferenceAsList[springcloudconnection.AuthenticationAttributes](scc.ref.Append("authentication"))
}

func (scc springCloudConnectionAttributes) SecretStore() terra.ListValue[springcloudconnection.SecretStoreAttributes] {
	return terra.ReferenceAsList[springcloudconnection.SecretStoreAttributes](scc.ref.Append("secret_store"))
}

func (scc springCloudConnectionAttributes) Timeouts() springcloudconnection.TimeoutsAttributes {
	return terra.ReferenceAsSingle[springcloudconnection.TimeoutsAttributes](scc.ref.Append("timeouts"))
}

type springCloudConnectionState struct {
	ClientType       string                                      `json:"client_type"`
	Id               string                                      `json:"id"`
	Name             string                                      `json:"name"`
	SpringCloudId    string                                      `json:"spring_cloud_id"`
	TargetResourceId string                                      `json:"target_resource_id"`
	VnetSolution     string                                      `json:"vnet_solution"`
	Authentication   []springcloudconnection.AuthenticationState `json:"authentication"`
	SecretStore      []springcloudconnection.SecretStoreState    `json:"secret_store"`
	Timeouts         *springcloudconnection.TimeoutsState        `json:"timeouts"`
}
