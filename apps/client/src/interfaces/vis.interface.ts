export interface VisItem {
  id: string;
  content: string;
  start: Date;
  end?: Date;
  title: string;
  type: 'range' | 'point' | 'box' | 'background';
  className?: string;
  group?: string;
  align?: string;
  selectable?: boolean;
  subgroup?: string;
  limitSize?: boolean;
}
