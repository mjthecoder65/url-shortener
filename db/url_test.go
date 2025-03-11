package db

import (
	"context"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/mjthecoder65/url-shortener/utils"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
)

func createRandomShortURL(t *testing.T) URL {
	arg := CreateShortURLParams{
		URL: faker.URL(),
	}

	shourtUrl, err := testQueries.CreateShortURL(context.Background(), testConfig, arg)
	require.NoError(t, err)
	require.NotEmpty(t, shourtUrl)
	require.Equal(t, arg.URL, arg.URL)
	require.Equal(t, len(shourtUrl.ShortCode), testConfig.ShortCodeLength)
	require.NotEmpty(t, shourtUrl.ID)

	return shourtUrl
}

func TestCreateShortURL(t *testing.T) {
	createRandomShortURL(t)
}

func TestGetShortUrl(t *testing.T) {
	shortUrl := createRandomShortURL(t)

	result, err := testQueries.GetShortURL(context.Background(), shortUrl.ShortCode)
	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Equal(t, result.ShortCode, shortUrl.ShortCode)
	require.Equal(t, shortUrl.URL, result.URL)

	shortCode := utils.GenerateShortCode(testConfig)
	require.Equal(t, len(shortCode), testConfig.ShortCodeLength)

	result, err = testQueries.GetShortURL(context.Background(), shortCode)
	require.Error(t, err)
	require.Empty(t, result)
	require.Equal(t, err, mongo.ErrNoDocuments)

}

func TestUpdateShortUrl(t *testing.T) {
	shortUrl := createRandomShortURL(t)

	arg := UpdateShortURLParams{
		ShortCode: shortUrl.ShortCode,
		URL:       faker.URL(),
	}

	result, err := testQueries.UpdateShortURL(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Equal(t, result.URL, arg.URL)
	require.Equal(t, shortUrl.ShortCode, arg.ShortCode)
}

func TestDeleteShortUrl(t *testing.T) {
	shortUrl := createRandomShortURL(t)
	err := testQueries.DeleteShortURL(context.Background(), shortUrl.ShortCode)
	require.NoError(t, err)

	err = testQueries.DeleteShortURL(context.Background(), shortUrl.ShortCode)
	require.Error(t, err)
	require.Equal(t, err, mongo.ErrNoDocuments)
}
