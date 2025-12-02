-- AlterTable
ALTER TABLE "EventEdit" ALTER COLUMN "prevContent" DROP NOT NULL,
ALTER COLUMN "newContent" DROP NOT NULL;

-- AlterTable
ALTER TABLE "Image" ADD COLUMN     "description" TEXT;
