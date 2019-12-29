package repositories

import (
	"fmt"
	"github.com/cryptosalamander/dream_factory/dtos"
	"github.com/cryptosalamander/dream_factory/models"
	"math"
	"strings"
)


func (r *DreamRepository) SaveMember(member *models.Member) RepositoryResult {
	err := r.db.Save(member).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: member}
}

func (r *DreamRepository) FindAllMembers() RepositoryResult {
	var members models.Members
	err := r.db.Find(&members).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}
	return RepositoryResult{Result: &members}
}

func (r *DreamRepository) FindMemberById(id string) RepositoryResult {
	var member models.Member

	err := r.db.Where(&models.Member{Member_id: id}).Take(&member).Error
	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &member}
}

func (r *DreamRepository) DeleteMemberById(id string) RepositoryResult {
	err := r.db.Delete(&models.Member{Member_id: id}).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: nil}
}

func (r *DreamRepository) DeleteMembersByIds(ids *[]string) RepositoryResult {
	err := r.db.Where("ID IN (?)", *ids).Delete(&models.Members{}).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: nil}
}

func (r *DreamRepository) MemberPagination(pagination *dtos.Pagination) (RepositoryResult, int) {
	var members models.Members

	totalRows, totalPages, fromRow, toRow := 0, 0, 0, 0

	offset := pagination.Page * pagination.Limit

	// get data with limit, offset & order
	find := r.db.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)

	// generate where query
	searchs := pagination.Searchs

	if searchs != nil {
		for _, value := range searchs {
			column := value.Column
			action := value.Action
			query := value.Query

			switch action {
			case "equals":
				whereQuery := fmt.Sprintf("%s = ?", column)
				find = find.Where(whereQuery, query)
				break
			case "contains":
				whereQuery := fmt.Sprintf("%s LIKE ?", column)
				find = find.Where(whereQuery, "%"+query+"%")
				break
			case "in":
				whereQuery := fmt.Sprintf("%s IN (?)", column)
				queryArray := strings.Split(query, ",")
				find = find.Where(whereQuery, queryArray)
				break
			}
		}
	}

	find = find.Find(&members)

	// has error find data
	errFind := find.Error

	if errFind != nil {
		return RepositoryResult{Error: errFind}, totalPages
	}

	pagination.Rows = members

	// count all data
	errCount := r.db.Model(&models.Member{}).Count(&totalRows).Error

	if errCount != nil {
		return RepositoryResult{Error: errCount}, totalPages
	}

	pagination.TotalRows = totalRows

	// calculate total pages
	totalPages = int(math.Ceil(float64(totalRows)/float64(pagination.Limit))) - 1

	if pagination.Page == 0 {
		// set from & to row on first page
		fromRow = 1
		toRow = pagination.Limit
	} else {
		if pagination.Page <= totalPages {
			// calculate from & to row
			fromRow = pagination.Page*pagination.Limit + 1
			toRow = (pagination.Page + 1) * pagination.Limit
		}
	}

	if toRow > totalRows {
		// set to row with total rows
		toRow = totalRows
	}

	pagination.FromRow = fromRow
	pagination.ToRow = toRow

	return RepositoryResult{Result: pagination}, totalPages
}
