// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	labserviceuser "github.com/golingon/terraproviders/azurerm/3.55.0/labserviceuser"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLabServiceUser creates a new instance of [LabServiceUser].
func NewLabServiceUser(name string, args LabServiceUserArgs) *LabServiceUser {
	return &LabServiceUser{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LabServiceUser)(nil)

// LabServiceUser represents the Terraform resource azurerm_lab_service_user.
type LabServiceUser struct {
	Name      string
	Args      LabServiceUserArgs
	state     *labServiceUserState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LabServiceUser].
func (lsu *LabServiceUser) Type() string {
	return "azurerm_lab_service_user"
}

// LocalName returns the local name for [LabServiceUser].
func (lsu *LabServiceUser) LocalName() string {
	return lsu.Name
}

// Configuration returns the configuration (args) for [LabServiceUser].
func (lsu *LabServiceUser) Configuration() interface{} {
	return lsu.Args
}

// DependOn is used for other resources to depend on [LabServiceUser].
func (lsu *LabServiceUser) DependOn() terra.Reference {
	return terra.ReferenceResource(lsu)
}

// Dependencies returns the list of resources [LabServiceUser] depends_on.
func (lsu *LabServiceUser) Dependencies() terra.Dependencies {
	return lsu.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LabServiceUser].
func (lsu *LabServiceUser) LifecycleManagement() *terra.Lifecycle {
	return lsu.Lifecycle
}

// Attributes returns the attributes for [LabServiceUser].
func (lsu *LabServiceUser) Attributes() labServiceUserAttributes {
	return labServiceUserAttributes{ref: terra.ReferenceResource(lsu)}
}

// ImportState imports the given attribute values into [LabServiceUser]'s state.
func (lsu *LabServiceUser) ImportState(av io.Reader) error {
	lsu.state = &labServiceUserState{}
	if err := json.NewDecoder(av).Decode(lsu.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lsu.Type(), lsu.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LabServiceUser] has state.
func (lsu *LabServiceUser) State() (*labServiceUserState, bool) {
	return lsu.state, lsu.state != nil
}

// StateMust returns the state for [LabServiceUser]. Panics if the state is nil.
func (lsu *LabServiceUser) StateMust() *labServiceUserState {
	if lsu.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lsu.Type(), lsu.LocalName()))
	}
	return lsu.state
}

// LabServiceUserArgs contains the configurations for azurerm_lab_service_user.
type LabServiceUserArgs struct {
	// AdditionalUsageQuota: string, optional
	AdditionalUsageQuota terra.StringValue `hcl:"additional_usage_quota,attr"`
	// Email: string, required
	Email terra.StringValue `hcl:"email,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LabId: string, required
	LabId terra.StringValue `hcl:"lab_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *labserviceuser.Timeouts `hcl:"timeouts,block"`
}
type labServiceUserAttributes struct {
	ref terra.Reference
}

// AdditionalUsageQuota returns a reference to field additional_usage_quota of azurerm_lab_service_user.
func (lsu labServiceUserAttributes) AdditionalUsageQuota() terra.StringValue {
	return terra.ReferenceAsString(lsu.ref.Append("additional_usage_quota"))
}

// Email returns a reference to field email of azurerm_lab_service_user.
func (lsu labServiceUserAttributes) Email() terra.StringValue {
	return terra.ReferenceAsString(lsu.ref.Append("email"))
}

// Id returns a reference to field id of azurerm_lab_service_user.
func (lsu labServiceUserAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lsu.ref.Append("id"))
}

// LabId returns a reference to field lab_id of azurerm_lab_service_user.
func (lsu labServiceUserAttributes) LabId() terra.StringValue {
	return terra.ReferenceAsString(lsu.ref.Append("lab_id"))
}

// Name returns a reference to field name of azurerm_lab_service_user.
func (lsu labServiceUserAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lsu.ref.Append("name"))
}

func (lsu labServiceUserAttributes) Timeouts() labserviceuser.TimeoutsAttributes {
	return terra.ReferenceAsSingle[labserviceuser.TimeoutsAttributes](lsu.ref.Append("timeouts"))
}

type labServiceUserState struct {
	AdditionalUsageQuota string                        `json:"additional_usage_quota"`
	Email                string                        `json:"email"`
	Id                   string                        `json:"id"`
	LabId                string                        `json:"lab_id"`
	Name                 string                        `json:"name"`
	Timeouts             *labserviceuser.TimeoutsState `json:"timeouts"`
}
