package hedera

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTokenDissociateTransaction_Execute(t *testing.T) {
	client := newTestClient(t)

	newKey, err := GeneratePrivateKey()
	assert.NoError(t, err)

	newBalance := NewHbar(1)

	assert.Equal(t, HbarUnits.Hbar.numberOfTinybar(), newBalance.tinybar)

	resp, err := NewAccountCreateTransaction().
		SetKey(newKey.PublicKey()).
		SetMaxTransactionFee(NewHbar(2)).
		SetInitialBalance(newBalance).
		Execute(client)
	assert.NoError(t, err)

	receipt, err := resp.GetReceipt(client)
	assert.NoError(t, err)

	accountID := *receipt.AccountID

	resp, err = NewTokenCreateTransaction().
		SetTokenName("ffff").
		SetTokenSymbol("F").
		SetDecimals(3).
		SetInitialSupply(1000000).
		SetTreasuryAccountID(client.GetOperatorAccountID()).
		SetAdminKey(client.GetOperatorPublicKey()).
		SetFreezeKey(client.GetOperatorPublicKey()).
		SetWipeKey(client.GetOperatorPublicKey()).
		SetKycKey(client.GetOperatorPublicKey()).
		SetSupplyKey(client.GetOperatorPublicKey()).
		SetFreezeDefault(false).
		Execute(client)
	assert.NoError(t, err)

	receipt, err = resp.GetReceipt(client)
	assert.NoError(t, err)

	tokenID := *receipt.TokenID

	associateTx, err := NewTokenAssociateTransaction().
		SetNodeAccountIDs([]AccountID{resp.NodeID}).
		SetAccountID(accountID).
		SetTokenIDs(tokenID).
		FreezeWith(client)
	assert.NoError(t, err)

	resp, err = associateTx.
		Sign(newKey).
		Execute(client)
	assert.NoError(t, err)

	_, err = resp.GetReceipt(client)
	assert.NoError(t, err)

	dissociateTx, err := NewTokenDissociateTransaction().
		SetNodeAccountIDs([]AccountID{resp.NodeID}).
		SetAccountID(accountID).
		SetTokenIDs(tokenID).
		FreezeWith(client)
	assert.NoError(t, err)

	resp, err = dissociateTx.
		Sign(newKey).
		Execute(client)
	assert.NoError(t, err)

	_, err = resp.GetReceipt(client)
	assert.NoError(t, err)

	resp, err = NewTokenDeleteTransaction().
		SetNodeAccountIDs([]AccountID{resp.NodeID}).
		SetTokenID(tokenID).
		Execute(client)
	assert.NoError(t, err)

	_, err = resp.GetReceipt(client)
	assert.NoError(t, err)

	tx, err := NewAccountDeleteTransaction().
		SetAccountID(accountID).
		SetTransferAccountID(client.GetOperatorAccountID()).
		FreezeWith(client)
	assert.NoError(t, err)

	resp, err = tx.
		Sign(newKey).
		Execute(client)
	assert.NoError(t, err)

	_, err = resp.GetReceipt(client)
	assert.NoError(t, err)
}

func Test_TokenDissociate_NoSigningOne(t *testing.T) {
	client := newTestClient(t)

	newKey, err := GeneratePrivateKey()
	assert.NoError(t, err)

	newBalance := NewHbar(1)

	assert.Equal(t, HbarUnits.Hbar.numberOfTinybar(), newBalance.tinybar)

	resp, err := NewAccountCreateTransaction().
		SetKey(newKey.PublicKey()).
		SetMaxTransactionFee(NewHbar(2)).
		SetInitialBalance(newBalance).
		Execute(client)
	assert.NoError(t, err)

	receipt, err := resp.GetReceipt(client)
	assert.NoError(t, err)

	accountID := *receipt.AccountID

	resp, err = NewTokenCreateTransaction().
		SetTokenName("ffff").
		SetTokenSymbol("F").
		SetDecimals(3).
		SetInitialSupply(1000000).
		SetTreasuryAccountID(client.GetOperatorAccountID()).
		SetAdminKey(client.GetOperatorPublicKey()).
		SetFreezeKey(client.GetOperatorPublicKey()).
		SetWipeKey(client.GetOperatorPublicKey()).
		SetKycKey(client.GetOperatorPublicKey()).
		SetSupplyKey(client.GetOperatorPublicKey()).
		SetFreezeDefault(false).
		Execute(client)
	assert.NoError(t, err)

	receipt, err = resp.GetReceipt(client)
	assert.NoError(t, err)

	tokenID := *receipt.TokenID

	_, err = NewTokenAssociateTransaction().
		SetNodeAccountIDs([]AccountID{resp.NodeID}).
		SetAccountID(accountID).
		SetTokenIDs(tokenID).
		Execute(client)
	assert.NoError(t, err)

	_, err = resp.GetReceipt(client)
	assert.NoError(t, err)

	_, err = NewTokenDissociateTransaction().
		SetNodeAccountIDs([]AccountID{resp.NodeID}).
		SetAccountID(accountID).
		SetTokenIDs(tokenID).
		Execute(client)
	assert.NoError(t, err)

	_, err = resp.GetReceipt(client)
	assert.NoError(t, err)

	resp, err = NewTokenDeleteTransaction().
		SetNodeAccountIDs([]AccountID{resp.NodeID}).
		SetTokenID(tokenID).
		Execute(client)
	assert.NoError(t, err)

	_, err = resp.GetReceipt(client)
	assert.NoError(t, err)

	_, err = NewAccountDeleteTransaction().
		SetAccountID(accountID).
		SetTransferAccountID(client.GetOperatorAccountID()).
		Execute(client)
	assert.NoError(t, err)

	_, err = resp.GetReceipt(client)
	assert.NoError(t, err)
}

func Test_TokenDissociate_NoTokenID(t *testing.T) {
	client := newTestClient(t)

	newKey, err := GeneratePrivateKey()
	assert.NoError(t, err)

	newBalance := NewHbar(1)

	assert.Equal(t, HbarUnits.Hbar.numberOfTinybar(), newBalance.tinybar)

	resp, err := NewAccountCreateTransaction().
		SetKey(newKey.PublicKey()).
		SetMaxTransactionFee(NewHbar(2)).
		SetInitialBalance(newBalance).
		Execute(client)
	assert.NoError(t, err)

	receipt, err := resp.GetReceipt(client)
	assert.NoError(t, err)

	accountID := *receipt.AccountID

	resp, err = NewTokenCreateTransaction().
		SetTokenName("ffff").
		SetTokenSymbol("F").
		SetDecimals(3).
		SetInitialSupply(1000000).
		SetTreasuryAccountID(client.GetOperatorAccountID()).
		SetAdminKey(client.GetOperatorPublicKey()).
		SetFreezeKey(client.GetOperatorPublicKey()).
		SetWipeKey(client.GetOperatorPublicKey()).
		SetKycKey(client.GetOperatorPublicKey()).
		SetSupplyKey(client.GetOperatorPublicKey()).
		SetFreezeDefault(false).
		Execute(client)
	assert.NoError(t, err)

	receipt, err = resp.GetReceipt(client)
	assert.NoError(t, err)

	tokenID := *receipt.TokenID

	_, err = NewTokenAssociateTransaction().
		SetNodeAccountIDs([]AccountID{resp.NodeID}).
		SetAccountID(accountID).
		SetTokenIDs(tokenID).
		Execute(client)
	assert.NoError(t, err)

	_, err = resp.GetReceipt(client)
	assert.NoError(t, err)

	_, err = NewTokenDissociateTransaction().
		SetNodeAccountIDs([]AccountID{resp.NodeID}).
		SetAccountID(accountID).
		Execute(client)
	assert.NoError(t, err)

	_, err = resp.GetReceipt(client)
	assert.NoError(t, err)

	resp, err = NewTokenDeleteTransaction().
		SetNodeAccountIDs([]AccountID{resp.NodeID}).
		SetTokenID(tokenID).
		Execute(client)
	assert.NoError(t, err)

	_, err = resp.GetReceipt(client)
	assert.NoError(t, err)

	_, err = NewAccountDeleteTransaction().
		SetAccountID(accountID).
		SetTransferAccountID(client.GetOperatorAccountID()).
		Execute(client)
	assert.NoError(t, err)

	_, err = resp.GetReceipt(client)
	assert.NoError(t, err)
}

func Test_TokenDissociate_NoAccountID(t *testing.T) {
	client := newTestClient(t)

	newKey, err := GeneratePrivateKey()
	assert.NoError(t, err)

	newBalance := NewHbar(1)

	assert.Equal(t, HbarUnits.Hbar.numberOfTinybar(), newBalance.tinybar)

	resp, err := NewAccountCreateTransaction().
		SetKey(newKey.PublicKey()).
		SetMaxTransactionFee(NewHbar(2)).
		SetInitialBalance(newBalance).
		Execute(client)
	assert.NoError(t, err)

	receipt, err := resp.GetReceipt(client)
	assert.NoError(t, err)

	accountID := *receipt.AccountID

	resp, err = NewTokenCreateTransaction().
		SetTokenName("ffff").
		SetTokenSymbol("F").
		SetDecimals(3).
		SetInitialSupply(1000000).
		SetTreasuryAccountID(client.GetOperatorAccountID()).
		SetAdminKey(client.GetOperatorPublicKey()).
		SetFreezeKey(client.GetOperatorPublicKey()).
		SetWipeKey(client.GetOperatorPublicKey()).
		SetKycKey(client.GetOperatorPublicKey()).
		SetSupplyKey(client.GetOperatorPublicKey()).
		SetFreezeDefault(false).
		Execute(client)
	assert.NoError(t, err)

	receipt, err = resp.GetReceipt(client)
	assert.NoError(t, err)

	tokenID := *receipt.TokenID

	_, err = NewTokenAssociateTransaction().
		SetNodeAccountIDs([]AccountID{resp.NodeID}).
		SetAccountID(accountID).
		SetTokenIDs(tokenID).
		Execute(client)
	assert.NoError(t, err)

	_, err = resp.GetReceipt(client)
	assert.NoError(t, err)

	_, err = NewTokenDissociateTransaction().
		SetNodeAccountIDs([]AccountID{resp.NodeID}).
		Execute(client)
	assert.Error(t, err)

	_, err = resp.GetReceipt(client)
	assert.NoError(t, err)

	resp, err = NewTokenDeleteTransaction().
		SetNodeAccountIDs([]AccountID{resp.NodeID}).
		SetTokenID(tokenID).
		Execute(client)
	assert.NoError(t, err)

	_, err = resp.GetReceipt(client)
	assert.NoError(t, err)

	_, err = NewAccountDeleteTransaction().
		SetAccountID(accountID).
		SetTransferAccountID(client.GetOperatorAccountID()).
		Execute(client)
	assert.NoError(t, err)

	_, err = resp.GetReceipt(client)
	assert.NoError(t, err)
}
