package repository

import (
	"database/sql"
	"log"
	"market_place/collections"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

// references : https://github.com/Fadli2001/go-unit-test-testify/blob/master/repository/customer_db_repository_test.go

// TO DO :
// 1. why testify? and not golang testing?
// 2. why you should create suite?
// 3. mock -> must using mocking, because not connect database

var dummyBankAccounts = []collections.BankAccountInput{
	{
		ID:                "1",
		UserID:            "1",
		BankName:          "AAA",
		BankAccountName:   "Account AAA",
		BankAccountNumber: "123123",
	},
	{
		ID:                "2",
		UserID:            "1",
		BankName:          "BBB",
		BankAccountName:   "Account BBB",
		BankAccountNumber: "123123",
	},
}

type BankAccountRepositoryTestSuite struct {
	suite.Suite
	mockDb  *sql.DB
	mockSql sqlmock.Sqlmock
}

func (suite *BankAccountRepositoryTestSuite) SetupTest() {
	mockDb, mockSql, err := sqlmock.New()
	if err != nil {
		log.Fatalln("An error when opening a database connection")
	}
	suite.mockDb = mockDb
	suite.mockSql = mockSql
}

func (suite *BankAccountRepositoryTestSuite) TearDownTest() {
	suite.mockDb.Close()
}

func (suite *BankAccountRepositoryTestSuite) TestBankAccountGetByIDSuccess() {
	ba := dummyBankAccounts[0]
	rows := sqlmock.NewRows([]string{"id", "name", "account_name", "account_number"}).AddRow(ba.ID, ba.BankName, ba.BankAccountName, ba.BankAccountNumber)

	suite.mockSql.ExpectQuery(regexp.QuoteMeta(`SELECT TEXT(id), name, account_name, account_number FROM bank_accounts WHERE id = $1 AND user_id = $2 AND deleted_at is null;`)).WillReturnRows(rows)

	repo := NewBankAccountRepository(suite.mockDb)

	actual, err := repo.GetByID(ba.ID, ba.UserID)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), actual)
}

func (suite *BankAccountRepositoryTestSuite) TestBankAccountCreateSuccess() {
	ba := dummyBankAccounts[0]
	suite.mockSql.ExpectExec(regexp.QuoteMeta(`INSERT INTO bank_accounts (name, account_name, account_number, user_id, created_at, updated_at) VALUES ($1, $2, $3, $4, current_timestamp, current_timestamp);`)).WithArgs("AAA", "Account AAA", "123123", "1").WillReturnResult(sqlmock.NewResult(1, 1))
	repo := NewBankAccountRepository(suite.mockDb)
	err := repo.Create(ba)
	assert.Nil(suite.T(), err)
}

func (suite *BankAccountRepositoryTestSuite) TestBankAccountListSuccess() {
	rows := sqlmock.NewRows([]string{"id", "name", "account_name", "account_number"})

	for _, data := range dummyBankAccounts {
		rows.AddRow(data.ID, data.BankName, data.BankAccountName, data.BankAccountNumber)
	}

	suite.mockSql.ExpectQuery(regexp.QuoteMeta(`SELECT TEXT(id), name, account_name, account_number FROM bank_accounts WHERE user_id = $1 AND deleted_at is null;`)).WillReturnRows(rows)
	repo := NewBankAccountRepository(suite.mockDb)

	data, err := repo.List("1")

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 2, len(data))
	assert.Equal(suite.T(), "AAA", data[0].BankName)
}

func TestBankAccountRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(BankAccountRepositoryTestSuite))
}
