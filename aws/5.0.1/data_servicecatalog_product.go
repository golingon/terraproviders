// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataservicecatalogproduct "github.com/golingon/terraproviders/aws/5.0.1/dataservicecatalogproduct"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataServicecatalogProduct creates a new instance of [DataServicecatalogProduct].
func NewDataServicecatalogProduct(name string, args DataServicecatalogProductArgs) *DataServicecatalogProduct {
	return &DataServicecatalogProduct{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataServicecatalogProduct)(nil)

// DataServicecatalogProduct represents the Terraform data resource aws_servicecatalog_product.
type DataServicecatalogProduct struct {
	Name string
	Args DataServicecatalogProductArgs
}

// DataSource returns the Terraform object type for [DataServicecatalogProduct].
func (sp *DataServicecatalogProduct) DataSource() string {
	return "aws_servicecatalog_product"
}

// LocalName returns the local name for [DataServicecatalogProduct].
func (sp *DataServicecatalogProduct) LocalName() string {
	return sp.Name
}

// Configuration returns the configuration (args) for [DataServicecatalogProduct].
func (sp *DataServicecatalogProduct) Configuration() interface{} {
	return sp.Args
}

// Attributes returns the attributes for [DataServicecatalogProduct].
func (sp *DataServicecatalogProduct) Attributes() dataServicecatalogProductAttributes {
	return dataServicecatalogProductAttributes{ref: terra.ReferenceDataResource(sp)}
}

// DataServicecatalogProductArgs contains the configurations for aws_servicecatalog_product.
type DataServicecatalogProductArgs struct {
	// AcceptLanguage: string, optional
	AcceptLanguage terra.StringValue `hcl:"accept_language,attr"`
	// Id: string, required
	Id terra.StringValue `hcl:"id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *dataservicecatalogproduct.Timeouts `hcl:"timeouts,block"`
}
type dataServicecatalogProductAttributes struct {
	ref terra.Reference
}

// AcceptLanguage returns a reference to field accept_language of aws_servicecatalog_product.
func (sp dataServicecatalogProductAttributes) AcceptLanguage() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("accept_language"))
}

// Arn returns a reference to field arn of aws_servicecatalog_product.
func (sp dataServicecatalogProductAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("arn"))
}

// CreatedTime returns a reference to field created_time of aws_servicecatalog_product.
func (sp dataServicecatalogProductAttributes) CreatedTime() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("created_time"))
}

// Description returns a reference to field description of aws_servicecatalog_product.
func (sp dataServicecatalogProductAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("description"))
}

// Distributor returns a reference to field distributor of aws_servicecatalog_product.
func (sp dataServicecatalogProductAttributes) Distributor() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("distributor"))
}

// HasDefaultPath returns a reference to field has_default_path of aws_servicecatalog_product.
func (sp dataServicecatalogProductAttributes) HasDefaultPath() terra.BoolValue {
	return terra.ReferenceAsBool(sp.ref.Append("has_default_path"))
}

// Id returns a reference to field id of aws_servicecatalog_product.
func (sp dataServicecatalogProductAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("id"))
}

// Name returns a reference to field name of aws_servicecatalog_product.
func (sp dataServicecatalogProductAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("name"))
}

// Owner returns a reference to field owner of aws_servicecatalog_product.
func (sp dataServicecatalogProductAttributes) Owner() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("owner"))
}

// Status returns a reference to field status of aws_servicecatalog_product.
func (sp dataServicecatalogProductAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("status"))
}

// SupportDescription returns a reference to field support_description of aws_servicecatalog_product.
func (sp dataServicecatalogProductAttributes) SupportDescription() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("support_description"))
}

// SupportEmail returns a reference to field support_email of aws_servicecatalog_product.
func (sp dataServicecatalogProductAttributes) SupportEmail() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("support_email"))
}

// SupportUrl returns a reference to field support_url of aws_servicecatalog_product.
func (sp dataServicecatalogProductAttributes) SupportUrl() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("support_url"))
}

// Tags returns a reference to field tags of aws_servicecatalog_product.
func (sp dataServicecatalogProductAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sp.ref.Append("tags"))
}

// Type returns a reference to field type of aws_servicecatalog_product.
func (sp dataServicecatalogProductAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("type"))
}

func (sp dataServicecatalogProductAttributes) Timeouts() dataservicecatalogproduct.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataservicecatalogproduct.TimeoutsAttributes](sp.ref.Append("timeouts"))
}
