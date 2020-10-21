package main

import (
	"fmt"
	"os"
	"strconv"
	"the_ledger_co/app/bank"
)

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
		return fmt.Errorf("Should provide commands")
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

	valF, _ := strconv.ParseFloat(os.Args[6], 64)
	loan.RateOfInterest = valF

	err, loan_data := loan.BorrowLoan()

	if err != nil {
		fmt.Printf("\033[1;31m%s\033[0m \n", err)
	} else {
		fmt.Printf("\033[1;32m%s\033[0m \n", " Your Loan has been approved")

		fmt.Println("-------------- Loan Details -------------------------")
		fmt.Println("Bank Name : ", loan_data.BankName)
		fmt.Println("Borrower Name : ", loan_data.BorrowerName)
		fmt.Println("Principal Amount : ", loan_data.PrincipalAmount)
		fmt.Println("No Of Years : ", loan_data.NoOfYears)
		fmt.Println("Rate Of Interest : ", loan_data.RateOfInterest)

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

// go run main.go LOAN IDIDI bank 1000 1 2
// go run main.go LOAN IDIDI Dale 10000 1 1.2
// go run main.go PAYMENT IDIDI bank 1000 5
// go run main.go BALANCE IDIDI bank 0
