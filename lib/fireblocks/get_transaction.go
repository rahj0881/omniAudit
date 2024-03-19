package fireblocks

import (
	"context"
	"encoding/json"
	"net/http"
	"path/filepath"

	"github.com/omni-network/omni/lib/errors"
)

func (c Client) GetTransactionByID(ctx context.Context, transactionID string) (*TransactionResponse, error) {
	var res TransactionResponse

	endpoint := filepath.Join(transactionEndpoint, transactionID)
	jwtToken, err := c.GenJWTToken(endpoint, nil)
	if err != nil {
		return nil, err
	}

	response, err := c.http.SendRequest(
		ctx,
		endpoint,
		http.MethodGet,
		nil,
		c.getHeaders(jwtToken),
	)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(response), &res)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshaling response")
	}

	return &res, nil
}
