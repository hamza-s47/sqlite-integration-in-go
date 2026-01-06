package lead

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hamza-s47/sqlite-integration-in-go/database"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   string `gorm:"default:'Not provided'" json:"phone"`
}

func GetLeads(c *gin.Context) {
	db := database.DBConn
	var leads []Lead
	if err := db.Find(&leads).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch leads"})
		return
	}
	c.JSON(http.StatusOK, leads)
}

func GetLead(c *gin.Context) {
	id := c.Param("id")
	db := database.DBConn
	var lead Lead
	if err := db.Find(&lead, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Lead not found"})
		return
	}
	c.JSON(http.StatusOK, lead)
}

func NewLead(c *gin.Context) {
	db := database.DBConn
	lead := new(Lead)

	if err := c.ShouldBindJSON(lead); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	if err := db.Create(lead).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create lead",
		})
		return
	}

	c.JSON(http.StatusCreated, lead)
}

func DeleteLead(c *gin.Context) {
	id := c.Param("id")
	db := database.DBConn

	var lead Lead
	if err := db.Where("ID=?", id).Delete(&lead).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete lead"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Message": "Lead deleted successfully!", "status": "200"})
}
