package connectErr

import "errors"

var (
	SSH_AUTH_ERROR             = errors.New("happen severe error")
	SSH_AUTH_DENIED            = errors.New("server not accept public key validate")
	SSH_AUTH_PARTIAL           = errors.New("still need other validation")
	SSH_AUTH_SUCCESS           = errors.New("public key validate success,can use ssh_userauth_publickey")
	SSH_AUTH_AGAIN             = errors.New("call it later")
	SSH_AUTH_INFO              = errors.New("server ask some question")
	SSH_FX_OK                  = errors.New("no error")
	SSH_FX_EOF                 = errors.New("happen end singal")
	SSH_FX_NO_SUCH_FILE        = errors.New("file not exist")
	SSH_FX_PERMISSION_DENIED   = errors.New("file permission denied")
	SSH_FX_FAILURE             = errors.New("match file failure")
	SSH_FX_BAD_MESSAGE         = errors.New("get bad message from server")
	SSH_FX_NO_CONNECTION       = errors.New("connection not good")
	SSH_FX_CONNECTION_LOST     = errors.New("lost connection")
	SSH_FX_OP_UNSUPPORTED      = errors.New("operation not supported")
	SSH_FX_INVALID_HANDLE      = errors.New("file handler invalid")
	SSH_FX_NO_SUCH_PATH        = errors.New("no such file or path")
	SSH_FX_FILE_ALREADY_EXISTS = errors.New("file already exist")
	SSH_FX_WRITE_PROTECT       = errors.New("write protected os")
	SSH_FX_NO_MEDIA            = errors.New("remote driver no media")
)

func GetConnectSSHErr(err error) error {

	switch err {
	case SSH_AUTH_ERROR:
		return SSH_AUTH_ERROR
	case SSH_AUTH_DENIED:
		return SSH_AUTH_DENIED
	case SSH_AUTH_PARTIAL:
		return SSH_AUTH_PARTIAL
	case SSH_AUTH_SUCCESS:
		return SSH_AUTH_SUCCESS
	case SSH_AUTH_AGAIN:
		return SSH_AUTH_AGAIN
	case SSH_AUTH_INFO:
		return SSH_AUTH_INFO
	case SSH_FX_EOF:
		return SSH_FX_EOF
	case SSH_FX_NO_SUCH_FILE:
		return SSH_FX_NO_SUCH_FILE
	case SSH_FX_PERMISSION_DENIED:
		return SSH_FX_PERMISSION_DENIED
	case SSH_FX_FAILURE:
		return SSH_FX_FAILURE
	case SSH_FX_BAD_MESSAGE:
		return SSH_FX_BAD_MESSAGE
	case SSH_FX_NO_CONNECTION:
		return SSH_FX_NO_CONNECTION
	case SSH_FX_CONNECTION_LOST:
		return SSH_FX_CONNECTION_LOST
	case SSH_FX_OP_UNSUPPORTED:
		return SSH_FX_OP_UNSUPPORTED
	case SSH_FX_INVALID_HANDLE:
		return SSH_FX_INVALID_HANDLE
	case SSH_FX_NO_SUCH_PATH:
		return SSH_FX_NO_SUCH_PATH
	case SSH_FX_FILE_ALREADY_EXISTS:
		return SSH_FX_FILE_ALREADY_EXISTS
	case SSH_FX_WRITE_PROTECT:
		return SSH_FX_WRITE_PROTECT
	case SSH_FX_NO_MEDIA:
		return SSH_FX_NO_MEDIA
	default:
		return err
	}
}
