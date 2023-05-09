// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	apigeeenvgroup "github.com/golingon/terraproviders/googlebeta/4.63.1/apigeeenvgroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApigeeEnvgroup creates a new instance of [ApigeeEnvgroup].
func NewApigeeEnvgroup(name string, args ApigeeEnvgroupArgs) *ApigeeEnvgroup {
	return &ApigeeEnvgroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApigeeEnvgroup)(nil)

// ApigeeEnvgroup represents the Terraform resource google_apigee_envgroup.
type ApigeeEnvgroup struct {
	Name      string
	Args      ApigeeEnvgroupArgs
	state     *apigeeEnvgroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApigeeEnvgroup].
func (ae *ApigeeEnvgroup) Type() string {
	return "google_apigee_envgroup"
}

// LocalName returns the local name for [ApigeeEnvgroup].
func (ae *ApigeeEnvgroup) LocalName() string {
	return ae.Name
}

// Configuration returns the configuration (args) for [ApigeeEnvgroup].
func (ae *ApigeeEnvgroup) Configuration() interface{} {
	return ae.Args
}

// DependOn is used for other resources to depend on [ApigeeEnvgroup].
func (ae *ApigeeEnvgroup) DependOn() terra.Reference {
	return terra.ReferenceResource(ae)
}

// Dependencies returns the list of resources [ApigeeEnvgroup] depends_on.
func (ae *ApigeeEnvgroup) Dependencies() terra.Dependencies {
	return ae.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApigeeEnvgroup].
func (ae *ApigeeEnvgroup) LifecycleManagement() *terra.Lifecycle {
	return ae.Lifecycle
}

// Attributes returns the attributes for [ApigeeEnvgroup].
func (ae *ApigeeEnvgroup) Attributes() apigeeEnvgroupAttributes {
	return apigeeEnvgroupAttributes{ref: terra.ReferenceResource(ae)}
}

// ImportState imports the given attribute values into [ApigeeEnvgroup]'s state.
func (ae *ApigeeEnvgroup) ImportState(av io.Reader) error {
	ae.state = &apigeeEnvgroupState{}
	if err := json.NewDecoder(av).Decode(ae.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ae.Type(), ae.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApigeeEnvgroup] has state.
func (ae *ApigeeEnvgroup) State() (*apigeeEnvgroupState, bool) {
	return ae.state, ae.state != nil
}

// StateMust returns the state for [ApigeeEnvgroup]. Panics if the state is nil.
func (ae *ApigeeEnvgroup) StateMust() *apigeeEnvgroupState {
	if ae.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ae.Type(), ae.LocalName()))
	}
	return ae.state
}

// ApigeeEnvgroupArgs contains the configurations for google_apigee_envgroup.
type ApigeeEnvgroupArgs struct {
	// Hostnames: list of string, optional
	Hostnames terra.ListValue[terra.StringValue] `hcl:"hostnames,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// OrgId: string, required
	OrgId terra.StringValue `hcl:"org_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *apigeeenvgroup.Timeouts `hcl:"timeouts,block"`
}
type apigeeEnvgroupAttributes struct {
	ref terra.Reference
}

// Hostnames returns a reference to field hostnames of google_apigee_envgroup.
func (ae apigeeEnvgroupAttributes) Hostnames() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ae.ref.Append("hostnames"))
}

// Id returns a reference to field id of google_apigee_envgroup.
func (ae apigeeEnvgroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ae.ref.Append("id"))
}

// Name returns a reference to field name of google_apigee_envgroup.
func (ae apigeeEnvgroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ae.ref.Append("name"))
}

// OrgId returns a reference to field org_id of google_apigee_envgroup.
func (ae apigeeEnvgroupAttributes) OrgId() terra.StringValue {
	return terra.ReferenceAsString(ae.ref.Append("org_id"))
}

func (ae apigeeEnvgroupAttributes) Timeouts() apigeeenvgroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apigeeenvgroup.TimeoutsAttributes](ae.ref.Append("timeouts"))
}

type apigeeEnvgroupState struct {
	Hostnames []string                      `json:"hostnames"`
	Id        string                        `json:"id"`
	Name      string                        `json:"name"`
	OrgId     string                        `json:"org_id"`
	Timeouts  *apigeeenvgroup.TimeoutsState `json:"timeouts"`
}
