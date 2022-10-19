/*
 * ecodata - OpenAPI 3.0
 *
 * API to store contamination events and data from different sources
 *
 * API version: 1.0.0
 * Contact: nestor.salvador@gdl.cinvestav.mx
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateUser - Create user
func CreateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// CreateUsersWithListInput - Creates list of users with given input array
func CreateUsersWithListInput(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteUser - Delete user
func DeleteUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetUserByName - Get user by user name
func GetUserByName(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// LoginUser - Logs user into the system
func LoginUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// LogoutUser - Logs out current logged in user session
func LogoutUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// UpdateUser - Update user
func UpdateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
