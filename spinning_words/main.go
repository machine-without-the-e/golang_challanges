package main


import (
  "strings"
)

func spinWords(words string) string {

  //Split words by space character
  words_split := strings.Split(words, " ")

  //For each word in words we want to see the length 
  for index, word := range words_split {
    if(len(word) < 5) {
      continue;
    }

    words_split[index] = flipWord(word)
  }
  // Join the words, and return words with flipped
  return strings.Join(words_split, " ")
}

func flipWord(word string) string {
  word_split := strings.Split(word, "")

  reversed := ""
  for _, char := range word_split {
    reversed = char + reversed
  }
  return reversed
}

func main() {

}
