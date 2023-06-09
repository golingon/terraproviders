// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	cloudidentitygroup "github.com/golingon/terraproviders/googlebeta/4.66.0/cloudidentitygroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudIdentityGroup creates a new instance of [CloudIdentityGroup].
func NewCloudIdentityGroup(name string, args CloudIdentityGroupArgs) *CloudIdentityGroup {
	return &CloudIdentityGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudIdentityGroup)(nil)

// CloudIdentityGroup represents the Terraform resource google_cloud_identity_group.
type CloudIdentityGroup struct {
	Name      string
	Args      CloudIdentityGroupArgs
	state     *cloudIdentityGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CloudIdentityGroup].
func (cig *CloudIdentityGroup) Type() string {
	return "google_cloud_identity_group"
}

// LocalName returns the local name for [CloudIdentityGroup].
func (cig *CloudIdentityGroup) LocalName() string {
	return cig.Name
}

// Configuration returns the configuration (args) for [CloudIdentityGroup].
func (cig *CloudIdentityGroup) Configuration() interface{} {
	return cig.Args
}

// DependOn is used for other resources to depend on [CloudIdentityGroup].
func (cig *CloudIdentityGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(cig)
}

// Dependencies returns the list of resources [CloudIdentityGroup] depends_on.
func (cig *CloudIdentityGroup) Dependencies() terra.Dependencies {
	return cig.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CloudIdentityGroup].
func (cig *CloudIdentityGroup) LifecycleManagement() *terra.Lifecycle {
	return cig.Lifecycle
}

// Attributes returns the attributes for [CloudIdentityGroup].
func (cig *CloudIdentityGroup) Attributes() cloudIdentityGroupAttributes {
	return cloudIdentityGroupAttributes{ref: terra.ReferenceResource(cig)}
}

// ImportState imports the given attribute values into [CloudIdentityGroup]'s state.
func (cig *CloudIdentityGroup) ImportState(av io.Reader) error {
	cig.state = &cloudIdentityGroupState{}
	if err := json.NewDecoder(av).Decode(cig.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cig.Type(), cig.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CloudIdentityGroup] has state.
func (cig *CloudIdentityGroup) State() (*cloudIdentityGroupState, bool) {
	return cig.state, cig.state != nil
}

// StateMust returns the state for [CloudIdentityGroup]. Panics if the state is nil.
func (cig *CloudIdentityGroup) StateMust() *cloudIdentityGroupState {
	if cig.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cig.Type(), cig.LocalName()))
	}
	return cig.state
}

// CloudIdentityGroupArgs contains the configurations for google_cloud_identity_group.
type CloudIdentityGroupArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InitialGroupConfig: string, optional
	InitialGroupConfig terra.StringValue `hcl:"initial_group_config,attr"`
	// Labels: map of string, required
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr" validate:"required"`
	// Parent: string, required
	Parent terra.StringValue `hcl:"parent,attr" validate:"required"`
	// GroupKey: required
	GroupKey *cloudidentitygroup.GroupKey `hcl:"group_key,block" validate:"required"`
	// Timeouts: optional
	Timeouts *cloudidentitygroup.Timeouts `hcl:"timeouts,block"`
}
type cloudIdentityGroupAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_cloud_identity_group.
func (cig cloudIdentityGroupAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(cig.ref.Append("create_time"))
}

// Description returns a reference to field description of google_cloud_identity_group.
func (cig cloudIdentityGroupAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cig.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of google_cloud_identity_group.
func (cig cloudIdentityGroupAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(cig.ref.Append("display_name"))
}

// Id returns a reference to field id of google_cloud_identity_group.
func (cig cloudIdentityGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cig.ref.Append("id"))
}

// InitialGroupConfig returns a reference to field initial_group_config of google_cloud_identity_group.
func (cig cloudIdentityGroupAttributes) InitialGroupConfig() terra.StringValue {
	return terra.ReferenceAsString(cig.ref.Append("initial_group_config"))
}

// Labels returns a reference to field labels of google_cloud_identity_group.
func (cig cloudIdentityGroupAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cig.ref.Append("labels"))
}

// Name returns a reference to field name of google_cloud_identity_group.
func (cig cloudIdentityGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cig.ref.Append("name"))
}

// Parent returns a reference to field parent of google_cloud_identity_group.
func (cig cloudIdentityGroupAttributes) Parent() terra.StringValue {
	return terra.ReferenceAsString(cig.ref.Append("parent"))
}

// UpdateTime returns a reference to field update_time of google_cloud_identity_group.
func (cig cloudIdentityGroupAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(cig.ref.Append("update_time"))
}

func (cig cloudIdentityGroupAttributes) GroupKey() terra.ListValue[cloudidentitygroup.GroupKeyAttributes] {
	return terra.ReferenceAsList[cloudidentitygroup.GroupKeyAttributes](cig.ref.Append("group_key"))
}

func (cig cloudIdentityGroupAttributes) Timeouts() cloudidentitygroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cloudidentitygroup.TimeoutsAttributes](cig.ref.Append("timeouts"))
}

type cloudIdentityGroupState struct {
	CreateTime         string                             `json:"create_time"`
	Description        string                             `json:"description"`
	DisplayName        string                             `json:"display_name"`
	Id                 string                             `json:"id"`
	InitialGroupConfig string                             `json:"initial_group_config"`
	Labels             map[string]string                  `json:"labels"`
	Name               string                             `json:"name"`
	Parent             string                             `json:"parent"`
	UpdateTime         string                             `json:"update_time"`
	GroupKey           []cloudidentitygroup.GroupKeyState `json:"group_key"`
	Timeouts           *cloudidentitygroup.TimeoutsState  `json:"timeouts"`
}
