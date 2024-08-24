package main

import (
  "testing"
)

type testcase struct {
  expected bool
  ip string
}

func Test_Is_valid_ip(t *testing.T) {
  tests := []testcase{
    testcase{
      expected: true,
      ip: "1.2.3.4",
    },
    testcase{
      expected: true,
      ip: "123.45.67.89",
    },
    testcase{
      expected: false,
      ip: "1.2.3",
    },
    testcase{
      expected: false,
      ip: "1.2.3.4.5",
    },
    testcase{
      expected: false,
      ip: "123.456.78.90",
    },
    testcase{
      expected: false,
      ip: "123.045.067.089",
    },
  }

  for _,test := range tests {
    result := Is_valid_ip(test.ip)

    if (result == test.expected) {
      continue;
    }

    t.Errorf("Fail: `%t` returned but expected `%t` for input `%s`",
      result,
      test.expected,
      test.ip,
    )
  }
}
