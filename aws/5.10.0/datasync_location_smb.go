// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	datasynclocationsmb "github.com/golingon/terraproviders/aws/5.10.0/datasynclocationsmb"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDatasyncLocationSmb creates a new instance of [DatasyncLocationSmb].
func NewDatasyncLocationSmb(name string, args DatasyncLocationSmbArgs) *DatasyncLocationSmb {
	return &DatasyncLocationSmb{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DatasyncLocationSmb)(nil)

// DatasyncLocationSmb represents the Terraform resource aws_datasync_location_smb.
type DatasyncLocationSmb struct {
	Name      string
	Args      DatasyncLocationSmbArgs
	state     *datasyncLocationSmbState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DatasyncLocationSmb].
func (dls *DatasyncLocationSmb) Type() string {
	return "aws_datasync_location_smb"
}

// LocalName returns the local name for [DatasyncLocationSmb].
func (dls *DatasyncLocationSmb) LocalName() string {
	return dls.Name
}

// Configuration returns the configuration (args) for [DatasyncLocationSmb].
func (dls *DatasyncLocationSmb) Configuration() interface{} {
	return dls.Args
}

// DependOn is used for other resources to depend on [DatasyncLocationSmb].
func (dls *DatasyncLocationSmb) DependOn() terra.Reference {
	return terra.ReferenceResource(dls)
}

// Dependencies returns the list of resources [DatasyncLocationSmb] depends_on.
func (dls *DatasyncLocationSmb) Dependencies() terra.Dependencies {
	return dls.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DatasyncLocationSmb].
func (dls *DatasyncLocationSmb) LifecycleManagement() *terra.Lifecycle {
	return dls.Lifecycle
}

// Attributes returns the attributes for [DatasyncLocationSmb].
func (dls *DatasyncLocationSmb) Attributes() datasyncLocationSmbAttributes {
	return datasyncLocationSmbAttributes{ref: terra.ReferenceResource(dls)}
}

// ImportState imports the given attribute values into [DatasyncLocationSmb]'s state.
func (dls *DatasyncLocationSmb) ImportState(av io.Reader) error {
	dls.state = &datasyncLocationSmbState{}
	if err := json.NewDecoder(av).Decode(dls.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dls.Type(), dls.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DatasyncLocationSmb] has state.
func (dls *DatasyncLocationSmb) State() (*datasyncLocationSmbState, bool) {
	return dls.state, dls.state != nil
}

// StateMust returns the state for [DatasyncLocationSmb]. Panics if the state is nil.
func (dls *DatasyncLocationSmb) StateMust() *datasyncLocationSmbState {
	if dls.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dls.Type(), dls.LocalName()))
	}
	return dls.state
}

// DatasyncLocationSmbArgs contains the configurations for aws_datasync_location_smb.
type DatasyncLocationSmbArgs struct {
	// AgentArns: set of string, required
	AgentArns terra.SetValue[terra.StringValue] `hcl:"agent_arns,attr" validate:"required"`
	// Domain: string, optional
	Domain terra.StringValue `hcl:"domain,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Password: string, required
	Password terra.StringValue `hcl:"password,attr" validate:"required"`
	// ServerHostname: string, required
	ServerHostname terra.StringValue `hcl:"server_hostname,attr" validate:"required"`
	// Subdirectory: string, required
	Subdirectory terra.StringValue `hcl:"subdirectory,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// User: string, required
	User terra.StringValue `hcl:"user,attr" validate:"required"`
	// MountOptions: optional
	MountOptions *datasynclocationsmb.MountOptions `hcl:"mount_options,block"`
}
type datasyncLocationSmbAttributes struct {
	ref terra.Reference
}

// AgentArns returns a reference to field agent_arns of aws_datasync_location_smb.
func (dls datasyncLocationSmbAttributes) AgentArns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](dls.ref.Append("agent_arns"))
}

// Arn returns a reference to field arn of aws_datasync_location_smb.
func (dls datasyncLocationSmbAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(dls.ref.Append("arn"))
}

// Domain returns a reference to field domain of aws_datasync_location_smb.
func (dls datasyncLocationSmbAttributes) Domain() terra.StringValue {
	return terra.ReferenceAsString(dls.ref.Append("domain"))
}

// Id returns a reference to field id of aws_datasync_location_smb.
func (dls datasyncLocationSmbAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dls.ref.Append("id"))
}

// Password returns a reference to field password of aws_datasync_location_smb.
func (dls datasyncLocationSmbAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(dls.ref.Append("password"))
}

// ServerHostname returns a reference to field server_hostname of aws_datasync_location_smb.
func (dls datasyncLocationSmbAttributes) ServerHostname() terra.StringValue {
	return terra.ReferenceAsString(dls.ref.Append("server_hostname"))
}

// Subdirectory returns a reference to field subdirectory of aws_datasync_location_smb.
func (dls datasyncLocationSmbAttributes) Subdirectory() terra.StringValue {
	return terra.ReferenceAsString(dls.ref.Append("subdirectory"))
}

// Tags returns a reference to field tags of aws_datasync_location_smb.
func (dls datasyncLocationSmbAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dls.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_datasync_location_smb.
func (dls datasyncLocationSmbAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dls.ref.Append("tags_all"))
}

// Uri returns a reference to field uri of aws_datasync_location_smb.
func (dls datasyncLocationSmbAttributes) Uri() terra.StringValue {
	return terra.ReferenceAsString(dls.ref.Append("uri"))
}

// User returns a reference to field user of aws_datasync_location_smb.
func (dls datasyncLocationSmbAttributes) User() terra.StringValue {
	return terra.ReferenceAsString(dls.ref.Append("user"))
}

func (dls datasyncLocationSmbAttributes) MountOptions() terra.ListValue[datasynclocationsmb.MountOptionsAttributes] {
	return terra.ReferenceAsList[datasynclocationsmb.MountOptionsAttributes](dls.ref.Append("mount_options"))
}

type datasyncLocationSmbState struct {
	AgentArns      []string                                `json:"agent_arns"`
	Arn            string                                  `json:"arn"`
	Domain         string                                  `json:"domain"`
	Id             string                                  `json:"id"`
	Password       string                                  `json:"password"`
	ServerHostname string                                  `json:"server_hostname"`
	Subdirectory   string                                  `json:"subdirectory"`
	Tags           map[string]string                       `json:"tags"`
	TagsAll        map[string]string                       `json:"tags_all"`
	Uri            string                                  `json:"uri"`
	User           string                                  `json:"user"`
	MountOptions   []datasynclocationsmb.MountOptionsState `json:"mount_options"`
}