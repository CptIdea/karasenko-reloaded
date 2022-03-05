package controller

import (
	"karasenko-reloaded/collectors/friends"
	"karasenko-reloaded/collectors/groups"
	"log"
)

type Controller struct {
	friends *friends.Collector
	groups  *groups.Collector
	store   Store
}

func NewController(friends *friends.Collector, groups *groups.Collector, store Store) *Controller {
	return &Controller{friends: friends, groups: groups, store: store}
}

type Store interface {
	GetUserWithUncheckedGroups() (int, error)
	GetUserWithUncheckedFriends() (int, error)
}

func (c *Controller) Start() error {
	for {
		user, err := c.store.GetUserWithUncheckedGroups()
		if err != nil {
			log.Println("Not groupsChecked users not found. Collecting new users...")
			newUser, err := c.store.GetUserWithUncheckedFriends()
			if err != nil {
				return err
			}

			collected, err := c.friends.CollectFriendsFrom(newUser)
			if err != nil {
				return err
			}
			log.Println("Collected users:", collected)
			continue
		}

		err = c.groups.GetUserGroups(user)
		if err != nil {
			return err
		}
		log.Println("Checked user", user)
	}
}
