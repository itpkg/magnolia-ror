package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
)

//Page page info
type Page struct {
	Locale      string
	Title       string
	SubTitle    string
	Keywords    string
	Description string
	Copyright   string

	Home   string
	Author Author
	Links  []Link
}

//Author author info
type Author struct {
	Name  string
	Email string
}

//Link link info
type Link struct {
	Title string
	Href  string
}

func (p *Engine) getInfo(c *gin.Context) (interface{}, error) {
	locale := c.MustGet("locale").(*language.Tag).String()
	var page Page
	err := p.Dao.Get(fmt.Sprintf("%s://site/info", locale), &page)
	return page, err
}

func (p *Engine) postInstall(c *gin.Context) (interface{}, error) {
	var count uint
	err := p.Db.Model(&User{}).Count(&count).Error
	if err == nil {
		if count > 0 {
			err = errors.New("table users not empty")
		}
	}

	var user *User
	var fm FmSignUp
	if err == nil {
		err = c.Bind(&fm)
	}
	if err == nil {
		if fm.Password != fm.RePassword {
			err = errors.New("passwords not match")
		}
	}
	if err == nil {
		user, err = p.Dao.SignUp(fm.Email, fm.Name, fm.Password)
	}
	if err == nil {
		err = p.Db.Model(user).Update("confirmed_at", time.Now()).Error
	}
	if err == nil {
		for _, rn := range []string{"root", "admin"} {
			var role *Role
			role, err = p.Dao.Role(rn, "-", 0)
			if err == nil {
				err = p.Dao.Allow(role.ID, user.ID, 10, 0, 0)
			}
			if err != nil {
				break
			}
		}
	}
	if err == nil {
		err = p.Db.Model(user).Update("confirmed_at", time.Now()).Error
	}
	return user, err
}
