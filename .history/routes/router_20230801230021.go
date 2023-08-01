type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc func(*gin.Context)
}

type routes struct {
	router *gin.Engine
}

type Routes []R	oute
