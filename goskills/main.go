package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "strconv"
) 

func getFloats(fileName string) ( []float64, error) {
  var numbers []float64
  file, err := os.Open(fileName)
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    number, err := strconv.ParseFloat(line, 64)
    if err != nil {
      log.Fatal(err)
    }
    numbers = append(numbers, number)
  }
  err = file.Close()
  if scanner.Err() != nil {
    log.Fatal(err)
  }
  return numbers, nil
}

func main() {
  numbers, err := getFloats("./floats.txt")
  if err != nil {
    log.Fatal(err)
  }
  var sum float64 = 0
  for _, number := range numbers {
    sum += number
  }
  average := sum/float64(len(numbers))
  fmt.Printf("Average is: %0.2f\n", average)
}
  
