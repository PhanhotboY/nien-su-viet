-- CreateSchema
CREATE SCHEMA IF NOT EXISTS "public";

-- CreateEnum
CREATE TYPE "EVENT_DATE_TYPE" AS ENUM ('EXACT', 'APPROXIMATE');

-- CreateTable
CREATE TABLE "User" (
    "id" TEXT NOT NULL,
    "email" TEXT NOT NULL,
    "image" TEXT,
    "name" TEXT NOT NULL,
    "role" TEXT,

    CONSTRAINT "User_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "HistoricalEvent" (
    "id" UUID NOT NULL,
    "fromYear" SMALLINT NOT NULL,
    "fromMonth" SMALLINT,
    "fromDay" SMALLINT,
    "toYear" SMALLINT,
    "toMonth" SMALLINT,
    "toDay" SMALLINT,
    "name" VARCHAR(255) NOT NULL,
    "content" TEXT NOT NULL,
    "thumbnailId" UUID,
    "authorId" TEXT NOT NULL,
    "createdAt" TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(6) NOT NULL,
    "fromDateType" "EVENT_DATE_TYPE" NOT NULL DEFAULT 'EXACT',
    "toDateType" "EVENT_DATE_TYPE" NOT NULL DEFAULT 'EXACT',

    CONSTRAINT "HistoricalEvent_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "EventCategory" (
    "id" UUID NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "slug" VARCHAR(255) NOT NULL,
    "description" TEXT,
    "createdAt" TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(6) NOT NULL,

    CONSTRAINT "EventCategory_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "EventCategories" (
    "eventId" UUID NOT NULL,
    "categoryId" UUID NOT NULL,

    CONSTRAINT "EventCategories_pkey" PRIMARY KEY ("eventId","categoryId")
);

-- CreateTable
CREATE TABLE "EventEdit" (
    "id" UUID NOT NULL,
    "eventId" UUID NOT NULL,
    "editorId" TEXT NOT NULL,
    "editedAt" TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "prevContent" TEXT,
    "newContent" TEXT,
    "prevFromYear" SMALLINT NOT NULL,
    "prevFromMonth" SMALLINT,
    "prevFromDay" SMALLINT,
    "prevToYear" SMALLINT,
    "prevToMonth" SMALLINT,
    "prevToDay" SMALLINT,
    "newFromYear" SMALLINT,
    "newFromMonth" SMALLINT,
    "newFromDay" SMALLINT,
    "newToYear" SMALLINT,
    "newToMonth" SMALLINT,
    "newToDay" SMALLINT,

    CONSTRAINT "EventEdit_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "EventPeriod" (
    "id" UUID NOT NULL,
    "fromYear" SMALLINT NOT NULL,
    "fromMonth" SMALLINT,
    "fromDay" SMALLINT,
    "toYear" SMALLINT NOT NULL,
    "toMonth" SMALLINT,
    "toDay" SMALLINT,
    "name" VARCHAR(255) NOT NULL,
    "slug" VARCHAR(255) NOT NULL,
    "description" TEXT,

    CONSTRAINT "EventPeriod_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "Image" (
    "id" UUID NOT NULL,
    "caption" TEXT,
    "publicUrl" VARCHAR(255) NOT NULL,
    "description" TEXT,

    CONSTRAINT "Image_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE UNIQUE INDEX "User_email_key" ON "User"("email");

-- CreateIndex
CREATE INDEX "idx_event_from_year" ON "HistoricalEvent"("fromYear");

-- CreateIndex
CREATE INDEX "idx_event_from_month" ON "HistoricalEvent"("fromMonth");

-- CreateIndex
CREATE INDEX "idx_event_from_day" ON "HistoricalEvent"("fromDay");

-- CreateIndex
CREATE INDEX "idx_event_to_year" ON "HistoricalEvent"("toYear");

-- CreateIndex
CREATE INDEX "idx_event_to_month" ON "HistoricalEvent"("toMonth");

-- CreateIndex
CREATE INDEX "idx_event_to_day" ON "HistoricalEvent"("toDay");

-- CreateIndex
CREATE UNIQUE INDEX "EventCategory_name_key" ON "EventCategory"("name");

-- CreateIndex
CREATE UNIQUE INDEX "EventCategory_slug_key" ON "EventCategory"("slug");

-- CreateIndex
CREATE INDEX "idx_event_category_slug" ON "EventCategory"("slug");

-- CreateIndex
CREATE INDEX "idx_categories_by_category" ON "EventCategories"("categoryId");

-- CreateIndex
CREATE INDEX "idx_categories_by_event" ON "EventCategories"("eventId");

-- CreateIndex
CREATE INDEX "idx_edit_event" ON "EventEdit"("eventId");

-- CreateIndex
CREATE INDEX "idx_edit_editor" ON "EventEdit"("editorId");

-- CreateIndex
CREATE UNIQUE INDEX "EventPeriod_name_key" ON "EventPeriod"("name");

-- CreateIndex
CREATE UNIQUE INDEX "EventPeriod_slug_key" ON "EventPeriod"("slug");

-- AddForeignKey
ALTER TABLE "HistoricalEvent" ADD CONSTRAINT "HistoricalEvent_authorId_fkey" FOREIGN KEY ("authorId") REFERENCES "User"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "HistoricalEvent" ADD CONSTRAINT "HistoricalEvent_thumbnailId_fkey" FOREIGN KEY ("thumbnailId") REFERENCES "Image"("id") ON DELETE SET NULL ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "EventCategories" ADD CONSTRAINT "EventCategories_categoryId_fkey" FOREIGN KEY ("categoryId") REFERENCES "EventCategory"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "EventCategories" ADD CONSTRAINT "EventCategories_eventId_fkey" FOREIGN KEY ("eventId") REFERENCES "HistoricalEvent"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "EventEdit" ADD CONSTRAINT "EventEdit_editorId_fkey" FOREIGN KEY ("editorId") REFERENCES "User"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "EventEdit" ADD CONSTRAINT "EventEdit_eventId_fkey" FOREIGN KEY ("eventId") REFERENCES "HistoricalEvent"("id") ON DELETE CASCADE ON UPDATE CASCADE;

