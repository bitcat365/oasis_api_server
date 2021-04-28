package handlers

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc"
	"net/http"
	"strconv"

	lgr "github.com/SimplyVC/oasis_api_server/src/logger"
	"github.com/SimplyVC/oasis_api_server/src/responses"
	"github.com/SimplyVC/oasis_api_server/src/rpc"
	governance "github.com/oasisprotocol/oasis-core/go/governance/api"
)

// loadGovernanceClient loads governance client and returns it
func loadGovernanceClient(socket string) (*grpc.ClientConn, governance.Backend) {

	// Attempt to load connection with governance client
	connection, governanceClient, err := rpc.GovernanceClient(socket)
	if err != nil {
		lgr.Error.Println("Failed to establish connection to governance"+
			" client: ", err)
		return nil, nil
	}
	return connection, governanceClient
}

// GetActiveProposals returns a list of all proposals that have not yet closed.
func GetActiveProposals(w http.ResponseWriter, r *http.Request) {

	// Add header so that received knows they're receiving JSON
	w.Header().Add("Content-Type", "application/json")

	// Retrieving name of node from query request
	nodeName := r.URL.Query().Get("name")
	confirmation, socket := checkNodeName(nodeName)
	if confirmation == false {

		// Stop code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Node name requested doesn't exist"})
		return
	}

	// Retrieving height from query request
	recvHeight := r.URL.Query().Get("height")
	height := checkHeight(recvHeight)
	if height == -1 {

		// Stop code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Unexepcted value found, height needs to be " +
				"string of int!"})
		return
	}

	// Attempt to load connection with governance client
	connection, ro := loadGovernanceClient(socket)

	// Close connection once code underneath executes
	defer connection.Close()

	// If null object was retrieved send response
	if ro == nil {

		// Stop code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to establish connection using socket: " +
				socket})
		return
	}

	// Retrieve ActiveProposals at specific block height
	proposals, err := ro.ActiveProposals(context.Background(), height)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to get ActiveProposals!"})
		lgr.Error.Println("Request at /api/governance/activeproposals failed "+
			"to retrieve ActiveProposals : ", err)
		return
	}

	// Responding with retrieved ActiveProposals
	lgr.Info.Println("Request at /api/governance/activeproposals responding with" +
		" ActiveProposals!")
	json.NewEncoder(w).Encode(responses.ProposalsResponse{
		Proposals: proposals})
}

// GetProposals returns a list of all proposals.
func GetProposals(w http.ResponseWriter, r *http.Request) {

	// Add header so that received knows they're receiving JSON
	w.Header().Add("Content-Type", "application/json")

	// Retrieving name of node from query request
	nodeName := r.URL.Query().Get("name")
	confirmation, socket := checkNodeName(nodeName)
	if confirmation == false {

		// Stop code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Node name requested doesn't exist"})
		return
	}

	// Retrieving height from query request
	recvHeight := r.URL.Query().Get("height")
	height := checkHeight(recvHeight)
	if height == -1 {

		// Stop code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Unexepcted value found, height needs to be " +
				"string of int!"})
		return
	}

	// Attempt to load connection with governance client
	connection, ro := loadGovernanceClient(socket)

	// Close connection once code underneath executes
	defer connection.Close()

	// If null object was retrieved send response
	if ro == nil {

		// Stop code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to establish connection using socket: " +
				socket})
		return
	}

	// Retrieve Proposals at specific block height
	proposals, err := ro.Proposals(context.Background(), height)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to get Proposals!"})
		lgr.Error.Println("Request at /api/governance/proposals failed "+
			"to retrieve Proposals : ", err)
		return
	}

	// Responding with retrieved Proposals
	lgr.Info.Println("Request at /api/governance/proposals responding with" +
		" Proposals!")
	json.NewEncoder(w).Encode(responses.ProposalsResponse{
		Proposals: proposals})
}

// GetProposal looks up a specific proposal.
func GetProposal(w http.ResponseWriter, r *http.Request) {

	// Add header so that received knows they're receiving JSON
	w.Header().Add("Content-Type", "application/json")

	// Retrieving name of node from query request
	nodeName := r.URL.Query().Get("name")
	confirmation, socket := checkNodeName(nodeName)
	if confirmation == false {

		// Stop code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Node name requested doesn't exist"})
		return
	}

	// Retrieving height from query request
	recvHeight := r.URL.Query().Get("height")
	height := checkHeight(recvHeight)
	if height == -1 {

		// Stop code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Unexepcted value found, height needs to be " +
				"string of int!"})
		return
	}

	// Attempt to load connection with governance client
	connection, ro := loadGovernanceClient(socket)

	// Close connection once code underneath executes
	defer connection.Close()

	// If null object was retrieved send response
	if ro == nil {

		// Stop code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to establish connection using socket: " +
				socket})
		return
	}

	id := r.URL.Query().Get("id")
	if len(id) == 0 {

		// Stop code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Unexpected value found, id needs to be a string representing an int!"})
		return
	}
	var proposalId uint64
	proposalId, _ = strconv.ParseUint(id, 10, 64)

	query := governance.ProposalQuery{Height: height, ProposalID: proposalId}

	// Retrieve Proposals at specific block height
	proposal, err := ro.Proposal(context.Background(), &query)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to get Proposal!"})
		lgr.Error.Println("Request at /api/governance/proposal failed "+
			"to retrieve Proposal : ", err)
		return
	}

	// Responding with retrieved Proposals
	lgr.Info.Println("Request at /api/governance/proposal responding with" +
		" Proposal!")
	json.NewEncoder(w).Encode(responses.ProposalResponse{
		Proposal: proposal})
}

// GetVotes looks up a specific proposal.
func GetVotes(w http.ResponseWriter, r *http.Request) {

	// Add header so that received knows they're receiving JSON
	w.Header().Add("Content-Type", "application/json")

	// Retrieving name of node from query request
	nodeName := r.URL.Query().Get("name")
	confirmation, socket := checkNodeName(nodeName)
	if confirmation == false {

		// Stop code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Node name requested doesn't exist"})
		return
	}

	// Retrieving height from query request
	recvHeight := r.URL.Query().Get("height")
	height := checkHeight(recvHeight)
	if height == -1 {

		// Stop code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Unexepcted value found, height needs to be " +
				"string of int!"})
		return
	}

	// Attempt to load connection with governance client
	connection, ro := loadGovernanceClient(socket)

	// Close connection once code underneath executes
	defer connection.Close()

	// If null object was retrieved send response
	if ro == nil {

		// Stop code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to establish connection using socket: " +
				socket})
		return
	}

	id := r.URL.Query().Get("id")
	if len(id) == 0 {

		// Stop code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Unexpected value found, id needs to be a string representing an int!"})
		return
	}
	var proposalId uint64
	proposalId, _ = strconv.ParseUint(id, 10, 64)

	query := governance.ProposalQuery{Height: height, ProposalID: proposalId}

	// Retrieve Votes at a specific proposal
	votes, err := ro.Votes(context.Background(), &query)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to get Votes!"})
		lgr.Error.Println("Request at /api/governance/votes failed "+
			"to retrieve Votes : ", err)
		return
	}

	// Responding with retrieved Votes
	lgr.Info.Println("Request at /api/governance/votes responding with" +
		" Votes!")
	json.NewEncoder(w).Encode(responses.VotesResponse{
		Votes: votes})
}
