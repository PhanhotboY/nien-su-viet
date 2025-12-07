export default function Quote({
  data,
}: {
  data: { text: string; caption: string; alignment: string };
}) {
  return (
    <blockquote className={`text-${data.alignment}`}>
      “{data.text}”
      {data.caption && (
        <span dangerouslySetInnerHTML={{ __html: ` — ${data.caption}` }}></span>
      )}
    </blockquote>
  );
}
