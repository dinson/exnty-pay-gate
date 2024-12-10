package enum

type Txn string
type TxnStatus string // success, failed, cancelled, initialized

const (
	TxnDeposit    Txn = "deposit"
	TxnWithdrawal Txn = "withdrawal"

	TxnStatusSuccess     TxnStatus = "success"
	TxnStatusFailed      TxnStatus = "failed"
	TxnStatusCancelled   TxnStatus = "cancelled"
	TxnStatusInitialized TxnStatus = "initialized"
)

func (t Txn) String() string { return string(t) }

func (t TxnStatus) String() string { return string(t) }
