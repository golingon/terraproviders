// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	securitycentercontact "github.com/golingon/terraproviders/azurerm/3.63.0/securitycentercontact"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSecurityCenterContact creates a new instance of [SecurityCenterContact].
func NewSecurityCenterContact(name string, args SecurityCenterContactArgs) *SecurityCenterContact {
	return &SecurityCenterContact{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SecurityCenterContact)(nil)

// SecurityCenterContact represents the Terraform resource azurerm_security_center_contact.
type SecurityCenterContact struct {
	Name      string
	Args      SecurityCenterContactArgs
	state     *securityCenterContactState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SecurityCenterContact].
func (scc *SecurityCenterContact) Type() string {
	return "azurerm_security_center_contact"
}

// LocalName returns the local name for [SecurityCenterContact].
func (scc *SecurityCenterContact) LocalName() string {
	return scc.Name
}

// Configuration returns the configuration (args) for [SecurityCenterContact].
func (scc *SecurityCenterContact) Configuration() interface{} {
	return scc.Args
}

// DependOn is used for other resources to depend on [SecurityCenterContact].
func (scc *SecurityCenterContact) DependOn() terra.Reference {
	return terra.ReferenceResource(scc)
}

// Dependencies returns the list of resources [SecurityCenterContact] depends_on.
func (scc *SecurityCenterContact) Dependencies() terra.Dependencies {
	return scc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SecurityCenterContact].
func (scc *SecurityCenterContact) LifecycleManagement() *terra.Lifecycle {
	return scc.Lifecycle
}

// Attributes returns the attributes for [SecurityCenterContact].
func (scc *SecurityCenterContact) Attributes() securityCenterContactAttributes {
	return securityCenterContactAttributes{ref: terra.ReferenceResource(scc)}
}

// ImportState imports the given attribute values into [SecurityCenterContact]'s state.
func (scc *SecurityCenterContact) ImportState(av io.Reader) error {
	scc.state = &securityCenterContactState{}
	if err := json.NewDecoder(av).Decode(scc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", scc.Type(), scc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SecurityCenterContact] has state.
func (scc *SecurityCenterContact) State() (*securityCenterContactState, bool) {
	return scc.state, scc.state != nil
}

// StateMust returns the state for [SecurityCenterContact]. Panics if the state is nil.
func (scc *SecurityCenterContact) StateMust() *securityCenterContactState {
	if scc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", scc.Type(), scc.LocalName()))
	}
	return scc.state
}

// SecurityCenterContactArgs contains the configurations for azurerm_security_center_contact.
type SecurityCenterContactArgs struct {
	// AlertNotifications: bool, required
	AlertNotifications terra.BoolValue `hcl:"alert_notifications,attr" validate:"required"`
	// AlertsToAdmins: bool, required
	AlertsToAdmins terra.BoolValue `hcl:"alerts_to_admins,attr" validate:"required"`
	// Email: string, required
	Email terra.StringValue `hcl:"email,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Phone: string, optional
	Phone terra.StringValue `hcl:"phone,attr"`
	// Timeouts: optional
	Timeouts *securitycentercontact.Timeouts `hcl:"timeouts,block"`
}
type securityCenterContactAttributes struct {
	ref terra.Reference
}

// AlertNotifications returns a reference to field alert_notifications of azurerm_security_center_contact.
func (scc securityCenterContactAttributes) AlertNotifications() terra.BoolValue {
	return terra.ReferenceAsBool(scc.ref.Append("alert_notifications"))
}

// AlertsToAdmins returns a reference to field alerts_to_admins of azurerm_security_center_contact.
func (scc securityCenterContactAttributes) AlertsToAdmins() terra.BoolValue {
	return terra.ReferenceAsBool(scc.ref.Append("alerts_to_admins"))
}

// Email returns a reference to field email of azurerm_security_center_contact.
func (scc securityCenterContactAttributes) Email() terra.StringValue {
	return terra.ReferenceAsString(scc.ref.Append("email"))
}

// Id returns a reference to field id of azurerm_security_center_contact.
func (scc securityCenterContactAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(scc.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_security_center_contact.
func (scc securityCenterContactAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(scc.ref.Append("name"))
}

// Phone returns a reference to field phone of azurerm_security_center_contact.
func (scc securityCenterContactAttributes) Phone() terra.StringValue {
	return terra.ReferenceAsString(scc.ref.Append("phone"))
}

func (scc securityCenterContactAttributes) Timeouts() securitycentercontact.TimeoutsAttributes {
	return terra.ReferenceAsSingle[securitycentercontact.TimeoutsAttributes](scc.ref.Append("timeouts"))
}

type securityCenterContactState struct {
	AlertNotifications bool                                 `json:"alert_notifications"`
	AlertsToAdmins     bool                                 `json:"alerts_to_admins"`
	Email              string                               `json:"email"`
	Id                 string                               `json:"id"`
	Name               string                               `json:"name"`
	Phone              string                               `json:"phone"`
	Timeouts           *securitycentercontact.TimeoutsState `json:"timeouts"`
}
