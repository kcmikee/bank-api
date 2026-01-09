package sqlc

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/kcmikee/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createTestAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomInt(0, 10000),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomInt(0, 10000),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createTestAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := createTestAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt.Time, account2.CreatedAt.Time, 1e9)
}

func TestUpdateAccount(t *testing.T) {
	account1 := createTestAccount(t)

	arg := UpdateAccountParams{
		ID:      account1.ID,
		Balance: util.RandomInt(0, 2000),
	}

	account2, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, arg.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt.Time, account2.CreatedAt.Time, 1e9)
}

func TestDeleteAccount(t *testing.T) {
	account1 := createTestAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.Error(t, err)
	require.Equal(t, pgx.ErrNoRows, err)
	require.Empty(t, account2)
}

func TestListAccount(t *testing.T) {
	// Create accounts with the same owner
	owner := util.RandomOwner()
	for i := 0; i < 10; i++ {
		arg := CreateAccountParams{
			Owner:    owner,
			Balance:  util.RandomInt(0, 10000),
			Currency: util.RandomCurrency(),
		}

		account, err := testQueries.CreateAccount(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, account)
	}

	listArg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), listArg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
		// require.Equal(t, owner, account.Owner)
	}
}
