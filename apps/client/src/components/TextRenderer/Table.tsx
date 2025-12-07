export default function Table({ data }: { data: any }) {
  console.log(data);
  return (
    <table>
      <tbody>
        {data.content.map((row: any, rowIndex: number) => (
          <tr key={rowIndex}>
            {row.map((cell: any, cellIndex: number) => (
              <td
                key={cellIndex}
                style={{ width: `${100 / data.content[0].length}%` }}
              >
                {cell}
              </td>
            ))}
          </tr>
        ))}
      </tbody>
    </table>
  );
}
