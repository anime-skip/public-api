package utils

func GraphQLError(message string) map[string]interface{} {
	return map[string]interface{}{
		"errors": []map[string]interface{}{
			{
				"message": "X-Client-ID header is required",
			},
		},
	}
}
