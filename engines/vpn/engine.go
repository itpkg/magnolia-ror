package vpn

import (
	"github.com/facebookgo/inject"
	"github.com/gin-gonic/gin"
	"github.com/itpkg/magnolia/web"
	"github.com/jinzhu/gorm"
	"github.com/urfave/cli"
)

//Engine engine
type Engine struct {
}

//Map map objects
func (p *Engine) Map(*inject.Graph) error {
	return nil

}

//Mount mount
func (p *Engine) Mount(*gin.Engine) {

}

//Migrate db:migrate
func (p *Engine) Migrate(*gorm.DB) {}

//Seed db:seed
func (p *Engine) Seed() {}

//Worker register job handler
func (p *Engine) Worker() {}

//Shell command line options
func (p *Engine) Shell() []cli.Command {
	return []cli.Command{}
}

func init() {
	web.Register(&Engine{})
}
