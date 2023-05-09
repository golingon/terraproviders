// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	iamworkforcepool "github.com/golingon/terraproviders/googlebeta/4.63.1/iamworkforcepool"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIamWorkforcePool creates a new instance of [IamWorkforcePool].
func NewIamWorkforcePool(name string, args IamWorkforcePoolArgs) *IamWorkforcePool {
	return &IamWorkforcePool{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IamWorkforcePool)(nil)

// IamWorkforcePool represents the Terraform resource google_iam_workforce_pool.
type IamWorkforcePool struct {
	Name      string
	Args      IamWorkforcePoolArgs
	state     *iamWorkforcePoolState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IamWorkforcePool].
func (iwp *IamWorkforcePool) Type() string {
	return "google_iam_workforce_pool"
}

// LocalName returns the local name for [IamWorkforcePool].
func (iwp *IamWorkforcePool) LocalName() string {
	return iwp.Name
}

// Configuration returns the configuration (args) for [IamWorkforcePool].
func (iwp *IamWorkforcePool) Configuration() interface{} {
	return iwp.Args
}

// DependOn is used for other resources to depend on [IamWorkforcePool].
func (iwp *IamWorkforcePool) DependOn() terra.Reference {
	return terra.ReferenceResource(iwp)
}

// Dependencies returns the list of resources [IamWorkforcePool] depends_on.
func (iwp *IamWorkforcePool) Dependencies() terra.Dependencies {
	return iwp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IamWorkforcePool].
func (iwp *IamWorkforcePool) LifecycleManagement() *terra.Lifecycle {
	return iwp.Lifecycle
}

// Attributes returns the attributes for [IamWorkforcePool].
func (iwp *IamWorkforcePool) Attributes() iamWorkforcePoolAttributes {
	return iamWorkforcePoolAttributes{ref: terra.ReferenceResource(iwp)}
}

// ImportState imports the given attribute values into [IamWorkforcePool]'s state.
func (iwp *IamWorkforcePool) ImportState(av io.Reader) error {
	iwp.state = &iamWorkforcePoolState{}
	if err := json.NewDecoder(av).Decode(iwp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", iwp.Type(), iwp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IamWorkforcePool] has state.
func (iwp *IamWorkforcePool) State() (*iamWorkforcePoolState, bool) {
	return iwp.state, iwp.state != nil
}

// StateMust returns the state for [IamWorkforcePool]. Panics if the state is nil.
func (iwp *IamWorkforcePool) StateMust() *iamWorkforcePoolState {
	if iwp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", iwp.Type(), iwp.LocalName()))
	}
	return iwp.state
}

// IamWorkforcePoolArgs contains the configurations for google_iam_workforce_pool.
type IamWorkforcePoolArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Disabled: bool, optional
	Disabled terra.BoolValue `hcl:"disabled,attr"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Parent: string, required
	Parent terra.StringValue `hcl:"parent,attr" validate:"required"`
	// SessionDuration: string, optional
	SessionDuration terra.StringValue `hcl:"session_duration,attr"`
	// WorkforcePoolId: string, required
	WorkforcePoolId terra.StringValue `hcl:"workforce_pool_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *iamworkforcepool.Timeouts `hcl:"timeouts,block"`
}
type iamWorkforcePoolAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_iam_workforce_pool.
func (iwp iamWorkforcePoolAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(iwp.ref.Append("description"))
}

// Disabled returns a reference to field disabled of google_iam_workforce_pool.
func (iwp iamWorkforcePoolAttributes) Disabled() terra.BoolValue {
	return terra.ReferenceAsBool(iwp.ref.Append("disabled"))
}

// DisplayName returns a reference to field display_name of google_iam_workforce_pool.
func (iwp iamWorkforcePoolAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(iwp.ref.Append("display_name"))
}

// Id returns a reference to field id of google_iam_workforce_pool.
func (iwp iamWorkforcePoolAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(iwp.ref.Append("id"))
}

// Location returns a reference to field location of google_iam_workforce_pool.
func (iwp iamWorkforcePoolAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(iwp.ref.Append("location"))
}

// Name returns a reference to field name of google_iam_workforce_pool.
func (iwp iamWorkforcePoolAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(iwp.ref.Append("name"))
}

// Parent returns a reference to field parent of google_iam_workforce_pool.
func (iwp iamWorkforcePoolAttributes) Parent() terra.StringValue {
	return terra.ReferenceAsString(iwp.ref.Append("parent"))
}

// SessionDuration returns a reference to field session_duration of google_iam_workforce_pool.
func (iwp iamWorkforcePoolAttributes) SessionDuration() terra.StringValue {
	return terra.ReferenceAsString(iwp.ref.Append("session_duration"))
}

// State returns a reference to field state of google_iam_workforce_pool.
func (iwp iamWorkforcePoolAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(iwp.ref.Append("state"))
}

// WorkforcePoolId returns a reference to field workforce_pool_id of google_iam_workforce_pool.
func (iwp iamWorkforcePoolAttributes) WorkforcePoolId() terra.StringValue {
	return terra.ReferenceAsString(iwp.ref.Append("workforce_pool_id"))
}

func (iwp iamWorkforcePoolAttributes) Timeouts() iamworkforcepool.TimeoutsAttributes {
	return terra.ReferenceAsSingle[iamworkforcepool.TimeoutsAttributes](iwp.ref.Append("timeouts"))
}

type iamWorkforcePoolState struct {
	Description     string                          `json:"description"`
	Disabled        bool                            `json:"disabled"`
	DisplayName     string                          `json:"display_name"`
	Id              string                          `json:"id"`
	Location        string                          `json:"location"`
	Name            string                          `json:"name"`
	Parent          string                          `json:"parent"`
	SessionDuration string                          `json:"session_duration"`
	State           string                          `json:"state"`
	WorkforcePoolId string                          `json:"workforce_pool_id"`
	Timeouts        *iamworkforcepool.TimeoutsState `json:"timeouts"`
}
