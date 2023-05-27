// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataservicecatalogprovisioningartifacts "github.com/golingon/terraproviders/aws/5.0.1/dataservicecatalogprovisioningartifacts"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataServicecatalogProvisioningArtifacts creates a new instance of [DataServicecatalogProvisioningArtifacts].
func NewDataServicecatalogProvisioningArtifacts(name string, args DataServicecatalogProvisioningArtifactsArgs) *DataServicecatalogProvisioningArtifacts {
	return &DataServicecatalogProvisioningArtifacts{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataServicecatalogProvisioningArtifacts)(nil)

// DataServicecatalogProvisioningArtifacts represents the Terraform data resource aws_servicecatalog_provisioning_artifacts.
type DataServicecatalogProvisioningArtifacts struct {
	Name string
	Args DataServicecatalogProvisioningArtifactsArgs
}

// DataSource returns the Terraform object type for [DataServicecatalogProvisioningArtifacts].
func (spa *DataServicecatalogProvisioningArtifacts) DataSource() string {
	return "aws_servicecatalog_provisioning_artifacts"
}

// LocalName returns the local name for [DataServicecatalogProvisioningArtifacts].
func (spa *DataServicecatalogProvisioningArtifacts) LocalName() string {
	return spa.Name
}

// Configuration returns the configuration (args) for [DataServicecatalogProvisioningArtifacts].
func (spa *DataServicecatalogProvisioningArtifacts) Configuration() interface{} {
	return spa.Args
}

// Attributes returns the attributes for [DataServicecatalogProvisioningArtifacts].
func (spa *DataServicecatalogProvisioningArtifacts) Attributes() dataServicecatalogProvisioningArtifactsAttributes {
	return dataServicecatalogProvisioningArtifactsAttributes{ref: terra.ReferenceDataResource(spa)}
}

// DataServicecatalogProvisioningArtifactsArgs contains the configurations for aws_servicecatalog_provisioning_artifacts.
type DataServicecatalogProvisioningArtifactsArgs struct {
	// AcceptLanguage: string, optional
	AcceptLanguage terra.StringValue `hcl:"accept_language,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ProductId: string, required
	ProductId terra.StringValue `hcl:"product_id,attr" validate:"required"`
	// ProvisioningArtifactDetails: min=0
	ProvisioningArtifactDetails []dataservicecatalogprovisioningartifacts.ProvisioningArtifactDetails `hcl:"provisioning_artifact_details,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataservicecatalogprovisioningartifacts.Timeouts `hcl:"timeouts,block"`
}
type dataServicecatalogProvisioningArtifactsAttributes struct {
	ref terra.Reference
}

// AcceptLanguage returns a reference to field accept_language of aws_servicecatalog_provisioning_artifacts.
func (spa dataServicecatalogProvisioningArtifactsAttributes) AcceptLanguage() terra.StringValue {
	return terra.ReferenceAsString(spa.ref.Append("accept_language"))
}

// Id returns a reference to field id of aws_servicecatalog_provisioning_artifacts.
func (spa dataServicecatalogProvisioningArtifactsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(spa.ref.Append("id"))
}

// ProductId returns a reference to field product_id of aws_servicecatalog_provisioning_artifacts.
func (spa dataServicecatalogProvisioningArtifactsAttributes) ProductId() terra.StringValue {
	return terra.ReferenceAsString(spa.ref.Append("product_id"))
}

func (spa dataServicecatalogProvisioningArtifactsAttributes) ProvisioningArtifactDetails() terra.ListValue[dataservicecatalogprovisioningartifacts.ProvisioningArtifactDetailsAttributes] {
	return terra.ReferenceAsList[dataservicecatalogprovisioningartifacts.ProvisioningArtifactDetailsAttributes](spa.ref.Append("provisioning_artifact_details"))
}

func (spa dataServicecatalogProvisioningArtifactsAttributes) Timeouts() dataservicecatalogprovisioningartifacts.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataservicecatalogprovisioningartifacts.TimeoutsAttributes](spa.ref.Append("timeouts"))
}
