
CREATE SCHEMA IF NOT EXISTS `sampleDB` DEFAULT CHARACTER SET utf8 ;
USE `sampleDB` ;

-- -----------------------------------------------------
-- Table `dojo_api`.`user`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `sampleDB`.`sample` (
  `id` VARCHAR(128) NOT NULL COMMENT 'ユーザID',
  `name` VARCHAR(64) NOT NULL COMMENT 'ユーザ名',
  `createdAt` datetime DEFAULT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB
COMMENT = 'ユーザ';
