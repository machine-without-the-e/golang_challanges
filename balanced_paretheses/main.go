package main

func main() {

}

func BalancedParens(n int) []string {

  parens := []string{
    "(",
    ")",

    "(",
    ")",
  }
  perms := []string{}

  
  for parentIndex, paren := range parens {
    perms = append(perms, getPerms(paren, remove(parens, parentIndex)))
  }

  return perms
}


func remove(s []string, i int) []string {
    s[i] = s[len(s)-1]
    return s[:len(s)-1]
}

func getPerms(perm string, avaliableCharacters []string) []string {
  if(len(avaliableCharacters) == 1) {
    return avaliableCharacters[0]
  }

  for _, avalibleCharacter := range avaliableCharacters {
    return avalibleCharacter + getPerms(perm, avaliableCharacters[1:])
  }

  return perm
}

func CheckValidParens(paren string) {
  //do some processing to check validity
}
