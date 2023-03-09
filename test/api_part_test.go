/*
sevDesk API

Testing PartApiService

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

func Test_openapi_PartApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PartApiService CreatePart", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PartApi.CreatePart(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PartApiService GetPartById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var partId int32

		resp, httpRes, err := apiClient.PartApi.GetPartById(context.Background(), partId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PartApiService GetParts", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PartApi.GetParts(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PartApiService PartGetStock", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var partId int32

		resp, httpRes, err := apiClient.PartApi.PartGetStock(context.Background(), partId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PartApiService UpdatePart", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var partId int32

		resp, httpRes, err := apiClient.PartApi.UpdatePart(context.Background(), partId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
