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

export { decodeHtmlEntities };
