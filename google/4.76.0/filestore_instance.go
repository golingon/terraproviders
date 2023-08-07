// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	filestoreinstance "github.com/golingon/terraproviders/google/4.76.0/filestoreinstance"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewFilestoreInstance creates a new instance of [FilestoreInstance].
func NewFilestoreInstance(name string, args FilestoreInstanceArgs) *FilestoreInstance {
	return &FilestoreInstance{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FilestoreInstance)(nil)

// FilestoreInstance represents the Terraform resource google_filestore_instance.
type FilestoreInstance struct {
	Name      string
	Args      FilestoreInstanceArgs
	state     *filestoreInstanceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [FilestoreInstance].
func (fi *FilestoreInstance) Type() string {
	return "google_filestore_instance"
}

// LocalName returns the local name for [FilestoreInstance].
func (fi *FilestoreInstance) LocalName() string {
	return fi.Name
}

// Configuration returns the configuration (args) for [FilestoreInstance].
func (fi *FilestoreInstance) Configuration() interface{} {
	return fi.Args
}

// DependOn is used for other resources to depend on [FilestoreInstance].
func (fi *FilestoreInstance) DependOn() terra.Reference {
	return terra.ReferenceResource(fi)
}

// Dependencies returns the list of resources [FilestoreInstance] depends_on.
func (fi *FilestoreInstance) Dependencies() terra.Dependencies {
	return fi.DependsOn
}

// LifecycleManagement returns the lifecycle block for [FilestoreInstance].
func (fi *FilestoreInstance) LifecycleManagement() *terra.Lifecycle {
	return fi.Lifecycle
}

// Attributes returns the attributes for [FilestoreInstance].
func (fi *FilestoreInstance) Attributes() filestoreInstanceAttributes {
	return filestoreInstanceAttributes{ref: terra.ReferenceResource(fi)}
}

// ImportState imports the given attribute values into [FilestoreInstance]'s state.
func (fi *FilestoreInstance) ImportState(av io.Reader) error {
	fi.state = &filestoreInstanceState{}
	if err := json.NewDecoder(av).Decode(fi.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", fi.Type(), fi.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [FilestoreInstance] has state.
func (fi *FilestoreInstance) State() (*filestoreInstanceState, bool) {
	return fi.state, fi.state != nil
}

// StateMust returns the state for [FilestoreInstance]. Panics if the state is nil.
func (fi *FilestoreInstance) StateMust() *filestoreInstanceState {
	if fi.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", fi.Type(), fi.LocalName()))
	}
	return fi.state
}

// FilestoreInstanceArgs contains the configurations for google_filestore_instance.
type FilestoreInstanceArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KmsKeyName: string, optional
	KmsKeyName terra.StringValue `hcl:"kms_key_name,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Tier: string, required
	Tier terra.StringValue `hcl:"tier,attr" validate:"required"`
	// Zone: string, optional
	Zone terra.StringValue `hcl:"zone,attr"`
	// FileShares: required
	FileShares *filestoreinstance.FileShares `hcl:"file_shares,block" validate:"required"`
	// Networks: min=1
	Networks []filestoreinstance.Networks `hcl:"networks,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *filestoreinstance.Timeouts `hcl:"timeouts,block"`
}
type filestoreInstanceAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_filestore_instance.
func (fi filestoreInstanceAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(fi.ref.Append("create_time"))
}

// Description returns a reference to field description of google_filestore_instance.
func (fi filestoreInstanceAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(fi.ref.Append("description"))
}

// Etag returns a reference to field etag of google_filestore_instance.
func (fi filestoreInstanceAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(fi.ref.Append("etag"))
}

// Id returns a reference to field id of google_filestore_instance.
func (fi filestoreInstanceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(fi.ref.Append("id"))
}

// KmsKeyName returns a reference to field kms_key_name of google_filestore_instance.
func (fi filestoreInstanceAttributes) KmsKeyName() terra.StringValue {
	return terra.ReferenceAsString(fi.ref.Append("kms_key_name"))
}

// Labels returns a reference to field labels of google_filestore_instance.
func (fi filestoreInstanceAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](fi.ref.Append("labels"))
}

// Location returns a reference to field location of google_filestore_instance.
func (fi filestoreInstanceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(fi.ref.Append("location"))
}

// Name returns a reference to field name of google_filestore_instance.
func (fi filestoreInstanceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(fi.ref.Append("name"))
}

// Project returns a reference to field project of google_filestore_instance.
func (fi filestoreInstanceAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(fi.ref.Append("project"))
}

// Tier returns a reference to field tier of google_filestore_instance.
func (fi filestoreInstanceAttributes) Tier() terra.StringValue {
	return terra.ReferenceAsString(fi.ref.Append("tier"))
}

// Zone returns a reference to field zone of google_filestore_instance.
func (fi filestoreInstanceAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(fi.ref.Append("zone"))
}

func (fi filestoreInstanceAttributes) FileShares() terra.ListValue[filestoreinstance.FileSharesAttributes] {
	return terra.ReferenceAsList[filestoreinstance.FileSharesAttributes](fi.ref.Append("file_shares"))
}

func (fi filestoreInstanceAttributes) Networks() terra.ListValue[filestoreinstance.NetworksAttributes] {
	return terra.ReferenceAsList[filestoreinstance.NetworksAttributes](fi.ref.Append("networks"))
}

func (fi filestoreInstanceAttributes) Timeouts() filestoreinstance.TimeoutsAttributes {
	return terra.ReferenceAsSingle[filestoreinstance.TimeoutsAttributes](fi.ref.Append("timeouts"))
}

type filestoreInstanceState struct {
	CreateTime  string                              `json:"create_time"`
	Description string                              `json:"description"`
	Etag        string                              `json:"etag"`
	Id          string                              `json:"id"`
	KmsKeyName  string                              `json:"kms_key_name"`
	Labels      map[string]string                   `json:"labels"`
	Location    string                              `json:"location"`
	Name        string                              `json:"name"`
	Project     string                              `json:"project"`
	Tier        string                              `json:"tier"`
	Zone        string                              `json:"zone"`
	FileShares  []filestoreinstance.FileSharesState `json:"file_shares"`
	Networks    []filestoreinstance.NetworksState   `json:"networks"`
	Timeouts    *filestoreinstance.TimeoutsState    `json:"timeouts"`
}
