package groups

import (
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"log"
)

type Collector struct {
	store Store
	vk    VkGroupsGetter
}

func NewCollector(store Store, vk VkGroupsGetter) *Collector {
	return &Collector{store: store, vk: vk}
}

type Store interface {
	SetUserGroupsChecked(id int) error
	AddUserGroups(id int, groups []int) error
}

type VkGroupsGetter interface {
	UsersGetSubscriptions(params api.Params) (response api.UsersGetSubscriptionsResponse, err error)
}

func (c *Collector) GetUserGroups(id int) error {

	defer func(store Store, id int) {
		err := store.SetUserGroupsChecked(id)
		if err != nil {
			log.Println(err)
		}
	}(c.store, id)
	resp, err := c.vk.UsersGetSubscriptions(params.NewUsersGetSubscriptionsBuilder().UserID(id).Params)
	if err != nil {
		log.Println(id, err)
		return nil
	}

	err = c.store.AddUserGroups(id, resp.Groups.Items)
	if err != nil {
		return err
	}

	return nil
}
