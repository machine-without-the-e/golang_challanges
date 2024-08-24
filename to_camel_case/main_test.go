package main

import (
  "testing"
)

func Test_to_camel_case(t *testing.T) {
  tests := [][]string{
    {"the-stealth-warrior", "theStealthWarrior"},
    {"The_Stealth_Warrior", "TheStealthWarrior"},
    {"The_Stealth-Warrior", "TheStealthWarrior"},
  }

  for _, test := range tests {
    result := ToCamelCase(test[0]);

    if (result == test[1]) {
      continue
    }

    t.Errorf("Expected `%s` but got `%s` for %s",
      test[1],
      result,
      test[0],
    )
  }
}

