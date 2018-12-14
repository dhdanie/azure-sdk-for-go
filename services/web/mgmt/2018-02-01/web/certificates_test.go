package web

import (
	"context"
	"fmt"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"testing"
)

const CertsSubscriptionId = "d48c5f68-09d3-45f2-b401-694f33f1472b"

func TestCertificatesClient_List(t *testing.T) {
	testList(t)
}

func testList(t *testing.T) {
	authorizer, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		t.Errorf("Error creating authorizer - %s", err.Error())
		return
	}

	client := NewCertificatesClient(CertsSubscriptionId)
	client.Authorizer = authorizer

	certs, err := client.List(context.Background())
	if err != nil {
		t.Errorf("Error retrieving certificates list - %s", err.Error())
		return
	}

	for certs.NotDone() {
		for _, cert := range certs.Values() {
			fmt.Printf("%s - %s\n", *cert.ID, *cert.Name)
		}
		certs.NextWithContext(context.Background())
	}
}
