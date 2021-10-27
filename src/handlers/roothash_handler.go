package handlers

import (
	"context"
	"encoding/json"
	lgr "github.com/SimplyVC/oasis_api_server/src/logger"
	"github.com/SimplyVC/oasis_api_server/src/responses"
	"github.com/SimplyVC/oasis_api_server/src/rpc"
	common_namespace "github.com/oasisprotocol/oasis-core/go/common"
	roothash "github.com/oasisprotocol/oasis-core/go/roothash/api"
	"google.golang.org/grpc"
	"net/http"
)

// loadRoothashClient loads roothash client and returns it
func loadRoothashClient(socket string) (*grpc.ClientConn, roothash.Backend) {

	// Attempt to load connection with roothash client
	connection, roothashClient, err := rpc.RoothashClient(socket)
	if err != nil {
		lgr.Error.Println("Failed to establish connection to roothash"+
			" client: ", err)
		return nil, nil
	}
	return connection, roothashClient
}

// GetLatestBlock returns the latest block.
//
// The metadata contained in this block can be further used to get
// the latest state from the storage backend.
func GetLatestBlock(w http.ResponseWriter, r *http.Request) {

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

	// Note Make sure that private key that is being sent is coded properly
	// Example A1X90rT/WK4AOTh/dJsUlOqNDV/nXM6ZU+h+blS9pto= should be
	// A1X90rT/WK4AOTh/dJsUlOqNDV/nXM6ZU%2Bh%2BblS9pto=
	var nameSpace common_namespace.Namespace
	nmspace := r.URL.Query().Get("namespace")
	if len(nmspace) == 0 {
		// Stop code here no need to establish connection and reply
		lgr.Warning.Println("Request at /api/roothash/latestblock failed" +
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

	// Attempt to load connection with roothash client
	connection, ro := loadRoothashClient(socket)

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

	// Creating query that will be used to return runtime by it's namespace
	query := roothash.RuntimeRequest{Height: height, RuntimeID: nameSpace}

	// Retrieving latest block object using above query
	roothashLatestBlock, err := ro.GetLatestBlock(context.Background(), &query)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to get Roothash Latest Block!"})
		lgr.Error.Println("Request at /api/roothash/latestblock failed "+
			"to retrieve Roothash Latest Block : ", err)
		return
	}

	// Responding with latestblock object retrieved above
	lgr.Info.Println("Request at /api/roothash/latestblock responding with " +
		"Roothash Latest Block!")
	json.NewEncoder(w).Encode(responses.RoothashLatestBlockResponse{
		LatestBlock: roothashLatestBlock})
}

// GetRuntimeState returns the given runtime's state.
func GetRuntimeState(w http.ResponseWriter, r *http.Request) {

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

	// Note Make sure that private key that is being sent is coded properly
	// Example A1X90rT/WK4AOTh/dJsUlOqNDV/nXM6ZU+h+blS9pto= should be
	// A1X90rT/WK4AOTh/dJsUlOqNDV/nXM6ZU%2Bh%2BblS9pto=
	var nameSpace common_namespace.Namespace
	nmspace := r.URL.Query().Get("namespace")
	if len(nmspace) == 0 {
		// Stop code here no need to establish connection and reply
		lgr.Warning.Println("Request at /api/roothash/runtimestate failed" +
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

	// Attempt to load connection with roothash client
	connection, ro := loadRoothashClient(socket)

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

	// Creating query that will be used to return runtime by it's namespace
	query := roothash.RuntimeRequest{Height: height, RuntimeID: nameSpace}

	// Retrieving runtime state object using above query
	roothashRuntimeState, err := ro.GetRuntimeState(context.Background(), &query)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to get Roothash Runtime State!"})
		lgr.Error.Println("Request at /api/roothash/runtimestate failed "+
			"to retrieve Roothash Runtime State : ", err)
		return
	}

	// Responding with runtime object retrieved above
	lgr.Info.Println("Request at /api/roothash/runtimestate responding with " +
		"Roothash Runtime State!")
	json.NewEncoder(w).Encode(responses.RuntimeStateResponse{
		RuntimeState: roothashRuntimeState})
}

// GetRoothashEvents returns the events at specified block height.
func GetRoothashEvents(w http.ResponseWriter, r *http.Request) {

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

	// Attempt to load connection with roothash client
	connection, ro := loadRoothashClient(socket)

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

	// Retrieving events object using above query
	roothashEvents, err := ro.GetEvents(context.Background(), height)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to get Roothash Events!"})
		lgr.Error.Println("Request at /api/roothash/events failed "+
			"to retrieve Roothash Events : ", err)
		return
	}

	// Responding with events object retrieved above
	lgr.Info.Println("Request at /api/roothash/events responding with " +
		"Roothash Events!")
	json.NewEncoder(w).Encode(responses.RoothashEventsResponse{
		RoothashEvents: roothashEvents})
}
