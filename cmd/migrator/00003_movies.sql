-- +goose Up
-- +goose StatementBegin
create table "movies"
(
    "id"      serial primary key,
    "imdb_id" integer      not null,
    "title"   varchar(100) not null
);

create table "guild_movies"
(
    "id"          serial primary key,
    "movie_id"    integer     not null,
    "added_by_id" varchar(20) not null,
    "added_at"    timestamptz not null,
    "watched_at"  timestamptz not null,
    "rating"      smallint    not null check ("rating" in (-1, 0, 1)) default 0,

    constraint "fk_movies" foreign key ("movie_id")
        references "movies" ("id") on delete cascade
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table "guild_movies";
drop table "movies";
-- +goose StatementEnd
