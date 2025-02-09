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

create role admin;
grant create, usage on schema public to admin;
grant all privileges on all tables in schema public to admin;

create user admin1 with PASSWORD 'admin1' in role admin;