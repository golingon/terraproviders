// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	ecrpublicrepository "github.com/golingon/terraproviders/aws/4.60.0/ecrpublicrepository"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewEcrpublicRepository(name string, args EcrpublicRepositoryArgs) *EcrpublicRepository {
	return &EcrpublicRepository{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EcrpublicRepository)(nil)

type EcrpublicRepository struct {
	Name  string
	Args  EcrpublicRepositoryArgs
	state *ecrpublicRepositoryState
}

func (er *EcrpublicRepository) Type() string {
	return "aws_ecrpublic_repository"
}

func (er *EcrpublicRepository) LocalName() string {
	return er.Name
}

func (er *EcrpublicRepository) Configuration() interface{} {
	return er.Args
}

func (er *EcrpublicRepository) Attributes() ecrpublicRepositoryAttributes {
	return ecrpublicRepositoryAttributes{ref: terra.ReferenceResource(er)}
}

func (er *EcrpublicRepository) ImportState(av io.Reader) error {
	er.state = &ecrpublicRepositoryState{}
	if err := json.NewDecoder(av).Decode(er.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", er.Type(), er.LocalName(), err)
	}
	return nil
}

func (er *EcrpublicRepository) State() (*ecrpublicRepositoryState, bool) {
	return er.state, er.state != nil
}

func (er *EcrpublicRepository) StateMust() *ecrpublicRepositoryState {
	if er.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", er.Type(), er.LocalName()))
	}
	return er.state
}

func (er *EcrpublicRepository) DependOn() terra.Reference {
	return terra.ReferenceResource(er)
}

type EcrpublicRepositoryArgs struct {
	// ForceDestroy: bool, optional
	ForceDestroy terra.BoolValue `hcl:"force_destroy,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// RepositoryName: string, required
	RepositoryName terra.StringValue `hcl:"repository_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// CatalogData: optional
	CatalogData *ecrpublicrepository.CatalogData `hcl:"catalog_data,block"`
	// Timeouts: optional
	Timeouts *ecrpublicrepository.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that EcrpublicRepository depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type ecrpublicRepositoryAttributes struct {
	ref terra.Reference
}

func (er ecrpublicRepositoryAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(er.ref.Append("arn"))
}

func (er ecrpublicRepositoryAttributes) ForceDestroy() terra.BoolValue {
	return terra.ReferenceBool(er.ref.Append("force_destroy"))
}

func (er ecrpublicRepositoryAttributes) Id() terra.StringValue {
	return terra.ReferenceString(er.ref.Append("id"))
}

func (er ecrpublicRepositoryAttributes) RegistryId() terra.StringValue {
	return terra.ReferenceString(er.ref.Append("registry_id"))
}

func (er ecrpublicRepositoryAttributes) RepositoryName() terra.StringValue {
	return terra.ReferenceString(er.ref.Append("repository_name"))
}

func (er ecrpublicRepositoryAttributes) RepositoryUri() terra.StringValue {
	return terra.ReferenceString(er.ref.Append("repository_uri"))
}

func (er ecrpublicRepositoryAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](er.ref.Append("tags"))
}

func (er ecrpublicRepositoryAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](er.ref.Append("tags_all"))
}

func (er ecrpublicRepositoryAttributes) CatalogData() terra.ListValue[ecrpublicrepository.CatalogDataAttributes] {
	return terra.ReferenceList[ecrpublicrepository.CatalogDataAttributes](er.ref.Append("catalog_data"))
}

func (er ecrpublicRepositoryAttributes) Timeouts() ecrpublicrepository.TimeoutsAttributes {
	return terra.ReferenceSingle[ecrpublicrepository.TimeoutsAttributes](er.ref.Append("timeouts"))
}

type ecrpublicRepositoryState struct {
	Arn            string                                 `json:"arn"`
	ForceDestroy   bool                                   `json:"force_destroy"`
	Id             string                                 `json:"id"`
	RegistryId     string                                 `json:"registry_id"`
	RepositoryName string                                 `json:"repository_name"`
	RepositoryUri  string                                 `json:"repository_uri"`
	Tags           map[string]string                      `json:"tags"`
	TagsAll        map[string]string                      `json:"tags_all"`
	CatalogData    []ecrpublicrepository.CatalogDataState `json:"catalog_data"`
	Timeouts       *ecrpublicrepository.TimeoutsState     `json:"timeouts"`
}
