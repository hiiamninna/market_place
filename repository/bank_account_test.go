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

func TestBankAccountRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(BankAccountRepositoryTestSuite))
}
