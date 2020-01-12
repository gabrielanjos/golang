-- Table: public.dadoscliente

-- DROP TABLE public.dadoscliente;

CREATE TABLE public.dadoscliente
(
    id bigint NOT NULL DEFAULT nextval('dadoscliente_id_seq'::regclass),
    cpf text COLLATE pg_catalog."default",
    cpfvalido boolean,
    private integer,
    incompleto integer,
    dataultimacompra text COLLATE pg_catalog."default",
    ticketmedio numeric,
    ticketultimacompra numeric,
    lojamaisfrequente text COLLATE pg_catalog."default",
    lojamaisfrequentevalido boolean,
    lojadaultimacompra text COLLATE pg_catalog."default",
    lojadaultimacompravalido boolean,
    CONSTRAINT dadoscliente_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE public.dadoscliente
    OWNER to postgres;

-- FUNCTION: public.higienizacaodedados(text) # HIGIENIZA O TEXTO RECEBIDO

-- DROP FUNCTION public.higienizacaodedados(text);

CREATE OR REPLACE FUNCTION public.higienizacaodedados(
	text)
    RETURNS text
    LANGUAGE 'plpgsql'

    COST 100
    STABLE 
    
AS $BODY$

DECLARE
input_string ALIAS for $1;
BEGIN

input_string := lower(input_string);
input_string := translate(input_string, 'âãäåāăą', 'aaaaaaa');
input_string := translate(input_string, 'èééêëēĕėęě', 'eeeeeeeeee');
input_string := translate(input_string, 'ìíîïìĩīĭ', 'iiiiiiii');
input_string := translate(input_string, 'óôõöōŏő', 'ooooooo');
input_string := translate(input_string, 'ùúûüũūŭů', 'uuuuuuuu');
input_string := translate(input_string, '.-/', '');

RETURN input_string;
END;

$BODY$;

ALTER FUNCTION public.higienizacaodedados(text)
    OWNER TO postgres;


-- FUNCTION: public.higienizaultimoinsert() # ATUALIZA O ULTIDO REGISTRO INSERIDO COM A HIGIENIZAÇÃO DOS DADOS

-- DROP FUNCTION public.higienizaultimoinsert(); 

CREATE FUNCTION public.higienizaultimoinsert()
    RETURNS trigger
    LANGUAGE 'plpgsql'
    COST 100
    VOLATILE NOT LEAKPROOF
AS $BODY$
BEGIN
update dadoscliente 
set cpf = higienizacaodedados(cpf), 
lojamaisfrequente = higienizacaodedados(lojamaisfrequente), 
lojadaultimacompra = higienizacaodedados(lojadaultimacompra), 
dataultimacompra = higienizacaodedados(dataultimacompra) 
where id = (select max(id) from dadoscliente);
RETURN NEW;
END;
$BODY$;

ALTER FUNCTION public.higienizaultimoinsert()
    OWNER TO postgres;

-- Trigger: higienizaultimoinsert # TRIGGER QUE CHAMA A FUNÇÃO PARA HIGIENIZAR O ULTIMO REGISTRO INSERIDO NA TABELA

-- DROP TRIGGER higienizaultimoinsert ON public.dadoscliente; 

CREATE TRIGGER higienizaultimoinsert 
    BEFORE INSERT
    ON public.dadoscliente
    FOR EACH ROW
    EXECUTE PROCEDURE public.higienizaultimoinsert();