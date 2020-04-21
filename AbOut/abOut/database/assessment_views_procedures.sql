USE `assessment`;

/**
* View Current Programs
* This view will ouput all current programs in the system
*/

CREATE OR REPLACE VIEW programs__list_current__vw 
AS
    SELECT 
        programs.abbrev,
        programs.name,
        semesters.name AS 'current semester'
    FROM 
        programs 
        JOIN intervals 
            ON programs.interval_id = intervals.id
        JOIN semesters
            ON programs.current_semester_id = semesters.id
    WHERE
        intervals.end IS NULL;

/**
 * View Program
 * This stored procedure will list the details about a specified program,
 * which is determined by the input program abbrev.
 */

DROP PROCEDURE IF EXISTS programs__list_one__sp;

DELIMITER //
 
CREATE PROCEDURE programs__list_one__sp
    (IN abbrev CHAR(5),
    OUT status INT,
    OUT error_message CHAR(50))

sp: BEGIN

    SET status = 0;
    SET error_message = NULL;
	
    SELECT 
        @pcount := COUNT(*)
    FROM 
        programs
    WHERE 
        programs.abbrev = abbrev;
    
    IF @pcount = 0 THEN
        SET status = 1;
        SET error_message = "unknown program abbreviation";
        LEAVE sp;
    END IF;
    
    SELECT 
        programs.abbrev, 
        programs.name, 
        semesters.name AS "current_semester"
    FROM 
        programs 
        JOIN semesters
            ON programs.current_semester_id = semesters.id
    WHERE
        programs.abbrev = abbrev
    LIMIT 1;
        
END //
 
DELIMITER ;

        
/**
 * View Outcomes
 * This view will output all current outcomes in the system 
 * (meaning all outcomes which have their end semester as null).
 */

CREATE OR REPLACE VIEW outcomes__list_all__vw
AS
	SELECT
		prefixes.text AS 'prefix',
		outcomes.identifier,
		outcomes.text,
        semesters.name AS 'begin',
        "" AS 'end'
	FROM outcomes
		JOIN prefixes
			ON outcomes.prefix_id = prefixes.id
		JOIN intervals 
			ON outcomes.interval_id = intervals.id
		JOIN semesters
			ON semesters.id = intervals.begin
	WHERE
		intervals.end IS NULL;

/**
 * View Outcome
 * This stored procedure will list the details about a specified outcome,
 * which is determined by the input outcome prefix and identifier.
 */

DROP PROCEDURE IF EXISTS outcomes__list_one__sp;

DELIMITER //
 
CREATE PROCEDURE outcomes__list_one__sp
	(IN pre CHAR(10), 
	IN idnt CHAR(10))
BEGIN
	SELECT
		prefixes.text AS 'prefix',
		outcomes.identifier,
		outcomes.text,
		begin_semester.name AS 'begin',
		IFNULL(end_semester.name, "") AS 'end'
	FROM
		outcomes
		JOIN prefixes 
			ON outcomes.prefix_id = prefixes.id
		JOIN intervals
			ON outcomes.interval_id = intervals.id
		JOIN semesters AS begin_semester
			ON intervals.begin = begin_semester.id
		LEFT JOIN semesters AS end_semester
			ON intervals.end = end_semester.id
	WHERE
		prefixes.text = pre
		AND outcomes.identifier = idnt
	LIMIT 1;

END //
 
DELIMITER ;

/**
 * View Outcomes in Program
 * This stored procedure will list all the outcomes associated with the given program
 */

DROP PROCEDURE IF EXISTS programs_outcomes__outcomes_in_program__sp;

DELIMITER //
 
CREATE PROCEDURE programs_outcomes__outcomes_in_program__sp
    (IN program_abbrev CHAR (5),
    OUT status INT,
    OUT error_message CHAR(50))
sp: BEGIN
    SET status = 0;
    SET error_message = NULL;
    
    -- Ensure the given program exists
    SELECT EXISTS 
    (
        SELECT
            1
        FROM
            programs
        WHERE
            abbrev = program_abbrev
    )
    INTO
        @program_exists;

    -- Leave if the program input is invalid
    IF @program_exists = 0 THEN
        SET status = 1;
        SET error_message = "Given program is invalid";
        LEAVE sp;
    END IF;
    
    -- Grab the associated outcomes
    SELECT
        prefixes.text AS prefix,
        outcomes.identifier,
        outcomes.text
    FROM 
        programs
        JOIN program_outcomes
            ON programs.id = program_outcomes.program_id
        JOIN outcomes
            ON program_outcomes.outcome_id = outcomes.id
        JOIN prefixes
            ON outcomes.prefix_id = prefixes.id
    WHERE
        programs.abbrev = program_abbrev;

END //
 
DELIMITER ;

/**
 * Associate Outcome
 * This stored procedure associates a known outcome with a known program.
 */

DROP PROCEDURE IF EXISTS program_outcomes__associate_outcome__sp;

DELIMITER //
CREATE PROCEDURE program_outcomes__associate_outcome__sp 
    (IN program_abbrev CHAR (5),
    IN prefix_text CHAR (5),
    IN outcome_identifier CHAR (5),
    OUT status INT,
    OUT error_message CHAR(50))
sp: BEGIN

    SET status = 0;
    SET error_message = NULL;

    -- Locate and store the outcome_id and count up if it is found
    SELECT 
        @outcome_id := outcomes.id,
        @ocount := COUNT(*)
    FROM 
        outcomes
        JOIN prefixes 
            ON outcomes.prefix_id = prefixes.id
    WHERE 
        outcome_identifier = outcomes.identifier 
        AND prefix_text = prefixes.text;

    -- Check if the outcome input was valid, exit if not
    IF @ocount = 0 THEN
        SET status = 1;
        SET error_message = "outcome does not exist";
        LEAVE sp;
    END IF;

    -- Locate and store the program_id and count up if it was found
    SELECT 
        @program_id := id,
        @pcount := COUNT(*),
        @current_semester := current_semester_id
    FROM 
        programs
    WHERE 
        abbrev = program_abbrev;

    -- Verify if the program input was valid, exit if not
    IF @pcount = 0 THEN
        SET status = 1;
        SET error_message = "program does not exist";
        LEAVE sp;
    END IF;

    -- Check if the given input has already been associated
    SELECT
        @exists_already := COUNT(*)
    FROM 
        program_outcomes
    WHERE 
        @program_id = program_id 
        AND @outcome_id = outcome_id;

    -- Exit if an association already exists
    IF @exists_already != 0 THEN
        SET status = 1;
        SET error_message = "association already exists";
        LEAVE sp;
    END IF;

    -- Find and store the current interval
    SELECT 
        @interval_id := id 
    FROM 
        intervals
    WHERE 
        @current_semester = begin
        AND end IS NULL;

    -- create the current interval if it does not exist
    IF @interval_id IS NULL THEN
        INSERT INTO 
            intervals (begin, end)
        VALUES 
            (@current_semester, NULL);
        
        SELECT @current_semester;

        -- must select the interval id if it did not exists at first
        SELECT 
            @interval_id := id 
        FROM 
            intervals
        WHERE 
            @current_semester = begin
            AND end IS NULL;
    END IF;

    -- insert the valid input into the table
    INSERT INTO 
        program_outcomes (program_id, outcome_id, interval_id)
    VALUES 
        (@program_id, @outcome_id, @interval_id);
END //

DELIMITER ;

/**
 * Disassociate Outcome
 * This stored procedure will remove an association between a program and an outcome
 */

DROP PROCEDURE IF EXISTS program_outcomes__disassociate_outcome__sp;

DELIMITER //
CREATE PROCEDURE program_outcomes__disassociate_outcome__sp 
    (IN program_abbrev CHAR (5),
    IN prefix_text CHAR (5),
    IN outcome_identifier CHAR (5),
    OUT status INT,
    OUT error_message CHAR(50))
sp: BEGIN

    SET status = 0;
    SET error_message = NULL;
    
    -- Check to see if the program exists
    SELECT EXISTS 
    (
		SELECT
			1
		FROM
			programs
		WHERE
			abbrev = program_abbrev
    )
    INTO
		@program_exists;
        
	-- If the program does not exist, leave
    IF @program_exists = 0 THEN
		SET status = 1;
        SET error_message = "Program does not exist.";
        LEAVE sp;
    END IF;
    
    -- Check to see if the outcome exists
    SELECT EXISTS
    (
		SELECT
			1
		FROM
			outcomes
            JOIN prefixes
				ON outcomes.prefix_id = prefixes.id
		WHERE
			prefixes.text = prefix_text
            AND outcomes.identifier = outcome_identifier
    )
    INTO
		@outcome_exists;
	
    -- If the outcome does not exist, leave
    IF @outcome_exists = 0 THEN
		SET status = 1;
        SET error_message = "Outcome does not exist.";
        LEAVE sp;
    END IF;
    
    -- Get the id of the program
    SELECT
		id
	INTO
		@program_id
	FROM
		programs
	WHERE
		abbrev = program_abbrev
	LIMIT 1;
    
    -- Get the id of the outcome
    SELECT
		outcomes.id
	INTO
		@outcome_id
	FROM
		outcomes
        JOIN prefixes
			ON outcomes.prefix_id = prefixes.id
	WHERE
		prefixes.text = prefix_text
        AND outcomes.identifier = outcome_identifier
	LIMIT 1;
    
    -- Check to see if the association exists
    SELECT EXISTS
    (
		SELECT
			1
		FROM
			program_outcomes
		WHERE
			program_id = @program_id
            AND outcome_id = @outcome_id
    )
    INTO
		@association_exists;
        
	-- If the association does not exist, leave
    IF @association_exists = 0 THEN
		SET status = 1;
        SET error_message = "Association does not exist.";
        LEAVE sp;
    END IF;
    
    -- Delete the association
    DELETE FROM program_outcomes
    WHERE 
		program_id = @program_id
        AND outcome_id = @outcome_id;
    
END //

DELIMITER ;

/**
 * Create Outcome
 * This stored procedure will add a new outcome to the system.
 */

DROP PROCEDURE IF EXISTS outcomes__create_outcome__sp;

DELIMITER //
 
CREATE PROCEDURE outcomes__create_outcome__sp
	(IN pre CHAR(5), 
	IN idnt CHAR(5),
    IN txt VARCHAR(300),
    OUT status INT,
    OUT error_message VARCHAR(100))
sp: BEGIN

	SET status = 0;
    SET error_message = NULL;
    
    -- Check if the prefix exists
    SELECT EXISTS
    (
		SELECT
			id
		FROM
			prefixes
		WHERE
			text = pre
    ) INTO @prefix_exists;
        
	-- If the prefix doesn't exist, leave
    IF @prefix_exists = 0 THEN
		SET status = 1;
        SET error_message = "Prefix does not exist";
        LEAVE sp;
	END IF;
    
    -- Get the prefix id
    SELECT
		id
	INTO
		@prefix_id
	FROM
		prefixes
	WHERE
		text = pre
	LIMIT 1;
    
    -- Determine if the outcome already exists
    SELECT EXISTS
	(
		SELECT
			outcomes.id
		FROM
			outcomes
            JOIN prefixes
				ON outcomes.prefix_id = prefixes.id
		WHERE
			outcomes.identifier = idnt
            AND prefixes.text = pre
	) INTO @outcome_exists;
    
    -- If the outcome input was a duplicate, then exit
    IF @outcome_exists = 1 THEN
		SET status = 1;
        SET error_message = "Outcome already exists with this prefix and identifier";
        LEAVE sp;
    END IF;

	-- Get the current semester id
	SELECT
		MAX(id)
	INTO
		@current_semester
	FROM
		semesters;
        
	-- Query for an interval that begins with the current semester and ends with NULL
    SELECT EXISTS
    (
		SELECT
			id
		FROM
			intervals
		WHERE
			begin = @current_semester
            AND end IS NULL
    ) INTO @interval_exists;

	-- If such an interval doesn't exist yet, create it.
	IF @interval_exists = 0 THEN
        INSERT INTO 
            intervals (begin, end)
        VALUES 
            (@current_semester, NULL);
    END IF;
    
    -- Get the interval id
    SELECT
		id
	INTO
		@interval_id
	FROM
		intervals
	WHERE
		begin = @current_semester
        AND end IS NULL;
        
	INSERT INTO outcomes
		(prefix_id, identifier, text, interval_id)
	VALUES
		(@prefix_id, idnt, txt, @interval_id);
END //
 
DELIMITER ;

/**
 * Update Outcome
 * This stored procedure updates the text of the specified outcome
 * which is located using the prefix text and the identifier of the outcome
 */
 
DROP PROCEDURE IF EXISTS outcomes__update_outcome__sp;

DELIMITER //
 
CREATE PROCEDURE outcomes__update_outcome__sp
    (IN o_prefix CHAR(5),
    IN o_identifier CHAR(5),
    IN new_text VARCHAR(300),
    OUT status INT,
    OUT error_message CHAR(50))
sp: BEGIN
	
    SET status = 0;
    SET error_message = NULL;
	
    -- assigns 1 to @ocount if the  given the outcome prefix and identifier correctly specifiy an existing outcome
    -- and stores the id of the given prefix for later use
    SELECT 
        @prefix_id := outcomes.prefix_id,
        @ocount := COUNT(*)
    FROM 
	    outcomes
        JOIN prefixes
            ON outcomes.prefix_id = prefixes.id
    WHERE
        o_identifier = outcomes.identifier
        AND o_prefix = prefixes.text;
        
    -- does the specified outcome exist? leave if no
    IF @ocount = 0 THEN
        SET status = 1;
        SET error_message = "unknown outcome prefix and/or outcome identifier";
        LEAVE sp;
	END IF;
    
    IF new_text IS NULL THEN
        SET status = 1;
        SET error_message = "encountered null new text";
        LEAVE sp;
    END IF;

    -- update the outcome text
    UPDATE
        outcomes
    SET
        text = new_text
    WHERE
        o_identifier = identifier
        AND @prefix_id = prefix_id;
        
END //
 
DELIMITER ;

/**
 * Delete Outcome
 * 
 */
 
DROP PROCEDURE IF EXISTS outcomes__delete_outcome__sp;

DELIMITER //
 
CREATE PROCEDURE outcomes__delete_outcome__sp
    (IN pre CHAR(5),
    IN idnt CHAR(5),
    OUT status INT,
    OUT error_message CHAR(50))
sp: BEGIN
	
    SET status = 0;
    SET error_message = NULL;
    
    -- Check if the outcome exists or not
    SELECT EXISTS
	(
		SELECT
			1
		FROM
			outcomes
            JOIN prefixes
				ON outcomes.prefix_id = prefixes.id
		WHERE
			outcomes.identifier = idnt
            AND prefixes.text = pre
	) INTO @outcome_exists;
    
    -- If the outcome doesn't exist, leave
    IF @outcome_exists = 0 THEN
		SET status = 1;
        SET error_message = "Outcome does not exist.";
        LEAVE sp;
    END IF;
    
    -- Check if the outcome has any associations
    SELECT
		program_outcomes.id
	INTO
		@association_id
	FROM
		program_outcomes
        JOIN outcomes
			ON program_outcomes.outcome_id = outcomes.id
		JOIN prefixes
			ON outcomes.prefix_id = prefixes.id
	WHERE
		outcomes.identifier = idnt
        AND prefixes.text = pre
	LIMIT 1;
    
    -- If there are any associations, leave
    IF @association_id IS NOT NULL THEN
		SET status = 1;
        SET error_message = "Can't delete outcome that has an association.";
        LEAVE sp;
    END IF;
    
    DELETE FROM outcomes
    WHERE 
		identifier = idnt
        AND prefix_id = (SELECT 
							prefixes.id 
						FROM 
							prefixes 
                        WHERE 
							prefixes.text = pre);
        
END //
 
DELIMITER ;