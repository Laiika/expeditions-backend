create table if not exists expeditions_members
(
    id            int generated always as identity primary key,
    expedition_id int not null,
    member_id     int not null,

    foreign key (expedition_id) references expeditions(id)
    on delete cascade,
    foreign key (member_id) references members(id)
    on delete cascade
);

create table if not exists expeditions_curators
(
    id            int generated always as identity primary key,
    expedition_id int not null,
    curator_id    int not null,

    foreign key (expedition_id) references expeditions(id)
    on delete cascade,
    foreign key (curator_id) references curators(id)
    on delete cascade
);
