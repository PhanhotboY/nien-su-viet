import { decodeHtmlEntities } from '@/helper/renderer.helper';

export default function Paragraph({
  data,
  tunes,
}: {
  data: { text: string };
  tunes?: { textAlign: { alignment: string } };
}) {
  const decodedText = decodeHtmlEntities(data.text);

  return (
    <p
      className={`edjs-paragraph text-${tunes?.textAlign.alignment}`}
      dangerouslySetInnerHTML={{ __html: decodedText }}
    ></p>
  );
}
