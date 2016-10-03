package auth

import (
	"crypto/aes"

	"github.com/SermoDigital/jose/crypto"
	"github.com/facebookgo/inject"
	"github.com/itpkg/magnolia/cache"
	"github.com/itpkg/magnolia/i18n"
	"github.com/itpkg/magnolia/jobber"
	"github.com/itpkg/magnolia/web"
	"github.com/jinzhu/gorm"
	logging "github.com/op/go-logging"
	"github.com/spf13/viper"
)

//Engine engine
type Engine struct {
	Cache             cache.Store            `inject:""`
	Db                *gorm.DB               `inject:""`
	Dao               *Dao                   `inject:""`
	Jobber            jobber.Jobber          `inject:""`
	Logger            *logging.Logger        `inject:""`
	Jwt               *Jwt                   `inject:""`
	PasswordEncryptor *web.PasswordEncryptor `inject:""`
	I18n              *i18n.I18n             `inject:""`
}

//Map map objects
func (p *Engine) Map(inj *inject.Graph) error {

	cip, err := aes.NewCipher([]byte(viper.GetString("secrets.aes")))
	if err != nil {
		return err
	}

	return inj.Provide(
		&inject.Object{Value: cip},
		&inject.Object{Value: []byte(viper.GetString("secrets.jwt")), Name: "jwt.key"},
		&inject.Object{Value: crypto.SigningMethodHS512, Name: "jwt.method"},
		&inject.Object{Value: &cache.RedisStore{}},
		&inject.Object{Value: &jobber.RedisJobber{
			Timeout:  viper.GetInt("workers.timeout"),
			Handlers: make(map[string]jobber.Handler),
		}},
	)

}

//Migrate db:migrate
func (p *Engine) Migrate(db *gorm.DB) {
	i18n.Migrate(db)
	db.AutoMigrate(
		&Setting{}, &Notice{}, &LeaveWord{},
		&User{}, &Role{}, &Permission{}, &Log{},
	)

	db.Model(&User{}).AddUniqueIndex("idx_user_provider_type_id", "provider_type", "provider_id")
	db.Model(&Role{}).AddUniqueIndex("idx_roles_name_resource_type_id", "name", "resource_type", "resource_id")
	db.Model(&Permission{}).AddUniqueIndex("idx_permissions_user_role", "user_id", "role_id")
}

//Seed db:seed
func (p *Engine) Seed() {}

func init() {
	web.Register(&Engine{})
}
