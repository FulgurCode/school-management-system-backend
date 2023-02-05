package adminHelpers

import (
	"github.com/FulgurCode/school-management-system/helpers"
	"github.com/FulgurCode/school-management-system/helpers/databaseHelpers"
	"github.com/FulgurCode/school-management-system/models"
)

// Create id and add teacher to database
func AddTeacher(teacher models.Teacher) error {
	// Adding unique id to teacher user
	var id = helpers.GetUniqueId()
	teacher.Id = id

	// Adding Teacher user to database
	var err = databaseHelpers.AddTeacher(teacher)
	return err
}
