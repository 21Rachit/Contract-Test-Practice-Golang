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

	// Verify the Provider with local Pact Files
	_, err := pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL:            "http://localhost:9001",
		BrokerURL:                  "https://testhpe.pactflow.io/",
		BrokerToken:                "<pact-token>",
		ProviderVersion:            "1.0.0",
		PactURLs:                   []string{"https://testhpe.pactflow.io/pacts/provider/example-server/consumer/example-client/version/1.0.0"},
		PublishVerificationResults: true,
	})

	if err != nil {
		fmt.Println("error in server test is ", err)
	}
}
