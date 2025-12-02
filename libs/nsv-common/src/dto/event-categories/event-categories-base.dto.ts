import { Exclude, Expose } from 'class-transformer';

@Exclude()
export class EventCategoriesBaseDto {
  @Expose()
  eventId!: string;

  @Expose()
  categoryId!: string;
}
