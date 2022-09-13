package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"scraping_challenge/services/cometco-scraper/api/converter"
	"scraping_challenge/services/cometco-scraper/api/model"
)

type CometServer struct {
	cometService CometService
	logger       *zap.Logger
}

func NewCometServer(
	cometService CometService,
	logger *zap.Logger,
) *CometServer {
	return &CometServer{
		cometService: cometService,
		logger:       logger.With(zap.String("server", "comet")),
	}
}

func (s *CometServer) Register(router *gin.Engine) {
	cometRouter := router.Group("/comet")
	tasksRoutes := cometRouter.Group("/tasks")
	{
		tasksRoutes.POST("/", s.CreateTask)
		tasksRoutes.GET("/", s.GetTasks)
		tasksRoutes.GET("/:id", s.GetTaskByID)
	}
	freelancerProfileRoutes := cometRouter.Group("/freelancer-profile")
	{
		freelancerProfileRoutes.GET("/:id", s.GetUserProfile)
	}
}

func (s *CometServer) CreateTask(c *gin.Context) {
	ctx := c.Request.Context()
	req := model.CreateTaskRequest{}
	if c.ShouldBindJSON(&req) != nil {
		s.errorResponse(c, model.NewErrorResponse(http.StatusBadRequest, errors.New("invalid request")))
		return
	}

	taskID, err := s.cometService.CreateTask(ctx, req.Login, req.Password)
	if err != nil {
		s.errorResponse(c, converter.FromError(err))
		return
	}

	c.JSON(http.StatusOK, model.CreateTaskResponse{
		ID: taskID,
	})
}

func (s *CometServer) GetTaskByID(c *gin.Context) {
	resp, err := s.cometService.GetTask(c.Request.Context(), c.Param("id"))
	if err != nil {
		s.errorResponse(c, converter.FromError(err))
		return
	}

	c.JSON(http.StatusOK, converter.ToGetTaskResponse(resp))
}

func (s *CometServer) GetTasks(c *gin.Context) {
	query := model.GetTaskListQuery{}
	if c.ShouldBind(&query) != nil {
		s.errorResponse(c, model.NewErrorResponse(http.StatusBadRequest, errors.New("invalid request")))
		return
	}

	resp, err := s.cometService.GetTasks(c.Request.Context(), query.Limit, query.Offset)
	if err != nil {
		s.errorResponse(c, converter.FromError(err))
		return
	}

	c.JSON(http.StatusOK, converter.ToGetTasksResponse(resp))
}

func (s *CometServer) GetUserProfile(c *gin.Context) {
	resp, err := s.cometService.GetUserProfile(c.Request.Context(), c.Param("id"))
	if err != nil {
		s.errorResponse(c, converter.FromError(err))
		return
	}

	c.JSON(http.StatusOK, converter.ToCometFreelancerProfile(resp))
}

func (s *CometServer) errorResponse(c *gin.Context, errResp model.ErrorResponse) {
	c.AbortWithStatusJSON(errResp.Code, errResp)
}
