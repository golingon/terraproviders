// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azuread

import (
	"encoding/json"
	"fmt"
	serviceprincipalpassword "github.com/golingon/terraproviders/azuread/2.39.0/serviceprincipalpassword"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewServicePrincipalPassword creates a new instance of [ServicePrincipalPassword].
func NewServicePrincipalPassword(name string, args ServicePrincipalPasswordArgs) *ServicePrincipalPassword {
	return &ServicePrincipalPassword{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ServicePrincipalPassword)(nil)

// ServicePrincipalPassword represents the Terraform resource azuread_service_principal_password.
type ServicePrincipalPassword struct {
	Name      string
	Args      ServicePrincipalPasswordArgs
	state     *servicePrincipalPasswordState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ServicePrincipalPassword].
func (spp *ServicePrincipalPassword) Type() string {
	return "azuread_service_principal_password"
}

// LocalName returns the local name for [ServicePrincipalPassword].
func (spp *ServicePrincipalPassword) LocalName() string {
	return spp.Name
}

// Configuration returns the configuration (args) for [ServicePrincipalPassword].
func (spp *ServicePrincipalPassword) Configuration() interface{} {
	return spp.Args
}

// DependOn is used for other resources to depend on [ServicePrincipalPassword].
func (spp *ServicePrincipalPassword) DependOn() terra.Reference {
	return terra.ReferenceResource(spp)
}

// Dependencies returns the list of resources [ServicePrincipalPassword] depends_on.
func (spp *ServicePrincipalPassword) Dependencies() terra.Dependencies {
	return spp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ServicePrincipalPassword].
func (spp *ServicePrincipalPassword) LifecycleManagement() *terra.Lifecycle {
	return spp.Lifecycle
}

// Attributes returns the attributes for [ServicePrincipalPassword].
func (spp *ServicePrincipalPassword) Attributes() servicePrincipalPasswordAttributes {
	return servicePrincipalPasswordAttributes{ref: terra.ReferenceResource(spp)}
}

// ImportState imports the given attribute values into [ServicePrincipalPassword]'s state.
func (spp *ServicePrincipalPassword) ImportState(av io.Reader) error {
	spp.state = &servicePrincipalPasswordState{}
	if err := json.NewDecoder(av).Decode(spp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", spp.Type(), spp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ServicePrincipalPassword] has state.
func (spp *ServicePrincipalPassword) State() (*servicePrincipalPasswordState, bool) {
	return spp.state, spp.state != nil
}

// StateMust returns the state for [ServicePrincipalPassword]. Panics if the state is nil.
func (spp *ServicePrincipalPassword) StateMust() *servicePrincipalPasswordState {
	if spp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", spp.Type(), spp.LocalName()))
	}
	return spp.state
}

// ServicePrincipalPasswordArgs contains the configurations for azuread_service_principal_password.
type ServicePrincipalPasswordArgs struct {
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// EndDate: string, optional
	EndDate terra.StringValue `hcl:"end_date,attr"`
	// EndDateRelative: string, optional
	EndDateRelative terra.StringValue `hcl:"end_date_relative,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// RotateWhenChanged: map of string, optional
	RotateWhenChanged terra.MapValue[terra.StringValue] `hcl:"rotate_when_changed,attr"`
	// ServicePrincipalId: string, required
	ServicePrincipalId terra.StringValue `hcl:"service_principal_id,attr" validate:"required"`
	// StartDate: string, optional
	StartDate terra.StringValue `hcl:"start_date,attr"`
	// Timeouts: optional
	Timeouts *serviceprincipalpassword.Timeouts `hcl:"timeouts,block"`
}
type servicePrincipalPasswordAttributes struct {
	ref terra.Reference
}

// DisplayName returns a reference to field display_name of azuread_service_principal_password.
func (spp servicePrincipalPasswordAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(spp.ref.Append("display_name"))
}

// EndDate returns a reference to field end_date of azuread_service_principal_password.
func (spp servicePrincipalPasswordAttributes) EndDate() terra.StringValue {
	return terra.ReferenceAsString(spp.ref.Append("end_date"))
}

// EndDateRelative returns a reference to field end_date_relative of azuread_service_principal_password.
func (spp servicePrincipalPasswordAttributes) EndDateRelative() terra.StringValue {
	return terra.ReferenceAsString(spp.ref.Append("end_date_relative"))
}

// Id returns a reference to field id of azuread_service_principal_password.
func (spp servicePrincipalPasswordAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(spp.ref.Append("id"))
}

// KeyId returns a reference to field key_id of azuread_service_principal_password.
func (spp servicePrincipalPasswordAttributes) KeyId() terra.StringValue {
	return terra.ReferenceAsString(spp.ref.Append("key_id"))
}

// RotateWhenChanged returns a reference to field rotate_when_changed of azuread_service_principal_password.
func (spp servicePrincipalPasswordAttributes) RotateWhenChanged() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](spp.ref.Append("rotate_when_changed"))
}

// ServicePrincipalId returns a reference to field service_principal_id of azuread_service_principal_password.
func (spp servicePrincipalPasswordAttributes) ServicePrincipalId() terra.StringValue {
	return terra.ReferenceAsString(spp.ref.Append("service_principal_id"))
}

// StartDate returns a reference to field start_date of azuread_service_principal_password.
func (spp servicePrincipalPasswordAttributes) StartDate() terra.StringValue {
	return terra.ReferenceAsString(spp.ref.Append("start_date"))
}

// Value returns a reference to field value of azuread_service_principal_password.
func (spp servicePrincipalPasswordAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(spp.ref.Append("value"))
}

func (spp servicePrincipalPasswordAttributes) Timeouts() serviceprincipalpassword.TimeoutsAttributes {
	return terra.ReferenceAsSingle[serviceprincipalpassword.TimeoutsAttributes](spp.ref.Append("timeouts"))
}

type servicePrincipalPasswordState struct {
	DisplayName        string                                  `json:"display_name"`
	EndDate            string                                  `json:"end_date"`
	EndDateRelative    string                                  `json:"end_date_relative"`
	Id                 string                                  `json:"id"`
	KeyId              string                                  `json:"key_id"`
	RotateWhenChanged  map[string]string                       `json:"rotate_when_changed"`
	ServicePrincipalId string                                  `json:"service_principal_id"`
	StartDate          string                                  `json:"start_date"`
	Value              string                                  `json:"value"`
	Timeouts           *serviceprincipalpassword.TimeoutsState `json:"timeouts"`
}
