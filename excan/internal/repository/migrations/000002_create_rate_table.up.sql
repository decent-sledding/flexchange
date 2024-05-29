CREATE TABLE IF NOT EXISTS rate(
    rateid     bool    PRIMARY KEY DEFAULT true
    , usd      float   NOT NULL
    , uah      float   NOT NULL
    , base     CHAR(3) NOT NULL 
    CONSTRAINT rate_uq CHECK (rateid)
);