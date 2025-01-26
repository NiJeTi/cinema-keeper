-- +goose Up
-- +goose StatementBegin
create table "movies"
(
    "id"         serial primary key,
    "imdb_id"    varchar(16)  not null unique,
    "title"      varchar(100) not null,
    "year"       varchar(10)  not null,
    "genre"      varchar(50)  not null,
    "director"   varchar(100) not null,
    "plot"       text         not null,
    "poster_url" varchar(256) not null
);

create table "guild_movies"
(
    "id"          serial primary key,
    "movie_id"    integer     not null,
    "guild_id"    varchar(20) not null,
    "added_by_id" varchar(20) not null,
    "added_at"    timestamptz not null,
    "watched_at"  timestamptz,
    "rating"      int2 check ("rating" in (-1, 1)),

    constraint "fk_movies" foreign key ("movie_id")
        references "movies" ("id") on delete cascade
);

create index "movies_imdb_id" on "movies" ("imdb_id");
create index "movies_title" on "movies" ("title");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop index "movies_title";
drop index "movies_imdb_id";

drop table "guild_movies";
drop table "movies";
-- +goose StatementEnd
