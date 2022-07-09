CREATE table users
(
    id serial not null unique,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE table todo_list
(
    id serial not null unique,
    title varchar(255) not null,
    description varchar(255)
);

CREATE table users_lists
(
    id      serial                                          not null unique,
    user_id int references users (id) on delete cascade     not null,
    list_is int references todo_list (id) on delete cascade not null
);


CREATE table todo_items
(
    id serial not   null unique,
    title varchar(255) not null,
    description varchar(255),
    done boolean not null default false
);

CREATE table lists_items
(
    id      serial                                          not null unique,
    item_id int references todo_items (id) on delete cascade     not null,
    list_is int references todo_list (id) on delete cascade not null
);
