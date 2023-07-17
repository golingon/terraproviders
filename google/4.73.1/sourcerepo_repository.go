// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	sourcereporepository "github.com/golingon/terraproviders/google/4.73.1/sourcereporepository"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSourcerepoRepository creates a new instance of [SourcerepoRepository].
func NewSourcerepoRepository(name string, args SourcerepoRepositoryArgs) *SourcerepoRepository {
	return &SourcerepoRepository{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SourcerepoRepository)(nil)

// SourcerepoRepository represents the Terraform resource google_sourcerepo_repository.
type SourcerepoRepository struct {
	Name      string
	Args      SourcerepoRepositoryArgs
	state     *sourcerepoRepositoryState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SourcerepoRepository].
func (sr *SourcerepoRepository) Type() string {
	return "google_sourcerepo_repository"
}

// LocalName returns the local name for [SourcerepoRepository].
func (sr *SourcerepoRepository) LocalName() string {
	return sr.Name
}

// Configuration returns the configuration (args) for [SourcerepoRepository].
func (sr *SourcerepoRepository) Configuration() interface{} {
	return sr.Args
}

// DependOn is used for other resources to depend on [SourcerepoRepository].
func (sr *SourcerepoRepository) DependOn() terra.Reference {
	return terra.ReferenceResource(sr)
}

// Dependencies returns the list of resources [SourcerepoRepository] depends_on.
func (sr *SourcerepoRepository) Dependencies() terra.Dependencies {
	return sr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SourcerepoRepository].
func (sr *SourcerepoRepository) LifecycleManagement() *terra.Lifecycle {
	return sr.Lifecycle
}

// Attributes returns the attributes for [SourcerepoRepository].
func (sr *SourcerepoRepository) Attributes() sourcerepoRepositoryAttributes {
	return sourcerepoRepositoryAttributes{ref: terra.ReferenceResource(sr)}
}

// ImportState imports the given attribute values into [SourcerepoRepository]'s state.
func (sr *SourcerepoRepository) ImportState(av io.Reader) error {
	sr.state = &sourcerepoRepositoryState{}
	if err := json.NewDecoder(av).Decode(sr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sr.Type(), sr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SourcerepoRepository] has state.
func (sr *SourcerepoRepository) State() (*sourcerepoRepositoryState, bool) {
	return sr.state, sr.state != nil
}

// StateMust returns the state for [SourcerepoRepository]. Panics if the state is nil.
func (sr *SourcerepoRepository) StateMust() *sourcerepoRepositoryState {
	if sr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sr.Type(), sr.LocalName()))
	}
	return sr.state
}

// SourcerepoRepositoryArgs contains the configurations for google_sourcerepo_repository.
type SourcerepoRepositoryArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// PubsubConfigs: min=0
	PubsubConfigs []sourcereporepository.PubsubConfigs `hcl:"pubsub_configs,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *sourcereporepository.Timeouts `hcl:"timeouts,block"`
}
type sourcerepoRepositoryAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_sourcerepo_repository.
func (sr sourcerepoRepositoryAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sr.ref.Append("id"))
}

// Name returns a reference to field name of google_sourcerepo_repository.
func (sr sourcerepoRepositoryAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sr.ref.Append("name"))
}

// Project returns a reference to field project of google_sourcerepo_repository.
func (sr sourcerepoRepositoryAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(sr.ref.Append("project"))
}

// Size returns a reference to field size of google_sourcerepo_repository.
func (sr sourcerepoRepositoryAttributes) Size() terra.NumberValue {
	return terra.ReferenceAsNumber(sr.ref.Append("size"))
}

// Url returns a reference to field url of google_sourcerepo_repository.
func (sr sourcerepoRepositoryAttributes) Url() terra.StringValue {
	return terra.ReferenceAsString(sr.ref.Append("url"))
}

func (sr sourcerepoRepositoryAttributes) PubsubConfigs() terra.SetValue[sourcereporepository.PubsubConfigsAttributes] {
	return terra.ReferenceAsSet[sourcereporepository.PubsubConfigsAttributes](sr.ref.Append("pubsub_configs"))
}

func (sr sourcerepoRepositoryAttributes) Timeouts() sourcereporepository.TimeoutsAttributes {
	return terra.ReferenceAsSingle[sourcereporepository.TimeoutsAttributes](sr.ref.Append("timeouts"))
}

type sourcerepoRepositoryState struct {
	Id            string                                    `json:"id"`
	Name          string                                    `json:"name"`
	Project       string                                    `json:"project"`
	Size          float64                                   `json:"size"`
	Url           string                                    `json:"url"`
	PubsubConfigs []sourcereporepository.PubsubConfigsState `json:"pubsub_configs"`
	Timeouts      *sourcereporepository.TimeoutsState       `json:"timeouts"`
}
