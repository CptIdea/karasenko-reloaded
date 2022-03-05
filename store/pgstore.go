package store

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"karasenko-reloaded/pkg"

	"github.com/jackc/pgx/v4"
)

type Store struct {
	db *pgx.Conn
}

func NewStore(db *pgx.Conn) *Store {
	return &Store{db: db}
}

func (s *Store) AddUserFriends(id int, friends []int) error {
	_, err := s.db.Exec(context.Background(), `insert into user_friends (vk_id, friends) VALUES ($1, $2)`, id, friends)
	return err
}

func (s *Store) SetUserFriendsChecked(id int) error {
	_, err := s.db.Exec(context.Background(), `update user_friends set  friends_checked = true where vk_id=$1;`, id)
	return err
}

func (s *Store) GetUserFriends(id int) ([]int, error) {
	var friends []int
	err := s.db.QueryRow(context.Background(), `select friends from user_friends where vk_id=$1`, id).Scan(&friends)

	return friends, err
}

func (s *Store) GetUserWithUncheckedGroups() (int, error) {
	var user int
	err := s.db.QueryRow(context.Background(), `select vk_id from user_friends where groups_checked=false limit 1`).Scan(&user)

	return user, err
}

func (s *Store) GetUserWithUncheckedFriends() (int, error) {
	var user int
	err := s.db.QueryRow(context.Background(), `select vk_id from user_friends where friends_checked=false limit 1`).Scan(&user)

	return user, err
}

func (s *Store) SetUserGroupsChecked(id int) error {
	_, err := s.db.Exec(context.Background(), `update user_friends set  groups_checked = true where vk_id=$1;`, id)
	return err
}

func (s *Store) AddUserGroups(id int, groups []int) error {
	_, err := s.db.Exec(context.Background(), `insert into user_groups (vk_id, groups) VALUES ($1, $2)`, id, groups)
	return err
}

func (s *Store) RequestCacheGet(id int, cityId int, sex int) ([]int, error) {
	var users []int
	err := s.db.QueryRow(context.Background(), `select * from request_cache where vk_id = $1 and city_id =$2 and sex=$3;`, id, cityId, sex).Scan(&users)
	return users, err
}

func (s *Store) RequestCacheSet(id int, cityId int, sex int, users []int) error {
	_, err := s.db.Exec(context.Background(), `insert into request_cache (vk_id, city_id, vk_ids_list,sex) values ($1,$2,$3,$4);`, id, cityId, users, sex)
	return err
}

func (s *Store) GetTargetUsers(id int, minimalMutual int, limit int) ([]pkg.User, error) {
	var users []pkg.User
	err := pgxscan.Select(context.Background(), s.db, &users,
		`WITH vars (target_groups) as (
values ((select groups
from user_groups
where vk_id = $1))
)
SELECT vk_id,
target_groups & groups as mutual_groups
from user_groups,
vars
where cardinality(target_groups & groups) > $2
order by cardinality(target_groups & groups) desc
limit $3;`,
		id, minimalMutual, limit)
	return users, err
}

func (s *Store) VkCacheGet(id int) (pkg.User, error) {
	var user pkg.User
	err := s.db.QueryRow(context.Background(), `select * from user_vk_cache where vk_id = $1;`, id).Scan(&user)
	return user, err
}

func (s *Store) VkCacheSet(id int, cityId int, sex int, name string) error {
	_, err := s.db.Exec(context.Background(), `insert into user_vk_cache (vk_id, city_id,sex,name) values ($1,$2,$3,$4);`, id, cityId, sex, name)
	return err
}
