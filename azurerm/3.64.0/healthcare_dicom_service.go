// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	healthcaredicomservice "github.com/golingon/terraproviders/azurerm/3.64.0/healthcaredicomservice"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewHealthcareDicomService creates a new instance of [HealthcareDicomService].
func NewHealthcareDicomService(name string, args HealthcareDicomServiceArgs) *HealthcareDicomService {
	return &HealthcareDicomService{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*HealthcareDicomService)(nil)

// HealthcareDicomService represents the Terraform resource azurerm_healthcare_dicom_service.
type HealthcareDicomService struct {
	Name      string
	Args      HealthcareDicomServiceArgs
	state     *healthcareDicomServiceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [HealthcareDicomService].
func (hds *HealthcareDicomService) Type() string {
	return "azurerm_healthcare_dicom_service"
}

// LocalName returns the local name for [HealthcareDicomService].
func (hds *HealthcareDicomService) LocalName() string {
	return hds.Name
}

// Configuration returns the configuration (args) for [HealthcareDicomService].
func (hds *HealthcareDicomService) Configuration() interface{} {
	return hds.Args
}

// DependOn is used for other resources to depend on [HealthcareDicomService].
func (hds *HealthcareDicomService) DependOn() terra.Reference {
	return terra.ReferenceResource(hds)
}

// Dependencies returns the list of resources [HealthcareDicomService] depends_on.
func (hds *HealthcareDicomService) Dependencies() terra.Dependencies {
	return hds.DependsOn
}

// LifecycleManagement returns the lifecycle block for [HealthcareDicomService].
func (hds *HealthcareDicomService) LifecycleManagement() *terra.Lifecycle {
	return hds.Lifecycle
}

// Attributes returns the attributes for [HealthcareDicomService].
func (hds *HealthcareDicomService) Attributes() healthcareDicomServiceAttributes {
	return healthcareDicomServiceAttributes{ref: terra.ReferenceResource(hds)}
}

// ImportState imports the given attribute values into [HealthcareDicomService]'s state.
func (hds *HealthcareDicomService) ImportState(av io.Reader) error {
	hds.state = &healthcareDicomServiceState{}
	if err := json.NewDecoder(av).Decode(hds.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", hds.Type(), hds.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [HealthcareDicomService] has state.
func (hds *HealthcareDicomService) State() (*healthcareDicomServiceState, bool) {
	return hds.state, hds.state != nil
}

// StateMust returns the state for [HealthcareDicomService]. Panics if the state is nil.
func (hds *HealthcareDicomService) StateMust() *healthcareDicomServiceState {
	if hds.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", hds.Type(), hds.LocalName()))
	}
	return hds.state
}

// HealthcareDicomServiceArgs contains the configurations for azurerm_healthcare_dicom_service.
type HealthcareDicomServiceArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PublicNetworkAccessEnabled: bool, optional
	PublicNetworkAccessEnabled terra.BoolValue `hcl:"public_network_access_enabled,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// WorkspaceId: string, required
	WorkspaceId terra.StringValue `hcl:"workspace_id,attr" validate:"required"`
	// Authentication: min=0
	Authentication []healthcaredicomservice.Authentication `hcl:"authentication,block" validate:"min=0"`
	// PrivateEndpoint: min=0
	PrivateEndpoint []healthcaredicomservice.PrivateEndpoint `hcl:"private_endpoint,block" validate:"min=0"`
	// Identity: optional
	Identity *healthcaredicomservice.Identity `hcl:"identity,block"`
	// Timeouts: optional
	Timeouts *healthcaredicomservice.Timeouts `hcl:"timeouts,block"`
}
type healthcareDicomServiceAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_healthcare_dicom_service.
func (hds healthcareDicomServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(hds.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_healthcare_dicom_service.
func (hds healthcareDicomServiceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(hds.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_healthcare_dicom_service.
func (hds healthcareDicomServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(hds.ref.Append("name"))
}

// PublicNetworkAccessEnabled returns a reference to field public_network_access_enabled of azurerm_healthcare_dicom_service.
func (hds healthcareDicomServiceAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(hds.ref.Append("public_network_access_enabled"))
}

// ServiceUrl returns a reference to field service_url of azurerm_healthcare_dicom_service.
func (hds healthcareDicomServiceAttributes) ServiceUrl() terra.StringValue {
	return terra.ReferenceAsString(hds.ref.Append("service_url"))
}

// Tags returns a reference to field tags of azurerm_healthcare_dicom_service.
func (hds healthcareDicomServiceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](hds.ref.Append("tags"))
}

// WorkspaceId returns a reference to field workspace_id of azurerm_healthcare_dicom_service.
func (hds healthcareDicomServiceAttributes) WorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(hds.ref.Append("workspace_id"))
}

func (hds healthcareDicomServiceAttributes) Authentication() terra.ListValue[healthcaredicomservice.AuthenticationAttributes] {
	return terra.ReferenceAsList[healthcaredicomservice.AuthenticationAttributes](hds.ref.Append("authentication"))
}

func (hds healthcareDicomServiceAttributes) PrivateEndpoint() terra.SetValue[healthcaredicomservice.PrivateEndpointAttributes] {
	return terra.ReferenceAsSet[healthcaredicomservice.PrivateEndpointAttributes](hds.ref.Append("private_endpoint"))
}

func (hds healthcareDicomServiceAttributes) Identity() terra.ListValue[healthcaredicomservice.IdentityAttributes] {
	return terra.ReferenceAsList[healthcaredicomservice.IdentityAttributes](hds.ref.Append("identity"))
}

func (hds healthcareDicomServiceAttributes) Timeouts() healthcaredicomservice.TimeoutsAttributes {
	return terra.ReferenceAsSingle[healthcaredicomservice.TimeoutsAttributes](hds.ref.Append("timeouts"))
}

type healthcareDicomServiceState struct {
	Id                         string                                        `json:"id"`
	Location                   string                                        `json:"location"`
	Name                       string                                        `json:"name"`
	PublicNetworkAccessEnabled bool                                          `json:"public_network_access_enabled"`
	ServiceUrl                 string                                        `json:"service_url"`
	Tags                       map[string]string                             `json:"tags"`
	WorkspaceId                string                                        `json:"workspace_id"`
	Authentication             []healthcaredicomservice.AuthenticationState  `json:"authentication"`
	PrivateEndpoint            []healthcaredicomservice.PrivateEndpointState `json:"private_endpoint"`
	Identity                   []healthcaredicomservice.IdentityState        `json:"identity"`
	Timeouts                   *healthcaredicomservice.TimeoutsState         `json:"timeouts"`
}
