/*
sevDesk API

Testing CheckAccountTransactionApiService

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

func Test_openapi_CheckAccountTransactionApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CheckAccountTransactionApiService CreateTransaction", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CheckAccountTransactionApi.CreateTransaction(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CheckAccountTransactionApiService DeleteCheckAccountTransaction", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var checkAccountTransactionId int32

		resp, httpRes, err := apiClient.CheckAccountTransactionApi.DeleteCheckAccountTransaction(context.Background(), checkAccountTransactionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CheckAccountTransactionApiService GetCheckAccountTransactionById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var checkAccountTransactionId int32

		resp, httpRes, err := apiClient.CheckAccountTransactionApi.GetCheckAccountTransactionById(context.Background(), checkAccountTransactionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CheckAccountTransactionApiService GetTransactions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CheckAccountTransactionApi.GetTransactions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CheckAccountTransactionApiService UpdateCheckAccountTransaction", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var checkAccountTransactionId int32

		resp, httpRes, err := apiClient.CheckAccountTransactionApi.UpdateCheckAccountTransaction(context.Background(), checkAccountTransactionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
