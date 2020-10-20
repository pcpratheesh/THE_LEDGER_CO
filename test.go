package main

import (
	"fmt"
	"the_ledger_co/app"
	"the_ledger_co/app/models"
)

func main() {
	app := app.InitApp()

	emil__id := "3P05S"
	var pendingrecord []models.EmiPaymentDetailLedger

	pendingemiobjectInstance := app.DB
	pendingemiobjectInstance = pendingemiobjectInstance.Where("emi_id = ? AND payment_status = ? AND id < ?", emil__id, 0, 281)
	pendingemiobjectInstance = pendingemiobjectInstance.Order("id ASC")
	pendingemiobjectInstance.Find(&pendingrecord)

	fmt.Println(pendingrecord)

}
