Create table todos if not exists (
    Id Serial primary key,
    Title varchar(50) not null,
    Completed boolean default false,
    Created at timestamp default current_timestamp,
    Updated at timestamp default current_timestamp
);