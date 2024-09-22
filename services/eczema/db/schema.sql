create table Event (
    time integer not null primary key,
    duration integer not null,

    state real not null,
    variant integer not null,
    location integer not null,
    action integer not null
);
