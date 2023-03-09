/*
sevDesk API

Testing VoucherApiService

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

func Test_openapi_VoucherApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test VoucherApiService BookVoucher", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var voucherId int32

		resp, httpRes, err := apiClient.VoucherApi.BookVoucher(context.Background(), voucherId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VoucherApiService CreateVoucherByFactory", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.VoucherApi.CreateVoucherByFactory(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VoucherApiService GetVoucherById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var voucherId int32

		resp, httpRes, err := apiClient.VoucherApi.GetVoucherById(context.Background(), voucherId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VoucherApiService GetVouchers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.VoucherApi.GetVouchers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VoucherApiService UpdateVoucher", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var voucherId int32

		resp, httpRes, err := apiClient.VoucherApi.UpdateVoucher(context.Background(), voucherId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VoucherApiService VoucherUploadFile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.VoucherApi.VoucherUploadFile(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}