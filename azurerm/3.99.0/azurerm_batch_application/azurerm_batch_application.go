// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_batch_application

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource azurerm_batch_application.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermBatchApplicationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (aba *Resource) Type() string {
	return "azurerm_batch_application"
}

// LocalName returns the local name for [Resource].
func (aba *Resource) LocalName() string {
	return aba.Name
}

// Configuration returns the configuration (args) for [Resource].
func (aba *Resource) Configuration() interface{} {
	return aba.Args
}

// DependOn is used for other resources to depend on [Resource].
func (aba *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(aba)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (aba *Resource) Dependencies() terra.Dependencies {
	return aba.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (aba *Resource) LifecycleManagement() *terra.Lifecycle {
	return aba.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (aba *Resource) Attributes() azurermBatchApplicationAttributes {
	return azurermBatchApplicationAttributes{ref: terra.ReferenceResource(aba)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (aba *Resource) ImportState(state io.Reader) error {
	aba.state = &azurermBatchApplicationState{}
	if err := json.NewDecoder(state).Decode(aba.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aba.Type(), aba.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (aba *Resource) State() (*azurermBatchApplicationState, bool) {
	return aba.state, aba.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (aba *Resource) StateMust() *azurermBatchApplicationState {
	if aba.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aba.Type(), aba.LocalName()))
	}
	return aba.state
}

// Args contains the configurations for azurerm_batch_application.
type Args struct {
	// AccountName: string, required
	AccountName terra.StringValue `hcl:"account_name,attr" validate:"required"`
	// AllowUpdates: bool, optional
	AllowUpdates terra.BoolValue `hcl:"allow_updates,attr"`
	// DefaultVersion: string, optional
	DefaultVersion terra.StringValue `hcl:"default_version,attr"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermBatchApplicationAttributes struct {
	ref terra.Reference
}

// AccountName returns a reference to field account_name of azurerm_batch_application.
func (aba azurermBatchApplicationAttributes) AccountName() terra.StringValue {
	return terra.ReferenceAsString(aba.ref.Append("account_name"))
}

// AllowUpdates returns a reference to field allow_updates of azurerm_batch_application.
func (aba azurermBatchApplicationAttributes) AllowUpdates() terra.BoolValue {
	return terra.ReferenceAsBool(aba.ref.Append("allow_updates"))
}

// DefaultVersion returns a reference to field default_version of azurerm_batch_application.
func (aba azurermBatchApplicationAttributes) DefaultVersion() terra.StringValue {
	return terra.ReferenceAsString(aba.ref.Append("default_version"))
}

// DisplayName returns a reference to field display_name of azurerm_batch_application.
func (aba azurermBatchApplicationAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(aba.ref.Append("display_name"))
}

// Id returns a reference to field id of azurerm_batch_application.
func (aba azurermBatchApplicationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aba.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_batch_application.
func (aba azurermBatchApplicationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aba.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_batch_application.
func (aba azurermBatchApplicationAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(aba.ref.Append("resource_group_name"))
}

func (aba azurermBatchApplicationAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](aba.ref.Append("timeouts"))
}

type azurermBatchApplicationState struct {
	AccountName       string         `json:"account_name"`
	AllowUpdates      bool           `json:"allow_updates"`
	DefaultVersion    string         `json:"default_version"`
	DisplayName       string         `json:"display_name"`
	Id                string         `json:"id"`
	Name              string         `json:"name"`
	ResourceGroupName string         `json:"resource_group_name"`
	Timeouts          *TimeoutsState `json:"timeouts"`
}
