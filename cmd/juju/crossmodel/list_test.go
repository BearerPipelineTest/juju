// Copyright 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package crossmodel_test

import (
	"fmt"
	"strings"

	"github.com/juju/cmd"
	"github.com/juju/errors"
	jc "github.com/juju/testing/checkers"
	gc "gopkg.in/check.v1"
	"gopkg.in/juju/charm.v6-unstable"

	"github.com/juju/juju/cmd/juju/crossmodel"
	model "github.com/juju/juju/model/crossmodel"
	"github.com/juju/juju/testing"
)

type ListSuite struct {
	BaseCrossModelSuite

	mockAPI *mockListAPI

	services  []model.OfferedServiceDetailsResult
	endpoints []charm.Relation
}

var _ = gc.Suite(&ListSuite{})

func (s *ListSuite) SetUpTest(c *gc.C) {
	s.BaseCrossModelSuite.SetUpTest(c)

	s.endpoints = []charm.Relation{
		{Name: "mysql", Interface: "db2", Role: charm.RoleRequirer},
		{Name: "log", Interface: "http", Role: charm.RoleProvider},
	}

	s.services = []model.OfferedServiceDetailsResult{
		{Result: s.createServiceItem("hosted-db2", "local", 0)},
	}

	s.mockAPI = &mockListAPI{
		list: func(filters ...model.OfferedServiceFilter) ([]model.OfferedServiceDetailsResult, error) {
			return s.services, nil
		},
	}
}

func (s *ListSuite) TestListError(c *gc.C) {
	msg := "fail api"

	s.mockAPI.list = func(filters ...model.OfferedServiceFilter) ([]model.OfferedServiceDetailsResult, error) {
		return nil, errors.New(msg)
	}

	_, err := s.runList(c, nil)
	c.Assert(err, gc.ErrorMatches, fmt.Sprintf(".*%v.*", msg))
}

func (s *ListSuite) TestListFormatError(c *gc.C) {
	s.services = append(s.services, model.OfferedServiceDetailsResult{Result: s.createServiceItem("zdi^%", "different_store", 33)})

	_, err := s.runList(c, nil)
	c.Assert(err, gc.ErrorMatches, ".*failed to format.*")
}

func (s *ListSuite) TestListDirectories(c *gc.C) {
	// Insert in random order to check sorting.
	s.services = append(s.services, model.OfferedServiceDetailsResult{Result: s.createServiceItem("zdiff-db2", "differentstore", 33)})
	s.services = append(s.services, model.OfferedServiceDetailsResult{Result: s.createServiceItem("adiff-db2", "vendor", 23)})

	s.assertValidList(
		c,
		nil,
		// Default format is tabular
		`
differentstore
SERVICE    CHARM  CONNECTED  STORE           URL                ENDPOINT  INTERFACE  ROLE
zdiff-db2  db2    33         differentstore  /u/fred/zdiff-db2  log       http       provider
                                                                mysql     db2        requirer
local
SERVICE     CHARM  CONNECTED  STORE  URL                 ENDPOINT  INTERFACE  ROLE
hosted-db2  db2    0          local  /u/fred/hosted-db2  log       http       provider
                                                         mysql     db2        requirer
vendor
SERVICE    CHARM  CONNECTED  STORE   URL                ENDPOINT  INTERFACE  ROLE
adiff-db2  db2    23         vendor  /u/fred/adiff-db2  log       http       provider
                                                        mysql     db2        requirer

`[1:],
		"",
	)
}

func (s *ListSuite) TestListWithErrors(c *gc.C) {
	msg := "here is the error"
	s.services = append(s.services, model.OfferedServiceDetailsResult{Error: errors.New(msg)})

	s.assertValidList(
		c,
		nil,
		`
local
SERVICE     CHARM  CONNECTED  STORE  URL                 ENDPOINT  INTERFACE  ROLE
hosted-db2  db2    0          local  /u/fred/hosted-db2  log       http       provider
                                                         mysql     db2        requirer

`[1:],
		msg,
	)
}

func (s *ListSuite) TestListYAML(c *gc.C) {
	// Since services are in the map and ordering is unreliable, ensure that there is only one endpoint.
	// We only need one to demonstrate display anyway :D
	s.services[0].Result.Endpoints = []charm.Relation{{Name: "mysql", Interface: "db2", Role: charm.RoleRequirer}}

	s.assertValidList(
		c,
		[]string{"--format", "yaml"},
		`
local:
  hosted-db2:
    charm: db2
    connected: 0
    store: local
    url: /u/fred/hosted-db2
    endpoints:
      mysql:
        interface: db2
        role: requirer
`[1:],
		"",
	)
}

func (s *ListSuite) createServiceItem(name, store string, count int) *model.OfferedServiceDetails {
	return &model.OfferedServiceDetails{
		ServiceName:    name,
		ServiceURL:     fmt.Sprintf("%s:%s%s", store, "/u/fred/", name),
		CharmName:      "db2",
		Endpoints:      s.endpoints,
		ConnectedCount: count,
	}
}

func (s *ListSuite) runList(c *gc.C, args []string) (*cmd.Context, error) {
	return testing.RunCommand(c, crossmodel.NewListEndpointsCommandForTest(s.mockAPI), args...)
}

func (s *ListSuite) assertValidList(c *gc.C, args []string, expectedValid, expectedErr string) {
	context, err := s.runList(c, args)
	c.Assert(err, jc.ErrorIsNil)

	obtainedErr := strings.Replace(testing.Stderr(context), "\n", "", -1)
	c.Assert(obtainedErr, gc.Matches, expectedErr)

	obtainedValid := testing.Stdout(context)
	c.Assert(obtainedValid, gc.Matches, expectedValid)
}

type mockListAPI struct {
	list func(filters ...model.OfferedServiceFilter) ([]model.OfferedServiceDetailsResult, error)
}

func (s mockListAPI) Close() error {
	return nil
}

func (s mockListAPI) ListOffers(filters ...model.OfferedServiceFilter) ([]model.OfferedServiceDetailsResult, error) {
	return s.list(filters...)
}
