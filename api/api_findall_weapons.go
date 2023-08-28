package api

import (
  "CwIndustries/dao"
  "context"
  "github.com/aws/aws-lambda-go/events"
  "net/http"

  
  "encoding/json"
)


// @Summary      GET Weapons input: Weapons
// @Description  GET Weapons API
// @Tags         GET Weapons - Weapons
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.Model_Weapons "Status OK"
// @Success      202  {array}   dto.Model_Weapons "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /FindallWeapons [GET]


    func FindallWeaponsApi(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse,error) {


inputObj, err := dao.DB_FindallWeapons()
    if err != nil {
        return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest, Body: err.Error()},nil
    }


responseBodyJSON, _ := json.Marshal(inputObj)

    return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(responseBodyJSON),
	},nil

}
