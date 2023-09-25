package postgres

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	my_errors "git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/errors"
	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/models"
	repoModels "git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/repository/postgres/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ClientRepository struct {
	dbA *gorm.DB
	dbT *gorm.DB
	dbC *gorm.DB
}

func NewCR(dbA *gorm.DB, dbT *gorm.DB, dbC *gorm.DB) *ClientRepository {
	return &ClientRepository{
		dbA: dbA,
		dbT: dbT,
		dbC: dbC,
	}
}

func (c *ClientRepository) Create(ctx context.Context, client *models.Client) error {
	data, err := json.Marshal(client.Data)
	if err != nil {
		return fmt.Errorf("json data marshall: %w", err)
	}

	clientDB := &repoModels.Client{
		UUID:             client.ID,
		Name:             client.Name,
		Surname:          client.Surname,
		Patronymic:       client.Patronymic,
		Gender:           client.Gender,
		BirthDate:        client.BirthDate.UTC(),
		RegistrationDate: client.RegistrationDate.UTC(),
		Email:            client.Email,
		Data:             string(data),
	}

	res := c.dbC.WithContext(ctx).Table("clients").Create(clientDB)
	if res.Error != nil {
		return fmt.Errorf("insert: %w", res.Error)
	}

	return nil
}

func (c *ClientRepository) Delete(ctx context.Context, id uuid.UUID) error {
	res := c.dbA.WithContext(ctx).Table("clients").Where("uuid = ?", id).Delete(&repoModels.Client{})
	if res.Error != nil {
		return fmt.Errorf("delete: %w", res.Error)
	}

	return nil
}

func (c *ClientRepository) GetAll(ctx context.Context) ([]*models.Client, error) {
	var clientsDB []*repoModels.Client

	res := c.dbT.WithContext(ctx).Table("clients").Find(&clientsDB)
	if res.Error != nil {
		return nil, fmt.Errorf("select: %w", res.Error)
	}

	clientsLogic := make([]*models.Client, 0, len(clientsDB))
	for _, c := range clientsDB {
		var data map[string]string
		err := json.Unmarshal([]byte(c.Data), &data)
		if err != nil {
			return nil, fmt.Errorf("json unmarshall data: %w", err)
		}

		clientsLogic = append(clientsLogic, &models.Client{
			ID:               c.UUID,
			Name:             c.Name,
			Surname:          c.Surname,
			Patronymic:       c.Patronymic,
			Gender:           c.Gender,
			BirthDate:        c.BirthDate,
			RegistrationDate: c.RegistrationDate,
			Email:            c.Email,
			Data:             data,
		})
	}

	return clientsLogic, nil
}

func (c *ClientRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Client, error) {
	client := repoModels.Client{}

	res := c.dbT.WithContext(ctx).Table("clients").Where("uuid = ?", id).Take(&client)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("select: %w", my_errors.ErrUserNotFound)
	}
	if res.Error != nil {
		return nil, fmt.Errorf("select: %w", res.Error)
	}

	var data map[string]string
	if client.Data != "" {
		err := json.Unmarshal([]byte(client.Data), &data)
		if err != nil {
			return nil, fmt.Errorf("json unmarshal data: %w", err)
		}
	}

	resClient := models.Client{
		ID:               client.UUID,
		Name:             client.Name,
		Surname:          client.Surname,
		Patronymic:       client.Patronymic,
		Gender:           client.Gender,
		BirthDate:        client.BirthDate,
		RegistrationDate: client.RegistrationDate,
		Email:            client.Email,
		Data:             data,
	}

	return &resClient, nil
}

func (c *ClientRepository) GetByIDs(ctx context.Context, ids []uuid.UUID) ([]*models.Client, error) {
	var clientsDB []*repoModels.Client

	res := c.dbT.WithContext(ctx).Table("clients").Where("uuid in ?", ids).Find(&clientsDB)
	if res.Error != nil {
		return nil, fmt.Errorf("select: %w", res.Error)
	}

	clientsLogic := make([]*models.Client, 0, len(clientsDB))
	for _, c := range clientsDB {
		var data map[string]string
		err := json.Unmarshal([]byte(c.Data), &data)
		if err != nil {
			return nil, fmt.Errorf("json unmarshall data: %w", err)
		}

		clientsLogic = append(clientsLogic, &models.Client{
			ID:               c.UUID,
			Name:             c.Name,
			Surname:          c.Surname,
			Patronymic:       c.Patronymic,
			Gender:           c.Gender,
			BirthDate:        c.BirthDate,
			RegistrationDate: c.RegistrationDate,
			Email:            c.Email,
			Data:             data,
		})
	}

	return clientsLogic, nil
}

func (c *ClientRepository) Filter(ctx context.Context, f []*models.FieldFilter) ([]uuid.UUID, error) {
	if len(f) == 0 {
		var uuidsStr []string

		res := c.dbT.WithContext(ctx).Table("clients").Distinct("uuid").Find(&uuidsStr)
		if res.Error != nil {
			return nil, fmt.Errorf("select: %w", res.Error)
		}

		uuids := make([]uuid.UUID, 0, len(uuidsStr))
		for _, u := range uuidsStr {
			uid, err := uuid.Parse(u)
			if err != nil {
				return nil, fmt.Errorf("parse uuid: %w", err)
			}

			uuids = append(uuids, uid)
		}

		return uuids, nil
	}

	query := make([]string, len(f))

	for i, v := range f {
		switch v.Field {
		case models.ByAge:
			query[i] += "extract(years from age(now(), birth_date))"
			switch v.Cmp {
			case models.Equal, models.Greater, models.Less:
				query[i] += " " + string(v.Cmp) + " " + v.Value1
			case models.Between:
				query[i] += " " + string(v.Cmp) + " " + v.Value1 + " and " + v.Value2
			default:
				return nil, fmt.Errorf("no matching filters: %w", my_errors.ErrFilter)
			}
		case models.ByBD:
			query[i] += "birth_date"
			switch v.Cmp {
			case models.Equal, models.Greater, models.Less:
				query[i] += " " + string(v.Cmp) + " '" + v.Value1 + "'"
			case models.Between:
				query[i] += " " + string(v.Cmp) + " '" + v.Value1 + "' and '" + v.Value2 + "'"
			default:
				return nil, fmt.Errorf("no matching filters: %w", my_errors.ErrFilter)
			}
		case models.ByGender:
			query[i] += "gender"
			if v.Cmp == models.Equal {
				query[i] += " = '" + v.Value1 + "'"
			} else {
				return nil, fmt.Errorf("no matching filters: %w", my_errors.ErrFilter)
			}
		case models.ByName:
			query[i] += "name"
			if v.Cmp == models.Equal {
				query[i] += " = '" + v.Value1 + "'"
			} else {
				return nil, fmt.Errorf("no matching filters: %w", my_errors.ErrFilter)
			}
		default:
			return nil, fmt.Errorf("no matching filters: %w", my_errors.ErrFilter)
		}
	}

	queryLine := strings.Join(query, " and ")

	var uuidsStr []string

	res := c.dbT.WithContext(ctx).Table("clients").Distinct("uuid").Where(queryLine).Find(&uuidsStr)
	if res.Error != nil {
		return nil, fmt.Errorf("select: %w", res.Error)
	}

	uuids := make([]uuid.UUID, 0, len(uuidsStr))
	for _, u := range uuidsStr {
		uid, err := uuid.Parse(u)
		if err != nil {
			return nil, fmt.Errorf("parse uuid: %w", err)
		}

		uuids = append(uuids, uid)
	}

	return uuids, nil
}

func (c *ClientRepository) GetStats(ctx context.Context) (*models.ClientStat, error) {
	var clientStats *repoModels.ClientStat

	res := c.dbT.WithContext(ctx).Table("clients").Select("count(*) as num, extract( years from avg(age(birth_date))) as avg_age").Find(&clientStats)
	if res.Error != nil {
		return nil, fmt.Errorf("select: %w", res.Error)
	}

	clientStatsL := &models.ClientStat{
		Num:    clientStats.Num,
		AvgAge: clientStats.AvgAge,
	}

	return clientStatsL, nil
}
