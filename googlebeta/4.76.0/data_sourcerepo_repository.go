// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	datasourcereporepository "github.com/golingon/terraproviders/googlebeta/4.76.0/datasourcereporepository"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataSourcerepoRepository creates a new instance of [DataSourcerepoRepository].
func NewDataSourcerepoRepository(name string, args DataSourcerepoRepositoryArgs) *DataSourcerepoRepository {
	return &DataSourcerepoRepository{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSourcerepoRepository)(nil)

// DataSourcerepoRepository represents the Terraform data resource google_sourcerepo_repository.
type DataSourcerepoRepository struct {
	Name string
	Args DataSourcerepoRepositoryArgs
}

// DataSource returns the Terraform object type for [DataSourcerepoRepository].
func (sr *DataSourcerepoRepository) DataSource() string {
	return "google_sourcerepo_repository"
}

// LocalName returns the local name for [DataSourcerepoRepository].
func (sr *DataSourcerepoRepository) LocalName() string {
	return sr.Name
}

// Configuration returns the configuration (args) for [DataSourcerepoRepository].
func (sr *DataSourcerepoRepository) Configuration() interface{} {
	return sr.Args
}

// Attributes returns the attributes for [DataSourcerepoRepository].
func (sr *DataSourcerepoRepository) Attributes() dataSourcerepoRepositoryAttributes {
	return dataSourcerepoRepositoryAttributes{ref: terra.ReferenceDataResource(sr)}
}

// DataSourcerepoRepositoryArgs contains the configurations for google_sourcerepo_repository.
type DataSourcerepoRepositoryArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// PubsubConfigs: min=0
	PubsubConfigs []datasourcereporepository.PubsubConfigs `hcl:"pubsub_configs,block" validate:"min=0"`
}
type dataSourcerepoRepositoryAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_sourcerepo_repository.
func (sr dataSourcerepoRepositoryAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sr.ref.Append("id"))
}

// Name returns a reference to field name of google_sourcerepo_repository.
func (sr dataSourcerepoRepositoryAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sr.ref.Append("name"))
}

// Project returns a reference to field project of google_sourcerepo_repository.
func (sr dataSourcerepoRepositoryAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(sr.ref.Append("project"))
}

// Size returns a reference to field size of google_sourcerepo_repository.
func (sr dataSourcerepoRepositoryAttributes) Size() terra.NumberValue {
	return terra.ReferenceAsNumber(sr.ref.Append("size"))
}

// Url returns a reference to field url of google_sourcerepo_repository.
func (sr dataSourcerepoRepositoryAttributes) Url() terra.StringValue {
	return terra.ReferenceAsString(sr.ref.Append("url"))
}

func (sr dataSourcerepoRepositoryAttributes) PubsubConfigs() terra.SetValue[datasourcereporepository.PubsubConfigsAttributes] {
	return terra.ReferenceAsSet[datasourcereporepository.PubsubConfigsAttributes](sr.ref.Append("pubsub_configs"))
}
