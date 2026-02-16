package main

import (
	"net/http"
	"strconv"

	"github.com/bbombelka/contactbook/types"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func validateRequestBody(addUserBody types.AddUserRequestBody, c *gin.Context) bool {

	if addUserBody.PhoneNumber == nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "phoneNumber field is missing!"})
		return false
	}
	if addUserBody.Name == nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "name field is missing!"})
		return false
	}
	if addUserBody.Address == nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "address field is missing!"})
		return false
	}
	if addUserBody.Username == nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "username field is missing!"})
		return false
	}

	return true
}

func extractUserId(c *gin.Context) string {
	userId := c.Param("id")

	return userId
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getUpdateFields(updateUserData types.UpdateUserRequestBody) bson.M {
	updatedFields := bson.M{}

	if updateUserData.Name != nil {
		updatedFields["name"] = updateUserData.Name
	}

	if updateUserData.Username != nil {
		updatedFields["username"] = updateUserData.Username
	}

	if updateUserData.Address != nil {
		updatedFields["address"] = updateUserData.Address
	}

	if updateUserData.PhoneNumber != nil {
		updatedFields["phoneNumber"] = updateUserData.PhoneNumber
	}

	return updatedFields
}

func extractQueryParams(c *gin.Context) (int64, int64) {
	limit := c.Query("limit")
	skip := c.Query("skip")

	if limit == "" {
		limit = "10"
	}

	if skip == "" {
		skip = "0"
	}

	intLimit, _ := strconv.ParseInt(limit, 10, 64)
	intSkip, _ := strconv.ParseInt(skip, 10, 64)

	return intLimit, intSkip
}
