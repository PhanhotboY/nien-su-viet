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
      {data.items.map((item, index) => (
        <li key={index}>
          {data.style === 'checklist' ? (
            <>
              <input type="checkbox" checked={item.meta.checked} disabled />
              <p>{item.content}</p>
            </>
          ) : (
            item.content
          )}

          {item.items && (
            <List data={{ items: item.items, meta: {}, style: data.style }} />
          )}
        </li>
      ))}
    </Wrapper>
  );
}
