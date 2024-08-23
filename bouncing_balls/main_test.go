package main

import (
  "testing"
)

type testData struct {
  expected int
  values []float64
}

func testBouncingBall(t *testing.T, expected []testData) {


  for _, test_data := range expected {
    result := BouncingBall(test_data.values[0],
      test_data.values[1],
      test_data.values[2],
    );

    if (result != test_data.expected) {
      t.Fatalf(`Data of (%f, %f, %f) should return %d, but got %d`,
        test_data.values[0],
        test_data.values[1],
        test_data.values[2],
        test_data.expected,
        result,
      );
    }
  }
}

func Test_height_must_be_greater_than_0(t *testing.T) {
  expected := []testData{
    testData{expected: -1, values: []float64{0,0,0}},
    testData{expected: -1, values: []float64{-44,0,0}},
  }

  testBouncingBall(t, expected);
}

func Test_bounce_must_be_greater_than_0_and_less_than_1(t *testing.T) {
  
  expected := []testData{
    testData{expected: -1, values: []float64{1,1,0}},
    testData{expected: -1, values: []float64{1,0,0}},
    testData{expected: -1, values: []float64{1,44,0}},
  }

  testBouncingBall(t, expected);

}

func Test_window_must_be_less_than_height(t *testing.T) {
  
  expected := []testData{
    testData{expected: -1, values: []float64{1,0.66,55}},
    testData{expected: -1, values: []float64{1,0.66,1}},
    testData{expected: -1, values: []float64{4,0.66,4}},
  }

  testBouncingBall(t, expected);

}


func Test_bouncing(t *testing.T) {
  
  expected := []testData{
    testData{expected: 3, values: []float64{3,0.66,1.5}},

    testData{expected: 3, values: []float64{40,0.4,10}},
  }

  testBouncingBall(t, expected);
}
