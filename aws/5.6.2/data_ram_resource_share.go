// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataramresourceshare "github.com/golingon/terraproviders/aws/5.6.2/dataramresourceshare"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataRamResourceShare creates a new instance of [DataRamResourceShare].
func NewDataRamResourceShare(name string, args DataRamResourceShareArgs) *DataRamResourceShare {
	return &DataRamResourceShare{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataRamResourceShare)(nil)

// DataRamResourceShare represents the Terraform data resource aws_ram_resource_share.
type DataRamResourceShare struct {
	Name string
	Args DataRamResourceShareArgs
}

// DataSource returns the Terraform object type for [DataRamResourceShare].
func (rrs *DataRamResourceShare) DataSource() string {
	return "aws_ram_resource_share"
}

// LocalName returns the local name for [DataRamResourceShare].
func (rrs *DataRamResourceShare) LocalName() string {
	return rrs.Name
}

// Configuration returns the configuration (args) for [DataRamResourceShare].
func (rrs *DataRamResourceShare) Configuration() interface{} {
	return rrs.Args
}

// Attributes returns the attributes for [DataRamResourceShare].
func (rrs *DataRamResourceShare) Attributes() dataRamResourceShareAttributes {
	return dataRamResourceShareAttributes{ref: terra.ReferenceDataResource(rrs)}
}

// DataRamResourceShareArgs contains the configurations for aws_ram_resource_share.
type DataRamResourceShareArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceOwner: string, required
	ResourceOwner terra.StringValue `hcl:"resource_owner,attr" validate:"required"`
	// ResourceShareStatus: string, optional
	ResourceShareStatus terra.StringValue `hcl:"resource_share_status,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Filter: min=0
	Filter []dataramresourceshare.Filter `hcl:"filter,block" validate:"min=0"`
}
type dataRamResourceShareAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_ram_resource_share.
func (rrs dataRamResourceShareAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(rrs.ref.Append("arn"))
}

// Id returns a reference to field id of aws_ram_resource_share.
func (rrs dataRamResourceShareAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rrs.ref.Append("id"))
}

// Name returns a reference to field name of aws_ram_resource_share.
func (rrs dataRamResourceShareAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rrs.ref.Append("name"))
}

// OwningAccountId returns a reference to field owning_account_id of aws_ram_resource_share.
func (rrs dataRamResourceShareAttributes) OwningAccountId() terra.StringValue {
	return terra.ReferenceAsString(rrs.ref.Append("owning_account_id"))
}

// ResourceOwner returns a reference to field resource_owner of aws_ram_resource_share.
func (rrs dataRamResourceShareAttributes) ResourceOwner() terra.StringValue {
	return terra.ReferenceAsString(rrs.ref.Append("resource_owner"))
}

// ResourceShareStatus returns a reference to field resource_share_status of aws_ram_resource_share.
func (rrs dataRamResourceShareAttributes) ResourceShareStatus() terra.StringValue {
	return terra.ReferenceAsString(rrs.ref.Append("resource_share_status"))
}

// Status returns a reference to field status of aws_ram_resource_share.
func (rrs dataRamResourceShareAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(rrs.ref.Append("status"))
}

// Tags returns a reference to field tags of aws_ram_resource_share.
func (rrs dataRamResourceShareAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rrs.ref.Append("tags"))
}

func (rrs dataRamResourceShareAttributes) Filter() terra.SetValue[dataramresourceshare.FilterAttributes] {
	return terra.ReferenceAsSet[dataramresourceshare.FilterAttributes](rrs.ref.Append("filter"))
}
