package auth

import (
	"bytes"
	"encoding/gob"
)

//Worker register worker handler
func (p *Engine) Worker() {
	p.Jobber.Register("email", func(args []byte) error {
		var buf bytes.Buffer
		dec := gob.NewDecoder(&buf)
		buf.Write(args)
		var model map[string]string
		if err := dec.Decode(&model); err != nil {
			return err
		}
		if IsProduction() {
			//TODO
		} else {
			p.Logger.Infof(
				"send email to: %s\n%s\n%s",
				model["to"],
				model["title"],
				model["body"],
			)
		}

		return nil
	})
}
