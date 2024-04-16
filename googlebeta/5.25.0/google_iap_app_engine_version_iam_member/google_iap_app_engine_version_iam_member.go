// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_iap_app_engine_version_iam_member

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource google_iap_app_engine_version_iam_member.
type Resource struct {
	Name      string
	Args      Args
	state     *googleIapAppEngineVersionIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (giaevim *Resource) Type() string {
	return "google_iap_app_engine_version_iam_member"
}

// LocalName returns the local name for [Resource].
func (giaevim *Resource) LocalName() string {
	return giaevim.Name
}

// Configuration returns the configuration (args) for [Resource].
func (giaevim *Resource) Configuration() interface{} {
	return giaevim.Args
}

// DependOn is used for other resources to depend on [Resource].
func (giaevim *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(giaevim)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (giaevim *Resource) Dependencies() terra.Dependencies {
	return giaevim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (giaevim *Resource) LifecycleManagement() *terra.Lifecycle {
	return giaevim.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (giaevim *Resource) Attributes() googleIapAppEngineVersionIamMemberAttributes {
	return googleIapAppEngineVersionIamMemberAttributes{ref: terra.ReferenceResource(giaevim)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (giaevim *Resource) ImportState(state io.Reader) error {
	giaevim.state = &googleIapAppEngineVersionIamMemberState{}
	if err := json.NewDecoder(state).Decode(giaevim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", giaevim.Type(), giaevim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (giaevim *Resource) State() (*googleIapAppEngineVersionIamMemberState, bool) {
	return giaevim.state, giaevim.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (giaevim *Resource) StateMust() *googleIapAppEngineVersionIamMemberState {
	if giaevim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", giaevim.Type(), giaevim.LocalName()))
	}
	return giaevim.state
}

// Args contains the configurations for google_iap_app_engine_version_iam_member.
type Args struct {
	// AppId: string, required
	AppId terra.StringValue `hcl:"app_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Service: string, required
	Service terra.StringValue `hcl:"service,attr" validate:"required"`
	// VersionId: string, required
	VersionId terra.StringValue `hcl:"version_id,attr" validate:"required"`
	// Condition: optional
	Condition *Condition `hcl:"condition,block"`
}

type googleIapAppEngineVersionIamMemberAttributes struct {
	ref terra.Reference
}

// AppId returns a reference to field app_id of google_iap_app_engine_version_iam_member.
func (giaevim googleIapAppEngineVersionIamMemberAttributes) AppId() terra.StringValue {
	return terra.ReferenceAsString(giaevim.ref.Append("app_id"))
}

// Etag returns a reference to field etag of google_iap_app_engine_version_iam_member.
func (giaevim googleIapAppEngineVersionIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(giaevim.ref.Append("etag"))
}

// Id returns a reference to field id of google_iap_app_engine_version_iam_member.
func (giaevim googleIapAppEngineVersionIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(giaevim.ref.Append("id"))
}

// Member returns a reference to field member of google_iap_app_engine_version_iam_member.
func (giaevim googleIapAppEngineVersionIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(giaevim.ref.Append("member"))
}

// Project returns a reference to field project of google_iap_app_engine_version_iam_member.
func (giaevim googleIapAppEngineVersionIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(giaevim.ref.Append("project"))
}

// Role returns a reference to field role of google_iap_app_engine_version_iam_member.
func (giaevim googleIapAppEngineVersionIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(giaevim.ref.Append("role"))
}

// Service returns a reference to field service of google_iap_app_engine_version_iam_member.
func (giaevim googleIapAppEngineVersionIamMemberAttributes) Service() terra.StringValue {
	return terra.ReferenceAsString(giaevim.ref.Append("service"))
}

// VersionId returns a reference to field version_id of google_iap_app_engine_version_iam_member.
func (giaevim googleIapAppEngineVersionIamMemberAttributes) VersionId() terra.StringValue {
	return terra.ReferenceAsString(giaevim.ref.Append("version_id"))
}

func (giaevim googleIapAppEngineVersionIamMemberAttributes) Condition() terra.ListValue[ConditionAttributes] {
	return terra.ReferenceAsList[ConditionAttributes](giaevim.ref.Append("condition"))
}

type googleIapAppEngineVersionIamMemberState struct {
	AppId     string           `json:"app_id"`
	Etag      string           `json:"etag"`
	Id        string           `json:"id"`
	Member    string           `json:"member"`
	Project   string           `json:"project"`
	Role      string           `json:"role"`
	Service   string           `json:"service"`
	VersionId string           `json:"version_id"`
	Condition []ConditionState `json:"condition"`
}
