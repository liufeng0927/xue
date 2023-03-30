package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"xuedengwang/core/ormx"
	"xuedengwang/dal/query"
	"xuedengwang/service/course/rpc/course"
	"xuedengwang/service/gateway/api/internal/config"
	"xuedengwang/service/gateway/api/internal/middleware"
	"xuedengwang/service/gradeclass/rpc/gradeclass"
	"xuedengwang/service/role/rpc/role"
	"xuedengwang/service/score/rpc/score"
	"xuedengwang/service/student/rpc/student"
	"xuedengwang/service/teacher/rpc/teacher"
	"xuedengwang/service/user/rpc/user"
)

type ServiceContext struct {
	Config                config.Config
	UserRpc               user.User
	SetUidToCtxMiddleware rest.Middleware
	JwtAuth               rest.Middleware
	Query                 *query.Query
	RoleRpc               role.Role
	CourseRpc             course.Course
	TeacherRpc            teacher.Teacher
	StudentRpc            student.Student
	GradeClassRpc         gradeclass.Gradeclass
	ScoresRpc             score.Score
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UserRpc:       user.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
		RoleRpc:       role.NewRole(zrpc.MustNewClient(c.RoleRpcConf)),
		CourseRpc:     course.NewCourse(zrpc.MustNewClient(c.CourseRpcConf)),
		TeacherRpc:    teacher.NewTeacher(zrpc.MustNewClient(c.TeacherRpcConf)),
		StudentRpc:    student.NewStudent(zrpc.MustNewClient(c.StudentRpcConf)),
		GradeClassRpc: gradeclass.NewGradeclass(zrpc.MustNewClient(c.GradeClassRpcConf)),
		ScoresRpc:     score.NewScore(zrpc.MustNewClient(c.ScoresRpcConf)),
		JwtAuth:       middleware.NewJwtAuthMiddleware(c.JwtAuth.AccessSecret).Handle,
		Query:         query.Use(ormx.NewMySQL(&c.DB)),
	}
}
