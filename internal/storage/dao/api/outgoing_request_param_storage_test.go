package api_dao_test

import (
	"sort"
	"testing"

	db_model "github.com/HidenYT/api-shortener/internal/storage/db-model/api"
	"github.com/stretchr/testify/require"
)

func createRequestParam(t *testing.T, configID uint) *db_model.OutgoingRequestParam {
	param := &db_model.OutgoingRequestParam{
		Name:                    "someParam",
		Value:                   "someValue",
		OutgoingRequestConfigID: configID,
	}
	err := testOutgoingRequestParamDAO.Create(param)

	require.NoError(t, err)
	require.NotEmpty(t, param)

	require.NotZero(t, param.ID)
	require.NotZero(t, param.CreatedAt)

	return param
}

func assertOutgoingRequestParamsEqual(t *testing.T, expected, actual *db_model.OutgoingRequestParam) {
	require.Equal(t, expected.ID, actual.ID)
	require.Equal(t, expected.Name, actual.Name)
	require.Equal(t, expected.Value, actual.Value)
	require.Equal(t, expected.OutgoingRequestConfigID, actual.OutgoingRequestConfigID)
}

func TestCreateRequestParam(t *testing.T) {
	configID := createRequestConfig(t, createShortenedAPI(t).ID).ID
	createRequestParam(t, configID)
}

func TestGetRequestParam(t *testing.T) {
	configID := createRequestConfig(t, createShortenedAPI(t).ID).ID
	param1 := createRequestParam(t, configID)

	param2, err := testOutgoingRequestParamDAO.Get(param1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, param2)
	assertOutgoingRequestParamsEqual(t, param1, param2)
}

func TestGetAllRequestParamsByConfigID(t *testing.T) {
	configID := createRequestConfig(t, createShortenedAPI(t).ID).ID
	params1 := []*db_model.OutgoingRequestParam{
		createRequestParam(t, configID),
		createRequestParam(t, configID),
		createRequestParam(t, configID),
	}

	params2, err := testOutgoingRequestParamDAO.GetAllByConfigID(configID)
	require.NoError(t, err)
	require.Equal(t, len(params1), len(params2))
	sort.Slice(params1, func(i, j int) bool { return params1[i].ID < params1[j].ID })
	sort.Slice(params2, func(i, j int) bool { return params2[i].ID < params2[j].ID })
	for i := range params1 {
		assertOutgoingRequestParamsEqual(t, params1[i], params2[i])
	}
}

func TestUpdateRequestParam(t *testing.T) {
	configID := createRequestConfig(t, createShortenedAPI(t).ID).ID
	param1 := createRequestParam(t, configID)

	param1.Name = "SomeUnpredictableName"
	param1.Value = "SomeUnpredictableValue"

	err := testOutgoingRequestParamDAO.Update(param1)
	require.NoError(t, err)

	param2, err := testOutgoingRequestParamDAO.Get(param1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, param2)
	assertOutgoingRequestParamsEqual(t, param1, param2)
}

func TestDeleteRequestParam(t *testing.T) {
	configID := createRequestConfig(t, createShortenedAPI(t).ID).ID
	param1 := createRequestParam(t, configID)

	err := testOutgoingRequestParamDAO.Delete(param1.ID)
	require.NoError(t, err)

	param2, err := testOutgoingRequestParamDAO.Get(param1.ID)
	require.Error(t, err)
	require.Empty(t, param2)
}
