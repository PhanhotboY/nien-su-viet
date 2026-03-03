import { tryParseJSONObject } from '@/helper/renderer.helper';

const DEFAULT_EXCERPT_LENGTH = 150;

type IEditorParagraphData = {
  text: string;
};
type IEditorHeadingData = {
  text: string;
  level: number;
};
type IEditorImageData = {
  file: { url: string };
  caption: string;
};
type IEditorListData = {
  items: { content: string }[];
  meta: { counterType: '' | 'numeric' };
  style: 'ordered' | 'unordered';
};
type IEditorLinkToolData = {
  link: string;
  meta: { title: string; description: string; image: { url: string } };
};

interface IEditorBlock {
  type: string;
  data:
    | IEditorHeadingData
    | IEditorParagraphData
    | IEditorImageData
    | IEditorListData
    | IEditorLinkToolData;
}

const extractTextFromBlock = (block: IEditorBlock): string => {
  let text = '';

  switch (block.type) {
    case 'paragraph':
      text = (block.data as IEditorParagraphData).text || '';
      break;
    case 'header':
      text = (block.data as IEditorHeadingData).text || '';
      break;
    case 'list':
      text =
        (block.data as IEditorListData).items
          ?.map((item) => item.content)
          .join(' ') || '';
      break;
    case 'linkTool':
      text = (block.data as IEditorLinkToolData).meta?.title || '';
      break;
    default:
      text = '';
      break;
  }
  return text.replace(/<[^>]+>/g, '').trim(); // Remove HTML tags and trim whitespace
};

const getExcerptStr = (content: string, excerptLenght: number): string => {
  const contentObj = tryParseJSONObject(content);
  if (!contentObj)
    return (
      content
        .replace(/<[^>]+>/g, '')
        .trim()
        .slice(0, excerptLenght || DEFAULT_EXCERPT_LENGTH) + '...'
    );

  const { blocks } = contentObj;
  let text = '';
  for (let i = 0; i < blocks.length; i++) {
    const content = extractTextFromBlock(blocks[i]);
    if (!content) continue;
    console.log('Extracted text content: ', content);

    if (
      text.length + content.length <=
      (excerptLenght || DEFAULT_EXCERPT_LENGTH)
    ) {
      text += content + ' ';
    } else {
      const remainingLength =
        (excerptLenght || DEFAULT_EXCERPT_LENGTH) - text.length;
      const truncatedText = content.slice(0, remainingLength);
      text += `${truncatedText}...`;

      // If the block is not a paragraph, we just stop adding more blocks

      break;
    }
  }
  return text;
};

export { getExcerptStr };
