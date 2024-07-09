create table if not exists members
(
    id           int generated always as identity primary key,
    name         text not null,
    phone_number text not null,
    login        text unique not null,
    password     text not null
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

create table if not exists expeditions_members
(
    id            int generated always as identity primary key,
    expedition_id int not null,
    member_id     int not null,

    foreign key (expedition_id) references expeditions(id) on delete cascade,
    foreign key (member_id) references members(id) on delete cascade
);