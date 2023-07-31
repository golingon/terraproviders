// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	bigtableinstance "github.com/golingon/terraproviders/google/4.75.1/bigtableinstance"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBigtableInstance creates a new instance of [BigtableInstance].
func NewBigtableInstance(name string, args BigtableInstanceArgs) *BigtableInstance {
	return &BigtableInstance{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BigtableInstance)(nil)

// BigtableInstance represents the Terraform resource google_bigtable_instance.
type BigtableInstance struct {
	Name      string
	Args      BigtableInstanceArgs
	state     *bigtableInstanceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BigtableInstance].
func (bi *BigtableInstance) Type() string {
	return "google_bigtable_instance"
}

// LocalName returns the local name for [BigtableInstance].
func (bi *BigtableInstance) LocalName() string {
	return bi.Name
}

// Configuration returns the configuration (args) for [BigtableInstance].
func (bi *BigtableInstance) Configuration() interface{} {
	return bi.Args
}

// DependOn is used for other resources to depend on [BigtableInstance].
func (bi *BigtableInstance) DependOn() terra.Reference {
	return terra.ReferenceResource(bi)
}

// Dependencies returns the list of resources [BigtableInstance] depends_on.
func (bi *BigtableInstance) Dependencies() terra.Dependencies {
	return bi.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BigtableInstance].
func (bi *BigtableInstance) LifecycleManagement() *terra.Lifecycle {
	return bi.Lifecycle
}

// Attributes returns the attributes for [BigtableInstance].
func (bi *BigtableInstance) Attributes() bigtableInstanceAttributes {
	return bigtableInstanceAttributes{ref: terra.ReferenceResource(bi)}
}

// ImportState imports the given attribute values into [BigtableInstance]'s state.
func (bi *BigtableInstance) ImportState(av io.Reader) error {
	bi.state = &bigtableInstanceState{}
	if err := json.NewDecoder(av).Decode(bi.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bi.Type(), bi.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BigtableInstance] has state.
func (bi *BigtableInstance) State() (*bigtableInstanceState, bool) {
	return bi.state, bi.state != nil
}

// StateMust returns the state for [BigtableInstance]. Panics if the state is nil.
func (bi *BigtableInstance) StateMust() *bigtableInstanceState {
	if bi.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bi.Type(), bi.LocalName()))
	}
	return bi.state
}

// BigtableInstanceArgs contains the configurations for google_bigtable_instance.
type BigtableInstanceArgs struct {
	// DeletionProtection: bool, optional
	DeletionProtection terra.BoolValue `hcl:"deletion_protection,attr"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceType: string, optional
	InstanceType terra.StringValue `hcl:"instance_type,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Cluster: min=0
	Cluster []bigtableinstance.Cluster `hcl:"cluster,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *bigtableinstance.Timeouts `hcl:"timeouts,block"`
}
type bigtableInstanceAttributes struct {
	ref terra.Reference
}

// DeletionProtection returns a reference to field deletion_protection of google_bigtable_instance.
func (bi bigtableInstanceAttributes) DeletionProtection() terra.BoolValue {
	return terra.ReferenceAsBool(bi.ref.Append("deletion_protection"))
}

// DisplayName returns a reference to field display_name of google_bigtable_instance.
func (bi bigtableInstanceAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(bi.ref.Append("display_name"))
}

// Id returns a reference to field id of google_bigtable_instance.
func (bi bigtableInstanceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bi.ref.Append("id"))
}

// InstanceType returns a reference to field instance_type of google_bigtable_instance.
func (bi bigtableInstanceAttributes) InstanceType() terra.StringValue {
	return terra.ReferenceAsString(bi.ref.Append("instance_type"))
}

// Labels returns a reference to field labels of google_bigtable_instance.
func (bi bigtableInstanceAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](bi.ref.Append("labels"))
}

// Name returns a reference to field name of google_bigtable_instance.
func (bi bigtableInstanceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(bi.ref.Append("name"))
}

// Project returns a reference to field project of google_bigtable_instance.
func (bi bigtableInstanceAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(bi.ref.Append("project"))
}

func (bi bigtableInstanceAttributes) Cluster() terra.ListValue[bigtableinstance.ClusterAttributes] {
	return terra.ReferenceAsList[bigtableinstance.ClusterAttributes](bi.ref.Append("cluster"))
}

func (bi bigtableInstanceAttributes) Timeouts() bigtableinstance.TimeoutsAttributes {
	return terra.ReferenceAsSingle[bigtableinstance.TimeoutsAttributes](bi.ref.Append("timeouts"))
}

type bigtableInstanceState struct {
	DeletionProtection bool                            `json:"deletion_protection"`
	DisplayName        string                          `json:"display_name"`
	Id                 string                          `json:"id"`
	InstanceType       string                          `json:"instance_type"`
	Labels             map[string]string               `json:"labels"`
	Name               string                          `json:"name"`
	Project            string                          `json:"project"`
	Cluster            []bigtableinstance.ClusterState `json:"cluster"`
	Timeouts           *bigtableinstance.TimeoutsState `json:"timeouts"`
}
