package hedera

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestScheduleCreateTransaction_Execute(t *testing.T) {
	client := newTestClient(t, false)

	newKey, err := GeneratePrivateKey()
	assert.NoError(t, err)

	newBalance := NewHbar(1)

	assert.Equal(t, HbarUnits.Hbar.numberOfTinybar(), newBalance.tinybar)

	transactionID := TransactionIDGenerate(client.GetOperatorAccountID())

	tx := NewAccountCreateTransaction().
		SetTransactionID(transactionID).
		SetKey(newKey.PublicKey()).
		SetMaxTransactionFee(NewHbar(2)).
		SetInitialBalance(newBalance)

	assert.NoError(t, err)

	scheduleTx, err := tx.Schedule()
	assert.NoError(t, err)

	resp, err := scheduleTx.
		SetPayerAccountID(client.GetOperatorAccountID()).
		SetAdminKey(client.GetOperatorPublicKey()).
		Execute(client)
	assert.NoError(t, err)

	receipt, err := resp.GetReceipt(client)
	assert.NoError(t, err)

	info, err := NewScheduleInfoQuery().
		SetScheduleID(*receipt.ScheduleID).
		SetQueryPayment(NewHbar(2)).
		Execute(client)
	assert.NoError(t, err)

	infoTx, err := info.GetTransaction()
	assert.NoError(t, err)

	assert.NotNil(t, infoTx)

	tx2, err := NewScheduleDeleteTransaction().
		SetScheduleID(*receipt.ScheduleID).
		FreezeWith(client)
	assert.NoError(t, err)

	resp, err = tx2.
		Sign(newKey).
		Execute(client)
	assert.NoError(t, err)

	_, err = resp.GetReceipt(client)
	assert.NoError(t, err)
}

func TestScheduleCreateTransaction_SetTransaction_Execute(t *testing.T) {
	client := newTestClient(t, false)

	newKey, err := GeneratePrivateKey()
	assert.NoError(t, err)

	newBalance := NewHbar(1)

	assert.Equal(t, HbarUnits.Hbar.numberOfTinybar(), newBalance.tinybar)

	transactionID := TransactionIDGenerate(client.GetOperatorAccountID())

	tx := NewAccountCreateTransaction().
		SetKey(newKey.PublicKey()).
		SetTransactionID(transactionID).
		SetMaxTransactionFee(NewHbar(2)).
		SetInitialBalance(newBalance)

	scheduleTx, err := NewScheduleCreateTransaction().
		SetPayerAccountID(client.GetOperatorAccountID()).
		SetAdminKey(client.GetOperatorPublicKey()).
		SetScheduledTransaction(tx)
	assert.NoError(t, err)

	resp, err := scheduleTx.Execute(client)
	assert.NoError(t, err)

	receipt, err := resp.GetReceipt(client)
	assert.NoError(t, err)

	_, err = NewScheduleInfoQuery().
		SetScheduleID(*receipt.ScheduleID).
		Execute(client)
	assert.NoError(t, err)

	tx2, err := NewScheduleDeleteTransaction().
		SetScheduleID(*receipt.ScheduleID).
		FreezeWith(client)
	assert.NoError(t, err)

	resp, err = tx2.
		Sign(newKey).
		Execute(client)
	assert.NoError(t, err)

	_, err = resp.GetReceipt(client)
	assert.NoError(t, err)
}


func TestScheduleCreateTransaction_Signature_Execute(t *testing.T) {
	client := newTestClient(t, false)

	newKey, err := GeneratePrivateKey()
	assert.NoError(t, err)

	newBalance := NewHbar(1)

	assert.Equal(t, HbarUnits.Hbar.numberOfTinybar(), newBalance.tinybar)

	accountTransactionID := TransactionIDGenerate(client.GetOperatorAccountID())

	tx, err := NewAccountCreateTransaction().
		SetKey(newKey.PublicKey()).
		SetMaxTransactionFee(NewHbar(2)).
		SetInitialBalance(newBalance).
		SetTransactionID(accountTransactionID).
		SetNodeAccountIDs([]AccountID{{Account: 3}}).
		FreezeWith(client)
	assert.NoError(t, err)

	txBytes, err := tx.ToBytes()
	assert.NoError(t, err)

	signature1, err := newKey.SignTransaction(&tx.Transaction)

	tx2, err := TransactionFromBytes(txBytes)
	assert.NoError(t, err)

	var tx3 *ScheduleCreateTransaction

	switch newTx := tx2.(type) {
	case AccountCreateTransaction:
		tx3, err = newTx.Schedule()
		assert.NoError(t, err)
	}

	assert.NotNil(t, tx3)

	tx3, err = tx3.
		SetAdminKey(newKey.PublicKey()).
		SetPayerAccountID(client.GetOperatorAccountID()).
		SetTransactionValidDuration(30 * time.Second).
		FreezeWith(client)
	assert.NoError(t, err)

	tx3 = tx3.Sign(newKey)

	resp, err := tx3.Execute(client)
	assert.NoError(t, err)

	receipt, err := resp.GetReceipt(client)
	assert.NoError(t, err)

	scheduleID := *receipt.ScheduleID

	resp, err = NewScheduleSignTransaction().
		SetScheduleID(scheduleID).
		SetNodeAccountIDs([]AccountID{resp.NodeID}).
		AddScheduleSignature(newKey.PublicKey(), signature1).
		Execute(client)
	assert.NoError(t, err)

	receipt, err = resp.GetReceipt(client)
	assert.NoError(t, err)

	//tx4, err := NewScheduleDeleteTransaction().
	//	SetScheduleID(scheduleID).
	//	FreezeWith(client)
	//assert.NoError(t, err)
	//
	//resp, err = tx4.
	//	Sign(newKey).
	//	Execute(client)
	//assert.NoError(t, err)
	//
	//_, err = resp.GetReceipt(client)
	//assert.NoError(t, err)
}