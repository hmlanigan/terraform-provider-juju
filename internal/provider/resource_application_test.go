// Copyright 2023 Canonical Ltd.
// Licensed under the Apache License, Version 2.0, see LICENCE file for details.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	internaltesting "github.com/juju/terraform-provider-juju/internal/testing"
)

func TestAcc_ResourceApplication(t *testing.T) {
	modelName := acctest.RandomWithPrefix("tf-test-application")
	appName := "test-app"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: frameworkProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceApplicationBasic(modelName, appName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("juju_application.this", "model", modelName),
					resource.TestCheckResourceAttr("juju_application.this", "name", appName),
					resource.TestCheckResourceAttr("juju_application.this", "charm.#", "1"),
					resource.TestCheckResourceAttr("juju_application.this", "charm.0.name", "jameinel-ubuntu-lite"),
					resource.TestCheckResourceAttr("juju_application.this", "trust", "true"),
					resource.TestCheckResourceAttr("juju_application.this", "expose.#", "1"),
				),
			},
			{
				// cores constraint is not valid in K8s
				SkipFunc: func() (bool, error) {
					return testingCloud != LXDCloudTesting, nil
				},
				Config: testAccResourceApplicationConstraints(modelName, "arch=amd64 cores=1 mem=4096M"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("juju_application.this", "model", modelName),
					resource.TestCheckResourceAttr("juju_application.this", "constraints", "arch=amd64 cores=1 mem=4096M"),
				),
			},
			{
				// specific constraints for k8s
				SkipFunc: func() (bool, error) {
					// Skipping this test due to a juju bug in 2.9:
					// Unable to create application, got error: charm
					// "state changing too quickly; try again soon"
					//
					return true, nil
					//return testingCloud != MicroK8sTesting, nil
				},
				Config: testAccResourceApplicationConstraints(modelName, "arch=amd64 mem=4096M"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("juju_application.this", "model", modelName),
					resource.TestCheckResourceAttr("juju_application.this", "constraints", "arch=amd64 mem=4096M"),
				),
			},
			{
				SkipFunc: func() (bool, error) {
					return testingCloud != LXDCloudTesting, nil
				},
				Config: testAccResourceApplicationConstraintsSubordinate(modelName, "arch=amd64 cores=1 mem=4096M"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("juju_application.this", "model", modelName),
					resource.TestCheckResourceAttr("juju_application.this", "constraints", "arch=amd64 cores=1 mem=4096M"),
				),
			},
			{
				ImportStateVerify: true,
				ImportState:       true,
				ResourceName:      "juju_application.this",
			},
		},
	})
}

func TestAcc_ResourceApplication_Updates(t *testing.T) {
	modelName := acctest.RandomWithPrefix("tf-test-application")
	appName := "jameinel-ubuntu-lite"
	if testingCloud != LXDCloudTesting {
		appName = "hello-kubecon"
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: frameworkProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceApplicationUpdates(modelName, 1, true, "machinename"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("juju_application.this", "model", modelName),
					resource.TestCheckResourceAttr("juju_application.this", "charm.#", "1"),
					resource.TestCheckResourceAttr("juju_application.this", "charm.0.name", appName),
					resource.TestCheckResourceAttr("juju_application.this", "units", "1"),
					resource.TestCheckResourceAttr("juju_application.this", "expose.#", "1"),
					// (juanmanuel-tirado) Uncomment and test when running
					// a different charm with other config
					//resource.TestCheckResourceAttr("juju_application.this", "config.hostname", "machinename"),
				),
			},
			{
				SkipFunc: func() (bool, error) {
					return testingCloud != LXDCloudTesting, nil
				},
				Config: testAccResourceApplicationUpdates(modelName, 2, true, "machinename"),
				Check:  resource.TestCheckResourceAttr("juju_application.this", "units", "2"),
			},
			{
				SkipFunc: func() (bool, error) {
					return testingCloud != LXDCloudTesting, nil
				},
				Config: testAccResourceApplicationUpdates(modelName, 2, true, "machinename"),
				Check:  resource.TestCheckResourceAttr("juju_application.this", "charm.0.revision", "10"),
			},
			{
				SkipFunc: func() (bool, error) {
					return testingCloud != MicroK8sTesting, nil
				},
				Config: testAccResourceApplicationUpdates(modelName, 2, true, "machinename"),
				Check:  resource.TestCheckResourceAttr("juju_application.this", "charm.0.revision", "19"),
			},
			{
				Config: testAccResourceApplicationUpdates(modelName, 2, false, "machinename"),
				Check:  resource.TestCheckResourceAttr("juju_application.this", "expose.#", "0"),
			},
			{
				Config: testAccResourceApplicationUpdates(modelName, 2, true, "machinename"),
				Check:  resource.TestCheckResourceAttr("juju_application.this", "expose.#", "1"),
			},
			{
				ImportStateVerify: true,
				ImportState:       true,
				ResourceName:      "juju_application.this",
			},
		},
	})
}

// TestAcc_ResourceApplication_UpdatesRevisionConfig will test the revision update that have new config parameters on
// the charm. The test will check that the config is updated and the revision is updated as well.
func TestAcc_ResourceApplication_UpdatesRevisionConfig(t *testing.T) {
	if testingCloud != LXDCloudTesting {
		t.Skip(t.Name() + " only runs with LXD")
	}
	modelName := acctest.RandomWithPrefix("tf-test-application")
	appName := "github-runner"
	configParamName := "runner-storage"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: frameworkProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceApplicationWithRevisionAndConfig(modelName, appName, 88, ""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("juju_application."+appName, "model", modelName),
					resource.TestCheckResourceAttr("juju_application."+appName, "charm.#", "1"),
					resource.TestCheckResourceAttr("juju_application."+appName, "charm.0.name", appName),
					resource.TestCheckResourceAttr("juju_application."+appName, "charm.0.revision", "88"),
				),
			},
			{
				Config: testAccResourceApplicationWithRevisionAndConfig(modelName, appName, 95, configParamName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("juju_application."+appName, "charm.0.revision", "95"),
					resource.TestCheckResourceAttr("juju_application."+appName, "config."+configParamName, configParamName+"-value"),
				),
			},
		},
	})
}

func TestAcc_CharmUpdates(t *testing.T) {
	modelName := acctest.RandomWithPrefix("tf-test-charmupdates")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: frameworkProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceApplicationUpdatesCharm(modelName, "latest/stable"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("juju_application.this", "charm.0.channel", "latest/stable"),
				),
			},
			{
				// move to latest/edge
				Config: testAccResourceApplicationUpdatesCharm(modelName, "latest/edge"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("juju_application.this", "charm.0.channel", "latest/edge"),
				),
			},
			{
				// move back to latest/stable
				Config: testAccResourceApplicationUpdatesCharm(modelName, "latest/stable"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("juju_application.this", "charm.0.channel", "latest/stable"),
				),
			},
		},
	})
}

func TestAcc_ResourceApplication_Minimal(t *testing.T) {
	modelName := acctest.RandomWithPrefix("tf-test-application")
	var charmName string
	if testingCloud == LXDCloudTesting {
		charmName = "juju-qa-test"
	} else {
		charmName = "hello-juju"
	}
	resourceName := "juju_application.testapp"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: frameworkProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceApplicationBasic_Minimal(modelName, charmName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "model", modelName),
					resource.TestCheckResourceAttr(resourceName, "name", charmName),
					resource.TestCheckResourceAttr(resourceName, "charm.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "charm.0.name", charmName),
				),
			},
			// TODO: Fix the provider so this tests succeeds.
			// Temporarily not testing, - Fixing via JUJU-5472
			//{
			//ImportStateVerify: true,
			//ImportState:       true,
			//ResourceName:      resourceName,
			//},
		},
	})
}

func TestAcc_ResourceApplication_UpgradeProvider(t *testing.T) {
	modelName := acctest.RandomWithPrefix("tf-test-application")
	appName := "test-app"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },

		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"juju": {
						VersionConstraint: TestProviderStableVersion,
						Source:            "juju/juju",
					},
				},
				Config: testAccResourceApplicationBasic(modelName, appName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("juju_application.this", "model", modelName),
					resource.TestCheckResourceAttr("juju_application.this", "name", appName),
					resource.TestCheckResourceAttr("juju_application.this", "charm.#", "1"),
					resource.TestCheckResourceAttr("juju_application.this", "charm.0.name", "jameinel-ubuntu-lite"),
					resource.TestCheckResourceAttr("juju_application.this", "trust", "true"),
					resource.TestCheckResourceAttr("juju_application.this", "expose.#", "1"),
				),
			},
			{
				ProtoV6ProviderFactories: frameworkProviderFactories,
				Config:                   testAccResourceApplicationBasic(modelName, appName),
				PlanOnly:                 true,
			},
		},
	})
}

func testAccResourceApplicationBasic_Minimal(modelName, charmName string) string {
	return fmt.Sprintf(`
		resource "juju_model" "testmodel" {
		  name = %q
		}
		
		resource "juju_application" "testapp" {
		  model = juju_model.testmodel.name
		  charm {
			name = %q
		  }
		}
		`, modelName, charmName)
}

func testAccResourceApplicationBasic(modelName, appName string) string {
	if testingCloud == LXDCloudTesting {
		return fmt.Sprintf(`
		resource "juju_model" "this" {
		  name = %q
		}
		
		resource "juju_application" "this" {
		  model = juju_model.this.name
		  name = %q
		  charm {
			name = "jameinel-ubuntu-lite"
		  }
		  trust = true
		  expose{}
		}
		`, modelName, appName)
	} else {
		// if we have a K8s deployment we need the machine hostname
		return fmt.Sprintf(`
		resource "juju_model" "this" {
		  name = %q
		}
		
		resource "juju_application" "this" {
		  model = juju_model.this.name
		  name = %q
		  charm {
			name = "jameinel-ubuntu-lite"
		  }
		  trust = true
		  expose{}
		  config = {
			juju-external-hostname="myhostname"
		  }
		}
		`, modelName, appName)
	}
}

func testAccResourceApplicationWithRevisionAndConfig(modelName, appName string, revision int, configParamName string) string {
	return internaltesting.GetStringFromTemplateWithData(
		"testAccResourceApplicationWithRevisionAndConfig",
		`
resource "juju_model" "{{.ModelName}}" {
  name = "{{.ModelName}}"
}

resource "juju_application" "{{.AppName}}" {
  name  = "{{.AppName}}"
  model = juju_model.{{.ModelName}}.name

  charm {
    name     = "{{.AppName}}"
    revision = {{.Revision}}
    channel  = "latest/edge"
  }

  {{ if ne .ConfigParamName "" }}
  config = {
    {{.ConfigParamName}} = "{{.ConfigParamName}}-value"
  }
  {{ end }}

  units = 1
}
`, internaltesting.TemplateData{
			"ModelName":       modelName,
			"AppName":         appName,
			"Revision":        fmt.Sprintf("%d", revision),
			"ConfigParamName": configParamName,
		})
}

func testAccResourceApplicationUpdates(modelName string, units int, expose bool, hostname string) string {
	exposeStr := "expose{}"
	if !expose {
		exposeStr = ""
	}

	if testingCloud == LXDCloudTesting {
		return fmt.Sprintf(`
		resource "juju_model" "this" {
		  name = %q
		}
		
		resource "juju_application" "this" {
		  model = juju_model.this.name
		  units = %d
		  name = "test-app"
		  charm {
			name     = "jameinel-ubuntu-lite"
		  }
		  trust = true
		  %s
		  # config = {
		  #	 hostname = "%s"
		  # }
		}
		`, modelName, units, exposeStr, hostname)
	} else {
		return fmt.Sprintf(`
		resource "juju_model" "this" {
		  name = %q
		}
		
		resource "juju_application" "this" {
		  model = juju_model.this.name
		  units = %d
		  name = "test-app"
		  charm {
			name     = "hello-kubecon"
		  }
		  trust = true
		  %s
		  config = {
		  	# hostname = "%s"
			juju-external-hostname="myhostname"
		  }
		}
		`, modelName, units, exposeStr, hostname)
	}
}

func testAccResourceApplicationUpdatesCharm(modelName string, channel string) string {
	if testingCloud == LXDCloudTesting {
		return fmt.Sprintf(`
		resource "juju_model" "this" {
		  name = %q
		}
		
		resource "juju_application" "this" {
		  model = juju_model.this.name
		  name = "test-app"
		  charm {
			name     = "ubuntu"
			channel = %q
		  }
		}
		`, modelName, channel)
	} else {
		return fmt.Sprintf(`
		resource "juju_model" "this" {
		  name = %q
		}
		
		resource "juju_application" "this" {
		  model = juju_model.this.name
		  name = "test-app"
		  charm {
			name     = "hello-kubecon"
			channel = %q
		  }
		}
		`, modelName, channel)
	}
}

// testAccResourceApplicationConstraints will return two set for constraint
// applications. The version to be used in K8s sets the juju-external-hostname
// because we set the expose parameter.
func testAccResourceApplicationConstraints(modelName string, constraints string) string {
	if testingCloud == LXDCloudTesting {
		return fmt.Sprintf(`
resource "juju_model" "this" {
  name = %q
}

resource "juju_application" "this" {
  model = juju_model.this.name
  units = 0
  name = "test-app"
  charm {
    name     = "jameinel-ubuntu-lite"
    revision = 10
  }
  
  trust = true 
  expose{}
  constraints = "%s"
}
`, modelName, constraints)
	} else {
		return fmt.Sprintf(`
resource "juju_model" "this" {
  name = %q
}

resource "juju_application" "this" {
  model = juju_model.this.name
  name = "test-app"
  charm {
    name     = "jameinel-ubuntu-lite"
	revision = 10
  }
  trust = true
  expose{}
  constraints = "%s"
  config = {
    juju-external-hostname="myhostname"
  }
}
`, modelName, constraints)
	}
}

func testAccResourceApplicationConstraintsSubordinate(modelName string, constraints string) string {
	return fmt.Sprintf(`
resource "juju_model" "this" {
  name = %q
}

resource "juju_application" "this" {
  model = juju_model.this.name
  units = 0
  name = "test-app"
  charm {
    name     = "jameinel-ubuntu-lite"
    revision = 10
  }
  trust = true
  expose{}
  constraints = "%s"
}

resource "juju_application" "subordinate" {
  model = juju_model.this.name
  units = 0
  name = "test-subordinate"
  charm {
    name = "nrpe"
    revision = 96
    }
} 
`, modelName, constraints)
}
