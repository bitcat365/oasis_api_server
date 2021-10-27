package handlers

import (
	"context"
	"encoding/json"
	lgr "github.com/SimplyVC/oasis_api_server/src/logger"
	"github.com/SimplyVC/oasis_api_server/src/responses"
	"github.com/SimplyVC/oasis_api_server/src/rpc"
	"github.com/oasisprotocol/oasis-core/go/common"
	common_namespace "github.com/oasisprotocol/oasis-core/go/common"
	runtime "github.com/oasisprotocol/oasis-core/go/runtime/client/api"
	"google.golang.org/grpc"
	"net/http"
	"strconv"
)

// loadRuntimeClient loads runtime client and returns it
func loadRuntimeClient(socket string) (*grpc.ClientConn, runtime.RuntimeClient) {

	// Attempt to load connection with runtime client
	connection, runtimeClient, err := rpc.RuntimeClient(socket)
	if err != nil {
		lgr.Error.Println("Failed to establish connection to runtime"+
			" client: ", err)
		return nil, nil
	}
	return connection, runtimeClient
}

// GetRuntimeBlock returns the events at specified block height.
func GetRuntimeBlock(w http.ResponseWriter, r *http.Request) {

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
	recvRound := r.URL.Query().Get("round")
	if len(recvRound) == 0 {

		// Stop code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Unexepcted value found, height needs to be " +
				"string of int!"})
		return
	}
	var round uint64
	round, _ = (strconv.ParseUint(recvRound, 10, 64))

	var id common.Namespace
	_id := r.URL.Query().Get("id")
	if err := id.UnmarshalHex(_id); err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "failed to decode runtime id"})
		return
	}

	// Attempt to load connection with runtime client
	connection, ro := loadRuntimeClient(socket)

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

	query := runtime.GetBlockRequest{RuntimeID: id, Round: round}

	// Retrieving events object using above query
	runtimeBlock, err := ro.GetBlock(context.Background(), &query)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to get Runtime Events!"})
		lgr.Error.Println("Request at /api/runtime/block failed "+
			"to retrieve Runtime Events : ", err)
		return
	}

	// Responding with events object retrieved above
	lgr.Info.Println("Request at /api/runtime/block responding with " +
		"Runtime Events!")
	json.NewEncoder(w).Encode(responses.RuntimeBlockResponse{
		RuntimeBlock: runtimeBlock})
}

// GetRuntimeTransactions returns the events at specified block height.
func GetRuntimeTransactions(w http.ResponseWriter, r *http.Request) {

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
	recvRound := r.URL.Query().Get("round")
	if len(recvRound) == 0 {

		// Stop code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Unexepcted value found, height needs to be " +
				"string of int!"})
		return
	}
	var round uint64
	round, _ = (strconv.ParseUint(recvRound, 10, 64))

	// Note Make sure that private key that is being sent is coded properly
	// Example A1X90rT/WK4AOTh/dJsUlOqNDV/nXM6ZU+h+blS9pto= should be
	// A1X90rT/WK4AOTh/dJsUlOqNDV/nXM6ZU%2Bh%2BblS9pto=
	var nameSpace common_namespace.Namespace
	nmspace := r.URL.Query().Get("namespace")
	if len(nmspace) == 0 {
		// Stop code here no need to establish connection and reply
		lgr.Warning.Println("Request at /api/runtime/transactions failed" +
			", namespace can't be empty!")
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "namespace can't be empty!"})
		return
	}

	// Unmarshal received text into namespace object
	err := nameSpace.UnmarshalText([]byte(nmspace))
	if err != nil {
		lgr.Error.Println("Failed to UnmarshalText into Namespace", err)
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to UnmarshalText into Namespace."})
		return
	}

	// Attempt to load connection with runtime client
	connection, ro := loadRuntimeClient(socket)

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

	query := runtime.GetTransactionsRequest{RuntimeID: nameSpace, Round: round}

	// Retrieving events object using above query
	runtimeTransactions, err := ro.GetTransactions(context.Background(), &query)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to get Runtime Transactions!"})
		lgr.Error.Println("Request at /api/runtime/transactions failed "+
			"to retrieve Runtime Transactions : ", err)
		return
	}

	// Responding with events object retrieved above
	lgr.Info.Println("Request at /api/runtime/transactions responding with " +
		"Runtime Transactions!")
	json.NewEncoder(w).Encode(responses.RuntimeTransactionsResponse{
		RuntimeTransactions: runtimeTransactions})
}

// GetRuntimeEvents returns the events at specified block height.
func GetRuntimeEvents(w http.ResponseWriter, r *http.Request) {

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
	recvRound := r.URL.Query().Get("round")
	if len(recvRound) == 0 {

		// Stop code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Unexepcted value found, height needs to be " +
				"string of int!"})
		return
	}
	var round uint64
	round, _ = (strconv.ParseUint(recvRound, 10, 64))

	// Note Make sure that private key that is being sent is coded properly
	// Example A1X90rT/WK4AOTh/dJsUlOqNDV/nXM6ZU+h+blS9pto= should be
	// A1X90rT/WK4AOTh/dJsUlOqNDV/nXM6ZU%2Bh%2BblS9pto=
	var nameSpace common_namespace.Namespace
	nmspace := r.URL.Query().Get("namespace")
	if len(nmspace) == 0 {
		// Stop code here no need to establish connection and reply
		lgr.Warning.Println("Request at /api/runtime/events failed" +
			", namespace can't be empty!")
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "namespace can't be empty!"})
		return
	}

	// Unmarshal received text into namespace object
	err := nameSpace.UnmarshalText([]byte(nmspace))
	if err != nil {
		lgr.Error.Println("Failed to UnmarshalText into Namespace", err)
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to UnmarshalText into Namespace."})
		return
	}

	// Attempt to load connection with runtime client
	connection, ro := loadRuntimeClient(socket)

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

	query := runtime.GetEventsRequest{RuntimeID: nameSpace, Round: round}

	// Retrieving events object using above query
	runtimeEvents, err := ro.GetEvents(context.Background(), &query)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to get Runtime Events!"})
		lgr.Error.Println("Request at /api/runtime/events failed "+
			"to retrieve Runtime Events : ", err)
		return
	}

	// Responding with events object retrieved above
	lgr.Info.Println("Request at /api/runtime/events responding with " +
		"Runtime Events!")
	json.NewEncoder(w).Encode(responses.RuntimeEventsResponse{
		RuntimeEvents: runtimeEvents})
}

// RuntimeQuery method.
func RuntimeQuery(w http.ResponseWriter, r *http.Request) {

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
	recvRound := r.URL.Query().Get("round")
	if len(recvRound) == 0 {

		// Stop code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to retrieve round, not specified!"})
		return
	}
	var round uint64
	round, _ = (strconv.ParseUint(recvRound, 10, 64))

	// Note Make sure that private key that is being sent is coded properly
	// Example A1X90rT/WK4AOTh/dJsUlOqNDV/nXM6ZU+h+blS9pto= should be
	// A1X90rT/WK4AOTh/dJsUlOqNDV/nXM6ZU%2Bh%2BblS9pto=
	var nameSpace common_namespace.Namespace
	nmspace := r.URL.Query().Get("namespace")
	if len(nmspace) == 0 {
		// Stop code here no need to establish connection and reply
		lgr.Warning.Println("Request at /api/runtime/query failed" +
			", namespace can't be empty!")
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "namespace can't be empty!"})
		return
	}

	// Unmarshal received text into namespace object
	err := nameSpace.UnmarshalText([]byte(nmspace))
	if err != nil {
		lgr.Error.Println("Failed to UnmarshalText into Namespace", err)
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to UnmarshalText into Namespace."})
		return
	}

	// Retrieving method from query request
	method := r.URL.Query().Get("method")
	if len(method) == 0 {

		// Stop code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to retrieve method, not specified!"})
		return
	}

	// Attempt to load connection with runtime client
	connection, ro := loadRuntimeClient(socket)

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

	args := r.URL.Query().Get("args")
	if args == "" {
		// Stop code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "No Consensus Key Provided"})
		return
	}

	query := runtime.QueryRequest{RuntimeID: nameSpace, Round: round, Method: method, Args: []byte(args)}

	// Retrieving query object using above query
	runtimeQuery, err := ro.Query(context.Background(), &query)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to get Runtime Query!"})
		lgr.Error.Println("Request at /api/runtime/query failed "+
			"to retrieve Runtime Query : ", err)
		return
	}

	// Responding with query object retrieved above
	lgr.Info.Println("Request at /api/runtime/query responding with " +
		"Runtime Query!")
	json.NewEncoder(w).Encode(responses.RuntimeQueryResponse{
		RuntimeQuery: runtimeQuery})
}
