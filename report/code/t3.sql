create table if not exists artifacts
(
    id          int generated always as identity primary key,
    location_id int not null,
    name        text not null,
    age         int not null,

    foreign key (location_id) references locations(id)
    on delete cascade
);

create table if not exists equipments
(
    id            int generated always as identity primary key,
    expedition_id int not null,
    name          text not null,
    amount        int not null,

    foreign key (expedition_id) references expeditions(id)
    on delete cascade
);

create table if not exists expeditions_leaders
(
    id            int generated always as identity primary key,
    expedition_id int not null,
    leader_id     int not null,

    foreign key (expedition_id) references expeditions(id)
    on delete cascade,
    foreign key (leader_id) references leaders(id)
    on delete cascade
);