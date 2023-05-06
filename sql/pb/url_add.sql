CREATE OR REPLACE FUNCTION url_add(_ul character varying, _uid character varying)
	RETURNS text
	LANGUAGE plpgsql
AS $function$
DECLARE
	_u urls;
BEGIN
	SELECT *
	FROM urls
	WHERE link = _ul
	INTO _u;

	IF _u.id ISNULL THEN
		INSERT INTO urls(uid, link) VALUES(_uid, _ul);
		RETURN _uid;
	ELSE
		RETURN _u.uid;
	END IF;
END;
$function$