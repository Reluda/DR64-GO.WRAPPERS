Welcome to the GO Wrapper!
This is a simple wrapper made to wrap your PYINSTALL FILES into smaller and Azure Registered Apps

1. PUT ICON IN:
rsrc -ico DR64.ico //Change to your own Icon (Or keep mine!)

2. ASSIGN ROLE (AZURE DETAILS REQUIRED)
az role assignment create --assignee YourAzureAppClientID --role Reader --scope /subscriptions/YourAzureSubscriptionID

3. SET VARIABLES
$env:AZURE_TENANT_ID="YourAzureTenantID"
$env:AZURE_CLIENT_ID="YourAzureAppClientID"
$env:AZURE_CLIENT_SECRET="YourAzureAppSecretValue"

4. BUILD
go build -ldflags="-H windowsgui" -o YourAppName.exe

