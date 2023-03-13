DROP DATABASE IF EXISTS app;
CREATE DATABASE app;
USE app;

SET NAMES utf8mb4;

DROP TABLE IF EXISTS `questions`;
CREATE TABLE `questions`
(
    `id`    varchar(36)  NOT NULL,
    `question` LONGTEXT  NOT NULL,
    `answer`   LONGTEXT,
    `answerer` VARCHAR(36),
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;

INSERT INTO `questions` (`id`, `question`, `answer`, `answerer`) VALUES ("60c2fce5-6bd2-e8df-972f-9e7d1fd453c3", "こんにちは", "",""),("18ac915e-a0d3-8c01-aa6a-52e5ae35c0e6", "こんばんは", "おいしい","7265b13d-9e06-42f6-98e3-41ea742f8fb2")