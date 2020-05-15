package flag

// UserFlag is response code and message for user controller
type UserFlag struct {
	Message string
	Flag    string
	Error
}

var (
	/********************************
	 * GET MANY USER ENDPOINT
	 ********************************/

	// GetUsersSuccess is response code and message
	// for get users with success result
	GetUsersSuccess UserFlag = UserFlag{
		Message: "Success get all users",
	}

	/********************************
	 * GET USER ENDPOINT
	 ********************************/

	// GetUserSuccess is response code and message
	// for get user with success result
	GetUserSuccess UserFlag = UserFlag{
		Message: "Success get user",
	}

	// GetUserNotFound is response code and message
	// For get user with user not found
	// in database
	GetUserNotFound UserFlag = UserFlag{
		Message: "User not found",
		Error: Error{
			Message: "User with this ID not enough",
			Flag:    "USER_NOT_FOUND",
		},
	}

	// GetUserInvalidParamID is response code and message
	// For get user with invalid param uri
	GetUserInvalidParamID UserFlag = UserFlag{
		Flag: "BAD_REQUEST",
		Error: Error{
			Message: "Param must be of type integer, required",
			Flag:    "INVALID_BODY",
		},
	}

	/********************************
	 * CREATE USER ENDPOINT
	 ********************************/

	// CreateUserSuccess is response code and message
	// for create user with success result
	CreateUserSuccess UserFlag = UserFlag{
		Message: "Success create user",
	}

	// CreateUserAlreadyExist is response code and message
	// for create user with user already
	// exist in database
	CreateUserAlreadyExist UserFlag = UserFlag{
		Flag: "BAD_REQUEST",
		Error: Error{
			Message: "User with this email already exist",
			Flag:    "USER_ALREADY_EXIST",
		},
	}

	// CreateUserInvalidBody is response code and message
	// for create user with invalid body
	CreateUserInvalidBody UserFlag = UserFlag{
		Flag: "BAD_REQUEST",
		Error: Error{
			Flag: "INVALID_BODY",
		},
	}

	/********************************
	 * UPDATE USER ENDPOINT
	 ********************************/

	// UpdateUserSuccess is response code and message
	// for update user with success result
	UpdateUserSuccess UserFlag = UserFlag{
		Message: "Success update user",
	}

	// UpdateUserNotExist is response code and message
	// for update user with user not
	// exist in database
	UpdateUserNotExist UserFlag = UserFlag{
		Flag: "BAD_REQUEST",
		Error: Error{
			Message: "User with this id not exist",
			Flag:    "USER_NOT_FOUND",
		},
	}

	// UpdateUserInvalidBody is response code and message
	// for update user with user not
	// exist in database
	UpdateUserInvalidBody UserFlag = UserFlag{
		Flag: "BAD_REQUEST",
		Error: Error{
			Flag: "INVALID_BODY",
		},
	}

	// UpdateUserInvlidParamURI is response code and message
	// for update user with invalid param uri
	UpdateUserInvlidParamURI UserFlag = UserFlag{
		Flag: "BAD_REQUEST",
		Error: Error{
			Message: "Param must be of type integer, required",
			Flag:    "INVALID_BODY",
		},
	}

	/********************************
	 * DELETE USER ENDPOINT
	 ********************************/

	// DeleteUserSuccess is response code and message
	// for delete user with success result
	DeleteUserSuccess UserFlag = UserFlag{
		Message: "Success delete user",
	}

	// DeleteUserNotExist is response code and message
	// for delete user with user not
	// exist in database
	DeleteUserNotExist UserFlag = UserFlag{
		Flag: "BAD_REQUEST",
		Error: Error{
			Message: "User with this id not exist",
			Flag:    "USER_NOT_FOUND",
		},
	}

	// DeleteUserInvalidParamURI is response code and message
	// for delete user with invalid param uri
	DeleteUserInvalidParamURI UserFlag = UserFlag{
		Flag: "BAD_REQUEST",
		Error: Error{
			Message: "Param must be of type integer, required",
			Flag:    "INVALID_BODY",
		},
	}
)
