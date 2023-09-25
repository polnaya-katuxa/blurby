package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	mycontext "git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/context"

	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/interfaces"
	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/models"
	openapi "git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/server"
	"github.com/google/uuid"
)

const (
	format = "2006-01-02"
	male   = "male"
)

type Server struct {
	openapi.DefaultApiService
	profileLogic   interfaces.ProfileLogic
	eventLogic     interfaces.EventLogic
	eventTypeLogic interfaces.EventTypeLogic
	filterLogic    interfaces.FilterLogic
	adLogic        interfaces.AdLogic
	clientLogic    interfaces.ClientLogic
	statsLogic     interfaces.StatsLogic
}

func New(pl interfaces.ProfileLogic, el interfaces.EventLogic, fl interfaces.FilterLogic,
	al interfaces.AdLogic, cl interfaces.ClientLogic, sl interfaces.StatsLogic, etl interfaces.EventTypeLogic,
) *Server {
	return &Server{
		profileLogic:   pl,
		eventLogic:     el,
		filterLogic:    fl,
		adLogic:        al,
		clientLogic:    cl,
		statsLogic:     sl,
		eventTypeLogic: etl,
	}
}

func (s *Server) Login(ctx context.Context, req openapi.AuthRequest) (openapi.ImplResponse, error) {
	token, err := s.profileLogic.Login(ctx, req.Login, req.Password)
	if err != nil {
		return toErrorResponse(err, "Cannot login.")
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.AuthResponse{Token: token},
	}, nil
}

func (s *Server) Register(ctx context.Context, req openapi.AuthRequest) (openapi.ImplResponse, error) {
	token, err := s.profileLogic.Register(ctx, &models.User{Login: req.Login}, req.Password)
	if err != nil {
		return toErrorResponse(err, "Cannot register..")
	}

	return openapi.ImplResponse{
		Code: http.StatusCreated,
		Body: openapi.AuthResponse{Token: token},
	}, nil
}

func (s *Server) CreateEvent(ctx context.Context, req openapi.CreateEventRequest) (openapi.ImplResponse, error) {
	clientID, err := uuid.Parse(req.Event.ClientID)
	if err != nil {
		return toErrorResponse(err, "Invalid data.")
	}

	err = s.eventLogic.Create(ctx, &models.Event{
		ClientID: clientID,
		Alias:    req.Event.Alias,
	})
	if err != nil {
		return toErrorResponse(err, "Cannot create event.")
	}

	return openapi.ImplResponse{
		Code: http.StatusAccepted,
		Body: openapi.CreateEventResponse{Passed: true},
	}, nil
}

func validateFilters(f []openapi.Filter) ([]models.Filter, error) {
	filters := make([]models.Filter, len(f))
	for i, v := range f {
		var field models.FieldType
		var cmp models.CmpType

		ft := models.ByField
		if v.Type == string(models.ByEvent) {
			ft = models.ByEvent

			if v.Filter.Alias == "" {
				return nil, fmt.Errorf("parse alias")
			}
			if v.Filter.Rate < 1 {
				return nil, fmt.Errorf("parse rate")
			}
		} else {
			switch v.Filter.Field {
			case string(models.ByName):
				field = models.ByName
			case string(models.ByBD):
				field = models.ByBD
			case string(models.ByGender):
				field = models.ByGender
			case string(models.ByAge):
				field = models.ByAge
			default:
				return nil, fmt.Errorf("parse field")
			}
			switch v.Filter.Cmp {
			case string(models.Greater):
				cmp = models.Greater
			case string(models.Less):
				cmp = models.Less
			case string(models.Equal):
				cmp = models.Equal
			case string(models.Between):
				cmp = models.Between
			default:
				return nil, fmt.Errorf("parse cmp")
			}

			if v.Filter.Value1 == "" {
				return nil, fmt.Errorf("parse value #1")
			}

			if cmp == models.Between && v.Filter.Value2 == "" {
				return nil, fmt.Errorf("parse value #2")
			}
		}

		var d time.Duration
		if v.Filter.Span != "" {
			var err error
			d, err = time.ParseDuration(v.Filter.Span)
			if err != nil {
				return nil, fmt.Errorf("parse filter span: %w", err)
			}
		}

		f := models.Filter{
			Type: ft,
			EventFilter: &models.EventFilter{
				Alias: v.Filter.Alias,
				Span:  d,
				Rate:  int(v.Filter.Rate),
			},
			FieldFilter: &models.FieldFilter{
				Field:  field,
				Cmp:    cmp,
				Value1: v.Filter.Value1,
				Value2: v.Filter.Value2,
			},
		}

		filters[i] = f
	}

	return filters, nil
}

func (s *Server) Filter(ctx context.Context, req openapi.FilterRequest) (openapi.ImplResponse, error) {
	filters, err := validateFilters(req.Filters)
	if err != nil {
		return toErrorResponse(err, "Invalid data.")
	}

	//filters := make([]models.Filter, len(req.Filters))
	//for i, v := range req.Filters {
	//	var field models.FieldType
	//	var cmp models.CmpType
	//
	//	ft := models.ByField
	//	if v.Type == string(models.ByEvent) {
	//		ft = models.ByEvent
	//	} else {
	//		switch v.Filter.Field {
	//		case string(models.ByName):
	//			field = models.ByName
	//		case string(models.ByBD):
	//			field = models.ByBD
	//		case string(models.ByGender):
	//			field = models.ByGender
	//		case string(models.ByAge):
	//			field = models.ByAge
	//		default:
	//			return toErrorResponse(fmt.Errorf("parse field type"), "Invalid data.")
	//		}
	//		switch v.Filter.Cmp {
	//		case string(models.Greater):
	//			cmp = models.Greater
	//		case string(models.Less):
	//			cmp = models.Less
	//		case string(models.Equal):
	//			cmp = models.Equal
	//		case string(models.Between):
	//			cmp = models.Between
	//		default:
	//			return toErrorResponse(fmt.Errorf("parse cmp type"), "Invalid data.")
	//		}
	//	}
	//
	//	var d time.Duration
	//	if v.Filter.Span != "" {
	//		var err error
	//		d, err = time.ParseDuration(v.Filter.Span)
	//		if err != nil {
	//			return toErrorResponse(err, "Invalid data.")
	//		}
	//	}
	//
	//	f := models.Filter{
	//		Type: ft,
	//		EventFilter: &models.EventFilter{
	//			Alias: v.Filter.Alias,
	//			Span:  d,
	//			Rate:  int(v.Filter.Rate),
	//		},
	//		FieldFilter: &models.FieldFilter{
	//			Field:  field,
	//			Cmp:    cmp,
	//			Value1: v.Filter.Value1,
	//			Value2: v.Filter.Value2,
	//		},
	//	}
	//
	//	filters[i] = f
	//}

	uuids, err := s.filterLogic.Filter(ctx, filters)
	if err != nil {
		return toErrorResponse(err, "Invalid data.")
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.FilterResponse{Count: int32(len(uuids))},
	}, nil
}

func (s *Server) CreateAd(ctx context.Context, req openapi.CreateAdRequest) (openapi.ImplResponse, error) {
	userID, err := uuid.Parse(req.Ad.UserID)
	if err != nil {
		return toErrorResponse(err, "Invalid data.")
	}

	schSpan, err := time.ParseDuration(req.Ad.Schedule.Span)
	if err != nil {
		return toErrorResponse(err, "Invalid data.")
	}

	fil, err := validateFilters(req.Ad.Filters)
	if err != nil {
		return toErrorResponse(err, "Invalid data.")
	}

	//fil := make([]models.Filter, 0, len(req.Ad.Filters))
	//for _, f := range req.Ad.Filters {
	//	span, err := time.ParseDuration(f.Filter.Span)
	//	if err != nil {
	//		return toErrorResponse(err, "Invalid data.")
	//	}
	//
	//	var c models.CmpType
	//	var l models.FieldType
	//	var t models.FilterType
	//
	//	if f.Type == "by event" {
	//		t = models.ByEvent
	//	} else {
	//		t = models.ByField
	//		switch f.Filter.Field {
	//		case "age":
	//			l = models.ByAge
	//		case "gender":
	//			l = models.ByGender
	//		case "birth_date":
	//			l = models.ByBD
	//		case "name":
	//			l = models.ByName
	//		default:
	//			return toErrorResponse(fmt.Errorf("parse field type"), "Invalid data.")
	//		}
	//
	//		switch f.Filter.Cmp {
	//		case ">":
	//			c = models.Greater
	//		case "<":
	//			c = models.Less
	//		case "=":
	//			c = models.Equal
	//		case "between":
	//			c = models.Between
	//		default:
	//			return toErrorResponse(fmt.Errorf("parse cmp type"), "Invalid data.")
	//		}
	//	}
	//
	//	fil = append(fil, models.Filter{
	//		Type: t,
	//		EventFilter: &models.EventFilter{
	//			Alias: f.Filter.Alias,
	//			Span:  span,
	//			Rate:  int(f.Filter.Rate),
	//		},
	//		FieldFilter: &models.FieldFilter{
	//			Field:  l,
	//			Cmp:    c,
	//			Value1: f.Filter.Value1,
	//			Value2: f.Filter.Value2,
	//		},
	//	})
	//}

	err = s.adLogic.Create(ctx, &models.Ad{
		Content: req.Ad.Content,
		Filters: fil,
		UserID:  userID,
		Schedule: &models.Schedule{
			Finished: false,
			Periodic: req.Ad.Schedule.Periodic,
			Span:     schSpan,
		},
	})
	if err != nil {
		return toErrorResponse(err, "Cannot create ad.")
	}

	return openapi.ImplResponse{
		Code: http.StatusCreated,
		Body: openapi.CreateAdResponse{Created: true},
	}, nil
}

func (s *Server) CreateClient(ctx context.Context, req openapi.CreateClientRequest) (openapi.ImplResponse, error) {
	bd, err := time.Parse(time.RFC3339, req.Client.BirthDate)
	if err != nil {
		return toErrorResponse(err, "Invalid data.")
	}

	rd, err := time.Parse(time.RFC3339, req.Client.RegistrationDate)
	if err != nil {
		return toErrorResponse(err, "Invalid data.")
	}

	id, err := uuid.Parse(req.Client.Id)
	if err != nil {
		return toErrorResponse(err, "Invalid data.")
	}

	var g models.Gender
	if req.Client.Gender == male {
		g = models.Male
	} else {
		g = models.Female
	}

	err = s.clientLogic.Create(ctx, &models.Client{
		ID:               id,
		Name:             req.Client.Name,
		Surname:          req.Client.Surname,
		Patronymic:       req.Client.Patronymic,
		Gender:           g,
		BirthDate:        bd,
		RegistrationDate: rd,
		Email:            req.Client.Email,
		Data:             req.Client.Data,
	})
	if err != nil {
		return toErrorResponse(err, "Cannot create client.")
	}

	return openapi.ImplResponse{
		Code: http.StatusCreated,
		Body: openapi.CreateClientResponse{Created: true},
	}, nil
}

func (s *Server) DeleteClient(ctx context.Context, req openapi.DeleteClientRequest) (openapi.ImplResponse, error) {
	clientID, err := uuid.Parse(req.Id)
	if err != nil {
		return toErrorResponse(err, "Invalid data.")
	}

	err = s.clientLogic.Delete(ctx, clientID)
	if err != nil {
		return toErrorResponse(err, "Cannot delete client.")
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.DeleteClientResponse{Deleted: true},
	}, nil
}

func (s *Server) DeleteUser(ctx context.Context, req openapi.DeleteUserRequest) (openapi.ImplResponse, error) {
	err := s.profileLogic.Delete(ctx, req.Login)
	if err != nil {
		return toErrorResponse(err, "Cannot delete client.")
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.DeleteUserResponse{Deleted: true},
	}, nil
}

func (s *Server) GetAds(ctx context.Context) (openapi.ImplResponse, error) {
	ads, err := s.adLogic.GetAll(ctx)
	if err != nil {
		return toErrorResponse(err, "Cannot get ads.")
	}

	adsOA := make([]openapi.Ad, 0, len(ads))
	for _, a := range ads {
		fil := make([]openapi.Filter, 0, len(a.Filters))
		for _, f := range a.Filters {
			var t string
			if f.Type == models.ByEvent {
				t = "by event"
			} else {
				t = "by field"
			}

			var l string
			switch f.FieldFilter.Field {
			case models.ByAge:
				l = "age"
			case models.ByGender:
				l = "gender"
			case models.ByBD:
				l = "birth_date"
			case models.ByName:
				l = "name"
			default:
				l = ""
			}

			var c string
			switch f.FieldFilter.Cmp {
			case models.Greater:
				c = ">"
			case models.Less:
				c = "<"
			case models.Equal:
				c = "="
			case models.Between:
				c = "between"
			default:
				c = ""
			}

			fil = append(fil, openapi.Filter{
				Type: t,
				Filter: openapi.FilterFilter{
					Alias:  f.EventFilter.Alias,
					Span:   f.EventFilter.Span.String(),
					Rate:   int32(f.EventFilter.Rate),
					Field:  l,
					Cmp:    c,
					Value1: f.FieldFilter.Value1,
					Value2: f.FieldFilter.Value2,
				},
			})
		}

		adsOA = append(adsOA, openapi.Ad{
			Content:    a.Content,
			Filters:    fil,
			CreateTime: a.CreateTime.Format(time.RFC3339),
			UserID:     a.UserID.String(),
			Schedule: openapi.Schedule{
				Periodic: a.Schedule.Periodic,
				Finished: a.Schedule.Finished,
				NextTime: a.Schedule.NextTime.Format(time.RFC3339),
				Span:     a.Schedule.Span.String(),
			},
		})
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.GetAdsResponse{Ads: adsOA},
	}, nil
}

func (s *Server) GetClient(ctx context.Context, req openapi.GetClientRequest) (openapi.ImplResponse, error) {
	clientID, err := uuid.Parse(req.Id)
	if err != nil {
		return toErrorResponse(err, "Invalid data.")
	}

	c, err := s.clientLogic.Get(ctx, clientID)
	if err != nil {
		return toErrorResponse(err, "Cannot get client.")
	}

	var g string
	if c.Gender == models.Male {
		g = male
	} else {
		g = "female"
	}

	cOA := openapi.Client{
		Id:               c.ID.String(),
		Name:             c.Name,
		Surname:          c.Surname,
		Patronymic:       c.Patronymic,
		Gender:           g,
		BirthDate:        c.BirthDate.Format(time.RFC3339),
		RegistrationDate: c.RegistrationDate.Format(time.RFC3339),
		Email:            c.Email,
		Data:             c.Data,
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.GetClientResponse{Client: cOA},
	}, nil
}

func (s *Server) GetClients(ctx context.Context) (openapi.ImplResponse, error) {
	clients, err := s.clientLogic.GetAll(ctx)
	if err != nil {
		return toErrorResponse(err, "Cannot get clients.")
	}

	cOA := make([]openapi.Client, 0, len(clients))
	for _, c := range clients {
		var g string
		if c.Gender == models.Male {
			g = male
		} else {
			g = "female"
		}

		cOA = append(cOA, openapi.Client{
			Id:               c.ID.String(),
			Name:             c.Name,
			Surname:          c.Surname,
			Patronymic:       c.Patronymic,
			Gender:           g,
			BirthDate:        c.BirthDate.Format(time.RFC3339),
			RegistrationDate: c.RegistrationDate.Format(time.RFC3339),
			Email:            c.Email,
			Data:             c.Data,
		})
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.GetClientsResponse{Clients: cOA},
	}, nil
}

func (s *Server) GetUser(ctx context.Context, req openapi.GetUserRequest) (openapi.ImplResponse, error) {
	u, err := s.profileLogic.GetByLogin(ctx, req.Login)
	if err != nil {
		return toErrorResponse(err, "Cannot get user.")
	}

	uOA := openapi.User{
		Id:      u.ID.String(),
		Login:   u.Login,
		IsAdmin: u.IsAdmin,
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.GetUserResponse{User: uOA},
	}, nil
}

func (s *Server) GetUserByID(ctx context.Context, req openapi.GetUserByIdRequest) (openapi.ImplResponse, error) {
	userID, err := uuid.Parse(req.Id)
	if err != nil {
		return toErrorResponse(err, "Invalid data.")
	}

	u, err := s.profileLogic.GetByID(ctx, userID)
	if err != nil {
		return toErrorResponse(err, "Cannot get user.")
	}

	uOA := openapi.User{
		Id:      u.ID.String(),
		Login:   u.Login,
		IsAdmin: u.IsAdmin,
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.GetUserResponse{User: uOA},
	}, nil
}

func (s *Server) GetUsers(ctx context.Context) (openapi.ImplResponse, error) {
	users, err := s.profileLogic.GetAll(ctx)
	if err != nil {
		return toErrorResponse(err, "Cannot get users.")
	}

	uOA := make([]openapi.User, 0, len(users))
	for _, u := range users {
		uOA = append(uOA, openapi.User{
			Id:      u.ID.String(),
			Login:   u.Login,
			IsAdmin: u.IsAdmin,
		})
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.GetUsersResponse{Users: uOA},
	}, nil
}

func (s *Server) GrantUserAdmin(ctx context.Context, req openapi.GrantAdminRequest) (openapi.ImplResponse, error) {
	err := s.profileLogic.GrantAdmin(ctx, req.Login)
	if err != nil {
		return toErrorResponse(err, "Cannot grant admin.")
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.GrantAdminResponse{Granted: true},
	}, nil
}

func (s *Server) UserInfo(ctx context.Context) (openapi.ImplResponse, error) {
	user, err := mycontext.UserFromContext(ctx)
	if err != nil {
		return toErrorResponse(err, "User is not authorized.")
	}

	u := openapi.User{
		Id:      user.ID.String(),
		Login:   user.Login,
		IsAdmin: user.IsAdmin,
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.UserInfoResponse{User: u},
	}, nil
}

func (s *Server) GetClientStats(ctx context.Context) (openapi.ImplResponse, error) {
	statsC, err := s.statsLogic.GetClientStats(ctx)
	if err != nil {
		return toErrorResponse(err, "Cannot get client stats.")
	}

	statsCOA := openapi.ClientStats{
		Num:    int32(statsC.Num),
		AvgAge: int32(statsC.AvgAge),
	}

	statsA, err := s.statsLogic.GetAdStats(ctx)
	if err != nil {
		return toErrorResponse(err, "Cannot get ad stats.")
	}

	statsAOA := make([]openapi.AdStat, 0, len(statsA))
	for _, v := range statsA {
		statAOA := openapi.AdStat{
			Num:  int32(v.Num),
			Date: v.Date.Format(format),
		}

		statsAOA = append(statsAOA, statAOA)
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.GetClientStatsResponse{ClientStats: statsCOA, AdStats: statsAOA},
	}, nil
}

func (s *Server) CreateEventType(ctx context.Context, _ string, req openapi.CreateEventTypeRequest) (openapi.ImplResponse, error) {
	et := &models.EventType{
		Name:  req.EventType.Name,
		Alias: req.EventType.Alias,
	}

	err := s.eventTypeLogic.Create(ctx, et)
	if err != nil {
		return toErrorResponse(err, "Cannot create event type.")
	}

	return openapi.ImplResponse{
		Code: http.StatusCreated,
		Body: openapi.CreateEventTypeResponse{Created: true},
	}, nil
}

func (s *Server) GetEventTypes(ctx context.Context) (openapi.ImplResponse, error) {
	ets, err := s.eventTypeLogic.GetAll(ctx)
	if err != nil {
		return toErrorResponse(err, "Cannot get event types. Create event types before making ads.")
	}

	etOA := make([]openapi.EventType, 0, len(ets))
	for _, u := range ets {
		etOA = append(etOA, openapi.EventType{
			Id:    u.ID.String(),
			Name:  u.Name,
			Alias: u.Alias,
		})
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.GetEventTypesResponse{EventTypes: etOA},
	}, nil
}
