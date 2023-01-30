// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/SilleCao/golang/go-micro-demo/internal/modules/sys/model"
)

func newSysMenu(db *gorm.DB, opts ...gen.DOOption) sysMenu {
	_sysMenu := sysMenu{}

	_sysMenu.sysMenuDo.UseDB(db, opts...)
	_sysMenu.sysMenuDo.UseModel(&model.SysMenu{})

	tableName := _sysMenu.sysMenuDo.TableName()
	_sysMenu.ALL = field.NewAsterisk(tableName)
	_sysMenu.ID = field.NewInt64(tableName, "id")
	_sysMenu.Pid = field.NewInt64(tableName, "pid")
	_sysMenu.Name = field.NewString(tableName, "name")
	_sysMenu.URL = field.NewString(tableName, "url")
	_sysMenu.Permissions = field.NewString(tableName, "permissions")
	_sysMenu.Type = field.NewInt32(tableName, "type")
	_sysMenu.Icon = field.NewString(tableName, "icon")
	_sysMenu.Sort = field.NewInt32(tableName, "sort")
	_sysMenu.Creator = field.NewInt64(tableName, "creator")
	_sysMenu.CreateDate = field.NewTime(tableName, "create_date")
	_sysMenu.Updater = field.NewInt64(tableName, "updater")
	_sysMenu.UpdateDate = field.NewTime(tableName, "update_date")

	_sysMenu.fillFieldMap()

	return _sysMenu
}

type sysMenu struct {
	sysMenuDo sysMenuDo

	ALL         field.Asterisk
	ID          field.Int64  // id
	Pid         field.Int64  // 上级ID，一级菜单为0
	Name        field.String // 名称
	URL         field.String // 菜单URL
	Permissions field.String // 授权(多个用逗号分隔，如：sys:user:list,sys:user:save)
	Type        field.Int32  // 类型   0：菜单   1：按钮
	Icon        field.String // 菜单图标
	Sort        field.Int32  // 排序
	Creator     field.Int64  // 创建者
	CreateDate  field.Time   // 创建时间
	Updater     field.Int64  // 更新者
	UpdateDate  field.Time   // 更新时间

	fieldMap map[string]field.Expr
}

func (s sysMenu) Table(newTableName string) *sysMenu {
	s.sysMenuDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysMenu) As(alias string) *sysMenu {
	s.sysMenuDo.DO = *(s.sysMenuDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysMenu) updateTableName(table string) *sysMenu {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.Pid = field.NewInt64(table, "pid")
	s.Name = field.NewString(table, "name")
	s.URL = field.NewString(table, "url")
	s.Permissions = field.NewString(table, "permissions")
	s.Type = field.NewInt32(table, "type")
	s.Icon = field.NewString(table, "icon")
	s.Sort = field.NewInt32(table, "sort")
	s.Creator = field.NewInt64(table, "creator")
	s.CreateDate = field.NewTime(table, "create_date")
	s.Updater = field.NewInt64(table, "updater")
	s.UpdateDate = field.NewTime(table, "update_date")

	s.fillFieldMap()

	return s
}

func (s *sysMenu) WithContext(ctx context.Context) *sysMenuDo { return s.sysMenuDo.WithContext(ctx) }

func (s sysMenu) TableName() string { return s.sysMenuDo.TableName() }

func (s sysMenu) Alias() string { return s.sysMenuDo.Alias() }

func (s *sysMenu) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysMenu) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 12)
	s.fieldMap["id"] = s.ID
	s.fieldMap["pid"] = s.Pid
	s.fieldMap["name"] = s.Name
	s.fieldMap["url"] = s.URL
	s.fieldMap["permissions"] = s.Permissions
	s.fieldMap["type"] = s.Type
	s.fieldMap["icon"] = s.Icon
	s.fieldMap["sort"] = s.Sort
	s.fieldMap["creator"] = s.Creator
	s.fieldMap["create_date"] = s.CreateDate
	s.fieldMap["updater"] = s.Updater
	s.fieldMap["update_date"] = s.UpdateDate
}

func (s sysMenu) clone(db *gorm.DB) sysMenu {
	s.sysMenuDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysMenu) replaceDB(db *gorm.DB) sysMenu {
	s.sysMenuDo.ReplaceDB(db)
	return s
}

type sysMenuDo struct{ gen.DO }

func (s sysMenuDo) Debug() *sysMenuDo {
	return s.withDO(s.DO.Debug())
}

func (s sysMenuDo) WithContext(ctx context.Context) *sysMenuDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysMenuDo) ReadDB() *sysMenuDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysMenuDo) WriteDB() *sysMenuDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysMenuDo) Session(config *gorm.Session) *sysMenuDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysMenuDo) Clauses(conds ...clause.Expression) *sysMenuDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysMenuDo) Returning(value interface{}, columns ...string) *sysMenuDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysMenuDo) Not(conds ...gen.Condition) *sysMenuDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysMenuDo) Or(conds ...gen.Condition) *sysMenuDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysMenuDo) Select(conds ...field.Expr) *sysMenuDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysMenuDo) Where(conds ...gen.Condition) *sysMenuDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysMenuDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *sysMenuDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s sysMenuDo) Order(conds ...field.Expr) *sysMenuDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysMenuDo) Distinct(cols ...field.Expr) *sysMenuDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysMenuDo) Omit(cols ...field.Expr) *sysMenuDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysMenuDo) Join(table schema.Tabler, on ...field.Expr) *sysMenuDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysMenuDo) LeftJoin(table schema.Tabler, on ...field.Expr) *sysMenuDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysMenuDo) RightJoin(table schema.Tabler, on ...field.Expr) *sysMenuDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysMenuDo) Group(cols ...field.Expr) *sysMenuDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysMenuDo) Having(conds ...gen.Condition) *sysMenuDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysMenuDo) Limit(limit int) *sysMenuDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysMenuDo) Offset(offset int) *sysMenuDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysMenuDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *sysMenuDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysMenuDo) Unscoped() *sysMenuDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysMenuDo) Create(values ...*model.SysMenu) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysMenuDo) CreateInBatches(values []*model.SysMenu, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysMenuDo) Save(values ...*model.SysMenu) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysMenuDo) First() (*model.SysMenu, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysMenu), nil
	}
}

func (s sysMenuDo) Take() (*model.SysMenu, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysMenu), nil
	}
}

func (s sysMenuDo) Last() (*model.SysMenu, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysMenu), nil
	}
}

func (s sysMenuDo) Find() ([]*model.SysMenu, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysMenu), err
}

func (s sysMenuDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysMenu, err error) {
	buf := make([]*model.SysMenu, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysMenuDo) FindInBatches(result *[]*model.SysMenu, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysMenuDo) Attrs(attrs ...field.AssignExpr) *sysMenuDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysMenuDo) Assign(attrs ...field.AssignExpr) *sysMenuDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysMenuDo) Joins(fields ...field.RelationField) *sysMenuDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysMenuDo) Preload(fields ...field.RelationField) *sysMenuDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysMenuDo) FirstOrInit() (*model.SysMenu, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysMenu), nil
	}
}

func (s sysMenuDo) FirstOrCreate() (*model.SysMenu, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysMenu), nil
	}
}

func (s sysMenuDo) FindByPage(offset int, limit int) (result []*model.SysMenu, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s sysMenuDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysMenuDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysMenuDo) Delete(models ...*model.SysMenu) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysMenuDo) withDO(do gen.Dao) *sysMenuDo {
	s.DO = *do.(*gen.DO)
	return s
}
