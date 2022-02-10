package authentication

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/darllantissei/api-cartesian/application/models"
	"github.com/darllantissei/api-cartesian/application/utils"
)

type AuthenticationService struct {
	PersistenceDB    IAuthenticationPersistenceDB
	PersistenceCache IAuthenticationPersistenceCache
	Utils            utils.IUtilsService
}

func (a *AuthenticationService) GetAuthentication(auth models.Authentication) (models.Authentication, error) {

	var (
		err error
	)

	err = a.checkGetAuthentication(auth)

	if err != nil {
		return models.Authentication{}, err
	}

	auth, err = a.getAuthentication(auth)

	if err != nil {
		return models.Authentication{}, err
	}

	return auth, nil

}

func (a *AuthenticationService) FetchAuthentication(auth models.Authentication) (models.Authentication, error) {

	var (
		errCollector []string
	)

	authCache, errCache := a.PersistenceCache.FetchAuthentication(auth)

	if errCache != nil {
		errCollector = append(errCollector, errCache.Error())
		authCache = models.Authentication{}
	}

	if reflect.DeepEqual(authCache, models.Authentication{}) {

		authDB, errDB := a.PersistenceDB.GetAuthentication(auth)

		if errDB != nil {
			errCollector = append(errCollector, errDB.Error())

			return models.Authentication{}, errors.New(strings.Join(errCollector, ";"))
		}

		if !authDB.Enabled {
			return models.Authentication{}, errors.New("usuário sem permissão de acesso ou não habilitado")
		}

		auth = authDB

		errCache = a.PersistenceCache.SetAuthenticationByID(authDB)
		if errCache != nil {
			errCollector = append(errCollector, fmt.Sprintf("Error ao salvar autenticação em cache. Detalhes: %s", errCache.Error()))
			return models.Authentication{}, errors.New(strings.Join(errCollector, ";"))
		}

	} else {

		auth = authCache

	}

	if !auth.Enabled {
		return models.Authentication{}, errors.New("usuário sem permissão de acesso ou não habilitado")
	}

	return auth, nil

}
