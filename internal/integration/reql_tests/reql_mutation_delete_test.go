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

// Tests deletes of selections
func TestMutationDeleteSuite(t *testing.T) {
	suite.Run(t, new(MutationDeleteSuite))
}

type MutationDeleteSuite struct {
	suite.Suite

	session *r.Session
}

func (suite *MutationDeleteSuite) SetupTest() {
	suite.T().Log("Setting up MutationDeleteSuite")
	// Use imports to prevent errors
	_ = time.Time{}
	_ = compare.AnythingIsFine

	session, err := r.Connect(r.ConnectOpts{
		Address: url,
	})
	suite.Require().NoError(err, "Error returned when connecting to server")
	suite.session = session

	r.DBDrop("db_mut_del").Exec(suite.session)
	err = r.DBCreate("db_mut_del").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("db_mut_del").Wait().Exec(suite.session)
	suite.Require().NoError(err)

	r.DB("db_mut_del").TableDrop("table_test_mutation_delete").Exec(suite.session)
	err = r.DB("db_mut_del").TableCreate("table_test_mutation_delete").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("db_mut_del").Table("table_test_mutation_delete").Wait().Exec(suite.session)
	suite.Require().NoError(err)
}

func (suite *MutationDeleteSuite) TearDownSuite() {
	suite.T().Log("Tearing down MutationDeleteSuite")

	if suite.session != nil {
		r.DB("rethinkdb").Table("_debug_scratch").Delete().Exec(suite.session)
		r.DB("db_mut_del").TableDrop("table_test_mutation_delete").Exec(suite.session)
		r.DBDrop("db_mut_del").Exec(suite.session)

		suite.session.Close()
	}
}

func (suite *MutationDeleteSuite) TestCases() {
	suite.T().Log("Running MutationDeleteSuite: Tests deletes of selections")

	table_test_mutation_delete := r.DB("db_mut_del").Table("table_test_mutation_delete")
	_ = table_test_mutation_delete // Prevent any noused variable errors

	{
		// mutation/delete.yaml line #7
		/* ({'deleted':0,'replaced':0,'unchanged':0,'errors':0,'skipped':0,'inserted':100}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"deleted": 0, "replaced": 0, "unchanged": 0, "errors": 0, "skipped": 0, "inserted": 100}
		/* table_test_mutation_delete.insert([{'id':i} for i in xrange(100)]) */

		suite.T().Log("About to run line #7: table_test_mutation_delete.Insert((func() []interface{} {\n    res := []interface{}{}\n    for iterator_ := 0; iterator_ < 100; iterator_++ {\n        i := iterator_\n        res = append(res, map[interface{}]interface{}{'id': i, })\n    }\n    return res\n}()))")

		runAndAssert(suite.Suite, expected_, table_test_mutation_delete.Insert((func() []interface{} {
			res := []interface{}{}
			for iterator_ := 0; iterator_ < 100; iterator_++ {
				i := iterator_
				res = append(res, map[interface{}]interface{}{"id": i})
			}
			return res
		}())), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #7")
	}

	{
		// mutation/delete.yaml line #19
		/* 100 */
		var expected_ int = 100
		/* table_test_mutation_delete.count() */

		suite.T().Log("About to run line #19: table_test_mutation_delete.Count()")

		runAndAssert(suite.Suite, expected_, table_test_mutation_delete.Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #19")
	}

	{
		// mutation/delete.yaml line #24
		/* ({'deleted':1,'replaced':0,'unchanged':0,'errors':0,'skipped':0,'inserted':0}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"deleted": 1, "replaced": 0, "unchanged": 0, "errors": 0, "skipped": 0, "inserted": 0}
		/* table_test_mutation_delete.get(12).delete() */

		suite.T().Log("About to run line #24: table_test_mutation_delete.Get(12).Delete()")

		runAndAssert(suite.Suite, expected_, table_test_mutation_delete.Get(12).Delete(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #24")
	}

	{
		// mutation/delete.yaml line #31
		/* err('ReqlQueryLogicError', 'Durability option `wrong` unrecognized (options are "hard" and "soft").', [0]) */
		var expected_ Err = err("ReqlQueryLogicError", "Durability option `wrong` unrecognized (options are \"hard\" and \"soft\").")
		/* table_test_mutation_delete.skip(50).delete(durability='wrong') */

		suite.T().Log("About to run line #31: table_test_mutation_delete.Skip(50).Delete().OptArgs(r.DeleteOpts{Durability: 'wrong', })")

		runAndAssert(suite.Suite, expected_, table_test_mutation_delete.Skip(50).Delete().OptArgs(r.DeleteOpts{Durability: "wrong"}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #31")
	}

	{
		// mutation/delete.yaml line #38
		/* ({'deleted':49,'replaced':0,'unchanged':0,'errors':0,'skipped':0,'inserted':0}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"deleted": 49, "replaced": 0, "unchanged": 0, "errors": 0, "skipped": 0, "inserted": 0}
		/* table_test_mutation_delete.skip(50).delete(durability='soft') */

		suite.T().Log("About to run line #38: table_test_mutation_delete.Skip(50).Delete().OptArgs(r.DeleteOpts{Durability: 'soft', })")

		runAndAssert(suite.Suite, expected_, table_test_mutation_delete.Skip(50).Delete().OptArgs(r.DeleteOpts{Durability: "soft"}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #38")
	}

	{
		// mutation/delete.yaml line #45
		/* ({'deleted':50,'replaced':0,'unchanged':0,'errors':0,'skipped':0,'inserted':0}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"deleted": 50, "replaced": 0, "unchanged": 0, "errors": 0, "skipped": 0, "inserted": 0}
		/* table_test_mutation_delete.delete(durability='hard') */

		suite.T().Log("About to run line #45: table_test_mutation_delete.Delete().OptArgs(r.DeleteOpts{Durability: 'hard', })")

		runAndAssert(suite.Suite, expected_, table_test_mutation_delete.Delete().OptArgs(r.DeleteOpts{Durability: "hard"}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #45")
	}

	{
		// mutation/delete.yaml line #49
		/* err('ReqlQueryLogicError', 'Expected type SELECTION but found DATUM:', [0]) */
		var expected_ Err = err("ReqlQueryLogicError", "Expected type SELECTION but found DATUM:")
		/* r.expr([1, 2]).delete() */

		suite.T().Log("About to run line #49: r.Expr([]interface{}{1, 2}).Delete()")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2}).Delete(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #49")
	}
}
