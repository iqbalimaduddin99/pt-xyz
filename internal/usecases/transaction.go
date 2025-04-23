package usecases

import (
	"database/sql"
	"fmt"
	"pt-xyz/internal/entities"
	"pt-xyz/internal/repository"
	"pt-xyz/pkg"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)


type ServiceTransaction struct {
	db *sqlx.DB
	repo repository.RepositoryTransaction
	repoTransactionProduct repository.RepositoryTransactionProduct
	repoLoanLimit repository.RepositoryLoanLimit
	repoLoanInstallment repository.RepositoryLoanInstallment
	repoMasterProductXYZ repository.RepositoryMasterProductXYZ
}

func NewServiceTransaction(db *sqlx.DB, repo repository.RepositoryTransaction, repoTransactionProduct repository.RepositoryTransactionProduct,  repoLoanLimit repository.RepositoryLoanLimit, repoLoanInstallment repository.RepositoryLoanInstallment, repoMasterProductXYZ repository.RepositoryMasterProductXYZ) *ServiceTransaction {
	return &ServiceTransaction{db: db, repo: repo, repoTransactionProduct: repoTransactionProduct, repoLoanLimit: repoLoanLimit, repoLoanInstallment: repoLoanInstallment, repoMasterProductXYZ: repoMasterProductXYZ}
}

func (s *ServiceTransaction) CreateTransaction(transaction *entities.TransactionTableReq,claims *pkg.Claims) (string, error) {
	
	tx, err := s.db.Beginx()
	if err != nil {
		// log.Println("tx failed")
	}


	// _, err := s.db.Exec("START TRANSACTION;")
    // if err != nil {
	// 	return "", err
    // }


	if err != nil {
		return "", err
	} 

	defer func() {
        if p := recover(); p != nil {
            tx.Rollback()
            panic(p) 
        } else if err != nil {
            tx.Rollback()
        }
    }()
	var totalPrice float64
	var transactionTbl entities.TransactionTable

	if (!transaction.IsExternalCompany) {
		
		transactionTbl.CompanyID = 111111 //internal code
		transactionTbl.ID = uuid.Nil
		
			for index, product := range transaction.TransactionProducts {
				productData, err := s.repoMasterProductXYZ.GetMasterProductForTransactionById(tx, product.ProductCompanyID)
				if err != nil {
					return "", err
				}
				if (productData.Stock == 0) {
					tx.Rollback()
					return "", fmt.Errorf("Run out of stock")
				}
				if (index == 0) {
					transactionTbl.ConsumerID = claims.ID
					transactionTbl.CompanyName = productData.CompanyName
					transactionTbl.CompanyCategory = productData.CompanyCategory
					transactionTbl.ContactNumber = productData.ContactNumber
					transactionTbl.AdminFee = productData.AdminFee
					idTransaction, err := s.repo.CreateTransaction(tx, &transactionTbl)

					if err != nil {
						return "", err
					}
					transactionTbl.ID = idTransaction
				}

				product.AssetName = productData.AssetName 
				product.ProductCompanyID = productData.ID
				product.Price = productData.Price 
				
				product.OTR = productData.OTR 
				product.TransactionID = transactionTbl.ID
				
				err = s.repoTransactionProduct.CreateTransactionProduct(tx, &product)

				totalPrice += productData.Price
			}

	} else {
		for _, product := range transaction.TransactionProducts {
			totalPrice += product.Price 
		}
	}

	transaction.TotalPrice = totalPrice

	existsLimitLoan, err := s.repoLoanLimit.GetLoanLimitByIDTransaction(tx, claims.ID)
	if err != nil  {
		if (err == sql.ErrNoRows) {
				return "", fmt.Errorf("You don't have limit")
		}
			return "", err
    }
	if existsLimitLoan == nil  {
		tx.Rollback();
		return "", fmt.Errorf("You don't have limit")
    }

	if (existsLimitLoan != nil && existsLimitLoan.LimitLoan < totalPrice) {
		tx.Rollback();
		return "", fmt.Errorf("You don't have limit")
	}
	

	installmentAmount := totalPrice / float64(existsLimitLoan.TenorAmount)

	interestRate := 2.0 / 100 * installmentAmount

	currentDate := time.Now()

	for i := 0; i < existsLimitLoan.TenorAmount; i++ {
		loanInstallment := entities.LoanInstallment{
			TransactionID:   transactionTbl.ID,
			ConsumerID:      claims.ID,
			InstallmentAmount: installmentAmount,
			InterestRate:    interestRate,
			Tenor:           currentDate.AddDate(0, i, 0), 
		}
		err = s.repoLoanInstallment.CreateLoanInstallment(tx, &loanInstallment)
		if err != nil {
			return "", err
		}
	}
	//Test pingin error aja 
	// return "", errors.New("Pinginn error aja")

	// _, commitErr := s.db.Exec("COMMIT;")
    // if commitErr != nil {
	// 	return "", err
    // }
	
    err = tx.Commit();
	return "Succes Create Transaction", nil

}


