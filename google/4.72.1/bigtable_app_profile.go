// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	bigtableappprofile "github.com/golingon/terraproviders/google/4.72.1/bigtableappprofile"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBigtableAppProfile creates a new instance of [BigtableAppProfile].
func NewBigtableAppProfile(name string, args BigtableAppProfileArgs) *BigtableAppProfile {
	return &BigtableAppProfile{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BigtableAppProfile)(nil)

// BigtableAppProfile represents the Terraform resource google_bigtable_app_profile.
type BigtableAppProfile struct {
	Name      string
	Args      BigtableAppProfileArgs
	state     *bigtableAppProfileState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BigtableAppProfile].
func (bap *BigtableAppProfile) Type() string {
	return "google_bigtable_app_profile"
}

// LocalName returns the local name for [BigtableAppProfile].
func (bap *BigtableAppProfile) LocalName() string {
	return bap.Name
}

// Configuration returns the configuration (args) for [BigtableAppProfile].
func (bap *BigtableAppProfile) Configuration() interface{} {
	return bap.Args
}

// DependOn is used for other resources to depend on [BigtableAppProfile].
func (bap *BigtableAppProfile) DependOn() terra.Reference {
	return terra.ReferenceResource(bap)
}

// Dependencies returns the list of resources [BigtableAppProfile] depends_on.
func (bap *BigtableAppProfile) Dependencies() terra.Dependencies {
	return bap.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BigtableAppProfile].
func (bap *BigtableAppProfile) LifecycleManagement() *terra.Lifecycle {
	return bap.Lifecycle
}

// Attributes returns the attributes for [BigtableAppProfile].
func (bap *BigtableAppProfile) Attributes() bigtableAppProfileAttributes {
	return bigtableAppProfileAttributes{ref: terra.ReferenceResource(bap)}
}

// ImportState imports the given attribute values into [BigtableAppProfile]'s state.
func (bap *BigtableAppProfile) ImportState(av io.Reader) error {
	bap.state = &bigtableAppProfileState{}
	if err := json.NewDecoder(av).Decode(bap.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bap.Type(), bap.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BigtableAppProfile] has state.
func (bap *BigtableAppProfile) State() (*bigtableAppProfileState, bool) {
	return bap.state, bap.state != nil
}

// StateMust returns the state for [BigtableAppProfile]. Panics if the state is nil.
func (bap *BigtableAppProfile) StateMust() *bigtableAppProfileState {
	if bap.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bap.Type(), bap.LocalName()))
	}
	return bap.state
}

// BigtableAppProfileArgs contains the configurations for google_bigtable_app_profile.
type BigtableAppProfileArgs struct {
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
	SingleClusterRouting *bigtableappprofile.SingleClusterRouting `hcl:"single_cluster_routing,block"`
	// Timeouts: optional
	Timeouts *bigtableappprofile.Timeouts `hcl:"timeouts,block"`
}
type bigtableAppProfileAttributes struct {
	ref terra.Reference
}

// AppProfileId returns a reference to field app_profile_id of google_bigtable_app_profile.
func (bap bigtableAppProfileAttributes) AppProfileId() terra.StringValue {
	return terra.ReferenceAsString(bap.ref.Append("app_profile_id"))
}

// Description returns a reference to field description of google_bigtable_app_profile.
func (bap bigtableAppProfileAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(bap.ref.Append("description"))
}

// Id returns a reference to field id of google_bigtable_app_profile.
func (bap bigtableAppProfileAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bap.ref.Append("id"))
}

// IgnoreWarnings returns a reference to field ignore_warnings of google_bigtable_app_profile.
func (bap bigtableAppProfileAttributes) IgnoreWarnings() terra.BoolValue {
	return terra.ReferenceAsBool(bap.ref.Append("ignore_warnings"))
}

// Instance returns a reference to field instance of google_bigtable_app_profile.
func (bap bigtableAppProfileAttributes) Instance() terra.StringValue {
	return terra.ReferenceAsString(bap.ref.Append("instance"))
}

// MultiClusterRoutingClusterIds returns a reference to field multi_cluster_routing_cluster_ids of google_bigtable_app_profile.
func (bap bigtableAppProfileAttributes) MultiClusterRoutingClusterIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](bap.ref.Append("multi_cluster_routing_cluster_ids"))
}

// MultiClusterRoutingUseAny returns a reference to field multi_cluster_routing_use_any of google_bigtable_app_profile.
func (bap bigtableAppProfileAttributes) MultiClusterRoutingUseAny() terra.BoolValue {
	return terra.ReferenceAsBool(bap.ref.Append("multi_cluster_routing_use_any"))
}

// Name returns a reference to field name of google_bigtable_app_profile.
func (bap bigtableAppProfileAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(bap.ref.Append("name"))
}

// Project returns a reference to field project of google_bigtable_app_profile.
func (bap bigtableAppProfileAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(bap.ref.Append("project"))
}

func (bap bigtableAppProfileAttributes) SingleClusterRouting() terra.ListValue[bigtableappprofile.SingleClusterRoutingAttributes] {
	return terra.ReferenceAsList[bigtableappprofile.SingleClusterRoutingAttributes](bap.ref.Append("single_cluster_routing"))
}

func (bap bigtableAppProfileAttributes) Timeouts() bigtableappprofile.TimeoutsAttributes {
	return terra.ReferenceAsSingle[bigtableappprofile.TimeoutsAttributes](bap.ref.Append("timeouts"))
}

type bigtableAppProfileState struct {
	AppProfileId                  string                                         `json:"app_profile_id"`
	Description                   string                                         `json:"description"`
	Id                            string                                         `json:"id"`
	IgnoreWarnings                bool                                           `json:"ignore_warnings"`
	Instance                      string                                         `json:"instance"`
	MultiClusterRoutingClusterIds []string                                       `json:"multi_cluster_routing_cluster_ids"`
	MultiClusterRoutingUseAny     bool                                           `json:"multi_cluster_routing_use_any"`
	Name                          string                                         `json:"name"`
	Project                       string                                         `json:"project"`
	SingleClusterRouting          []bigtableappprofile.SingleClusterRoutingState `json:"single_cluster_routing"`
	Timeouts                      *bigtableappprofile.TimeoutsState              `json:"timeouts"`
}
