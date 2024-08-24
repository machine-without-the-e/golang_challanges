package main

import (
  "strings"
  "strconv"
)

func Is_valid_ip(ip string) bool {
  ip_parts := strings.Split(ip, ".")

  if (len(ip_parts) != 4) {
    return false;
  }

  for _, part := range ip_parts {
    number, err := strconv.Atoi(part)

    if (err != nil) {
      return false;
    }
  
    if (number > 255 || number < 0) {
      return false;
    }
    
    // Filters out leading zeros by converting the int back to string and comparing with original
    if (strconv.Itoa(number) != part) {
      return false;
    }
  }

  return true;
}



func main() {

}

