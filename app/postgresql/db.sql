create table users
(
    id           serial primary key,
    account      varchar(50) unique not null,
    birthdate    date default null,
    gender       smallint default 1 check (gender >= 1 and gender <= 3),
    password     varchar(50) not null,
    nickname     varchar(50),
    country      varchar(3),
    address      varchar(100),
    region_code  varchar(10),
    phone_number varchar(20),
    create_at    timestamp(0) with time zone default now() not null,
    update_at    timestamp(0) with time zone default now() not null
);

create table user_hobbies
(
    id        serial primary key,
    user_id   int not null,
    hobby     varchar(50),
    create_at timestamp(0) with time zone default now() not null,
    update_at timestamp(0) with time zone default now() not null
);

create table user_jobs
(
    id        serial primary key,
    user_id   int not null,
    job       varchar(50),
    create_at timestamp(0) with time zone default now() not null,
    update_at timestamp(0) with time zone default now() not null
);

create or replace
function update_update_at()
returns trigger as $$
begin
  new.update_at = now();
  return new;
end;
$$ language plpgsql;

drop trigger if exists update_users_update_at on users;
create trigger update_users_update_at
    before update
    on users
    for each row execute procedure update_update_at();

drop trigger if exists update_user_hobbies_update_at on user_hobbies;

create trigger update_user_hobbies_update_at
    before update
    on user_hobbies
    for each row execute procedure update_update_at();

drop trigger if exists update_user_jobs_update_at on user_jobs;

create trigger update_user_jobs_update_at
    before update
    on user_jobs
    for each row execute procedure update_update_at();