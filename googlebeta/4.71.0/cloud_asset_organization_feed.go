// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	cloudassetorganizationfeed "github.com/golingon/terraproviders/googlebeta/4.71.0/cloudassetorganizationfeed"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudAssetOrganizationFeed creates a new instance of [CloudAssetOrganizationFeed].
func NewCloudAssetOrganizationFeed(name string, args CloudAssetOrganizationFeedArgs) *CloudAssetOrganizationFeed {
	return &CloudAssetOrganizationFeed{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudAssetOrganizationFeed)(nil)

// CloudAssetOrganizationFeed represents the Terraform resource google_cloud_asset_organization_feed.
type CloudAssetOrganizationFeed struct {
	Name      string
	Args      CloudAssetOrganizationFeedArgs
	state     *cloudAssetOrganizationFeedState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CloudAssetOrganizationFeed].
func (caof *CloudAssetOrganizationFeed) Type() string {
	return "google_cloud_asset_organization_feed"
}

// LocalName returns the local name for [CloudAssetOrganizationFeed].
func (caof *CloudAssetOrganizationFeed) LocalName() string {
	return caof.Name
}

// Configuration returns the configuration (args) for [CloudAssetOrganizationFeed].
func (caof *CloudAssetOrganizationFeed) Configuration() interface{} {
	return caof.Args
}

// DependOn is used for other resources to depend on [CloudAssetOrganizationFeed].
func (caof *CloudAssetOrganizationFeed) DependOn() terra.Reference {
	return terra.ReferenceResource(caof)
}

// Dependencies returns the list of resources [CloudAssetOrganizationFeed] depends_on.
func (caof *CloudAssetOrganizationFeed) Dependencies() terra.Dependencies {
	return caof.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CloudAssetOrganizationFeed].
func (caof *CloudAssetOrganizationFeed) LifecycleManagement() *terra.Lifecycle {
	return caof.Lifecycle
}

// Attributes returns the attributes for [CloudAssetOrganizationFeed].
func (caof *CloudAssetOrganizationFeed) Attributes() cloudAssetOrganizationFeedAttributes {
	return cloudAssetOrganizationFeedAttributes{ref: terra.ReferenceResource(caof)}
}

// ImportState imports the given attribute values into [CloudAssetOrganizationFeed]'s state.
func (caof *CloudAssetOrganizationFeed) ImportState(av io.Reader) error {
	caof.state = &cloudAssetOrganizationFeedState{}
	if err := json.NewDecoder(av).Decode(caof.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", caof.Type(), caof.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CloudAssetOrganizationFeed] has state.
func (caof *CloudAssetOrganizationFeed) State() (*cloudAssetOrganizationFeedState, bool) {
	return caof.state, caof.state != nil
}

// StateMust returns the state for [CloudAssetOrganizationFeed]. Panics if the state is nil.
func (caof *CloudAssetOrganizationFeed) StateMust() *cloudAssetOrganizationFeedState {
	if caof.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", caof.Type(), caof.LocalName()))
	}
	return caof.state
}

// CloudAssetOrganizationFeedArgs contains the configurations for google_cloud_asset_organization_feed.
type CloudAssetOrganizationFeedArgs struct {
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
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// OrgId: string, required
	OrgId terra.StringValue `hcl:"org_id,attr" validate:"required"`
	// Condition: optional
	Condition *cloudassetorganizationfeed.Condition `hcl:"condition,block"`
	// FeedOutputConfig: required
	FeedOutputConfig *cloudassetorganizationfeed.FeedOutputConfig `hcl:"feed_output_config,block" validate:"required"`
	// Timeouts: optional
	Timeouts *cloudassetorganizationfeed.Timeouts `hcl:"timeouts,block"`
}
type cloudAssetOrganizationFeedAttributes struct {
	ref terra.Reference
}

// AssetNames returns a reference to field asset_names of google_cloud_asset_organization_feed.
func (caof cloudAssetOrganizationFeedAttributes) AssetNames() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](caof.ref.Append("asset_names"))
}

// AssetTypes returns a reference to field asset_types of google_cloud_asset_organization_feed.
func (caof cloudAssetOrganizationFeedAttributes) AssetTypes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](caof.ref.Append("asset_types"))
}

// BillingProject returns a reference to field billing_project of google_cloud_asset_organization_feed.
func (caof cloudAssetOrganizationFeedAttributes) BillingProject() terra.StringValue {
	return terra.ReferenceAsString(caof.ref.Append("billing_project"))
}

// ContentType returns a reference to field content_type of google_cloud_asset_organization_feed.
func (caof cloudAssetOrganizationFeedAttributes) ContentType() terra.StringValue {
	return terra.ReferenceAsString(caof.ref.Append("content_type"))
}

// FeedId returns a reference to field feed_id of google_cloud_asset_organization_feed.
func (caof cloudAssetOrganizationFeedAttributes) FeedId() terra.StringValue {
	return terra.ReferenceAsString(caof.ref.Append("feed_id"))
}

// Id returns a reference to field id of google_cloud_asset_organization_feed.
func (caof cloudAssetOrganizationFeedAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(caof.ref.Append("id"))
}

// Name returns a reference to field name of google_cloud_asset_organization_feed.
func (caof cloudAssetOrganizationFeedAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(caof.ref.Append("name"))
}

// OrgId returns a reference to field org_id of google_cloud_asset_organization_feed.
func (caof cloudAssetOrganizationFeedAttributes) OrgId() terra.StringValue {
	return terra.ReferenceAsString(caof.ref.Append("org_id"))
}

func (caof cloudAssetOrganizationFeedAttributes) Condition() terra.ListValue[cloudassetorganizationfeed.ConditionAttributes] {
	return terra.ReferenceAsList[cloudassetorganizationfeed.ConditionAttributes](caof.ref.Append("condition"))
}

func (caof cloudAssetOrganizationFeedAttributes) FeedOutputConfig() terra.ListValue[cloudassetorganizationfeed.FeedOutputConfigAttributes] {
	return terra.ReferenceAsList[cloudassetorganizationfeed.FeedOutputConfigAttributes](caof.ref.Append("feed_output_config"))
}

func (caof cloudAssetOrganizationFeedAttributes) Timeouts() cloudassetorganizationfeed.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cloudassetorganizationfeed.TimeoutsAttributes](caof.ref.Append("timeouts"))
}

type cloudAssetOrganizationFeedState struct {
	AssetNames       []string                                           `json:"asset_names"`
	AssetTypes       []string                                           `json:"asset_types"`
	BillingProject   string                                             `json:"billing_project"`
	ContentType      string                                             `json:"content_type"`
	FeedId           string                                             `json:"feed_id"`
	Id               string                                             `json:"id"`
	Name             string                                             `json:"name"`
	OrgId            string                                             `json:"org_id"`
	Condition        []cloudassetorganizationfeed.ConditionState        `json:"condition"`
	FeedOutputConfig []cloudassetorganizationfeed.FeedOutputConfigState `json:"feed_output_config"`
	Timeouts         *cloudassetorganizationfeed.TimeoutsState          `json:"timeouts"`
}
