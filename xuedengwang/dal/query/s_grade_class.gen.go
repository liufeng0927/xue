// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"xuedengwang/dal/model"
)

func newGradeClass(db *gorm.DB, opts ...gen.DOOption) gradeClass {
	_gradeClass := gradeClass{}

	_gradeClass.gradeClassDo.UseDB(db, opts...)
	_gradeClass.gradeClassDo.UseModel(&model.GradeClass{})

	tableName := _gradeClass.gradeClassDo.TableName()
	_gradeClass.ALL = field.NewAsterisk(tableName)
	_gradeClass.ID = field.NewInt64(tableName, "id")
	_gradeClass.CreateBy = field.NewInt64(tableName, "create_by")
	_gradeClass.CreateTime = field.NewTime(tableName, "create_time")
	_gradeClass.Remarks = field.NewString(tableName, "remarks")
	_gradeClass.UpdateBy = field.NewInt64(tableName, "update_by")
	_gradeClass.UpdateTime = field.NewTime(tableName, "update_time")
	_gradeClass.Clazz = field.NewInt32(tableName, "clazz")
	_gradeClass.Code = field.NewString(tableName, "code")
	_gradeClass.Grade = field.NewInt32(tableName, "grade")
	_gradeClass.Name = field.NewString(tableName, "name")

	_gradeClass.fillFieldMap()

	return _gradeClass
}

type gradeClass struct {
	gradeClassDo gradeClassDo

	ALL        field.Asterisk
	ID         field.Int64
	CreateBy   field.Int64
	CreateTime field.Time // 创建时间
	Remarks    field.String
	UpdateBy   field.Int64
	UpdateTime field.Time // 更新时间
	Clazz      field.Int32
	Code       field.String
	Grade      field.Int32
	Name       field.String

	fieldMap map[string]field.Expr
}

func (g gradeClass) Table(newTableName string) *gradeClass {
	g.gradeClassDo.UseTable(newTableName)
	return g.updateTableName(newTableName)
}

func (g gradeClass) As(alias string) *gradeClass {
	g.gradeClassDo.DO = *(g.gradeClassDo.As(alias).(*gen.DO))
	return g.updateTableName(alias)
}

func (g *gradeClass) updateTableName(table string) *gradeClass {
	g.ALL = field.NewAsterisk(table)
	g.ID = field.NewInt64(table, "id")
	g.CreateBy = field.NewInt64(table, "create_by")
	g.CreateTime = field.NewTime(table, "create_time")
	g.Remarks = field.NewString(table, "remarks")
	g.UpdateBy = field.NewInt64(table, "update_by")
	g.UpdateTime = field.NewTime(table, "update_time")
	g.Clazz = field.NewInt32(table, "clazz")
	g.Code = field.NewString(table, "code")
	g.Grade = field.NewInt32(table, "grade")
	g.Name = field.NewString(table, "name")

	g.fillFieldMap()

	return g
}

func (g *gradeClass) WithContext(ctx context.Context) IGradeClassDo {
	return g.gradeClassDo.WithContext(ctx)
}

func (g gradeClass) TableName() string { return g.gradeClassDo.TableName() }

func (g gradeClass) Alias() string { return g.gradeClassDo.Alias() }

func (g *gradeClass) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := g.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (g *gradeClass) fillFieldMap() {
	g.fieldMap = make(map[string]field.Expr, 10)
	g.fieldMap["id"] = g.ID
	g.fieldMap["create_by"] = g.CreateBy
	g.fieldMap["create_time"] = g.CreateTime
	g.fieldMap["remarks"] = g.Remarks
	g.fieldMap["update_by"] = g.UpdateBy
	g.fieldMap["update_time"] = g.UpdateTime
	g.fieldMap["clazz"] = g.Clazz
	g.fieldMap["code"] = g.Code
	g.fieldMap["grade"] = g.Grade
	g.fieldMap["name"] = g.Name
}

func (g gradeClass) clone(db *gorm.DB) gradeClass {
	g.gradeClassDo.ReplaceConnPool(db.Statement.ConnPool)
	return g
}

func (g gradeClass) replaceDB(db *gorm.DB) gradeClass {
	g.gradeClassDo.ReplaceDB(db)
	return g
}

type gradeClassDo struct{ gen.DO }

type IGradeClassDo interface {
	gen.SubQuery
	Debug() IGradeClassDo
	WithContext(ctx context.Context) IGradeClassDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IGradeClassDo
	WriteDB() IGradeClassDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IGradeClassDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IGradeClassDo
	Not(conds ...gen.Condition) IGradeClassDo
	Or(conds ...gen.Condition) IGradeClassDo
	Select(conds ...field.Expr) IGradeClassDo
	Where(conds ...gen.Condition) IGradeClassDo
	Order(conds ...field.Expr) IGradeClassDo
	Distinct(cols ...field.Expr) IGradeClassDo
	Omit(cols ...field.Expr) IGradeClassDo
	Join(table schema.Tabler, on ...field.Expr) IGradeClassDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IGradeClassDo
	RightJoin(table schema.Tabler, on ...field.Expr) IGradeClassDo
	Group(cols ...field.Expr) IGradeClassDo
	Having(conds ...gen.Condition) IGradeClassDo
	Limit(limit int) IGradeClassDo
	Offset(offset int) IGradeClassDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IGradeClassDo
	Unscoped() IGradeClassDo
	Create(values ...*model.GradeClass) error
	CreateInBatches(values []*model.GradeClass, batchSize int) error
	Save(values ...*model.GradeClass) error
	First() (*model.GradeClass, error)
	Take() (*model.GradeClass, error)
	Last() (*model.GradeClass, error)
	Find() ([]*model.GradeClass, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.GradeClass, err error)
	FindInBatches(result *[]*model.GradeClass, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.GradeClass) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IGradeClassDo
	Assign(attrs ...field.AssignExpr) IGradeClassDo
	Joins(fields ...field.RelationField) IGradeClassDo
	Preload(fields ...field.RelationField) IGradeClassDo
	FirstOrInit() (*model.GradeClass, error)
	FirstOrCreate() (*model.GradeClass, error)
	FindByPage(offset int, limit int) (result []*model.GradeClass, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IGradeClassDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (g gradeClassDo) Debug() IGradeClassDo {
	return g.withDO(g.DO.Debug())
}

func (g gradeClassDo) WithContext(ctx context.Context) IGradeClassDo {
	return g.withDO(g.DO.WithContext(ctx))
}

func (g gradeClassDo) ReadDB() IGradeClassDo {
	return g.Clauses(dbresolver.Read)
}

func (g gradeClassDo) WriteDB() IGradeClassDo {
	return g.Clauses(dbresolver.Write)
}

func (g gradeClassDo) Session(config *gorm.Session) IGradeClassDo {
	return g.withDO(g.DO.Session(config))
}

func (g gradeClassDo) Clauses(conds ...clause.Expression) IGradeClassDo {
	return g.withDO(g.DO.Clauses(conds...))
}

func (g gradeClassDo) Returning(value interface{}, columns ...string) IGradeClassDo {
	return g.withDO(g.DO.Returning(value, columns...))
}

func (g gradeClassDo) Not(conds ...gen.Condition) IGradeClassDo {
	return g.withDO(g.DO.Not(conds...))
}

func (g gradeClassDo) Or(conds ...gen.Condition) IGradeClassDo {
	return g.withDO(g.DO.Or(conds...))
}

func (g gradeClassDo) Select(conds ...field.Expr) IGradeClassDo {
	return g.withDO(g.DO.Select(conds...))
}

func (g gradeClassDo) Where(conds ...gen.Condition) IGradeClassDo {
	return g.withDO(g.DO.Where(conds...))
}

func (g gradeClassDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IGradeClassDo {
	return g.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (g gradeClassDo) Order(conds ...field.Expr) IGradeClassDo {
	return g.withDO(g.DO.Order(conds...))
}

func (g gradeClassDo) Distinct(cols ...field.Expr) IGradeClassDo {
	return g.withDO(g.DO.Distinct(cols...))
}

func (g gradeClassDo) Omit(cols ...field.Expr) IGradeClassDo {
	return g.withDO(g.DO.Omit(cols...))
}

func (g gradeClassDo) Join(table schema.Tabler, on ...field.Expr) IGradeClassDo {
	return g.withDO(g.DO.Join(table, on...))
}

func (g gradeClassDo) LeftJoin(table schema.Tabler, on ...field.Expr) IGradeClassDo {
	return g.withDO(g.DO.LeftJoin(table, on...))
}

func (g gradeClassDo) RightJoin(table schema.Tabler, on ...field.Expr) IGradeClassDo {
	return g.withDO(g.DO.RightJoin(table, on...))
}

func (g gradeClassDo) Group(cols ...field.Expr) IGradeClassDo {
	return g.withDO(g.DO.Group(cols...))
}

func (g gradeClassDo) Having(conds ...gen.Condition) IGradeClassDo {
	return g.withDO(g.DO.Having(conds...))
}

func (g gradeClassDo) Limit(limit int) IGradeClassDo {
	return g.withDO(g.DO.Limit(limit))
}

func (g gradeClassDo) Offset(offset int) IGradeClassDo {
	return g.withDO(g.DO.Offset(offset))
}

func (g gradeClassDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IGradeClassDo {
	return g.withDO(g.DO.Scopes(funcs...))
}

func (g gradeClassDo) Unscoped() IGradeClassDo {
	return g.withDO(g.DO.Unscoped())
}

func (g gradeClassDo) Create(values ...*model.GradeClass) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Create(values)
}

func (g gradeClassDo) CreateInBatches(values []*model.GradeClass, batchSize int) error {
	return g.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (g gradeClassDo) Save(values ...*model.GradeClass) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Save(values)
}

func (g gradeClassDo) First() (*model.GradeClass, error) {
	if result, err := g.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.GradeClass), nil
	}
}

func (g gradeClassDo) Take() (*model.GradeClass, error) {
	if result, err := g.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.GradeClass), nil
	}
}

func (g gradeClassDo) Last() (*model.GradeClass, error) {
	if result, err := g.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.GradeClass), nil
	}
}

func (g gradeClassDo) Find() ([]*model.GradeClass, error) {
	result, err := g.DO.Find()
	return result.([]*model.GradeClass), err
}

func (g gradeClassDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.GradeClass, err error) {
	buf := make([]*model.GradeClass, 0, batchSize)
	err = g.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (g gradeClassDo) FindInBatches(result *[]*model.GradeClass, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return g.DO.FindInBatches(result, batchSize, fc)
}

func (g gradeClassDo) Attrs(attrs ...field.AssignExpr) IGradeClassDo {
	return g.withDO(g.DO.Attrs(attrs...))
}

func (g gradeClassDo) Assign(attrs ...field.AssignExpr) IGradeClassDo {
	return g.withDO(g.DO.Assign(attrs...))
}

func (g gradeClassDo) Joins(fields ...field.RelationField) IGradeClassDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Joins(_f))
	}
	return &g
}

func (g gradeClassDo) Preload(fields ...field.RelationField) IGradeClassDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Preload(_f))
	}
	return &g
}

func (g gradeClassDo) FirstOrInit() (*model.GradeClass, error) {
	if result, err := g.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.GradeClass), nil
	}
}

func (g gradeClassDo) FirstOrCreate() (*model.GradeClass, error) {
	if result, err := g.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.GradeClass), nil
	}
}

func (g gradeClassDo) FindByPage(offset int, limit int) (result []*model.GradeClass, count int64, err error) {
	result, err = g.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = g.Offset(-1).Limit(-1).Count()
	return
}

func (g gradeClassDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	err = g.Offset(offset).Limit(limit).Scan(result)
	return
}

func (g gradeClassDo) Scan(result interface{}) (err error) {
	return g.DO.Scan(result)
}

func (g gradeClassDo) Delete(models ...*model.GradeClass) (result gen.ResultInfo, err error) {
	return g.DO.Delete(models)
}

func (g *gradeClassDo) withDO(do gen.Dao) *gradeClassDo {
	g.DO = *do.(*gen.DO)
	return g
}