CREATE TABLE "investments" (
    "id" bigserial,
    "amount" int8,
    "success" bool,
    PRIMARY KEY ("id")
);

CREATE TABLE "credit" (
    "id" bigserial,
    "investment_id" int8,
    "type" varchar(32),
    "quantity" int4,
    PRIMARY KEY ("id"),
    CONSTRAINT "fk_credit_investments_1" FOREIGN KEY ("investment_id") REFERENCES "investments" ("id")
);