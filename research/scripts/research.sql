create table if not exists locations
(
    id           int generated always as identity primary key,
    name         text not null,
    country      text not null,
    nearest_town text not null
);

create table if not exists expeditions
(
    id          int generated always as identity primary key,
    location_id int not null,
    start_date  date not null,
    end_date    date not null,

    foreign key (location_id) references locations(id) on delete cascade
);
