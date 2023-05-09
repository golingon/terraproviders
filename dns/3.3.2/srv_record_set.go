// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dns

import (
	"encoding/json"
	"fmt"
	srvrecordset "github.com/golingon/terraproviders/dns/3.3.2/srvrecordset"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSrvRecordSet creates a new instance of [SrvRecordSet].
func NewSrvRecordSet(name string, args SrvRecordSetArgs) *SrvRecordSet {
	return &SrvRecordSet{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SrvRecordSet)(nil)

// SrvRecordSet represents the Terraform resource dns_srv_record_set.
type SrvRecordSet struct {
	Name      string
	Args      SrvRecordSetArgs
	state     *srvRecordSetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SrvRecordSet].
func (srs *SrvRecordSet) Type() string {
	return "dns_srv_record_set"
}

// LocalName returns the local name for [SrvRecordSet].
func (srs *SrvRecordSet) LocalName() string {
	return srs.Name
}

// Configuration returns the configuration (args) for [SrvRecordSet].
func (srs *SrvRecordSet) Configuration() interface{} {
	return srs.Args
}

// DependOn is used for other resources to depend on [SrvRecordSet].
func (srs *SrvRecordSet) DependOn() terra.Reference {
	return terra.ReferenceResource(srs)
}

// Dependencies returns the list of resources [SrvRecordSet] depends_on.
func (srs *SrvRecordSet) Dependencies() terra.Dependencies {
	return srs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SrvRecordSet].
func (srs *SrvRecordSet) LifecycleManagement() *terra.Lifecycle {
	return srs.Lifecycle
}

// Attributes returns the attributes for [SrvRecordSet].
func (srs *SrvRecordSet) Attributes() srvRecordSetAttributes {
	return srvRecordSetAttributes{ref: terra.ReferenceResource(srs)}
}

// ImportState imports the given attribute values into [SrvRecordSet]'s state.
func (srs *SrvRecordSet) ImportState(av io.Reader) error {
	srs.state = &srvRecordSetState{}
	if err := json.NewDecoder(av).Decode(srs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", srs.Type(), srs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SrvRecordSet] has state.
func (srs *SrvRecordSet) State() (*srvRecordSetState, bool) {
	return srs.state, srs.state != nil
}

// StateMust returns the state for [SrvRecordSet]. Panics if the state is nil.
func (srs *SrvRecordSet) StateMust() *srvRecordSetState {
	if srs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", srs.Type(), srs.LocalName()))
	}
	return srs.state
}

// SrvRecordSetArgs contains the configurations for dns_srv_record_set.
type SrvRecordSetArgs struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Ttl: number, optional
	Ttl terra.NumberValue `hcl:"ttl,attr"`
	// Zone: string, required
	Zone terra.StringValue `hcl:"zone,attr" validate:"required"`
	// Srv: min=0
	Srv []srvrecordset.Srv `hcl:"srv,block" validate:"min=0"`
}
type srvRecordSetAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of dns_srv_record_set.
func (srs srvRecordSetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(srs.ref.Append("id"))
}

// Name returns a reference to field name of dns_srv_record_set.
func (srs srvRecordSetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(srs.ref.Append("name"))
}

// Ttl returns a reference to field ttl of dns_srv_record_set.
func (srs srvRecordSetAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceAsNumber(srs.ref.Append("ttl"))
}

// Zone returns a reference to field zone of dns_srv_record_set.
func (srs srvRecordSetAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(srs.ref.Append("zone"))
}

func (srs srvRecordSetAttributes) Srv() terra.SetValue[srvrecordset.SrvAttributes] {
	return terra.ReferenceAsSet[srvrecordset.SrvAttributes](srs.ref.Append("srv"))
}

type srvRecordSetState struct {
	Id   string                  `json:"id"`
	Name string                  `json:"name"`
	Ttl  float64                 `json:"ttl"`
	Zone string                  `json:"zone"`
	Srv  []srvrecordset.SrvState `json:"srv"`
}
