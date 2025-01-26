-- +goose Up
-- +goose StatementBegin
create table "movies"
(
    "id"      serial primary key,
    "imdb_id" varchar(9)   not null unique,
    "title"   varchar(100) not null
);

create table "guild_movies"
(
    "id"          serial primary key,
    "movie_id"    integer     not null,
    "guild_id"    varchar(20) not null,
    "added_by_id" varchar(20) not null,
    "added_at"    timestamptz not null,
    "watched_at"  timestamptz,
    "rating"      smallint check ("rating" in (-1, 1)),

    constraint "fk_movies" foreign key ("movie_id")
        references "movies" ("id") on delete cascade
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table "guild_movies";
drop table "movies";
-- +goose StatementEnd
