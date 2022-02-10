package authentication

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/darllantissei/api-cartesian/application/models"
	"github.com/eucatur/go-toolbox/text"
)

func (a *AuthenticationService) getAuthentication(auth models.Authentication) (models.Authentication, error) {

	var (
		errCollector []string
	)

	auth.Password = a.Utils.ParseToSHA256(auth.Password)

	authCache, errCache := a.PersistenceCache.GetAuthentication(auth)

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

		errCache = a.PersistenceCache.SetAutentication(authDB)
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

func (a *AuthenticationService) checkGetAuthentication(auth models.Authentication) error {

	if text.StringIsEmptyOrWhiteSpace(auth.UserName) || text.StringIsEmptyOrWhiteSpace(auth.Password) {
		return errors.New("favor informe usuário e senha")
	}

	return nil

}
