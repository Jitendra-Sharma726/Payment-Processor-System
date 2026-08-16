package main

import (
  "encoding/csv"
  "fmt"
  "os"
  "strconv"

  // 1. Define the PaymentMethod interface 
  type PaymentMethod interface {
    Pay(amount float64)  (receipt string, fee float64, err error)
  }

  //2. Define the structs
  type CreditCard struct {
    CardNumber string
  }

  type Paypal struct {
    Email string
  }

  // 3. Implement the Interface for CreditCard
  func (c*CreditCard) Pay(amount float64) (string, float64, error) {
    //Fee: 2.5%
    fee := amount * 0.025
    receipt := fmt.Sprintf("Paid $%.2f via CreditCard (%s)", amount, c.CardNumber)
    return receipt, fee, nil
  }

  //4. Implement the Interface for PayPal
  func (p *PayPal) Pay(amount float64) (string, float64, error) {
    //Fee: 1.5% + 0.30
    fee := (amount * 0.015) + 0.30
    recipt := fmt.Sprintf("Paid $%.2f via PayPal (%s)", amount, p.Email)
    return receipt, fee, nil

  //5.The Factory Function
  func GetPaymentMethod(methodType, details string) (PaymentMethod, error) {
    switch methodType {
      case "credit_card":
         return &CreditCard{CardNumber: details}, nil
      case "paypal":
         return &PayPal{Email: details}, nil
      default:
         return nil, fmt.Errorf("unknown payment method %q", methodType)
      }
}

  func main() {
    if len(os.Args) < 2 {
      fmt.Println("Usage: go run main.go <csv_file>")
      return
    }

    fileName := os.Args[1]

 //6. Open the CSV File
 file, err := os.Open(fileName)
 if err != nil {
       fmt.Printf("Error Opening file: %v\n", err)
       return
  }
  defer file.Close()


//7. Parse the CSV Data
reader := csv.NewReader(file)


//Tell the CSV reader to allow variable row lengths so we can handle malformed rows manually
reader.FieldsPerRecord = -1

records, err := reader.ReadAll()
if err!= nil {
  fmt.Printf("Error reading CSV: %v\n", err)
  return
}

var totalFees float64

fmt.Println("=== Processing Batch Transactions ===")
fmt.Println()

//8. Polymorphic Processing Loop
// Start at i=1 to skip the header row
for i := 1; i < len(records); i++ {
    row := records[i]

  // Ensure the row has excatly 4 columns
  if len(row) != 4 {
    fmt.Printf("[Row %d] Failed: malformed csv data\n", i)
    continue
  }

  txnID := row[0]
  amountStr := row[1]
  methodType := row[2]
  details := row[3]

  //9. Parse amount safely 
  amount, err := strconv.ParseFloat(amountStr, 64)
  if err != nil {
    fmt.Printf("[%s] Failed: invalid amount format\n", txnID)
    continue
  }

  // 10.Instantiate via Factory
  processor, err := GetPaymentMethod(methodType, details)
  if err != nil {
     fmt.Printf("[%s] Failed: %v\n", txnID, err)
    continue
  }

  //11.Execute Polymorphic Method
  receipt, fee, err := processor.Pay(amount)
  if err != nil {
    fmt.Printf("[%s] Failed: %v\n", txnID, err)
    continue
  }

  //12. Aggregate and Print
  totalFees += fee
  fmt.Printf("[%s] Success: %s | Fee: $%.2f\n", txnID, receipt, fee)
}

fmt.Println("\n=== End of Batch ===")
fmt.Printf("Total Gateway Fees Collected $%.2f\n", totalFees)
}
  



    















  




  

    









  

    
    








    

    











    





      















  
