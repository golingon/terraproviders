// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	cloudbuildv2repository "github.com/golingon/terraproviders/googlebeta/5.24.0/cloudbuildv2repository"
	"io"
)

// NewCloudbuildv2Repository creates a new instance of [Cloudbuildv2Repository].
func NewCloudbuildv2Repository(name string, args Cloudbuildv2RepositoryArgs) *Cloudbuildv2Repository {
	return &Cloudbuildv2Repository{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Cloudbuildv2Repository)(nil)

// Cloudbuildv2Repository represents the Terraform resource google_cloudbuildv2_repository.
type Cloudbuildv2Repository struct {
	Name      string
	Args      Cloudbuildv2RepositoryArgs
	state     *cloudbuildv2RepositoryState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Cloudbuildv2Repository].
func (cr *Cloudbuildv2Repository) Type() string {
	return "google_cloudbuildv2_repository"
}

// LocalName returns the local name for [Cloudbuildv2Repository].
func (cr *Cloudbuildv2Repository) LocalName() string {
	return cr.Name
}

// Configuration returns the configuration (args) for [Cloudbuildv2Repository].
func (cr *Cloudbuildv2Repository) Configuration() interface{} {
	return cr.Args
}

// DependOn is used for other resources to depend on [Cloudbuildv2Repository].
func (cr *Cloudbuildv2Repository) DependOn() terra.Reference {
	return terra.ReferenceResource(cr)
}

// Dependencies returns the list of resources [Cloudbuildv2Repository] depends_on.
func (cr *Cloudbuildv2Repository) Dependencies() terra.Dependencies {
	return cr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Cloudbuildv2Repository].
func (cr *Cloudbuildv2Repository) LifecycleManagement() *terra.Lifecycle {
	return cr.Lifecycle
}

// Attributes returns the attributes for [Cloudbuildv2Repository].
func (cr *Cloudbuildv2Repository) Attributes() cloudbuildv2RepositoryAttributes {
	return cloudbuildv2RepositoryAttributes{ref: terra.ReferenceResource(cr)}
}

// ImportState imports the given attribute values into [Cloudbuildv2Repository]'s state.
func (cr *Cloudbuildv2Repository) ImportState(av io.Reader) error {
	cr.state = &cloudbuildv2RepositoryState{}
	if err := json.NewDecoder(av).Decode(cr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cr.Type(), cr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Cloudbuildv2Repository] has state.
func (cr *Cloudbuildv2Repository) State() (*cloudbuildv2RepositoryState, bool) {
	return cr.state, cr.state != nil
}

// StateMust returns the state for [Cloudbuildv2Repository]. Panics if the state is nil.
func (cr *Cloudbuildv2Repository) StateMust() *cloudbuildv2RepositoryState {
	if cr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cr.Type(), cr.LocalName()))
	}
	return cr.state
}

// Cloudbuildv2RepositoryArgs contains the configurations for google_cloudbuildv2_repository.
type Cloudbuildv2RepositoryArgs struct {
	// Annotations: map of string, optional
	Annotations terra.MapValue[terra.StringValue] `hcl:"annotations,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ParentConnection: string, required
	ParentConnection terra.StringValue `hcl:"parent_connection,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// RemoteUri: string, required
	RemoteUri terra.StringValue `hcl:"remote_uri,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *cloudbuildv2repository.Timeouts `hcl:"timeouts,block"`
}
type cloudbuildv2RepositoryAttributes struct {
	ref terra.Reference
}

// Annotations returns a reference to field annotations of google_cloudbuildv2_repository.
func (cr cloudbuildv2RepositoryAttributes) Annotations() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cr.ref.Append("annotations"))
}

// CreateTime returns a reference to field create_time of google_cloudbuildv2_repository.
func (cr cloudbuildv2RepositoryAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("create_time"))
}

// EffectiveAnnotations returns a reference to field effective_annotations of google_cloudbuildv2_repository.
func (cr cloudbuildv2RepositoryAttributes) EffectiveAnnotations() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cr.ref.Append("effective_annotations"))
}

// Etag returns a reference to field etag of google_cloudbuildv2_repository.
func (cr cloudbuildv2RepositoryAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("etag"))
}

// Id returns a reference to field id of google_cloudbuildv2_repository.
func (cr cloudbuildv2RepositoryAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("id"))
}

// Location returns a reference to field location of google_cloudbuildv2_repository.
func (cr cloudbuildv2RepositoryAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("location"))
}

// Name returns a reference to field name of google_cloudbuildv2_repository.
func (cr cloudbuildv2RepositoryAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("name"))
}

// ParentConnection returns a reference to field parent_connection of google_cloudbuildv2_repository.
func (cr cloudbuildv2RepositoryAttributes) ParentConnection() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("parent_connection"))
}

// Project returns a reference to field project of google_cloudbuildv2_repository.
func (cr cloudbuildv2RepositoryAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("project"))
}

// RemoteUri returns a reference to field remote_uri of google_cloudbuildv2_repository.
func (cr cloudbuildv2RepositoryAttributes) RemoteUri() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("remote_uri"))
}

// UpdateTime returns a reference to field update_time of google_cloudbuildv2_repository.
func (cr cloudbuildv2RepositoryAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("update_time"))
}

func (cr cloudbuildv2RepositoryAttributes) Timeouts() cloudbuildv2repository.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cloudbuildv2repository.TimeoutsAttributes](cr.ref.Append("timeouts"))
}

type cloudbuildv2RepositoryState struct {
	Annotations          map[string]string                     `json:"annotations"`
	CreateTime           string                                `json:"create_time"`
	EffectiveAnnotations map[string]string                     `json:"effective_annotations"`
	Etag                 string                                `json:"etag"`
	Id                   string                                `json:"id"`
	Location             string                                `json:"location"`
	Name                 string                                `json:"name"`
	ParentConnection     string                                `json:"parent_connection"`
	Project              string                                `json:"project"`
	RemoteUri            string                                `json:"remote_uri"`
	UpdateTime           string                                `json:"update_time"`
	Timeouts             *cloudbuildv2repository.TimeoutsState `json:"timeouts"`
}
