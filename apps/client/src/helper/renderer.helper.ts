import { sanitizeHtml } from '@/lib/utils';

function decodeHtmlEntities(text: string): string {
  if (typeof window === 'undefined') {
    // Server-side decoding
    return sanitizeHtml(
      text
        .replace(/&amp;/g, '&')
        .replace(/&lt;/g, '<')
        .replace(/&gt;/g, '>')
        .replace(/&quot;/g, '"')
        .replace(/&#039;/g, "'")
        .replace(/&nbsp;/g, ' ')
        .replace(/<script[^>]*>([\S\s]*?)<\/script>/gim, ''),
    ); // Remove script tags
  }
  // Client-side decoding
  const textarea = document.createElement('textarea');
  textarea.innerHTML = text;
  return textarea.value;
}

function tryParseJSONObject(jsonString: string) {
  try {
    var o = JSON.parse(jsonString);

    // Handle non-exception-throwing cases:
    // Neither JSON.parse(false) or JSON.parse(1234) throw errors, hence the type-checking,
    // but... JSON.parse(null) returns null, and typeof null === "object",
    // so we must check for that, too. Thankfully, null is falsey, so this suffices:
    if (o && typeof o === 'object') {
      return o;
    }
  } catch (e) {}

  return false;
}

export { tryParseJSONObject, decodeHtmlEntities };
