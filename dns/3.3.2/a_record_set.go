// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dns

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewARecordSet creates a new instance of [ARecordSet].
func NewARecordSet(name string, args ARecordSetArgs) *ARecordSet {
	return &ARecordSet{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ARecordSet)(nil)

// ARecordSet represents the Terraform resource dns_a_record_set.
type ARecordSet struct {
	Name      string
	Args      ARecordSetArgs
	state     *aRecordSetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ARecordSet].
func (ars *ARecordSet) Type() string {
	return "dns_a_record_set"
}

// LocalName returns the local name for [ARecordSet].
func (ars *ARecordSet) LocalName() string {
	return ars.Name
}

// Configuration returns the configuration (args) for [ARecordSet].
func (ars *ARecordSet) Configuration() interface{} {
	return ars.Args
}

// DependOn is used for other resources to depend on [ARecordSet].
func (ars *ARecordSet) DependOn() terra.Reference {
	return terra.ReferenceResource(ars)
}

// Dependencies returns the list of resources [ARecordSet] depends_on.
func (ars *ARecordSet) Dependencies() terra.Dependencies {
	return ars.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ARecordSet].
func (ars *ARecordSet) LifecycleManagement() *terra.Lifecycle {
	return ars.Lifecycle
}

// Attributes returns the attributes for [ARecordSet].
func (ars *ARecordSet) Attributes() aRecordSetAttributes {
	return aRecordSetAttributes{ref: terra.ReferenceResource(ars)}
}

// ImportState imports the given attribute values into [ARecordSet]'s state.
func (ars *ARecordSet) ImportState(av io.Reader) error {
	ars.state = &aRecordSetState{}
	if err := json.NewDecoder(av).Decode(ars.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ars.Type(), ars.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ARecordSet] has state.
func (ars *ARecordSet) State() (*aRecordSetState, bool) {
	return ars.state, ars.state != nil
}

// StateMust returns the state for [ARecordSet]. Panics if the state is nil.
func (ars *ARecordSet) StateMust() *aRecordSetState {
	if ars.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ars.Type(), ars.LocalName()))
	}
	return ars.state
}

// ARecordSetArgs contains the configurations for dns_a_record_set.
type ARecordSetArgs struct {
	// Addresses: set of string, required
	Addresses terra.SetValue[terra.StringValue] `hcl:"addresses,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Ttl: number, optional
	Ttl terra.NumberValue `hcl:"ttl,attr"`
	// Zone: string, required
	Zone terra.StringValue `hcl:"zone,attr" validate:"required"`
}
type aRecordSetAttributes struct {
	ref terra.Reference
}

// Addresses returns a reference to field addresses of dns_a_record_set.
func (ars aRecordSetAttributes) Addresses() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ars.ref.Append("addresses"))
}

// Id returns a reference to field id of dns_a_record_set.
func (ars aRecordSetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ars.ref.Append("id"))
}

// Name returns a reference to field name of dns_a_record_set.
func (ars aRecordSetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ars.ref.Append("name"))
}

// Ttl returns a reference to field ttl of dns_a_record_set.
func (ars aRecordSetAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceAsNumber(ars.ref.Append("ttl"))
}

// Zone returns a reference to field zone of dns_a_record_set.
func (ars aRecordSetAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(ars.ref.Append("zone"))
}

type aRecordSetState struct {
	Addresses []string `json:"addresses"`
	Id        string   `json:"id"`
	Name      string   `json:"name"`
	Ttl       float64  `json:"ttl"`
	Zone      string   `json:"zone"`
}