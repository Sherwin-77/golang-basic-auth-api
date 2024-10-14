package resources

import "github.com/sherwin-77/golang-basic-auth-api/models"

type UserResource struct {
	ModelResource
}

func (r *UserResource) mapResource(model models.User) map[string]interface{} {
	return map[string]interface{}{
		"id":         model.ID,
		"username":   model.Username,
		"email":      model.Email,
		"created_at": model.CreatedAt,
		"updated_at": model.UpdatedAt,

		"roles": model.Roles,
		"todos": model.Todos,
	}
}

func (r *UserResource) Make(model models.User) JsonResponse {
	return JsonResponse{
		Data:    r.mapResource(model),
		Message: "Success",
	}
}

type UserIndexResource struct {
	ModelResource
}

func (r *UserIndexResource) Collections(models []models.User) JsonResponse {
	var resources []interface{}

	for _, model := range models {
		resources = append(resources, r.mapResource(model))
	}

	return JsonResponse{
		Data:    resources,
		Message: "Success",
	}
}
