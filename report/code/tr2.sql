create or replace trigger check_expedition_dates_trigger
before insert on expeditions_members
for each row
execute function check_expedition_dates();