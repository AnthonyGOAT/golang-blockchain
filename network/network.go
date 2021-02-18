package network

const (
    protocol = "tcp"
    version = 1
    commandLength = 12
)

var (
    nodeAddress string
    minerAddress string
    KnownNodes = []string{"localhost:3000"}
    blocksInTransit = [][]byte{}
    memoryPool = make(map[string]blockchain.Transaction)
)
