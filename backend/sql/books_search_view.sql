CREATE OR REPLACE VIEW BOOKS_SEARCH AS
SELECT
    BOOKS.*
  , CONCAT(ARRAY_TO_STRING(ARRAY(SELECT S.NAME
                                 FROM
                                     SUBJECTS S
                                         LEFT JOIN SUBJECT2_BOOKS S2B ON S2B.SUBJECT_ID = S.ID
                                 WHERE S2B.BOOK_ID = BOOKS.ID), ' ')
        , ' '
        , ARRAY_TO_STRING(ARRAY(SELECT A.NAME
                                FROM
                                    AUTHORS A
                                        LEFT JOIN AUTHOR2_BOOKS A2B ON A2B.AUTHOR_ID = A.ID
                                WHERE A2B.BOOK_ID = BOOKS.ID), ' ')
        , ' '
        , COLLECTIONS.TITLE
        , ' '
        , BOOKS.TITLE) AS SEARCH_TERMS
FROM
    BOOKS
        LEFT JOIN COLLECTIONS ON BOOKS.COLLECTION_ID = COLLECTIONS.ID;
