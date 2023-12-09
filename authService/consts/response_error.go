package consts

const (
	Success = 20000 //Normal return
	//controller the error code for the layer starts at 4000
	BindingJsonErr = 40000 // Bind json failed
	GetFileErr     = 40007 // Failed to get uploaded file
	PathParamErr   = 40008 // url parameter error
	UpdateIdErr    = 40009 // Update ID error
	NotFoundErr    = 40010 // Object not found

	//service Layer error codes start at 50000	common.GenShortResponse(c, consts.Success, response, "")
	ServerErr             = 50000 //
	ValidateErr           = 50005 // Data validation failed
	UserNameOrPasswordErr = 50006 // User name or password error
	PermissionErr         = 50007 // No permissions
	TokenValidErr         = 50008 // token void
	FileUploadErr         = 50009 // File upload failed
	MigrationErr          = 50010 // Upgrade failed

	//repository the error code for the repository layer starts at 60000
	DBInsertErr = 60000 // Data insertion failure
	DBUpdateErr = 60001 // Data Update Failed
	DBSelectErr = 60002 // Data Query failed
	DBDeleteErr = 60003 // Data deletion failed
)
