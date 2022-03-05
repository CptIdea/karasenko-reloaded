package friends

import (
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"log"
)

type Collector struct {
	vk    VkFriendsGetter
	store Store
}

func NewCollector(vk VkFriendsGetter, store Store) *Collector {
	return &Collector{vk: vk, store: store}
}

type Store interface {
	AddUserFriends(id int, friends []int) error
	SetUserFriendsChecked(id int) error
	GetUserFriends(id int) ([]int, error)
}

type VkFriendsGetter interface {
	FriendsGet(params api.Params) (response api.FriendsGetResponse, err error)
}

func (c *Collector) CollectFriendsFrom(id int) (int, error) {
	friends, err := c.store.GetUserFriends(id)
	if err != nil {
		return 0, err
	}
	defer func(store Store, id int) {
		err := store.SetUserFriendsChecked(id)
		if err != nil {
			log.Println(err)
		}
	}(c.store, id)

	var counter int

	for _, friend := range friends {
		if _, err := c.store.GetUserFriends(friend); err != nil {
			newFriends, err := c.vk.FriendsGet(params.NewFriendsGetBuilder().UserID(friend).Params)
			if err != nil {
				log.Println(friend, err)
			} else {
				log.Println("Getted friends from user", friend)
				counter++
			}

			err = c.store.AddUserFriends(friend, newFriends.Items)
			if err != nil {
				return 0, err
			}
		}

	}

	return counter, nil
}
