# MySQL Style Guide (https://github.com/adam-p/markdown-here/wiki/Markdown-Cheatsheet)

## General

- Indents for the main part of a query should be 4 spaces (i.e. when in a CTE)
  - Further indents should be two spaces (except for predicates, which should line up with the WHERE keyword)
- No tabs should be used - only spaces. Your editor should be setup to convert tabs to spaces - see our onboarding template for more details
- Lines of SQL should be no longer than 80 characters
- Commas should be at the end-of-line (EOL) as a right comma, with the exception of temporary filters in the WHERE clause for specific values.
```SQL
-- Good
SELECT
    deleted     AS is_deleted, -- EOL right comma 
    accountId   AS account_id
FROM table
WHERE 
    is_deleted = false 
    AND account_id NOT IN (
        '232'
        , '234' -- left comma
        , '425'
    )

-- Bad
SELECT
    deleted     AS is_deleted, -- EOL right comma 
    accountId   AS account_id
FROM table
WHERE 
    is_deleted = false 
    AND account_id NOT IN ('232', '234', '425')
```
- When ```SELECT```ing, always give each column its own row, with the exception of ```SELECT *``` which can be on a single row
- ```DISTINCT``` should be included on the same row as ```SELECT```
- The ```AS``` keyword should be used when projecting a field or table name
- When aliasing use ```AS```, strive to align the original column names on a single vertical line and the ```AS``` keyword on a separate vertical line
- Fields should be stated before aggregates / window functions
- Ordering and grouping by a number (eg. group by 1, 2) is preferred
- Prefer ```WHERE``` to ```HAVING``` when either would suffice
- Prefer accessing JSON using the bracket syntax, e.g. ```data_by_row['id']::bigint as id_value```
- Never use ```USING``` in joins. On Snowflake, it will produce inaccurate results
- Prefer ```UNION ALL``` to ```UNION```. This is because a ```UNION``` could indicate upstream data integrity issue that are better solved elsewhere.
- Prefer ```!=``` to ```<>```. This is because ```!=``` is more common in other programming languages and reads like "not equal" which is how we're more likely to speak
- Consider performance. Understand the difference between ```LIKE``` vs ```ILIKE```, ```IS``` vs ```=```, and ```NOT``` vs ```!``` vs ```<>```. Use appropriately
- Prefer ```lower(column) LIKE '%match%' to column ILIKE '%Match%'```. This lowers the chance of stray capital letters leading to an unexpected result
- Prefer using ```VARCHAR``` when casting this data type and not specifying a length. A column only consumes storage for the amount of actual data stored.
- Familiarize yourself with ```the DRY Principal```. Leverage CTEs, jinja and macros in dbt, and snippets in Sisense. If you type the same line twice, it needs to be maintained in two places
- DO NOT OPTIMIZE FOR A SMALLER NUMBER OF LINES OF CODE. NEWLINES ARE CHEAP. BRAIN TIME IS EXPENSIVE.

## Comments

- When making single line comments in a model use the ```--``` syntax
- When making multi-line comments in a model use the ```/* */``` syntax
- Respect the character line limit when making comments. Move to a new line or to the model documentation if the comment is too long
- dbt model comments should live in the model documentation
- Calculations made in SQL should have a brief description of what's going on and a link to the handbook defining the metric (and how it's calculated)
- Instead of leaving ```TODO``` comments, create new issues for improvement

## JOINs

- Be explicit when joining, e.g. use ```LEFT JOIN``` instead of ```JOIN```. (Default joins are ```INNER```)
- Prefix the table name to a column when joining, otherwise omit
- Specify the order of a join with the FROM table first and JOIN table second:
```SQL
-- Good
FROM source
    LEFT JOIN other_source 
        ON source.id = other_source.id
WHERE ...

-- Bad
FROM source
    LEFT JOIN other_source 
        ON other_source.id = source.id
WHERE ...
```

## Field Naming and Reference Conventions

- Field names should all be lowercased.
- When joining tables and referencing columns from both, strongly prefer to reference the full table name instead of an alias. When the table name is long (~20), try to rename the CTE if possible, and lastly consider aliasing to something descriptive.
```SQL
-- Good
SELECT 
    budget_forecast_cogs_opex.account_id,
    -- 15 more columns
    date_details.fiscal_year,
    date_details.fiscal_quarter,
    date_details.fiscal_quarter_name,
    cost_category.cost_category_level_1,
    cost_category.cost_category_level_2
FROM budget_forecast_cogs_opex
    LEFT JOIN date_details
        ON date_details.first_day_of_month = budget_forecast_cogs_opex.accounting_period
    LEFT JOIN cost_category
        ON budget_forecast_cogs_opex.unique_account_name = cost_category.unique_account_name

-- Ok, but not preferred. Consider renaming the CTE in lieu of aliasing
SELECT 
    bfcopex.account_id,
    -- 15 more columns
    date_details.fiscal_year,
    date_details.fiscal_quarter,
    date_details.fiscal_quarter_name,
    cost_category.cost_category_level_1,
    cost_category.cost_category_level_2
FROM budget_forecast_cogs_opex bfcopex
    LEFT JOIN date_details
        ON date_details.first_day_of_month = bfcopex.accounting_period
    LEFT JOIN cost_category
        ON bfcopex.unique_account_name = cost_category.unique_account_name

-- Bad
SELECT 
    a.*,
    -- 15 more columns
    b.fiscal_year,
    b.fiscal_quarter,
    b.fiscal_quarter_name,
    c.cost_category_level_1,
    c.cost_category_level_2
FROM budget_forecast_cogs_opex a
    LEFT JOIN date_details b
        ON b.first_day_of_month = a.accounting_period
    LEFT JOIN cost_category c
        ON b.unique_account_name = c.unique_account_name
```

- All field names should be snake-cased
```SQL
-- Good
SELECT
    dvcecreatedtstamp AS device_created_timestamp
FROM table

-- Bad
SELECT
    dvcecreatedtstamp AS DeviceCreatedTimestamp
FROM table
```

### Dates

- Timestamps should end with ```_at```, e.g. ```deal_closed_at```, and should always be in UTC
- Dates should end with ```_date```, e.g. ```deal_closed_date```
- Months should be indicated as such and should always be truncated to a date format, e.g. ```deal_closed_month```
- Always avoid key words like ```date``` or ```month``` as a column name
- Prefer the explicit date function over ```date_part```, but prefer ```date_part``` over ```extract```, e.g. ```DAYOFWEEK(created_at)``` > ```DATE_PART(dayofweek, 'created_at')``` > ```EXTRACT(dow FROM created_at)```
  - Note that selecting a date's part is different from truncating the date. ```date_trunc('month', created_at)``` will produce the calendar month ('2019-01-01' for '2019-01-25') while ```SELECT date_part('month', '2019-01-25'::date)``` will produce the number 1
- Be careful using ```DATEDIFF```, as the results are often non-intuitive.
  - For example, ```SELECT DATEDIFF('days', '2001-12-01 23:59:59.999', '2001-12-02 00:00:00.000')``` returns ```1``` even though the timestamps are different by one millisecond.
  - Similarly, ```SELECT DATEDIFF('days', '2001-12-01 00:00:00.001', '2001-12-01 23:59:59.999')``` return ```0``` even though the timestamps are nearly an entire day apart.
  - Using the appropriate interval with the ```DATEDIFF``` function will ensure you are getting the right results. For example, ```DATEDIFF('days', '2001-12-01 23:59:59.999', '2001-12-02 00:00:00.000')``` will provide a ```1 day interval``` and ```DATEDIFF('ms', '2001-12-01 23:59:59.999', '2001-12-02 00:00:00.000')``` will provide a ```1 millisecond interval```.

## Use CTEs (Common Table Expressions), not subqueries

- CTEs make SQL more readable and are more performant
- Use CTEs to reference other tables. Think of these as import statements
- CTEs should be placed at the top of the query
- Where performance permits, CTEs should perform a single, logical unit of work
- CTE names should be as concise as possible while still being clear
  - Avoid long names like ```replace_sfdc_account_id_with_master_record_id``` and prefer a shorter name with a comment in the CTE. This will help avoid table aliasing in joins
- CTEs with confusing or noteable logic should be commented in file and documented in dbt docs
- CTEs that are duplicated across models should be pulled out into their own models
- Leave an empty row above and below the query statement
- CTEs should be formatted as follows:
```SQL
WITH events AS ( -- think of these select statements as your import statements.
     
  ...
     
), filtered_events AS ( -- CTE comments go here
     
  ...
     
) 
     
SELECT * -- you should always aim to "select * from final" for your last model
FROM filtered_events
```

## Stored procedures and views

- Stored procedures should named as ```table_name__functionality__sp```, e.g. ```program_outcomes__associate_outcome__sp```
  - The table name should preferably be the table that is most affected by the procedure, though this is up to the discretion of the person naming the procedure
- Stored procedures should named as ```table_name__functionality__vw```, e.g. ```programs__list_all__vw```
  - The table name should preferably be the table that is most affected by the procedure, though this is up to the discretion of the person naming the procedure

## Functions

- Function names and keywords should all be capitalized
- Prefer ```IFNULL``` TO ```NVL```
- Prefer ```IFF``` to a single line ```CASE WHEN``` statement
- Prefer ```IFF``` to selecting a boolean statement ```(amount < 10) AS is_less_than_ten```