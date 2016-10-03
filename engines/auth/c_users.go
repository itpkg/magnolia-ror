package auth

import (
	"bytes"
	"errors"
	"fmt"
	"text/template"
	"time"

	"golang.org/x/text/language"

	"github.com/SermoDigital/jose/jws"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func (p *Engine) postUserSignIn(c *gin.Context) (interface{}, error) {
	var fm FmSignIn
	var tk []byte
	var user *User

	err := c.Bind(&fm)
	if err == nil {
		user, err = p.Dao.SignIn(fm.Email, fm.Password)
	}
	if err == nil {
		tk, err = p.Jwt.Sum(p.Dao.UserClaims(user), 7)
	}
	return gin.H{"user": user, "token": string(tk)}, err
}

func (p *Engine) postUserSignUp(c *gin.Context) (interface{}, error) {
	lng := c.MustGet("locale").(*language.Tag)
	var user *User
	var fm FmSignUp

	err := c.Bind(&fm)
	if err == nil {
		if fm.Password != fm.RePassword {
			err = errors.New("passwords not match")
		}
	}
	if err == nil {
		user, err = p.Dao.SignUp(fm.Email, fm.Name, fm.Password)
	}
	if err == nil {
		err = p.sendMail(lng, user, "confirm")
	}
	return user, err
}

func (p *Engine) postUserConfirm(c *gin.Context) (interface{}, error) {
	lng := c.MustGet("locale").(*language.Tag)

	var user *User
	var fm FmEmail
	err := c.Bind(&fm)
	if err == nil {
		user, err = p.Dao.GetUserByEmail(fm.Email)
	}
	if err == nil {
		if user.IsConfirmed() {
			err = fmt.Errorf("user [%s] was confirmed", fm.Email)
		}
	}
	if err == nil {
		err = p.sendMail(lng, user, "confirm")
	}
	return nil, err
}

func (p *Engine) postUserUnlock(c *gin.Context) (interface{}, error) {
	lng := c.MustGet("locale").(*language.Tag)

	var user *User
	var fm FmEmail
	err := c.Bind(&fm)
	if err == nil {
		user, err = p.Dao.GetUserByEmail(fm.Email)
	}
	if err == nil {
		if !user.IsLocked() {
			err = fmt.Errorf("user [%s] wasn't locked", fm.Email)
		}
	}
	if err == nil {
		err = p.sendMail(lng, user, "unlock")
	}
	return nil, err
}

func (p *Engine) postUserForgotPassword(c *gin.Context) (interface{}, error) {
	lng := c.MustGet("locale").(*language.Tag)

	var user *User
	var fm FmEmail

	err := c.Bind(&fm)
	if err == nil {
		user, err = p.Dao.GetUserByEmail(fm.Email)
	}
	if err == nil {
		err = p.sendMail(lng, user, "change-password")
	}
	return nil, err
}

func (p *Engine) postUserChangePassword(c *gin.Context) (interface{}, error) {
	var user *User
	var fm FmChangePassword
	err := c.Bind(&fm)
	if err == nil {
		if fm.Password != fm.RePassword {
			err = errors.New("passwords not match")
		}
	}
	var data map[string]interface{}
	data, err = p.Jwt.Validate([]byte(fm.Token))
	if err == nil {
		act := data["act"].(string)
		if act != "change-password" {
			err = fmt.Errorf("unknown action %s", act)
		}
	}
	if err == nil {
		user, err = p.Dao.GetUserByUID(data["uid"].(string))
	}
	if err == nil {
		if !user.IsAvailable() {
			err = fmt.Errorf("user [%s] wasn't available", user.Email)
		}
	}
	if err == nil {
		var password string
		password, err = p.PasswordEncryptor.Sum([]byte(fm.Password), 8)
		if err == nil {
			err = p.Db.Model(user).Update("password", password).Error
		}
	}
	return nil, err
}

func (p *Engine) getUserConfirm(c *gin.Context) (string, error) {
	var user *User
	var fm FmToken
	var data map[string]interface{}
	err := c.Bind(&fm)

	if err == nil {
		data, err = p.Jwt.Validate([]byte(fm.Token))
	}
	if err == nil {
		act := data["act"].(string)
		if act != "confirm" {
			err = fmt.Errorf("unknown action %s", act)
		}
	}
	if err == nil {
		user, err = p.Dao.GetUserByUID(data["uid"].(string))
	}
	if err == nil {
		if user.IsConfirmed() {
			err = fmt.Errorf("user %s was confirmed", user.Email)
		}
	}
	if err == nil {
		err = p.Db.Model(user).Update("confirmed_at", time.Now()).Error
	}

	return viper.GetString("home.front"), err
}

func (p *Engine) getUserUnlock(c *gin.Context) (string, error) {
	var user *User
	var fm FmToken
	var data map[string]interface{}
	err := c.Bind(&fm)

	if err == nil {
		data, err = p.Jwt.Validate([]byte(fm.Token))
	}
	if err == nil {
		act := data["act"].(string)
		if act != "unlock" {
			err = fmt.Errorf("unknown action %s", act)
		}
	}
	if err == nil {
		user, err = p.Dao.GetUserByUID(data["uid"].(string))
	}
	if err == nil {
		if !user.IsLocked() {
			err = fmt.Errorf("user %s was confirmed", user.Email)
		}
	}
	if err == nil {
		err = p.Db.Model(user).Update("locked_at", nil).Error
	}
	return viper.GetString("home.front"), err
}

//-----------------------------------------------------------------------------

func (p *Engine) sendMail(locale *language.Tag, user *User, action string) error {
	cm := jws.Claims{}
	cm["uid"] = user.UID
	cm["act"] = action
	var token string
	if tkn, err := p.Jwt.Sum(cm, 1); err == nil {
		token = string(tkn)
	} else {
		return err
	}

	var link string
	args := make(map[string]string)
	args["to"] = user.Email
	switch action {
	case "confirm":
		link = fmt.Sprintf("%s/users/confirm?token=%s", viper.GetString("home.backend"), token)
	case "unlock":
		link = fmt.Sprintf("%s/users/unlock?token=%s", viper.GetString("home.backend"), token)
	case "change-password":
		link = fmt.Sprintf("%s/users/change-password?token=%s", viper.GetString("home.front"), token)
	default:
		return fmt.Errorf("bad action %s", action)
	}

	var page Page
	if e := p.Dao.Get(fmt.Sprintf("%s://site/info", locale), &page); e != nil {
		p.Logger.Error(e)
	}

	if title, err := p.parse(p.I18n.T(locale, fmt.Sprintf("mail.auth.users.%s.title", action)), struct {
		Page Page
	}{
		Page: page,
	}); err == nil {
		args["title"] = title
	} else {
		return err
	}

	if body, err := p.parse(p.I18n.T(locale, fmt.Sprintf("mail.auth.users.%s.body", action)), struct {
		Page Page
		Link string
	}{
		Page: page,
		Link: link,
	}); err == nil {
		args["body"] = body
	} else {
		return err
	}

	return p.Jobber.Push("email", args)
}

func (p *Engine) parse(tpl string, args interface{}) (string, error) {
	var buf bytes.Buffer
	t := template.Must(template.New("").Parse(tpl))
	e := t.Execute(&buf, args)
	return buf.String(), e
}
