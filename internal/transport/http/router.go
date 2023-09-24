package http

func (s *Server) SetUpRoute() {
	v1 := s.App.Group("/api/v1")

	user := v1.Group("/users")
	user.POST("/sign-up", s.handler.SignUp)
	user.POST("/sign-in", s.handler.SignIn)

	news := v1.Group("/news")
	news.GET("", s.handler.GetAllNews)

	survey := v1.Group("/survey")
	survey.GET("", s.handler.GetSurveyList)
	survey.GET("/:survey_id", s.handler.GetSurvey)
	survey.POST("", s.handler.SaveSurveyAnswer)

	photo := v1.Group("/photo")
	photo.POST("/upload", s.handler.UploadPhoto)
	photo.GET("/download", s.handler.DownloadPhoto)
}
