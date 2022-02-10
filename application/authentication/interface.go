package authentication

import "github.com/darllantissei/api-cartesian/application/models"

type IAuthenticationService interface {
	GetAuthentication(auth models.Authentication) (models.Authentication, error)
	FetchAuthentication(auth models.Authentication) (models.Authentication, error)
}

type IAuthenticationReaderDB interface {
	GetAuthentication(auth models.Authentication) (models.Authentication, error)
}

type IAuthenticationWriterDB interface {
}

type IAuthenticationReaderCache interface {
	GetAuthentication(auth models.Authentication) (models.Authentication, error)
	FetchAuthentication(auth models.Authentication) (models.Authentication, error)
}

type IAuthenticationWriterCache interface {
	SetAutentication(auth models.Authentication) error
	SetAuthenticationByID(auth models.Authentication) error
}

type IAuthenticationPersistenceDB interface {
	IAuthenticationReaderDB
	IAuthenticationWriterDB
}

type IAuthenticationPersistenceCache interface {
	IAuthenticationReaderCache
	IAuthenticationWriterCache
}
