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

// Tests for the basic usage of the multiplication operation
func TestMathLogicMulSuite(t *testing.T) {
	suite.Run(t, new(MathLogicMulSuite))
}

type MathLogicMulSuite struct {
	suite.Suite

	session *r.Session
}

func (suite *MathLogicMulSuite) SetupTest() {
	suite.T().Log("Setting up MathLogicMulSuite")
	// Use imports to prevent errors
	_ = time.Time{}
	_ = compare.AnythingIsFine

	session, err := r.Connect(r.ConnectOpts{
		Address: url,
	})
	suite.Require().NoError(err, "Error returned when connecting to server")
	suite.session = session

	r.DBDrop("db_logic_mul").Exec(suite.session)
	err = r.DBCreate("db_logic_mul").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("db_logic_mul").Wait().Exec(suite.session)
	suite.Require().NoError(err)

}

func (suite *MathLogicMulSuite) TearDownSuite() {
	suite.T().Log("Tearing down MathLogicMulSuite")

	if suite.session != nil {
		r.DB("rethinkdb").Table("_debug_scratch").Delete().Exec(suite.session)
		r.DBDrop("db_logic_mul").Exec(suite.session)

		suite.session.Close()
	}
}

func (suite *MathLogicMulSuite) TestCases() {
	suite.T().Log("Running MathLogicMulSuite: Tests for the basic usage of the multiplication operation")

	{
		// math_logic/mul.yaml line #6
		/* 2 */
		var expected_ int = 2
		/* r.expr(1) * 2 */

		suite.T().Log("About to run line #6: r.Expr(1).Mul(2)")

		runAndAssert(suite.Suite, expected_, r.Expr(1).Mul(2), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #6")
	}

	{
		// math_logic/mul.yaml line #7
		/* 2 */
		var expected_ int = 2
		/* 1 * r.expr(2) */

		suite.T().Log("About to run line #7: r.Mul(1, r.Expr(2))")

		runAndAssert(suite.Suite, expected_, r.Mul(1, r.Expr(2)), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #7")
	}

	{
		// math_logic/mul.yaml line #8
		/* 2 */
		var expected_ int = 2
		/* r.expr(1).mul(2) */

		suite.T().Log("About to run line #8: r.Expr(1).Mul(2)")

		runAndAssert(suite.Suite, expected_, r.Expr(1).Mul(2), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #8")
	}

	{
		// math_logic/mul.yaml line #15
		/* 1 */
		var expected_ int = 1
		/* r.expr(-1) * -1 */

		suite.T().Log("About to run line #15: r.Expr(-1).Mul(-1)")

		runAndAssert(suite.Suite, expected_, r.Expr(-1).Mul(-1), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #15")
	}

	{
		// math_logic/mul.yaml line #21
		/* 6.75 */
		var expected_ float64 = 6.75
		/* r.expr(1.5) * 4.5 */

		suite.T().Log("About to run line #21: r.Expr(1.5).Mul(4.5)")

		runAndAssert(suite.Suite, expected_, r.Expr(1.5).Mul(4.5), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #21")
	}

	{
		// math_logic/mul.yaml line #25
		/* [1,2,3,1,2,3,1,2,3] */
		var expected_ []interface{} = []interface{}{1, 2, 3, 1, 2, 3, 1, 2, 3}
		/* r.expr([1,2,3]) * 3 */

		suite.T().Log("About to run line #25: r.Expr([]interface{}{1, 2, 3}).Mul(3)")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2, 3}).Mul(3), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #25")
	}

	{
		// math_logic/mul.yaml line #30
		/* 120 */
		var expected_ int = 120
		/* r.expr(1).mul(2,3,4,5) */

		suite.T().Log("About to run line #30: r.Expr(1).Mul(2, 3, 4, 5)")

		runAndAssert(suite.Suite, expected_, r.Expr(1).Mul(2, 3, 4, 5), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #30")
	}

	{
		// math_logic/mul.yaml line #46
		/* err('ReqlQueryLogicError', 'Expected type NUMBER but found STRING.', [0]) */
		var expected_ Err = err("ReqlQueryLogicError", "Expected type NUMBER but found STRING.")
		/* r.expr('a') * 0.8 */

		suite.T().Log("About to run line #46: r.Expr('a').Mul(0.8)")

		runAndAssert(suite.Suite, expected_, r.Expr("a").Mul(0.8), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #46")
	}

	{
		// math_logic/mul.yaml line #50
		/* err('ReqlQueryLogicError', 'Expected type NUMBER but found STRING.', [1]) */
		var expected_ Err = err("ReqlQueryLogicError", "Expected type NUMBER but found STRING.")
		/* r.expr(1) * 'a' */

		suite.T().Log("About to run line #50: r.Expr(1).Mul('a')")

		runAndAssert(suite.Suite, expected_, r.Expr(1).Mul("a"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #50")
	}

	{
		// math_logic/mul.yaml line #54
		/* err('ReqlQueryLogicError', 'Expected type NUMBER but found STRING.', [0]) */
		var expected_ Err = err("ReqlQueryLogicError", "Expected type NUMBER but found STRING.")
		/* r.expr('b') * 'a' */

		suite.T().Log("About to run line #54: r.Expr('b').Mul('a')")

		runAndAssert(suite.Suite, expected_, r.Expr("b").Mul("a"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #54")
	}

	{
		// math_logic/mul.yaml line #58
		/* err('ReqlQueryLogicError', 'Number not an integer: 1.5', [0]) */
		var expected_ Err = err("ReqlQueryLogicError", "Number not an integer: 1.5")
		/* r.expr([]) * 1.5 */

		suite.T().Log("About to run line #58: r.Expr([]interface{}{}).Mul(1.5)")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{}).Mul(1.5), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #58")
	}
}
