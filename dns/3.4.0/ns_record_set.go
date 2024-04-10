// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package dns

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewNsRecordSet creates a new instance of [NsRecordSet].
func NewNsRecordSet(name string, args NsRecordSetArgs) *NsRecordSet {
	return &NsRecordSet{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NsRecordSet)(nil)

// NsRecordSet represents the Terraform resource dns_ns_record_set.
type NsRecordSet struct {
	Name      string
	Args      NsRecordSetArgs
	state     *nsRecordSetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NsRecordSet].
func (nrs *NsRecordSet) Type() string {
	return "dns_ns_record_set"
}

// LocalName returns the local name for [NsRecordSet].
func (nrs *NsRecordSet) LocalName() string {
	return nrs.Name
}

// Configuration returns the configuration (args) for [NsRecordSet].
func (nrs *NsRecordSet) Configuration() interface{} {
	return nrs.Args
}

// DependOn is used for other resources to depend on [NsRecordSet].
func (nrs *NsRecordSet) DependOn() terra.Reference {
	return terra.ReferenceResource(nrs)
}

// Dependencies returns the list of resources [NsRecordSet] depends_on.
func (nrs *NsRecordSet) Dependencies() terra.Dependencies {
	return nrs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NsRecordSet].
func (nrs *NsRecordSet) LifecycleManagement() *terra.Lifecycle {
	return nrs.Lifecycle
}

// Attributes returns the attributes for [NsRecordSet].
func (nrs *NsRecordSet) Attributes() nsRecordSetAttributes {
	return nsRecordSetAttributes{ref: terra.ReferenceResource(nrs)}
}

// ImportState imports the given attribute values into [NsRecordSet]'s state.
func (nrs *NsRecordSet) ImportState(av io.Reader) error {
	nrs.state = &nsRecordSetState{}
	if err := json.NewDecoder(av).Decode(nrs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nrs.Type(), nrs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NsRecordSet] has state.
func (nrs *NsRecordSet) State() (*nsRecordSetState, bool) {
	return nrs.state, nrs.state != nil
}

// StateMust returns the state for [NsRecordSet]. Panics if the state is nil.
func (nrs *NsRecordSet) StateMust() *nsRecordSetState {
	if nrs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nrs.Type(), nrs.LocalName()))
	}
	return nrs.state
}

// NsRecordSetArgs contains the configurations for dns_ns_record_set.
type NsRecordSetArgs struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Nameservers: set of string, required
	Nameservers terra.SetValue[terra.StringValue] `hcl:"nameservers,attr" validate:"required"`
	// Ttl: number, optional
	Ttl terra.NumberValue `hcl:"ttl,attr"`
	// Zone: string, required
	Zone terra.StringValue `hcl:"zone,attr" validate:"required"`
}
type nsRecordSetAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of dns_ns_record_set.
func (nrs nsRecordSetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nrs.ref.Append("id"))
}

// Name returns a reference to field name of dns_ns_record_set.
func (nrs nsRecordSetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(nrs.ref.Append("name"))
}

// Nameservers returns a reference to field nameservers of dns_ns_record_set.
func (nrs nsRecordSetAttributes) Nameservers() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](nrs.ref.Append("nameservers"))
}

// Ttl returns a reference to field ttl of dns_ns_record_set.
func (nrs nsRecordSetAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceAsNumber(nrs.ref.Append("ttl"))
}

// Zone returns a reference to field zone of dns_ns_record_set.
func (nrs nsRecordSetAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(nrs.ref.Append("zone"))
}

type nsRecordSetState struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Nameservers []string `json:"nameservers"`
	Ttl         float64  `json:"ttl"`
	Zone        string   `json:"zone"`
}
