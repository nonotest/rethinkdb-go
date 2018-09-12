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

// Tests updates of selections
func TestMutationUpdateSuite(t *testing.T) {
	suite.Run(t, new(MutationUpdateSuite))
}

type MutationUpdateSuite struct {
	suite.Suite

	session *r.Session
}

func (suite *MutationUpdateSuite) SetupTest() {
	suite.T().Log("Setting up MutationUpdateSuite")
	// Use imports to prevent errors
	_ = time.Time{}
	_ = compare.AnythingIsFine

	session, err := r.Connect(r.ConnectOpts{
		Address: url,
	})
	suite.Require().NoError(err, "Error returned when connecting to server")
	suite.session = session

	r.DBDrop("db_mut_upd").Exec(suite.session)
	err = r.DBCreate("db_mut_upd").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("db_mut_upd").Wait().Exec(suite.session)
	suite.Require().NoError(err)

	r.DB("db_mut_upd").TableDrop("table_test_mutation_update").Exec(suite.session)
	err = r.DB("db_mut_upd").TableCreate("table_test_mutation_update").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("db_mut_upd").Table("table_test_mutation_update").Wait().Exec(suite.session)
	suite.Require().NoError(err)
	r.DB("db_mut_upd").TableDrop("table_test_mutation_update2").Exec(suite.session)
	err = r.DB("db_mut_upd").TableCreate("table_test_mutation_update2").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("db_mut_upd").Table("table_test_mutation_update2").Wait().Exec(suite.session)
	suite.Require().NoError(err)
}

func (suite *MutationUpdateSuite) TearDownSuite() {
	suite.T().Log("Tearing down MutationUpdateSuite")

	if suite.session != nil {
		r.DB("rethinkdb").Table("_debug_scratch").Delete().Exec(suite.session)
		r.DB("db_mut_upd").TableDrop("table_test_mutation_update").Exec(suite.session)
		r.DB("db_mut_upd").TableDrop("table_test_mutation_update2").Exec(suite.session)
		r.DBDrop("db_mut_upd").Exec(suite.session)

		suite.session.Close()
	}
}

func (suite *MutationUpdateSuite) TestCases() {
	suite.T().Log("Running MutationUpdateSuite: Tests updates of selections")

	table_test_mutation_update := r.DB("db_mut_upd").Table("table_test_mutation_update")
	_ = table_test_mutation_update // Prevent any noused variable errors
	table_test_mutation_update2 := r.DB("db_mut_upd").Table("table_test_mutation_update2")
	_ = table_test_mutation_update2 // Prevent any noused variable errors

	{
		// mutation/update.yaml line #6
		/* ({'deleted':0.0,'replaced':0.0,'unchanged':0.0,'errors':0.0,'skipped':0.0,'inserted':100}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"deleted": 0.0, "replaced": 0.0, "unchanged": 0.0, "errors": 0.0, "skipped": 0.0, "inserted": 100}
		/* table_test_mutation_update.insert([{'id':i} for i in xrange(100)]) */

		suite.T().Log("About to run line #6: table_test_mutation_update.Insert((func() []interface{} {\n    res := []interface{}{}\n    for iterator_ := 0; iterator_ < 100; iterator_++ {\n        i := iterator_\n        res = append(res, map[interface{}]interface{}{'id': i, })\n    }\n    return res\n}()))")

		runAndAssert(suite.Suite, expected_, table_test_mutation_update.Insert((func() []interface{} {
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
		suite.T().Log("Finished running line #6")
	}

	{
		// mutation/update.yaml line #18
		/* 100 */
		var expected_ int = 100
		/* table_test_mutation_update.count() */

		suite.T().Log("About to run line #18: table_test_mutation_update.Count()")

		runAndAssert(suite.Suite, expected_, table_test_mutation_update.Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #18")
	}

	{
		// mutation/update.yaml line #21
		/* ({'deleted':0.0,'replaced':0.0,'unchanged':0.0,'errors':0.0,'skipped':0.0,'inserted':100}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"deleted": 0.0, "replaced": 0.0, "unchanged": 0.0, "errors": 0.0, "skipped": 0.0, "inserted": 100}
		/* table_test_mutation_update2.insert([{'id':i, 'foo':{'bar':i}} for i in xrange(100)]) */

		suite.T().Log("About to run line #21: table_test_mutation_update2.Insert((func() []interface{} {\n    res := []interface{}{}\n    for iterator_ := 0; iterator_ < 100; iterator_++ {\n        i := iterator_\n        res = append(res, map[interface{}]interface{}{'id': i, 'foo': map[interface{}]interface{}{'bar': i, }, })\n    }\n    return res\n}()))")

		runAndAssert(suite.Suite, expected_, table_test_mutation_update2.Insert((func() []interface{} {
			res := []interface{}{}
			for iterator_ := 0; iterator_ < 100; iterator_++ {
				i := iterator_
				res = append(res, map[interface{}]interface{}{"id": i, "foo": map[interface{}]interface{}{"bar": i}})
			}
			return res
		}())), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #21")
	}

	{
		// mutation/update.yaml line #33
		/* 100 */
		var expected_ int = 100
		/* table_test_mutation_update2.count() */

		suite.T().Log("About to run line #33: table_test_mutation_update2.Count()")

		runAndAssert(suite.Suite, expected_, table_test_mutation_update2.Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #33")
	}

	{
		// mutation/update.yaml line #37
		/* {'deleted':0.0,'replaced':0.0,'unchanged':1,'errors':0.0,'skipped':0.0,'inserted':0.0} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"deleted": 0.0, "replaced": 0.0, "unchanged": 1, "errors": 0.0, "skipped": 0.0, "inserted": 0.0}
		/* table_test_mutation_update.get(12).update(lambda row:row) */

		suite.T().Log("About to run line #37: table_test_mutation_update.Get(12).Update(func(row r.Term) interface{} { return row})")

		runAndAssert(suite.Suite, expected_, table_test_mutation_update.Get(12).Update(func(row r.Term) interface{} { return row }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #37")
	}

	{
		// mutation/update.yaml line #43
		/* {'deleted':0.0,'replaced':1,'unchanged':0.0,'errors':0.0,'skipped':0.0,'inserted':0.0} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"deleted": 0.0, "replaced": 1, "unchanged": 0.0, "errors": 0.0, "skipped": 0.0, "inserted": 0.0}
		/* table_test_mutation_update.get(12).update(lambda row:{'a':row['id'] + 1}, durability='soft') */

		suite.T().Log("About to run line #43: table_test_mutation_update.Get(12).Update(func(row r.Term) interface{} { return map[interface{}]interface{}{'a': row.AtIndex('id').Add(1), }}).OptArgs(r.UpdateOpts{Durability: 'soft', })")

		runAndAssert(suite.Suite, expected_, table_test_mutation_update.Get(12).Update(func(row r.Term) interface{} { return map[interface{}]interface{}{"a": row.AtIndex("id").Add(1)} }).OptArgs(r.UpdateOpts{Durability: "soft"}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #43")
	}

	{
		// mutation/update.yaml line #48
		/* {'id':12, 'a':13} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"id": 12, "a": 13}
		/* table_test_mutation_update.get(12) */

		suite.T().Log("About to run line #48: table_test_mutation_update.Get(12)")

		runAndAssert(suite.Suite, expected_, table_test_mutation_update.Get(12), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #48")
	}

	{
		// mutation/update.yaml line #52
		/* {'deleted':0.0,'replaced':1,'unchanged':0.0,'errors':0.0,'skipped':0.0,'inserted':0.0} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"deleted": 0.0, "replaced": 1, "unchanged": 0.0, "errors": 0.0, "skipped": 0.0, "inserted": 0.0}
		/* table_test_mutation_update.get(12).update(lambda row:{'a':row['id'] + 2}, durability='hard') */

		suite.T().Log("About to run line #52: table_test_mutation_update.Get(12).Update(func(row r.Term) interface{} { return map[interface{}]interface{}{'a': row.AtIndex('id').Add(2), }}).OptArgs(r.UpdateOpts{Durability: 'hard', })")

		runAndAssert(suite.Suite, expected_, table_test_mutation_update.Get(12).Update(func(row r.Term) interface{} { return map[interface{}]interface{}{"a": row.AtIndex("id").Add(2)} }).OptArgs(r.UpdateOpts{Durability: "hard"}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #52")
	}

	{
		// mutation/update.yaml line #57
		/* {'id':12, 'a':14} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"id": 12, "a": 14}
		/* table_test_mutation_update.get(12) */

		suite.T().Log("About to run line #57: table_test_mutation_update.Get(12)")

		runAndAssert(suite.Suite, expected_, table_test_mutation_update.Get(12), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #57")
	}

	{
		// mutation/update.yaml line #61
		/* err('ReqlQueryLogicError', 'Durability option `wrong` unrecognized (options are "hard" and "soft").', [0]) */
		var expected_ Err = err("ReqlQueryLogicError", "Durability option `wrong` unrecognized (options are \"hard\" and \"soft\").")
		/* table_test_mutation_update.get(12).update(lambda row:{'a':row['id'] + 3}, durability='wrong') */

		suite.T().Log("About to run line #61: table_test_mutation_update.Get(12).Update(func(row r.Term) interface{} { return map[interface{}]interface{}{'a': row.AtIndex('id').Add(3), }}).OptArgs(r.UpdateOpts{Durability: 'wrong', })")

		runAndAssert(suite.Suite, expected_, table_test_mutation_update.Get(12).Update(func(row r.Term) interface{} { return map[interface{}]interface{}{"a": row.AtIndex("id").Add(3)} }).OptArgs(r.UpdateOpts{Durability: "wrong"}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #61")
	}

	{
		// mutation/update.yaml line #66
		/* {'id':12, 'a':14} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"id": 12, "a": 14}
		/* table_test_mutation_update.get(12) */

		suite.T().Log("About to run line #66: table_test_mutation_update.Get(12)")

		runAndAssert(suite.Suite, expected_, table_test_mutation_update.Get(12), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #66")
	}

	{
		// mutation/update.yaml line #70
		/* {'deleted':0.0,'replaced':1,'unchanged':0.0,'errors':0.0,'skipped':0.0,'inserted':0.0} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"deleted": 0.0, "replaced": 1, "unchanged": 0.0, "errors": 0.0, "skipped": 0.0, "inserted": 0.0}
		/* table_test_mutation_update.get(12).update(lambda row:{'a':row['id']}) */

		suite.T().Log("About to run line #70: table_test_mutation_update.Get(12).Update(func(row r.Term) interface{} { return map[interface{}]interface{}{'a': row.AtIndex('id'), }})")

		runAndAssert(suite.Suite, expected_, table_test_mutation_update.Get(12).Update(func(row r.Term) interface{} { return map[interface{}]interface{}{"a": row.AtIndex("id")} }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #70")
	}

	{
		// mutation/update.yaml line #75
		/* {'id':12, 'a':12} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"id": 12, "a": 12}
		/* table_test_mutation_update.get(12) */

		suite.T().Log("About to run line #75: table_test_mutation_update.Get(12)")

		runAndAssert(suite.Suite, expected_, table_test_mutation_update.Get(12), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #75")
	}

	{
		// mutation/update.yaml line #79
		/* {'deleted':0.0,'replaced':1,'unchanged':0.0,'errors':0.0,'skipped':0.0,'inserted':0.0} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"deleted": 0.0, "replaced": 1, "unchanged": 0.0, "errors": 0.0, "skipped": 0.0, "inserted": 0.0}
		/* table_test_mutation_update.get(12).update({'a':r.literal()}) */

		suite.T().Log("About to run line #79: table_test_mutation_update.Get(12).Update(map[interface{}]interface{}{'a': r.Literal(), })")

		runAndAssert(suite.Suite, expected_, table_test_mutation_update.Get(12).Update(map[interface{}]interface{}{"a": r.Literal()}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #79")
	}

	{
		// mutation/update.yaml line #84
		/* {'deleted':0.0,'replaced':10,'unchanged':0,'errors':0.0,'skipped':0.0,'inserted':0.0} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"deleted": 0.0, "replaced": 10, "unchanged": 0, "errors": 0.0, "skipped": 0.0, "inserted": 0.0}
		/* table_test_mutation_update.between(10, 20).update(lambda row:{'a':row['id']}) */

		suite.T().Log("About to run line #84: table_test_mutation_update.Between(10, 20).Update(func(row r.Term) interface{} { return map[interface{}]interface{}{'a': row.AtIndex('id'), }})")

		runAndAssert(suite.Suite, expected_, table_test_mutation_update.Between(10, 20).Update(func(row r.Term) interface{} { return map[interface{}]interface{}{"a": row.AtIndex("id")} }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #84")
	}

	{
		// mutation/update.yaml line #89
		/* {'deleted':0.0,'replaced':0.0,'unchanged':10,'errors':0.0,'skipped':0.0,'inserted':0.0} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"deleted": 0.0, "replaced": 0.0, "unchanged": 10, "errors": 0.0, "skipped": 0.0, "inserted": 0.0}
		/* table_test_mutation_update.filter(lambda row:(row['id'] >= 10) & (row['id'] < 20)).update(lambda row:{'a':row['id']}) */

		suite.T().Log("About to run line #89: table_test_mutation_update.Filter(func(row r.Term) interface{} { return row.AtIndex('id').Ge(10).And(row.AtIndex('id').Lt(20))}).Update(func(row r.Term) interface{} { return map[interface{}]interface{}{'a': row.AtIndex('id'), }})")

		runAndAssert(suite.Suite, expected_, table_test_mutation_update.Filter(func(row r.Term) interface{} { return row.AtIndex("id").Ge(10).And(row.AtIndex("id").Lt(20)) }).Update(func(row r.Term) interface{} { return map[interface{}]interface{}{"a": row.AtIndex("id")} }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #89")
	}

	{
		// mutation/update.yaml line #94
		/* {'deleted':0.0,'replaced':10,'unchanged':0.0,'errors':0.0,'skipped':0.0,'inserted':0.0} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"deleted": 0.0, "replaced": 10, "unchanged": 0.0, "errors": 0.0, "skipped": 0.0, "inserted": 0.0}
		/* table_test_mutation_update.filter(lambda row:(row['id'] >= 10) & (row['id'] < 20)).update(lambda row:{'b':row['id']}) */

		suite.T().Log("About to run line #94: table_test_mutation_update.Filter(func(row r.Term) interface{} { return row.AtIndex('id').Ge(10).And(row.AtIndex('id').Lt(20))}).Update(func(row r.Term) interface{} { return map[interface{}]interface{}{'b': row.AtIndex('id'), }})")

		runAndAssert(suite.Suite, expected_, table_test_mutation_update.Filter(func(row r.Term) interface{} { return row.AtIndex("id").Ge(10).And(row.AtIndex("id").Lt(20)) }).Update(func(row r.Term) interface{} { return map[interface{}]interface{}{"b": row.AtIndex("id")} }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #94")
	}

	{
		// mutation/update.yaml line #100
		/* {'deleted':0.0,'replaced':10,'unchanged':0,'errors':0.0,'skipped':0.0,'inserted':0.0} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"deleted": 0.0, "replaced": 10, "unchanged": 0, "errors": 0.0, "skipped": 0.0, "inserted": 0.0}
		/* table_test_mutation_update.between(10, 20).update({'a':r.literal()}) */

		suite.T().Log("About to run line #100: table_test_mutation_update.Between(10, 20).Update(map[interface{}]interface{}{'a': r.Literal(), })")

		runAndAssert(suite.Suite, expected_, table_test_mutation_update.Between(10, 20).Update(map[interface{}]interface{}{"a": r.Literal()}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #100")
	}

	{
		// mutation/update.yaml line #104
		/* {'first_error':"Primary key `id` cannot be changed (`{\n\t\"id\":\t1\n}` -> `{\n\t\"d\":\t1,\n\t\"id\":\t2\n}`).",'deleted':0.0,'replaced':0.0,'unchanged':0.0,'errors':1,'skipped':0.0,'inserted':0.0} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"first_error": "Primary key `id` cannot be changed (`{\n\t\"id\":\t1\n}` -> `{\n\t\"d\":\t1,\n\t\"id\":\t2\n}`).", "deleted": 0.0, "replaced": 0.0, "unchanged": 0.0, "errors": 1, "skipped": 0.0, "inserted": 0.0}
		/* table_test_mutation_update.get(1).update({'id':2,'d':1}) */

		suite.T().Log("About to run line #104: table_test_mutation_update.Get(1).Update(map[interface{}]interface{}{'id': 2, 'd': 1, })")

		runAndAssert(suite.Suite, expected_, table_test_mutation_update.Get(1).Update(map[interface{}]interface{}{"id": 2, "d": 1}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #104")
	}

	{
		// mutation/update.yaml line #108
		/* {'deleted':0.0,'replaced':1,'unchanged':0.0,'errors':0.0,'skipped':0.0,'inserted':0.0} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"deleted": 0.0, "replaced": 1, "unchanged": 0.0, "errors": 0.0, "skipped": 0.0, "inserted": 0.0}
		/* table_test_mutation_update.get(1).update({'id':r.row['id'],'d':'b'}) */

		suite.T().Log("About to run line #108: table_test_mutation_update.Get(1).Update(map[interface{}]interface{}{'id': r.Row.AtIndex('id'), 'd': 'b', })")

		runAndAssert(suite.Suite, expected_, table_test_mutation_update.Get(1).Update(map[interface{}]interface{}{"id": r.Row.AtIndex("id"), "d": "b"}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #108")
	}

	{
		// mutation/update.yaml line #114
		/* {'deleted':0.0,'replaced':0.0,'unchanged':1,'errors':0.0,'skipped':0.0,'inserted':0.0} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"deleted": 0.0, "replaced": 0.0, "unchanged": 1, "errors": 0.0, "skipped": 0.0, "inserted": 0.0}
		/* table_test_mutation_update.get(1).update(r.row.merge({'d':'b'})) */

		suite.T().Log("About to run line #114: table_test_mutation_update.Get(1).Update(r.Row.Merge(map[interface{}]interface{}{'d': 'b', }))")

		runAndAssert(suite.Suite, expected_, table_test_mutation_update.Get(1).Update(r.Row.Merge(map[interface{}]interface{}{"d": "b"})), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #114")
	}

	{
		// mutation/update.yaml line #119
		/* err('ReqlQueryLogicError', 'Could not prove argument deterministic.  Maybe you want to use the non_atomic flag?', [0]) */
		var expected_ Err = err("ReqlQueryLogicError", "Could not prove argument deterministic.  Maybe you want to use the non_atomic flag?")
		/* table_test_mutation_update.get(1).update({'d':r.js('5')}) */

		suite.T().Log("About to run line #119: table_test_mutation_update.Get(1).Update(map[interface{}]interface{}{'d': r.JS('5'), })")

		runAndAssert(suite.Suite, expected_, table_test_mutation_update.Get(1).Update(map[interface{}]interface{}{"d": r.JS("5")}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #119")
	}

	{
		// mutation/update.yaml line #122
		/* err('ReqlQueryLogicError', 'Could not prove argument deterministic.  Maybe you want to use the non_atomic flag?', [0]) */
		var expected_ Err = err("ReqlQueryLogicError", "Could not prove argument deterministic.  Maybe you want to use the non_atomic flag?")
		/* table_test_mutation_update.get(1).update({'d':table_test_mutation_update.nth(0)}) */

		suite.T().Log("About to run line #122: table_test_mutation_update.Get(1).Update(map[interface{}]interface{}{'d': table_test_mutation_update.Nth(0), })")

		runAndAssert(suite.Suite, expected_, table_test_mutation_update.Get(1).Update(map[interface{}]interface{}{"d": table_test_mutation_update.Nth(0)}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #122")
	}

	{
		// mutation/update.yaml line #125
		/* {'deleted':0.0,'replaced':1,'unchanged':0.0,'errors':0.0,'skipped':0.0,'inserted':0.0} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"deleted": 0.0, "replaced": 1, "unchanged": 0.0, "errors": 0.0, "skipped": 0.0, "inserted": 0.0}
		/* table_test_mutation_update.get(1).update({'d':r.js('5')}, non_atomic=True) */

		suite.T().Log("About to run line #125: table_test_mutation_update.Get(1).Update(map[interface{}]interface{}{'d': r.JS('5'), }).OptArgs(r.UpdateOpts{NonAtomic: true, })")

		runAndAssert(suite.Suite, expected_, table_test_mutation_update.Get(1).Update(map[interface{}]interface{}{"d": r.JS("5")}).OptArgs(r.UpdateOpts{NonAtomic: true}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #125")
	}

	{
		// mutation/update.yaml line #137
		/* {'deleted':0.0,'replaced':100,'unchanged':0,'errors':0.0,'skipped':0.0,'inserted':0.0} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"deleted": 0.0, "replaced": 100, "unchanged": 0, "errors": 0.0, "skipped": 0.0, "inserted": 0.0}
		/* table_test_mutation_update.update(lambda row:{'a':row['id']}) */

		suite.T().Log("About to run line #137: table_test_mutation_update.Update(func(row r.Term) interface{} { return map[interface{}]interface{}{'a': row.AtIndex('id'), }})")

		runAndAssert(suite.Suite, expected_, table_test_mutation_update.Update(func(row r.Term) interface{} { return map[interface{}]interface{}{"a": row.AtIndex("id")} }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #137")
	}

	{
		// mutation/update.yaml line #143
		/* {'deleted':0.0,'replaced':100,'unchanged':0,'errors':0.0,'skipped':0.0,'inserted':0.0} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"deleted": 0.0, "replaced": 100, "unchanged": 0, "errors": 0.0, "skipped": 0.0, "inserted": 0.0}
		/* table_test_mutation_update.update({'a':r.literal()}) */

		suite.T().Log("About to run line #143: table_test_mutation_update.Update(map[interface{}]interface{}{'a': r.Literal(), })")

		runAndAssert(suite.Suite, expected_, table_test_mutation_update.Update(map[interface{}]interface{}{"a": r.Literal()}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #143")
	}

	{
		// mutation/update.yaml line #147
		/* {'deleted':0.0,'replaced':99,'unchanged':1,'errors':0.0,'skipped':0.0,'inserted':0.0} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"deleted": 0.0, "replaced": 99, "unchanged": 1, "errors": 0.0, "skipped": 0.0, "inserted": 0.0}
		/* table_test_mutation_update2.update({'foo':{'bar':2}}) */

		suite.T().Log("About to run line #147: table_test_mutation_update2.Update(map[interface{}]interface{}{'foo': map[interface{}]interface{}{'bar': 2, }, })")

		runAndAssert(suite.Suite, expected_, table_test_mutation_update2.Update(map[interface{}]interface{}{"foo": map[interface{}]interface{}{"bar": 2}}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #147")
	}

	{
		// mutation/update.yaml line #150
		/* {'deleted':0.0,'replaced':0,'unchanged':100,'errors':0.0,'skipped':0.0,'inserted':0.0} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"deleted": 0.0, "replaced": 0, "unchanged": 100, "errors": 0.0, "skipped": 0.0, "inserted": 0.0}
		/* table_test_mutation_update2.update({'foo':r.literal({'bar':2})}) */

		suite.T().Log("About to run line #150: table_test_mutation_update2.Update(map[interface{}]interface{}{'foo': r.Literal(map[interface{}]interface{}{'bar': 2, }), })")

		runAndAssert(suite.Suite, expected_, table_test_mutation_update2.Update(map[interface{}]interface{}{"foo": r.Literal(map[interface{}]interface{}{"bar": 2})}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #150")
	}

	{
		// mutation/update.yaml line #156
		/* {'id':0,'foo':{'bar':2}} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"id": 0, "foo": map[interface{}]interface{}{"bar": 2}}
		/* table_test_mutation_update2.order_by('id').nth(0) */

		suite.T().Log("About to run line #156: table_test_mutation_update2.OrderBy('id').Nth(0)")

		runAndAssert(suite.Suite, expected_, table_test_mutation_update2.OrderBy("id").Nth(0), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #156")
	}

	{
		// mutation/update.yaml line #159
		/* {'deleted':0.0,'replaced':100,'unchanged':0,'errors':0.0,'skipped':0.0,'inserted':0.0} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"deleted": 0.0, "replaced": 100, "unchanged": 0, "errors": 0.0, "skipped": 0.0, "inserted": 0.0}
		/* table_test_mutation_update2.update({'foo':{'buzz':2}}) */

		suite.T().Log("About to run line #159: table_test_mutation_update2.Update(map[interface{}]interface{}{'foo': map[interface{}]interface{}{'buzz': 2, }, })")

		runAndAssert(suite.Suite, expected_, table_test_mutation_update2.Update(map[interface{}]interface{}{"foo": map[interface{}]interface{}{"buzz": 2}}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #159")
	}

	{
		// mutation/update.yaml line #162
		/* {'id':0,'foo':{'bar':2,'buzz':2}} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"id": 0, "foo": map[interface{}]interface{}{"bar": 2, "buzz": 2}}
		/* table_test_mutation_update2.order_by('id').nth(0) */

		suite.T().Log("About to run line #162: table_test_mutation_update2.OrderBy('id').Nth(0)")

		runAndAssert(suite.Suite, expected_, table_test_mutation_update2.OrderBy("id").Nth(0), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #162")
	}

	{
		// mutation/update.yaml line #165
		/* {'deleted':0.0,'replaced':100,'unchanged':0,'errors':0.0,'skipped':0.0,'inserted':0.0} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"deleted": 0.0, "replaced": 100, "unchanged": 0, "errors": 0.0, "skipped": 0.0, "inserted": 0.0}
		/* table_test_mutation_update2.update({'foo':r.literal(1)}) */

		suite.T().Log("About to run line #165: table_test_mutation_update2.Update(map[interface{}]interface{}{'foo': r.Literal(1), })")

		runAndAssert(suite.Suite, expected_, table_test_mutation_update2.Update(map[interface{}]interface{}{"foo": r.Literal(1)}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #165")
	}

	{
		// mutation/update.yaml line #168
		/* {'id':0,'foo':1} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"id": 0, "foo": 1}
		/* table_test_mutation_update2.order_by('id').nth(0) */

		suite.T().Log("About to run line #168: table_test_mutation_update2.OrderBy('id').Nth(0)")

		runAndAssert(suite.Suite, expected_, table_test_mutation_update2.OrderBy("id").Nth(0), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #168")
	}
}
