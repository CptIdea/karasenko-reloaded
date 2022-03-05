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

insert into user_friends (vk_id, friends, friends_checked, groups_checked) VALUES (362145315,'{31480508, 35294456, 29534144, 40638632, 150550417, 20629724, 209366260, 125004421, 92876084, 210608903, 166774454, 154306815, 157081760, 30666517, 114913332, 109575676, 148438786, 154906069, 140336241, 51016572, 82887947, 110242131, 121574455, 121141827, 132147085, 88350989, 111587102, 179664673, 78515560, 67171472, 82476651, 119359773, 162121263, 109796404, 121732667, 166142266, 75538001, 169431149, 128637780, 39243732, 106352936, 68670236, 58808669, 52171320, 91933860, 32295218, 122560283, 62830729, 140105161, 77715974, 169626752, 34662672, 183190591, 135856652, 16711345, 135897440, 2158488, 183563391, 86217001, 25336774, 180547025, 76031716, 107232016, 164934159, 170172287, 197495754, 30618587, 157645793, 98093162, 162318648, 154430577, 162573051, 44180355, 114503206, 138347372, 129642215, 162148795, 135682415, 207762396, 174299957, 161665691, 183540171, 130486220, 131710892, 26713492, 104763752, 153164937, 191306907, 157572003, 52755582, 26018968, 207532056, 151186434, 109239782, 63731512, 147415323, 133194853, 76229642, 185349205, 206353295, 96007970, 84064469, 153014773, 170240710, 174193655, 106631887, 170257648, 196605827, 206456855, 206046622, 65847147, 206081459, 166359672, 71755823, 88550574, 128707424, 31649688, 68809723, 142582865, 168389251, 164516058, 116143547, 22822305, 80869263, 52652740, 96948046, 31219902, 194752494, 169710947, 202866815, 194978009, 165643679, 56929213, 199954890, 104224126, 202857559, 202576242, 168354344, 170569428, 171597505, 160204864, 184475011, 84528882, 34300760, 40059402, 163500389, 128763224, 179032552, 196520157, 162113393, 174045069, 171027291, 183522239, 94551105, 107386670, 98210264, 68372674, 180388405, 186626379, 138117729, 105494325, 57595712, 58385845, 59356375, 81335553, 85928959, 94387079, 100786550, 109069686, 111825283, 124123414, 125057453, 131571666, 132528669, 137011084, 140336673, 142559517, 149289390, 149529071, 151864030, 152320618, 156357049, 158080996, 158868136, 160151110, 160239461, 161501232, 163247882, 164197441, 166244206, 166441561, 166600751, 170377356, 170670690, 171002293, 172276612, 178640969, 182472331, 183752300, 188069951, 189660722, 192946559, 194354238}',false,false);

CREATE EXTENSION intarray;