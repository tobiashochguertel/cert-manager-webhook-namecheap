package main

import (
	"os"
	"testing"

	// "github.com/jetstack/cert-manager/test/acme/dns"
	acmetest "github.com/cert-manager/cert-manager/test/acme"
	// "github.com/cert-manager/webhook-example/example"
)

// go get github.com/cert-manager/cert-manager/test/acme@v1.12.6
// go get github.com/cert-manager/webhook-example/example

var (
	zone = os.Getenv("TEST_ZONE_NAME")
)

func TestRunsSuite(t *testing.T) {
	// The manifest path should contain a file named config.json that is a
	// snippet of valid configuration that should be included on the
	// ChallengeRequest passed as part of the test cases.
	//

	// Uncomment the below fixture when implementing your custom DNS provider
	fixture := acmetest.NewFixture(&namecheapDNSProviderSolver{},
		acmetest.SetResolvedZone(zone),
		acmetest.SetAllowAmbientCredentials(false),
		acmetest.SetManifestPath("testdata/namecheap"),
		// acmetest.SetBinariesPath("_test/kubebuilder/bin"),
	)
	/* 	solver := example.New("59351")
	   	fixture := acmetest.NewFixture(solver,
	   		acmetest.SetResolvedZone("example.com."),
	   		acmetest.SetManifestPath("testdata/my-custom-solver"),
	   		acmetest.SetDNSServer("127.0.0.1:59351"),
	   		acmetest.SetUseAuthoritative(false),
	   	) */
	//need to uncomment and  RunConformance delete runBasic and runExtended once https://github.com/cert-manager/cert-manager/pull/4835 is merged
	fixture.RunConformance(t)
	/* 	fixture.RunBasic(t)
	   	fixture.RunExtended(t) */

	/*
		 	fixture := dns.NewFixture(&namecheapDNSProviderSolver{},
				dns.SetResolvedZone(zone),
				dns.SetAllowAmbientCredentials(false),
				dns.SetManifestPath("testdata/namecheap"),
				dns.SetBinariesPath("_test/kubebuilder/bin"),
			)

			fixture.RunConformance(t)
	*/
}
