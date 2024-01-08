CREATE OR REPLACE FUNCTION gen_random_int(min int, max int)
RETURNS INT AS
$$
DECLARE
  result INT;
BEGIN
  result = floor(random() * (max - min + 1) + 1);

  RETURN result;
END;
$$ LANGUAGE plpgsql;

--
CREATE OR REPLACE FUNCTION gen_random_string(
    num_characters INT
) RETURNS VARCHAR AS
$$
DECLARE
    characters VARCHAR[] := ARRAY['a','b','c','d','e','f','g','h','i','j','k','l','m','n','o','p','q','r','s','t','u','v','w','x','y','z'];
    random_string VARCHAR := '';
    random_index INT;
BEGIN
    FOR i IN 1..num_characters LOOP
        random_index := floor(random() * array_length(characters, 1) + 1);
        random_string := random_string || characters[random_index];
    END LOOP;

    RETURN random_string;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION gen_random_date()
RETURNS TIMESTAMPTZ AS
$$
DECLARE
    result_timestamp TIMESTAMPTZ;
BEGIN
    -- Construct a timestamp using the provided year, month, and day
  result_timestamp = (timestamp '2024-01-01 20:00:00' - floor(random()*365))

    RETURN result_timestamp;
END;
$$ LANGUAGE plpgsql;