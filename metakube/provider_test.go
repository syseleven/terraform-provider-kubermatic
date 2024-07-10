package metakube

import (
	"fmt"
	"os"
	"testing"
	"text/template"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const (
	testNamePrefix = "tf-acc-test-"

	testEnvOtherUserEmail = "METAKUBE_ANOTHER_USER_EMAIL"

	testEnvK8sVersionOpenstack = "METAKUBE_K8S_VERSION_OS"
	testEnvK8sVersionAWS       = "METAKUBE_K8S_VERSION_AWS"
	testEnvK8sOlderVersion     = "METAKUBE_K8S_OLDER_VERSION"

	testEnvProjectID   = "METAKUBE_PROJECT_ID"
	testEnvProjectName = "METAKUBE_PROJECT_NAME"

	testEnvOpenstackNodeDC                       = "METAKUBE_OPENSTACK_NODE_DC"
	testEnvOpenstackApplicationCredentialsID     = "METAKUBE_OPENSTACK_APPLICATION_CREDENTIALS_ID"
	testEnvOpenstackApplicationCredentialsSecret = "METAKUBE_OPENSTACK_APPLICATION_CREDENTIALS_SECRET"
	testEnvOpenstackUsername                     = "METAKUBE_OPENSTACK_USERNAME"
	testEnvOpenstackAuthURL                      = "METAKUBE_OPENSTACK_AUTH_URL"
	testEnvOpenstackPassword                     = "METAKUBE_OPENSTACK_PASSWORD"
	testEnvOpenstackProjectID                    = "METAKUBE_OPENSTACK_PROJECT_ID"
	testEnvOpenstackProjectName                  = "METAKUBE_OPENSTACK_PROJECT_NAME"
	testEnvOpenstackRegion                       = "METAKUBE_OPENSTACK_REGION"
	testEnvOpenstackImage                        = "METAKUBE_OPENSTACK_IMAGE"
	testEnvOpenstackImage2                       = "METAKUBE_OPENSTACK_IMAGE2"
	testEnvOpenstackFlavor                       = "METAKUBE_OPENSTACK_FLAVOR"

	testEnvAzureNodeDC         = "METAKUBE_AZURE_NODE_DC"
	testEnvAzureNodeSize       = "METAKUBE_AZURE_NODE_SIZE"
	testEnvAzureClientID       = "METAKUBE_AZURE_CLIENT_ID"
	testEnvAzureClientSecret   = "METAKUBE_AZURE_CLIENT_SECRET"
	testEnvAzureTenantID       = "METAKUBE_AZURE_TENANT_ID"
	testEnvAzureSubscriptionID = "METAKUBE_AZURE_SUBSCRIPTION_ID"

	testEnvAWSAccessKeyID      = "METAKUBE_AWS_ACCESS_KEY_ID"
	testAWSSecretAccessKey     = "METAKUBE_AWS_ACCESS_KEY_SECRET"
	testEnvAWSVPCID            = "METAKUBE_AWS_VPC_ID"
	testEnvAWSNodeDC           = "METAKUBE_AWS_NODE_DC"
	testEnvAWSInstanceType     = "METAKUBE_AWS_INSTANCE_TYPE"
	testEnvAWSSubnetID         = "METAKUBE_AWS_SUBNET_ID"
	testEnvAWSAvailabilityZone = "METAKUBE_AWS_AVAILABILITY_ZONE"
	testEnvAWSDiskSize         = "METAKUBE_AWS_DISK_SIZE"
)

var (
	testAccProviders map[string]*schema.Provider
	testAccProvider  *schema.Provider
)

func init() {
	resource.AddTestSweepers("metakube_cluster", &resource.Sweeper{
		Name: "metakube_cluster",
		F:    testSweepClusters,
	})
	resource.AddTestSweepers("metakube_sshkey", &resource.Sweeper{
		Name: "metakube_sshkey",
		F:    testSweepSSHKeys,
	})
}
func TestMain(m *testing.M) {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"metakube": testAccProvider,
	}
	resource.TestMain(m)
}

func testAccPreCheckForOpenstack(t *testing.T) {
	t.Helper()
	testAccPreCheck(t)
	checkEnv(t, testEnvK8sVersionOpenstack)
	checkEnv(t, testEnvOpenstackApplicationCredentialsID)
	checkEnv(t, testEnvOpenstackApplicationCredentialsSecret)
	checkEnv(t, testEnvOpenstackProjectID)
	checkEnv(t, testEnvOpenstackProjectName)
	checkEnv(t, testEnvOpenstackRegion)
	checkEnv(t, testEnvOpenstackNodeDC)
	checkEnv(t, testEnvOpenstackImage)
	checkEnv(t, testEnvOpenstackImage2)
	checkEnv(t, testEnvOpenstackFlavor)
	checkEnv(t, testEnvOpenstackAuthURL)
	checkEnv(t, testEnvK8sOlderVersion)
	checkEnv(t, testEnvProjectID)
}

func testAccPreCheckForAWS(t *testing.T) {
	t.Helper()
	testAccPreCheck(t)
	checkEnv(t, testEnvAWSAccessKeyID)
	checkEnv(t, testEnvK8sVersionAWS)
	checkEnv(t, testAWSSecretAccessKey)
	checkEnv(t, testEnvAWSVPCID)
	checkEnv(t, testEnvAWSNodeDC)
	checkEnv(t, testEnvK8sOlderVersion)
	checkEnv(t, testEnvProjectID)
}

func testAccPreCheck(t *testing.T) {
	t.Helper()
	checkEnv(t, "METAKUBE_HOST")
	checkEnv(t, "METAKUBE_TOKEN")
}

func checkEnv(t *testing.T, n string) {
	t.Helper()
	if v := os.Getenv(n); v == "" {
		t.Fatalf("%s must be set for acceptance tests", n)
	}
}

func makeRandomName() string {
	return randomName(testNamePrefix, 5)
}

func randomName(prefix string, length int) string {
	return fmt.Sprintf("%s%s", prefix, acctest.RandString(length))
}

func testResourceInstanceState(name string, check func(*terraform.InstanceState) error) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		m := s.RootModule()
		if rs, ok := m.Resources[name]; ok {
			is := rs.Primary
			if is == nil {
				return fmt.Errorf("no instance: %s", name)
			}

			return check(is)
		}
		return fmt.Errorf("not found: %s", name)

	}
}

func mustParseTemplate(name, text string) *template.Template {
	r, err := template.New(name).Parse(text)
	if err != nil {
		panic(err)
	}
	return r
}
