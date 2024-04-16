// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_bigtable_app_profile

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

// Resource represents the Terraform resource google_bigtable_app_profile.
type Resource struct {
	Name      string
	Args      Args
	state     *googleBigtableAppProfileState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gbap *Resource) Type() string {
	return "google_bigtable_app_profile"
}

// LocalName returns the local name for [Resource].
func (gbap *Resource) LocalName() string {
	return gbap.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gbap *Resource) Configuration() interface{} {
	return gbap.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gbap *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gbap)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gbap *Resource) Dependencies() terra.Dependencies {
	return gbap.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gbap *Resource) LifecycleManagement() *terra.Lifecycle {
	return gbap.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gbap *Resource) Attributes() googleBigtableAppProfileAttributes {
	return googleBigtableAppProfileAttributes{ref: terra.ReferenceResource(gbap)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gbap *Resource) ImportState(state io.Reader) error {
	gbap.state = &googleBigtableAppProfileState{}
	if err := json.NewDecoder(state).Decode(gbap.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gbap.Type(), gbap.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gbap *Resource) State() (*googleBigtableAppProfileState, bool) {
	return gbap.state, gbap.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gbap *Resource) StateMust() *googleBigtableAppProfileState {
	if gbap.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gbap.Type(), gbap.LocalName()))
	}
	return gbap.state
}

// Args contains the configurations for google_bigtable_app_profile.
type Args struct {
	// AppProfileId: string, required
	AppProfileId terra.StringValue `hcl:"app_profile_id,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IgnoreWarnings: bool, optional
	IgnoreWarnings terra.BoolValue `hcl:"ignore_warnings,attr"`
	// Instance: string, optional
	Instance terra.StringValue `hcl:"instance,attr"`
	// MultiClusterRoutingClusterIds: list of string, optional
	MultiClusterRoutingClusterIds terra.ListValue[terra.StringValue] `hcl:"multi_cluster_routing_cluster_ids,attr"`
	// MultiClusterRoutingUseAny: bool, optional
	MultiClusterRoutingUseAny terra.BoolValue `hcl:"multi_cluster_routing_use_any,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// SingleClusterRouting: optional
	SingleClusterRouting *SingleClusterRouting `hcl:"single_cluster_routing,block"`
	// StandardIsolation: optional
	StandardIsolation *StandardIsolation `hcl:"standard_isolation,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googleBigtableAppProfileAttributes struct {
	ref terra.Reference
}

// AppProfileId returns a reference to field app_profile_id of google_bigtable_app_profile.
func (gbap googleBigtableAppProfileAttributes) AppProfileId() terra.StringValue {
	return terra.ReferenceAsString(gbap.ref.Append("app_profile_id"))
}

// Description returns a reference to field description of google_bigtable_app_profile.
func (gbap googleBigtableAppProfileAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(gbap.ref.Append("description"))
}

// Id returns a reference to field id of google_bigtable_app_profile.
func (gbap googleBigtableAppProfileAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gbap.ref.Append("id"))
}

// IgnoreWarnings returns a reference to field ignore_warnings of google_bigtable_app_profile.
func (gbap googleBigtableAppProfileAttributes) IgnoreWarnings() terra.BoolValue {
	return terra.ReferenceAsBool(gbap.ref.Append("ignore_warnings"))
}

// Instance returns a reference to field instance of google_bigtable_app_profile.
func (gbap googleBigtableAppProfileAttributes) Instance() terra.StringValue {
	return terra.ReferenceAsString(gbap.ref.Append("instance"))
}

// MultiClusterRoutingClusterIds returns a reference to field multi_cluster_routing_cluster_ids of google_bigtable_app_profile.
func (gbap googleBigtableAppProfileAttributes) MultiClusterRoutingClusterIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](gbap.ref.Append("multi_cluster_routing_cluster_ids"))
}

// MultiClusterRoutingUseAny returns a reference to field multi_cluster_routing_use_any of google_bigtable_app_profile.
func (gbap googleBigtableAppProfileAttributes) MultiClusterRoutingUseAny() terra.BoolValue {
	return terra.ReferenceAsBool(gbap.ref.Append("multi_cluster_routing_use_any"))
}

// Name returns a reference to field name of google_bigtable_app_profile.
func (gbap googleBigtableAppProfileAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gbap.ref.Append("name"))
}

// Project returns a reference to field project of google_bigtable_app_profile.
func (gbap googleBigtableAppProfileAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gbap.ref.Append("project"))
}

func (gbap googleBigtableAppProfileAttributes) SingleClusterRouting() terra.ListValue[SingleClusterRoutingAttributes] {
	return terra.ReferenceAsList[SingleClusterRoutingAttributes](gbap.ref.Append("single_cluster_routing"))
}

func (gbap googleBigtableAppProfileAttributes) StandardIsolation() terra.ListValue[StandardIsolationAttributes] {
	return terra.ReferenceAsList[StandardIsolationAttributes](gbap.ref.Append("standard_isolation"))
}

func (gbap googleBigtableAppProfileAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gbap.ref.Append("timeouts"))
}

type googleBigtableAppProfileState struct {
	AppProfileId                  string                      `json:"app_profile_id"`
	Description                   string                      `json:"description"`
	Id                            string                      `json:"id"`
	IgnoreWarnings                bool                        `json:"ignore_warnings"`
	Instance                      string                      `json:"instance"`
	MultiClusterRoutingClusterIds []string                    `json:"multi_cluster_routing_cluster_ids"`
	MultiClusterRoutingUseAny     bool                        `json:"multi_cluster_routing_use_any"`
	Name                          string                      `json:"name"`
	Project                       string                      `json:"project"`
	SingleClusterRouting          []SingleClusterRoutingState `json:"single_cluster_routing"`
	StandardIsolation             []StandardIsolationState    `json:"standard_isolation"`
	Timeouts                      *TimeoutsState              `json:"timeouts"`
}
