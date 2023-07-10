// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	privateendpointapplicationsecuritygroupassociation "github.com/golingon/terraproviders/azurerm/3.64.0/privateendpointapplicationsecuritygroupassociation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPrivateEndpointApplicationSecurityGroupAssociation creates a new instance of [PrivateEndpointApplicationSecurityGroupAssociation].
func NewPrivateEndpointApplicationSecurityGroupAssociation(name string, args PrivateEndpointApplicationSecurityGroupAssociationArgs) *PrivateEndpointApplicationSecurityGroupAssociation {
	return &PrivateEndpointApplicationSecurityGroupAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PrivateEndpointApplicationSecurityGroupAssociation)(nil)

// PrivateEndpointApplicationSecurityGroupAssociation represents the Terraform resource azurerm_private_endpoint_application_security_group_association.
type PrivateEndpointApplicationSecurityGroupAssociation struct {
	Name      string
	Args      PrivateEndpointApplicationSecurityGroupAssociationArgs
	state     *privateEndpointApplicationSecurityGroupAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PrivateEndpointApplicationSecurityGroupAssociation].
func (peasga *PrivateEndpointApplicationSecurityGroupAssociation) Type() string {
	return "azurerm_private_endpoint_application_security_group_association"
}

// LocalName returns the local name for [PrivateEndpointApplicationSecurityGroupAssociation].
func (peasga *PrivateEndpointApplicationSecurityGroupAssociation) LocalName() string {
	return peasga.Name
}

// Configuration returns the configuration (args) for [PrivateEndpointApplicationSecurityGroupAssociation].
func (peasga *PrivateEndpointApplicationSecurityGroupAssociation) Configuration() interface{} {
	return peasga.Args
}

// DependOn is used for other resources to depend on [PrivateEndpointApplicationSecurityGroupAssociation].
func (peasga *PrivateEndpointApplicationSecurityGroupAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(peasga)
}

// Dependencies returns the list of resources [PrivateEndpointApplicationSecurityGroupAssociation] depends_on.
func (peasga *PrivateEndpointApplicationSecurityGroupAssociation) Dependencies() terra.Dependencies {
	return peasga.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PrivateEndpointApplicationSecurityGroupAssociation].
func (peasga *PrivateEndpointApplicationSecurityGroupAssociation) LifecycleManagement() *terra.Lifecycle {
	return peasga.Lifecycle
}

// Attributes returns the attributes for [PrivateEndpointApplicationSecurityGroupAssociation].
func (peasga *PrivateEndpointApplicationSecurityGroupAssociation) Attributes() privateEndpointApplicationSecurityGroupAssociationAttributes {
	return privateEndpointApplicationSecurityGroupAssociationAttributes{ref: terra.ReferenceResource(peasga)}
}

// ImportState imports the given attribute values into [PrivateEndpointApplicationSecurityGroupAssociation]'s state.
func (peasga *PrivateEndpointApplicationSecurityGroupAssociation) ImportState(av io.Reader) error {
	peasga.state = &privateEndpointApplicationSecurityGroupAssociationState{}
	if err := json.NewDecoder(av).Decode(peasga.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", peasga.Type(), peasga.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PrivateEndpointApplicationSecurityGroupAssociation] has state.
func (peasga *PrivateEndpointApplicationSecurityGroupAssociation) State() (*privateEndpointApplicationSecurityGroupAssociationState, bool) {
	return peasga.state, peasga.state != nil
}

// StateMust returns the state for [PrivateEndpointApplicationSecurityGroupAssociation]. Panics if the state is nil.
func (peasga *PrivateEndpointApplicationSecurityGroupAssociation) StateMust() *privateEndpointApplicationSecurityGroupAssociationState {
	if peasga.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", peasga.Type(), peasga.LocalName()))
	}
	return peasga.state
}

// PrivateEndpointApplicationSecurityGroupAssociationArgs contains the configurations for azurerm_private_endpoint_application_security_group_association.
type PrivateEndpointApplicationSecurityGroupAssociationArgs struct {
	// ApplicationSecurityGroupId: string, required
	ApplicationSecurityGroupId terra.StringValue `hcl:"application_security_group_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PrivateEndpointId: string, required
	PrivateEndpointId terra.StringValue `hcl:"private_endpoint_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *privateendpointapplicationsecuritygroupassociation.Timeouts `hcl:"timeouts,block"`
}
type privateEndpointApplicationSecurityGroupAssociationAttributes struct {
	ref terra.Reference
}

// ApplicationSecurityGroupId returns a reference to field application_security_group_id of azurerm_private_endpoint_application_security_group_association.
func (peasga privateEndpointApplicationSecurityGroupAssociationAttributes) ApplicationSecurityGroupId() terra.StringValue {
	return terra.ReferenceAsString(peasga.ref.Append("application_security_group_id"))
}

// Id returns a reference to field id of azurerm_private_endpoint_application_security_group_association.
func (peasga privateEndpointApplicationSecurityGroupAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(peasga.ref.Append("id"))
}

// PrivateEndpointId returns a reference to field private_endpoint_id of azurerm_private_endpoint_application_security_group_association.
func (peasga privateEndpointApplicationSecurityGroupAssociationAttributes) PrivateEndpointId() terra.StringValue {
	return terra.ReferenceAsString(peasga.ref.Append("private_endpoint_id"))
}

func (peasga privateEndpointApplicationSecurityGroupAssociationAttributes) Timeouts() privateendpointapplicationsecuritygroupassociation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[privateendpointapplicationsecuritygroupassociation.TimeoutsAttributes](peasga.ref.Append("timeouts"))
}

type privateEndpointApplicationSecurityGroupAssociationState struct {
	ApplicationSecurityGroupId string                                                            `json:"application_security_group_id"`
	Id                         string                                                            `json:"id"`
	PrivateEndpointId          string                                                            `json:"private_endpoint_id"`
	Timeouts                   *privateendpointapplicationsecuritygroupassociation.TimeoutsState `json:"timeouts"`
}
