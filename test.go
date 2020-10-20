package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// app := app.InitApp()

	// emil__id := "3P05S"
	// var pendingrecord []models.EmiPaymentDetailLedger

	// pendingemiobjectInstance := app.DB
	// pendingemiobjectInstance = pendingemiobjectInstance.Where("emi_id = ? AND payment_status = ? AND id < ?", emil__id, 0, 281)
	// pendingemiobjectInstance = pendingemiobjectInstance.Order("id ASC")
	// pendingemiobjectInstance.Find(&pendingrecord)

	// fmt.Println(pendingrecord)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)

}
