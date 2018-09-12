// Code generated by gen_tests.py and process_polyglot.py.
// Do not edit this file directly.
// The template for this file is located at:
// ../template.go.tpl
package reql_tests

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	r "gopkg.in/rethinkdb/rethinkdb-go.v4"
	"gopkg.in/rethinkdb/rethinkdb-go.v4/internal/compare"
)

// sindex nulls in strings
func TestSindexNullsinstringsSuite(t *testing.T) {
	suite.Run(t, new(SindexNullsinstringsSuite))
}

type SindexNullsinstringsSuite struct {
	suite.Suite

	session *r.Session
}

func (suite *SindexNullsinstringsSuite) SetupTest() {
	suite.T().Log("Setting up SindexNullsinstringsSuite")
	// Use imports to prevent errors
	_ = time.Time{}
	_ = compare.AnythingIsFine

	session, err := r.Connect(r.ConnectOpts{
		Address: url,
	})
	suite.Require().NoError(err, "Error returned when connecting to server")
	suite.session = session

	r.DBDrop("db_sindex_null").Exec(suite.session)
	err = r.DBCreate("db_sindex_null").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("db_sindex_null").Wait().Exec(suite.session)
	suite.Require().NoError(err)

	r.DB("db_sindex_null").TableDrop("table_test_sindex_nullstr").Exec(suite.session)
	err = r.DB("db_sindex_null").TableCreate("table_test_sindex_nullstr").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("db_sindex_null").Table("table_test_sindex_nullstr").Wait().Exec(suite.session)
	suite.Require().NoError(err)
}

func (suite *SindexNullsinstringsSuite) TearDownSuite() {
	suite.T().Log("Tearing down SindexNullsinstringsSuite")

	if suite.session != nil {
		r.DB("rethinkdb").Table("_debug_scratch").Delete().Exec(suite.session)
		r.DB("db_sindex_null").TableDrop("table_test_sindex_nullstr").Exec(suite.session)
		r.DBDrop("db_sindex_null").Exec(suite.session)

		suite.session.Close()
	}
}

func (suite *SindexNullsinstringsSuite) TestCases() {
	suite.T().Log("Running SindexNullsinstringsSuite: sindex nulls in strings")

	table_test_sindex_nullstr := r.DB("db_sindex_null").Table("table_test_sindex_nullstr")
	_ = table_test_sindex_nullstr // Prevent any noused variable errors

	{
		// sindex/nullsinstrings.yaml line #4
		/* ({"created":1}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"created": 1}
		/* table_test_sindex_nullstr.index_create("key") */

		suite.T().Log("About to run line #4: table_test_sindex_nullstr.IndexCreate('key')")

		runAndAssert(suite.Suite, expected_, table_test_sindex_nullstr.IndexCreate("key"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #4")
	}

	{
		// sindex/nullsinstrings.yaml line #6
		/* ([{"ready":true}]) */
		var expected_ []interface{} = []interface{}{map[interface{}]interface{}{"ready": true}}
		/* table_test_sindex_nullstr.index_wait().pluck("ready") */

		suite.T().Log("About to run line #6: table_test_sindex_nullstr.IndexWait().Pluck('ready')")

		runAndAssert(suite.Suite, expected_, table_test_sindex_nullstr.IndexWait().Pluck("ready"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #6")
	}

	{
		// sindex/nullsinstrings.yaml line #10
		/* ({"inserted":2}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"inserted": 2}
		/* table_test_sindex_nullstr.insert([{"id":1,"key":["a","b"]},{"id":2,"key":["a\u0000Sb"]}]).pluck("inserted") */

		suite.T().Log("About to run line #10: table_test_sindex_nullstr.Insert([]interface{}{map[interface{}]interface{}{'id': 1, 'key': []interface{}{'a', 'b'}, }, map[interface{}]interface{}{'id': 2, 'key': []interface{}{'a\\u0000Sb'}, }}).Pluck('inserted')")

		runAndAssert(suite.Suite, expected_, table_test_sindex_nullstr.Insert([]interface{}{map[interface{}]interface{}{"id": 1, "key": []interface{}{"a", "b"}}, map[interface{}]interface{}{"id": 2, "key": []interface{}{"a\u0000Sb"}}}).Pluck("inserted"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #10")
	}

	{
		// sindex/nullsinstrings.yaml line #13
		/* ([{"id":2}]) */
		var expected_ []interface{} = []interface{}{map[interface{}]interface{}{"id": 2}}
		/* table_test_sindex_nullstr.get_all(["a\u0000Sb"], index="key").pluck("id") */

		suite.T().Log("About to run line #13: table_test_sindex_nullstr.GetAll([]interface{}{'a\\u0000Sb'}).OptArgs(r.GetAllOpts{Index: 'key', }).Pluck('id')")

		runAndAssert(suite.Suite, expected_, table_test_sindex_nullstr.GetAll([]interface{}{"a\u0000Sb"}).OptArgs(r.GetAllOpts{Index: "key"}).Pluck("id"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #13")
	}

	{
		// sindex/nullsinstrings.yaml line #18
		/* ([{"id":1}]) */
		var expected_ []interface{} = []interface{}{map[interface{}]interface{}{"id": 1}}
		/* table_test_sindex_nullstr.get_all(["a","b"], index="key").pluck("id") */

		suite.T().Log("About to run line #18: table_test_sindex_nullstr.GetAll([]interface{}{'a', 'b'}).OptArgs(r.GetAllOpts{Index: 'key', }).Pluck('id')")

		runAndAssert(suite.Suite, expected_, table_test_sindex_nullstr.GetAll([]interface{}{"a", "b"}).OptArgs(r.GetAllOpts{Index: "key"}).Pluck("id"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #18")
	}
}
