// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dns

import (
	"encoding/json"
	"fmt"
	mxrecordset "github.com/golingon/terraproviders/dns/3.3.0/mxrecordset"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMxRecordSet creates a new instance of [MxRecordSet].
func NewMxRecordSet(name string, args MxRecordSetArgs) *MxRecordSet {
	return &MxRecordSet{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MxRecordSet)(nil)

// MxRecordSet represents the Terraform resource dns_mx_record_set.
type MxRecordSet struct {
	Name      string
	Args      MxRecordSetArgs
	state     *mxRecordSetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MxRecordSet].
func (mrs *MxRecordSet) Type() string {
	return "dns_mx_record_set"
}

// LocalName returns the local name for [MxRecordSet].
func (mrs *MxRecordSet) LocalName() string {
	return mrs.Name
}

// Configuration returns the configuration (args) for [MxRecordSet].
func (mrs *MxRecordSet) Configuration() interface{} {
	return mrs.Args
}

// DependOn is used for other resources to depend on [MxRecordSet].
func (mrs *MxRecordSet) DependOn() terra.Reference {
	return terra.ReferenceResource(mrs)
}

// Dependencies returns the list of resources [MxRecordSet] depends_on.
func (mrs *MxRecordSet) Dependencies() terra.Dependencies {
	return mrs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MxRecordSet].
func (mrs *MxRecordSet) LifecycleManagement() *terra.Lifecycle {
	return mrs.Lifecycle
}

// Attributes returns the attributes for [MxRecordSet].
func (mrs *MxRecordSet) Attributes() mxRecordSetAttributes {
	return mxRecordSetAttributes{ref: terra.ReferenceResource(mrs)}
}

// ImportState imports the given attribute values into [MxRecordSet]'s state.
func (mrs *MxRecordSet) ImportState(av io.Reader) error {
	mrs.state = &mxRecordSetState{}
	if err := json.NewDecoder(av).Decode(mrs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mrs.Type(), mrs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MxRecordSet] has state.
func (mrs *MxRecordSet) State() (*mxRecordSetState, bool) {
	return mrs.state, mrs.state != nil
}

// StateMust returns the state for [MxRecordSet]. Panics if the state is nil.
func (mrs *MxRecordSet) StateMust() *mxRecordSetState {
	if mrs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mrs.Type(), mrs.LocalName()))
	}
	return mrs.state
}

// MxRecordSetArgs contains the configurations for dns_mx_record_set.
type MxRecordSetArgs struct {
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Ttl: number, optional
	Ttl terra.NumberValue `hcl:"ttl,attr"`
	// Zone: string, required
	Zone terra.StringValue `hcl:"zone,attr" validate:"required"`
	// Mx: min=0
	Mx []mxrecordset.Mx `hcl:"mx,block" validate:"min=0"`
}
type mxRecordSetAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of dns_mx_record_set.
func (mrs mxRecordSetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mrs.ref.Append("id"))
}

// Name returns a reference to field name of dns_mx_record_set.
func (mrs mxRecordSetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mrs.ref.Append("name"))
}

// Ttl returns a reference to field ttl of dns_mx_record_set.
func (mrs mxRecordSetAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceAsNumber(mrs.ref.Append("ttl"))
}

// Zone returns a reference to field zone of dns_mx_record_set.
func (mrs mxRecordSetAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(mrs.ref.Append("zone"))
}

func (mrs mxRecordSetAttributes) Mx() terra.SetValue[mxrecordset.MxAttributes] {
	return terra.ReferenceAsSet[mxrecordset.MxAttributes](mrs.ref.Append("mx"))
}

type mxRecordSetState struct {
	Id   string                `json:"id"`
	Name string                `json:"name"`
	Ttl  float64               `json:"ttl"`
	Zone string                `json:"zone"`
	Mx   []mxrecordset.MxState `json:"mx"`
}
