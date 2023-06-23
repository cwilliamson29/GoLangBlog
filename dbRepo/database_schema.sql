USE blog;

CREATE TABLE `users` (
                         `id` INT NOT NULL AUTO_INCREMENT,
                         `name` varchar(45) NOT NULL,
                         `email` varchar(125) NOT NULL,
                         `password` varchar(256) NOT NULL,
                         `user_type` int NOT NULL,
                         `banned` BOOL DEFAULT 0 NOT NULL;
                         `acct_created` TIMESTAMP NOT NULL,
                         `last_login` TIMESTAMP NOT NULL,
                         PRIMARY KEY (`id`)
);

CREATE TABLE `user_types` (
                              `id` int NOT NULL AUTO_INCREMENT,
                              `type` varchar(45) NOT NULL,
                              PRIMARY KEY (`id`)
);

CREATE TABLE `posts` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `title` varchar(45) NOT NULL,
    `content` TEXT NOT NULL,
    `user_id` INT NOT NULL,
    `category` INT NOT NULL,
    `sub_category` INT NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `comments`(
    `id`  INT NOT NULL AUTO_INCREMENT,
    `post_id` INT NOT NULL,
    `user_id` INT NOT NULL,
    `content` TEXT NOT NULL,
    `comment_date` TIMESTAMP NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `category` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `name` varchar(45) NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `sub_category` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `name` varchar(45) NOT NULL,
    `parent_category` INT NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `navbar`(
    `id` INT NOT NULL AUTO_INCREMENT,
    `name` varchar(45) NOT NULL,
    `position` INT NOT NULL,
    `menu_id` int,
    `single_page_id` int,
    PRIMARY KEY (`id`)
);

CREATE TABLE `menu` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `name` varchar(45) NOT NULL,
    `target` varchar(256) NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `menu_item`(
    `id` INT NOT NULL AUTO_INCREMENT,
    `name` varchar(45) NOT NULL,
    `target` varchar(256) NOT NULL,
    `menu_parent_id` INT NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `site_settings` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `logo` varchar(256) NOT NULL,
    `brand_logo` varchar(256) NOT NULL,
    `site_name` varchar(125) NOT NULL,
    `site_desc` TEXT NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `single_page` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `title` varchar(125) NOT NULL,
    `content` TEXT NOT NULL,
    `user_id` INT NOT NULL,
    PRIMARY KEY (`id`)
);

ALTER TABLE `users` ADD CONSTRAINT `users_fk0` FOREIGN KEY (`user_type`) REFERENCES `user_types`(`id`);

ALTER TABLE `navbar` ADD CONSTRAINT `navbar_fk0` FOREIGN KEY (`menu_id`) REFERENCES `menu`(`id`);

ALTER TABLE `navbar` ADD CONSTRAINT `navbar_fk1` FOREIGN KEY (`single_page_id`) REFERENCES `single_page`(`id`);

ALTER TABLE `menu_item` ADD CONSTRAINT `menu_item_fk0` FOREIGN KEY (`menu_parent_id`) REFERENCES `menu`(`id`);

ALTER TABLE `posts` ADD CONSTRAINT `posts_fk0` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`);

ALTER TABLE `posts` ADD CONSTRAINT `posts_fk1` FOREIGN KEY (`category`) REFERENCES `category`(`id`);

ALTER TABLE `posts` ADD CONSTRAINT `posts_fk2` FOREIGN KEY (`sub_category`) REFERENCES `sub_category`(`id`);

ALTER TABLE `comments` ADD CONSTRAINT `comments_fk0` FOREIGN KEY (`posts_id`) REFERENCES `posts`(`id`);

ALTER TABLE `comments` ADD CONSTRAINT `comments_fk1` FOREIGN KEY (`users_id`) REFERENCES `users`(`id`);

ALTER TABLE `sub_category` ADD CONSTRAINT `sub_category_fk0` FOREIGN KEY (`parent_category`) REFERENCES `category`(`id`);

ALTER TABLE `single_page` ADD CONSTRAINT `single_page_fk0` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`);

insert into user_types(type)
values(
       "user",
       "monderator",
       "admin"
);


insert into users(name, email, password, acct_created, last_login, user_type)
-- Password is "TurtleDove"
values("Admin", "admin@admin.com", "$2a$12$hEBSoCizHAMXhX3Z2Sb0j./gJB0ojZOsgll4Eq4BijI94vtwxIDJi", NOW(), NOW(), 3);

insert into category(name)
values("main");

insert into sub_category(name, parent_category)
    value("sub main", 1);

INSERT INTO posts(title, content, user_id, category, sub_category)
VALUES("The effects of heaven on the human brain", "To do or not to do, this blog post has seen better days of an over explained concept of the human thought process.", 1, 1,1);

INSERT INTO posts(title, content, user_id, category, sub_category)
VALUES("The effects of hell on the human brain", "To do or not to do, this blog post has seen better days of an over explained concept of the human thought process.", 1, 1,1);

INSERT INTO posts(title, content, user_id, category, sub_category)
VALUES("The effects of programming on the human brain", "To do or not to do, this blog post has seen better days of an over explained concept of the human thought process.", 1, 1,1);



