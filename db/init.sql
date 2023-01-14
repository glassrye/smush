CREATE TABLE filestat (
    id BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
    orig_sum varchar(80) NOT NULL UNIQUE,
    comp_sum varchar(80) NOT NULL UNIQUE,
    oname varchar(255) NOT NULL,
    cname varchar(255) NOT NULL,
    cur_loc varchar(255),
    mod_date TIMESTAMP,
    PRIMARY KEY (id)
)