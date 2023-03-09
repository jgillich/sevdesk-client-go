/*
sevDesk API

Testing AccountingContactApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_AccountingContactApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AccountingContactApiService CreateAccountingContact", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AccountingContactApi.CreateAccountingContact(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountingContactApiService DeleteAccountingContact", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountingContactId int32

		resp, httpRes, err := apiClient.AccountingContactApi.DeleteAccountingContact(context.Background(), accountingContactId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountingContactApiService GetAccountingContact", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AccountingContactApi.GetAccountingContact(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountingContactApiService GetAccountingContactById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountingContactId int32

		resp, httpRes, err := apiClient.AccountingContactApi.GetAccountingContactById(context.Background(), accountingContactId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountingContactApiService UpdateAccountingContact", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountingContactId int32

		resp, httpRes, err := apiClient.AccountingContactApi.UpdateAccountingContact(context.Background(), accountingContactId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}