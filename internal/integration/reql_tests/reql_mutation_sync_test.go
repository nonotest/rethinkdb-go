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

// Tests syncing tables
func TestMutationSyncSuite(t *testing.T) {
	suite.Run(t, new(MutationSyncSuite))
}

type MutationSyncSuite struct {
	suite.Suite

	session *r.Session
}

func (suite *MutationSyncSuite) SetupTest() {
	suite.T().Log("Setting up MutationSyncSuite")
	// Use imports to prevent errors
	_ = time.Time{}
	_ = compare.AnythingIsFine

	session, err := r.Connect(r.ConnectOpts{
		Address: url,
	})
	suite.Require().NoError(err, "Error returned when connecting to server")
	suite.session = session

	r.DBDrop("db_mut_sync").Exec(suite.session)
	err = r.DBCreate("db_mut_sync").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("db_mut_sync").Wait().Exec(suite.session)
	suite.Require().NoError(err)

}

func (suite *MutationSyncSuite) TearDownSuite() {
	suite.T().Log("Tearing down MutationSyncSuite")

	if suite.session != nil {
		r.DB("rethinkdb").Table("_debug_scratch").Delete().Exec(suite.session)
		r.DBDrop("db_mut_sync").Exec(suite.session)

		suite.session.Close()
	}
}

func (suite *MutationSyncSuite) TestCases() {
	suite.T().Log("Running MutationSyncSuite: Tests syncing tables")

	{
		// mutation/sync.yaml line #5
		/* partial({'tables_created':1}) */
		var expected_ compare.Expected = compare.PartialMatch(map[interface{}]interface{}{"tables_created": 1})
		/* r.db('test').table_create('test1') */

		suite.T().Log("About to run line #5: r.DB('test').TableCreate('test1')")

		runAndAssert(suite.Suite, expected_, r.DB("db_mut_sync").TableCreate("test1"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #5")
	}

	{
		// mutation/sync.yaml line #7
		/* partial({'tables_created':1}) */
		var expected_ compare.Expected = compare.PartialMatch(map[interface{}]interface{}{"tables_created": 1})
		/* r.db('test').table_create('test1soft') */

		suite.T().Log("About to run line #7: r.DB('test').TableCreate('test1soft')")

		runAndAssert(suite.Suite, expected_, r.DB("db_mut_sync").TableCreate("test1soft"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #7")
	}

	{
		// mutation/sync.yaml line #9
		/* {'skipped':0, 'deleted':0, 'unchanged':0, 'errors':0, 'replaced':1, 'inserted':0} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"skipped": 0, "deleted": 0, "unchanged": 0, "errors": 0, "replaced": 1, "inserted": 0}
		/* r.db('test').table('test1soft').config().update({'durability':'soft'}) */

		suite.T().Log("About to run line #9: r.DB('test').Table('test1soft').Config().Update(map[interface{}]interface{}{'durability': 'soft', })")

		runAndAssert(suite.Suite, expected_, r.DB("db_mut_sync").Table("test1soft").Config().Update(map[interface{}]interface{}{"durability": "soft"}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #9")
	}

	// mutation/sync.yaml line #11
	// tbl = r.db('test').table('test1')
	suite.T().Log("Possibly executing: var tbl r.Term = r.DB('test').Table('test1')")

	tbl := r.DB("db_mut_sync").Table("test1")
	_ = tbl // Prevent any noused variable errors

	// mutation/sync.yaml line #12
	// tbl_soft = r.db('test').table('test1soft')
	suite.T().Log("Possibly executing: var tbl_soft r.Term = r.DB('test').Table('test1soft')")

	tbl_soft := r.DB("db_mut_sync").Table("test1soft")
	_ = tbl_soft // Prevent any noused variable errors

	{
		// mutation/sync.yaml line #13
		/* partial({'created':1}) */
		var expected_ compare.Expected = compare.PartialMatch(map[interface{}]interface{}{"created": 1})
		/* tbl.index_create('x') */

		suite.T().Log("About to run line #13: tbl.IndexCreate('x')")

		runAndAssert(suite.Suite, expected_, tbl.IndexCreate("x"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #13")
	}

	{
		// mutation/sync.yaml line #15
		/* [{'ready':True, 'index':'x'}] */
		var expected_ []interface{} = []interface{}{map[interface{}]interface{}{"ready": true, "index": "x"}}
		/* tbl.index_wait('x').pluck('index', 'ready') */

		suite.T().Log("About to run line #15: tbl.IndexWait('x').Pluck('index', 'ready')")

		runAndAssert(suite.Suite, expected_, tbl.IndexWait("x").Pluck("index", "ready"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #15")
	}

	{
		// mutation/sync.yaml line #19
		/* {'synced':1} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"synced": 1}
		/* tbl.sync() */

		suite.T().Log("About to run line #19: tbl.Sync()")

		runAndAssert(suite.Suite, expected_, tbl.Sync(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #19")
	}

	{
		// mutation/sync.yaml line #21
		/* {'synced':1} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"synced": 1}
		/* tbl_soft.sync() */

		suite.T().Log("About to run line #21: tbl_soft.Sync()")

		runAndAssert(suite.Suite, expected_, tbl_soft.Sync(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #21")
	}

	{
		// mutation/sync.yaml line #23
		/* {'synced':1} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"synced": 1}
		/* tbl.sync() */

		suite.T().Log("About to run line #23: tbl.Sync()")

		runAndAssert(suite.Suite, expected_, tbl.Sync(), suite.session, r.RunOpts{
			Durability:     "soft",
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #23")
	}

	{
		// mutation/sync.yaml line #27
		/* {'synced':1} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"synced": 1}
		/* tbl.sync() */

		suite.T().Log("About to run line #27: tbl.Sync()")

		runAndAssert(suite.Suite, expected_, tbl.Sync(), suite.session, r.RunOpts{
			Durability:     "hard",
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #27")
	}

	{
		// mutation/sync.yaml line #48
		/* partial({'tables_dropped':1}) */
		var expected_ compare.Expected = compare.PartialMatch(map[interface{}]interface{}{"tables_dropped": 1})
		/* r.db('test').table_drop('test1') */

		suite.T().Log("About to run line #48: r.DB('test').TableDrop('test1')")

		runAndAssert(suite.Suite, expected_, r.DB("db_mut_sync").TableDrop("test1"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #48")
	}

	{
		// mutation/sync.yaml line #50
		/* partial({'tables_dropped':1}) */
		var expected_ compare.Expected = compare.PartialMatch(map[interface{}]interface{}{"tables_dropped": 1})
		/* r.db('test').table_drop('test1soft') */

		suite.T().Log("About to run line #50: r.DB('test').TableDrop('test1soft')")

		runAndAssert(suite.Suite, expected_, r.DB("db_mut_sync").TableDrop("test1soft"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #50")
	}
}
