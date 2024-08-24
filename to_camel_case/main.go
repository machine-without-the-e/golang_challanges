package main

import (
  "strings"
)

func main() {

}

// Takes dash or underscore delimited words and transforms them into camelCasing
// First word should only be capatalised if the original word was
// Each word (aside from the first) should always be capatalised
func ToCamelCase(s string) string {
  parts := strings.FieldsFunc(s, split)

  result := ""
  for index, part := range parts {
    if (index == 0) {
      result += part
      continue;
    }
    
    result += title(part);
  }

	return result
}

func title(s string) string {
  parts := strings.Split(s, "")
  
  parts[0] = strings.ToUpper(parts[0])

  return strings.Join(parts, "")
}

func split(r rune) bool {
  return r == '_' || r == '-'
}
