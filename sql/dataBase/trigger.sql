	CREATE or REPLACE public.messages_notify_trigger() RETURNS trigger AS $$
	DECLARE
		BEGIN
			PERFORM pg_notify(CAST ('new_message' AS text), row_to_json(NEW)::text), 
			RETURN new;
		END;
	$$ LANGUAGE plpgsql;

	CREATE TRIGGER messages_new_trigger AFTER INSERT ON public.messages
	FOR EACH ROW EXECUTE PROCEDURE public.messages_notify_trigger();
    
    /* IF used AFTER UPDATE, you can keep it for all, or compare OLD and NEW data */
    /*WHEN (OLD.text IS DISTINCT FROM NEW.text)*/
	
    /* TRIGGER can be dropped same as table*/
	/*DROP TRIGGER messages_new_trigger AFTER INSERT ON public.messages*/