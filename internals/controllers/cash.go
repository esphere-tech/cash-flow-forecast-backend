package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/waltertaya/cash-flow-forecast-backend/internals/db"
	"github.com/waltertaya/cash-flow-forecast-backend/internals/models"
)

func GetEntries(c *gin.Context) {
	userID := c.GetString("user_id")
	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	var entries []models.CashEntry
	if err := db.DB.Where("user_id = ?", parsedUserID).Find(&entries).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve entries"})
		return
	}

	c.JSON(200, entries)
}

func CreateEntry(c *gin.Context) {
	userID := c.GetString("user_id")
	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	var input struct {
		Type        string  `json:"type" binding:"required,oneof=inflow outflow"`
		Amount      float64 `json:"amount" binding:"required"`
		Category    string  `json:"category"`
		Description string  `json:"description"`
		Date        string  `json:"date" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	entry := models.CashEntry{
		UserID:      parsedUserID,
		Type:        input.Type,
		Amount:      input.Amount,
		Category:    input.Category,
		Description: input.Description,
		Date:        input.Date,
	}

	if err := db.DB.Create(&entry).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create entry"})
		return
	}

	c.JSON(201, entry)
}

func CreateEntries(c *gin.Context) {
	userID := c.GetString("user_id")
	parsedUserID, err := uuid.Parse(userID)

	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	var input []struct {
		Type        string  `json:"type" binding:"required,oneof=inflow outflow"`
		Amount      float64 `json:"amount" binding:"required"`
		Category    string  `json:"category"`
		Description string  `json:"description"`
		Date        string  `json:"date" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var entries []models.CashEntry
	for _, item := range input {
		entry := models.CashEntry{
			UserID:      parsedUserID,
			Type:        item.Type,
			Amount:      item.Amount,
			Category:    item.Category,
			Description: item.Description,
			Date:        item.Date,
		}
		entries = append(entries, entry)
	}

	if err := db.DB.Create(&entries).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create entries"})
		return
	}

	c.JSON(201, entries)
}

func UpdateEntry(c *gin.Context) {
	userID := c.GetString("user_id")
	parsedUserID, err := uuid.Parse(userID)

	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	entryID := c.Param("id")
	parsedEntryID, err := uuid.Parse(entryID)

	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid entry ID"})
		return
	}

	var input struct {
		Type        string  `json:"type" binding:"required,oneof=inflow outflow"`
		Amount      float64 `json:"amount" binding:"required"`
		Category    string  `json:"category"`
		Description string  `json:"description"`
		Date        string  `json:"date" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var entry models.CashEntry
	if err := db.DB.Where("id = ? AND user_id = ?", parsedEntryID, parsedUserID).First(&entry).Error; err != nil {
		c.JSON(404, gin.H{"error": "Entry not found"})
		return
	}

	entry.Type = input.Type
	entry.Amount = input.Amount
	entry.Category = input.Category
	entry.Description = input.Description
	entry.Date = input.Date

	if err := db.DB.Save(&entry).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update entry"})
		return
	}

	c.JSON(200, entry)
}

func DeleteEntry(c *gin.Context) {
	userID := c.GetString("user_id")
	parsedUserID, err := uuid.Parse(userID)

	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	entryID := c.Param("id")
	parsedEntryID, err := uuid.Parse(entryID)

	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid entry ID"})
		return
	}

	if err := db.DB.Where("id = ? AND user_id = ?", parsedEntryID, parsedUserID).Delete(&models.CashEntry{}).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete entry"})
		return
	}

	c.JSON(200, gin.H{"message": "Entry deleted successfully"})
}
