-- +goose Up
-- +goose StatementBegin
create table if not exists "quotes"
(
    "id"          serial primary key,
    "author_id"   varchar(20) not null,
    "text"        text        not null,
    "guild_id"    varchar(20) not null,
    "added_by_id" varchar(20) not null,
    "timestamp"   timestamptz not null
);

create index if not exists "quotes_author_id"
    on quotes using hash ("author_id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop index if exists "quotes_author_id";
drop table if exists "quotes";
-- +goose StatementEnd
