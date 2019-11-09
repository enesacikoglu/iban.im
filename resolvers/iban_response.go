package resolvers

import (
	"strconv"

	graphql "github.com/graph-gophers/graphql-go"

	"github.com/monocash/iban.im/model"
)

// IbanResponse is the user response type
type IbanResponse struct {
	i *model.Iban
}

// ID for IbanResponse
func (r *IbanResponse) ID() graphql.ID {
	id := strconv.Itoa(int(r.i.ID))
	return graphql.ID(id)
}

// Text for IbanResponse
func (r *IbanResponse) Text() string {
	return r.i.Text
}

// Password for IbanResponse
func (r *IbanResponse) Password() string {
	return r.i.Password
}

// Handle for IbanResponse
func (r *IbanResponse) Handle() string {
	return r.i.Handle
}

// CreatedAt for IbanResponse
func (r *IbanResponse) CreatedAt() string {
	return r.i.CreatedAt.String()
}

// UpdatedAt for IbanResponse
func (r *IbanResponse) UpdatedAt() string {
	return r.i.UpdatedAt.String()
}
