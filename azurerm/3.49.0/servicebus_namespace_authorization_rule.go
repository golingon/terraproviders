// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	servicebusnamespaceauthorizationrule "github.com/golingon/terraproviders/azurerm/3.49.0/servicebusnamespaceauthorizationrule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewServicebusNamespaceAuthorizationRule(name string, args ServicebusNamespaceAuthorizationRuleArgs) *ServicebusNamespaceAuthorizationRule {
	return &ServicebusNamespaceAuthorizationRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ServicebusNamespaceAuthorizationRule)(nil)

type ServicebusNamespaceAuthorizationRule struct {
	Name  string
	Args  ServicebusNamespaceAuthorizationRuleArgs
	state *servicebusNamespaceAuthorizationRuleState
}

func (snar *ServicebusNamespaceAuthorizationRule) Type() string {
	return "azurerm_servicebus_namespace_authorization_rule"
}

func (snar *ServicebusNamespaceAuthorizationRule) LocalName() string {
	return snar.Name
}

func (snar *ServicebusNamespaceAuthorizationRule) Configuration() interface{} {
	return snar.Args
}

func (snar *ServicebusNamespaceAuthorizationRule) Attributes() servicebusNamespaceAuthorizationRuleAttributes {
	return servicebusNamespaceAuthorizationRuleAttributes{ref: terra.ReferenceResource(snar)}
}

func (snar *ServicebusNamespaceAuthorizationRule) ImportState(av io.Reader) error {
	snar.state = &servicebusNamespaceAuthorizationRuleState{}
	if err := json.NewDecoder(av).Decode(snar.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", snar.Type(), snar.LocalName(), err)
	}
	return nil
}

func (snar *ServicebusNamespaceAuthorizationRule) State() (*servicebusNamespaceAuthorizationRuleState, bool) {
	return snar.state, snar.state != nil
}

func (snar *ServicebusNamespaceAuthorizationRule) StateMust() *servicebusNamespaceAuthorizationRuleState {
	if snar.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", snar.Type(), snar.LocalName()))
	}
	return snar.state
}

func (snar *ServicebusNamespaceAuthorizationRule) DependOn() terra.Reference {
	return terra.ReferenceResource(snar)
}

type ServicebusNamespaceAuthorizationRuleArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Listen: bool, optional
	Listen terra.BoolValue `hcl:"listen,attr"`
	// Manage: bool, optional
	Manage terra.BoolValue `hcl:"manage,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NamespaceId: string, required
	NamespaceId terra.StringValue `hcl:"namespace_id,attr" validate:"required"`
	// Send: bool, optional
	Send terra.BoolValue `hcl:"send,attr"`
	// Timeouts: optional
	Timeouts *servicebusnamespaceauthorizationrule.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that ServicebusNamespaceAuthorizationRule depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type servicebusNamespaceAuthorizationRuleAttributes struct {
	ref terra.Reference
}

func (snar servicebusNamespaceAuthorizationRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceString(snar.ref.Append("id"))
}

func (snar servicebusNamespaceAuthorizationRuleAttributes) Listen() terra.BoolValue {
	return terra.ReferenceBool(snar.ref.Append("listen"))
}

func (snar servicebusNamespaceAuthorizationRuleAttributes) Manage() terra.BoolValue {
	return terra.ReferenceBool(snar.ref.Append("manage"))
}

func (snar servicebusNamespaceAuthorizationRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceString(snar.ref.Append("name"))
}

func (snar servicebusNamespaceAuthorizationRuleAttributes) NamespaceId() terra.StringValue {
	return terra.ReferenceString(snar.ref.Append("namespace_id"))
}

func (snar servicebusNamespaceAuthorizationRuleAttributes) PrimaryConnectionString() terra.StringValue {
	return terra.ReferenceString(snar.ref.Append("primary_connection_string"))
}

func (snar servicebusNamespaceAuthorizationRuleAttributes) PrimaryConnectionStringAlias() terra.StringValue {
	return terra.ReferenceString(snar.ref.Append("primary_connection_string_alias"))
}

func (snar servicebusNamespaceAuthorizationRuleAttributes) PrimaryKey() terra.StringValue {
	return terra.ReferenceString(snar.ref.Append("primary_key"))
}

func (snar servicebusNamespaceAuthorizationRuleAttributes) SecondaryConnectionString() terra.StringValue {
	return terra.ReferenceString(snar.ref.Append("secondary_connection_string"))
}

func (snar servicebusNamespaceAuthorizationRuleAttributes) SecondaryConnectionStringAlias() terra.StringValue {
	return terra.ReferenceString(snar.ref.Append("secondary_connection_string_alias"))
}

func (snar servicebusNamespaceAuthorizationRuleAttributes) SecondaryKey() terra.StringValue {
	return terra.ReferenceString(snar.ref.Append("secondary_key"))
}

func (snar servicebusNamespaceAuthorizationRuleAttributes) Send() terra.BoolValue {
	return terra.ReferenceBool(snar.ref.Append("send"))
}

func (snar servicebusNamespaceAuthorizationRuleAttributes) Timeouts() servicebusnamespaceauthorizationrule.TimeoutsAttributes {
	return terra.ReferenceSingle[servicebusnamespaceauthorizationrule.TimeoutsAttributes](snar.ref.Append("timeouts"))
}

type servicebusNamespaceAuthorizationRuleState struct {
	Id                             string                                              `json:"id"`
	Listen                         bool                                                `json:"listen"`
	Manage                         bool                                                `json:"manage"`
	Name                           string                                              `json:"name"`
	NamespaceId                    string                                              `json:"namespace_id"`
	PrimaryConnectionString        string                                              `json:"primary_connection_string"`
	PrimaryConnectionStringAlias   string                                              `json:"primary_connection_string_alias"`
	PrimaryKey                     string                                              `json:"primary_key"`
	SecondaryConnectionString      string                                              `json:"secondary_connection_string"`
	SecondaryConnectionStringAlias string                                              `json:"secondary_connection_string_alias"`
	SecondaryKey                   string                                              `json:"secondary_key"`
	Send                           bool                                                `json:"send"`
	Timeouts                       *servicebusnamespaceauthorizationrule.TimeoutsState `json:"timeouts"`
}
