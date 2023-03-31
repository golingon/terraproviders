// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	loadtest "github.com/golingon/terraproviders/azurerm/3.49.0/loadtest"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewLoadTest(name string, args LoadTestArgs) *LoadTest {
	return &LoadTest{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LoadTest)(nil)

type LoadTest struct {
	Name  string
	Args  LoadTestArgs
	state *loadTestState
}

func (lt *LoadTest) Type() string {
	return "azurerm_load_test"
}

func (lt *LoadTest) LocalName() string {
	return lt.Name
}

func (lt *LoadTest) Configuration() interface{} {
	return lt.Args
}

func (lt *LoadTest) Attributes() loadTestAttributes {
	return loadTestAttributes{ref: terra.ReferenceResource(lt)}
}

func (lt *LoadTest) ImportState(av io.Reader) error {
	lt.state = &loadTestState{}
	if err := json.NewDecoder(av).Decode(lt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lt.Type(), lt.LocalName(), err)
	}
	return nil
}

func (lt *LoadTest) State() (*loadTestState, bool) {
	return lt.state, lt.state != nil
}

func (lt *LoadTest) StateMust() *loadTestState {
	if lt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lt.Type(), lt.LocalName()))
	}
	return lt.state
}

func (lt *LoadTest) DependOn() terra.Reference {
	return terra.ReferenceResource(lt)
}

type LoadTestArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Identity: optional
	Identity *loadtest.Identity `hcl:"identity,block"`
	// Timeouts: optional
	Timeouts *loadtest.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that LoadTest depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type loadTestAttributes struct {
	ref terra.Reference
}

func (lt loadTestAttributes) DataPlaneUri() terra.StringValue {
	return terra.ReferenceString(lt.ref.Append("data_plane_uri"))
}

func (lt loadTestAttributes) Description() terra.StringValue {
	return terra.ReferenceString(lt.ref.Append("description"))
}

func (lt loadTestAttributes) Id() terra.StringValue {
	return terra.ReferenceString(lt.ref.Append("id"))
}

func (lt loadTestAttributes) Location() terra.StringValue {
	return terra.ReferenceString(lt.ref.Append("location"))
}

func (lt loadTestAttributes) Name() terra.StringValue {
	return terra.ReferenceString(lt.ref.Append("name"))
}

func (lt loadTestAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(lt.ref.Append("resource_group_name"))
}

func (lt loadTestAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](lt.ref.Append("tags"))
}

func (lt loadTestAttributes) Identity() terra.ListValue[loadtest.IdentityAttributes] {
	return terra.ReferenceList[loadtest.IdentityAttributes](lt.ref.Append("identity"))
}

func (lt loadTestAttributes) Timeouts() loadtest.TimeoutsAttributes {
	return terra.ReferenceSingle[loadtest.TimeoutsAttributes](lt.ref.Append("timeouts"))
}

type loadTestState struct {
	DataPlaneUri      string                   `json:"data_plane_uri"`
	Description       string                   `json:"description"`
	Id                string                   `json:"id"`
	Location          string                   `json:"location"`
	Name              string                   `json:"name"`
	ResourceGroupName string                   `json:"resource_group_name"`
	Tags              map[string]string        `json:"tags"`
	Identity          []loadtest.IdentityState `json:"identity"`
	Timeouts          *loadtest.TimeoutsState  `json:"timeouts"`
}
