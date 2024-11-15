package ipfs

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ipld/go-car/v2/blockstore"
)

type Performance struct {
	Assigned int `json:"assigned"`
	Included int `json:"included"`
}

type Validator struct {
	Perf    Performance `json:"perf"`
	Slashed bool        `json:"slashed"`
}

type Operator struct {
	Distributed int64             `json:"distributed"`
	Stuck       bool              `json:"stuck"`
	Validators  map[int]Validator `json:"validators"`
}

type Blockstamp struct {
	BlockHash      string `json:"block_hash"`
	BlockNumber    int    `json:"block_number"`
	BlockTimestamp int64  `json:"block_timestamp"`
	RefEpoch       int    `json:"ref_epoch"`
	RefSlot        int    `json:"ref_slot"`
	SlotNumber     int    `json:"slot_number"`
	StateRoot      string `json:"state_root"`
}

type Report struct {
	Blockstamp    Blockstamp       `json:"blockstamp"`
	Distributable int64            `json:"distributable"`
	Frame         [2]int           `json:"frame"`
	Operators     map[int]Operator `json:"operators"`
	Threshold     float64          `json:"threshold"`
}

type IPFSAdapter struct {
	GatewayURL string
}

func NewIPFSAdapter(gatewayURL string) *IPFSAdapter {
	return &IPFSAdapter{
		GatewayURL: gatewayURL,
	}
}

// FetchAndParseIpfs fetches data from IPFS using a CID, decodes it from CAR format, and returns the content as a Report.
func (ia *IPFSAdapter) FetchAndParseIpfs(cid string) (Report, error) {
	url := fmt.Sprintf("%s/ipfs/%s?format=car", ia.GatewayURL, cid)
	fmt.Printf("Fetching IPFS data from URL: %s\n", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return Report{}, fmt.Errorf("failed to create request: %w", err)
	}

	// Set the Accept header to request CAR format
	req.Header.Set("Accept", "application/vnd.ipld.car")
	fmt.Println("Request header set to accept CAR format")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error fetching data from IPFS: %v\n", err)
		return Report{}, fmt.Errorf("failed to fetch data from IPFS: %w", err)
	}
	defer resp.Body.Close()

	fmt.Printf("Received response with status: %s\n", resp.Status)
	if resp.StatusCode != http.StatusOK {
		return Report{}, fmt.Errorf("unexpected response status: %s", resp.Status)
	}

	// Read the CAR data into a buffer
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return Report{}, fmt.Errorf("failed to read response body: %w", err)
	}
	fmt.Printf("Read %d bytes from response\n", len(body))

	// Open a CAR blockstore from the data
	store, err := blockstore.NewReadOnly(bytes.NewReader(body), nil)
	if err != nil {
		fmt.Printf("Error opening CAR blockstore: %v\n", err)
		return Report{}, fmt.Errorf("failed to open CAR blockstore: %w", err)
	}
	fmt.Println("Successfully opened CAR blockstore")

	// Extract the root CID(s) from the CAR file
	roots, err := store.Roots()
	if err != nil {
		fmt.Printf("Error retrieving roots from CAR file: %v\n", err)
		return Report{}, fmt.Errorf("failed to get roots from CAR file: %w", err)
	}
	fmt.Printf("Retrieved %d root(s) from CAR file\n", len(roots))

	// Assuming the first root contains the report data we need
	root := roots[0]
	fmt.Printf("Using root CID: %s\n", root)

	block, err := store.Get(context.Background(), root)
	if err != nil {
		fmt.Printf("Error retrieving block from CAR file: %v\n", err)
		return Report{}, fmt.Errorf("failed to get block from CAR file: %w", err)
	}
	fmt.Println("Successfully retrieved block data")

	// print block.RawData
	fmt.Printf("Block raw data: %s\n", block.RawData())

	// Parse the block data as JSON to the Report struct
	var report Report
	err = json.Unmarshal(block.RawData(), &report)
	if err != nil {
		fmt.Printf("Error unmarshalling block data to Report: %v\n", err)
		return Report{}, fmt.Errorf("failed to unmarshal block data: %w", err)
	}
	fmt.Println("Successfully unmarshalled block data to Report struct")

	return report, nil
}
