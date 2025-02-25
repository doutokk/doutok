package casbin

import (
	_ "embed"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/doutokk/doutok/app/auth/biz/dal"
	"github.com/doutokk/doutok/app/auth/biz/dal/mysql"
	"github.com/doutokk/doutok/common/utils"
	"log"
	"strings"
)

var (
	enforcer *casbin.Enforcer
	err      error
	//go:embed auth_model.conf
	modelFile []byte
	//go:embed auth_policy.csv
	policyFile []byte
)

/*
按如下约定:
  1. 所有策略只针对角色组设置
  2. 用户关联到组(一个用户可以有多个组)
+-------+-------+-----------+--------+----+----+----+
| ptype | v0    | v1        | v2     | v3 | v4 | v5 |
+-------+-------+-----------+--------+----+----+----+
| p     | admin | /api/user | GET    |    |    |    |
+-------+-------+-----------+--------+----+----+----+
| p     | admin | /api/user | DELETE |    |    |    |
+-------+-------+-----------+--------+----+----+----+
| p     | user  | /api/user | GET    |    |    |    |
+-------+-------+-----------+--------+----+----+----+
| ...   | ...   | ...       |        |    |    |    |
+-------+-------+-----------+--------+----+----+----+
| g     | leo   | admin     |        |    |    |    |
+-------+-------+-----------+--------+----+----+----+
| g     | leo2  | admin     |        |    |    |    |
+-------+-------+-----------+--------+----+----+----+
| g     | leo3  | user      |        |    |    |    |
+-------+-------+-----------+--------+----+----+----+
*/

func Init() {
	utils.WriteFile("infra/casbin/auth_model.conf", modelFile)
	utils.WriteFile("infra/casbin/auth_policy.csv", policyFile)

	// 鉴权中间件
	adapter, err := gormadapter.NewAdapterByDB(mysql.DB)
	if err != nil {
		panic("auth init failed")
	}

	m, err := model.NewModelFromFile("infra/casbin/auth_model.conf")
	enforcer, err = casbin.NewEnforcer(m, adapter)
	if err != nil {
		hlog.Error(err)
		panic("auth init failed")
	}
	err = enforcer.LoadPolicy()

	if err != nil {
		panic("auth init failed")
	}
}

// (RoleName, Url, Method) 对应于 `CasbinRule` 表中的 (v0, v1, v2)
type RolePolicy struct {
	RoleName string `gorm:"column:v0"`
	Url      string `gorm:"column:v1"`
	Method   string `gorm:"column:v2"`
}

// 拿着csv用Ai生成就行,放在cmd中
func InitPolicy() {
	dal.Init()
	utils.WriteFile("infra/casbin/auth_model.conf", modelFile)
	utils.WriteFile("infra/casbin/auth_policy.csv", policyFile)

	// 鉴权中间件
	adapter, err := gormadapter.NewAdapterByDB(mysql.DB)
	if err != nil {
		panic("auth init failed")
	}

	m, err := model.NewModelFromFile("infra/casbin/auth_model.conf")
	enforcer, err = casbin.NewEnforcer(m, adapter)
	if err != nil {
		hlog.Error(err)
		panic("auth init failed")
	}

	rolePolicies := []RolePolicy{
		{"base", "/user/login", "POST"},
		{"base", "/user/register", "POST"},
		{"base", "/product/*", "GET"},
		{"base", "/product", "GET"},
		{"base", "/product", "POST"},

		{"user", "/cart", "GET"},
		{"user", "/cart/edit", "POST"},
		{"user", "/order", "POST"},
		{"user", "/order/*", "GET"},
		{"user", "/order", "GET"},
		{"user", "/payment", "POST"},
	}

	// 遍历并插入
	for _, rp := range rolePolicies {
		err := CreateRolePolicy(rp)
		if err != nil {
			log.Printf("Failed to insert policy for role %s with url %s and method %s: %v\n", rp.RoleName, rp.Url, rp.Method, err)
		} else {
			log.Printf("Successfully inserted policy for role %s with url %s and method %s\n", rp.RoleName, rp.Url, rp.Method)
		}
	}

	CreateUserRole("test", "user")
	CreateUserRole("user", "base")
}

// 创建角色组权限, 已有的会忽略ca
func CreateRolePolicy(r RolePolicy) error {
	// 不直接操作数据库，利用enforcer简化操作
	err := enforcer.LoadPolicy()
	if err != nil {
		return err
	}
	_, err = enforcer.AddPolicy(r.RoleName, r.Url, r.Method)
	if err != nil {
		return err
	}
	return enforcer.SavePolicy()
}

// 修改角色组权限
func UpdateRolePolicy(old, new RolePolicy) error {
	_, err := enforcer.UpdatePolicy([]string{old.RoleName, old.Url, old.Method},
		[]string{new.RoleName, new.Url, new.Method})
	if err != nil {
		return err
	}
	return enforcer.SavePolicy()
}

// 删除角色组权限
func DeleteRolePolicy(r RolePolicy) error {
	_, err := enforcer.RemovePolicy(r.RoleName, r.Url, r.Method)
	if err != nil {
		return err
	}
	return enforcer.SavePolicy()
}

// 角色组中添加用户, 没有组默认创建
func CreateUserRole(username, rolename string) error {
	_, err := enforcer.AddGroupingPolicy(username, rolename)
	if err != nil {
		return err
	}
	return enforcer.SavePolicy()
}

// 角色组中删除用户
func DeleteUserRole(username, rolename string) error {
	_, err := enforcer.RemoveGroupingPolicy(username, rolename)
	if err != nil {
		return err
	}
	return enforcer.SavePolicy()
}

// 这里还需要uri和method
func CheckAuthByRBAC(sub string, obj string, act string) bool {
	// todo:现在是魔改路径,取出obj ？和？后面的参数
	obj = processObj(obj)
	ok, err := enforcer.Enforce(sub, obj, act)
	if err != nil {
		hlog.Error("checkAuthByRBAC failed: ", err)
		return false
	}
	if !ok {
		return false
	}
	return true
}

// 取出路径的?后面的参数
func processObj(obj string) string {
	index := strings.Index(obj, "?")
	if index > -1 {
		return obj[:index]
	}
	return obj
}
