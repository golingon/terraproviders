// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	healthcarefhirservice "github.com/golingon/terraproviders/azurerm/3.52.0/healthcarefhirservice"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewHealthcareFhirService creates a new instance of [HealthcareFhirService].
func NewHealthcareFhirService(name string, args HealthcareFhirServiceArgs) *HealthcareFhirService {
	return &HealthcareFhirService{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*HealthcareFhirService)(nil)

// HealthcareFhirService represents the Terraform resource azurerm_healthcare_fhir_service.
type HealthcareFhirService struct {
	Name      string
	Args      HealthcareFhirServiceArgs
	state     *healthcareFhirServiceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [HealthcareFhirService].
func (hfs *HealthcareFhirService) Type() string {
	return "azurerm_healthcare_fhir_service"
}

// LocalName returns the local name for [HealthcareFhirService].
func (hfs *HealthcareFhirService) LocalName() string {
	return hfs.Name
}

// Configuration returns the configuration (args) for [HealthcareFhirService].
func (hfs *HealthcareFhirService) Configuration() interface{} {
	return hfs.Args
}

// DependOn is used for other resources to depend on [HealthcareFhirService].
func (hfs *HealthcareFhirService) DependOn() terra.Reference {
	return terra.ReferenceResource(hfs)
}

// Dependencies returns the list of resources [HealthcareFhirService] depends_on.
func (hfs *HealthcareFhirService) Dependencies() terra.Dependencies {
	return hfs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [HealthcareFhirService].
func (hfs *HealthcareFhirService) LifecycleManagement() *terra.Lifecycle {
	return hfs.Lifecycle
}

// Attributes returns the attributes for [HealthcareFhirService].
func (hfs *HealthcareFhirService) Attributes() healthcareFhirServiceAttributes {
	return healthcareFhirServiceAttributes{ref: terra.ReferenceResource(hfs)}
}

// ImportState imports the given attribute values into [HealthcareFhirService]'s state.
func (hfs *HealthcareFhirService) ImportState(av io.Reader) error {
	hfs.state = &healthcareFhirServiceState{}
	if err := json.NewDecoder(av).Decode(hfs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", hfs.Type(), hfs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [HealthcareFhirService] has state.
func (hfs *HealthcareFhirService) State() (*healthcareFhirServiceState, bool) {
	return hfs.state, hfs.state != nil
}

// StateMust returns the state for [HealthcareFhirService]. Panics if the state is nil.
func (hfs *HealthcareFhirService) StateMust() *healthcareFhirServiceState {
	if hfs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", hfs.Type(), hfs.LocalName()))
	}
	return hfs.state
}

// HealthcareFhirServiceArgs contains the configurations for azurerm_healthcare_fhir_service.
type HealthcareFhirServiceArgs struct {
	// AccessPolicyObjectIds: set of string, optional
	AccessPolicyObjectIds terra.SetValue[terra.StringValue] `hcl:"access_policy_object_ids,attr"`
	// ConfigurationExportStorageAccountName: string, optional
	ConfigurationExportStorageAccountName terra.StringValue `hcl:"configuration_export_storage_account_name,attr"`
	// ContainerRegistryLoginServerUrl: set of string, optional
	ContainerRegistryLoginServerUrl terra.SetValue[terra.StringValue] `hcl:"container_registry_login_server_url,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Kind: string, optional
	Kind terra.StringValue `hcl:"kind,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// WorkspaceId: string, required
	WorkspaceId terra.StringValue `hcl:"workspace_id,attr" validate:"required"`
	// Authentication: required
	Authentication *healthcarefhirservice.Authentication `hcl:"authentication,block" validate:"required"`
	// Cors: optional
	Cors *healthcarefhirservice.Cors `hcl:"cors,block"`
	// Identity: optional
	Identity *healthcarefhirservice.Identity `hcl:"identity,block"`
	// OciArtifact: min=0
	OciArtifact []healthcarefhirservice.OciArtifact `hcl:"oci_artifact,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *healthcarefhirservice.Timeouts `hcl:"timeouts,block"`
}
type healthcareFhirServiceAttributes struct {
	ref terra.Reference
}

// AccessPolicyObjectIds returns a reference to field access_policy_object_ids of azurerm_healthcare_fhir_service.
func (hfs healthcareFhirServiceAttributes) AccessPolicyObjectIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](hfs.ref.Append("access_policy_object_ids"))
}

// ConfigurationExportStorageAccountName returns a reference to field configuration_export_storage_account_name of azurerm_healthcare_fhir_service.
func (hfs healthcareFhirServiceAttributes) ConfigurationExportStorageAccountName() terra.StringValue {
	return terra.ReferenceAsString(hfs.ref.Append("configuration_export_storage_account_name"))
}

// ContainerRegistryLoginServerUrl returns a reference to field container_registry_login_server_url of azurerm_healthcare_fhir_service.
func (hfs healthcareFhirServiceAttributes) ContainerRegistryLoginServerUrl() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](hfs.ref.Append("container_registry_login_server_url"))
}

// Id returns a reference to field id of azurerm_healthcare_fhir_service.
func (hfs healthcareFhirServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(hfs.ref.Append("id"))
}

// Kind returns a reference to field kind of azurerm_healthcare_fhir_service.
func (hfs healthcareFhirServiceAttributes) Kind() terra.StringValue {
	return terra.ReferenceAsString(hfs.ref.Append("kind"))
}

// Location returns a reference to field location of azurerm_healthcare_fhir_service.
func (hfs healthcareFhirServiceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(hfs.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_healthcare_fhir_service.
func (hfs healthcareFhirServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(hfs.ref.Append("name"))
}

// PublicNetworkAccessEnabled returns a reference to field public_network_access_enabled of azurerm_healthcare_fhir_service.
func (hfs healthcareFhirServiceAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(hfs.ref.Append("public_network_access_enabled"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_healthcare_fhir_service.
func (hfs healthcareFhirServiceAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(hfs.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_healthcare_fhir_service.
func (hfs healthcareFhirServiceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](hfs.ref.Append("tags"))
}

// WorkspaceId returns a reference to field workspace_id of azurerm_healthcare_fhir_service.
func (hfs healthcareFhirServiceAttributes) WorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(hfs.ref.Append("workspace_id"))
}

func (hfs healthcareFhirServiceAttributes) Authentication() terra.ListValue[healthcarefhirservice.AuthenticationAttributes] {
	return terra.ReferenceAsList[healthcarefhirservice.AuthenticationAttributes](hfs.ref.Append("authentication"))
}

func (hfs healthcareFhirServiceAttributes) Cors() terra.ListValue[healthcarefhirservice.CorsAttributes] {
	return terra.ReferenceAsList[healthcarefhirservice.CorsAttributes](hfs.ref.Append("cors"))
}

func (hfs healthcareFhirServiceAttributes) Identity() terra.ListValue[healthcarefhirservice.IdentityAttributes] {
	return terra.ReferenceAsList[healthcarefhirservice.IdentityAttributes](hfs.ref.Append("identity"))
}

func (hfs healthcareFhirServiceAttributes) OciArtifact() terra.ListValue[healthcarefhirservice.OciArtifactAttributes] {
	return terra.ReferenceAsList[healthcarefhirservice.OciArtifactAttributes](hfs.ref.Append("oci_artifact"))
}

func (hfs healthcareFhirServiceAttributes) Timeouts() healthcarefhirservice.TimeoutsAttributes {
	return terra.ReferenceAsSingle[healthcarefhirservice.TimeoutsAttributes](hfs.ref.Append("timeouts"))
}

type healthcareFhirServiceState struct {
	AccessPolicyObjectIds                 []string                                    `json:"access_policy_object_ids"`
	ConfigurationExportStorageAccountName string                                      `json:"configuration_export_storage_account_name"`
	ContainerRegistryLoginServerUrl       []string                                    `json:"container_registry_login_server_url"`
	Id                                    string                                      `json:"id"`
	Kind                                  string                                      `json:"kind"`
	Location                              string                                      `json:"location"`
	Name                                  string                                      `json:"name"`
	PublicNetworkAccessEnabled            bool                                        `json:"public_network_access_enabled"`
	ResourceGroupName                     string                                      `json:"resource_group_name"`
	Tags                                  map[string]string                           `json:"tags"`
	WorkspaceId                           string                                      `json:"workspace_id"`
	Authentication                        []healthcarefhirservice.AuthenticationState `json:"authentication"`
	Cors                                  []healthcarefhirservice.CorsState           `json:"cors"`
	Identity                              []healthcarefhirservice.IdentityState       `json:"identity"`
	OciArtifact                           []healthcarefhirservice.OciArtifactState    `json:"oci_artifact"`
	Timeouts                              *healthcarefhirservice.TimeoutsState        `json:"timeouts"`
}
