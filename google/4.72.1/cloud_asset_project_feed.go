// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	cloudassetprojectfeed "github.com/golingon/terraproviders/google/4.72.1/cloudassetprojectfeed"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudAssetProjectFeed creates a new instance of [CloudAssetProjectFeed].
func NewCloudAssetProjectFeed(name string, args CloudAssetProjectFeedArgs) *CloudAssetProjectFeed {
	return &CloudAssetProjectFeed{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudAssetProjectFeed)(nil)

// CloudAssetProjectFeed represents the Terraform resource google_cloud_asset_project_feed.
type CloudAssetProjectFeed struct {
	Name      string
	Args      CloudAssetProjectFeedArgs
	state     *cloudAssetProjectFeedState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CloudAssetProjectFeed].
func (capf *CloudAssetProjectFeed) Type() string {
	return "google_cloud_asset_project_feed"
}

// LocalName returns the local name for [CloudAssetProjectFeed].
func (capf *CloudAssetProjectFeed) LocalName() string {
	return capf.Name
}

// Configuration returns the configuration (args) for [CloudAssetProjectFeed].
func (capf *CloudAssetProjectFeed) Configuration() interface{} {
	return capf.Args
}

// DependOn is used for other resources to depend on [CloudAssetProjectFeed].
func (capf *CloudAssetProjectFeed) DependOn() terra.Reference {
	return terra.ReferenceResource(capf)
}

// Dependencies returns the list of resources [CloudAssetProjectFeed] depends_on.
func (capf *CloudAssetProjectFeed) Dependencies() terra.Dependencies {
	return capf.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CloudAssetProjectFeed].
func (capf *CloudAssetProjectFeed) LifecycleManagement() *terra.Lifecycle {
	return capf.Lifecycle
}

// Attributes returns the attributes for [CloudAssetProjectFeed].
func (capf *CloudAssetProjectFeed) Attributes() cloudAssetProjectFeedAttributes {
	return cloudAssetProjectFeedAttributes{ref: terra.ReferenceResource(capf)}
}

// ImportState imports the given attribute values into [CloudAssetProjectFeed]'s state.
func (capf *CloudAssetProjectFeed) ImportState(av io.Reader) error {
	capf.state = &cloudAssetProjectFeedState{}
	if err := json.NewDecoder(av).Decode(capf.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", capf.Type(), capf.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CloudAssetProjectFeed] has state.
func (capf *CloudAssetProjectFeed) State() (*cloudAssetProjectFeedState, bool) {
	return capf.state, capf.state != nil
}

// StateMust returns the state for [CloudAssetProjectFeed]. Panics if the state is nil.
func (capf *CloudAssetProjectFeed) StateMust() *cloudAssetProjectFeedState {
	if capf.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", capf.Type(), capf.LocalName()))
	}
	return capf.state
}

// CloudAssetProjectFeedArgs contains the configurations for google_cloud_asset_project_feed.
type CloudAssetProjectFeedArgs struct {
	// AssetNames: list of string, optional
	AssetNames terra.ListValue[terra.StringValue] `hcl:"asset_names,attr"`
	// AssetTypes: list of string, optional
	AssetTypes terra.ListValue[terra.StringValue] `hcl:"asset_types,attr"`
	// BillingProject: string, optional
	BillingProject terra.StringValue `hcl:"billing_project,attr"`
	// ContentType: string, optional
	ContentType terra.StringValue `hcl:"content_type,attr"`
	// FeedId: string, required
	FeedId terra.StringValue `hcl:"feed_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Condition: optional
	Condition *cloudassetprojectfeed.Condition `hcl:"condition,block"`
	// FeedOutputConfig: required
	FeedOutputConfig *cloudassetprojectfeed.FeedOutputConfig `hcl:"feed_output_config,block" validate:"required"`
	// Timeouts: optional
	Timeouts *cloudassetprojectfeed.Timeouts `hcl:"timeouts,block"`
}
type cloudAssetProjectFeedAttributes struct {
	ref terra.Reference
}

// AssetNames returns a reference to field asset_names of google_cloud_asset_project_feed.
func (capf cloudAssetProjectFeedAttributes) AssetNames() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](capf.ref.Append("asset_names"))
}

// AssetTypes returns a reference to field asset_types of google_cloud_asset_project_feed.
func (capf cloudAssetProjectFeedAttributes) AssetTypes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](capf.ref.Append("asset_types"))
}

// BillingProject returns a reference to field billing_project of google_cloud_asset_project_feed.
func (capf cloudAssetProjectFeedAttributes) BillingProject() terra.StringValue {
	return terra.ReferenceAsString(capf.ref.Append("billing_project"))
}

// ContentType returns a reference to field content_type of google_cloud_asset_project_feed.
func (capf cloudAssetProjectFeedAttributes) ContentType() terra.StringValue {
	return terra.ReferenceAsString(capf.ref.Append("content_type"))
}

// FeedId returns a reference to field feed_id of google_cloud_asset_project_feed.
func (capf cloudAssetProjectFeedAttributes) FeedId() terra.StringValue {
	return terra.ReferenceAsString(capf.ref.Append("feed_id"))
}

// Id returns a reference to field id of google_cloud_asset_project_feed.
func (capf cloudAssetProjectFeedAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(capf.ref.Append("id"))
}

// Name returns a reference to field name of google_cloud_asset_project_feed.
func (capf cloudAssetProjectFeedAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(capf.ref.Append("name"))
}

// Project returns a reference to field project of google_cloud_asset_project_feed.
func (capf cloudAssetProjectFeedAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(capf.ref.Append("project"))
}

func (capf cloudAssetProjectFeedAttributes) Condition() terra.ListValue[cloudassetprojectfeed.ConditionAttributes] {
	return terra.ReferenceAsList[cloudassetprojectfeed.ConditionAttributes](capf.ref.Append("condition"))
}

func (capf cloudAssetProjectFeedAttributes) FeedOutputConfig() terra.ListValue[cloudassetprojectfeed.FeedOutputConfigAttributes] {
	return terra.ReferenceAsList[cloudassetprojectfeed.FeedOutputConfigAttributes](capf.ref.Append("feed_output_config"))
}

func (capf cloudAssetProjectFeedAttributes) Timeouts() cloudassetprojectfeed.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cloudassetprojectfeed.TimeoutsAttributes](capf.ref.Append("timeouts"))
}

type cloudAssetProjectFeedState struct {
	AssetNames       []string                                      `json:"asset_names"`
	AssetTypes       []string                                      `json:"asset_types"`
	BillingProject   string                                        `json:"billing_project"`
	ContentType      string                                        `json:"content_type"`
	FeedId           string                                        `json:"feed_id"`
	Id               string                                        `json:"id"`
	Name             string                                        `json:"name"`
	Project          string                                        `json:"project"`
	Condition        []cloudassetprojectfeed.ConditionState        `json:"condition"`
	FeedOutputConfig []cloudassetprojectfeed.FeedOutputConfigState `json:"feed_output_config"`
	Timeouts         *cloudassetprojectfeed.TimeoutsState          `json:"timeouts"`
}
