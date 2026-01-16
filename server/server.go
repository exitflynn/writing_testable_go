package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/workshop/writing-testable-go/shortener"
)

type Server struct {
	store  shortener.URLStore
	router *gin.Engine
}

func New(store shortener.URLStore) *Server {
	s := &Server{store: store}
	s.setupRoutes()
	return s
}

func (s *Server) setupRoutes() {
	gin.SetMode(gin.ReleaseMode)
	s.router = gin.New()
	s.router.Use(gin.Recovery())

	s.router.Static("/static", "./server/static")
	s.router.GET("/", s.handleIndex)
	s.router.POST("/shorten", s.handleShorten)
	s.router.GET("/:code", s.handleRedirect)
}

func (s *Server) Run(addr string) error {
	return s.router.Run(addr)
}

func (s *Server) handleIndex(c *gin.Context) {
	c.File("./server/static/index.html")
}

type shortenRequest struct {
	URL string `json:"url"`
}

type shortenResponse struct {
	ShortCode string `json:"short_code"`
	ShortURL  string `json:"short_url"`
}

func (s *Server) handleShorten(c *gin.Context) {
	var req shortenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	mapping, err := shortener.CreateMapping(req.URL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := s.store.Save(mapping); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save"})
		return
	}

	c.JSON(http.StatusOK, shortenResponse{
		ShortCode: mapping.ShortCode,
		ShortURL:  "http://localhost:8080/" + mapping.ShortCode,
	})
}

func (s *Server) handleRedirect(c *gin.Context) {
	code := c.Param("code")

	mapping, err := s.store.Get(code)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	c.Redirect(http.StatusMovedPermanently, mapping.LongURL)
}
