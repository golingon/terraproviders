// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	databatchapplication "github.com/golingon/terraproviders/azurerm/3.68.0/databatchapplication"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataBatchApplication creates a new instance of [DataBatchApplication].
func NewDataBatchApplication(name string, args DataBatchApplicationArgs) *DataBatchApplication {
	return &DataBatchApplication{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataBatchApplication)(nil)

// DataBatchApplication represents the Terraform data resource azurerm_batch_application.
type DataBatchApplication struct {
	Name string
	Args DataBatchApplicationArgs
}

// DataSource returns the Terraform object type for [DataBatchApplication].
func (ba *DataBatchApplication) DataSource() string {
	return "azurerm_batch_application"
}

// LocalName returns the local name for [DataBatchApplication].
func (ba *DataBatchApplication) LocalName() string {
	return ba.Name
}

// Configuration returns the configuration (args) for [DataBatchApplication].
func (ba *DataBatchApplication) Configuration() interface{} {
	return ba.Args
}

// Attributes returns the attributes for [DataBatchApplication].
func (ba *DataBatchApplication) Attributes() dataBatchApplicationAttributes {
	return dataBatchApplicationAttributes{ref: terra.ReferenceDataResource(ba)}
}

// DataBatchApplicationArgs contains the configurations for azurerm_batch_application.
type DataBatchApplicationArgs struct {
	// AccountName: string, required
	AccountName terra.StringValue `hcl:"account_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *databatchapplication.Timeouts `hcl:"timeouts,block"`
}
type dataBatchApplicationAttributes struct {
	ref terra.Reference
}

// AccountName returns a reference to field account_name of azurerm_batch_application.
func (ba dataBatchApplicationAttributes) AccountName() terra.StringValue {
	return terra.ReferenceAsString(ba.ref.Append("account_name"))
}

// AllowUpdates returns a reference to field allow_updates of azurerm_batch_application.
func (ba dataBatchApplicationAttributes) AllowUpdates() terra.BoolValue {
	return terra.ReferenceAsBool(ba.ref.Append("allow_updates"))
}

// DefaultVersion returns a reference to field default_version of azurerm_batch_application.
func (ba dataBatchApplicationAttributes) DefaultVersion() terra.StringValue {
	return terra.ReferenceAsString(ba.ref.Append("default_version"))
}

// DisplayName returns a reference to field display_name of azurerm_batch_application.
func (ba dataBatchApplicationAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(ba.ref.Append("display_name"))
}

// Id returns a reference to field id of azurerm_batch_application.
func (ba dataBatchApplicationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ba.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_batch_application.
func (ba dataBatchApplicationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ba.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_batch_application.
func (ba dataBatchApplicationAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ba.ref.Append("resource_group_name"))
}

func (ba dataBatchApplicationAttributes) Timeouts() databatchapplication.TimeoutsAttributes {
	return terra.ReferenceAsSingle[databatchapplication.TimeoutsAttributes](ba.ref.Append("timeouts"))
}
