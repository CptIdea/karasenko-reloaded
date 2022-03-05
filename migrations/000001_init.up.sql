create table user_groups
(
    vk_id  integer   not null
        constraint user_groups_pk
            primary key,
    groups integer[] not null
);

alter table user_groups
    owner to postgres;

create unique index user_groups_vk_id_uindex
    on user_groups (vk_id);

create index user_groups_groups_index
    on user_groups using gin (groups);

create table user_friends
(
    vk_id           integer               not null
        constraint user_friends_pk
            primary key,
    friends         integer[],
    friends_checked boolean default false not null,
    groups_checked  boolean default false
);

alter table user_friends
    owner to postgres;

create unique index user_friends_vk_id_uindex
    on user_friends (vk_id);

create index user_friends_friends_index
    on user_friends using gin (friends);

