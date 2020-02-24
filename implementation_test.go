package main

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

/* Example */

func (s *MySuite) ExamplePrefixToPostfix(c *C) {
	res, err := PrefixToPostfix("* + 22 45 - 3 17")
	c.Check(err, Equals, nil)
	c.Check(res, Equals, "22 4 + 3 17 - *")
}

/* Simple test cases */

func (s *MySuite) TestPrefixToPostfixSimpleFirst(c *C) {
	res, err := PrefixToPostfix("- * 3 5 / 7 6")
	c.Check(err, Equals, nil)
	c.Check(res, Equals, "3 5 * 7 6 / -")
}

func (s *MySuite) TestPrefixToPostfixSimpleSecond(c *C) {
	res, err := PrefixToPostfix("+ ^ 3 2 + 5 6")
	c.Check(err, Equals, nil)
	c.Check(res, Equals, "3 2 ^ 5 6 + +")
}

/* Complicated test cases */

func (s *MySuite) TestPrefixToPostfixComplicatedFirst(c *C) {
	res, err := PrefixToPostfix("- / * 4 8 22 - - ^ 8 4 1 + 3 * 54 3")
	c.Check(err, Equals, nil)
	c.Check(res, Equals, "4 8 * 22 / 8 4 ^ 1 - 3 54 3 * + - -")
}

func (s *MySuite) TestPrefixToPostfixComplicatedSecond(c *C) {
	res, err := PrefixToPostfix("/ 3 + + * * 3 1 4 3 - ^ 7 2 * 1 9")
	c.Check(err, Equals, nil)
	c.Check(res, Equals, "3 3 1 * 4 * 3 + 7 2 ^ 1 9 * - + /")
}

/* Error test cases */

func (s *MySuite) TestPrefixToPostfixErrorEmptyLine(c *C) {
	_, err := PrefixToPostfix("  ")
	c.Check(err.Error(), Equals, "Error: empty line entered!")
}

func (s *MySuite) TestPrefixToPostfixErrorInvalidOperand(c *C) {
	_, err := PrefixToPostfix("- * e 5 / 7 6")
	c.Check(err.Error(), Equals, "Error: operand is not valid!")
}

func (s *MySuite) TestPrefixToPostfixErrorFewOperands(c *C) {
	_, err := PrefixToPostfix("- 2")
	c.Check(err.Error(), Equals, "Error: not enough operands!")
}

func (s *MySuite) TestPrefixToPostfixErrorFewOperators(c *C) {
	_, err := PrefixToPostfix("- 2 7 5")
	c.Check(err.Error(), Equals, "Error: input is not valid")
}
