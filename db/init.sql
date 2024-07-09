-- ТАБЛИЦЫ

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
grant insert, update, delete on public.members to leader;
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

create or replace function check_analyzer()
returns trigger
as $$
begin
    if new.analyzer_id is null then
        return new;
end if;

    if (select count(*) from sensors where analyzer_id = new.analyzer_id) >=
    (select ats.max_sensors from analyzer_types ats join
    (select type from gas_analyzers where id = new.analyzer_id) ga on ga.type = ats.name) then
        raise exception 'gas analyzer with id % is already has max number of sensors', new.analyzer_id;
end if;

    if new.gas not in
    (select gas from types_gases where analyzer_type in
    (select type from gas_analyzers where id = new.analyzer_id)) then
        raise exception 'gas analyzer cannot work with %', new.gas;
end if;

return new;
end;
$$ language plpgsql;

create or replace trigger check_analyzer_trigger
before insert or update on sensors
                               for each row execute function check_analyzer();

-- ИНДЕКСЫ

create index idx_expeditions_members_member_id on expeditions_members(member_id);
create index idx_expeditions_members_expedition_id on expeditions_members(expedition_id);