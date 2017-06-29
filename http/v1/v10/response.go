package v10

import (
	"github.com/amsokol/go-ignite-client/http/types"
)

// response is common for all methods response header
type response struct {
	SuccessStatus types.SuccessStatus `json:"successStatus"`
	Error         string              `json:"error"`
	SessionToken  types.SessionToken  `json:"sessionToken"`
}
