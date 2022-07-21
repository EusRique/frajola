package factory

import (
	"github.com/EusRique/frajola/application/usecase"
	"github.com/EusRique/frajola/infrastructure/repository"
	"github.com/jinzhu/gorm"
)

func DocumentUseCaseFactory(database *gorm.DB) usecase.DocumentUseCase {
	documentRepository := repository.DocumentRepositoryDb{Db: database}

	documentUseCase := usecase.DocumentUseCase{
		DocumentRepository: &documentRepository,
	}

	return documentUseCase
}
