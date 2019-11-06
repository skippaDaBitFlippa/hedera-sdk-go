package hedera

import "google.golang.org/grpc"

const defaultMaxTransactionFee uint64 = 100_000_000 // 1 Hbar

type Client struct {
	// todo: support multiple nodes
	nodeID AccountID
	maxTransactionFee uint64
	maxQueryPayment uint64
	operator *operator
	conn *grpc.ClientConn
}

type operator struct {
	accountID AccountID
	privateKey PrivateKey
}

func NewClient(nodeId AccountID, address string) (*Client, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := Client {
		nodeId,
		defaultMaxTransactionFee,
		0,
		nil,
		conn,
	}

	return &client, nil
}

func (client *Client) Close() error {
	return client.conn.Close()
}

func (client *Client) SetOperator(accountId AccountID, privateKey PrivateKey) *Client {
	operator := operator {
		accountId,
		privateKey,
	}

	client.operator = &operator

	return client
}

func (client *Client) SetMaxTransactionFee(tinyBars uint64) *Client {
	client.maxTransactionFee = tinyBars
	return client
}

func (client *Client) SetMaxQueryPayment(tinyBars uint64) *Client {
	client.maxTransactionFee = tinyBars
	return client
}

func (client *Client) OperatorId() *AccountID {
	if client.operator == nil {
		return nil
	}

	return &client.operator.accountID
}

func (client *Client) OperatorPrivateKey() *PrivateKey {
	if client.operator == nil {
		return nil
	}

	return &client.operator.privateKey
}

func (client *Client) MaxTransactionFee() uint64 {
	return client.maxTransactionFee
}

func (client *Client) MaxQueryPayment() uint64 {
	return client.maxQueryPayment
}
