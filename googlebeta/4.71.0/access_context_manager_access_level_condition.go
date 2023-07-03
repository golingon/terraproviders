// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	accesscontextmanageraccesslevelcondition "github.com/golingon/terraproviders/googlebeta/4.71.0/accesscontextmanageraccesslevelcondition"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAccessContextManagerAccessLevelCondition creates a new instance of [AccessContextManagerAccessLevelCondition].
func NewAccessContextManagerAccessLevelCondition(name string, args AccessContextManagerAccessLevelConditionArgs) *AccessContextManagerAccessLevelCondition {
	return &AccessContextManagerAccessLevelCondition{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AccessContextManagerAccessLevelCondition)(nil)

// AccessContextManagerAccessLevelCondition represents the Terraform resource google_access_context_manager_access_level_condition.
type AccessContextManagerAccessLevelCondition struct {
	Name      string
	Args      AccessContextManagerAccessLevelConditionArgs
	state     *accessContextManagerAccessLevelConditionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AccessContextManagerAccessLevelCondition].
func (acmalc *AccessContextManagerAccessLevelCondition) Type() string {
	return "google_access_context_manager_access_level_condition"
}

// LocalName returns the local name for [AccessContextManagerAccessLevelCondition].
func (acmalc *AccessContextManagerAccessLevelCondition) LocalName() string {
	return acmalc.Name
}

// Configuration returns the configuration (args) for [AccessContextManagerAccessLevelCondition].
func (acmalc *AccessContextManagerAccessLevelCondition) Configuration() interface{} {
	return acmalc.Args
}

// DependOn is used for other resources to depend on [AccessContextManagerAccessLevelCondition].
func (acmalc *AccessContextManagerAccessLevelCondition) DependOn() terra.Reference {
	return terra.ReferenceResource(acmalc)
}

// Dependencies returns the list of resources [AccessContextManagerAccessLevelCondition] depends_on.
func (acmalc *AccessContextManagerAccessLevelCondition) Dependencies() terra.Dependencies {
	return acmalc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AccessContextManagerAccessLevelCondition].
func (acmalc *AccessContextManagerAccessLevelCondition) LifecycleManagement() *terra.Lifecycle {
	return acmalc.Lifecycle
}

// Attributes returns the attributes for [AccessContextManagerAccessLevelCondition].
func (acmalc *AccessContextManagerAccessLevelCondition) Attributes() accessContextManagerAccessLevelConditionAttributes {
	return accessContextManagerAccessLevelConditionAttributes{ref: terra.ReferenceResource(acmalc)}
}

// ImportState imports the given attribute values into [AccessContextManagerAccessLevelCondition]'s state.
func (acmalc *AccessContextManagerAccessLevelCondition) ImportState(av io.Reader) error {
	acmalc.state = &accessContextManagerAccessLevelConditionState{}
	if err := json.NewDecoder(av).Decode(acmalc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", acmalc.Type(), acmalc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AccessContextManagerAccessLevelCondition] has state.
func (acmalc *AccessContextManagerAccessLevelCondition) State() (*accessContextManagerAccessLevelConditionState, bool) {
	return acmalc.state, acmalc.state != nil
}

// StateMust returns the state for [AccessContextManagerAccessLevelCondition]. Panics if the state is nil.
func (acmalc *AccessContextManagerAccessLevelCondition) StateMust() *accessContextManagerAccessLevelConditionState {
	if acmalc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", acmalc.Type(), acmalc.LocalName()))
	}
	return acmalc.state
}

// AccessContextManagerAccessLevelConditionArgs contains the configurations for google_access_context_manager_access_level_condition.
type AccessContextManagerAccessLevelConditionArgs struct {
	// AccessLevel: string, required
	AccessLevel terra.StringValue `hcl:"access_level,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IpSubnetworks: list of string, optional
	IpSubnetworks terra.ListValue[terra.StringValue] `hcl:"ip_subnetworks,attr"`
	// Members: list of string, optional
	Members terra.ListValue[terra.StringValue] `hcl:"members,attr"`
	// Negate: bool, optional
	Negate terra.BoolValue `hcl:"negate,attr"`
	// Regions: list of string, optional
	Regions terra.ListValue[terra.StringValue] `hcl:"regions,attr"`
	// RequiredAccessLevels: list of string, optional
	RequiredAccessLevels terra.ListValue[terra.StringValue] `hcl:"required_access_levels,attr"`
	// DevicePolicy: optional
	DevicePolicy *accesscontextmanageraccesslevelcondition.DevicePolicy `hcl:"device_policy,block"`
	// Timeouts: optional
	Timeouts *accesscontextmanageraccesslevelcondition.Timeouts `hcl:"timeouts,block"`
}
type accessContextManagerAccessLevelConditionAttributes struct {
	ref terra.Reference
}

// AccessLevel returns a reference to field access_level of google_access_context_manager_access_level_condition.
func (acmalc accessContextManagerAccessLevelConditionAttributes) AccessLevel() terra.StringValue {
	return terra.ReferenceAsString(acmalc.ref.Append("access_level"))
}

// Id returns a reference to field id of google_access_context_manager_access_level_condition.
func (acmalc accessContextManagerAccessLevelConditionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(acmalc.ref.Append("id"))
}

// IpSubnetworks returns a reference to field ip_subnetworks of google_access_context_manager_access_level_condition.
func (acmalc accessContextManagerAccessLevelConditionAttributes) IpSubnetworks() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](acmalc.ref.Append("ip_subnetworks"))
}

// Members returns a reference to field members of google_access_context_manager_access_level_condition.
func (acmalc accessContextManagerAccessLevelConditionAttributes) Members() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](acmalc.ref.Append("members"))
}

// Negate returns a reference to field negate of google_access_context_manager_access_level_condition.
func (acmalc accessContextManagerAccessLevelConditionAttributes) Negate() terra.BoolValue {
	return terra.ReferenceAsBool(acmalc.ref.Append("negate"))
}

// Regions returns a reference to field regions of google_access_context_manager_access_level_condition.
func (acmalc accessContextManagerAccessLevelConditionAttributes) Regions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](acmalc.ref.Append("regions"))
}

// RequiredAccessLevels returns a reference to field required_access_levels of google_access_context_manager_access_level_condition.
func (acmalc accessContextManagerAccessLevelConditionAttributes) RequiredAccessLevels() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](acmalc.ref.Append("required_access_levels"))
}

func (acmalc accessContextManagerAccessLevelConditionAttributes) DevicePolicy() terra.ListValue[accesscontextmanageraccesslevelcondition.DevicePolicyAttributes] {
	return terra.ReferenceAsList[accesscontextmanageraccesslevelcondition.DevicePolicyAttributes](acmalc.ref.Append("device_policy"))
}

func (acmalc accessContextManagerAccessLevelConditionAttributes) Timeouts() accesscontextmanageraccesslevelcondition.TimeoutsAttributes {
	return terra.ReferenceAsSingle[accesscontextmanageraccesslevelcondition.TimeoutsAttributes](acmalc.ref.Append("timeouts"))
}

type accessContextManagerAccessLevelConditionState struct {
	AccessLevel          string                                                       `json:"access_level"`
	Id                   string                                                       `json:"id"`
	IpSubnetworks        []string                                                     `json:"ip_subnetworks"`
	Members              []string                                                     `json:"members"`
	Negate               bool                                                         `json:"negate"`
	Regions              []string                                                     `json:"regions"`
	RequiredAccessLevels []string                                                     `json:"required_access_levels"`
	DevicePolicy         []accesscontextmanageraccesslevelcondition.DevicePolicyState `json:"device_policy"`
	Timeouts             *accesscontextmanageraccesslevelcondition.TimeoutsState      `json:"timeouts"`
}