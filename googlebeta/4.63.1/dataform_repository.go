// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	dataformrepository "github.com/golingon/terraproviders/googlebeta/4.63.1/dataformrepository"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataformRepository creates a new instance of [DataformRepository].
func NewDataformRepository(name string, args DataformRepositoryArgs) *DataformRepository {
	return &DataformRepository{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataformRepository)(nil)

// DataformRepository represents the Terraform resource google_dataform_repository.
type DataformRepository struct {
	Name      string
	Args      DataformRepositoryArgs
	state     *dataformRepositoryState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataformRepository].
func (dr *DataformRepository) Type() string {
	return "google_dataform_repository"
}

// LocalName returns the local name for [DataformRepository].
func (dr *DataformRepository) LocalName() string {
	return dr.Name
}

// Configuration returns the configuration (args) for [DataformRepository].
func (dr *DataformRepository) Configuration() interface{} {
	return dr.Args
}

// DependOn is used for other resources to depend on [DataformRepository].
func (dr *DataformRepository) DependOn() terra.Reference {
	return terra.ReferenceResource(dr)
}

// Dependencies returns the list of resources [DataformRepository] depends_on.
func (dr *DataformRepository) Dependencies() terra.Dependencies {
	return dr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataformRepository].
func (dr *DataformRepository) LifecycleManagement() *terra.Lifecycle {
	return dr.Lifecycle
}

// Attributes returns the attributes for [DataformRepository].
func (dr *DataformRepository) Attributes() dataformRepositoryAttributes {
	return dataformRepositoryAttributes{ref: terra.ReferenceResource(dr)}
}

// ImportState imports the given attribute values into [DataformRepository]'s state.
func (dr *DataformRepository) ImportState(av io.Reader) error {
	dr.state = &dataformRepositoryState{}
	if err := json.NewDecoder(av).Decode(dr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dr.Type(), dr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataformRepository] has state.
func (dr *DataformRepository) State() (*dataformRepositoryState, bool) {
	return dr.state, dr.state != nil
}

// StateMust returns the state for [DataformRepository]. Panics if the state is nil.
func (dr *DataformRepository) StateMust() *dataformRepositoryState {
	if dr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dr.Type(), dr.LocalName()))
	}
	return dr.state
}

// DataformRepositoryArgs contains the configurations for google_dataform_repository.
type DataformRepositoryArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// GitRemoteSettings: optional
	GitRemoteSettings *dataformrepository.GitRemoteSettings `hcl:"git_remote_settings,block"`
	// Timeouts: optional
	Timeouts *dataformrepository.Timeouts `hcl:"timeouts,block"`
}
type dataformRepositoryAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_dataform_repository.
func (dr dataformRepositoryAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dr.ref.Append("id"))
}

// Name returns a reference to field name of google_dataform_repository.
func (dr dataformRepositoryAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dr.ref.Append("name"))
}

// Project returns a reference to field project of google_dataform_repository.
func (dr dataformRepositoryAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dr.ref.Append("project"))
}

// Region returns a reference to field region of google_dataform_repository.
func (dr dataformRepositoryAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(dr.ref.Append("region"))
}

func (dr dataformRepositoryAttributes) GitRemoteSettings() terra.ListValue[dataformrepository.GitRemoteSettingsAttributes] {
	return terra.ReferenceAsList[dataformrepository.GitRemoteSettingsAttributes](dr.ref.Append("git_remote_settings"))
}

func (dr dataformRepositoryAttributes) Timeouts() dataformrepository.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataformrepository.TimeoutsAttributes](dr.ref.Append("timeouts"))
}

type dataformRepositoryState struct {
	Id                string                                      `json:"id"`
	Name              string                                      `json:"name"`
	Project           string                                      `json:"project"`
	Region            string                                      `json:"region"`
	GitRemoteSettings []dataformrepository.GitRemoteSettingsState `json:"git_remote_settings"`
	Timeouts          *dataformrepository.TimeoutsState           `json:"timeouts"`
}
