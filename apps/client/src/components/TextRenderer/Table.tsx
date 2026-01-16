import { decodeHtmlEntities } from '@/helper/renderer.helper';

export default function Table({
  data,
}: {
  data: {
    content: string[][];
    withHeadings: boolean;
    stretched: boolean;
  };
}) {
  return (
    <table className="table-auto">
      <tbody>
        {data.content.map((row, rowIndex: number) => (
          <tr
            key={rowIndex}
            className={data.withHeadings && rowIndex === 0 ? 'font-bold' : ''}
          >
            {row.map((cell, cellIndex: number) => {
              const content = decodeHtmlEntities(cell);
              return (
                <td
                  key={cellIndex}
                  dangerouslySetInnerHTML={{ __html: content }}
                ></td>
              );
            })}
          </tr>
        ))}
      </tbody>
    </table>
  );
}
