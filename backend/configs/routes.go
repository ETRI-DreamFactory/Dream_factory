package configs

import (
	"github.com/cryptosalamander/dream_factory/helpers"
	"github.com/cryptosalamander/dream_factory/models"
	"github.com/cryptosalamander/dream_factory/repositories"
	"github.com/cryptosalamander/dream_factory/services"
	"net/http"

	//"github.com/cryptosalamander/dream_factory/dtos"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(DreamRepository *repositories.DreamRepository) *gin.Engine {
	route := gin.Default()

	// create route /create endpoint

	route.POST("/create", func(context *gin.Context) {
		// initialization contact model
		var contact models.Contact

		//validate json
		err := context.ShouldBindJSON(&context)

		//validation errors
		if err != nil {
			// generate validation errors response
			response := helpers.GenerateValidationResponse(err)
			context.JSON(http.StatusBadRequest, response)
			return
		}

		// default http status code = 200
		code := http.StatusOK

		// save contact & get it's response
		response := services.CreateContact(&contact, *DreamRepository)

		// save contact failed
		if !response.Success {
			//change http status code to 400
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.POST("/createmember", func(context *gin.Context) {
		// initialization contact model
		var member models.Member

		//validate json
		err := context.ShouldBindJSON(&member)

		//validation errors
		if err != nil {
			// generate validation errors response
			response := helpers.GenerateValidationResponse(err)
			context.JSON(http.StatusBadRequest, response)
			return
		}

		// default http status code = 200
		code := http.StatusOK

		// save contact & get it's response
		response := services.CreateMember(&member, *DreamRepository)

		// save contact failed
		if !response.Success {
			//change http status code to 400
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.GET("/", func(context *gin.Context) {
		code := http.StatusOK
		response := services.FindAllContacts(*DreamRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.GET("/getmembers", func(context *gin.Context) {
		code := http.StatusOK
		response := services.FindAllMembers(*DreamRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.GET("/show/:id", func(context *gin.Context) {
		id := context.Param("id")

		code := http.StatusOK
		response := services.FindContactById(id, *DreamRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.GET("/showmember/:id", func(context *gin.Context) {
		id := context.Param("id")

		code := http.StatusOK
		response := services.FindMemberById(id, *DreamRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.PUT("/update/:id", func(context *gin.Context) {
		id := context.Param("id")
		var contact models.Contact
		err := context.ShouldBindJSON(&contact)

		// validation errors

		if err != nil {
			response := helpers.GenerateValidationResponse(err)
			context.JSON(http.StatusBadRequest, response)
			return
		}

		code := http.StatusOK

		response := services.UpdateContactById(id, &contact, *DreamRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.PUT("/updatemember/:id", func(context *gin.Context) {
		id := context.Param("id")
		var member models.Member
		err := context.ShouldBindJSON(&member)

		// validation errors

		if err != nil {
			response := helpers.GenerateValidationResponse(err)
			context.JSON(http.StatusBadRequest, response)
			return
		}

		code := http.StatusOK

		response := services.UpdateMemberById(id, &member, *DreamRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.DELETE("/delete/:id", func(context *gin.Context) {
		id := context.Param("id")
		code := http.StatusOK

		response := services.DeleteOneContactById(id, *DreamRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.DELETE("/deletemember/:id", func(context *gin.Context) {
		id := context.Param("id")
		code := http.StatusOK

		response := services.DeleteMemberById(id, *DreamRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.GET("/pagination", func(context *gin.Context) {
		code := http.StatusOK

		pagination := helpers.GeneratePaginationRequest(context)

		response := services.Pagination(*DreamRepository, context, pagination)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.GET("/memberpagination", func(context *gin.Context) {
		code := http.StatusOK

		pagination := helpers.GeneratePaginationRequest(context)

		response := services.MemberPagination(*DreamRepository, context, pagination)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	return route

}
