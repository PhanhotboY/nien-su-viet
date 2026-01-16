import { decodeHtmlEntities } from '@/helper/renderer.helper';

type ListItem = {
  content: string;
  meta: { checked?: boolean };
  items?: ListItem[];
};

export default function List({
  data,
}: {
  data: {
    items: ListItem[];
    meta: { counterType?: '' | 'numeric' };
    style: 'ordered' | 'unordered' | 'checklist';
  };
}) {
  const Wrapper = data.style === 'ordered' ? 'ol' : 'ul';

  return (
    <Wrapper className={data.style === 'checklist' ? 'list-none' : ''}>
      {data.items.map((item, index) => {
        const content = decodeHtmlEntities(item.content);
        return (
          <li key={index}>
            {data.style === 'checklist' && (
              <input type="checkbox" checked={item.meta.checked} disabled />
            )}
            <p dangerouslySetInnerHTML={{ __html: content }}></p>

            {item.items && (
              <List data={{ items: item.items, meta: {}, style: data.style }} />
            )}
          </li>
        );
      })}
    </Wrapper>
  );
}
