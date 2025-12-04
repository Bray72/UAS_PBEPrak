package handler

import (
	"clean-arch/app/model"
	"clean-arch/app/service"
	"clean-arch/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type AchievementHandler struct {
	service service.AchievementService
}

func NewAchievementHandler(svc service.AchievementService) *AchievementHandler {
	return &AchievementHandler{service: svc}
}

// GetAchievements godoc
// @Summary List user's achievements
// @Description Get list of achievements for logged in user
// @Tags Achievements
// @Security Bearer
// @Success 200 {object} model.ListAchievementResponse
// @Failure 401 {object} map[string]interface{}
// @Router /api/v1/achievements [get]
func (h *AchievementHandler) GetAchievements(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)

	achievements, err := h.service.GetAchievementsByUserID(userID)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch achievements", nil)
	}

	return utils.SuccessResponse(c, http.StatusOK, "Achievements retrieved successfully", achievements)
}

// GetAchievementDetail godoc
// @Summary Get achievement detail
// @Description Get detail of specific achievement by ID
// @Tags Achievements
// @Security Bearer
// @Param id path string true "Achievement ID"
// @Success 200 {object} model.DetailAchievementResponse
// @Failure 401 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /api/v1/achievements/{id} [get]
func (h *AchievementHandler) GetAchievementDetail(c *fiber.Ctx) error {
	id := c.Params("id")
	userID := c.Locals("user_id").(string)

	achievement, err := h.service.GetAchievementByID(id)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusNotFound, "Achievement not found", nil)
	}

	// Check ownership
	if achievement.UserID != userID {
		return utils.ErrorResponse(c, http.StatusForbidden, "Unauthorized", nil)
	}

	return utils.SuccessResponse(c, http.StatusOK, "Achievement retrieved successfully", achievement)
}

// CreateAchievement godoc
// @Summary Create new achievement
// @Description Create new achievement (stored in MongoDB with PostgreSQL reference)
// @Tags Achievements
// @Security Bearer
// @Accept json
// @Produce json
// @Param request body model.CreateAchievementRequest true "Achievement data"
// @Success 201 {object} model.DetailAchievementResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /api/v1/achievements [post]
func (h *AchievementHandler) CreateAchievement(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)

	var req model.CreateAchievementRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", nil)
	}

	achievement, err := h.service.CreateAchievement(userID, &req)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	return utils.SuccessResponse(c, http.StatusCreated, "Achievement created successfully", achievement)
}

// UpdateAchievement godoc
// @Summary Update achievement
// @Description Update achievement (only draft status can be updated)
// @Tags Achievements
// @Security Bearer
// @Accept json
// @Produce json
// @Param id path string true "Achievement ID"
// @Param request body model.UpdateAchievementRequest true "Achievement data"
// @Success 200 {object} model.DetailAchievementResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /api/v1/achievements/{id} [put]
func (h *AchievementHandler) UpdateAchievement(c *fiber.Ctx) error {
	id := c.Params("id")
	userID := c.Locals("user_id").(string)

	var req model.UpdateAchievementRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", nil)
	}

	achievement, err := h.service.UpdateAchievement(id, userID, &req)
	if err != nil {
		if err.Error() == "unauthorized" {
			return utils.ErrorResponse(c, http.StatusForbidden, "Unauthorized", nil)
		}
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	return utils.SuccessResponse(c, http.StatusOK, "Achievement updated successfully", achievement)
}

// DeleteAchievement godoc
// @Summary Delete achievement
// @Description Delete achievement (only draft status can be deleted)
// @Tags Achievements
// @Security Bearer
// @Param id path string true "Achievement ID"
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /api/v1/achievements/{id} [delete]
func (h *AchievementHandler) DeleteAchievement(c *fiber.Ctx) error {
	id := c.Params("id")
	userID := c.Locals("user_id").(string)

	err := h.service.DeleteAchievement(id, userID)
	if err != nil {
		if err.Error() == "unauthorized" {
			return utils.ErrorResponse(c, http.StatusForbidden, "Unauthorized", nil)
		}
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	return utils.SuccessResponse(c, http.StatusOK, "Achievement deleted successfully", nil)
}

// SubmitAchievement godoc
// @Summary Submit achievement for verification
// @Description Submit achievement for verification (change status from draft to submitted)
// @Tags Achievements
// @Security Bearer
// @Accept json
// @Produce json
// @Param id path string true "Achievement ID"
// @Param request body model.SubmitAchievementRequest true "Submit data"
// @Success 200 {object} model.SubmitAchievementResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /api/v1/achievements/{id}/submit [post]
func (h *AchievementHandler) SubmitAchievement(c *fiber.Ctx) error {
	id := c.Params("id")
	userID := c.Locals("user_id").(string)

	var req model.SubmitAchievementRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", nil)
	}

	achievement, err := h.service.SubmitAchievement(id, userID, &req)
	if err != nil {
		if err.Error() == "unauthorized" {
			return utils.ErrorResponse(c, http.StatusForbidden, "Unauthorized", nil)
		}
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	return utils.SuccessResponse(c, http.StatusOK, "Achievement submitted successfully", achievement)
}
