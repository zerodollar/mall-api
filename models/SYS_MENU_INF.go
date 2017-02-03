package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type SYSMENUINF struct {
	CATEGORY   string    `orm:"column(CATEGORY);size(20)"      json:"category"`
	Id         int       `orm:"column(ID);pk"                  json:"id"`
	NAME       string    `orm:"column(NAME);size(200)"         json:"name"`
	REMARK     string    `orm:"column(REMARK);size(200);null"  json:"remark"`
	PARENTID   string    `orm:"column(PARENT_ID);size(20);null" json:"parentid"`
	LEVEL      string    `orm:"column(LEVEL);size(2)"           json:"level"`
	POS        int       `orm:"column(POS);null"                json:"pos"`
	URL        string    `orm:"column(URL);size(200);null"      json:"url"`
	ICON       string    `orm:"column(ICON);size(100);null"     json:"icon"`
	STATUS     string    `orm:"column(STATUS);size(4);null"     json:"status"`
	CREATEDATE time.Time `orm:"column(CREATE_DATE);type(datetime);null;auto_now_add" json:"-"`
	CREATEEMP  string    `orm:"column(CREATE_EMP);size(10);null" json:"-"`
	MODIFYDATE time.Time `orm:"column(MODIFY_DATE);type(datetime);null" json:"-"`
	MODIFYEMP  string    `orm:"column(MODIFY_EMP);size(10);null" json:"-"`
}

func (t *SYSMENUINF) TableName() string {
	return "SYS_MENU_INF"
}

func init() {
	orm.RegisterModel(new(SYSMENUINF))
}

// AddSYSMENUINF insert a new SYSMENUINF into database and returns
// last inserted Id on success.
func AddSYSMENUINF(m *SYSMENUINF) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSYSMENUINFById retrieves SYSMENUINF by Id. Returns error if
// Id doesn't exist
func GetSYSMENUINFById(id int) (v *SYSMENUINF, err error) {
	o := orm.NewOrm()
	v = &SYSMENUINF{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSYSMENUINF retrieves all SYSMENUINF matches certain condition. Returns empty list if
// no records exist
func GetAllSYSMENUINF(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SYSMENUINF))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []SYSMENUINF
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateSYSMENUINF updates SYSMENUINF by Id and returns error if
// the record to be updated doesn't exist
func UpdateSYSMENUINFById(m *SYSMENUINF) (err error) {
	o := orm.NewOrm()
	v := SYSMENUINF{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSYSMENUINF deletes SYSMENUINF by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSYSMENUINF(id int) (err error) {
	o := orm.NewOrm()
	v := SYSMENUINF{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SYSMENUINF{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
