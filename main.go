package main

import (
	"fmt"
	"os"
	"strconv"
	"the_ledger_co/app/bank"
)

func main__() {
	// loan := bank.InitLoan()
	// loan.BankName = "MDB"
	// loan.BorrowerName = "Dale"
	// loan.PrincipalAmount = 10000
	// loan.NoOfYears = 1
	// loan.RateOfInterest = 2

	// _, _ = loan.BorrowLoan()

	// fetch abalance
	bl := bank.InitBalance()
	bl.BankName = "MDB"
	bl.BorrowerName = "Dale"
	bl.EmiNumber = 4

	response := bl.Balance()

	if response.Status == true {
		fmt.Println(response.BankName, response.BorrowerName, response.AmountPaid, response.NoOfEmiLeft)
	} else {
		fmt.Printf("\033[1;31m%s\033[0m \n", response.Error)
	}

	// payment
	// payment := bank.InitPayment()
	// payment.BankName = "MDB"
	// payment.BorrowerName = "Dale"
	// payment.EmiNumber = 0
	// payment.LumpSumAmount = 3000

	// response_payment := payment.Payment()
	// fmt.Println(response_payment)
}
func main() {

	err := ProcessCommands()
	if err != nil {
		fmt.Printf("\033[1;31m[error]\033[0m %s \n", err)
	}
}

/**
 * -------------------------------------------------------------------
 * Routing the functions based on the input comands
 * -------------------------------------------------------------------
 */
func ProcessCommands() error {

	allowedCommands := map[string]string{
		"LOAN":    "LOAN",
		"PAYMENT": "PAYMENT",
		"BALANCE": "BALANCE",
	}

	if len(os.Args) > 1 {

		//check comnand exist allowed commands
		if _, ok := allowedCommands[os.Args[1]]; !ok {
			return fmt.Errorf("Invalid command")
		}

		switch os.Args[1] {
		case "LOAN":
			return _process_loan()
		case "PAYMENT":
			return _process_payment()

		case "BALANCE":
			return _process_balance()
		}
	} else {

	}

	return nil
}

/**
 * -------------------------------------------------------------------
 * Process LOAN
 * -------------------------------------------------------------------
 */
func _process_loan() error {
	if len(os.Args) < 7 {
		return fmt.Errorf("insufficient parameters")
	}

	loan := bank.InitLoan()
	loan.BankName = os.Args[2]
	loan.BorrowerName = os.Args[3]

	val, _ := strconv.ParseFloat(os.Args[4], 64)
	loan.PrincipalAmount = val

	valI, _ := strconv.Atoi(os.Args[5])
	loan.NoOfYears = valI

	valI, _ = strconv.Atoi(os.Args[6])
	loan.RateOfInterest = valI

	err, loan_data := loan.BorrowLoan()

	if err != nil {
		fmt.Printf("\033[1;31m%s\033[0m \n", err)
	} else {
		fmt.Printf("\033[1;32m%s\033[0m \n", " Your Loan has been approved")

		fmt.Println("-------------- Loan Details -------------------------")
		fmt.Println("BankName : ", loan_data.BankName)
		fmt.Println("BorrowerName : ", loan_data.BorrowerName)
		fmt.Println("PrincipalAmount : ", loan_data.PrincipalAmount)
		fmt.Println("NoOfYears : ", loan_data.NoOfYears)
		fmt.Println("RateOfInterest : ", loan_data.RateOfInterest)

		fmt.Println("-------------- EMI Details  -------------------------")
		fmt.Println("Total Repay : ", loan_data.TotalRepay)
		fmt.Println("Total Emis : ", loan_data.NoOfEmis)
		fmt.Println("Monthly Payment : ", loan_data.MonthlyEmi)
	}
	return nil
}

/**
 * -------------------------------------------------------------------
 * Process Payments
 * -------------------------------------------------------------------
 */
func _process_payment() error {
	if len(os.Args) < 6 {
		return fmt.Errorf("insufficient parameters")
	}

	payment := bank.InitPayment()
	payment.BankName = os.Args[2]
	payment.BorrowerName = os.Args[3]

	valF, _ := strconv.ParseFloat(os.Args[4], 64)
	payment.LumpSumAmount = valF

	valI, _ := strconv.Atoi(os.Args[5])
	payment.EmiNumber = valI

	response_payment := payment.Payment()
	if response_payment.Status == true {
	} else {
		fmt.Printf("\033[1;31m%s\033[0m \n", response_payment.Error)
	}
	return nil
}

/**
 * -------------------------------------------------------------------
 * Process Balance
 * -------------------------------------------------------------------
 */
func _process_balance() error {
	if len(os.Args) < 5 {
		return fmt.Errorf("insufficient parameters")
	}

	bl := bank.InitBalance()
	bl.BankName = os.Args[2]
	bl.BorrowerName = os.Args[3]

	valI, _ := strconv.Atoi(os.Args[4])
	bl.EmiNumber = valI

	response := bl.Balance()

	if response.Status == true {
		fmt.Println(response.BankName, response.BorrowerName, response.AmountPaid, response.NoOfEmiLeft)
	} else {
		fmt.Printf("\033[1;31m%s\033[0m \n", response.Error)
	}

	return nil
}

// go run main.go LOAN IDIDI bank 1000 2 2
// go run main.go PAYMENT IDIDI bank 1000 5
