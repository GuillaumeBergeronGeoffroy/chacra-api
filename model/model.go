/*
	Package model ...
*/
package model

type Model interface {
	Validate() (err error)
	BeforeSave() (err error)
	AfterSave() (err error)
}
