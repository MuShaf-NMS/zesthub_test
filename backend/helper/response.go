package helper

// Helper to generate error response
func ErrorResponseBuilder(msg string, errs interface{}) map[string]interface{} {
	res := map[string]interface{}{
		"status": "Failed",
		"msg":    msg,
		"err":    errs,
	}
	return res

}

// Helper to generate respons
func ResponseBuilder(msg string, data interface{}) map[string]interface{} {
	res := map[string]interface{}{
		"status": "OK",
		"msg":    msg,
		"data":   data,
	}
	return res
}
