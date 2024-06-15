-- tables
create table "quotes"
(
    "id"          serial primary key,
    "author_id"   varchar(20) not null,
    "text"        text        not null,
    "guild_id"    varchar(20) not null,
    "added_by_id" varchar(20) not null,
    "timestamp"   timestamptz not null
);

-- indices
create index "quotes_author_id" on "quotes" using hash ("author_id");
