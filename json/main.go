// -----------------------------
// This is a sample code in Go:
// - Encoding Data to JSON format
// - Decoding Data from JSON format
// -----------------------------

package main

import (
  "fmt"
  "encoding/json"
)

type Data struct {
  // VaribleName, DataType, CustomName
  ID              int       `json:"id"`
  Username        string    `json:"username"`
  Email           string    `json:"email"`
}

func main() {
  // -----------------------------
  // Encode Data to JSON
  // -----------------------------

  // Fill in Dummy Data
  data := Data{ID: 1, Username: "dimaskiddo", Email: "dimas.restu@student.upi.edu"}

  // Encode Dummy Data in JSON format Using Marshal function
  // ret Will Return Encoded JSON Data in Byte format
  // err Will Return Any Error
  ret, err := json.Marshal(data)
  if err != nil {
    fmt.Println(err.Error())
  }

  // Print Converted Encoded JSON
  fmt.Println(string(ret))

  // Line Break
  fmt.Println("--------------------------------------------")

  // -----------------------------
  // Decode Data from JSON
  // -----------------------------

  // To Decode Data from JSON we will using Unmarshal function
  // and it need to store return value in some variable memory
  var extractData Data

  // Decode Encoded Dummy Data from JSON format Using Unmarshal function
  // err Will Return Any Error, Because err is reused then just use it 
  //     instead fill in as new variable (use = rather that :=)
  err = json.Unmarshal(ret, &extractData)
  if err != nil {
    fmt.Println(err.Error())
  }

  // Print Decoded Dummy Data from JSON
  fmt.Println(extractData)
}
