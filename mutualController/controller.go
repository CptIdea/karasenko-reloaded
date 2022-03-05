package mutual

import (
	"fmt"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"karasenko-reloaded/pkg"
)

type Controller struct {
	store Store
	vk    VkUserGetter
}

func NewController(store Store, vk VkUserGetter) *Controller {
	return &Controller{store: store, vk: vk}
}

type Store interface {
	RequestCacheGet(id int, cityId int, sex int) ([]int, error)
	RequestCacheSet(id int, cityId int, sex int, users []int) error
	GetTargetUsers(id int, minimalMutual int, limit int) ([]pkg.User, error)
	VkCacheGet(id int) (pkg.User, error)
	VkCacheSet(id int, cityId int, sex int, name string) error
}

type VkUserGetter interface {
	UsersGet(params api.Params) (response api.UsersGetResponse, err error)
}

func (c *Controller) GetUsersData(ids []int) ([]pkg.User, error) {
	var toCheck []string
	var users []pkg.User
	for _, id := range ids {
		user, err := c.store.VkCacheGet(id)
		if err != nil {
			toCheck = append(toCheck, fmt.Sprint(id))
			continue
		}

		users = append(users, user)
	}

	if len(toCheck) > 0 {
		vkUsers, err := c.vk.UsersGet(params.NewUsersGetBuilder().UserIDs(toCheck).Fields([]string{"sex", "city"}).Params)
		if err != nil {
			return nil, err
		}
		for _, user := range vkUsers {
			newUser := pkg.User{
				VkId:   user.ID,
				Name:   fmt.Sprintf("%s %s", user.FirstName, user.LastName),
				Sex:    user.Sex,
				CityId: user.City.ID,
			}
			users = append(users, newUser)

			err := c.store.VkCacheSet(newUser.VkId, newUser.CityId, newUser.Sex, newUser.Name)
			if err != nil {
				return nil, err
			}
		}
	}

	return users, nil
}

func (c *Controller) GetUserWithMutual(id int, cityId int, sex int) ([]pkg.User, error) {
	ids, err := c.store.RequestCacheGet(id, cityId, sex)
	if err == nil {
		return c.GetUsersData(ids)
	}

	rawUsers, err := c.store.GetTargetUsers(id, 5, 100)
	if err != nil {
		return nil, err
	}
	ids = []int{}
	for _, user := range rawUsers {
		ids = append(ids, user.VkId)
	}

	rawUsers, err = c.GetUsersData(ids)
	if err != nil {
		return nil, err
	}

	var users []pkg.User
	for _, user := range rawUsers {
		if sex != 0 && sex != user.Sex {
			continue
		}
		if cityId != 0 && cityId != user.CityId {
			continue
		}
		users = append(users, user)
	}

	ids = []int{}
	for _, user := range users {
		ids = append(ids, user.VkId)
	}
	err = c.store.RequestCacheSet(id, cityId, sex, ids)
	if err != nil {
		return nil, err
	}

	return users, nil
}
