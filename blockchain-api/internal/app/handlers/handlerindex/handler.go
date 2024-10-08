package handlerindex

import (
	"github.com/cronnoss/blockchain-api/blockchain-api/internal/app/handlers/ctx"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

// IndexHandler is a callback handler.
type IndexHandler struct {
	indexService IndexService
}

// New creates a new index handler.
func New(indexService IndexService) *IndexHandler {
	return &IndexHandler{
		indexService: indexService,
	}
}

// Get handles GET /indexes/:id request.
func (h *IndexHandler) Get(c *fiber.Ctx) error {

	// swagger:route GET /indexes/:id indexes getIndex
	//
	// Lists contract index info.
	//
	// This will show contract index info.
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http
	//
	//     Responses:
	//       default: body:error
	//       200: body:index
	//       400: body:error

	indexIDStr := c.Params("id")
	indexID, err := strconv.ParseInt(indexIDStr, 10, 64)
	if err != nil {
		return ctx.Error(c, http.StatusBadRequest, err)
	}

	index, getIndexErr := h.indexService.GetIndex(c.Context(), indexID)
	if getIndexErr != nil {
		return ctx.Error(c, http.StatusInternalServerError, getIndexErr)
	}

	return c.JSON(index)
}
