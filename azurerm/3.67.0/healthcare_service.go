// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	healthcareservice "github.com/golingon/terraproviders/azurerm/3.67.0/healthcareservice"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewHealthcareService creates a new instance of [HealthcareService].
func NewHealthcareService(name string, args HealthcareServiceArgs) *HealthcareService {
	return &HealthcareService{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*HealthcareService)(nil)

// HealthcareService represents the Terraform resource azurerm_healthcare_service.
type HealthcareService struct {
	Name      string
	Args      HealthcareServiceArgs
	state     *healthcareServiceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [HealthcareService].
func (hs *HealthcareService) Type() string {
	return "azurerm_healthcare_service"
}

// LocalName returns the local name for [HealthcareService].
func (hs *HealthcareService) LocalName() string {
	return hs.Name
}

// Configuration returns the configuration (args) for [HealthcareService].
func (hs *HealthcareService) Configuration() interface{} {
	return hs.Args
}

// DependOn is used for other resources to depend on [HealthcareService].
func (hs *HealthcareService) DependOn() terra.Reference {
	return terra.ReferenceResource(hs)
}

// Dependencies returns the list of resources [HealthcareService] depends_on.
func (hs *HealthcareService) Dependencies() terra.Dependencies {
	return hs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [HealthcareService].
func (hs *HealthcareService) LifecycleManagement() *terra.Lifecycle {
	return hs.Lifecycle
}

// Attributes returns the attributes for [HealthcareService].
func (hs *HealthcareService) Attributes() healthcareServiceAttributes {
	return healthcareServiceAttributes{ref: terra.ReferenceResource(hs)}
}

// ImportState imports the given attribute values into [HealthcareService]'s state.
func (hs *HealthcareService) ImportState(av io.Reader) error {
	hs.state = &healthcareServiceState{}
	if err := json.NewDecoder(av).Decode(hs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", hs.Type(), hs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [HealthcareService] has state.
func (hs *HealthcareService) State() (*healthcareServiceState, bool) {
	return hs.state, hs.state != nil
}

// StateMust returns the state for [HealthcareService]. Panics if the state is nil.
func (hs *HealthcareService) StateMust() *healthcareServiceState {
	if hs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", hs.Type(), hs.LocalName()))
	}
	return hs.state
}

// HealthcareServiceArgs contains the configurations for azurerm_healthcare_service.
type HealthcareServiceArgs struct {
	// AccessPolicyObjectIds: set of string, optional
	AccessPolicyObjectIds terra.SetValue[terra.StringValue] `hcl:"access_policy_object_ids,attr"`
	// CosmosdbKeyVaultKeyVersionlessId: string, optional
	CosmosdbKeyVaultKeyVersionlessId terra.StringValue `hcl:"cosmosdb_key_vault_key_versionless_id,attr"`
	// CosmosdbThroughput: number, optional
	CosmosdbThroughput terra.NumberValue `hcl:"cosmosdb_throughput,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Kind: string, optional
	Kind terra.StringValue `hcl:"kind,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PublicNetworkAccessEnabled: bool, optional
	PublicNetworkAccessEnabled terra.BoolValue `hcl:"public_network_access_enabled,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// AuthenticationConfiguration: optional
	AuthenticationConfiguration *healthcareservice.AuthenticationConfiguration `hcl:"authentication_configuration,block"`
	// CorsConfiguration: optional
	CorsConfiguration *healthcareservice.CorsConfiguration `hcl:"cors_configuration,block"`
	// Timeouts: optional
	Timeouts *healthcareservice.Timeouts `hcl:"timeouts,block"`
}
type healthcareServiceAttributes struct {
	ref terra.Reference
}

// AccessPolicyObjectIds returns a reference to field access_policy_object_ids of azurerm_healthcare_service.
func (hs healthcareServiceAttributes) AccessPolicyObjectIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](hs.ref.Append("access_policy_object_ids"))
}

// CosmosdbKeyVaultKeyVersionlessId returns a reference to field cosmosdb_key_vault_key_versionless_id of azurerm_healthcare_service.
func (hs healthcareServiceAttributes) CosmosdbKeyVaultKeyVersionlessId() terra.StringValue {
	return terra.ReferenceAsString(hs.ref.Append("cosmosdb_key_vault_key_versionless_id"))
}

// CosmosdbThroughput returns a reference to field cosmosdb_throughput of azurerm_healthcare_service.
func (hs healthcareServiceAttributes) CosmosdbThroughput() terra.NumberValue {
	return terra.ReferenceAsNumber(hs.ref.Append("cosmosdb_throughput"))
}

// Id returns a reference to field id of azurerm_healthcare_service.
func (hs healthcareServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(hs.ref.Append("id"))
}

// Kind returns a reference to field kind of azurerm_healthcare_service.
func (hs healthcareServiceAttributes) Kind() terra.StringValue {
	return terra.ReferenceAsString(hs.ref.Append("kind"))
}

// Location returns a reference to field location of azurerm_healthcare_service.
func (hs healthcareServiceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(hs.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_healthcare_service.
func (hs healthcareServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(hs.ref.Append("name"))
}

// PublicNetworkAccessEnabled returns a reference to field public_network_access_enabled of azurerm_healthcare_service.
func (hs healthcareServiceAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(hs.ref.Append("public_network_access_enabled"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_healthcare_service.
func (hs healthcareServiceAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(hs.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_healthcare_service.
func (hs healthcareServiceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](hs.ref.Append("tags"))
}

func (hs healthcareServiceAttributes) AuthenticationConfiguration() terra.ListValue[healthcareservice.AuthenticationConfigurationAttributes] {
	return terra.ReferenceAsList[healthcareservice.AuthenticationConfigurationAttributes](hs.ref.Append("authentication_configuration"))
}

func (hs healthcareServiceAttributes) CorsConfiguration() terra.ListValue[healthcareservice.CorsConfigurationAttributes] {
	return terra.ReferenceAsList[healthcareservice.CorsConfigurationAttributes](hs.ref.Append("cors_configuration"))
}

func (hs healthcareServiceAttributes) Timeouts() healthcareservice.TimeoutsAttributes {
	return terra.ReferenceAsSingle[healthcareservice.TimeoutsAttributes](hs.ref.Append("timeouts"))
}

type healthcareServiceState struct {
	AccessPolicyObjectIds            []string                                             `json:"access_policy_object_ids"`
	CosmosdbKeyVaultKeyVersionlessId string                                               `json:"cosmosdb_key_vault_key_versionless_id"`
	CosmosdbThroughput               float64                                              `json:"cosmosdb_throughput"`
	Id                               string                                               `json:"id"`
	Kind                             string                                               `json:"kind"`
	Location                         string                                               `json:"location"`
	Name                             string                                               `json:"name"`
	PublicNetworkAccessEnabled       bool                                                 `json:"public_network_access_enabled"`
	ResourceGroupName                string                                               `json:"resource_group_name"`
	Tags                             map[string]string                                    `json:"tags"`
	AuthenticationConfiguration      []healthcareservice.AuthenticationConfigurationState `json:"authentication_configuration"`
	CorsConfiguration                []healthcareservice.CorsConfigurationState           `json:"cors_configuration"`
	Timeouts                         *healthcareservice.TimeoutsState                     `json:"timeouts"`
}
