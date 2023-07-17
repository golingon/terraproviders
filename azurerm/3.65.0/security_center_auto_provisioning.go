// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	securitycenterautoprovisioning "github.com/golingon/terraproviders/azurerm/3.65.0/securitycenterautoprovisioning"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSecurityCenterAutoProvisioning creates a new instance of [SecurityCenterAutoProvisioning].
func NewSecurityCenterAutoProvisioning(name string, args SecurityCenterAutoProvisioningArgs) *SecurityCenterAutoProvisioning {
	return &SecurityCenterAutoProvisioning{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SecurityCenterAutoProvisioning)(nil)

// SecurityCenterAutoProvisioning represents the Terraform resource azurerm_security_center_auto_provisioning.
type SecurityCenterAutoProvisioning struct {
	Name      string
	Args      SecurityCenterAutoProvisioningArgs
	state     *securityCenterAutoProvisioningState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SecurityCenterAutoProvisioning].
func (scap *SecurityCenterAutoProvisioning) Type() string {
	return "azurerm_security_center_auto_provisioning"
}

// LocalName returns the local name for [SecurityCenterAutoProvisioning].
func (scap *SecurityCenterAutoProvisioning) LocalName() string {
	return scap.Name
}

// Configuration returns the configuration (args) for [SecurityCenterAutoProvisioning].
func (scap *SecurityCenterAutoProvisioning) Configuration() interface{} {
	return scap.Args
}

// DependOn is used for other resources to depend on [SecurityCenterAutoProvisioning].
func (scap *SecurityCenterAutoProvisioning) DependOn() terra.Reference {
	return terra.ReferenceResource(scap)
}

// Dependencies returns the list of resources [SecurityCenterAutoProvisioning] depends_on.
func (scap *SecurityCenterAutoProvisioning) Dependencies() terra.Dependencies {
	return scap.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SecurityCenterAutoProvisioning].
func (scap *SecurityCenterAutoProvisioning) LifecycleManagement() *terra.Lifecycle {
	return scap.Lifecycle
}

// Attributes returns the attributes for [SecurityCenterAutoProvisioning].
func (scap *SecurityCenterAutoProvisioning) Attributes() securityCenterAutoProvisioningAttributes {
	return securityCenterAutoProvisioningAttributes{ref: terra.ReferenceResource(scap)}
}

// ImportState imports the given attribute values into [SecurityCenterAutoProvisioning]'s state.
func (scap *SecurityCenterAutoProvisioning) ImportState(av io.Reader) error {
	scap.state = &securityCenterAutoProvisioningState{}
	if err := json.NewDecoder(av).Decode(scap.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", scap.Type(), scap.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SecurityCenterAutoProvisioning] has state.
func (scap *SecurityCenterAutoProvisioning) State() (*securityCenterAutoProvisioningState, bool) {
	return scap.state, scap.state != nil
}

// StateMust returns the state for [SecurityCenterAutoProvisioning]. Panics if the state is nil.
func (scap *SecurityCenterAutoProvisioning) StateMust() *securityCenterAutoProvisioningState {
	if scap.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", scap.Type(), scap.LocalName()))
	}
	return scap.state
}

// SecurityCenterAutoProvisioningArgs contains the configurations for azurerm_security_center_auto_provisioning.
type SecurityCenterAutoProvisioningArgs struct {
	// AutoProvision: string, required
	AutoProvision terra.StringValue `hcl:"auto_provision,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Timeouts: optional
	Timeouts *securitycenterautoprovisioning.Timeouts `hcl:"timeouts,block"`
}
type securityCenterAutoProvisioningAttributes struct {
	ref terra.Reference
}

// AutoProvision returns a reference to field auto_provision of azurerm_security_center_auto_provisioning.
func (scap securityCenterAutoProvisioningAttributes) AutoProvision() terra.StringValue {
	return terra.ReferenceAsString(scap.ref.Append("auto_provision"))
}

// Id returns a reference to field id of azurerm_security_center_auto_provisioning.
func (scap securityCenterAutoProvisioningAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(scap.ref.Append("id"))
}

func (scap securityCenterAutoProvisioningAttributes) Timeouts() securitycenterautoprovisioning.TimeoutsAttributes {
	return terra.ReferenceAsSingle[securitycenterautoprovisioning.TimeoutsAttributes](scap.ref.Append("timeouts"))
}

type securityCenterAutoProvisioningState struct {
	AutoProvision string                                        `json:"auto_provision"`
	Id            string                                        `json:"id"`
	Timeouts      *securitycenterautoprovisioning.TimeoutsState `json:"timeouts"`
}
