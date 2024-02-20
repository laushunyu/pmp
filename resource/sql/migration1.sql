create table task
(
    id          text
        constraint task_pk
            primary key,
    worker_id   text,
    created_at  integer default current_timestamp,
    finished_at integer,
    title       text,
    sprint      text,
    parent_id   text
);

create table task_record
(
    task_id   text                              not null,
    worked_at integer default current_timestamp not null,
    worker_id text                              not null,
    detail    text
);

create table worker
(
    id   text not null,
    name text not null
);