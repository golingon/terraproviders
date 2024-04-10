// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package dns

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewCnameRecord creates a new instance of [CnameRecord].
func NewCnameRecord(name string, args CnameRecordArgs) *CnameRecord {
	return &CnameRecord{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CnameRecord)(nil)

// CnameRecord represents the Terraform resource dns_cname_record.
type CnameRecord struct {
	Name      string
	Args      CnameRecordArgs
	state     *cnameRecordState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CnameRecord].
func (cr *CnameRecord) Type() string {
	return "dns_cname_record"
}

// LocalName returns the local name for [CnameRecord].
func (cr *CnameRecord) LocalName() string {
	return cr.Name
}

// Configuration returns the configuration (args) for [CnameRecord].
func (cr *CnameRecord) Configuration() interface{} {
	return cr.Args
}

// DependOn is used for other resources to depend on [CnameRecord].
func (cr *CnameRecord) DependOn() terra.Reference {
	return terra.ReferenceResource(cr)
}

// Dependencies returns the list of resources [CnameRecord] depends_on.
func (cr *CnameRecord) Dependencies() terra.Dependencies {
	return cr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CnameRecord].
func (cr *CnameRecord) LifecycleManagement() *terra.Lifecycle {
	return cr.Lifecycle
}

// Attributes returns the attributes for [CnameRecord].
func (cr *CnameRecord) Attributes() cnameRecordAttributes {
	return cnameRecordAttributes{ref: terra.ReferenceResource(cr)}
}

// ImportState imports the given attribute values into [CnameRecord]'s state.
func (cr *CnameRecord) ImportState(av io.Reader) error {
	cr.state = &cnameRecordState{}
	if err := json.NewDecoder(av).Decode(cr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cr.Type(), cr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CnameRecord] has state.
func (cr *CnameRecord) State() (*cnameRecordState, bool) {
	return cr.state, cr.state != nil
}

// StateMust returns the state for [CnameRecord]. Panics if the state is nil.
func (cr *CnameRecord) StateMust() *cnameRecordState {
	if cr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cr.Type(), cr.LocalName()))
	}
	return cr.state
}

// CnameRecordArgs contains the configurations for dns_cname_record.
type CnameRecordArgs struct {
	// Cname: string, required
	Cname terra.StringValue `hcl:"cname,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Ttl: number, optional
	Ttl terra.NumberValue `hcl:"ttl,attr"`
	// Zone: string, required
	Zone terra.StringValue `hcl:"zone,attr" validate:"required"`
}
type cnameRecordAttributes struct {
	ref terra.Reference
}

// Cname returns a reference to field cname of dns_cname_record.
func (cr cnameRecordAttributes) Cname() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("cname"))
}

// Id returns a reference to field id of dns_cname_record.
func (cr cnameRecordAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("id"))
}

// Name returns a reference to field name of dns_cname_record.
func (cr cnameRecordAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("name"))
}

// Ttl returns a reference to field ttl of dns_cname_record.
func (cr cnameRecordAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceAsNumber(cr.ref.Append("ttl"))
}

// Zone returns a reference to field zone of dns_cname_record.
func (cr cnameRecordAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("zone"))
}

type cnameRecordState struct {
	Cname string  `json:"cname"`
	Id    string  `json:"id"`
	Name  string  `json:"name"`
	Ttl   float64 `json:"ttl"`
	Zone  string  `json:"zone"`
}
