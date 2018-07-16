package ethereum

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// TransactionStatus is receipt status of transaction.
type TransactionStatus uint

const (
	// TransactionFailed is the status code of a transaction if execution failed.
	TransactionFailed = TransactionStatus(0)

	// TransactionSuccessful is the status code of a transaction if execution succeeded.
	TransactionSuccessful = TransactionStatus(1)
)

// Block holds information about Ethereum block.
type Block struct {
	Difficulty   *big.Int
	ExtraData    []byte
	GasLimit     *big.Int
	GasUsed      *big.Int
	Hash         common.Hash
	Miner        common.Address
	Number       *big.Int
	Timestamp    *big.Int
	Transactions Transactions
}

func (b *Block) String() string {
	str := fmt.Sprintf(`Block(#%v): {
Difficulty:	    %v
ExtraData:      %s
GasLimit:	    %v
GasUsed:	    %v
Hash:           %x
Miner:          %x
Timestamp:      %v
Transactions:
%v
}
`, b.Number, b.Difficulty, b.ExtraData, b.GasLimit, b.GasUsed, b.Hash[:], b.Miner[:], b.Timestamp, b.Transactions)
	return str
}

// Transactions slice type.
type Transactions []*Transaction

// Transaction holds information about Ethereum transaction.
type Transaction struct {
	BlockNumber      *big.Int
	From             common.Address
	GasLimit         *big.Int
	GasPrice         *big.Int
	GasUsed          *big.Int
	Hash             common.Hash
	Input            []byte
	Nonce            uint64
	To               *common.Address // nil means contract creation
	TransactionIndex uint64
	Value            *big.Int
	ContractAddress  *common.Address
	Status           *TransactionStatus
}

func (t *Transaction) String() string {
	var to string

	if t.To == nil {
		to = "[contract creation]"
	} else {
		to = fmt.Sprintf("%x", t.To[:])
	}

	var status string
	if t.Status == nil {
		status = "[unknown status]"
	} else {
		if *t.Status == TransactionSuccessful {
			status = "successful"
		} else {
			status = "failed"
		}
	}

	return fmt.Sprintf(`
	TX(%s)
	BlockNumber:     %#v
	TxIndex:         %v
	Contract:        %v
	ContractAddress: %v
	From:            %x
	To:              %s
	GasPrice:        %#v
	GasLimit:        %#v
	GasUsed:         %#v
	Input:           %x
	Nonce:           %v
	Value:           %#v
	Status:          %v
`,
		t.Hash.String(),
		t.BlockNumber,
		t.TransactionIndex,
		t.To == nil,
		t.ContractAddress,
		t.From[:],
		to,
		t.GasPrice,
		t.GasLimit,
		t.GasUsed,
		t.Input,
		t.Nonce,
		t.Value,
		status,
	)
}