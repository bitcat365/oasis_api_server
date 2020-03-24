package responses

import (
	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/disk"
	"github.com/mackerelio/go-osstat/memory"
	"github.com/mackerelio/go-osstat/network"
	common_signature "github.com/oasislabs/oasis-core/go/common/crypto/signature"
	common_entity "github.com/oasislabs/oasis-core/go/common/entity"
	common_node "github.com/oasislabs/oasis-core/go/common/node"
	common_quantity "github.com/oasislabs/oasis-core/go/common/quantity"
	consensus_api "github.com/oasislabs/oasis-core/go/consensus/api"
	epoch_api "github.com/oasislabs/oasis-core/go/epochtime/api"
	gen_api "github.com/oasislabs/oasis-core/go/genesis/api"
	registry_api "github.com/oasislabs/oasis-core/go/registry/api"
	scheduler_api "github.com/oasislabs/oasis-core/go/scheduler/api"
	staking_api "github.com/oasislabs/oasis-core/go/staking/api"
	mint_types "github.com/tendermint/tendermint/types"
)

// NetworkResponse responds with the network statistics of the system
type NetworkResponse struct {
	Network []network.Stats `json:"result"`
}

// CPUResponse responds with the CPU statistics of the system
type CPUResponse struct {
	CPU *cpu.Stats `json:"result"`
}

// DiskResponse responds with the memory statistics of the system
type DiskResponse struct {
	Disk []disk.Stats `json:"result"`
}

// MemoryResponse responds with the memory statistics of the system
type MemoryResponse struct {
	Memory *memory.Stats `json:"result"`
}

// SchedulerGenesisState responds with the scheduler genesis state
type SchedulerGenesisState struct {
	SchedulerGenesisState *scheduler_api.Genesis `json:"result"`
}

// CommitteesResponse responds with the Committees
type CommitteesResponse struct {
	Committee []*scheduler_api.Committee `json:"result"`
}

// ValidatorsResponse responds with Validators and their voting power
type ValidatorsResponse struct {
	Validators []*scheduler_api.Validator `json:"result"`
}

// IsSyncedResponse responds with a boolean signifying the synchronisation state of a node
type IsSyncedResponse struct {
	Synced bool `json:"result"`
}

// DebondingDelegationsResponse responds with debonding delegations for a specified public key
type DebondingDelegationsResponse struct {
	DebondingDelegations map[common_signature.PublicKey][]*staking_api.DebondingDelegation `json:"result"`
}

// DelegationsResponse responds with delegations for a public key
type DelegationsResponse struct {
	Delegations map[common_signature.PublicKey]*staking_api.Delegation `json:"result"`
}

// AccountResponse responds with an account
type AccountResponse struct {
	AccountInfo *staking_api.Account `json:"result"`
}

// AllAccountsResponse responds with a list of Accounts
type AllAccountsResponse struct {
	AllAccounts []common_signature.PublicKey `json:"result"`
}

// StakingGenesisResponse responds with a Staking Genesis File
type StakingGenesisResponse struct {
	GenesisStaking *staking_api.Genesis `json:"result"`
}

// QuantityResponse responds with a quantity
type QuantityResponse struct {
	Quantity *common_quantity.Quantity `json:"result"`
}

// RegistryEntityResponse responds with the details of a single Entity
type RegistryEntityResponse struct {
	Entity *common_entity.Entity `json:"result"`
}

// RegistryNodeResponse responds with the details of a single Node
type RegistryNodeResponse struct {
	Node *common_node.Node `json:"result"`
}

// RegistryGenesisResponse responds with the genesis state of the registry
type RegistryGenesisResponse struct {
	GenesisRegistry *registry_api.Genesis `json:"result"`
}

// NodelistResponse responds with a NodeList
type NodelistResponse struct {
	NodeList *registry_api.NodeList `json:"result"`
}

// RuntimeResponse responds with a single Runtime
type RuntimeResponse struct {
	Runtime *registry_api.Runtime `json:"result"`
}

// RuntimesResponse responds with Multiple Runtimes
type RuntimesResponse struct {
	Runtimes []*registry_api.Runtime `json:"result"`
}

// NodesResponse responding with Multiple Nodes
type NodesResponse struct {
	Nodes []*common_node.Node `json:"result"`
}

// EntitiesResponse responding with Multiple Entities
type EntitiesResponse struct {
	Entities []*common_entity.Entity `json:"result"`
}

// TransactionsResponse responds with all the transactions in a block
type TransactionsResponse struct {
	Transactions [][]byte `json:"result"`
}

// BlockHeaderResponse responds with a Tendermint Header Type
type BlockHeaderResponse struct {
	BlkHeader *mint_types.Header `json:"result"`
}

// BlockLastCommitResponse responds with a Tendermint Last Commit Type
type BlockLastCommitResponse struct {
	BlkLastCommit *mint_types.Commit `json:"result"`
}

// BlockResponse responds with a custom Block Response with an unmarshalled message
type BlockResponse struct {
	Blk *consensus_api.Block `json:"result"`
}

// EpochResponse responds with epcoh time
type EpochResponse struct {
	Ep epoch_api.EpochTime `json:"result"`
}

// ConsensusGenesisResponse with the consensus Genesis Document
type ConsensusGenesisResponse struct {
	GenJSON *gen_api.Document `json:"result"`
}

// SuccessResponse with a succeful result
type SuccessResponse struct {
	Result string `json:"result"`
}

// ErrorResponse repsonds with an error message that will be set
type ErrorResponse struct {
	Error string `json:"error"`
}

// ConnectionsResponse responds with all the connections configured
type ConnectionsResponse struct {
	Results []string `json:"result"`
}

// SuccessResponsed Assinging Variable Responses that do not need to be changed.
var SuccessResponsed = SuccessResponse{Result: "pong"}
