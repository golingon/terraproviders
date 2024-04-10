// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	servicecatalogtagoption "github.com/golingon/terraproviders/aws/5.44.0/servicecatalogtagoption"
	"io"
)

// NewServicecatalogTagOption creates a new instance of [ServicecatalogTagOption].
func NewServicecatalogTagOption(name string, args ServicecatalogTagOptionArgs) *ServicecatalogTagOption {
	return &ServicecatalogTagOption{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ServicecatalogTagOption)(nil)

// ServicecatalogTagOption represents the Terraform resource aws_servicecatalog_tag_option.
type ServicecatalogTagOption struct {
	Name      string
	Args      ServicecatalogTagOptionArgs
	state     *servicecatalogTagOptionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ServicecatalogTagOption].
func (sto *ServicecatalogTagOption) Type() string {
	return "aws_servicecatalog_tag_option"
}

// LocalName returns the local name for [ServicecatalogTagOption].
func (sto *ServicecatalogTagOption) LocalName() string {
	return sto.Name
}

// Configuration returns the configuration (args) for [ServicecatalogTagOption].
func (sto *ServicecatalogTagOption) Configuration() interface{} {
	return sto.Args
}

// DependOn is used for other resources to depend on [ServicecatalogTagOption].
func (sto *ServicecatalogTagOption) DependOn() terra.Reference {
	return terra.ReferenceResource(sto)
}

// Dependencies returns the list of resources [ServicecatalogTagOption] depends_on.
func (sto *ServicecatalogTagOption) Dependencies() terra.Dependencies {
	return sto.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ServicecatalogTagOption].
func (sto *ServicecatalogTagOption) LifecycleManagement() *terra.Lifecycle {
	return sto.Lifecycle
}

// Attributes returns the attributes for [ServicecatalogTagOption].
func (sto *ServicecatalogTagOption) Attributes() servicecatalogTagOptionAttributes {
	return servicecatalogTagOptionAttributes{ref: terra.ReferenceResource(sto)}
}

// ImportState imports the given attribute values into [ServicecatalogTagOption]'s state.
func (sto *ServicecatalogTagOption) ImportState(av io.Reader) error {
	sto.state = &servicecatalogTagOptionState{}
	if err := json.NewDecoder(av).Decode(sto.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sto.Type(), sto.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ServicecatalogTagOption] has state.
func (sto *ServicecatalogTagOption) State() (*servicecatalogTagOptionState, bool) {
	return sto.state, sto.state != nil
}

// StateMust returns the state for [ServicecatalogTagOption]. Panics if the state is nil.
func (sto *ServicecatalogTagOption) StateMust() *servicecatalogTagOptionState {
	if sto.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sto.Type(), sto.LocalName()))
	}
	return sto.state
}

// ServicecatalogTagOptionArgs contains the configurations for aws_servicecatalog_tag_option.
type ServicecatalogTagOptionArgs struct {
	// Active: bool, optional
	Active terra.BoolValue `hcl:"active,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Key: string, required
	Key terra.StringValue `hcl:"key,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *servicecatalogtagoption.Timeouts `hcl:"timeouts,block"`
}
type servicecatalogTagOptionAttributes struct {
	ref terra.Reference
}

// Active returns a reference to field active of aws_servicecatalog_tag_option.
func (sto servicecatalogTagOptionAttributes) Active() terra.BoolValue {
	return terra.ReferenceAsBool(sto.ref.Append("active"))
}

// Id returns a reference to field id of aws_servicecatalog_tag_option.
func (sto servicecatalogTagOptionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sto.ref.Append("id"))
}

// Key returns a reference to field key of aws_servicecatalog_tag_option.
func (sto servicecatalogTagOptionAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(sto.ref.Append("key"))
}

// Owner returns a reference to field owner of aws_servicecatalog_tag_option.
func (sto servicecatalogTagOptionAttributes) Owner() terra.StringValue {
	return terra.ReferenceAsString(sto.ref.Append("owner"))
}

// Value returns a reference to field value of aws_servicecatalog_tag_option.
func (sto servicecatalogTagOptionAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(sto.ref.Append("value"))
}

func (sto servicecatalogTagOptionAttributes) Timeouts() servicecatalogtagoption.TimeoutsAttributes {
	return terra.ReferenceAsSingle[servicecatalogtagoption.TimeoutsAttributes](sto.ref.Append("timeouts"))
}

type servicecatalogTagOptionState struct {
	Active   bool                                   `json:"active"`
	Id       string                                 `json:"id"`
	Key      string                                 `json:"key"`
	Owner    string                                 `json:"owner"`
	Value    string                                 `json:"value"`
	Timeouts *servicecatalogtagoption.TimeoutsState `json:"timeouts"`
}
