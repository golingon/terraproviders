// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	cloudassetfolderfeed "github.com/golingon/terraproviders/googlebeta/4.71.0/cloudassetfolderfeed"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudAssetFolderFeed creates a new instance of [CloudAssetFolderFeed].
func NewCloudAssetFolderFeed(name string, args CloudAssetFolderFeedArgs) *CloudAssetFolderFeed {
	return &CloudAssetFolderFeed{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudAssetFolderFeed)(nil)

// CloudAssetFolderFeed represents the Terraform resource google_cloud_asset_folder_feed.
type CloudAssetFolderFeed struct {
	Name      string
	Args      CloudAssetFolderFeedArgs
	state     *cloudAssetFolderFeedState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CloudAssetFolderFeed].
func (caff *CloudAssetFolderFeed) Type() string {
	return "google_cloud_asset_folder_feed"
}

// LocalName returns the local name for [CloudAssetFolderFeed].
func (caff *CloudAssetFolderFeed) LocalName() string {
	return caff.Name
}

// Configuration returns the configuration (args) for [CloudAssetFolderFeed].
func (caff *CloudAssetFolderFeed) Configuration() interface{} {
	return caff.Args
}

// DependOn is used for other resources to depend on [CloudAssetFolderFeed].
func (caff *CloudAssetFolderFeed) DependOn() terra.Reference {
	return terra.ReferenceResource(caff)
}

// Dependencies returns the list of resources [CloudAssetFolderFeed] depends_on.
func (caff *CloudAssetFolderFeed) Dependencies() terra.Dependencies {
	return caff.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CloudAssetFolderFeed].
func (caff *CloudAssetFolderFeed) LifecycleManagement() *terra.Lifecycle {
	return caff.Lifecycle
}

// Attributes returns the attributes for [CloudAssetFolderFeed].
func (caff *CloudAssetFolderFeed) Attributes() cloudAssetFolderFeedAttributes {
	return cloudAssetFolderFeedAttributes{ref: terra.ReferenceResource(caff)}
}

// ImportState imports the given attribute values into [CloudAssetFolderFeed]'s state.
func (caff *CloudAssetFolderFeed) ImportState(av io.Reader) error {
	caff.state = &cloudAssetFolderFeedState{}
	if err := json.NewDecoder(av).Decode(caff.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", caff.Type(), caff.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CloudAssetFolderFeed] has state.
func (caff *CloudAssetFolderFeed) State() (*cloudAssetFolderFeedState, bool) {
	return caff.state, caff.state != nil
}

// StateMust returns the state for [CloudAssetFolderFeed]. Panics if the state is nil.
func (caff *CloudAssetFolderFeed) StateMust() *cloudAssetFolderFeedState {
	if caff.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", caff.Type(), caff.LocalName()))
	}
	return caff.state
}

// CloudAssetFolderFeedArgs contains the configurations for google_cloud_asset_folder_feed.
type CloudAssetFolderFeedArgs struct {
	// AssetNames: list of string, optional
	AssetNames terra.ListValue[terra.StringValue] `hcl:"asset_names,attr"`
	// AssetTypes: list of string, optional
	AssetTypes terra.ListValue[terra.StringValue] `hcl:"asset_types,attr"`
	// BillingProject: string, required
	BillingProject terra.StringValue `hcl:"billing_project,attr" validate:"required"`
	// ContentType: string, optional
	ContentType terra.StringValue `hcl:"content_type,attr"`
	// FeedId: string, required
	FeedId terra.StringValue `hcl:"feed_id,attr" validate:"required"`
	// Folder: string, required
	Folder terra.StringValue `hcl:"folder,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Condition: optional
	Condition *cloudassetfolderfeed.Condition `hcl:"condition,block"`
	// FeedOutputConfig: required
	FeedOutputConfig *cloudassetfolderfeed.FeedOutputConfig `hcl:"feed_output_config,block" validate:"required"`
	// Timeouts: optional
	Timeouts *cloudassetfolderfeed.Timeouts `hcl:"timeouts,block"`
}
type cloudAssetFolderFeedAttributes struct {
	ref terra.Reference
}

// AssetNames returns a reference to field asset_names of google_cloud_asset_folder_feed.
func (caff cloudAssetFolderFeedAttributes) AssetNames() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](caff.ref.Append("asset_names"))
}

// AssetTypes returns a reference to field asset_types of google_cloud_asset_folder_feed.
func (caff cloudAssetFolderFeedAttributes) AssetTypes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](caff.ref.Append("asset_types"))
}

// BillingProject returns a reference to field billing_project of google_cloud_asset_folder_feed.
func (caff cloudAssetFolderFeedAttributes) BillingProject() terra.StringValue {
	return terra.ReferenceAsString(caff.ref.Append("billing_project"))
}

// ContentType returns a reference to field content_type of google_cloud_asset_folder_feed.
func (caff cloudAssetFolderFeedAttributes) ContentType() terra.StringValue {
	return terra.ReferenceAsString(caff.ref.Append("content_type"))
}

// FeedId returns a reference to field feed_id of google_cloud_asset_folder_feed.
func (caff cloudAssetFolderFeedAttributes) FeedId() terra.StringValue {
	return terra.ReferenceAsString(caff.ref.Append("feed_id"))
}

// Folder returns a reference to field folder of google_cloud_asset_folder_feed.
func (caff cloudAssetFolderFeedAttributes) Folder() terra.StringValue {
	return terra.ReferenceAsString(caff.ref.Append("folder"))
}

// FolderId returns a reference to field folder_id of google_cloud_asset_folder_feed.
func (caff cloudAssetFolderFeedAttributes) FolderId() terra.StringValue {
	return terra.ReferenceAsString(caff.ref.Append("folder_id"))
}

// Id returns a reference to field id of google_cloud_asset_folder_feed.
func (caff cloudAssetFolderFeedAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(caff.ref.Append("id"))
}

// Name returns a reference to field name of google_cloud_asset_folder_feed.
func (caff cloudAssetFolderFeedAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(caff.ref.Append("name"))
}

func (caff cloudAssetFolderFeedAttributes) Condition() terra.ListValue[cloudassetfolderfeed.ConditionAttributes] {
	return terra.ReferenceAsList[cloudassetfolderfeed.ConditionAttributes](caff.ref.Append("condition"))
}

func (caff cloudAssetFolderFeedAttributes) FeedOutputConfig() terra.ListValue[cloudassetfolderfeed.FeedOutputConfigAttributes] {
	return terra.ReferenceAsList[cloudassetfolderfeed.FeedOutputConfigAttributes](caff.ref.Append("feed_output_config"))
}

func (caff cloudAssetFolderFeedAttributes) Timeouts() cloudassetfolderfeed.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cloudassetfolderfeed.TimeoutsAttributes](caff.ref.Append("timeouts"))
}

type cloudAssetFolderFeedState struct {
	AssetNames       []string                                     `json:"asset_names"`
	AssetTypes       []string                                     `json:"asset_types"`
	BillingProject   string                                       `json:"billing_project"`
	ContentType      string                                       `json:"content_type"`
	FeedId           string                                       `json:"feed_id"`
	Folder           string                                       `json:"folder"`
	FolderId         string                                       `json:"folder_id"`
	Id               string                                       `json:"id"`
	Name             string                                       `json:"name"`
	Condition        []cloudassetfolderfeed.ConditionState        `json:"condition"`
	FeedOutputConfig []cloudassetfolderfeed.FeedOutputConfigState `json:"feed_output_config"`
	Timeouts         *cloudassetfolderfeed.TimeoutsState          `json:"timeouts"`
}
