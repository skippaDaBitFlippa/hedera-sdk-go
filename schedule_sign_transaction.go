package hedera

import (
	protobuf "github.com/golang/protobuf/proto"
	"time"

	"github.com/hashgraph/hedera-sdk-go/v2/proto"
)

type ScheduleSignTransaction struct {
	Transaction
	pb *proto.ScheduleSignTransactionBody
}

func NewScheduleSignTransaction() *ScheduleSignTransaction {
	pb := &proto.ScheduleSignTransactionBody{}

	transaction := ScheduleSignTransaction{
		pb:          pb,
		Transaction: newTransaction(),
	}

	transaction.SetMaxTransactionFee(NewHbar(5))

	return &transaction
}

func scheduleSignTransactionFromProtobuf(transaction Transaction, pb *proto.TransactionBody) ScheduleSignTransaction {
	return ScheduleSignTransaction{
		Transaction: transaction,
		pb:          pb.GetScheduleSign(),
	}
}

func (transaction *ScheduleSignTransaction) SetScheduleID(id ScheduleID) *ScheduleSignTransaction {
	transaction.requireNotFrozen()
	transaction.pb.ScheduleID = id.toProtobuf()

	return transaction
}

func (transaction *ScheduleSignTransaction) GetScheduleID() ScheduleID {
	return scheduleIDFromProtobuf(transaction.pb.ScheduleID)
}

func (transaction *ScheduleSignTransaction) AddScheduleSignature(publicKey PublicKey, signature []byte) *ScheduleSignTransaction {
	transaction.requireOneNodeAccountID()

	if transaction.pb.SigMap != nil {
		if transaction.pb.SigMap.SigPair != nil {
			transaction.pb.SigMap.SigPair = append(transaction.pb.SigMap.SigPair, publicKey.toSignaturePairProtobuf(signature))
		} else {
			transaction.pb.SigMap.SigPair = make([]*proto.SignaturePair, 0)
			transaction.pb.SigMap.SigPair = append(transaction.pb.SigMap.SigPair, publicKey.toSignaturePairProtobuf(signature))
		}
	} else {
		transaction.pb.SigMap = &proto.SignatureMap{
			SigPair: make([]*proto.SignaturePair, 0),
		}
		transaction.pb.SigMap.SigPair = append(transaction.pb.SigMap.SigPair, publicKey.toSignaturePairProtobuf(signature))
	}

	return transaction
}

func (transaction *ScheduleSignTransaction) GetScheduleSignatures() (map[*PublicKey][]byte, error) {
	signMap := make(map[*PublicKey][]byte, len(transaction.pb.GetSigMap().GetSigPair()))

	for _, sigPair := range transaction.pb.GetSigMap().GetSigPair() {
		key, err := PublicKeyFromBytes(sigPair.PubKeyPrefix)
		if err != nil {
			return make(map[*PublicKey][]byte, 0), err
		}
		switch sigPair.Signature.(type) {
		case *proto.SignaturePair_Contract:
			signMap[&key] = sigPair.GetContract()
		case *proto.SignaturePair_Ed25519:
			signMap[&key] = sigPair.GetEd25519()
		case *proto.SignaturePair_RSA_3072:
			signMap[&key] = sigPair.GetRSA_3072()
		case *proto.SignaturePair_ECDSA_384:
			signMap[&key] = sigPair.GetECDSA_384()
		}
	}

	return signMap, nil
}

func (transaction *ScheduleSignTransaction) Schedule() (*ScheduleCreateTransaction, error) {
	transaction.requireNotFrozen()

	txBytes, err := protobuf.Marshal(transaction.constructProtobuf())
	if err != nil {
		return &ScheduleCreateTransaction{}, err
	}

	return NewScheduleCreateTransaction().setTransactionBodyBytes(txBytes), nil
}

func (transaction *ScheduleSignTransaction) constructProtobuf() *proto.TransactionBody {
	return &proto.TransactionBody{
		TransactionID:            transaction.GetTransactionID().SetScheduled(true).toProtobuf(),
		NodeAccountID:            transaction.pbBody.GetNodeAccountID(),
		TransactionFee:           transaction.pbBody.GetTransactionFee(),
		TransactionValidDuration: transaction.pbBody.GetTransactionValidDuration(),
		GenerateRecord:           transaction.pbBody.GetGenerateRecord(),
		Memo:                     transaction.pbBody.GetMemo(),
		Data: &proto.TransactionBody_ScheduleSign{
			ScheduleSign: &proto.ScheduleSignTransactionBody{
				ScheduleID: transaction.pb.GetScheduleID(),
				SigMap:     transaction.pb.GetSigMap(),
			},
		},
	}
}

//
// The following methods must be copy-pasted/overriden at the bottom of **every** _transaction.go file
// We override the embedded fluent setter methods to return the outer type
//

func scheduleSignTransaction_getMethod(request request, channel *channel) method {
	return method{
		transaction: channel.getSchedule().SignSchedule,
	}
}

func (transaction *ScheduleSignTransaction) IsFrozen() bool {
	return transaction.isFrozen()
}

// Sign uses the provided privateKey to sign the transaction.
func (transaction *ScheduleSignTransaction) Sign(
	privateKey PrivateKey,
) *ScheduleSignTransaction {
	return transaction.SignWith(privateKey.PublicKey(), privateKey.Sign)
}

func (transaction *ScheduleSignTransaction) SignWithOperator(
	client *Client,
) (*ScheduleSignTransaction, error) {
	// If the transaction is not signed by the operator, we need
	// to sign the transaction with the operator

	if client == nil {
		return nil, errNoClientProvided
	} else if client.operator == nil {
		return nil, errClientOperatorSigning
	}
	if !transaction.IsFrozen() {
		_, err := transaction.FreezeWith(client)
		if err != nil {
			return transaction, err
		}
	}
	return transaction.SignWith(client.operator.publicKey, client.operator.signer), nil
}

// SignWith executes the TransactionSigner and adds the resulting signature data to the Transaction's signature map
// with the publicKey as the map key.
func (transaction *ScheduleSignTransaction) SignWith(
	publicKey PublicKey,
	signer TransactionSigner,
) *ScheduleSignTransaction {
	if !transaction.IsFrozen() {
		_, _ = transaction.Freeze()
	} else {
		transaction.transactions = make([]*proto.Transaction, 0)
	}

	if transaction.keyAlreadySigned(publicKey) {
		return transaction
	}

	for index := 0; index < len(transaction.signedTransactions); index++ {
		signature := signer(transaction.signedTransactions[index].GetBodyBytes())

		transaction.signedTransactions[index].SigMap.SigPair = append(
			transaction.signedTransactions[index].SigMap.SigPair,
			publicKey.toSignaturePairProtobuf(signature),
		)
	}

	return transaction
}

// Execute executes the Transaction with the provided client
func (transaction *ScheduleSignTransaction) Execute(
	client *Client,
) (TransactionResponse, error) {
	if client == nil || client.operator == nil {
		return TransactionResponse{}, errNoClientProvided
	}

	if transaction.freezeError != nil {
		return TransactionResponse{}, transaction.freezeError
	}

	if !transaction.IsFrozen() {
		_, err := transaction.FreezeWith(client)
		if err != nil {
			return TransactionResponse{}, err
		}
	}

	transactionID := transaction.GetTransactionID()

	if !client.GetOperatorAccountID().isZero() && client.GetOperatorAccountID().equals(*transactionID.AccountID) {
		transaction.SignWith(
			client.GetOperatorPublicKey(),
			client.operator.signer,
		)
	}

	resp, err := execute(
		client,
		request{
			transaction: &transaction.Transaction,
		},
		transaction_shouldRetry,
		transaction_makeRequest,
		transaction_advanceRequest,
		transaction_getNodeAccountID,
		scheduleSignTransaction_getMethod,
		transaction_mapResponseStatus,
		transaction_mapResponse,
	)

	if err != nil {
		return TransactionResponse{
			TransactionID: transaction.GetTransactionID(),
			NodeID:        resp.transaction.NodeID,
		}, err
	}

	hash, err := transaction.GetTransactionHash()

	return TransactionResponse{
		TransactionID: transaction.GetTransactionID(),
		NodeID:        resp.transaction.NodeID,
		Hash:          hash,
	}, nil
}

func (transaction *ScheduleSignTransaction) onFreeze(
	pbBody *proto.TransactionBody,
) bool {
	pbBody.Data = &proto.TransactionBody_ScheduleSign{
		ScheduleSign: transaction.pb,
	}

	return true
}

func (transaction *ScheduleSignTransaction) Freeze() (*ScheduleSignTransaction, error) {
	return transaction.FreezeWith(nil)
}

func (transaction *ScheduleSignTransaction) FreezeWith(client *Client) (*ScheduleSignTransaction, error) {
	if transaction.IsFrozen() {
		return transaction, nil
	}
	transaction.initFee(client)
	if err := transaction.initTransactionID(client); err != nil {
		return transaction, err
	}

	if !transaction.onFreeze(transaction.pbBody) {
		return transaction, nil
	}

	return transaction, transaction_freezeWith(&transaction.Transaction, client)
}

func (transaction *ScheduleSignTransaction) GetMaxTransactionFee() Hbar {
	return transaction.Transaction.GetMaxTransactionFee()
}

// SetMaxTransactionFee sets the max transaction fee for this ScheduleSignTransaction.
func (transaction *ScheduleSignTransaction) SetMaxTransactionFee(fee Hbar) *ScheduleSignTransaction {
	transaction.requireNotFrozen()
	transaction.Transaction.SetMaxTransactionFee(fee)
	return transaction
}

func (transaction *ScheduleSignTransaction) GetTransactionMemo() string {
	return transaction.Transaction.GetTransactionMemo()
}

// SetTransactionMemo sets the memo for this ScheduleSignTransaction.
func (transaction *ScheduleSignTransaction) SetTransactionMemo(memo string) *ScheduleSignTransaction {
	transaction.requireNotFrozen()
	transaction.Transaction.SetTransactionMemo(memo)
	return transaction
}

func (transaction *ScheduleSignTransaction) GetTransactionValidDuration() time.Duration {
	return transaction.Transaction.GetTransactionValidDuration()
}

// SetTransactionValidDuration sets the valid duration for this ScheduleSignTransaction.
func (transaction *ScheduleSignTransaction) SetTransactionValidDuration(duration time.Duration) *ScheduleSignTransaction {
	transaction.requireNotFrozen()
	transaction.Transaction.SetTransactionValidDuration(duration)
	return transaction
}

func (transaction *ScheduleSignTransaction) GetTransactionID() TransactionID {
	return transaction.Transaction.GetTransactionID()
}

// SetTransactionID sets the TransactionID for this ScheduleSignTransaction.
func (transaction *ScheduleSignTransaction) SetTransactionID(transactionID TransactionID) *ScheduleSignTransaction {
	transaction.requireNotFrozen()

	transaction.Transaction.SetTransactionID(transactionID)
	return transaction
}

// SetNodeAccountID sets the node AccountID for this ScheduleSignTransaction.
func (transaction *ScheduleSignTransaction) SetNodeAccountIDs(nodeID []AccountID) *ScheduleSignTransaction {
	transaction.requireNotFrozen()
	transaction.Transaction.SetNodeAccountIDs(nodeID)
	return transaction
}

func (transaction *ScheduleSignTransaction) SetMaxRetry(count int) *ScheduleSignTransaction {
	transaction.Transaction.SetMaxRetry(count)
	return transaction
}