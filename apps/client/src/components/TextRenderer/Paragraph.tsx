function decodeHtmlEntities(text: string): string {
  if (typeof window === 'undefined') {
    // Server-side decoding
    return text
      .replace(/&amp;/g, '&')
      .replace(/&lt;/g, '<')
      .replace(/&gt;/g, '>')
      .replace(/&quot;/g, '"')
      .replace(/&#039;/g, "'")
      .replace(/&nbsp;/g, ' ')
      .replace(/<script[^>]*>([\S\s]*?)<\/script>/gim, ''); // Remove script tags
  }
  // Client-side decoding
  const textarea = document.createElement('textarea');
  textarea.innerHTML = text;
  return textarea.value;
}

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
