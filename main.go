package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	. "github.com/ellcrys/cocoon/core/stub"
	"github.com/ellcrys/util"
)

// TotalMegaCoin is the total amount of digital current to issue or pay employees
const TotalMegaCoin = "250000000"

// Ledger is the name of the dedicated ledger for this contract
const Ledger = "megaledger2"

// ToJSON returns JSON encoded representation of an object
func ToJSON(obj interface{}) (bs []byte) {
	bs, _ = util.ToJSON(obj)
	return
}

// Account defines a structure for representing an account
type Account struct {
	ID        string
	Balance   float64
	FirstName string
	LastName  string
}

// Employee represents a MegaCorp employee
type Employee struct {
	AccountID string
	Position  string
}

// MegaCorp is a service for managing MegaCorp Inc.
// employees and salary payment
type MegaCorp struct {
	maxEmployee    int
	positionSalary map[string]float64
	salaryPayDay   int
}

// OnInit handles contract initialization
func (m *MegaCorp) OnInit() error {

	// Initialize fields
	m.maxEmployee = 100
	m.salaryPayDay = 29
	m.positionSalary = map[string]float64{
		"ceo":          200000.00,
		"coo":          170000.00,
		"cto":          150000.00,
		"frontend_dev": 120000.00,
		"backend_dev":  120000.00,
	}

	defer m.cron()

	// create ledger
	_, err := Me.NewLedger(Ledger, true, true)
	if err != nil {
		if !strings.Contains(err.Error(), "already exists") {
			return fmt.Errorf("failed to create ledger: %s", err.Error())
		}
		return nil
	}

	// persist the current coin supply
	if _, err = Me.Put(Ledger, "total_megacoin", []byte(TotalMegaCoin)); err != nil {
		return fmt.Errorf("failed to set coin supply")
	}

	return nil
}

// OnInvoke handles function calls from external clients
func (m *MegaCorp) OnInvoke(md Metadata, function string, params []string) ([]byte, error) {
	switch function {

	// create-account function creates a new account
	case "create-account":

		if len(params) < 2 {
			return nil, fmt.Errorf("first name and last name are required")
		}

		newAcct, err := m.createAccount(params[0], params[1])
		if err != nil {
			return nil, fmt.Errorf("failed to create account")
		}
		return ToJSON(newAcct), nil

	// create-employee function creates a new employee.
	// requires account id and position passed in params.
	case "create-employee":
		if len(params) < 2 {
			return nil, fmt.Errorf("account id and position are required")
		}
		newEmp, err := m.createEmployee(params[0], params[1])
		if err != nil {
			return nil, fmt.Errorf("failed to create employee: %s", err)
		}
		return ToJSON(newEmp), nil

	// get-all-employees fetches all employees
	case "get-all-employees":
		return ToJSON(m.getAllEmployees()), nil

	// get-account fetches an account
	case "get-account":
		if len(params) < 1 {
			return nil, fmt.Errorf("account is required")
		}
		acct, err := m.getAccount(params[0])
		if err != nil {
			return nil, fmt.Errorf("failed to get account: %s", err)
		}
		return ToJSON(acct), nil

	case "get-all-accounts":
		return ToJSON(m.getAllAccounts()), nil

	// get-total-supply fetches the total supply of megacoin
	case "get-total-supply":
		return []byte(m.getRemainingCoinSupply()), nil

	default:
		return nil, fmt.Errorf("unsupported function")
	}
}

// OnStop is called when the contract is stopped.
// This is where you should free resources.
func (m *MegaCorp) OnStop() {
	return
}

// createAccount creates a new account
func (m *MegaCorp) createAccount(firstName, lastName string) (acct *Account, err error) {
	acct = &Account{ID: util.UUID4(), FirstName: firstName, LastName: lastName}
	_, err = Me.Put(Ledger, fmt.Sprintf("account.%s", acct.ID), ToJSON(acct))
	return
}

// createEmployee creates a new employee. Requires a valid MegaCorp account
func (m *MegaCorp) createEmployee(accountID, position string) (emp *Employee, err error) {

	// get account
	_, err = Me.Get(Ledger, fmt.Sprintf("account.%s", accountID))
	if err != nil {
		if err.Error() == "transaction not found" {
			err = fmt.Errorf("account not found")
		}
		return
	}

	// ensure the position is valid
	if _, ok := m.positionSalary[position]; !ok {
		err = fmt.Errorf("position is unrecognised")
		return
	}

	emp = &Employee{
		AccountID: accountID,
		Position:  position,
	}

	// create employee
	_, err = Me.Put(Ledger, fmt.Sprintf("employee.%s.%s", accountID, position), ToJSON(emp))

	return
}

// getAccount gets an account
func (m *MegaCorp) getAccount(id string) (acct *Account, err error) {
	tx, err := Me.Get(Ledger, fmt.Sprintf("account.%s", id))
	util.FromJSON([]byte(tx.Value), &acct)
	return
}

// getAllAccounts fetches all accounts
func (m *MegaCorp) getAllAccounts() (accts []*Account) {
	rg := Me.NewRangeGetter(Ledger, "account", "", false)
	for rg.HasNext() {
		var act Account
		util.FromJSON([]byte(rg.Next().Value), &act)
		accts = append(accts, &act)
	}
	return
}

// getAllEmployees fetches all employees
func (m *MegaCorp) getAllEmployees() (emps []*Employee) {
	rg := Me.NewRangeGetter(Ledger, "employee", "", false)
	for rg.HasNext() {
		var emp Employee
		util.FromJSON([]byte(rg.Next().Value), &emp)
		emps = append(emps, &emp)
	}
	return
}

// saveAccount saves an account
func (m *MegaCorp) saveAccount(acct *Account) (err error) {
	_, err = Me.Put(Ledger, fmt.Sprintf("account.%s", acct.ID), ToJSON(acct))
	return
}

// getRemainingCoinSupply gets the number of remaining coins
func (m *MegaCorp) getRemainingCoinSupply() string {
	tx, _ := Me.Get(Ledger, "total_megacoin")
	return tx.Value
}

// deductTotalMegaCoinSupply deducts from the total megacoin supply
func (m *MegaCorp) deductTotalMegaCoinSupply(amt float64) (err error) {
	tx, err := Me.Get(Ledger, "total_megacoin")
	curTotalSupply, _ := strconv.ParseFloat(tx.Value, 64)
	newTotalSupply := curTotalSupply - amt
	newTotalSupplyBytes := []byte(strconv.FormatFloat(newTotalSupply, 'f', -1, 64))
	_, err = Me.Put(Ledger, "total_megacoin", newTotalSupplyBytes)
	return
}

// cron runs background tasks such as paying salaries when due
func (m *MegaCorp) cron() {
	t := time.NewTicker(60 * time.Second)
	for _ = range t.C {

		fmt.Println("pay salary")

		if time.Now().Day() == m.salaryPayDay {

			// get all the employees
			rg := Me.NewRangeGetter(Ledger, "employee", "", false)
			for rg.HasNext() {
				var emp Employee
				util.FromJSON([]byte(rg.Next().Value), &emp)

				// get account
				acct, _ := m.getAccount(emp.AccountID)
				amountToPay := m.positionSalary[emp.Position]
				acct.Balance += amountToPay

				// reduce total mega coin
				m.deductTotalMegaCoinSupply(amountToPay)

				// save changes by overriding previous account
				m.saveAccount(acct)
			}
		}
	}
}

func main() {
	Run(&MegaCorp{})
}
