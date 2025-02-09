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
    where em.member_id = new.member_id and
          not(end_d < ex.start_date or start_d > ex.end_date);

    if overlapping_count > 0 then
        raise exception
            'a expedition with the same member ' ||
            'and overlapping date already exists';
    end if;

    return new;
end;
$$ language plpgsql;