// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	dedicatedhost "github.com/golingon/terraproviders/azurerm/3.69.0/dedicatedhost"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDedicatedHost creates a new instance of [DedicatedHost].
func NewDedicatedHost(name string, args DedicatedHostArgs) *DedicatedHost {
	return &DedicatedHost{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DedicatedHost)(nil)

// DedicatedHost represents the Terraform resource azurerm_dedicated_host.
type DedicatedHost struct {
	Name      string
	Args      DedicatedHostArgs
	state     *dedicatedHostState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DedicatedHost].
func (dh *DedicatedHost) Type() string {
	return "azurerm_dedicated_host"
}

// LocalName returns the local name for [DedicatedHost].
func (dh *DedicatedHost) LocalName() string {
	return dh.Name
}

// Configuration returns the configuration (args) for [DedicatedHost].
func (dh *DedicatedHost) Configuration() interface{} {
	return dh.Args
}

// DependOn is used for other resources to depend on [DedicatedHost].
func (dh *DedicatedHost) DependOn() terra.Reference {
	return terra.ReferenceResource(dh)
}

// Dependencies returns the list of resources [DedicatedHost] depends_on.
func (dh *DedicatedHost) Dependencies() terra.Dependencies {
	return dh.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DedicatedHost].
func (dh *DedicatedHost) LifecycleManagement() *terra.Lifecycle {
	return dh.Lifecycle
}

// Attributes returns the attributes for [DedicatedHost].
func (dh *DedicatedHost) Attributes() dedicatedHostAttributes {
	return dedicatedHostAttributes{ref: terra.ReferenceResource(dh)}
}

// ImportState imports the given attribute values into [DedicatedHost]'s state.
func (dh *DedicatedHost) ImportState(av io.Reader) error {
	dh.state = &dedicatedHostState{}
	if err := json.NewDecoder(av).Decode(dh.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dh.Type(), dh.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DedicatedHost] has state.
func (dh *DedicatedHost) State() (*dedicatedHostState, bool) {
	return dh.state, dh.state != nil
}

// StateMust returns the state for [DedicatedHost]. Panics if the state is nil.
func (dh *DedicatedHost) StateMust() *dedicatedHostState {
	if dh.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dh.Type(), dh.LocalName()))
	}
	return dh.state
}

// DedicatedHostArgs contains the configurations for azurerm_dedicated_host.
type DedicatedHostArgs struct {
	// AutoReplaceOnFailure: bool, optional
	AutoReplaceOnFailure terra.BoolValue `hcl:"auto_replace_on_failure,attr"`
	// DedicatedHostGroupId: string, required
	DedicatedHostGroupId terra.StringValue `hcl:"dedicated_host_group_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LicenseType: string, optional
	LicenseType terra.StringValue `hcl:"license_type,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PlatformFaultDomain: number, required
	PlatformFaultDomain terra.NumberValue `hcl:"platform_fault_domain,attr" validate:"required"`
	// SkuName: string, required
	SkuName terra.StringValue `hcl:"sku_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *dedicatedhost.Timeouts `hcl:"timeouts,block"`
}
type dedicatedHostAttributes struct {
	ref terra.Reference
}

// AutoReplaceOnFailure returns a reference to field auto_replace_on_failure of azurerm_dedicated_host.
func (dh dedicatedHostAttributes) AutoReplaceOnFailure() terra.BoolValue {
	return terra.ReferenceAsBool(dh.ref.Append("auto_replace_on_failure"))
}

// DedicatedHostGroupId returns a reference to field dedicated_host_group_id of azurerm_dedicated_host.
func (dh dedicatedHostAttributes) DedicatedHostGroupId() terra.StringValue {
	return terra.ReferenceAsString(dh.ref.Append("dedicated_host_group_id"))
}

// Id returns a reference to field id of azurerm_dedicated_host.
func (dh dedicatedHostAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dh.ref.Append("id"))
}

// LicenseType returns a reference to field license_type of azurerm_dedicated_host.
func (dh dedicatedHostAttributes) LicenseType() terra.StringValue {
	return terra.ReferenceAsString(dh.ref.Append("license_type"))
}

// Location returns a reference to field location of azurerm_dedicated_host.
func (dh dedicatedHostAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dh.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_dedicated_host.
func (dh dedicatedHostAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dh.ref.Append("name"))
}

// PlatformFaultDomain returns a reference to field platform_fault_domain of azurerm_dedicated_host.
func (dh dedicatedHostAttributes) PlatformFaultDomain() terra.NumberValue {
	return terra.ReferenceAsNumber(dh.ref.Append("platform_fault_domain"))
}

// SkuName returns a reference to field sku_name of azurerm_dedicated_host.
func (dh dedicatedHostAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(dh.ref.Append("sku_name"))
}

// Tags returns a reference to field tags of azurerm_dedicated_host.
func (dh dedicatedHostAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dh.ref.Append("tags"))
}

func (dh dedicatedHostAttributes) Timeouts() dedicatedhost.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dedicatedhost.TimeoutsAttributes](dh.ref.Append("timeouts"))
}

type dedicatedHostState struct {
	AutoReplaceOnFailure bool                         `json:"auto_replace_on_failure"`
	DedicatedHostGroupId string                       `json:"dedicated_host_group_id"`
	Id                   string                       `json:"id"`
	LicenseType          string                       `json:"license_type"`
	Location             string                       `json:"location"`
	Name                 string                       `json:"name"`
	PlatformFaultDomain  float64                      `json:"platform_fault_domain"`
	SkuName              string                       `json:"sku_name"`
	Tags                 map[string]string            `json:"tags"`
	Timeouts             *dedicatedhost.TimeoutsState `json:"timeouts"`
}
