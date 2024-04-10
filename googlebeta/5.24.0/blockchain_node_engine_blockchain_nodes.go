// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	blockchainnodeengineblockchainnodes "github.com/golingon/terraproviders/googlebeta/5.24.0/blockchainnodeengineblockchainnodes"
	"io"
)

// NewBlockchainNodeEngineBlockchainNodes creates a new instance of [BlockchainNodeEngineBlockchainNodes].
func NewBlockchainNodeEngineBlockchainNodes(name string, args BlockchainNodeEngineBlockchainNodesArgs) *BlockchainNodeEngineBlockchainNodes {
	return &BlockchainNodeEngineBlockchainNodes{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BlockchainNodeEngineBlockchainNodes)(nil)

// BlockchainNodeEngineBlockchainNodes represents the Terraform resource google_blockchain_node_engine_blockchain_nodes.
type BlockchainNodeEngineBlockchainNodes struct {
	Name      string
	Args      BlockchainNodeEngineBlockchainNodesArgs
	state     *blockchainNodeEngineBlockchainNodesState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BlockchainNodeEngineBlockchainNodes].
func (bnebn *BlockchainNodeEngineBlockchainNodes) Type() string {
	return "google_blockchain_node_engine_blockchain_nodes"
}

// LocalName returns the local name for [BlockchainNodeEngineBlockchainNodes].
func (bnebn *BlockchainNodeEngineBlockchainNodes) LocalName() string {
	return bnebn.Name
}

// Configuration returns the configuration (args) for [BlockchainNodeEngineBlockchainNodes].
func (bnebn *BlockchainNodeEngineBlockchainNodes) Configuration() interface{} {
	return bnebn.Args
}

// DependOn is used for other resources to depend on [BlockchainNodeEngineBlockchainNodes].
func (bnebn *BlockchainNodeEngineBlockchainNodes) DependOn() terra.Reference {
	return terra.ReferenceResource(bnebn)
}

// Dependencies returns the list of resources [BlockchainNodeEngineBlockchainNodes] depends_on.
func (bnebn *BlockchainNodeEngineBlockchainNodes) Dependencies() terra.Dependencies {
	return bnebn.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BlockchainNodeEngineBlockchainNodes].
func (bnebn *BlockchainNodeEngineBlockchainNodes) LifecycleManagement() *terra.Lifecycle {
	return bnebn.Lifecycle
}

// Attributes returns the attributes for [BlockchainNodeEngineBlockchainNodes].
func (bnebn *BlockchainNodeEngineBlockchainNodes) Attributes() blockchainNodeEngineBlockchainNodesAttributes {
	return blockchainNodeEngineBlockchainNodesAttributes{ref: terra.ReferenceResource(bnebn)}
}

// ImportState imports the given attribute values into [BlockchainNodeEngineBlockchainNodes]'s state.
func (bnebn *BlockchainNodeEngineBlockchainNodes) ImportState(av io.Reader) error {
	bnebn.state = &blockchainNodeEngineBlockchainNodesState{}
	if err := json.NewDecoder(av).Decode(bnebn.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bnebn.Type(), bnebn.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BlockchainNodeEngineBlockchainNodes] has state.
func (bnebn *BlockchainNodeEngineBlockchainNodes) State() (*blockchainNodeEngineBlockchainNodesState, bool) {
	return bnebn.state, bnebn.state != nil
}

// StateMust returns the state for [BlockchainNodeEngineBlockchainNodes]. Panics if the state is nil.
func (bnebn *BlockchainNodeEngineBlockchainNodes) StateMust() *blockchainNodeEngineBlockchainNodesState {
	if bnebn.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bnebn.Type(), bnebn.LocalName()))
	}
	return bnebn.state
}

// BlockchainNodeEngineBlockchainNodesArgs contains the configurations for google_blockchain_node_engine_blockchain_nodes.
type BlockchainNodeEngineBlockchainNodesArgs struct {
	// BlockchainNodeId: string, required
	BlockchainNodeId terra.StringValue `hcl:"blockchain_node_id,attr" validate:"required"`
	// BlockchainType: string, optional
	BlockchainType terra.StringValue `hcl:"blockchain_type,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// ConnectionInfo: min=0
	ConnectionInfo []blockchainnodeengineblockchainnodes.ConnectionInfo `hcl:"connection_info,block" validate:"min=0"`
	// EthereumDetails: optional
	EthereumDetails *blockchainnodeengineblockchainnodes.EthereumDetails `hcl:"ethereum_details,block"`
	// Timeouts: optional
	Timeouts *blockchainnodeengineblockchainnodes.Timeouts `hcl:"timeouts,block"`
}
type blockchainNodeEngineBlockchainNodesAttributes struct {
	ref terra.Reference
}

// BlockchainNodeId returns a reference to field blockchain_node_id of google_blockchain_node_engine_blockchain_nodes.
func (bnebn blockchainNodeEngineBlockchainNodesAttributes) BlockchainNodeId() terra.StringValue {
	return terra.ReferenceAsString(bnebn.ref.Append("blockchain_node_id"))
}

// BlockchainType returns a reference to field blockchain_type of google_blockchain_node_engine_blockchain_nodes.
func (bnebn blockchainNodeEngineBlockchainNodesAttributes) BlockchainType() terra.StringValue {
	return terra.ReferenceAsString(bnebn.ref.Append("blockchain_type"))
}

// CreateTime returns a reference to field create_time of google_blockchain_node_engine_blockchain_nodes.
func (bnebn blockchainNodeEngineBlockchainNodesAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(bnebn.ref.Append("create_time"))
}

// EffectiveLabels returns a reference to field effective_labels of google_blockchain_node_engine_blockchain_nodes.
func (bnebn blockchainNodeEngineBlockchainNodesAttributes) EffectiveLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](bnebn.ref.Append("effective_labels"))
}

// Id returns a reference to field id of google_blockchain_node_engine_blockchain_nodes.
func (bnebn blockchainNodeEngineBlockchainNodesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bnebn.ref.Append("id"))
}

// Labels returns a reference to field labels of google_blockchain_node_engine_blockchain_nodes.
func (bnebn blockchainNodeEngineBlockchainNodesAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](bnebn.ref.Append("labels"))
}

// Location returns a reference to field location of google_blockchain_node_engine_blockchain_nodes.
func (bnebn blockchainNodeEngineBlockchainNodesAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(bnebn.ref.Append("location"))
}

// Name returns a reference to field name of google_blockchain_node_engine_blockchain_nodes.
func (bnebn blockchainNodeEngineBlockchainNodesAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(bnebn.ref.Append("name"))
}

// Project returns a reference to field project of google_blockchain_node_engine_blockchain_nodes.
func (bnebn blockchainNodeEngineBlockchainNodesAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(bnebn.ref.Append("project"))
}

// TerraformLabels returns a reference to field terraform_labels of google_blockchain_node_engine_blockchain_nodes.
func (bnebn blockchainNodeEngineBlockchainNodesAttributes) TerraformLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](bnebn.ref.Append("terraform_labels"))
}

// UpdateTime returns a reference to field update_time of google_blockchain_node_engine_blockchain_nodes.
func (bnebn blockchainNodeEngineBlockchainNodesAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(bnebn.ref.Append("update_time"))
}

func (bnebn blockchainNodeEngineBlockchainNodesAttributes) ConnectionInfo() terra.ListValue[blockchainnodeengineblockchainnodes.ConnectionInfoAttributes] {
	return terra.ReferenceAsList[blockchainnodeengineblockchainnodes.ConnectionInfoAttributes](bnebn.ref.Append("connection_info"))
}

func (bnebn blockchainNodeEngineBlockchainNodesAttributes) EthereumDetails() terra.ListValue[blockchainnodeengineblockchainnodes.EthereumDetailsAttributes] {
	return terra.ReferenceAsList[blockchainnodeengineblockchainnodes.EthereumDetailsAttributes](bnebn.ref.Append("ethereum_details"))
}

func (bnebn blockchainNodeEngineBlockchainNodesAttributes) Timeouts() blockchainnodeengineblockchainnodes.TimeoutsAttributes {
	return terra.ReferenceAsSingle[blockchainnodeengineblockchainnodes.TimeoutsAttributes](bnebn.ref.Append("timeouts"))
}

type blockchainNodeEngineBlockchainNodesState struct {
	BlockchainNodeId string                                                     `json:"blockchain_node_id"`
	BlockchainType   string                                                     `json:"blockchain_type"`
	CreateTime       string                                                     `json:"create_time"`
	EffectiveLabels  map[string]string                                          `json:"effective_labels"`
	Id               string                                                     `json:"id"`
	Labels           map[string]string                                          `json:"labels"`
	Location         string                                                     `json:"location"`
	Name             string                                                     `json:"name"`
	Project          string                                                     `json:"project"`
	TerraformLabels  map[string]string                                          `json:"terraform_labels"`
	UpdateTime       string                                                     `json:"update_time"`
	ConnectionInfo   []blockchainnodeengineblockchainnodes.ConnectionInfoState  `json:"connection_info"`
	EthereumDetails  []blockchainnodeengineblockchainnodes.EthereumDetailsState `json:"ethereum_details"`
	Timeouts         *blockchainnodeengineblockchainnodes.TimeoutsState         `json:"timeouts"`
}
