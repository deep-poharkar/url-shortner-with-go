type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc func(*gin.Context)
}