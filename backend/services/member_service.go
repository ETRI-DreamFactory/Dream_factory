package services

import (
	"fmt"
	"github.com/cryptosalamander/dream_factory/dtos"
	"github.com/cryptosalamander/dream_factory/models"
	"github.com/cryptosalamander/dream_factory/repositories"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
)

func CreateMember(member *models.Member, repository repositories.DreamRepository) dtos.Response {

	uuidResult, err := uuid.NewRandom()

	if err != nil {
		log.Fatalln(err)
	}

	member.Member_id = uuidResult.String()

	operationResult := repository.SaveMember(member)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.Member)

	return dtos.Response{Success: true, Data: data}
}

func FindAllMembers(repository repositories.DreamRepository) dtos.Response {
	operationResult := repository.FindAllMembers()

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var datas = operationResult.Result.(*models.Members)
	return dtos.Response{Success: true, Data: datas}
}

func FindMemberById(id string, repository repositories.DreamRepository) dtos.Response {
	operationResult := repository.FindMemberById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.Member)

	return dtos.Response{Success: true, Data: data}
}

func UpdateMemberById(id string, member *models.Member, repository repositories.DreamRepository) dtos.Response {
	existingContactResponse := FindMemberById(id, repository)

	if !existingContactResponse.Success {
		return existingContactResponse
	}

	existingMember := existingContactResponse.Data.(*models.Member)

	existingMember.Passwd = member.Passwd
	existingMember.Name = member.Name
	existingMember.Phone = member.Phone
	existingMember.Wallet = member.Wallet
	existingMember.Nickname = member.Nickname
	operationResult := repository.SaveMember(existingMember)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true, Data: operationResult.Result}
}

func DeleteMemberById(id string, repository repositories.DreamRepository) dtos.Response {
	operationResult := repository.DeleteMemberById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true}

}

func MemberPagination(repository repositories.DreamRepository, context *gin.Context, pagination *dtos.Pagination) dtos.Response {
	operationResult, totalPages := repository.MemberPagination(pagination)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*dtos.Pagination)

	// get current url path
	urlPath := context.Request.URL.Path

	// search query params
	searchQueryParams := ""

	for _, search := range pagination.Searchs {
		searchQueryParams += fmt.Sprintf("&%s.%s=%s", search.Column, search.Action, search.Query)
	}

	// set first & last page pagination response
	data.FirstPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, 0, pagination.Sort) + searchQueryParams
	data.LastPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, totalPages, pagination.Sort) + searchQueryParams

	if data.Page > 0 {
		// set previous page pagination response
		data.PreviousPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, data.Page-1, pagination.Sort) + searchQueryParams
	}

	if data.Page < totalPages {
		// set next page pagination response
		data.NextPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, data.Page+1, pagination.Sort) + searchQueryParams
	}

	if data.Page > totalPages {
		// reset previous page
		data.PreviousPage = ""
	}

	return dtos.Response{Success: true, Data: data}
}
