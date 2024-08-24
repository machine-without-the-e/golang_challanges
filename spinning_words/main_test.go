package main

import (
  "testing"
)

func Test_spinWords(t *testing.T) {
  result := spinWords("Hey fellow warriors");
  

  if (result != "Hey wollef sroirraw") {
    t.Errorf("spinwords expected `Hey wollef sroirraw` but got %s", result)
  }

 
  result = spinWords("This is a test");
  

  if (result != "This is a test") {
    t.Errorf("spinwords expected `This is a test` but got %s", result)
  }

  result = spinWords("This is another test");
  

  if (result != "This is rehtona test") {
    t.Errorf("spinwords expected `This is rehtona test` but got %s", result)
  }


  result = spinWords("please");
  

  if (result != "esaelp") {
    t.Errorf("spinwords expected `esaelp` but got %s", result)
  }


  result = spinWords("look");
  

  if (result != "look") {
    t.Errorf("spinwords expected `look` but got %s", result)
  }


  result = spinWords("itchy");
  

  if (result != "yhcti") {
    t.Errorf("spinwords expected `yhcti` but got %s", result)
  }
}

func Test_flipWord(t *testing.T) {
  result := flipWord("hello!")

  if (result != "!olleh") {
    t.Errorf("flipword expected `!olleh` but got %s", result)
  }
}
