package helper

import(
	"regexp"
)

// for exporting a function you have to Camelcasesisize the function name 

func Check_exp(exp *regexp.Regexp, txt string)bool{
	if exp.MatchString(txt){
		return true
	}
	return false
}