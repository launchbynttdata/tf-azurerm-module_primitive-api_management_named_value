package testimpl

import (
	"context"
	"os"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	apiManagement "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/launchbynttdata/lcaf-component-terratest/types"
	"github.com/stretchr/testify/assert"
)

func TestApiManagementModule(t *testing.T, ctx types.TestContext) {
	testApiManagementModule(t, ctx)
}

func TestComposableApiManagementModule(t *testing.T, ctx types.TestContext) {
	testApiManagementModule(t, ctx)
}

func testApiManagementModule(t *testing.T, ctx types.TestContext) {
	subscriptionId := os.Getenv("ARM_SUBSCRIPTION_ID")
	if len(subscriptionId) == 0 {
		t.Fatal("ARM_SUBSCRIPTION_ID environment variable is not set")
	}

	credential, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		t.Fatalf("Unable to get credentials: %e\n", err)
	}

	namedValueClient, err := apiManagement.NewNamedValueClient(subscriptionId, credential, &arm.ClientOptions{
		ClientOptions: azcore.ClientOptions{
			Cloud: cloud.AzurePublic,
		},
	})
	if err != nil {
		t.Fatalf("Error getting API Management named value client: %v", err)
	}

	resourceGroupName := terraform.Output(t, ctx.TerratestTerraformOptions(), "resource_group_name")
	serviceName := terraform.Output(t, ctx.TerratestTerraformOptions(), "api_management_name")
	namedValueName := terraform.Output(t, ctx.TerratestTerraformOptions(), "named_value_name")
	secretNamedValueName := terraform.Output(t, ctx.TerratestTerraformOptions(), "secret_named_value_name")

	t.Run("doesApiManagementNamedValueExist", func(t *testing.T) {
		namedValue, err := namedValueClient.Get(context.Background(), resourceGroupName, serviceName, namedValueName, nil)
		if err != nil {
			t.Fatalf("Error getting API Management named value: %v", err)
		}
		assert.Equal(t, namedValueName, *namedValue.Name)

		secretNamedValue, err := namedValueClient.Get(context.Background(), resourceGroupName, serviceName, secretNamedValueName, nil)
		if err != nil {
			t.Fatalf("Error getting API Management secret named value: %v", err)
		}
		assert.Equal(t, secretNamedValueName, *secretNamedValue.Name)
	})
}
