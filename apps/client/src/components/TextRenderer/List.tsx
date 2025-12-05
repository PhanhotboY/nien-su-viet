import { Checkbox } from '../ui/checkbox';

export default function List({
  data,
}: {
  data: {
    items: { content: string }[];
    meta: { counterType: '' | 'numeric'; checked?: boolean };
    style: 'ordered' | 'unordered' | 'checklist';
  };
}) {
  const Wrapper = data.style === 'ordered' ? 'ol' : 'ul';
  return (
    <Wrapper className={data.style === 'checklist' ? 'list-none' : ''}>
      {data.items.map((item, index) => (
        <li key={index}>
          {data.style === 'checklist' ? (
            <>
              <input type="checkbox" checked={data.meta.checked} disabled />
              <p>{item.content}</p>
            </>
          ) : (
            item.content
          )}
        </li>
      ))}
    </Wrapper>
  );
}
