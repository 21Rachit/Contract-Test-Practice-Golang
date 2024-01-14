package server

import (
	"fmt"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
)

func TestServerPact_Verification(t *testing.T) {
	// initialize PACT DSL
	pact := dsl.Pact{
		Provider: "example-server",
	}

	// verify Contract on server side
	_, err := pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL: "http://localhost:9001",
		PactURLs:        []string{"../client/pacts/example-client-example-server.json"},
	})

	if err != nil {
		fmt.Println("error in server test is ", err)
	}
}
