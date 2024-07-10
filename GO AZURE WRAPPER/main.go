package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/subscription/armsubscription"
)

const subscriptionID = "YourAzureSubscriptionID"

func main() {
	// Set environment variables for Azure authentication
	os.Setenv("AZURE_TENANT_ID", "YourAzureTenantID")
	os.Setenv("AZURE_CLIENT_ID", "YourAzureAppClientID")
	os.Setenv("AZURE_CLIENT_SECRET", "YourAzureAppSecretValue")

	// Authenticate with Azure
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}

	client, err := armsubscription.NewSubscriptionsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create subscription client: %v", err)
	}

	_, err = client.Get(context.TODO(), subscriptionID, nil)
	if err != nil {
		log.Fatalf("failed to get subscription: %v", err)
	}

	// Path to your executable
	cmd := exec.Command("PathToYourExecuatable")

	// Suppress the command window
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	// Execute the command and get the output
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Print the output
	fmt.Println(string(out))
}
