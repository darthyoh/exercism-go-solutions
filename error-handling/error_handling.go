package erratum

import "errors"

func Use(o ResourceOpener, input string) error {

  var (
    resource Resource
    err error
  )

  if resource, err = o(); err !=nil {
    if _, ok := err.(TransientError); ok {
      return Use(o, input)
    }
    return err 
  }

  defer resource.Close()

  frob(resource, input, &err)
  return err
}

func frob(resource Resource, input string, err *error) {
    defer func() {
    if p:= recover(); p!=nil {
      switch e:=p.(type) {
      case FrobError:
        resource.Defrob(e.defrobTag)
        *err = e
      case error:
        *err = e
      default :
        *err = errors.New("Other error")
      }
    }
  }()
	resource.Frob(input)
}
