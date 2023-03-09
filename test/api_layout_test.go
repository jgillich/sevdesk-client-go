/*
sevDesk API

Testing LayoutApiService

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

func Test_openapi_LayoutApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test LayoutApiService GetLetterpapersWithThumb", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LayoutApi.GetLetterpapersWithThumb(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LayoutApiService GetTemplates", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LayoutApi.GetTemplates(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LayoutApiService UpdateCreditNoteTemplate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var creditNoteId int32

		resp, httpRes, err := apiClient.LayoutApi.UpdateCreditNoteTemplate(context.Background(), creditNoteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LayoutApiService UpdateInvoiceTemplate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var invoiceId int32

		resp, httpRes, err := apiClient.LayoutApi.UpdateInvoiceTemplate(context.Background(), invoiceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LayoutApiService UpdateOrderTemplate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orderId int32

		resp, httpRes, err := apiClient.LayoutApi.UpdateOrderTemplate(context.Background(), orderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
