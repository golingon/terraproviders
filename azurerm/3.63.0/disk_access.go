// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	diskaccess "github.com/golingon/terraproviders/azurerm/3.63.0/diskaccess"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDiskAccess creates a new instance of [DiskAccess].
func NewDiskAccess(name string, args DiskAccessArgs) *DiskAccess {
	return &DiskAccess{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DiskAccess)(nil)

// DiskAccess represents the Terraform resource azurerm_disk_access.
type DiskAccess struct {
	Name      string
	Args      DiskAccessArgs
	state     *diskAccessState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DiskAccess].
func (da *DiskAccess) Type() string {
	return "azurerm_disk_access"
}

// LocalName returns the local name for [DiskAccess].
func (da *DiskAccess) LocalName() string {
	return da.Name
}

// Configuration returns the configuration (args) for [DiskAccess].
func (da *DiskAccess) Configuration() interface{} {
	return da.Args
}

// DependOn is used for other resources to depend on [DiskAccess].
func (da *DiskAccess) DependOn() terra.Reference {
	return terra.ReferenceResource(da)
}

// Dependencies returns the list of resources [DiskAccess] depends_on.
func (da *DiskAccess) Dependencies() terra.Dependencies {
	return da.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DiskAccess].
func (da *DiskAccess) LifecycleManagement() *terra.Lifecycle {
	return da.Lifecycle
}

// Attributes returns the attributes for [DiskAccess].
func (da *DiskAccess) Attributes() diskAccessAttributes {
	return diskAccessAttributes{ref: terra.ReferenceResource(da)}
}

// ImportState imports the given attribute values into [DiskAccess]'s state.
func (da *DiskAccess) ImportState(av io.Reader) error {
	da.state = &diskAccessState{}
	if err := json.NewDecoder(av).Decode(da.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", da.Type(), da.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DiskAccess] has state.
func (da *DiskAccess) State() (*diskAccessState, bool) {
	return da.state, da.state != nil
}

// StateMust returns the state for [DiskAccess]. Panics if the state is nil.
func (da *DiskAccess) StateMust() *diskAccessState {
	if da.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", da.Type(), da.LocalName()))
	}
	return da.state
}

// DiskAccessArgs contains the configurations for azurerm_disk_access.
type DiskAccessArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *diskaccess.Timeouts `hcl:"timeouts,block"`
}
type diskAccessAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_disk_access.
func (da diskAccessAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(da.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_disk_access.
func (da diskAccessAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(da.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_disk_access.
func (da diskAccessAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(da.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_disk_access.
func (da diskAccessAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(da.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_disk_access.
func (da diskAccessAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](da.ref.Append("tags"))
}

func (da diskAccessAttributes) Timeouts() diskaccess.TimeoutsAttributes {
	return terra.ReferenceAsSingle[diskaccess.TimeoutsAttributes](da.ref.Append("timeouts"))
}

type diskAccessState struct {
	Id                string                    `json:"id"`
	Location          string                    `json:"location"`
	Name              string                    `json:"name"`
	ResourceGroupName string                    `json:"resource_group_name"`
	Tags              map[string]string         `json:"tags"`
	Timeouts          *diskaccess.TimeoutsState `json:"timeouts"`
}
