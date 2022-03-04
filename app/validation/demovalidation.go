package validation

import (
	"carefree/app/agreement"
	"errors"
)

type DemoValidation struct {
}

func NewDemoValidation() *DemoValidation {
	return &DemoValidation{}
}

func (v *DemoValidation) HomeValidate(params agreement.HomeRequest) (err error) {
	if params.Uid == 0 {
		err = errors.New("uid is required")
		return err
	}

	return nil
}
