package resources

type JsonResponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type ModelResource struct {
}

func (r *ModelResource) mapResource(model interface{}) interface{} {
	return model
}

func (r *ModelResource) Make(model interface{}) JsonResponse {
	return JsonResponse{
		Data:    r.mapResource(model),
		Message: "Success",
	}
}

func (r *ModelResource) Collections(models []interface{}) JsonResponse {
	var resources []interface{}

	for _, model := range models {
		resources = append(resources, r.mapResource(model))
	}

	return JsonResponse{
		Data:    resources,
		Message: "Success",
	}
}
