package hedera

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFileUpdateTransaction_Execute(t *testing.T) {
	client := newTestClient(t)

	client.SetMaxTransactionFee(NewHbar(2))

	resp, err := NewFileCreateTransaction().
		SetKeys(client.GetOperatorPublicKey()).
		SetContents([]byte("Hello, World")).
		SetTransactionMemo("go sdk e2e tests").
		Execute(client)

	assert.NoError(t, err)

	receipt, err := resp.GetReceipt(client)
	assert.NoError(t, err)

	fileID := *receipt.FileID
	assert.NotNil(t, fileID)

	var newContents = []byte("Good Night, World")

	resp, err = NewFileUpdateTransaction().
		SetFileID(fileID).
		SetNodeAccountIDs([]AccountID{resp.NodeID}).
		SetContents(newContents).
		Execute(client)

	assert.NoError(t, err)

	_, err = resp.GetReceipt(client)
	assert.NoError(t, err)

	contents, err := NewFileContentsQuery().
		SetFileID(fileID).
		SetNodeAccountIDs([]AccountID{resp.NodeID}).
		Execute(client)
	assert.NoError(t, err)

	assert.Equal(t, newContents, contents)

	resp, err = NewFileDeleteTransaction().
		SetFileID(fileID).
		SetNodeAccountIDs([]AccountID{resp.NodeID}).
		Execute(client)
	assert.NoError(t, err)

	_, err = resp.GetReceipt(client)
	assert.NoError(t, err)
}

func Test_FileUpdate_3(t *testing.T) {
	client := newTestClient(t)

	client.SetMaxTransactionFee(NewHbar(2))

	resp, err := NewFileCreateTransaction().
		SetKeys(client.GetOperatorPublicKey()).
		SetContents([]byte("Hello, World")).
		SetTransactionMemo("go sdk e2e tests").
		Execute(client)

	assert.NoError(t, err)

	receipt, err := resp.GetReceipt(client)
	assert.NoError(t, err)

	fileID := *receipt.FileID
	assert.NotNil(t, fileID)

	resp, err = NewFileUpdateTransaction().
		SetNodeAccountIDs([]AccountID{resp.NodeID}).
		Execute(client)
	assert.NoError(t, err)

	_, err = resp.GetReceipt(client)
	assert.Error(t, err)

	resp, err = NewFileDeleteTransaction().
		SetFileID(fileID).
		SetNodeAccountIDs([]AccountID{resp.NodeID}).
		Execute(client)
	assert.NoError(t, err)

	_, err = resp.GetReceipt(client)
	assert.NoError(t, err)
}
