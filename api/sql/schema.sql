CREATE TABLE `characters` (
    `id` BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(255) NOT NULL,
    `image_url` VARCHAR(255) NOT NULL,
) ENGINE=InnoDB CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;

CREATE TABLE character_artifact (
    id int primary key auto_increment,
    character_id int,
    artifact_id int,
    FOREIGN KEY (character_id) REFERENCES character(id),
    FOREIGN KEY (artifact_id) REFERENCES artifact(id)
);

CREATE TABLE artifact (
    id int primary key auto_increment,
    name varchar(255)
);

CREATE TABLE user (
    id int primary key auto_increment,
    name varchar(255)
);

CREATE TABLE status (
    id int primary key auto_increment,
    character_id int,
    status varchar(255),
    FOREIGN KEY (character_id) REFERENCES character(id)
);

CREATE TABLE effect (
    id int primary key auto_increment,
    status_id int,
    effect varchar(255),
    FOREIGN KEY (status_id) REFERENCES status(id)
);
`