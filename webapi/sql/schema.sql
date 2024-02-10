USE `genshindb`;

-- 原神の所持キャラクター
CREATE TABLE `characters` (
    `id` BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(255) NOT NULL,
    `image_url` VARCHAR(255)
) ENGINE=InnoDB CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;

-- 聖遺物スコア
CREATE TABLE scores (
    `id` BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `user_id` BIGINT NOT NULL,
    `character_id` BIGINT NOT NULL,
    `score_type_id` INT,
    `score` BIGINT
) ENGINE=InnoDB CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;;

-- discordのユーザー情報
CREATE TABLE users (
    `id` BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(255) NOT NULL,
    `discord_id` VARCHAR(255) NOT NULL,
    `genshin_uuid` VARCHAR(255) NOT NULL
) ENGINE=InnoDB CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;;

-- 聖遺物
CREATE TABLE artifacts (
    `id` BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(255) NOT NULL,
    `part` ENUM('flower', 'plume', 'sands', 'goblet', 'circlet') NOT NULL
) ENGINE=InnoDB CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;;

-- 聖遺物の効果
CREATE TABLE effects (
    `id` BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `artifact_id` BIGINT NOT NULL,
    `user_id` VARCHAR(255) NOT NULL,
    `is_percentage` BOOLEAN NOT NULL,
    `value` DECIMAL(5, 1) NOT NULL
) ENGINE=InnoDB CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;;
