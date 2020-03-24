package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"google.golang.org/grpc"

	lgr "github.com/SimplyVC/oasis_api_server/src/logger"
	"github.com/SimplyVC/oasis_api_server/src/responses"
	"github.com/SimplyVC/oasis_api_server/src/rpc"
	"github.com/oasislabs/oasis-core/go/common/cbor"
	consensus "github.com/oasislabs/oasis-core/go/consensus/api"
	mint_api "github.com/oasislabs/oasis-core/go/consensus/tendermint/api"
)

// loadConsensusClient loads the consensus client and returns it
func loadConsensusClient(socket string) (*grpc.ClientConn, consensus.ClientBackend) {
	// Attempt to load a connection with the consensus client
	connection, consensusClient, err := rpc.ConsensusClient(socket)
	if err != nil {
		lgr.Error.Println("Failed to establish connection to the consensus client : ", err)
		return nil, nil
	}
	return connection, consensusClient
}

// GetConsensusStateToGenesis returns the genesis state at the specified block height for Consensus.
func GetConsensusStateToGenesis(w http.ResponseWriter, r *http.Request) {
	// Adding a header so that the receiver knows they are receiving a JSON structure
	w.Header().Add("Content-Type", "application/json")

	// Retrieving the name of the node from the query request
	nodeName := r.URL.Query().Get("name")
	confirmation, socket := checkNodeName(nodeName)
	if confirmation == false {
		// Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Node name requested doesn't exist"})
		return
	}

	// Retrieve the height from the query
	recvHeight := r.URL.Query().Get("height")
	height := checkHeight(recvHeight)
	if height == -1 {
		// Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Unexepcted value found, height needs to be string of int!"})
		return
	}

	// Attempt to load a connection with the consensus client
	connection, co := loadConsensusClient(socket)

	// Wait for the code underneath it to execute and then close the connection
	defer connection.Close()

	// If a null object was retrieved send response
	if co == nil {
		// Stop the code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to establish a connection using the socket : " + socket})
		return
	}

	// Retrieving the genesis state of the consensus object at the specified height
	consensusGenesis, err := co.StateToGenesis(context.Background(), height)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to get Genesis file of Block!"})
		lgr.Error.Println("Request at /api/GetStateToGenesis/ Failed to retrieve the genesis file : ", err)
		return
	}

	// Responding with the consensus genesis state object, retrieved above.
	lgr.Info.Println("Request at /api/GetStateToGenesis/ responding with a genesis file!")
	json.NewEncoder(w).Encode(responses.ConsensusGenesisResponse{GenJSON: consensusGenesis})
}

// GetEpoch returns the current epoch of a given block height
func GetEpoch(w http.ResponseWriter, r *http.Request) {
	// Adding a header so that the receiver knows they are receiving a JSON structure
	w.Header().Add("Content-Type", "application/json")

	// Retrieving the name of the node from the query request
	nodeName := r.URL.Query().Get("name")
	confirmation, socket := checkNodeName(nodeName)
	if confirmation == false {
		// Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Node name requested doesn't exist"})
		return
	}

	// Retrieve the height from the query
	recvHeight := r.URL.Query().Get("height")
	height := checkHeight(recvHeight)
	if height == -1 {
		// Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Unexepcted value found, height needs to be string of int!"})
		return
	}

	// Attempt to load a connection with the consensus client
	connection, co := loadConsensusClient(socket)

	// Wait for the code underneath it to execute and then close the connection
	defer connection.Close()

	// If a null object was retrieved send response
	if co == nil {
		// Stop the code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to establish a connection using the socket : " + socket})
		return
	}

	// Return the epcoh of the specific height
	epoch, err := co.GetEpoch(context.Background(), height)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to retrieve Epoch of Block!"})
		lgr.Error.Println("Request at /api/consensus/epoch/ Failed to retrieve Epoch : ", err)
		return
	}

	// Respond with the retrieved epoch above
	lgr.Info.Println("Request at /api/consensus/epoch/ responding with an Epoch!")
	json.NewEncoder(w).Encode(responses.EpochResponse{Ep: epoch})
}

// PingNode returns a consensus block at a specific height thus signifying that it was pinged.
func PingNode(w http.ResponseWriter, r *http.Request) {
	// Adding a header so that the receiver knows they are receiving a JSON structure
	w.Header().Add("Content-Type", "application/json")

	// Retrieving the name of the node from the query request
	nodeName := r.URL.Query().Get("name")
	confirmation, socket := checkNodeName(nodeName)
	if confirmation == false {
		// Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Node name requested doesn't exist"})
		return
	}

	// Setting the height to latest
	height := consensus.HeightLatest

	// Attempt to load a connection with the consensus client
	connection, co := loadConsensusClient(socket)

	// Wait for the code underneath it to execute and then close the connection
	defer connection.Close()

	// If a null object was retrieved send response
	if co == nil {
		// Stop the code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to establish a connection using the socket : " + socket})
		return
	}

	// Making sure that the error being retrieved is nill meaning that the api is pingable
	_, err := co.GetBlock(context.Background(), height)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to ping a node by retrieving heighest block height!"})
		lgr.Error.Println("Request at /api/pingnode/ Failed to ping node : ", err)
		return
	}

	// Responding with a Pong Response
	lgr.Info.Println("Request at /api/pingnode/ responding with Pong!")
	json.NewEncoder(w).Encode(responses.SuccessResponsed)
}

// GetBlock returns a consensus block at a specific height.
func GetBlock(w http.ResponseWriter, r *http.Request) {
	// Adding a header so that the receiver knows they are receiving a JSON structure
	w.Header().Add("Content-Type", "application/json")

	// Retrieving the name of the node from the query request
	nodeName := r.URL.Query().Get("name")
	confirmation, socket := checkNodeName(nodeName)
	if confirmation == false {
		// Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Node name requested doesn't exist"})
		return
	}

	// Retrieve the height from the query
	recvHeight := r.URL.Query().Get("height")
	height := checkHeight(recvHeight)
	if height == -1 {
		// Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Unexepcted value found, height needs to be string of int!"})
		return
	}

	// Attempt to load a connection with the consensus client
	connection, co := loadConsensusClient(socket)

	// Wait for the code underneath it to execute and then close the connection
	defer connection.Close()

	// If a null object was retrieved send response
	if co == nil {
		// Stop the code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to establish a connection using the socket : " + socket})
		return
	}

	// Retrieve the block at the specific height from the consensus client
	blk, err := co.GetBlock(context.Background(), height)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to retrieve Block!"})
		lgr.Error.Println("Request at /api/consensus/block/ Failed to retrieve a Block : ", err)
		return
	}

	// Responding with the retrieved block
	lgr.Info.Println("Request at /api/consensus/block/ responding with a Block!")
	json.NewEncoder(w).Encode(responses.BlockResponse{Blk: blk})
}

// GetBlockHeader returns a consensus block header at a specific height
func GetBlockHeader(w http.ResponseWriter, r *http.Request) {
	// Adding a header so that the receiver knows they are receiving a JSON structure
	w.Header().Add("Content-Type", "application/json")

	// Retrieving the name of the node from the query request
	nodeName := r.URL.Query().Get("name")
	confirmation, socket := checkNodeName(nodeName)
	if confirmation == false {
		// Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Node name requested doesn't exist"})
		return
	}

	// Retrieving the height from the query
	recvHeight := r.URL.Query().Get("height")
	height := checkHeight(recvHeight)
	if height == -1 {
		// Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Unexepcted value found, height needs to be string of int!"})
		return
	}

	// Attempt to load a connection with the consensus client
	connection, co := loadConsensusClient(socket)

	// Wait for the code underneath it to execute and then close the connection
	defer connection.Close()

	// If a null object was retrieved send response
	if co == nil {
		// Stop the code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to establish a connection using the socket : " + socket})
		return
	}

	// Retriving the Block at a specific height using the Consensus client
	blk, err := co.GetBlock(context.Background(), height)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to retrieve Block!"})
		lgr.Error.Println("Request at /api/consensus/blockheader/ Failed to retrieve a Block : ", err)
		return
	}

	// Creating a BlockMeta object
	var meta mint_api.BlockMeta
	if err := cbor.Unmarshal(blk.Meta, &meta); err != nil {
		lgr.Error.Println("Request at /api/consensus/blockheader/ Failed to Unmarshal Block Metadata : ", err)
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to Unmarshal Block Metadata!"})
		return
	}

	// Responds with block header retrieved above
	lgr.Info.Println("Request at /api/consensus/blockheader/ responding with a Block Header!")
	json.NewEncoder(w).Encode(responses.BlockHeaderResponse{BlkHeader: meta.Header})
}

// GetBlockLastCommit returns a consensus block last commit at a specific height
func GetBlockLastCommit(w http.ResponseWriter, r *http.Request) {
	// Adding a header so that the receiver knows they are receiving a JSON structure
	w.Header().Add("Content-Type", "application/json")

	// Retrieving the name of the node from the query request
	nodeName := r.URL.Query().Get("name")
	confirmation, socket := checkNodeName(nodeName)
	if confirmation == false {
		// Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Node name requested doesn't exist"})
		return
	}

	// Retrieving the height from the query
	recvHeight := r.URL.Query().Get("height")
	height := checkHeight(recvHeight)
	if height == -1 {
		// Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Unexepcted value found, height needs to be string of int!"})
		return
	}

	// Attempt to load a connection with the consensus client
	connection, co := loadConsensusClient(socket)

	// Wait for the code underneath it to execute and then close the connection
	defer connection.Close()

	// If a null object was retrieved send response
	if co == nil {
		// Stop the code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to establish a connection using the socket : " + socket})
		return
	}

	// Retrieve the block at a specific height from the consensus client
	blk, err := co.GetBlock(context.Background(), height)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to retrieve Block!"})
		lgr.Error.Println("Request at /api/consensus/blocklastcommit/ Failed to retrieve a Block : ", err)
		return
	}

	// Creating the BlockMeta object
	var meta mint_api.BlockMeta
	if err := cbor.Unmarshal(blk.Meta, &meta); err != nil {
		lgr.Error.Println("Request at /api/consensus/blocklastcommit/ Failed to Unmarshal Block Metadata : ", err)
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to Unmarshal Block Metadata!"})
		return
	}

	// Responds with the Block Last commit retrieved above
	lgr.Info.Println("Request at /api/consensus/blocklastcommit/ responding with a Block Last Commit!")
	json.NewEncoder(w).Encode(responses.BlockLastCommitResponse{BlkLastCommit: meta.LastCommit})
}

// GetTransactions returns a consensus block header at a specific height
func GetTransactions(w http.ResponseWriter, r *http.Request) {
	// Adding a header so that the receiver knows they are receiving a JSON structure
	w.Header().Add("Content-Type", "application/json")

	// Retrieving the name of the node from the query request
	nodeName := r.URL.Query().Get("name")
	confirmation, socket := checkNodeName(nodeName)
	if confirmation == false {
		// Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Node name requested doesn't exist"})
		return
	}

	// Retrieving the height from the query
	recvHeight := r.URL.Query().Get("height")
	height := checkHeight(recvHeight)
	if height == -1 {
		// Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Unexepcted value found, height needs to be string of int!"})
		return
	}

	// Attempt to load a connection with the consensus client
	connection, co := loadConsensusClient(socket)

	// Wait for the code underneath it to execute and then close the connection
	defer connection.Close()

	// If a null object was retrieved send response
	if co == nil {
		// Stop the code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to establish a connection using the socket : " + socket})
		return
	}

	// Use the consensus client to retrieve transactions at a specific block height
	transactions, err := co.GetTransactions(context.Background(), height)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to retrieve Transactions!"})
		lgr.Error.Println("Request at /api/consensus/transactions/ Failed to retrieve Transactions : ", err)
		return
	}

	// Responds with the transactions retrieved above
	lgr.Info.Println("Request at /api/consensus/transactions/ responding with all the transactions in the specified Block!")
	json.NewEncoder(w).Encode(responses.TransactionsResponse{Transactions: transactions})
}
