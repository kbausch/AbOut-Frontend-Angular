USE `assessment`;

/**
 * Fetch User
 * Retrieves a user's permissions from their CAS username.
 */

DROP PROCEDURE IF EXISTS permissions__fetch_user__sp;

DELIMITER //

CREATE PROCEDURE permissions__fetch_user__sp
    (IN usernameCAS CHAR(25),
    OUT is_super_user TINYINT)
BEGIN
    DECLARE user_id INT;
    
    SELECT
        users.id, users.superuser
    INTO
        user_id, is_super_user
    FROM
        users
    WHERE 
        users.cas = usernameCAS
    LIMIT 1;

    SELECT
        programs.name AS `program_name`,
        permissions.is_manager,
        permissions.is_observer
    FROM
        permissions
        JOIN programs
            ON programs.id = permissions.program_id
    WHERE
        permissions.user_id = user_id;
END //