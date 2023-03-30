package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

const dsn = "root:root@(127.0.0.1:3306)/xueden_student_mangement_system?charset=utf8mb4&parseTime=True&loc=Local"

// generate code
func main() {
	// specify the output directory (default: "./query")
	// ### if you want to query without context constrain, set mode gen.WithoutContext ###
	g := gen.NewGenerator(gen.Config{
		OutPath:      "./dal/query",
		ModelPkgPath: "./dal/model",
		Mode:         gen.WithDefaultQuery | gen.WithQueryInterface,
		//if you want the nullable field generation property to be pointer type, set FieldNullable true
		FieldNullable: true,
		//generate pointer when field has default value
		FieldCoverable: true,
		//if you want to generate index tags from database, set FieldWithIndexTag true
		FieldWithIndexTag: true,
		//if you want to generate type tags from database, set FieldWithTypeTag true
		FieldWithTypeTag: true,
	})

	// reuse the database connection in Project or create a connection here
	// if you want to use GenerateModel/GenerateModelAs, UseDB is necessary, or it will panic
	db, _ := gorm.Open(mysql.Open(dsn))
	g.UseDB(db)

	// apply diy interfaces on structs or table models

	//user 表和 role 表是Many2Many的关系
	userModel := g.GenerateModelAs("sys_user", "User",
		gen.FieldRelate(field.HasOne, "Role", g.GenerateModelAs("sys_role", "Role"),
			&field.RelateConfig{GORMTag: "foreignKey:RoleID"}),
	)
	g.ApplyInterface(func(IUserMethod) {}, userModel)

	//role表
	g.ApplyInterface(func(IUserMethod) {}, g.GenerateModelAs("sys_role", "Role"))

	//course表
	g.ApplyInterface(func(IUserMethod) {}, g.GenerateModelAs("s_course", "Course"))

	//grade_class表
	g.ApplyInterface(func(IUserMethod) {}, g.GenerateModelAs("s_grade_class", "GradeClass"))

	//student表
	studentModel := g.GenerateModelAs("s_student", "Student",
		gen.FieldRelate(field.HasOne, "GradeClass", g.GenerateModelAs("s_grade_class", "GradeClass"),
			&field.RelateConfig{GORMTag: "foreignKey:GradeClassID"}),
	)
	g.ApplyInterface(func(IUserMethod) {}, studentModel)

	//student_score表
	studentScoreModel := g.GenerateModelAs("s_student_score", "StudentScore",
		gen.FieldRelate(field.HasOne, "GradeClass", g.GenerateModelAs("s_grade_class", "GradeClass"),
			&field.RelateConfig{GORMTag: "foreignKey:GradeClassID"}),
		gen.FieldRelate(field.HasOne, "Student", g.GenerateModelAs("s_student", "Student"),
			&field.RelateConfig{GORMTag: "foreignKey:StudentID"}),
		gen.FieldRelate(field.HasOne, "Course", g.GenerateModelAs("s_course", "Course"),
			&field.RelateConfig{GORMTag: "foreignKey:CourseID"}),
	)
	g.ApplyInterface(func(IUserMethod) {}, studentScoreModel)

	//teacher表
	teacherModel := g.GenerateModelAs("s_teacher", "Teacher",
		gen.FieldRelate(field.HasOne, "Course", g.GenerateModelAs("s_course", "Course"),
			&field.RelateConfig{GORMTag: "foreignKey:CourseID"}),
	)
	g.ApplyInterface(func(IUserMethod) {}, teacherModel)

	//g.GenerateAllTable()

	// execute the action of code generation
	g.Execute()
}

type IUserMethod interface {
}
