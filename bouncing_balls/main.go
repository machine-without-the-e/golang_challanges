package main

func BouncingBall(height, bounce, window float64) int {

  if (height <= 0) {
    return -1;
  }

  if (bounce <= 0 || bounce >= 1) {
    return -1;
  }

  if (window >= height) {
    return -1;
  }


  times_seen := 0;
  for height > window {
    // Dropping
    times_seen++;

    // Bouncing
    height = height * bounce;

    // Bounced and passed window
    if(height > window) {
      times_seen++;
    }
  }


  return times_seen;
  
}

func main() {
} 
