// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	mediajob "github.com/golingon/terraproviders/azurerm/3.49.0/mediajob"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewMediaJob(name string, args MediaJobArgs) *MediaJob {
	return &MediaJob{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MediaJob)(nil)

type MediaJob struct {
	Name  string
	Args  MediaJobArgs
	state *mediaJobState
}

func (mj *MediaJob) Type() string {
	return "azurerm_media_job"
}

func (mj *MediaJob) LocalName() string {
	return mj.Name
}

func (mj *MediaJob) Configuration() interface{} {
	return mj.Args
}

func (mj *MediaJob) Attributes() mediaJobAttributes {
	return mediaJobAttributes{ref: terra.ReferenceResource(mj)}
}

func (mj *MediaJob) ImportState(av io.Reader) error {
	mj.state = &mediaJobState{}
	if err := json.NewDecoder(av).Decode(mj.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mj.Type(), mj.LocalName(), err)
	}
	return nil
}

func (mj *MediaJob) State() (*mediaJobState, bool) {
	return mj.state, mj.state != nil
}

func (mj *MediaJob) StateMust() *mediaJobState {
	if mj.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mj.Type(), mj.LocalName()))
	}
	return mj.state
}

func (mj *MediaJob) DependOn() terra.Reference {
	return terra.ReferenceResource(mj)
}

type MediaJobArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MediaServicesAccountName: string, required
	MediaServicesAccountName terra.StringValue `hcl:"media_services_account_name,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Priority: string, optional
	Priority terra.StringValue `hcl:"priority,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// TransformName: string, required
	TransformName terra.StringValue `hcl:"transform_name,attr" validate:"required"`
	// InputAsset: required
	InputAsset *mediajob.InputAsset `hcl:"input_asset,block" validate:"required"`
	// OutputAsset: min=1
	OutputAsset []mediajob.OutputAsset `hcl:"output_asset,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *mediajob.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that MediaJob depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type mediaJobAttributes struct {
	ref terra.Reference
}

func (mj mediaJobAttributes) Description() terra.StringValue {
	return terra.ReferenceString(mj.ref.Append("description"))
}

func (mj mediaJobAttributes) Id() terra.StringValue {
	return terra.ReferenceString(mj.ref.Append("id"))
}

func (mj mediaJobAttributes) MediaServicesAccountName() terra.StringValue {
	return terra.ReferenceString(mj.ref.Append("media_services_account_name"))
}

func (mj mediaJobAttributes) Name() terra.StringValue {
	return terra.ReferenceString(mj.ref.Append("name"))
}

func (mj mediaJobAttributes) Priority() terra.StringValue {
	return terra.ReferenceString(mj.ref.Append("priority"))
}

func (mj mediaJobAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(mj.ref.Append("resource_group_name"))
}

func (mj mediaJobAttributes) TransformName() terra.StringValue {
	return terra.ReferenceString(mj.ref.Append("transform_name"))
}

func (mj mediaJobAttributes) InputAsset() terra.ListValue[mediajob.InputAssetAttributes] {
	return terra.ReferenceList[mediajob.InputAssetAttributes](mj.ref.Append("input_asset"))
}

func (mj mediaJobAttributes) OutputAsset() terra.ListValue[mediajob.OutputAssetAttributes] {
	return terra.ReferenceList[mediajob.OutputAssetAttributes](mj.ref.Append("output_asset"))
}

func (mj mediaJobAttributes) Timeouts() mediajob.TimeoutsAttributes {
	return terra.ReferenceSingle[mediajob.TimeoutsAttributes](mj.ref.Append("timeouts"))
}

type mediaJobState struct {
	Description              string                      `json:"description"`
	Id                       string                      `json:"id"`
	MediaServicesAccountName string                      `json:"media_services_account_name"`
	Name                     string                      `json:"name"`
	Priority                 string                      `json:"priority"`
	ResourceGroupName        string                      `json:"resource_group_name"`
	TransformName            string                      `json:"transform_name"`
	InputAsset               []mediajob.InputAssetState  `json:"input_asset"`
	OutputAsset              []mediajob.OutputAssetState `json:"output_asset"`
	Timeouts                 *mediajob.TimeoutsState     `json:"timeouts"`
}
