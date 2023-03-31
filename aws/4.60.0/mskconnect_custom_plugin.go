// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	mskconnectcustomplugin "github.com/golingon/terraproviders/aws/4.60.0/mskconnectcustomplugin"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMskconnectCustomPlugin creates a new instance of [MskconnectCustomPlugin].
func NewMskconnectCustomPlugin(name string, args MskconnectCustomPluginArgs) *MskconnectCustomPlugin {
	return &MskconnectCustomPlugin{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MskconnectCustomPlugin)(nil)

// MskconnectCustomPlugin represents the Terraform resource aws_mskconnect_custom_plugin.
type MskconnectCustomPlugin struct {
	Name      string
	Args      MskconnectCustomPluginArgs
	state     *mskconnectCustomPluginState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MskconnectCustomPlugin].
func (mcp *MskconnectCustomPlugin) Type() string {
	return "aws_mskconnect_custom_plugin"
}

// LocalName returns the local name for [MskconnectCustomPlugin].
func (mcp *MskconnectCustomPlugin) LocalName() string {
	return mcp.Name
}

// Configuration returns the configuration (args) for [MskconnectCustomPlugin].
func (mcp *MskconnectCustomPlugin) Configuration() interface{} {
	return mcp.Args
}

// DependOn is used for other resources to depend on [MskconnectCustomPlugin].
func (mcp *MskconnectCustomPlugin) DependOn() terra.Reference {
	return terra.ReferenceResource(mcp)
}

// Dependencies returns the list of resources [MskconnectCustomPlugin] depends_on.
func (mcp *MskconnectCustomPlugin) Dependencies() terra.Dependencies {
	return mcp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MskconnectCustomPlugin].
func (mcp *MskconnectCustomPlugin) LifecycleManagement() *terra.Lifecycle {
	return mcp.Lifecycle
}

// Attributes returns the attributes for [MskconnectCustomPlugin].
func (mcp *MskconnectCustomPlugin) Attributes() mskconnectCustomPluginAttributes {
	return mskconnectCustomPluginAttributes{ref: terra.ReferenceResource(mcp)}
}

// ImportState imports the given attribute values into [MskconnectCustomPlugin]'s state.
func (mcp *MskconnectCustomPlugin) ImportState(av io.Reader) error {
	mcp.state = &mskconnectCustomPluginState{}
	if err := json.NewDecoder(av).Decode(mcp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mcp.Type(), mcp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MskconnectCustomPlugin] has state.
func (mcp *MskconnectCustomPlugin) State() (*mskconnectCustomPluginState, bool) {
	return mcp.state, mcp.state != nil
}

// StateMust returns the state for [MskconnectCustomPlugin]. Panics if the state is nil.
func (mcp *MskconnectCustomPlugin) StateMust() *mskconnectCustomPluginState {
	if mcp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mcp.Type(), mcp.LocalName()))
	}
	return mcp.state
}

// MskconnectCustomPluginArgs contains the configurations for aws_mskconnect_custom_plugin.
type MskconnectCustomPluginArgs struct {
	// ContentType: string, required
	ContentType terra.StringValue `hcl:"content_type,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Location: required
	Location *mskconnectcustomplugin.Location `hcl:"location,block" validate:"required"`
	// Timeouts: optional
	Timeouts *mskconnectcustomplugin.Timeouts `hcl:"timeouts,block"`
}
type mskconnectCustomPluginAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_mskconnect_custom_plugin.
func (mcp mskconnectCustomPluginAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(mcp.ref.Append("arn"))
}

// ContentType returns a reference to field content_type of aws_mskconnect_custom_plugin.
func (mcp mskconnectCustomPluginAttributes) ContentType() terra.StringValue {
	return terra.ReferenceAsString(mcp.ref.Append("content_type"))
}

// Description returns a reference to field description of aws_mskconnect_custom_plugin.
func (mcp mskconnectCustomPluginAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(mcp.ref.Append("description"))
}

// Id returns a reference to field id of aws_mskconnect_custom_plugin.
func (mcp mskconnectCustomPluginAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mcp.ref.Append("id"))
}

// LatestRevision returns a reference to field latest_revision of aws_mskconnect_custom_plugin.
func (mcp mskconnectCustomPluginAttributes) LatestRevision() terra.NumberValue {
	return terra.ReferenceAsNumber(mcp.ref.Append("latest_revision"))
}

// Name returns a reference to field name of aws_mskconnect_custom_plugin.
func (mcp mskconnectCustomPluginAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mcp.ref.Append("name"))
}

// State returns a reference to field state of aws_mskconnect_custom_plugin.
func (mcp mskconnectCustomPluginAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(mcp.ref.Append("state"))
}

func (mcp mskconnectCustomPluginAttributes) Location() terra.ListValue[mskconnectcustomplugin.LocationAttributes] {
	return terra.ReferenceAsList[mskconnectcustomplugin.LocationAttributes](mcp.ref.Append("location"))
}

func (mcp mskconnectCustomPluginAttributes) Timeouts() mskconnectcustomplugin.TimeoutsAttributes {
	return terra.ReferenceAsSingle[mskconnectcustomplugin.TimeoutsAttributes](mcp.ref.Append("timeouts"))
}

type mskconnectCustomPluginState struct {
	Arn            string                                 `json:"arn"`
	ContentType    string                                 `json:"content_type"`
	Description    string                                 `json:"description"`
	Id             string                                 `json:"id"`
	LatestRevision float64                                `json:"latest_revision"`
	Name           string                                 `json:"name"`
	State          string                                 `json:"state"`
	Location       []mskconnectcustomplugin.LocationState `json:"location"`
	Timeouts       *mskconnectcustomplugin.TimeoutsState  `json:"timeouts"`
}
