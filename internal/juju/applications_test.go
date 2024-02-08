// Copyright 2024 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package juju

// Basic imports
import (
	"context"
	"testing"

	"github.com/juju/juju/api"
	"github.com/juju/juju/api/base"
	"github.com/juju/juju/core/constraints"
	"github.com/juju/juju/core/model"
	"github.com/juju/juju/rpc/params"
	"github.com/juju/names/v4"
	"github.com/juju/version/v2"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type ApplicationSuite struct {
	suite.Suite

	VariableThatShouldStartAtFive int
	mockApplicationClient         *MockApplicationAPIClient
	mockClient                    *MockClientAPIClient
	mockSharedClient              *MockSharedClient
	mockConnection                *MockConnection
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (s *ApplicationSuite) SetupTest() {
	s.VariableThatShouldStartAtFive = 5
}

func (s *ApplicationSuite) setupMocks(t *testing.T) *gomock.Controller {
	ctlr := gomock.NewController(t)
	s.mockApplicationClient = NewMockApplicationAPIClient(ctlr)
	s.mockClient = NewMockClientAPIClient(ctlr)

	s.mockConnection = NewMockConnection(ctlr)
	s.mockConnection.EXPECT().Close().Return(nil).AnyTimes()

	s.mockSharedClient = NewMockSharedClient(ctlr)
	s.mockSharedClient.EXPECT().Debugf(gomock.Any(), gomock.Any()).AnyTimes()
	s.mockSharedClient.EXPECT().Errorf(gomock.Any(), gomock.Any()).AnyTimes()
	s.mockSharedClient.EXPECT().Tracef(gomock.Any(), gomock.Any()).AnyTimes()
	s.mockSharedClient.EXPECT().JujuLogger().Return(&jujuLoggerShim{}).AnyTimes()
	s.mockSharedClient.EXPECT().GetConnection(gomock.Any()).Return(s.mockConnection, nil).AnyTimes()
	return ctlr
}

func (s *ApplicationSuite) getApplicationsClient() applicationsClient {
	return applicationsClient{
		SharedClient:      s.mockSharedClient,
		controllerVersion: version.Number{},
		getApplicationAPIClient: func(_ base.APICallCloser) ApplicationAPIClient {
			return s.mockApplicationClient
		},
		getClientAPIClient: func(_ api.Connection) ClientAPIClient {
			return s.mockClient
		},
	}
}

// All methods that begin with "Test" are run as tests within a
// suite.
func (s *ApplicationSuite) TestReadApplicationRetry() {
	defer s.setupMocks(s.T()).Finish()
	s.mockSharedClient.EXPECT().ModelType(gomock.Any()).Return(model.IAAS, nil).AnyTimes()
	appName := "testapplication"
	statusResult := &params.FullStatus{
		Applications: map[string]params.ApplicationStatus{appName: {
			Charm: "ch:amd64/jammy/testcharm-5",
			Units: map[string]params.UnitStatus{"testapplication/0": {
				Machine: "0",
			}},
		}},
	}
	s.mockClient.EXPECT().Status(gomock.Any()).Return(statusResult, nil)
	aExp := s.mockApplicationClient.EXPECT()

	// First response is not found.
	aExp.ApplicationsInfo(gomock.Any()).Return([]params.ApplicationInfoResult{{
		Error: &params.Error{Message: `application "testapplication" not found`, Code: "not found"},
	}}, nil)
	// The second time return a real application.
	amdConst, _ := constraints.Parse("arch=amd64")
	infoResult := params.ApplicationInfoResult{
		Result: &params.ApplicationResult{
			Tag:         names.NewApplicationTag(appName).String(),
			Charm:       "ch:amd64/jammy/testcharm-5",
			Base:        params.Base{Name: "ubuntu", Channel: "22.04"},
			Channel:     "stable",
			Constraints: amdConst,
			Principal:   true,
		},
		Error: nil,
	}
	aExp.ApplicationsInfo(gomock.Any()).Return([]params.ApplicationInfoResult{infoResult}, nil)
	//aExp.GetConstraints(gomock.Any()).Return()
	getResult := &params.ApplicationGetResults{
		Application:       appName,
		CharmConfig:       nil,
		ApplicationConfig: nil,
		Charm:             "ch:amd64/jammy/testcharm-5",
		Base:              params.Base{Name: "ubuntu", Channel: "22.04"},
		Channel:           "stable",
		Constraints:       amdConst,
		EndpointBindings:  nil,
	}
	aExp.Get("master", appName).Return(getResult, nil)

	client := s.getApplicationsClient()
	resp, err := client.ReadApplicationWithRetryOnNotFound(context.Background(),
		&ReadApplicationInput{
			ModelName: "testmodel",
			AppName:   appName,
		})
	s.Assert().NoError(err)

	if err != nil {
		s.FailNowf("error found", " %+v", err)
	}
	if resp == nil {
		s.FailNow("nil response")
	}
	s.Assert().NotNil(resp)
	respName := resp.Name
	s.Assert().Equal(respName, appName)
	s.Assert().Equal(resp.Channel, "stable")

	//assert.Equal(suite.T(), 5, suite.VariableThatShouldStartAtFive)
	//suite.Equal(5, suite.VariableThatShouldStartAtFive)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestApplicationSuite(t *testing.T) {
	suite.Run(t, new(ApplicationSuite))
}
