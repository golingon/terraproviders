// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	securitycentercontact "github.com/golingon/terraproviders/azurerm/3.49.0/securitycentercontact"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewSecurityCenterContact(name string, args SecurityCenterContactArgs) *SecurityCenterContact {
	return &SecurityCenterContact{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SecurityCenterContact)(nil)

type SecurityCenterContact struct {
	Name  string
	Args  SecurityCenterContactArgs
	state *securityCenterContactState
}

func (scc *SecurityCenterContact) Type() string {
	return "azurerm_security_center_contact"
}

func (scc *SecurityCenterContact) LocalName() string {
	return scc.Name
}

func (scc *SecurityCenterContact) Configuration() interface{} {
	return scc.Args
}

func (scc *SecurityCenterContact) Attributes() securityCenterContactAttributes {
	return securityCenterContactAttributes{ref: terra.ReferenceResource(scc)}
}

func (scc *SecurityCenterContact) ImportState(av io.Reader) error {
	scc.state = &securityCenterContactState{}
	if err := json.NewDecoder(av).Decode(scc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", scc.Type(), scc.LocalName(), err)
	}
	return nil
}

func (scc *SecurityCenterContact) State() (*securityCenterContactState, bool) {
	return scc.state, scc.state != nil
}

func (scc *SecurityCenterContact) StateMust() *securityCenterContactState {
	if scc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", scc.Type(), scc.LocalName()))
	}
	return scc.state
}

func (scc *SecurityCenterContact) DependOn() terra.Reference {
	return terra.ReferenceResource(scc)
}

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
	// DependsOn contains resources that SecurityCenterContact depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type securityCenterContactAttributes struct {
	ref terra.Reference
}

func (scc securityCenterContactAttributes) AlertNotifications() terra.BoolValue {
	return terra.ReferenceBool(scc.ref.Append("alert_notifications"))
}

func (scc securityCenterContactAttributes) AlertsToAdmins() terra.BoolValue {
	return terra.ReferenceBool(scc.ref.Append("alerts_to_admins"))
}

func (scc securityCenterContactAttributes) Email() terra.StringValue {
	return terra.ReferenceString(scc.ref.Append("email"))
}

func (scc securityCenterContactAttributes) Id() terra.StringValue {
	return terra.ReferenceString(scc.ref.Append("id"))
}

func (scc securityCenterContactAttributes) Name() terra.StringValue {
	return terra.ReferenceString(scc.ref.Append("name"))
}

func (scc securityCenterContactAttributes) Phone() terra.StringValue {
	return terra.ReferenceString(scc.ref.Append("phone"))
}

func (scc securityCenterContactAttributes) Timeouts() securitycentercontact.TimeoutsAttributes {
	return terra.ReferenceSingle[securitycentercontact.TimeoutsAttributes](scc.ref.Append("timeouts"))
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
