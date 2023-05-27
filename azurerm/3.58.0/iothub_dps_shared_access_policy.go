// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	iothubdpssharedaccesspolicy "github.com/golingon/terraproviders/azurerm/3.58.0/iothubdpssharedaccesspolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIothubDpsSharedAccessPolicy creates a new instance of [IothubDpsSharedAccessPolicy].
func NewIothubDpsSharedAccessPolicy(name string, args IothubDpsSharedAccessPolicyArgs) *IothubDpsSharedAccessPolicy {
	return &IothubDpsSharedAccessPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IothubDpsSharedAccessPolicy)(nil)

// IothubDpsSharedAccessPolicy represents the Terraform resource azurerm_iothub_dps_shared_access_policy.
type IothubDpsSharedAccessPolicy struct {
	Name      string
	Args      IothubDpsSharedAccessPolicyArgs
	state     *iothubDpsSharedAccessPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IothubDpsSharedAccessPolicy].
func (idsap *IothubDpsSharedAccessPolicy) Type() string {
	return "azurerm_iothub_dps_shared_access_policy"
}

// LocalName returns the local name for [IothubDpsSharedAccessPolicy].
func (idsap *IothubDpsSharedAccessPolicy) LocalName() string {
	return idsap.Name
}

// Configuration returns the configuration (args) for [IothubDpsSharedAccessPolicy].
func (idsap *IothubDpsSharedAccessPolicy) Configuration() interface{} {
	return idsap.Args
}

// DependOn is used for other resources to depend on [IothubDpsSharedAccessPolicy].
func (idsap *IothubDpsSharedAccessPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(idsap)
}

// Dependencies returns the list of resources [IothubDpsSharedAccessPolicy] depends_on.
func (idsap *IothubDpsSharedAccessPolicy) Dependencies() terra.Dependencies {
	return idsap.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IothubDpsSharedAccessPolicy].
func (idsap *IothubDpsSharedAccessPolicy) LifecycleManagement() *terra.Lifecycle {
	return idsap.Lifecycle
}

// Attributes returns the attributes for [IothubDpsSharedAccessPolicy].
func (idsap *IothubDpsSharedAccessPolicy) Attributes() iothubDpsSharedAccessPolicyAttributes {
	return iothubDpsSharedAccessPolicyAttributes{ref: terra.ReferenceResource(idsap)}
}

// ImportState imports the given attribute values into [IothubDpsSharedAccessPolicy]'s state.
func (idsap *IothubDpsSharedAccessPolicy) ImportState(av io.Reader) error {
	idsap.state = &iothubDpsSharedAccessPolicyState{}
	if err := json.NewDecoder(av).Decode(idsap.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", idsap.Type(), idsap.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IothubDpsSharedAccessPolicy] has state.
func (idsap *IothubDpsSharedAccessPolicy) State() (*iothubDpsSharedAccessPolicyState, bool) {
	return idsap.state, idsap.state != nil
}

// StateMust returns the state for [IothubDpsSharedAccessPolicy]. Panics if the state is nil.
func (idsap *IothubDpsSharedAccessPolicy) StateMust() *iothubDpsSharedAccessPolicyState {
	if idsap.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", idsap.Type(), idsap.LocalName()))
	}
	return idsap.state
}

// IothubDpsSharedAccessPolicyArgs contains the configurations for azurerm_iothub_dps_shared_access_policy.
type IothubDpsSharedAccessPolicyArgs struct {
	// EnrollmentRead: bool, optional
	EnrollmentRead terra.BoolValue `hcl:"enrollment_read,attr"`
	// EnrollmentWrite: bool, optional
	EnrollmentWrite terra.BoolValue `hcl:"enrollment_write,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IothubDpsName: string, required
	IothubDpsName terra.StringValue `hcl:"iothub_dps_name,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RegistrationRead: bool, optional
	RegistrationRead terra.BoolValue `hcl:"registration_read,attr"`
	// RegistrationWrite: bool, optional
	RegistrationWrite terra.BoolValue `hcl:"registration_write,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ServiceConfig: bool, optional
	ServiceConfig terra.BoolValue `hcl:"service_config,attr"`
	// Timeouts: optional
	Timeouts *iothubdpssharedaccesspolicy.Timeouts `hcl:"timeouts,block"`
}
type iothubDpsSharedAccessPolicyAttributes struct {
	ref terra.Reference
}

// EnrollmentRead returns a reference to field enrollment_read of azurerm_iothub_dps_shared_access_policy.
func (idsap iothubDpsSharedAccessPolicyAttributes) EnrollmentRead() terra.BoolValue {
	return terra.ReferenceAsBool(idsap.ref.Append("enrollment_read"))
}

// EnrollmentWrite returns a reference to field enrollment_write of azurerm_iothub_dps_shared_access_policy.
func (idsap iothubDpsSharedAccessPolicyAttributes) EnrollmentWrite() terra.BoolValue {
	return terra.ReferenceAsBool(idsap.ref.Append("enrollment_write"))
}

// Id returns a reference to field id of azurerm_iothub_dps_shared_access_policy.
func (idsap iothubDpsSharedAccessPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(idsap.ref.Append("id"))
}

// IothubDpsName returns a reference to field iothub_dps_name of azurerm_iothub_dps_shared_access_policy.
func (idsap iothubDpsSharedAccessPolicyAttributes) IothubDpsName() terra.StringValue {
	return terra.ReferenceAsString(idsap.ref.Append("iothub_dps_name"))
}

// Name returns a reference to field name of azurerm_iothub_dps_shared_access_policy.
func (idsap iothubDpsSharedAccessPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(idsap.ref.Append("name"))
}

// PrimaryConnectionString returns a reference to field primary_connection_string of azurerm_iothub_dps_shared_access_policy.
func (idsap iothubDpsSharedAccessPolicyAttributes) PrimaryConnectionString() terra.StringValue {
	return terra.ReferenceAsString(idsap.ref.Append("primary_connection_string"))
}

// PrimaryKey returns a reference to field primary_key of azurerm_iothub_dps_shared_access_policy.
func (idsap iothubDpsSharedAccessPolicyAttributes) PrimaryKey() terra.StringValue {
	return terra.ReferenceAsString(idsap.ref.Append("primary_key"))
}

// RegistrationRead returns a reference to field registration_read of azurerm_iothub_dps_shared_access_policy.
func (idsap iothubDpsSharedAccessPolicyAttributes) RegistrationRead() terra.BoolValue {
	return terra.ReferenceAsBool(idsap.ref.Append("registration_read"))
}

// RegistrationWrite returns a reference to field registration_write of azurerm_iothub_dps_shared_access_policy.
func (idsap iothubDpsSharedAccessPolicyAttributes) RegistrationWrite() terra.BoolValue {
	return terra.ReferenceAsBool(idsap.ref.Append("registration_write"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_iothub_dps_shared_access_policy.
func (idsap iothubDpsSharedAccessPolicyAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(idsap.ref.Append("resource_group_name"))
}

// SecondaryConnectionString returns a reference to field secondary_connection_string of azurerm_iothub_dps_shared_access_policy.
func (idsap iothubDpsSharedAccessPolicyAttributes) SecondaryConnectionString() terra.StringValue {
	return terra.ReferenceAsString(idsap.ref.Append("secondary_connection_string"))
}

// SecondaryKey returns a reference to field secondary_key of azurerm_iothub_dps_shared_access_policy.
func (idsap iothubDpsSharedAccessPolicyAttributes) SecondaryKey() terra.StringValue {
	return terra.ReferenceAsString(idsap.ref.Append("secondary_key"))
}

// ServiceConfig returns a reference to field service_config of azurerm_iothub_dps_shared_access_policy.
func (idsap iothubDpsSharedAccessPolicyAttributes) ServiceConfig() terra.BoolValue {
	return terra.ReferenceAsBool(idsap.ref.Append("service_config"))
}

func (idsap iothubDpsSharedAccessPolicyAttributes) Timeouts() iothubdpssharedaccesspolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[iothubdpssharedaccesspolicy.TimeoutsAttributes](idsap.ref.Append("timeouts"))
}

type iothubDpsSharedAccessPolicyState struct {
	EnrollmentRead            bool                                       `json:"enrollment_read"`
	EnrollmentWrite           bool                                       `json:"enrollment_write"`
	Id                        string                                     `json:"id"`
	IothubDpsName             string                                     `json:"iothub_dps_name"`
	Name                      string                                     `json:"name"`
	PrimaryConnectionString   string                                     `json:"primary_connection_string"`
	PrimaryKey                string                                     `json:"primary_key"`
	RegistrationRead          bool                                       `json:"registration_read"`
	RegistrationWrite         bool                                       `json:"registration_write"`
	ResourceGroupName         string                                     `json:"resource_group_name"`
	SecondaryConnectionString string                                     `json:"secondary_connection_string"`
	SecondaryKey              string                                     `json:"secondary_key"`
	ServiceConfig             bool                                       `json:"service_config"`
	Timeouts                  *iothubdpssharedaccesspolicy.TimeoutsState `json:"timeouts"`
}
