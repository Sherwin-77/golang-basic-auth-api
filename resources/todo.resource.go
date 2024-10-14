package resources

import "github.com/sherwin-77/golang-basic-auth-api/models"

type TodoResource struct {
	ModelResource
}

func (r *TodoResource) Collections(models []models.Todo) JsonResponse {
	var resources []interface{}

	for _, model := range models {
		resources = append(resources, r.mapResource(model))
	}

	return JsonResponse{
		Data:    resources,
		Message: "Success",
	}
}
