create table `b_m_maker` 
    -- --------------------------------------------------
    --  Table Structure for `qax580go/models.BMMaker`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `b_m_maker` (
        `id` bigint AUTO_INCREMENT NOT NULL PRIMARY KEY,
        `m_id` varchar(255) NOT NULL DEFAULT '' ,
        `name` varchar(255) NOT NULL DEFAULT '' ,
        `lng` double precision NOT NULL DEFAULT 0 ,
        `lat` double precision NOT NULL DEFAULT 0 ,
        `describe` varchar(4096) NOT NULL DEFAULT '' ,
        `time` bigint NOT NULL DEFAULT 0 
    ) ENGINE=InnoDB;