package service

import (
	"strconv"
	"time"

	"todolist"
	"todolist/pkg/redisC"
	"todolist/pkg/repository"
)

type ListsService struct {
	repo  repository.TodoList
	cashe *redisC.RedisCashe
}

func NewListService(repo repository.TodoList, cashe *redisC.RedisCashe) *ListsService {
	return &ListsService{
		repo:  repo,
		cashe: cashe,
	}
}

func (s *ListsService) CreateList(userId int, list todolist.TodoList) (int, error) {
	listId, err := s.repo.CreateList(userId, list)
	if err != nil {
		return 0, err
	}
	return listId, nil
}

func (s *ListsService) GetAllLists(userId int) ([]todolist.TodoList, error) {
	value, err := s.cashe.Get("lists:" + strconv.Itoa(userId))
	if err == nil {
		cashedata, ok := value.(map[string]interface{})
		if ok {
			listsc := []todolist.TodoList{}
			list := todolist.TodoList{
				Id:          int(cashedata["id"].(float64)),
				Title:       cashedata["title"].(string),
				Description: cashedata["description"].(string),
			}
			listsc = append(listsc, list)

			return listsc, nil
		}
	}
	lists, err := s.repo.GetAllLists(userId)
	if err != nil {
		return nil, err
	}
	err = s.cashe.Set("lists:"+strconv.Itoa(userId), lists, time.Hour)
	if err != nil {
		return nil, err
	}

	return lists, nil
}

func (s *ListsService) GetListById(userId, listId int) (todolist.TodoList, error) {
	return s.repo.GetListById(userId, listId)
}

func (s *ListsService) Delete(userID, lisId int) error {
	return s.repo.Delete(userID, lisId)
}

func (s *ListsService) UpdateListInout(userid, listID int, input todolist.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.UpdateListInout(userid, listID, input)
}
