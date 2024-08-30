-- ТАБЛИЦЫ

create table if not exists admins
(
    id           int generated always as identity primary key,
    name         text not null,
    login        text unique not null,
    password     text not null
);

create table if not exists leaders
(
    id           int generated always as identity primary key,
    name         text not null,
    phone_number text not null,
    login        text unique not null,
    password     text not null
);

create table if not exists members
(
    id           int generated always as identity primary key,
    name         text not null,
    phone_number text not null,
    role         text not null,
    login        text unique not null,
    password     text not null
);

create table if not exists curators
(
    id           int generated always as identity primary key,
    name         text unique not null
);

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

create table if not exists artifacts
(
    id          int generated always as identity primary key,
    location_id int not null,
    name        text not null,
    age         int not null,

    foreign key (location_id) references locations(id) on delete cascade
);

create table if not exists equipments
(
    id            int generated always as identity primary key,
    expedition_id int not null,
    name          text not null,
    amount        int not null,

    foreign key (expedition_id) references expeditions(id) on delete cascade
);

create table if not exists expeditions_leaders
(
    id            int generated always as identity primary key,
    expedition_id int not null,
    leader_id     int not null,

    foreign key (expedition_id) references expeditions(id) on delete cascade,
    foreign key (leader_id) references leaders(id) on delete cascade
);

create table if not exists expeditions_members
(
    id            int generated always as identity primary key,
    expedition_id int not null,
    member_id     int not null,

    foreign key (expedition_id) references expeditions(id) on delete cascade,
    foreign key (member_id) references members(id) on delete cascade
);

create table if not exists expeditions_curators
(
    id            int generated always as identity primary key,
    expedition_id int not null,
    curator_id    int not null,

    foreign key (expedition_id) references expeditions(id) on delete cascade,
    foreign key (curator_id) references curators(id) on delete cascade
);

-- РОЛИ

-- Участник
create role member;
grant select on public.expeditions to member;
grant select on public.leaders to member;
grant select on public.members to member;
grant select on public.curators to member;
grant select on public.locations to member;
grant select on public.artifacts to member;
grant select on public.equipments to member;
grant select on public.expeditions_leaders to member;
grant select on public.expeditions_members to member;
grant select on public.expeditions_curators to member;

create user member1 with PASSWORD 'member1' in role member;

-- Руководитель
create role leader inherit;
grant member to leader;
grant insert, delete on public.members to leader;
grant insert, update, delete on public.expeditions to leader;
grant insert, delete on public.curators to leader;
grant insert, delete on public.locations to leader;
grant insert on public.artifacts to leader;
grant insert, delete on public.equipments to leader;
grant insert, delete on public.expeditions_members to leader;
grant insert, delete on public.expeditions_curators to leader;

create user leader1 with PASSWORD 'leader1' in role leader;

-- Администратор
create role admin;
grant create, usage on schema public to admin;
grant all privileges on all tables in schema public to admin;

create user admin1 with PASSWORD 'admin1' in role admin;

-- ТРИГГЕР

create or replace function check_expedition_dates()
returns trigger as $$
declare
    start_d date;
    end_d date;
    overlapping_count integer;
begin
    select start_date, end_date
    into start_d, end_d
    from expeditions
    where id = new.expedition_id;

    select count(*)
    into overlapping_count
    from expeditions ex
    join expeditions_members em on ex.id = em.expedition_id
    where em.member_id = new.member_id and not(end_d < ex.start_date or start_d > ex.end_date);

    if overlapping_count > 0 then
        raise exception
            'a expedition with the same member and overlapping date already exists';
    end if;

    return new;
end;
$$ language plpgsql;

create or replace trigger check_expedition_dates_trigger
before insert on expeditions_members
for each row
execute function check_expedition_dates();
