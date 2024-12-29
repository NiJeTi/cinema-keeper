-- +goose Up
-- +goose StatementBegin
drop index "quotes_author_id";
create index "quotes_author_id" on quotes ("author_id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop index "quotes_author_id";
create index if not exists "quotes_author_id"
    on quotes using hash ("author_id");
-- +goose StatementEnd
