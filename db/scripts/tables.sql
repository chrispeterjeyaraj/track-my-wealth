-- Create Table Income
CREATE TABLE IF NOT EXISTS income (
  "id" TEXT NOT NULL,
  "income" TEXT NOT NULL,
  "category" TEXT NOT NULL,

  PRIMARY KEY ("id")
);

-- Create Table Expense
CREATE TABLE IF NOT EXISTS expense (
  "id" TEXT NOT NULL,
  "expense" TEXT NOT NULL,
  "category" TEXT NOT NULL,

  PRIMARY KEY ("id")
);

-- Create Table Savings
CREATE TABLE IF NOT EXISTS savings (
  "id" TEXT NOT NULL,
  "saving" TEXT NOT NULL,
  "category" TEXT NOT NULL,

  PRIMARY KEY ("id")
);


-- Create Table User Preference
CREATE TABLE IF NOT EXISTS userpref (
  "id" TEXT NOT NULL,
  "setting" TEXT NOT NULL,
  "configuration" TEXT NOT NULL,

  PRIMARY KEY ("id")
);

