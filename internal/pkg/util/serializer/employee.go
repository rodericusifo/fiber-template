package serializer

import (
	"github.com/rodericusifo/fiber-template/internal/app/core/employee/controller/api/response"
	"github.com/rodericusifo/fiber-template/internal/app/core/employee/service/dto/output"
	"github.com/rodericusifo/fiber-template/internal/app/model/database/sql"
	"github.com/rodericusifo/fiber-template/internal/pkg/constant"
)

func SerializeEmployeeToEmployeeDTO(model *sql.Employee) *output.EmployeeDTO {
	return &output.EmployeeDTO{
		XID:       model.XID,
		Name:      model.Name,
		Email:     model.Email,
		Address:   model.Address,
		Age:       model.Age,
		Birthday:  model.Birthday,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}

func SerializeEmployeesToEmployeeDTOs(models []*sql.Employee) []*output.EmployeeDTO {
	result := make([]*output.EmployeeDTO, 0)

	for _, model := range models {
		result = append(result, &output.EmployeeDTO{
			XID:       model.XID,
			Name:      model.Name,
			Email:     model.Email,
			Address:   model.Address,
			Age:       model.Age,
			Birthday:  model.Birthday,
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
		})
	}

	return result
}

func SerializeEmployeeDTOToEmployeeResponse(dto *output.EmployeeDTO) *response.EmployeeResponse {
	birthdayTime := *dto.Birthday
	birthdayString := birthdayTime.Format(constant.DEFAULT_TIME_LAYOUT.(string))
	createdAtTime := dto.CreatedAt
	createdAtString := createdAtTime.Format(constant.DEFAULT_TIME_LAYOUT.(string))
	updatedAtTime := dto.UpdatedAt
	updatedAtString := updatedAtTime.Format(constant.DEFAULT_TIME_LAYOUT.(string))

	return &response.EmployeeResponse{
		XID:       dto.XID,
		Name:      dto.Name,
		Email:     dto.Email,
		Address:   dto.Address,
		Age:       dto.Age,
		Birthday:  &birthdayString,
		CreatedAt: createdAtString,
		UpdatedAt: updatedAtString,
	}
}

func SerializeEmployeeDTOsToEmployeeResponses(dtos []*output.EmployeeDTO) []*response.EmployeeResponse {
	result := make([]*response.EmployeeResponse, 0)

	for _, dto := range dtos {
		birthdayTime := *dto.Birthday
		birthdayString := birthdayTime.Format(constant.DEFAULT_TIME_LAYOUT.(string))
		createdAtTime := dto.CreatedAt
		createdAtString := createdAtTime.Format(constant.DEFAULT_TIME_LAYOUT.(string))
		updatedAtTime := dto.UpdatedAt
		updatedAtString := updatedAtTime.Format(constant.DEFAULT_TIME_LAYOUT.(string))

		result = append(result, &response.EmployeeResponse{
			XID:       dto.XID,
			Name:      dto.Name,
			Email:     dto.Email,
			Address:   dto.Address,
			Age:       dto.Age,
			Birthday:  &birthdayString,
			CreatedAt: createdAtString,
			UpdatedAt: updatedAtString,
		})
	}

	return result
}
