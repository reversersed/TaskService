package rest

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/reversersed/taskservice/internal/application/interfaces"
	"github.com/reversersed/taskservice/internal/interface/api/rest/dto/mapper"
	"github.com/reversersed/taskservice/internal/interface/api/rest/dto/request"
	_ "github.com/reversersed/taskservice/internal/interface/api/rest/dto/response"
	"github.com/reversersed/taskservice/pkg/middleware"
)

type validator interface {
	StructValidation(any) error
}
type taskController struct {
	service   interfaces.TaskService
	validator validator
}

func NewTaskController(g *gin.Engine, service interfaces.TaskService, validator validator) *taskController {
	controller := &taskController{
		service:   service,
		validator: validator,
	}

	group := g.Group("/tasks")
	{
		group.POST("", controller.CreateTask)
		group.GET("", controller.GetAllTasks)
		group.GET("/:id", controller.GetTask)
		group.PUT("/:id", controller.UpdateTask)
		group.DELETE("/:id", controller.DeleteTask)
	}

	return controller
}

// @Summary Create new task
// @Tags tasks
// @Produce json
// @Param body body request.CreateTaskRequest true "Task request. Due field must be UTC time presented in format: yyyy-MM-ddThh:mm:ss"
// @Success 201 {object} response.TaskResponse
// @Failure 400 {object} middleware.customError "Received bad request"
// @Failure 500 {object} middleware.customError "Internal error occured"
// @Router /tasks [post]
func (t *taskController) CreateTask(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 1*time.Second)
	defer cancel()

	var request request.CreateTaskRequest
	if err := c.BindJSON(&request); err != nil {
		c.Error(middleware.BadRequestError(err.Error()))
		return
	}
	if err := t.validator.StructValidation(&request); err != nil {
		c.Error(err)
		return
	}
	command, err := request.Command()
	if err != nil {
		c.Error(middleware.BadRequestError(err.Error()))
		return
	}
	result, err := t.service.Create(ctx, command)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, mapper.ToTaskResponse(result.Result))
}

// @Summary Get all tasks
// @Tags tasks
// @Produce json
// @Success 200 {array} response.TaskResponse
// @Failure 500 {object} middleware.customError "Internal error occured"
// @Router /tasks [get]
func (t *taskController) GetAllTasks(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 1*time.Second)
	defer cancel()

	list, err := t.service.GetAll(ctx)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, mapper.ToTaskListResponse(list.Result))
}

// @Summary Get task by id
// @Tags tasks
// @Produce json
// @Param id path int true "Task Id"
// @Success 200 {object} response.TaskResponse
// @Failure 404 {object} middleware.customError "Task not found"
// @Failure 500 {object} middleware.customError "Internal error occured"
// @Router /tasks/{id} [get]
func (t *taskController) GetTask(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 1*time.Second)
	defer cancel()

	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.Error(middleware.BadRequestError("id param must be presented and must be valid integer value"))
		return
	}

	result, err := t.service.Get(ctx, id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, mapper.ToTaskResponse(result.Result))
}

// @Summary Update specified task by id
// @Tags tasks
// @Produce json
// @Param id path int true "Task Id"
// @Param body body request.UpdateTaskRequest true "Task body. Due field must be UTC time presented in format: yyyy-MM-ddThh:mm:ss"
// @Success 200 {object} response.TaskResponse
// @Failure 400 {object} middleware.customError "Received bad request"
// @Failure 404 {object} middleware.customError "Task not found"
// @Failure 500 {object} middleware.customError "Internal error occured"
// @Router /tasks/{id} [put]
func (t *taskController) UpdateTask(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 1*time.Second)
	defer cancel()

	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.Error(middleware.BadRequestError("id param must be presented and must be valid integer value"))
		return
	}

	var request request.UpdateTaskRequest
	if err := c.BindJSON(&request); err != nil {
		c.Error(middleware.BadRequestError(err.Error()))
		return
	}
	request.Id = id

	if err := t.validator.StructValidation(&request); err != nil {
		c.Error(err)
		return
	}
	command, err := request.Command()
	if err != nil {
		c.Error(middleware.BadRequestError(err.Error()))
		return
	}
	result, err := t.service.Update(ctx, command)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, mapper.ToTaskResponse(result.Result))
}

// @Summary Delete task by id
// @Tags tasks
// @Param id path int true "Task Id"
// @Success 204
// @Failure 404 {object} middleware.customError "Task not found"
// @Failure 500 {object} middleware.customError "Internal error occured"
// @Router /tasks/{id} [delete]
func (t *taskController) DeleteTask(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 1*time.Second)
	defer cancel()

	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.Error(middleware.BadRequestError("id param must be presented and must be valid integer value"))
		return
	}

	err = t.service.Delete(ctx, id)
	if err != nil {
		c.Error(err)
		return
	}
	c.Status(http.StatusNoContent)
}
