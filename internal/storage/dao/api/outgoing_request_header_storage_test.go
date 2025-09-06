package api_dao_test

import (
	"sort"
	"testing"

	db_model "github.com/HidenYT/api-shortener/internal/storage/db-model/api"
	"github.com/stretchr/testify/require"
)

func createRequestHeader(t *testing.T, configID uint) *db_model.OutgoingRequestHeader {
	header := &db_model.OutgoingRequestHeader{
		Name:                    "someHeader",
		Value:                   "someValue",
		OutgoingRequestConfigID: configID,
	}
	err := testOutgoingRequestHeaderDAO.Create(header)

	require.NoError(t, err)
	require.NotEmpty(t, header)

	require.NotZero(t, header.ID)
	require.NotZero(t, header.CreatedAt)

	return header
}

func assertOutgoingRequestHeadersEqual(t *testing.T, expected, actual *db_model.OutgoingRequestHeader) {
	require.Equal(t, expected.ID, actual.ID)
	require.Equal(t, expected.Name, actual.Name)
	require.Equal(t, expected.Value, actual.Value)
	require.Equal(t, expected.OutgoingRequestConfigID, actual.OutgoingRequestConfigID)
}

func TestCreateRequestHeader(t *testing.T) {
	configID := createRequestConfig(t, createShortenedAPI(t).ID).ID
	createRequestHeader(t, configID)
}

func TestGetRequestHeader(t *testing.T) {
	configID := createRequestConfig(t, createShortenedAPI(t).ID).ID
	header1 := createRequestHeader(t, configID)

	header2, err := testOutgoingRequestHeaderDAO.Get(header1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, header2)
	assertOutgoingRequestHeadersEqual(t, header1, header2)
}

func TestGetAllRequestHeadersByConfigID(t *testing.T) {
	configID := createRequestConfig(t, createShortenedAPI(t).ID).ID
	headers1 := []*db_model.OutgoingRequestHeader{
		createRequestHeader(t, configID),
		createRequestHeader(t, configID),
		createRequestHeader(t, configID),
	}

	headers2, err := testOutgoingRequestHeaderDAO.GetAllByConfigID(configID)
	require.NoError(t, err)
	require.Equal(t, len(headers1), len(headers2))
	sort.Slice(headers1, func(i, j int) bool { return headers1[i].ID < headers1[j].ID })
	sort.Slice(headers2, func(i, j int) bool { return headers2[i].ID < headers2[j].ID })
	for i := range headers1 {
		assertOutgoingRequestHeadersEqual(t, headers1[i], headers2[i])
	}
}

func TestUpdateRequestHeader(t *testing.T) {
	configID := createRequestConfig(t, createShortenedAPI(t).ID).ID
	header1 := createRequestHeader(t, configID)

	header1.Name = "SomeUnpredictableName"
	header1.Value = "SomeUnpredictableValue"

	err := testOutgoingRequestHeaderDAO.Update(header1)
	require.NoError(t, err)

	header2, err := testOutgoingRequestHeaderDAO.Get(header1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, header2)
	assertOutgoingRequestHeadersEqual(t, header1, header2)
}

func TestDeleteRequestHeader(t *testing.T) {
	configID := createRequestConfig(t, createShortenedAPI(t).ID).ID
	header1 := createRequestHeader(t, configID)

	err := testOutgoingRequestHeaderDAO.Delete(header1.ID)
	require.NoError(t, err)

	header2, err := testOutgoingRequestHeaderDAO.Get(header1.ID)
	require.Error(t, err)
	require.Empty(t, header2)
}
