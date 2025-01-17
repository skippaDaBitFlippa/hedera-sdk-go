//go:build all || unit
// +build all unit

package hedera

/*-
 *
 * Hedera Go SDK
 *
 * Copyright (C) 2020 - 2022 Hedera Hashgraph, LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// testinf func (transfer TokenNftTransfer) ToBytes() and func NftTransferFromBytes(data []byte)
func TestUnitTokenNftTransferToBytes(t *testing.T) {
	t.Parallel()

	transfer := TokenNftTransfer{
		SenderAccountID:   AccountID{Account: 3},
		ReceiverAccountID: AccountID{Account: 4},
		SerialNumber:      5,
		IsApproved:        true,
	}

	transferBytes := transfer.ToBytes()
	transferFromBytes, err := NftTransferFromBytes(transferBytes)

	assert.NoError(t, err)
	assert.Equal(t, transfer, transferFromBytes)

	// test invalid data from and to bytes
	_, err = NftTransferFromBytes([]byte{1, 2, 3})
	assert.Error(t, err)

	// test nil data from bytes and to bytes
	_, err = NftTransferFromBytes(nil)
	assert.Error(t, err)
}
