CREATE TABLE IF NOT EXISTS injection_certifications (
    injection_certification_id serial primary key,
    dog_id bigint not null,
    type int not null,
    file text not null,
    reg_at timestamp not null,
    upd_at timestamp not null
);
