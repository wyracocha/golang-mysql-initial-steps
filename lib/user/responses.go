package user

type responseModel struct {
	status int
	message string
}
type GenericResponse struct {

}
type responseKeys struct {
	Ok responseModel
	NoContent responseModel
	InternalError responseModel
	NotFound responseModel
	BadRequest responseModel
	TimeOut responseModel
}

func Response () responseKeys {
	r := responseKeys {
		Ok:responseModel{
			status: 200,
			message: "Ok",
		},
		NoContent:responseModel{
			status: 204,
			message: "No Content",
		},
		InternalError:responseModel{
			status: 500,
			message: "Internal Server Error",
		},
		NotFound:responseModel{
			status: 404,
			message: "Not found",
		},
		BadRequest:responseModel{
			status: 400,
			message: "Bad request",
		},
		TimeOut:responseModel{
			status: 504,
			message: "Gateway TimeOut",
		},
	}
	return r
}