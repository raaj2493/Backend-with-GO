Create table if not exists todos  (
    Id Serial primary key,
    Title varchar(50) not null,
    Completed boolean default false,
    Created_at timestamp default current_timestamp,
    Updated_at timestamp default current_timestamp
);