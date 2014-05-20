package query

import (
	"fmt"
	. "launchpad.net/gocheck"
	"testing"
)

func Test(t *testing.T) {
	TestingT(t)
}

type QueryParserSuite struct{}

var _ = Suite(&QueryParserSuite{})

func (self *QueryParserSuite) TestParseBasicSelectQuery(c *C) {
	q, err := ParseQuery("select value, time from t where c = '5';")
	c.Assert(err, IsNil)
	fmt.Printf("%+v\n", q.GetColumnNames())
	//c.Assert(q.GetColumnNames(), DeepEquals, []string{"value", "time"})
	w := q.GetWhereClause()
	c.Assert(q.GetFromClause().TableName, Equals, "t")
	c.Assert(w.ColumnName, Equals, "c")
	c.Assert(w.Value, Equals, "5")
}

func (self *QueryParserSuite) TestParseSelectWithUpperCase(c *C) {
	q, err := ParseQuery("SELECT VALUE, TIME FROM t WHERE C = '5';")
	c.Assert(err, IsNil)
	c.Assert(q.GetColumnNames(), DeepEquals, []string{"VALUE", "TIME"})
	w := q.GetWhereClause()
	c.Assert(w.ColumnName, Equals, "C")
	c.Assert(w.Value, Equals, "5")
}


func (self *QueryParserSuite) TestParseSelectWithInequality(c *C) {
	q, err := ParseQuery("SELECT VALUE, TIME FROM t WHERE C < '5';")
	//q, err := ParseQuery("select value, time from t where c < 5;")
	c.Assert(err, IsNil)
	//c.Assert(q.GetColumnNames(), DeepEquals, []string{"value", "time"})
	w := q.GetWhereClause()
	c.Assert(q.GetFromClause().TableName, Equals, "t")
	c.Assert(w.ColumnName, Equals, "C")
	c.Assert(int(w.Op), Equals, LESS_THAN)
	// TODO: fix this
	c.Assert(w.Value, Equals, "5")
}

func (self *QueryParserSuite) TestParseSelectWithTimeCondition(c *C) {
	q, err := ParseQuery("select value, time FROM t WHERE time > 1;")
	//q, err := ParseQuery("select value, time from t where time > now() - 1d;")
	c.Assert(err, IsNil)
	c.Assert(q.GetColumnNames(), DeepEquals, []string{"value", "time"})
	w := q.GetWhereClause()
	c.Assert(q.GetFromClause().TableName, Equals, "t")
	c.Assert(w.ColumnName, Equals, "time")
	//c.Assert(w.Value, Equals, "5")
}


