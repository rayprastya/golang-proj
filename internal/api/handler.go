package api

import (
	"net/http"
	"project/shopifyx/internal/model"
	"project/shopifyx/internal/repository"

	"github.com/gin-gonic/gin"
)

type AlbumHandler struct {
	repo *repository.AlbumRepository
}

func NewAlbumHandler(repo *repository.AlbumRepository) *AlbumHandler {
	return &AlbumHandler{
		repo: repo,
	}
}
func (h *AlbumHandler) GetAlbums(c *gin.Context) {
	albums, err := h.repo.GetAlbums(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, albums)
}

func (h *AlbumHandler) PostAlbums(c *gin.Context) {
	var newAlbum model.Album
	if err := c.BindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.AddAlbum(c, newAlbum); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newAlbum)
}
