create database todoDB;

use todoDB;

create Table todo(
	id int NOT NULL auto_increment,
    name varchar(255) NOT NULL,
    complete Boolean DEFAULT false,
    CONSTRAINT PK_Todo primary key (id)
);

insert into todo (name) values ("Sesi-01");
