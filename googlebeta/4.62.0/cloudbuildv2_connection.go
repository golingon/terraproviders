// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	cloudbuildv2connection "github.com/golingon/terraproviders/googlebeta/4.62.0/cloudbuildv2connection"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudbuildv2Connection creates a new instance of [Cloudbuildv2Connection].
func NewCloudbuildv2Connection(name string, args Cloudbuildv2ConnectionArgs) *Cloudbuildv2Connection {
	return &Cloudbuildv2Connection{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Cloudbuildv2Connection)(nil)

// Cloudbuildv2Connection represents the Terraform resource google_cloudbuildv2_connection.
type Cloudbuildv2Connection struct {
	Name      string
	Args      Cloudbuildv2ConnectionArgs
	state     *cloudbuildv2ConnectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Cloudbuildv2Connection].
func (cc *Cloudbuildv2Connection) Type() string {
	return "google_cloudbuildv2_connection"
}

// LocalName returns the local name for [Cloudbuildv2Connection].
func (cc *Cloudbuildv2Connection) LocalName() string {
	return cc.Name
}

// Configuration returns the configuration (args) for [Cloudbuildv2Connection].
func (cc *Cloudbuildv2Connection) Configuration() interface{} {
	return cc.Args
}

// DependOn is used for other resources to depend on [Cloudbuildv2Connection].
func (cc *Cloudbuildv2Connection) DependOn() terra.Reference {
	return terra.ReferenceResource(cc)
}

// Dependencies returns the list of resources [Cloudbuildv2Connection] depends_on.
func (cc *Cloudbuildv2Connection) Dependencies() terra.Dependencies {
	return cc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Cloudbuildv2Connection].
func (cc *Cloudbuildv2Connection) LifecycleManagement() *terra.Lifecycle {
	return cc.Lifecycle
}

// Attributes returns the attributes for [Cloudbuildv2Connection].
func (cc *Cloudbuildv2Connection) Attributes() cloudbuildv2ConnectionAttributes {
	return cloudbuildv2ConnectionAttributes{ref: terra.ReferenceResource(cc)}
}

// ImportState imports the given attribute values into [Cloudbuildv2Connection]'s state.
func (cc *Cloudbuildv2Connection) ImportState(av io.Reader) error {
	cc.state = &cloudbuildv2ConnectionState{}
	if err := json.NewDecoder(av).Decode(cc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cc.Type(), cc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Cloudbuildv2Connection] has state.
func (cc *Cloudbuildv2Connection) State() (*cloudbuildv2ConnectionState, bool) {
	return cc.state, cc.state != nil
}

// StateMust returns the state for [Cloudbuildv2Connection]. Panics if the state is nil.
func (cc *Cloudbuildv2Connection) StateMust() *cloudbuildv2ConnectionState {
	if cc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cc.Type(), cc.LocalName()))
	}
	return cc.state
}

// Cloudbuildv2ConnectionArgs contains the configurations for google_cloudbuildv2_connection.
type Cloudbuildv2ConnectionArgs struct {
	// Annotations: map of string, optional
	Annotations terra.MapValue[terra.StringValue] `hcl:"annotations,attr"`
	// Disabled: bool, optional
	Disabled terra.BoolValue `hcl:"disabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// InstallationState: min=0
	InstallationState []cloudbuildv2connection.InstallationState `hcl:"installation_state,block" validate:"min=0"`
	// GithubConfig: optional
	GithubConfig *cloudbuildv2connection.GithubConfig `hcl:"github_config,block"`
	// GithubEnterpriseConfig: optional
	GithubEnterpriseConfig *cloudbuildv2connection.GithubEnterpriseConfig `hcl:"github_enterprise_config,block"`
	// Timeouts: optional
	Timeouts *cloudbuildv2connection.Timeouts `hcl:"timeouts,block"`
}
type cloudbuildv2ConnectionAttributes struct {
	ref terra.Reference
}

// Annotations returns a reference to field annotations of google_cloudbuildv2_connection.
func (cc cloudbuildv2ConnectionAttributes) Annotations() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cc.ref.Append("annotations"))
}

// CreateTime returns a reference to field create_time of google_cloudbuildv2_connection.
func (cc cloudbuildv2ConnectionAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("create_time"))
}

// Disabled returns a reference to field disabled of google_cloudbuildv2_connection.
func (cc cloudbuildv2ConnectionAttributes) Disabled() terra.BoolValue {
	return terra.ReferenceAsBool(cc.ref.Append("disabled"))
}

// Etag returns a reference to field etag of google_cloudbuildv2_connection.
func (cc cloudbuildv2ConnectionAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("etag"))
}

// Id returns a reference to field id of google_cloudbuildv2_connection.
func (cc cloudbuildv2ConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("id"))
}

// Location returns a reference to field location of google_cloudbuildv2_connection.
func (cc cloudbuildv2ConnectionAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("location"))
}

// Name returns a reference to field name of google_cloudbuildv2_connection.
func (cc cloudbuildv2ConnectionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("name"))
}

// Project returns a reference to field project of google_cloudbuildv2_connection.
func (cc cloudbuildv2ConnectionAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("project"))
}

// Reconciling returns a reference to field reconciling of google_cloudbuildv2_connection.
func (cc cloudbuildv2ConnectionAttributes) Reconciling() terra.BoolValue {
	return terra.ReferenceAsBool(cc.ref.Append("reconciling"))
}

// UpdateTime returns a reference to field update_time of google_cloudbuildv2_connection.
func (cc cloudbuildv2ConnectionAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("update_time"))
}

func (cc cloudbuildv2ConnectionAttributes) InstallationState() terra.ListValue[cloudbuildv2connection.InstallationStateAttributes] {
	return terra.ReferenceAsList[cloudbuildv2connection.InstallationStateAttributes](cc.ref.Append("installation_state"))
}

func (cc cloudbuildv2ConnectionAttributes) GithubConfig() terra.ListValue[cloudbuildv2connection.GithubConfigAttributes] {
	return terra.ReferenceAsList[cloudbuildv2connection.GithubConfigAttributes](cc.ref.Append("github_config"))
}

func (cc cloudbuildv2ConnectionAttributes) GithubEnterpriseConfig() terra.ListValue[cloudbuildv2connection.GithubEnterpriseConfigAttributes] {
	return terra.ReferenceAsList[cloudbuildv2connection.GithubEnterpriseConfigAttributes](cc.ref.Append("github_enterprise_config"))
}

func (cc cloudbuildv2ConnectionAttributes) Timeouts() cloudbuildv2connection.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cloudbuildv2connection.TimeoutsAttributes](cc.ref.Append("timeouts"))
}

type cloudbuildv2ConnectionState struct {
	Annotations            map[string]string                                    `json:"annotations"`
	CreateTime             string                                               `json:"create_time"`
	Disabled               bool                                                 `json:"disabled"`
	Etag                   string                                               `json:"etag"`
	Id                     string                                               `json:"id"`
	Location               string                                               `json:"location"`
	Name                   string                                               `json:"name"`
	Project                string                                               `json:"project"`
	Reconciling            bool                                                 `json:"reconciling"`
	UpdateTime             string                                               `json:"update_time"`
	InstallationState      []cloudbuildv2connection.InstallationStateState      `json:"installation_state"`
	GithubConfig           []cloudbuildv2connection.GithubConfigState           `json:"github_config"`
	GithubEnterpriseConfig []cloudbuildv2connection.GithubEnterpriseConfigState `json:"github_enterprise_config"`
	Timeouts               *cloudbuildv2connection.TimeoutsState                `json:"timeouts"`
}
