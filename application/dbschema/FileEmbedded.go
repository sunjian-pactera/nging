// @generated Do not edit this file, which is automatically generated by the generator.

package dbschema

import (
	"fmt"

	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	"github.com/webx-top/echo"
	"github.com/webx-top/echo/param"
	
)

type Slice_FileEmbedded []*FileEmbedded

func (s Slice_FileEmbedded) Range(fn func(m factory.Model) error ) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_FileEmbedded) RangeRaw(fn func(m *FileEmbedded) error ) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

// FileEmbedded 嵌入文件
type FileEmbedded struct {
	param   *factory.Param
	trans	*factory.Transaction
	objects []*FileEmbedded
	namer   func(string) string
	connID  int
	context echo.Context
	
	Id        	uint64  	`db:"id,omitempty,pk" bson:"id,omitempty" comment:"主键" json:"id" xml:"id"`
	Project   	string  	`db:"project" bson:"project" comment:"项目名" json:"project" xml:"project"`
	TableId   	uint64  	`db:"table_id" bson:"table_id" comment:"表主键" json:"table_id" xml:"table_id"`
	TableName 	string  	`db:"table_name" bson:"table_name" comment:"表名称" json:"table_name" xml:"table_name"`
	FieldName 	string  	`db:"field_name" bson:"field_name" comment:"字段名" json:"field_name" xml:"field_name"`
	FileIds   	string  	`db:"file_ids" bson:"file_ids" comment:"文件id列表" json:"file_ids" xml:"file_ids"`
	Embedded  	string  	`db:"embedded" bson:"embedded" comment:"是否(Y/N)为内嵌文件" json:"embedded" xml:"embedded"`
}

func (this *FileEmbedded) Trans() *factory.Transaction {
	return this.trans
}

func (this *FileEmbedded) Use(trans *factory.Transaction) factory.Model {
	this.trans = trans
	return this
}

func (this *FileEmbedded) SetContext(ctx echo.Context) factory.Model {
	this.context = ctx
	return this
}

func (this *FileEmbedded) Context() echo.Context {
	return this.context
}

func (this *FileEmbedded) SetConnID(connID int) factory.Model {
	this.connID = connID
	return this
}

func (this *FileEmbedded) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName,connID[0]).Use(this.trans)
	}
	return factory.NewModel(structName,this.connID).Use(this.trans)
}

func (this *FileEmbedded) Objects() []*FileEmbedded {
	if this.objects == nil {
		return nil
	}
	return this.objects[:]
}

func (this *FileEmbedded) NewObjects() factory.Ranger {
	return &Slice_FileEmbedded{}
}

func (this *FileEmbedded) InitObjects() *[]*FileEmbedded {
	this.objects = []*FileEmbedded{}
	return &this.objects
}

func (this *FileEmbedded) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(this.connID).SetTrans(this.trans).SetCollection(this.Name_()).SetModel(this)
}

func (this *FileEmbedded) SetNamer(namer func (string) string) factory.Model {
	this.namer = namer
	return this
}

func (this *FileEmbedded) Short_() string {
	return "file_embedded"
}

func (this *FileEmbedded) Struct_() string {
	return "FileEmbedded"
}

func (this *FileEmbedded) Name_() string {
	if this.namer != nil {
		return WithPrefix(this.namer(this.Short_()))
	}
	return WithPrefix(factory.TableNamerGet(this.Short_())(this))
}

func (this *FileEmbedded) Namer() func(string) string {
	return this.namer
}

func (this *FileEmbedded) CPAFrom(source factory.Model) factory.Model {
	this.SetContext(source.Context())
	this.Use(source.Trans())
	this.SetNamer(source.Namer())
	return this
}

func (this *FileEmbedded) SetParam(param *factory.Param) factory.Model {
	this.param = param
	return this
}

func (this *FileEmbedded) Param() *factory.Param {
	if this.param == nil {
		return this.NewParam()
	}
	return this.param
}

func (this *FileEmbedded) Get(mw func(db.Result) db.Result, args ...interface{}) error {
	return this.Param().SetArgs(args...).SetRecv(this).SetMiddleware(mw).One()
}

func (this *FileEmbedded) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.InitObjects()
	}
	return this.Param().SetArgs(args...).SetPage(page).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *FileEmbedded) GroupBy(keyField string, inputRows ...[]*FileEmbedded) map[string][]*FileEmbedded {
	var rows []*FileEmbedded
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string][]*FileEmbedded{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*FileEmbedded{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (this *FileEmbedded) KeyBy(keyField string, inputRows ...[]*FileEmbedded) map[string]*FileEmbedded {
	var rows []*FileEmbedded
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string]*FileEmbedded{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (this *FileEmbedded) AsKV(keyField string, valueField string, inputRows ...[]*FileEmbedded) map[string]interface{} {
	var rows []*FileEmbedded
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string]interface{}{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = dmap[valueField]
	}
	return r
}

func (this *FileEmbedded) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.InitObjects()
	}
	return this.Param().SetArgs(args...).SetOffset(offset).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *FileEmbedded) Add() (pk interface{}, err error) {
	this.Id = 0
	if len(this.Embedded) == 0 { this.Embedded = "Y" }
	err = DBI.Fire("creating", this, nil)
	if err != nil {
		return
	}
	pk, err = this.Param().SetSend(this).Insert()
	if err == nil && pk != nil {
		if v, y := pk.(uint64); y {
			this.Id = v
		} else if v, y := pk.(int64); y {
			this.Id = uint64(v)
		}
	}
	if err == nil {
		err = DBI.Fire("created", this, nil)
	}
	return
}

func (this *FileEmbedded) Edit(mw func(db.Result) db.Result, args ...interface{}) (err error) {
	
	if len(this.Embedded) == 0 { this.Embedded = "Y" }
	if err = DBI.Fire("updating", this, mw, args...); err != nil {
		return
	}
	if err = this.Setter(mw, args...).SetSend(this).Update(); err != nil {
		return
	}
	return DBI.Fire("updated", this, mw, args...)
}

func (this *FileEmbedded) Setter(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	return this.Param().SetArgs(args...).SetMiddleware(mw)
}

func (this *FileEmbedded) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (err error) {
	return this.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (this *FileEmbedded) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (err error) {
	
	if val, ok := kvset["embedded"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["embedded"] = "Y" } }
	m := *this
	m.FromRow(kvset)
	var editColumns []string
	for column := range kvset {
		editColumns = append(editColumns, column)
	}
	if err = DBI.FireUpdate("updating", &m, editColumns, mw, args...); err != nil {
		return
	}
	if err = this.Setter(mw, args...).SetSend(kvset).Update(); err != nil {
		return
	}
	return DBI.FireUpdate("updated", &m, editColumns, mw, args...)
}

func (this *FileEmbedded) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = this.Param().SetArgs(args...).SetSend(this).SetMiddleware(mw).Upsert(func() error { 
	if len(this.Embedded) == 0 { this.Embedded = "Y" }
		return DBI.Fire("updating", this, mw, args...)
	}, func() error { this.Id = 0
	if len(this.Embedded) == 0 { this.Embedded = "Y" }
		return DBI.Fire("creating", this, nil)
	})
	if err == nil && pk != nil {
		if v, y := pk.(uint64); y {
			this.Id = v
		} else if v, y := pk.(int64); y {
			this.Id = uint64(v)
		}
	}
	if err == nil {
		if pk == nil {
			err = DBI.Fire("updated", this, mw, args...)
		} else {
			err = DBI.Fire("created", this, nil)
		}
	} 
	return 
}

func (this *FileEmbedded) Delete(mw func(db.Result) db.Result, args ...interface{})  (err error) {
	
	if err = DBI.Fire("deleting", this, mw, args...); err != nil {
		return
	}
	if err = this.Param().SetArgs(args...).SetMiddleware(mw).Delete(); err != nil {
		return
	}
	return DBI.Fire("deleted", this, mw, args...)
}

func (this *FileEmbedded) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return this.Param().SetArgs(args...).SetMiddleware(mw).Count()
}

func (this *FileEmbedded) Reset() *FileEmbedded {
	this.Id = 0
	this.Project = ``
	this.TableId = 0
	this.TableName = ``
	this.FieldName = ``
	this.FileIds = ``
	this.Embedded = ``
	return this
}

func (this *FileEmbedded) AsMap() map[string]interface{} {
	r := map[string]interface{}{}
	r["Id"] = this.Id
	r["Project"] = this.Project
	r["TableId"] = this.TableId
	r["TableName"] = this.TableName
	r["FieldName"] = this.FieldName
	r["FileIds"] = this.FileIds
	r["Embedded"] = this.Embedded
	return r
}

func (this *FileEmbedded) FromRow(row map[string]interface{}) {
	for key, value := range row {
		switch key {
			case "id": this.Id = param.AsUint64(value)
			case "project": this.Project = param.AsString(value)
			case "table_id": this.TableId = param.AsUint64(value)
			case "table_name": this.TableName = param.AsString(value)
			case "field_name": this.FieldName = param.AsString(value)
			case "file_ids": this.FileIds = param.AsString(value)
			case "embedded": this.Embedded = param.AsString(value)
		}
	}
}

func (this *FileEmbedded) Set(key interface{}, value ...interface{}) {
	switch k := key.(type) {
		case map[string]interface{}:
			for kk, vv := range k {
				this.Set(kk, vv)
			}
		default:
			var (
				kk string
				vv interface{}
			)
			if k, y := key.(string); y {
				kk = k
			} else {
				kk = fmt.Sprint(key)
			}
			if len(value) > 0 {
				vv = value[0]
			}
			switch kk {
				case "Id": this.Id = param.AsUint64(vv)
				case "Project": this.Project = param.AsString(vv)
				case "TableId": this.TableId = param.AsUint64(vv)
				case "TableName": this.TableName = param.AsString(vv)
				case "FieldName": this.FieldName = param.AsString(vv)
				case "FileIds": this.FileIds = param.AsString(vv)
				case "Embedded": this.Embedded = param.AsString(vv)
			}
	}
}

func (this *FileEmbedded) AsRow() map[string]interface{} {
	r := map[string]interface{}{}
	r["id"] = this.Id
	r["project"] = this.Project
	r["table_id"] = this.TableId
	r["table_name"] = this.TableName
	r["field_name"] = this.FieldName
	r["file_ids"] = this.FileIds
	r["embedded"] = this.Embedded
	return r
}

func (this *FileEmbedded) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = this.AsRow()
	}
	return factory.BatchValidate(this.Short_(), kvset)
}

func (this *FileEmbedded) Validate(field string, value interface{}) error {
	return factory.Validate(this.Short_(), field, value)
}

