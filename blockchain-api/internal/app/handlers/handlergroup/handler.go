package handlergroup

import (
	"github.com/cronnoss/blockchain-api/blockchain-api/internal/app/handlers/ctx"
	"github.com/cronnoss/blockchain-api/blockchain-api/internal/pkg/models"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

// GroupHandler is a callback handler.
type GroupHandler struct {
	groupService GroupService
}

// New creates a new group handler.
func New(groupService GroupService) *GroupHandler {
	return &GroupHandler{
		groupService: groupService,
	}
}

// GetAll handles GET /groups request.
func (h *GroupHandler) GetAll(c *fiber.Ctx) error {

	// swagger:route GET /groups groups getGroups
	//
	// Lists all contract groups.
	//
	// This will show all available contract groups.
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http
	//
	//     Responses:
	//       default: body:error
	//       200: body:groupIDs
	//       400: body:error

	groups, err := h.groupService.GetGroupIDs(c.Context())
	if err != nil {
		return ctx.Error(c, http.StatusInternalServerError, err)
	}

	return c.JSON(models.Groups{GroupIDs: groups})
}

// Get handles GET /groups/:id request.
func (h *GroupHandler) Get(c *fiber.Ctx) error {

	// swagger:route GET /groups/:id groups getGroup
	//
	// Returns contract group info.
	//
	// This will show name and indexes for a smart contract group with provided ID (if any).
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http
	//
	//     Responses:
	//       default: body:error
	//       200: body:group
	//       400: body:error

	groupIDStr := c.Params("id")
	groupID, err := strconv.ParseInt(groupIDStr, 10, 64)
	if err != nil {
		return ctx.Error(c, http.StatusBadRequest, err)
	}

	group, getGroupErr := h.groupService.GetGroup(c.Context(), groupID)
	if getGroupErr != nil {
		return ctx.Error(c, http.StatusInternalServerError, getGroupErr)
	}

	return c.JSON(group)
}
