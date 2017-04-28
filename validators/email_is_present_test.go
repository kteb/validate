package validators_test

import (
	"testing"

	"github.com/markbates/validate"
	"github.com/stretchr/testify/require"
	. "github.com/markbates/validate/validators"
)


func Test_EmailIsPresent(t *testing.T) {
	r := require.New(t)

	var tests = []struct {
		email    string
		valid    bool
	}{
		{"", false},
		{"foo@bar.com", true},
		{"x@x.x", true},
		{"foo@bar.com.au", true},
		{"foo+bar@bar.com", true},
		{"foo@bar.coffee", true},
		{"foo@bar.中文网", true},
		{"invalidemail@", false},
		{"invalid.com", false},
		{"@invalid.com", false},
		{"test|123@m端ller.com", true},
		{"hans@m端ller.com", true},
		{"hans.m端ller@test.com", true},
		{"NathAn.daVIeS@DomaIn.cOM", true},
		{"NATHAN.DAVIES@DOMAIN.CO.UK", true},
	}
	for _, test := range tests {
		v := EmailIsPresent{Name: "email", Field: test.email}
		errors := validate.NewErrors()
		v.IsValid(errors)
		r.Equal(test.valid, !errors.HasAny())
	}
}
