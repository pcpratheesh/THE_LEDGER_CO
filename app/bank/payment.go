package bank

import (
	"fmt"
	"strconv"
	"the_ledger_co/app"
	"the_ledger_co/app/models"

	"gorm.io/gorm"
)

type PaymentRequest struct {
	BankName      string
	BorrowerName  string
	LumpSumAmount float64
	EmiNumber     int
}

type PaymentResponse struct {
	Status bool
	Error  error
}

func InitPayment() *PaymentRequest {
	models.AutoMigrateModel()
	return &PaymentRequest{}
}

/**
 * ---------------------------------------------------------------------------
 * Should process an emi payment
 *
 * Steps
 *		1 	-	Check the loan details exists
 * 		2 	-	if yes, get the unique emi id
 * 		3	-	Get emi next unpaid emi detail
 *		4	-	If added a lumpsum amount, extra amount should deduct from last emi
 * Condition check :
 * 				Payment amount should be greater than emi amount
 * 				Cannot pay 5th emi without paying the 4 emi
 *
 *
 * ---------------------------------------------------------------------------
 */
func (payment *PaymentRequest) Payment() PaymentResponse {
	//get the record of borrower detail
	result := map[string]interface{}{}
	var PaymentRespoObjec PaymentResponse

	app := app.InitApp()

	objectInstance := app.DB
	objectInstance = objectInstance.Model(&models.LoanDetailsLedger{})
	objectInstance = objectInstance.Where("bank_name = ? AND borrower_name = ?", payment.BankName, payment.BorrowerName)
	objectInstance = objectInstance.Find(&result)

	if len(result) > 0 {
		emil__id := result["emi_id"]

		// get the emi detail after the give emi no
		var record models.EmiPaymentDetailLedger
		var lumpsumAdded int

		emiobjectInstance := app.DB
		emiobjectInstance = emiobjectInstance.Where("emi_id = ?", emil__id)
		emiobjectInstance = emiobjectInstance.Limit(1)
		emiobjectInstance = emiobjectInstance.Offset(payment.EmiNumber - 1)
		emiobjectInstance = emiobjectInstance.Order("id ASC")
		emiobjectInstance.Find(&record)

		// check the payment already paid
		if record.PaymentStatus == 1 {
			PaymentRespoObjec.Status = false
			PaymentRespoObjec.Error = fmt.Errorf("This EMI has been paid already")
			return PaymentRespoObjec
		}

		// check there are pending emi payment for previous month
		var pendingrecord []models.EmiPaymentDetailLedger

		pendingemiobjectInstance := app.DB
		pendingemiobjectInstance = pendingemiobjectInstance.Where("emi_id = ? AND payment_status = ? AND id < ?", emil__id, 0, record.ID)
		pendingemiobjectInstance = pendingemiobjectInstance.Order("id ASC")
		pendingemiobjectInstance.Find(&pendingrecord)

		if len(pendingrecord) > 0 {
			PaymentRespoObjec.Status = false
			PaymentRespoObjec.Error = fmt.Errorf("There exists an unpaid emi of previous month. Please pay")
			return PaymentRespoObjec
		}

		// check lup sum amount greater than or equals
		if payment.LumpSumAmount < record.EmiAmount {
			PaymentRespoObjec.Status = false
			PaymentRespoObjec.Error = fmt.Errorf("Emi payment amount should be greater than or equals %v", record.EmiAmount)
			return PaymentRespoObjec

		}

		// check lump sum amound added
		if payment.LumpSumAmount > record.EmiAmount {
			lumpsumAdded = 1
		}

		// update the emi payment as received
		app.DB.Model(&models.EmiPaymentDetailLedger{}).Where("id = ?", record.ID).Updates(models.EmiPaymentDetailLedger{
			EmiPaidAmount: payment.LumpSumAmount,
			PaymentStatus: 1,
			AddedLumpsum:  lumpsumAdded,
		})

		if lumpsumAdded == 1 {
			extra__lumpsum := payment.LumpSumAmount - record.EmiAmount
			var lumpsum_record models.EmiPaymentDetailLedger
			var loop_updated_ids []string
			loop_updated_ids = append(loop_updated_ids, "0")

			for {
				lumpsum_record.ID = 0 // un initiallize

				emiobjectInstance := app.DB
				emiobjectInstance = emiobjectInstance.Where("emi_id = ?", emil__id, 0)
				emiobjectInstance = emiobjectInstance.Not(map[string]interface{}{"id": loop_updated_ids})
				emiobjectInstance = emiobjectInstance.Limit(1)
				emiobjectInstance = emiobjectInstance.Order("id DESC")
				emiobjectInstance.Find(&lumpsum_record)

				loop_updated_ids = append(loop_updated_ids, strconv.Itoa(int(lumpsum_record.ID)))

				if lumpsum_record.ID > 0 {

					if lumpsum_record.EmiAmount <= extra__lumpsum {
						app.DB.Model(&models.EmiPaymentDetailLedger{}).Where("id = ?", lumpsum_record.ID).Updates(map[string]interface{}{
							"payment_status": 1,
							"added_lumpsum":  1,
						})
						extra__lumpsum = extra__lumpsum - lumpsum_record.EmiAmount

						if extra__lumpsum <= 0 {
							break
						}
					} else {
						app.DB.Model(&models.EmiPaymentDetailLedger{}).Where("id = ?", lumpsum_record.ID).Updates(map[string]interface{}{
							"emi_amount":    gorm.Expr("emi_amount - ?", extra__lumpsum),
							"added_lumpsum": 1,
						})

						break

					}
				} else {
					break
				}

			}
		}

		PaymentRespoObjec.Status = true

	} else {
		PaymentRespoObjec.Status = false
		PaymentRespoObjec.Error = fmt.Errorf("Unable to find loan details of user with bank name")
	}

	return PaymentRespoObjec
}
