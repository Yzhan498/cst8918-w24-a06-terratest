package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// You normally want to run this under a separate "Testing" subscription
// For lab purposes you will use your assigned subscription under the Cloud Dev/Ops program tenant
var subscriptionID string = "45a61a7b-f92a-4559-8168-1d379738daf3"

func TestAzureLinuxVMCreation(t *testing.T) {
	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../",
		// Override the default terraform variables
		Vars: map[string]interface{}{
			"labelPrefix": "zhan0865",
		},
	}

	defer terraform.Destroy(t, terraformOptions)

	// Run `terraform init` and `terraform apply`. Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the value of output variable
	vmName := terraform.Output(t, terraformOptions, "vm_name")
	resourceGroupName := terraform.Output(t, terraformOptions, "resource_group_name")

	// Confirm VM exists
	assert.True(t, azure.VirtualMachineExists(t, vmName, resourceGroupName, subscriptionID))
}
func TestAzureLinuxVMNicExists(t *testing.T) {
    terraformOptions := &terraform.Options{
        // The path to where our Terraform code is located
        TerraformDir: "../",
        // Override the default terraform variables
        Vars: map[string]interface{}{
            "labelPrefix": "zhan0865",
        },
    }
 
    defer terraform.Destroy(t, terraformOptions)
 
    // Run `terraform init` and `terraform apply`. Fail the test if there are any errors.
    terraform.InitAndApply(t, terraformOptions)
 
    // Run `terraform output` to get the value of output variable
    vmName := terraform.Output(t, terraformOptions, "vm_name")
    nicName := terraform.Output(t, terraformOptions, "nic_name")
    resourceGroupName := terraform.Output(t, terraformOptions, "resource_group_name")
 
    // Confirm VM exists
    assert.True(t, azure.VirtualMachineExists(t, vmName, resourceGroupName, subscriptionID))
    // Check if NIC exists on the VM created
    nics := azure.GetVirtualMachineNics(t, vmName, resourceGroupName, subscriptionID)
   
    // Assert that at least one NIC exists
    assert.NotEmpty(t, nics, "At least one NIC should be attached to the VM")
 
    // Check if the NIC with the specified name exists
    var nicExists bool
    for _, nic := range nics {
        if nic == nicName {
            nicExists = true
            break
        }
    }
 
    // Assert that the NIC with the specified name exists
    assert.True(t, nicExists, "NIC with name %s should exist", nicName)
 
}

func TestUbuntuVersionOnAzureVM(t *testing.T) {
 terraformOptions := &terraform.Options{
     // The path to where our Terraform code is located
     TerraformDir: "../",
     // Override the default terraform variables
     Vars: map[string]interface{}{
         "labelPrefix": "zhan0865",
     },
 }
 
 defer terraform.Destroy(t, terraformOptions)
 
 // Run `terraform init` and `terraform apply`. Fail the test if there are any errors.
 terraform.InitAndApply(t, terraformOptions)
 
 // Run `terraform output` to get the value of output variable
 vmName := terraform.Output(t, terraformOptions, "vm_name")
 resourceGroupName := terraform.Output(t, terraformOptions, "resource_group_name")
 vmImage := azure.GetVirtualMachineImage(t, vmName, resourceGroupName, subscriptionID)
 vmImageVersionCorrect := vmImage.Version
    
 assert.True(t,  vmImageVersionCorrect == "latest")
}