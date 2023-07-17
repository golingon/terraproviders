// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datasshpublickey "github.com/golingon/terraproviders/azurerm/3.65.0/datasshpublickey"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataSshPublicKey creates a new instance of [DataSshPublicKey].
func NewDataSshPublicKey(name string, args DataSshPublicKeyArgs) *DataSshPublicKey {
	return &DataSshPublicKey{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSshPublicKey)(nil)

// DataSshPublicKey represents the Terraform data resource azurerm_ssh_public_key.
type DataSshPublicKey struct {
	Name string
	Args DataSshPublicKeyArgs
}

// DataSource returns the Terraform object type for [DataSshPublicKey].
func (spk *DataSshPublicKey) DataSource() string {
	return "azurerm_ssh_public_key"
}

// LocalName returns the local name for [DataSshPublicKey].
func (spk *DataSshPublicKey) LocalName() string {
	return spk.Name
}

// Configuration returns the configuration (args) for [DataSshPublicKey].
func (spk *DataSshPublicKey) Configuration() interface{} {
	return spk.Args
}

// Attributes returns the attributes for [DataSshPublicKey].
func (spk *DataSshPublicKey) Attributes() dataSshPublicKeyAttributes {
	return dataSshPublicKeyAttributes{ref: terra.ReferenceDataResource(spk)}
}

// DataSshPublicKeyArgs contains the configurations for azurerm_ssh_public_key.
type DataSshPublicKeyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *datasshpublickey.Timeouts `hcl:"timeouts,block"`
}
type dataSshPublicKeyAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_ssh_public_key.
func (spk dataSshPublicKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(spk.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_ssh_public_key.
func (spk dataSshPublicKeyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(spk.ref.Append("name"))
}

// PublicKey returns a reference to field public_key of azurerm_ssh_public_key.
func (spk dataSshPublicKeyAttributes) PublicKey() terra.StringValue {
	return terra.ReferenceAsString(spk.ref.Append("public_key"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_ssh_public_key.
func (spk dataSshPublicKeyAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(spk.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_ssh_public_key.
func (spk dataSshPublicKeyAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](spk.ref.Append("tags"))
}

func (spk dataSshPublicKeyAttributes) Timeouts() datasshpublickey.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datasshpublickey.TimeoutsAttributes](spk.ref.Append("timeouts"))
}
