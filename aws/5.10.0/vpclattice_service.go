// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	vpclatticeservice "github.com/golingon/terraproviders/aws/5.10.0/vpclatticeservice"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVpclatticeService creates a new instance of [VpclatticeService].
func NewVpclatticeService(name string, args VpclatticeServiceArgs) *VpclatticeService {
	return &VpclatticeService{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VpclatticeService)(nil)

// VpclatticeService represents the Terraform resource aws_vpclattice_service.
type VpclatticeService struct {
	Name      string
	Args      VpclatticeServiceArgs
	state     *vpclatticeServiceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VpclatticeService].
func (vs *VpclatticeService) Type() string {
	return "aws_vpclattice_service"
}

// LocalName returns the local name for [VpclatticeService].
func (vs *VpclatticeService) LocalName() string {
	return vs.Name
}

// Configuration returns the configuration (args) for [VpclatticeService].
func (vs *VpclatticeService) Configuration() interface{} {
	return vs.Args
}

// DependOn is used for other resources to depend on [VpclatticeService].
func (vs *VpclatticeService) DependOn() terra.Reference {
	return terra.ReferenceResource(vs)
}

// Dependencies returns the list of resources [VpclatticeService] depends_on.
func (vs *VpclatticeService) Dependencies() terra.Dependencies {
	return vs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VpclatticeService].
func (vs *VpclatticeService) LifecycleManagement() *terra.Lifecycle {
	return vs.Lifecycle
}

// Attributes returns the attributes for [VpclatticeService].
func (vs *VpclatticeService) Attributes() vpclatticeServiceAttributes {
	return vpclatticeServiceAttributes{ref: terra.ReferenceResource(vs)}
}

// ImportState imports the given attribute values into [VpclatticeService]'s state.
func (vs *VpclatticeService) ImportState(av io.Reader) error {
	vs.state = &vpclatticeServiceState{}
	if err := json.NewDecoder(av).Decode(vs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vs.Type(), vs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VpclatticeService] has state.
func (vs *VpclatticeService) State() (*vpclatticeServiceState, bool) {
	return vs.state, vs.state != nil
}

// StateMust returns the state for [VpclatticeService]. Panics if the state is nil.
func (vs *VpclatticeService) StateMust() *vpclatticeServiceState {
	if vs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vs.Type(), vs.LocalName()))
	}
	return vs.state
}

// VpclatticeServiceArgs contains the configurations for aws_vpclattice_service.
type VpclatticeServiceArgs struct {
	// AuthType: string, optional
	AuthType terra.StringValue `hcl:"auth_type,attr"`
	// CertificateArn: string, optional
	CertificateArn terra.StringValue `hcl:"certificate_arn,attr"`
	// CustomDomainName: string, optional
	CustomDomainName terra.StringValue `hcl:"custom_domain_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// DnsEntry: min=0
	DnsEntry []vpclatticeservice.DnsEntry `hcl:"dns_entry,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *vpclatticeservice.Timeouts `hcl:"timeouts,block"`
}
type vpclatticeServiceAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_vpclattice_service.
func (vs vpclatticeServiceAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(vs.ref.Append("arn"))
}

// AuthType returns a reference to field auth_type of aws_vpclattice_service.
func (vs vpclatticeServiceAttributes) AuthType() terra.StringValue {
	return terra.ReferenceAsString(vs.ref.Append("auth_type"))
}

// CertificateArn returns a reference to field certificate_arn of aws_vpclattice_service.
func (vs vpclatticeServiceAttributes) CertificateArn() terra.StringValue {
	return terra.ReferenceAsString(vs.ref.Append("certificate_arn"))
}

// CustomDomainName returns a reference to field custom_domain_name of aws_vpclattice_service.
func (vs vpclatticeServiceAttributes) CustomDomainName() terra.StringValue {
	return terra.ReferenceAsString(vs.ref.Append("custom_domain_name"))
}

// Id returns a reference to field id of aws_vpclattice_service.
func (vs vpclatticeServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vs.ref.Append("id"))
}

// Name returns a reference to field name of aws_vpclattice_service.
func (vs vpclatticeServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vs.ref.Append("name"))
}

// Status returns a reference to field status of aws_vpclattice_service.
func (vs vpclatticeServiceAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(vs.ref.Append("status"))
}

// Tags returns a reference to field tags of aws_vpclattice_service.
func (vs vpclatticeServiceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vs.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_vpclattice_service.
func (vs vpclatticeServiceAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vs.ref.Append("tags_all"))
}

func (vs vpclatticeServiceAttributes) DnsEntry() terra.ListValue[vpclatticeservice.DnsEntryAttributes] {
	return terra.ReferenceAsList[vpclatticeservice.DnsEntryAttributes](vs.ref.Append("dns_entry"))
}

func (vs vpclatticeServiceAttributes) Timeouts() vpclatticeservice.TimeoutsAttributes {
	return terra.ReferenceAsSingle[vpclatticeservice.TimeoutsAttributes](vs.ref.Append("timeouts"))
}

type vpclatticeServiceState struct {
	Arn              string                            `json:"arn"`
	AuthType         string                            `json:"auth_type"`
	CertificateArn   string                            `json:"certificate_arn"`
	CustomDomainName string                            `json:"custom_domain_name"`
	Id               string                            `json:"id"`
	Name             string                            `json:"name"`
	Status           string                            `json:"status"`
	Tags             map[string]string                 `json:"tags"`
	TagsAll          map[string]string                 `json:"tags_all"`
	DnsEntry         []vpclatticeservice.DnsEntryState `json:"dns_entry"`
	Timeouts         *vpclatticeservice.TimeoutsState  `json:"timeouts"`
}
