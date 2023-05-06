CREATE OR REPLACE FUNCTION url_get(_uid character varying)
	RETURNS text
	LANGUAGE plpgsql
AS $function$
DECLARE
	_u urls;
BEGIN
	SELECT *
	FROM urls
	WHERE uid = _uid
	INTO _u;

	IF _u.id ISNULL THEN
		RETURN '';
	ELSE
		RETURN _u.link;
	END IF;
END;
$function$