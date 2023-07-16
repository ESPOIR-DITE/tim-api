package userDetailsDomain

import (
	"errors"
	"net/http"
	"time"
)

type AccountDetails struct {
	Id                      string `json:"id" gorm:"primaryKey"`
	BankId                  string `json:"user.home.controller.bank.repository" sql:"bank_id"`
	CompanyRegisteredNumber string `json:"company_registered_number" sql:"company_registered_number"`
	TaxNumber               string `json:"tax_number" sql:"tax_number"`
	CreatedAt               time.Time
	UpdatedAt               time.Time
}

func (AccountDetails) TableName() string {
	return "account_detail"
}

func (u AccountDetails) Bind(r *http.Request) error {
	if u.CompanyRegisteredNumber == "" && u.TaxNumber == "" {
		return errors.New("missing required fields")
	}
	return nil
}

type UserBank struct {
	Id         string `json:"id" gorm:"primaryKey"`
	BankType   string `json:"bank_type" sql:"bank_type"`
	BankName   string `json:"bank_name" sql:"bank_name"`
	BranchCode string `json:"branch_code" sql:"bank_code"`
	BankNumber string `json:"bank_number" sql:"bank_number"`
	CvcCode    string `json:"cvc_code" sql:"cvc_code"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (UserBank) TableName() string {
	return "user_bank"
}

func (u UserBank) Bind(r *http.Request) error {
	if u.BranchCode == "" && u.BankType == "" && u.BankNumber == "" {
		return errors.New("missing required fields")
	}
	return nil
}
